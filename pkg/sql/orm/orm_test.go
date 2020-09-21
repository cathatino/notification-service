package orm

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/cathatino/notification-service/pkg/sql/pg"
	"github.com/stretchr/testify/require"
)

const (
	mockUserTableName = "mocked_user_tab"
)

var (
	config *pg.Config
	// mocking var
	mockUserTabColNameUserId   = "user_id"
	mockUserTabColNameUserName = "user_name"
	mockUserTabColNameCtime    = "ctime"
	mockUserTabColumns         = []string{
		mockUserTabColNameUserName,
		mockUserTabColNameCtime,
	}
)

type MockUserModel struct {
	UserId   int64  `db:"user_id"`
	UserName string `db:"user_name"`
	Ctime    uint32 `db:"ctime"`
}

func (m *MockUserModel) GetTableName() string {
	return mockUserTableName
}

func (m *MockUserModel) GetColumns(withPrimaryKey bool) []string {
	if withPrimaryKey {
		return append([]string{mockUserTabColNameUserId}, mockUserTabColumns...)
	}
	return mockUserTabColumns
}

func (m *MockUserModel) GetValues(withPrimaryKey bool) []interface{} {
	values := []interface{}{
		m.UserName,
		m.Ctime,
	}
	if withPrimaryKey {
		return append([]interface{}{m.UserId}, values...)
	}
	return values
}

func (m *MockUserModel) SetPrimaryKey(primaryKey int64) {
	m.UserId = primaryKey
}

func (m *MockUserModel) GetPrimaryKey() (string, int64) {
	return mockUserTabColNameUserId, m.UserId
}

func init() {
	// get env variables
	env := func(key, defaultValue string) string {
		if value := os.Getenv(key); value != "" {
			return value
		}
		return defaultValue
	}

	// init config
	config = &pg.Config{
		Host:            env("PSQL_TEST_HOST", "host"),
		Port:            env("PSQL_TEST_PORT", "5432"),
		User:            env("PSQL_TEST_USER", "user"),
		Password:        env("PSQL_TEST_PASSWORD", "password"),
		Dbname:          env("PSQL_TEST_DBNAME", "dbname"),
		MaxOpenConns:    10,
		MaxIdleConns:    10,
		ConnMaxLifetime: time.Hour,
	}
}

func fetchNewOrm(t *testing.T) ORM {
	conn, err := pg.New(config)
	if err != nil {
		t.Fatal(err)
	}
	return NewOrm(conn)
}

func TestOrmFind(t *testing.T) {
	orm := fetchNewOrm(t)

	ctx := context.Background()
	users := make([]MockUserModel, 0)

	err := orm.Find(ctx, &users, squirrel.Eq{"user_id": "1"})
	if err != nil {
		t.Fatal(err)
	}
	require.True(t, len(users) == 1)
}

func TestOrmCreate(t *testing.T) {
	orm := fetchNewOrm(t)

	ctx := context.Background()
	user := MockUserModel{
		UserName: fmt.Sprintf("user_name_%d", time.Now().Unix()),
		Ctime:    uint32(time.Now().Unix()),
	}

	err := orm.Create(ctx, &user)
	if err != nil {
		t.Fatal(err)
	}

	require.True(t, user.UserId > 0)
}
