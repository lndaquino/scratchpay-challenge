package routers

import (
	"github.com/labstack/echo"
	"github.com/lndaquino/scratchpay-challenge/api/controllers"
)

// SystemRoutes struct models a system level router
type SystemRoutes struct {
	practiceController *controllers.PracticeController
}

// NewSystemRoutes returns a SystemRoutes instance
func NewSystemRoutes(c *controllers.PracticeController) *SystemRoutes {
	return &SystemRoutes{
		practiceController: c,
	}
}

// MakeControllers setups the app routes
func (routes *SystemRoutes) MakeControllers() *echo.Echo {
	e := echo.New()

	e.GET("/", routes.practiceController.Search)

	return e
}
