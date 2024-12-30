package middleware

import (
	"net/http"

	casbin "github.com/casbin/casbin/v2"
	"github.com/spf13/cast"

	"github.com/jzero-io/jzero-admin/server/internal/auth"
)

type AuthxMiddleware struct {
	CasbinEnforcer *casbin.Enforcer
	Route2CodeFunc func(r *http.Request) string
}

func NewAuthxMiddleware(casbinEnforcer *casbin.Enforcer, route2codeFunc func(r *http.Request) string) *AuthxMiddleware {
	return &AuthxMiddleware{
		CasbinEnforcer: casbinEnforcer,
		Route2CodeFunc: route2codeFunc,
	}
}

func (m *AuthxMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authInfo, err := auth.Info(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		rctx := r.Context()

		subs := cast.ToStringSlice(authInfo.RoleIds)
		obj := m.Route2CodeFunc(r)

		// verify casbin rule
		if result := batchCheck(m.CasbinEnforcer, subs, obj); !result {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		next(w, r.WithContext(rctx))
	}
}

func batchCheck(cbn *casbin.Enforcer, subs []string, obj string) bool {
	var checkReq [][]any
	for _, v := range subs {
		checkReq = append(checkReq, []any{v, obj})
	}

	result, err := cbn.BatchEnforce(checkReq)
	if err != nil {
		return false
	}

	for _, v := range result {
		if v {
			return true
		}
	}

	return false
}
