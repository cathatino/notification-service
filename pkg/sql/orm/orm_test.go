package orm

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/Masterminds/squirrel"

	"github.com/cathatino/notification-service/pkg/sql/pg"
)

const (
	mockUserTableName = "mocked_user_tab"
)

var (
	config *pg.Config
	// mocking var
	mockUserTabColNameId       = "id"
	mockUserTabColNameUserName = "user_name"
	mockUserTabColNameCtime    = "ctime"
	mockUserTabColumns         = []string{
		mockUserTabColNameId,
		mockUserTabColNameUserName,
		mockUserTabColNameCtime,
	}
)

type MockUserModel struct {
	Id       int64  `db:"id"`
	UserName string `db:"user_name"`
	Ctime    uint32 `db:"ctime"`
}

func (m *MockUserModel) GetTableName() string {
	return mockUserTableName
}

func (m *MockUserModel) GetColumns() []string {
	return mockUserTabColumns
}

func (m *MockUserModel) GetValues() []interface{} {
	return []interface{}{
		m.Id,
		m.UserName,
		m.Ctime,
	}
}

func (m *MockUserModel) SetPrimaryKey(primaryKey int64) {
	m.Id = primaryKey
}

func (m *MockUserModel) GetPrimaryKey() (string, int64) {
	return mockUserTabColNameId, m.Id
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

	err := orm.Find(ctx, &users, squirrel.Eq{"id": 1})
	if err != nil {
		t.Fatal(err)
	}
}
