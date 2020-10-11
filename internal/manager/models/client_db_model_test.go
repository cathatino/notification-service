package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	mockedClientId          int64  = 123
	mockedClientCategory    int64  = 321
	mockedClientKey         string = "client_key"
	mockedClientDescription string = "client_description"
	mockedCtime             uint32 = 1000
	mockedMtime             uint32 = 2000
	mockedClientModel              = &ClientModel{
		ClientId:       mockedClientId,
		ClientCategory: mockedClientCategory,
		ClientKey:      mockedClientKey,
		Description:    mockedClientDescription,
		Ctime:          mockedCtime,
		Mtime:          mockedMtime,
	}
)

func testEq(a, b []interface{}) bool {

	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestGetTableName(t *testing.T) {
	require.True(t, mockedClientModel.GetTableName() == clientTableName)
}

func TestGetColumns(t *testing.T) {
	withPrimaryKey := true
	modelColumns := mockedClientModel.GetColumns(!withPrimaryKey)
	modelColumnsWithPrimaryKey := mockedClientModel.GetColumns(withPrimaryKey)
	require.True(t, len(modelColumns) == len(clientTabColumns))
	require.True(t, len(modelColumnsWithPrimaryKey) == len(clientTabColumns)+1)
	for idx, col := range modelColumns {
		require.True(t, clientTabColumns[idx] == col)
	}
}

func TestGetValues(t *testing.T) {
	withPrimaryKey := true
	modelValues := mockedClientModel.GetValues(!withPrimaryKey)
	require.True(t, len(modelValues) == len(clientTabColumns))
	mockedValues := []interface{}{
		mockedClientCategory,
		mockedClientKey,
		mockedClientDescription,
		mockedCtime,
		mockedMtime,
	}
	require.True(t, testEq(modelValues, mockedValues))
}

func TestPrimaryKey(t *testing.T) {
	var mockePrimaryKeyValue int64 = 11234
	mockedClientModel.SetPrimaryKey(mockePrimaryKeyValue)
	_, primaryKeyValue := mockedClientModel.GetPrimaryKey()
	require.True(t, primaryKeyValue == mockePrimaryKeyValue)
}
