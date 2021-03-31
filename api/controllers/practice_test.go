package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/lndaquino/scratchpay-challenge/datastore/repos"
	"github.com/lndaquino/scratchpay-challenge/pkg/domain/practice"
	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	repo := repos.NewPracticeInMemoryRepo()
	p := repo.SeedTest()

	usecase := practice.NewPracticeUseCase(repo)
	ctl := NewPracticeController(usecase)

	samples := []struct {
		name       string
		state      string
		from       string
		to         string
		statusCode int
		test       string
	}{
		{
			name:       "",
			state:      "",
			from:       "",
			to:         "",
			statusCode: http.StatusBadRequest,
			test:       "no search parameters",
		},
		{
			name:       p.Name,
			state:      "",
			from:       "",
			to:         "",
			statusCode: http.StatusOK,
			test:       "valid name",
		},
		{
			name:       "",
			state:      p.StateName,
			from:       "",
			to:         "",
			statusCode: http.StatusOK,
			test:       "valid state name",
		},
		{
			name:       "",
			state:      p.StateCode,
			from:       "",
			to:         "",
			statusCode: http.StatusOK,
			test:       "valid state code",
		},
		{
			name:       p.Name,
			state:      p.StateName,
			from:       "",
			to:         "",
			statusCode: http.StatusOK,
			test:       "valid name and state",
		},
		{
			name:       "",
			state:      "",
			from:       fmt.Sprintf("%02d:%02d", p.OpeningHour, p.OpeningMinute),
			to:         "",
			statusCode: http.StatusOK,
			test:       "valid from",
		},
		{
			name:       "",
			state:      "",
			from:       "",
			to:         fmt.Sprintf("%02d:%02d", p.ClosureHour, p.ClosureMinute),
			statusCode: http.StatusOK,
			test:       "valid to",
		},
		{
			name:       "",
			state:      "",
			from:       fmt.Sprintf("%02d:%02d", p.OpeningHour, p.OpeningMinute),
			to:         fmt.Sprintf("%02d:%02d", p.ClosureHour, p.ClosureMinute),
			statusCode: http.StatusOK,
			test:       "valid from and to",
		},
		{
			name:       p.Name,
			state:      "",
			from:       fmt.Sprintf("%02d:%02d", p.OpeningHour, p.OpeningMinute),
			to:         "",
			statusCode: http.StatusOK,
			test:       "valid name and from",
		},
		{
			name:       p.Name,
			state:      "",
			from:       "",
			to:         fmt.Sprintf("%02d:%02d", p.ClosureHour, p.ClosureMinute),
			statusCode: http.StatusOK,
			test:       "valid name and to",
		},
		{
			name:       "",
			state:      p.StateName,
			from:       fmt.Sprintf("%02d:%02d", p.OpeningHour, p.OpeningMinute),
			to:         "",
			statusCode: http.StatusOK,
			test:       "valid state and from",
		},
		{
			name:       "",
			state:      p.StateName,
			from:       "",
			to:         fmt.Sprintf("%02d:%02d", p.ClosureHour, p.ClosureMinute),
			statusCode: http.StatusOK,
			test:       "valid state and to",
		},
		{
			name:       p.Name,
			state:      p.StateName,
			from:       fmt.Sprintf("%02d:%02d", p.OpeningHour, p.OpeningMinute),
			to:         "",
			statusCode: http.StatusOK,
			test:       "valid name, state and from",
		},
		{
			name:       p.Name,
			state:      p.StateName,
			from:       "",
			to:         fmt.Sprintf("%02d:%02d", p.ClosureHour, p.ClosureMinute),
			statusCode: http.StatusOK,
			test:       "valid name, state and to",
		},
		{
			name:       p.Name,
			state:      p.StateName,
			from:       fmt.Sprintf("%02d:%02d", p.OpeningHour, p.OpeningMinute),
			to:         fmt.Sprintf("%02d:%02d", p.ClosureHour, p.ClosureMinute),
			statusCode: http.StatusOK,
			test:       "valid name, state, from and to",
		},
		{
			name:       "name not found",
			state:      "",
			from:       "",
			to:         "",
			statusCode: http.StatusNotFound,
			test:       "invalid name",
		},
		{
			name:       "",
			state:      "state not found",
			from:       "",
			to:         "",
			statusCode: http.StatusNotFound,
			test:       "invalid state",
		},
		{
			name:       "name not found",
			state:      "state not found",
			from:       "",
			to:         "",
			statusCode: http.StatusNotFound,
			test:       "invalid name and state",
		},
		{
			name:       "",
			state:      "",
			from:       fmt.Sprintf("%02d:%02d", p.ClosureHour, p.ClosureMinute),
			to:         "",
			statusCode: http.StatusNotFound,
			test:       "no availability",
		},
		{
			name:       "",
			state:      "",
			from:       "",
			to:         fmt.Sprintf("%02d:%02d", p.OpeningHour, p.OpeningMinute),
			statusCode: http.StatusNotFound,
			test:       "no availability",
		},
		{
			name:       "",
			state:      "",
			from:       "invalid from",
			to:         "",
			statusCode: http.StatusUnprocessableEntity,
			test:       "wrong from format",
		},
		{
			name:       "",
			state:      "",
			from:       "",
			to:         "invalid to",
			statusCode: http.StatusUnprocessableEntity,
			test:       "wrong to format",
		},
		{
			name:       "",
			state:      "",
			from:       fmt.Sprintf("%02d:%02d", p.ClosureHour, p.ClosureMinute),
			to:         fmt.Sprintf("%02d:%02d", p.OpeningHour, p.OpeningMinute),
			statusCode: http.StatusBadRequest,
			test:       "to earlier then from",
		},
	}

	for i, sample := range samples {
		t.Run(fmt.Sprintf("%d %s", i, sample.test), func(t *testing.T) {
			e := echo.New()

			req := httptest.NewRequest(echo.GET, "/", nil)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.QueryParams().Add("name", sample.name)
			c.QueryParams().Add("state", sample.state)
			c.QueryParams().Add("from", sample.from)
			c.QueryParams().Add("to", sample.to)

			ctl.Search(c)

			assert.Equal(t, sample.statusCode, rec.Code)
		})
	}
}
