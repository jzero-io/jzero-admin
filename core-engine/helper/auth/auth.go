package auth

import (
	"context"
)

type Auth struct {
	Uuid      string   `json:"uuid"`
	Username  string   `json:"username"`
	RoleUuids []string `json:"role_uuids"`
}

func Info(ctx context.Context) (Auth, error) {
	var auth Auth
	if v, ok := ctx.Value("uuid").(string); ok {
		auth.Uuid = v
	}
	if v, ok := ctx.Value("username").(string); ok {
		auth.Username = v
	}
	if v, ok := ctx.Value("role_uuids").([]any); ok {
		roleIds := make([]string, len(v))
		for i, n := range v {
			if str, ok := n.(string); ok {
				roleIds[i] = str
			}
		}
		auth.RoleUuids = roleIds
	}
	return auth, nil
}
