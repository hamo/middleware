package recovery

import (
	"runtime/debug"

	"gopkg.in/kataras/iris.v5"
)

var Handler = iris.HandlerFunc(func(ctx *iris.Context) {
	defer func() {
		if err := recover(); err != nil {

			ctx.Log("Recovery from panic: %s\n%s", err, debug.Stack())

			//ctx.Panic just sends  http status 500 by default, but you can change it by: iris.OnPanic(func( c *iris.Context){})
			ctx.Panic()
		}
	}()
	ctx.Next()
})

// New restores the server on internal server errors (panics)
// returns the middleware
//
// is here for compatiblity
func New() iris.HandlerFunc {
	return Handler
}
