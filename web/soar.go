package main

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	// app.Favicon("./assets/favicon.ico")

	// first parameter is the request path
	// second is the system directory
	//
	// app.HandleDir("/css", iris.Dir("./assets/css"))
	// app.HandleDir("/js",  iris.Dir("./assets/js"))
	// v1 := app.Party("/v1")
	indexFunc := func(ctx iris.Context) {
		ctx.Redirect("/static/index.html")
	}
	app.Get("/", indexFunc)
	app.Get("/static", indexFunc)

	routes := app.HandleDir("/static", iris.Dir("/usr/local/soar/web/static"), iris.DirOptions{
		Compress:   false,
		ShowList:   false,
		ShowHidden: false,
		Cache: iris.DirCacheOptions{
			Enable:          true,
			CompressIgnore:  iris.MatchImagesAssets,
			CompressMinSize: 300,
			Encodings:       []string{"gzip", "br" /* you can also add: deflate, snappy */},
		},
		DirList: iris.DirListRich(),
	})
	for _, r := range routes {
		fmt.Println(r)
	}

	// You can also register any index handler manually, order of registration does not matter:
	// v1.Get("/static", [...custom middleware...], func(ctx iris.Context) {
	//  [...custom code...]
	// 	ctx.ServeFile("./assets/index.html")
	// })

	// http://localhost:8080/v1/static
	// http://localhost:8080/v1/static/css/main.css
	// http://localhost:8080/v1/static/js/jquery-2.1.1.js
	// http://localhost:8080/v1/static/favicon.ico

	app.Listen(":8080")
}
