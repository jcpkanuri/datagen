package appconfig

import (
	"datagen/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConf(t *testing.T) {
	got := GetConf()
	assert.NotNil(t, got)
	assert.IsType(t, &got.Conns, &[]types.DbConn{})
}
