package conn

import (
	"fmt"
	"testing"
)

func TestConnect(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			return
		}
	}()
	Connect()
}
