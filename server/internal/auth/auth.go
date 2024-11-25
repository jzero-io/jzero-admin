package auth

import (
	"context"
	"encoding/json"
)

type Auth struct {
	Id       int     `json:"id"`
	Username string  `json:"username"`
	RoleIds  []int64 `json:"role_ids"`
}

func Info(ctx context.Context) (Auth, error) {
	var auth Auth
	if v, ok := ctx.Value("id").(json.Number); ok {
		id, _ := v.Int64()
		auth.Id = int(id)
	}
	if v, ok := ctx.Value("username").(string); ok {
		auth.Username = v
	}
	if v, ok := ctx.Value("role_ids").([]any); ok {
		roleIds := make([]int64, len(v))
		for i, n := range v {
			roleIds[i], _ = n.(json.Number).Int64()
		}
		auth.RoleIds = roleIds
	}
	return auth, nil
}
