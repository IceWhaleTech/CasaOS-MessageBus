package route

import "github.com/labstack/echo/v4"

func (r *APIRoute) SubscribeSIO(ctx echo.Context) error {
	server := r.services.SocketIOService.Server()
	server.ServeHTTP(ctx.Response(), ctx.Request())
	return nil
}

// unfortunately need to duplicate the func to support both `/socket.io` and `/socket.io/` (with a trailing slash) API endpoints
func (r *APIRoute) SubscribeSIO2(ctx echo.Context) error {
	return r.SubscribeSIO(ctx)
}

func (r *APIRoute) PollSIO(ctx echo.Context) error {
	server := r.services.SocketIOService.Server()
	server.ServeHTTP(ctx.Response(), ctx.Request())
	return nil
}

// unfortunately need to duplicate the func to support both `/socket.io` and `/socket.io/` (with a trailing slash) API endpoints
func (r *APIRoute) PollSIO2(ctx echo.Context) error {
	return r.PollSIO(ctx)
}
