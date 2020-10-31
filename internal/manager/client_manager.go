package manager

import (
	"context"

	"github.com/Masterminds/squirrel"

	"github.com/cathatino/notification-service/internal/manager/models"
	"github.com/cathatino/notification-service/pkg/cache/redis"
	"github.com/cathatino/notification-service/pkg/sql/connector"
	"github.com/cathatino/notification-service/pkg/sql/orm"
)

type ClientManager interface {
	GetClientByClientId(ctx context.Context, clientId int64) (*models.ClientModel, error)
}

type clientManager struct {
	ORM       orm.ORM
	RedisPool *redis.Pool // TODO: implement cache layer orm
}

func NewClientManager(
	con connector.Connector,
	redisPool *redis.Pool,
) ClientManager {
	return &clientManager{
		ORM:       orm.NewOrm(con),
		RedisPool: redisPool,
	}
}

func (cm *clientManager) GetClientByClientId(ctx context.Context, clientId int64) (
	*models.ClientModel,
	error,
) {
	clients := make([]models.ClientModel, 0)
	cm.RedisPool.
	if err := cm.ORM.Find(ctx, &clients, squirrel.Eq{"client_id": clientId}); err != nil {
		return nil, err
	}
	if length := len(clients); length == 0 {
		return nil, ErrRecordNotFound
	} else if length > 1 {
		return nil, ErrUnexpectedLengthFromDb
	}
	clientModelObj := &clients[0]
	return clientModelObj, nil
}
