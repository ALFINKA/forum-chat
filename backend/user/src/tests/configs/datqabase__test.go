package configs

import (
	"fmt"
	"testing"
	"user/src/configs"
)

func TestDatabaseConnect(t *testing.T) {
	fmt.Println(configs.ConnectDB())
}
