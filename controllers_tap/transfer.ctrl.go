package tapcontrollers

import (
	"net/http"

	"github.com/bittap-protocol/lnhub/lib/responses"
	"github.com/bittap-protocol/lnhub/lib/service"
	"github.com/labstack/echo/v4"
)

// TransferController : TransferController struct
type TransferController struct {
	svc *service.LndhubService
}

func NewTransferController(svc *service.LndhubService) *TransferController {
	return &TransferController{svc: svc}
}

// transfer request
type TransferRequestBody struct {
	Address string `json:"address"`
}

// transfer response
type TransferResponseBody struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

// Transfer godoc
// @Summary      Transfer assets
// @Description  Transfer assets to an address
// @Accept       json
// @Produce      json
// @Tags         Transfer
// @Param        address  body  string  true  "Address"
// @Success      200      {object}  TransferResponseBody
// @Failure      400      {object}  responses.ErrorResponse
// @Failure      500      {object}  responses.ErrorResponse
// @Router       /v2/transfer [post]

func (controller *TransferController) Transfer(c echo.Context) error {
	userId := c.Get("UserID").(int64)
	// payload is an event
	var body TransferRequestBody
	// load request payload, params into nostr.Event struct
	if err := c.Bind(&body); err != nil {
		c.Logger().Errorf("Failed to bind transfer request body: %v", err)
		return c.JSON(http.StatusBadRequest, responses.BadArgumentsError)
	}

	if err := c.Validate(&body); err != nil {
		c.Logger().Errorf("Failed to validate transfer request body: %v", err)
		return c.JSON(http.StatusBadRequest, responses.BadArgumentsError)
	}
	// transfer assets
	msg, status := controller.svc.TransferAssets(c.Request().Context(), uint64(userId), body.Address)
	if !status {
		c.Logger().Errorf("Failed to send assets: %v", msg)
		return c.JSON(http.StatusBadRequest, responses.BadArgumentsError)
	}
	return c.JSON(http.StatusOK, &TransferResponseBody{
		Message: msg,
		Status:  status,
	})
}
