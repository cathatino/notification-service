# notification-service
A notification service

## Skeleton
```text
├── LICENSE
├── README.md
├── go.mod
├── go.sum
├── internal
│   └── manager
│       ├── client_manager.go
│       ├── client_manager_test.go
│       ├── constants.go
│       ├── errors.go
│       └── models
│           ├── client_db_model.go
│           └── client_db_model_test.go
├── main.go
├── pkg
│   ├── cache
│   │   └── redis
│   │       ├── client.go
│   │       └── client_test.go
│   ├── sql
│   │   ├── connector
│   │   │   └── connector.go
│   │   ├── orm
│   │   │   ├── errors.go
│   │   │   ├── model.go
│   │   │   ├── orm.go
│   │   │   └── orm_test.go
│   │   └── pg
│   │       ├── config.go
│   │       ├── config_test.go
│   │       ├── connector.go
│   │       └── connector_test.go
│   └── utils
│       └── reflectutil
│           ├── reflect.go
│           └── reflect_test.go
└── script
    └── sql
        └── mocked_db.sql
```

## Testing

### Local Testing
use the cmd to run test under local
```bash
go test ./...
```

## Reference
[sqlx](https://github.com/jmoiron/sqlx)
[gvm](https://github.com/moovweb/gvm)
[gob](https://golang.org/pkg/encoding/gob/)

