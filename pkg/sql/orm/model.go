package orm

type Model interface {
	GetTableName() string
	GetColumns(withPrimaryKey bool) []string
	GetValues(withPrimaryKey bool) []interface{}
	SetPrimaryKey(int64)
	GetPrimaryKey() (column string, value int64)
}
