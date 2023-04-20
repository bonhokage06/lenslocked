package context

import (
	"context"

	"github.com/bonhokage06/lenslocked/models"
)

type key string

const (
	userKey key = "user"
)

func WithUser(ctx context.Context, userSession *models.UserSession) context.Context {
	return context.WithValue(ctx, userKey, userSession)
}
func User(ctx context.Context) *models.UserSession {
	val := ctx.Value(userKey)
	user, ok := val.(*models.UserSession)
	if !ok {
		return nil
	}
	return user
}
