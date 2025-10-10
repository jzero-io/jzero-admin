package errcodes

import (
	"github.com/jzero-io/jzero-admin/server/internal/errcodes/auth"
	"github.com/jzero-io/jzero-admin/server/internal/errcodes/manage"
)

func Register() {
	manage.RegisterManage()
	auth.RegisterAuth()
}
