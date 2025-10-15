package transport

import (
	tapcontrollers "github.com/bittap-protocol/lnhub/controllers_tap"
	"github.com/bittap-protocol/lnhub/lib/service"
	"github.com/labstack/echo/v4"
)

func RegisterTapEndpoints(svc *service.LndhubService, e *echo.Echo, secured *echo.Group, securedWithStrictRateLimit *echo.Group, strictRateLimitMiddleware echo.MiddlewareFunc, adminMw echo.MiddlewareFunc, logMw echo.MiddlewareFunc) {
	e.GET("/tap/universe-assets", tapcontrollers.NewUniverseController(svc).UniverseAssets, strictRateLimitMiddleware, logMw)
	secured.GET("/tap/balances/all", tapcontrollers.NewBalanceController(svc).Balances, strictRateLimitMiddleware, logMw)
	secured.POST("/tap/create-address", tapcontrollers.NewAddressController(svc).CreateAddress, strictRateLimitMiddleware, logMw)
	secured.POST("/tap/transfer", tapcontrollers.NewTransferController(svc).Transfer, strictRateLimitMiddleware, logMw)
	secured.GET("/tap/balance/:asset_id", tapcontrollers.NewBalanceController(svc).Balance)
}
