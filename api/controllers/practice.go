package controllers

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/lndaquino/scratchpay-challenge/pkg/entity"
	er "github.com/lndaquino/scratchpay-challenge/pkg/errors"
)

// PracticeController struct models a controller for practice routes
type PracticeController struct {
	usecase PracticeUsecase
}

// PracticeUsecase interface defines the practice usecase methods
type PracticeUsecase interface {
	Search(string, string, string, string) ([]*entity.Practice, error)
}

// NewPracticeController returns a practice controller instance
func NewPracticeController(usecase PracticeUsecase) *PracticeController {
	return &PracticeController{
		usecase: usecase,
	}
}

// Search handles requests on the GET practices route
func (ctl *PracticeController) Search(c echo.Context) error {
	name := c.QueryParam("name")
	state := c.QueryParam("state")
	from := c.QueryParam("from")
	to := c.QueryParam("to")

	if name == "" && state == "" && from == "" && to == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "You must set at least one search query parameter: [name] or [state] or [from] or [to]",
		})
	}

	practices, err := ctl.usecase.Search(strings.ToLower(name), state, from, to)
	if err != nil {
		switch err {
		case er.ErrNotFound:
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": err.Error(),
			})
		case er.ErrInvalidFrom:
			return c.JSON(http.StatusUnprocessableEntity, map[string]string{
				"error": err.Error(),
			})
		case er.ErrInvalidTo:
			return c.JSON(http.StatusUnprocessableEntity, map[string]string{
				"error": err.Error(),
			})
		case er.ErrInvalidParameters:
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}
	}

	response := parseSearchResponse(practices)

	return c.JSON(http.StatusOK, response)
}
