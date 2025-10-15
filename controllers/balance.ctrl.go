package controllers

import (
	"net/http"

	"github.com/bittap-protocol/lnhub/common"
	"github.com/bittap-protocol/lnhub/lib/responses"
	"github.com/bittap-protocol/lnhub/lib/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// BalanceController : BalanceController struct
type BalanceController struct {
	svc *service.LndhubService
}

func NewBalanceController(svc *service.LndhubService) *BalanceController {
	return &BalanceController{svc: svc}
}

type BalanceResponse struct {
	BTC struct {
		AvailableBalance int64
	}
}

func (controller *BalanceController) Balance(c echo.Context) error {
	userId := c.Get("UserID").(int64)
	assetParam := c.Param("asset_id")
	// default to bitcoin if error parsing the param
	if assetParam == "" {
		assetParam = common.BTC_ASSET_ID
	}

	balance, err := controller.svc.CurrentUserBalance(c.Request().Context(), assetParam, userId)
	if err != nil {
		c.Logger().Errorj(
			log.JSON{
				"message":        "failed to retrieve user balance",
				"lndhub_user_id": userId,
				"error":          err,
			},
		)
		return c.JSON(http.StatusBadRequest, responses.BadArgumentsError)
	}
	return c.JSON(http.StatusOK, &BalanceResponse{
		BTC: struct{ AvailableBalance int64 }{
			AvailableBalance: balance,
		},
	})
}
