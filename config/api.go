package config

import (
	"fmt"
	"os"
)

func GetAPIAddress() string {
	return fmt.Sprintf(":%s", os.Getenv("PORT"))
}
