package controller

import (
	"fmt"
	"net/http"

	"github.com/AmyangXYZ/SG_Sweetie/config"
	"github.com/AmyangXYZ/sweetygo"
)

// Upload API Handler.
func Upload(ctx *sweetygo.Context) error {
	saveDir := "uploadsfolder/"
	filename, err := ctx.SaveFile("file", config.RootDir+saveDir)
	fmt.Println(filename)
	if err != nil {
		return ctx.JSON(200, 0, "upload file error", nil)
	}
	filePath := "/" + saveDir + filename
	return ctx.JSON(200, 1, "success", filePath)
}

// Static files handler
func Static(ctx *sweetygo.Context) error {
	staticHandle := http.StripPrefix("/static",
		http.FileServer(http.Dir(config.RootDir+"/static")))
	staticHandle.ServeHTTP(ctx.Resp, ctx.Req)
	return nil
}

// Uploaded files handler
func Uploaded(ctx *sweetygo.Context) error {
	staticHandle := http.StripPrefix("/uploadsfolder",
		http.FileServer(http.Dir(config.RootDir+"/uploadsfolder")))
	staticHandle.ServeHTTP(ctx.Resp, ctx.Req)
	return nil
}
