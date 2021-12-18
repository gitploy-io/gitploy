//go:generate mockgen -source ./interface.go -destination ./mock/interactor.go -package mock

package slack

import (
	"context"

	"github.com/gitploy-io/gitploy/model/ent"
)

type (
	Interactor interface {
		FindUserByID(ctx context.Context, id int64) (*ent.User, error)

		FindChatUserByID(ctx context.Context, id string) (*ent.ChatUser, error)
		CreateChatUser(ctx context.Context, cu *ent.ChatUser) (*ent.ChatUser, error)
		UpdateChatUser(ctx context.Context, cu *ent.ChatUser) (*ent.ChatUser, error)
		DeleteChatUser(ctx context.Context, cu *ent.ChatUser) error

		SubscribeEvent(fn func(e *ent.Event)) error
		UnsubscribeEvent(fn func(e *ent.Event)) error
	}
)
