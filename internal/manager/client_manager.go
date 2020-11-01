package manager

import (
	"fmt"
	"bytes"
	"context"
	"encoding/gob"

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
	ORM            orm.ORM
	RedisConnector *redis.Connector
}

func NewClientManager(
	con connector.Connector,
	connector *redis.Connector,
) ClientManager {
	return &clientManager{
		ORM:            orm.NewOrm(con),
		RedisConnector: connector,
	}
}

func (cm *clientManager) GetClientByClientId(ctx context.Context, clientId int64) (
	*models.ClientModel,
	error,
) {
	// Read from cache
	cacheKey := fmt.Sprintf(GetUserClientByIdCacheFmt, clientId)
	if bytesSlice, _ := cm.RedisConnector.Get(cacheKey); len(bytesSlice) > 0 {
		buffer := bytes.NewBuffer(bytesSlice)
		decoder := gob.NewDecoder(buffer)
		var clientModelObj models.ClientModel
		err := decoder.Decode(&clientModelObj)
		return &clientModelObj, err
	}

	// Read from db
	clients := make([]models.ClientModel, 0)
	if err := cm.ORM.Find(ctx, &clients, squirrel.Eq{"client_id": clientId}); err != nil {
		return nil, err
	}
	if length := len(clients); length == 0 {
		return nil, ErrRecordNotFound
	} else if length > 1 {
		return nil, ErrUnexpectedLengthFromDb
	}
	clientModelObj := &clients[0]

	// Set to Cache
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
    encoder.Encode(clientModelObj)
	bytesSlice := buffer.Bytes()
	cm.RedisConnector.Set(cacheKey, bytesSlice)

	return clientModelObj, nil
}
