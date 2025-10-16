package tapcontrollers

import (
	"net/http"

	"github.com/bittap-protocol/lnhub/lib/responses"
	"github.com/bittap-protocol/lnhub/lib/service"
	"github.com/labstack/echo/v4"
)

// UniverseController : UniverseController struct
type UniverseController struct {
	svc *service.LndhubService
}

func NewUniverseController(svc *service.LndhubService) *UniverseController {
	return &UniverseController{svc: svc}
}

// universe assets response
// / universe response
type UniverseAssetsResponseBody struct {
	Assets map[string]string `json:"assets"`
}

// Universe godoc
// @Summary      Retrieve universe assets
// @Description  Retrieve universe assets
// @Accept       json
// @Produce      json
// @Tags         Taproot Assets
// @Success      200  {object}  UniverseAssetsResponseBody
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /tap/universe-assets [get]
func (controller *UniverseController) UniverseAssets(c echo.Context) error {
	data, err := controller.svc.GetUniverseAssetsJson(c.Request().Context())
	if err != nil {
		c.Logger().Errorf("Failed to retrieve universe assets: %v", err)
		return c.JSON(http.StatusBadRequest, responses.BadArgumentsError)
	}

	return c.JSON(http.StatusOK, &UniverseAssetsResponseBody{
		Assets: data,
	})
}
