# notification-service
A notification service

## Skeleton
```text
├── LICENSE
├── README.md
├── cmd
├── go.mod
├── go.sum
├── internal
│   └── manager
│       └── db
│           └── driver.go
└── pkg
    └── pg
        ├── config.go
        ├── config_test.go
        ├── connector.go
        └── connector_test.go
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
