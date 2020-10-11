package models

const (
	clientTableName                = "client_tab"
	clientTabColNameClientId       = "client_id"
	clientTabColNameClientCategory = "client_category"
	clientTabColNameClientKey      = "client_key"
	clientTabColNameDescription    = "description"
	clientTabColNameCtime          = "ctime"
	clientTabColNameMtime          = "mtime"
)

var (
	clientTabColumns = []string{
		clientTabColNameClientCategory,
		clientTabColNameClientKey,
		clientTabColNameDescription,
		clientTabColNameCtime,
		clientTabColNameMtime,
	}
)

type ClientModel struct {
	ClientId       int64  `db:"client_id"`
	ClientCategory int64  `db:"client_category"`
	ClientKey      string `db:"client_key"`
	Description    string `db:"description"`
	Ctime          uint32 `db:"ctime"`
	Mtime          uint32 `db:"mtime"`
}

func (c *ClientModel) GetTableName() string {
	return clientTableName
}

func (c *ClientModel) GetColumns(withPrimaryKey bool) []string {
	if withPrimaryKey {
		return append([]string{clientTabColNameClientId}, clientTabColumns...)
	}
	return clientTabColumns
}

func (c *ClientModel) GetValues(withPrimaryKey bool) []interface{} {
	values := []interface{}{
		c.ClientCategory,
		c.ClientKey,
		c.Description,
		c.Ctime,
		c.Mtime,
	}
	if withPrimaryKey {
		return append([]interface{}{c.ClientId}, values...)
	}
	return values
}

func (c *ClientModel) SetPrimaryKey(primaryKey int64) {
	c.ClientId = primaryKey
}

func (m *ClientModel) GetPrimaryKey() (string, int64) {
	return clientTabColNameClientId, m.ClientId
}
