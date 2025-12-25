package custom

import (
	"context"
	"encoding/json"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/jzero-io/jzero/core/stores/migrate"

	"github.com/jzero-io/jzero-admin/server/internal/global"
	"github.com/jzero-io/jzero-admin/server/internal/model"
	menutypes "github.com/jzero-io/jzero-admin/server/internal/types/v1/manage/menu"
)

type Custom struct{}

func New() *Custom {
	return &Custom{}
}

// Init Please add custom logic here.
func (c *Custom) Init() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	m, err := migrate.NewMigrate(global.ServiceContext.ConfigCenter.MustGetConfig().Sqlx.SqlConf, migrate.WithSourceAppendDriver(true))
	if err != nil {
		return err
	}

	if err = m.Up(); err != nil {
		return err
	}
	defer m.Close()

	// auto gen casbin rules
	if err = InitCasbinRule(ctx, global.ServiceContext.Model, global.ServiceContext.CasbinEnforcer); err != nil {
		return err
	}

	return nil
}

// Start Please add custom logic here.
func (c *Custom) Start() {}

// Stop Please add shut down logic here.
func (c *Custom) Stop() {}

func InitCasbinRule(ctx context.Context, model model.Model, enforcer *casbin.Enforcer) error {
	// get all role
	allRoles, err := model.ManageRole.FindByCondition(ctx, nil)
	if err != nil {
		return err
	}

	allMenus, err := model.ManageMenu.FindByCondition(ctx, nil)
	if err != nil {
		return err
	}

	// get role menu
	allRoleMenus, err := model.ManageRoleMenu.FindByCondition(ctx, nil)
	if err != nil {
		return err
	}

	var casbinRules [][]string
	for _, v := range allRoles {
		for _, arm := range allRoleMenus {
			if v.Uuid == arm.RoleUuid {
				for _, am := range allMenus {
					if arm.MenuUuid == am.Uuid {
						var permissions []menutypes.Permission
						err = json.Unmarshal([]byte(am.Permissions), &permissions)
						if err == nil {
							for _, perm := range permissions {
								if perm.Code != "" {
									if hasPolicy, _ := enforcer.HasPolicy(v.Uuid, perm.Code); !hasPolicy {
										casbinRules = append(casbinRules, []string{v.Uuid, perm.Code})
									}
								}
							}
						}
					}
				}
			}
		}
	}

	if len(casbinRules) > 0 {
		_, err = enforcer.AddPolicies(casbinRules)
		if err != nil {
			return err
		}
	}
	return nil
}
