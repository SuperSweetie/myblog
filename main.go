package main

import (
	"html/template"
	"strings"

	"github.com/AmyangXYZ/SG_Sweetie/config"
	"github.com/AmyangXYZ/SG_Sweetie/router"
	"github.com/AmyangXYZ/sweetygo"
)

// Host is for subdomain control
type Host struct {
	SG *sweetygo.SweetyGo
}

func main() {
	blog := sweetygo.New()
	blog.SetTemplates(config.RootDir+"templates", template.FuncMap{
		"unescaped":    unescaped,
		"space2hyphen": space2hyphen,
		"abstract":     abstract,
		"rmtag":        rmtag,
	})
	router.SetMiddlewares(blog)
	router.SetRouter(blog)
	blog.Run(":8888")
}

func unescaped(s string) interface{} {
	return template.HTML(s)
}

// for title in url, Hello World -> Hello-World
func space2hyphen(s string) string {
	return strings.Replace(s, " ", "-", -1)
}

// show abstract, splited by tag icon.
func abstract(s string) string {
	return strings.Split(s, "<p><i class=\"fa fa-tag fa-emoji\" title=\"tag\"></i></p>")[0]
}

// replace tag icon in content
func rmtag(s string) string {
	return strings.Replace(s, "<p><i class=\"fa fa-tag fa-emoji\" title=\"tag\"></i></p>", "", -1)
}
