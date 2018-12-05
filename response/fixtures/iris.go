package fixtures

import (
	"net/http"

	"github.com/kataras/iris"
	"github.com/rohmanhm/brover/response"
)

// SetupIrisServer fixtures
func SetupIrisServer() *iris.Application {
	app := iris.New()
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		r := response.NewIris(ctx)
		r.Errors(&response.Errors{
			Message: "route not found",
		}).StatusCode(http.StatusNotFound)

		r.Render()
	})
	app.Get(RouteText(RouteTestResponse), func(ctx iris.Context) {
		r := response.NewIris(ctx)
		r.Data([]string{
			"foo",
			"bar",
		})

		r.Render()
	})
	app.Get(RouteText(RouteTestResponseError), func(ctx iris.Context) {
		r := response.NewIris(ctx)
		r.Errors(&response.Errors{
			Message: "error message",
		})

		r.Render()
	})

	return app
}
