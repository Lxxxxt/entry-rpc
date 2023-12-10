package modules

import "testing"

func TestConn(t *testing.T) {
	err := Startup()
	if err != nil {
		panic(err)
	}
}
