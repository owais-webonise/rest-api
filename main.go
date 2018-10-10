package main

import (
	"github.com/webonise/rest-api/app"
)

func main() {
	a := app.App{}
	a.Initialize("root", "root", "rest_api_example")
	a.Run(":8080")
}
