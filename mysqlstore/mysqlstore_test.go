package mysqlstore

import (
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/tester"
)

const mysqlTestServer = "localhost:33006"

var newStore = func(_ *testing.T) sessions.Store {
	store, err := NewStore(mysqlTestServer, "test", "/", 3600, []byte("secret"))
	if err != nil {
		panic(err)
	}
	return store
}

func TestMySQL_SessionGetSet(t *testing.T) {
	tester.GetSet(t, newStore)
}

func TestMySQL_SessionDeleteKey(t *testing.T) {
	tester.DeleteKey(t, newStore)
}

func TestMySQL_SessionFlashes(t *testing.T) {
	tester.Flashes(t, newStore)
}

func TestMySQL_SessionClear(t *testing.T) {
	tester.Clear(t, newStore)
}

func TestMySQL_SessionOptions(t *testing.T) {
	tester.Options(t, newStore)
}

func TestMySQL_SessionMany(t *testing.T) {
	tester.Many(t, newStore)
}
