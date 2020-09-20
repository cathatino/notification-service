package orm

type Model interface {
	GetTableName() string
	GetColumns() []string
	GetValues() []interface{}
	SetPrimaryKey(int64)
	GetPrimaryKey(int64)
}
