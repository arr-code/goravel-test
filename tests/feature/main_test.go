package feature

import (
	"os"
	"testing"

	"github.com/goravel/framework/facades"
)

func TestMain(m *testing.M) {
	database, err := facades.Testing().Docker().Database("mysql")
	if err != nil {
		panic(err)
	}
	if err := database.Build(); err != nil {
		panic(err)
	}

	go func() {
		if err := facades.Route().Run(); err != nil {
			panic(err)
		}
	}()

	exit := m.Run()

	if err := database.Clear(); err != nil {
		panic(err)
	}

	os.Exit(exit)

}
