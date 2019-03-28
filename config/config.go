package config

import (
	"fmt"
	"os"
)

var (
	// RootDir of your app
	RootDir = "/root/go/src/github.com/AmyangXYZ/SG_Sweetie/"

	// SecretKey computes sg_token
	SecretKey string

	// DB addr and passwd.
	DB string
)

func init() {
	DB = fmt.Sprintf("root:%v@tcp(127.0.0.1:3306)/sg_sweetie?charset=utf8", "123!")
	SecretKey = os.Getenv("SecretKey")
}
