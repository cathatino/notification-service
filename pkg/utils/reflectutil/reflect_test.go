package reflectutil

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsPtr(t *testing.T) {
	userid := 1
	require.False(t, IsPtr(userid))
	require.True(t, IsPtr(&userid))
}
