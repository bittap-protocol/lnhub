package tapcontrollers

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
	Balance int64  `json:"balance"`
	AssetId string `json:"asset_id"`
}

// / get all balances
type BalancesResponse struct {
	Balances map[string]int64 `json:"balances"`
}

// Balance godoc
// @Summary      Retrieve balance
// @Description  Current user's balance in satoshi
// @Accept       json
// @Produce      json
// @Tags         Taproot Assets
// @Success      200  {object}  BalanceResponse
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /tap/balance/:asset_id [get]
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
		Balance: balance,
		AssetId: assetParam,
	})
}

// Balances godoc
// @Summary      Retrieve all balances
// @Description  Retrieve all user balances
// @Accept       json
// @Produce      json
// @Tags         Taproot Assets
// @Success      200  {object}  BalancesResponse
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /tap/balances/all [get]
func (controller *BalanceController) Balances(c echo.Context) error {
	userId := c.Get("UserID").(int64)

	balances, err := controller.svc.GetAllCurrentBalancesJson(c.Request().Context(), userId)
	if err != nil {
		c.Logger().Errorj(
			log.JSON{
				"message":        "failed to retrieve user balances",
				"lndhub_user_id": userId,
				"error":          err,
			},
		)
		return c.JSON(http.StatusBadRequest, responses.BadArgumentsError)
	}
	return c.JSON(http.StatusOK, &BalancesResponse{
		Balances: balances,
	})
}
