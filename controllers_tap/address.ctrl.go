package tapcontrollers

import (
	"net/http"
	"strconv"

	"github.com/bittap-protocol/lnhub/lib/responses"
	"github.com/bittap-protocol/lnhub/lib/service"
	"github.com/labstack/echo/v4"
)

// AddressController : AddressController struct
type AddressController struct {
	svc *service.LndhubService
}

func NewAddressController(svc *service.LndhubService) *AddressController {
	return &AddressController{svc: svc}
}

// get address request
type AddressRequestBody struct {
	AssetId string `json:"asset_id"`
	Amt     string `json:"amt"`
}

type AddressResponseBody struct {
	Address string `json:"address"`
}

// CreateAddress godoc
// @Summary      Get or create address
// @Description  Get or create address for deposit
// @Accept       json
// @Produce      json
// @Tags         Taproot Assets
// @Param        asset_id  body  string  true  "Asset ID"
// @Param        amt       body  string  true  "Amount"
// @Success      200      {object}  AddressResponseBody
// @Failure      400      {object}  responses.ErrorResponse
// @Failure      500      {object}  responses.ErrorResponse
// @Router       /tap/create-address [post]
// @Security     OAuth2Password
func (controller *AddressController) CreateAddress(c echo.Context) error {
	userId := c.Get("UserID").(int64)

	var body AddressRequestBody

	if err := c.Bind(&body); err != nil {
		c.Logger().Errorf("Failed to load tapcontrollers.CreateAddress request body: %v", err)
		return c.JSON(http.StatusBadRequest, responses.BadArgumentsError)
	}

	if err := c.Validate(&body); err != nil {
		c.Logger().Errorf("Invalid tapcontrollers.CreateAddress request body: %v", err)
		return c.JSON(http.StatusBadRequest, responses.BadArgumentsError)
	}
	// convert string amount to uint
	amt, err := strconv.ParseUint(body.Amt, 10, 64)
	if err != nil {
		c.Logger().Errorf("Invalid amount. Pass value as a string: %v", err)
		return c.JSON(http.StatusBadRequest, responses.BadArgumentsError)
	}
	result, err := controller.svc.FetchOrCreateAssetAddr(c.Request().Context(), uint64(userId), body.AssetId, amt)
	if err != nil {
		c.Logger().Errorf("error creating address: %v", err)
		return c.JSON(http.StatusBadRequest, responses.BadArgumentsError)
	}
	return c.JSON(http.StatusOK, &AddressResponseBody{
		Address: result,
	})
}
