package auth

import (
	"context"
	"encoding/json"
)

type Auth struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
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
	return auth, nil
}
