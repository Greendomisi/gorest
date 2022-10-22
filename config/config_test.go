package config_test

import (
	"github.com/Greendomisi/gorest/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	_, err := config.New("config.json")
	assert.NoError(t, err)
	_, err = config.New("config.jsox")
	assert.Error(t, err)
}
