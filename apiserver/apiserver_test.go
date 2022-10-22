package apiserver_test

import (
	"testing"

	"github.com/Greendomisi/gorest/apiserver"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestOptions(t *testing.T) {
	t.Parallel()
	t.Run("Options tests", func(t *testing.T) {
		s := apiserver.New(apiserver.AdressOpt("127.0.0.1:8080"), apiserver.RouterOpt(mux.NewRouter()))

		assert.Equal(t, s.Adress, "127.0.0.1:8080")
		s.Adress = "127.0.0.1:8081"
		assert.NotEqual(t, s.Adress, "127.0.0.1:8080")
	})
}
