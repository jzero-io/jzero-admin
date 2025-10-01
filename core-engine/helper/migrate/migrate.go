package migrate

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/eddieowens/opts"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type MigrateUpOpts struct {
	BeforeMigrateUpFuncs map[uint]func(version uint) error
	AfterMigrateUpFuncs  map[uint]func(version uint) error
	PluginName           string
}

func (opts MigrateUpOpts) DefaultOptions() MigrateUpOpts {
	return MigrateUpOpts{}
}

func WithBeforeMigrateUpFunc(mapFuncs map[uint]func(uint) error) opts.Opt[MigrateUpOpts] {
	return func(opts *MigrateUpOpts) {
		opts.BeforeMigrateUpFuncs = mapFuncs
	}
}

func WithAfterMigrateUpFunc(mapFunc map[uint]func(uint) error) opts.Opt[MigrateUpOpts] {
	return func(opts *MigrateUpOpts) {
		opts.AfterMigrateUpFuncs = mapFunc
	}
}

func WithPluginName(pluginName string) opts.Opt[MigrateUpOpts] {
	return func(opts *MigrateUpOpts) {
		opts.PluginName = pluginName
	}
}

func MigrateUp(ctx context.Context, c sqlx.SqlConf, op ...opts.Opt[MigrateUpOpts]) error {
	ops := opts.DefaultApply(op...)
	var (
		dataSource     = c.DataSource
		source         = filepath.Join("desc", "sql_migration")
		paramConnector string
	)

	if strings.Contains(dataSource, "?") {
		paramConnector = "&"
	} else {
		paramConnector = "?"
	}

	switch c.DriverName {
	case "mysql":
		dataSource = "mysql://" + c.DataSource
		if ops.PluginName != "" {
			source = filepath.Join("plugins", ops.PluginName, "desc", "sql_migration")
			dataSource = fmt.Sprintf("%s%sx-migrations-table=%s", dataSource, paramConnector, "schema_migrations_plugin_"+ops.PluginName)
		}
	case "pgx":
		dataSource = "pgx5://" + strings.TrimPrefix(c.DataSource, "postgres://")
		source = filepath.Join(source, "postgresql")
		if ops.PluginName != "" {
			source = filepath.Join("plugins", ops.PluginName, "desc", "sql_migration", "postgresql")
			dataSource = fmt.Sprintf("%s%sx-migrations-table=%s", dataSource, paramConnector, "schema_migrations_plugin_"+ops.PluginName)
		}
	}

	if err := migrateUp(ctx, source, dataSource, c, ops); err != nil {
		return err
	}
	return nil
}

type customFileSource struct {
	*file.File
	sqlConf sqlx.SqlConf
	ctx     context.Context
}

func (c *customFileSource) ReadUp(version uint) (r io.ReadCloser, identifier string, err error) {
	rc, id, err := c.File.ReadUp(version)
	if err != nil {
		return nil, "", err
	}

	content, err := io.ReadAll(rc)
	if err != nil {
		return nil, "", err
	}

	if err = rc.Close(); err != nil {
		return nil, "", err
	}

	return io.NopCloser(strings.NewReader(c.preprocessSQL(string(content)))), id, nil
}

func (c *customFileSource) ReadDown(version uint) (r io.ReadCloser, identifier string, err error) {
	rc, id, err := c.File.ReadDown(version)
	if err != nil {
		return nil, "", err
	}

	content, err := io.ReadAll(rc)
	if err != nil {
		return nil, "", err
	}

	if err = rc.Close(); err != nil {
		return nil, "", err
	}

	return io.NopCloser(strings.NewReader(c.preprocessSQL(string(content)))), id, nil
}

func (c *customFileSource) preprocessSQL(content string) string {
	if c.sqlConf.DriverName == "pgx" {
		// Match INSERT INTO ... VALUES ... patterns and add ON CONFLICT DO NOTHING
		insertRegex := regexp.MustCompile(`(?i)(INSERT\s+INTO\s+[^;]+VALUES[^;]*);`)
		content = insertRegex.ReplaceAllString(content, "$1 ON CONFLICT DO NOTHING;")
	} else {
		// For MySQL, replace INSERT INTO with INSERT IGNORE INTO
		insertRegex := regexp.MustCompile(`(?i)INSERT\s+INTO`)
		content = insertRegex.ReplaceAllString(content, "INSERT IGNORE INTO")
	}
	return content
}

func migrateUp(ctx context.Context, source, dataSource string, c sqlx.SqlConf, ops MigrateUpOpts) error {
	fileDriver := &file.File{}
	if !strings.HasPrefix(source, "file://") {
		source = "file://" + source
	}
	fileSource, err := fileDriver.Open(source)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil
		}
		return err
	}

	customSource := &customFileSource{
		File:    fileSource.(*file.File),
		sqlConf: c,
		ctx:     ctx,
	}

	m, err := migrate.NewWithSourceInstance("file", customSource, dataSource)
	if err != nil {
		return err
	}
	defer m.Close()

	if ops.BeforeMigrateUpFuncs == nil && ops.AfterMigrateUpFuncs == nil {
		err = m.Up()
		if errors.Is(err, fs.ErrNotExist) || errors.Is(err, migrate.ErrNilVersion) || errors.Is(err, migrate.ErrNoChange) {
			return nil
		}
		return err
	}
	// 获取当前版本
	currentVersion, _, err := m.Version()
	if err != nil {
		if errors.Is(err, migrate.ErrNilVersion) {
			// 不存在的话, 直接返回 Up
			if err = m.Up(); err != nil {
				if errors.Is(err, fs.ErrNotExist) || errors.Is(err, migrate.ErrNilVersion) || errors.Is(err, migrate.ErrNoChange) {
					return nil
				}
			}
			return err
		}
	}

	for {
		nextVersion, err := customSource.Next(currentVersion)
		if err == nil && nextVersion > currentVersion {
			if f, ok := ops.BeforeMigrateUpFuncs[nextVersion]; ok {
				if err = f(nextVersion); err != nil {
					return err
				}
			}
			if err = m.Steps(1); err != nil {
				return err
			}
			if f, ok := ops.AfterMigrateUpFuncs[nextVersion]; ok {
				if err = f(nextVersion); err != nil {
					if stepDownErr := m.Steps(-1); stepDownErr != nil {
						return stepDownErr
					}
					return err
				}
			}
			currentVersion = nextVersion
		} else {
			break
		}
	}

	return nil
}
