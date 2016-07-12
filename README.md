## Repository information

This repository contains the restful types context's rendering for [Iris web framework](https://github.com/kataras/iris) and [fasthttp](https://github.com/valyala/fasthttp).
The idea of this came from [unrolled/render](https://github.com/unrolled/render), which I had the time to convert it for fasthttp & iris & improve even more its performance and usability.



## Quick look

```go
// small example for json rendering, same for all other

package main

import (
	"github.com/kataras/iris"
)

func main() {
	iris.Get("/hi_json", func(ctx *iris.Context) {
		ctx.JSON(iris.StatusOK, iris.Map{
			"Name": "Iris",
			"Age":  2,
		}) // ctx.XML, ctx.Data, ctx.Text, ctx.JSONP...
	})

	iris.Listen(":8080")
}

```

## How to use

- Docs [here](https://kataras.gitbooks.io/iris/content/render_rest.html)
- Examples [here](https://github.com/iris-contrib/examples/tree/dev)

## License

This project is licensed under the MIT License.

License can be found [here](LICENSE).
