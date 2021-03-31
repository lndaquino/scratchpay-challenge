package main

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/lndaquino/scratchpay-challenge/api/controllers"
	"github.com/lndaquino/scratchpay-challenge/api/routers"
	"github.com/lndaquino/scratchpay-challenge/datastore/repos"
	"github.com/lndaquino/scratchpay-challenge/pkg/domain/practice"
)

func init() {
	godotenv.Load()
}

func main() {
	var e *echo.Echo
	var app *routers.SystemRoutes

	var err error
	app, err = setupApplication()
	if err != nil {
		panic("Error starting application ==> " + err.Error())
	}

	e = app.MakeControllers()
	server := &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodPatch, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAuthorization},
	}))

	e.Logger.Fatal(e.StartServer(server))
}

func setupApplication() (*routers.SystemRoutes, error) {
	practiceRepo := repos.NewPracticeInMemoryRepo()
	if err := practiceRepo.Seed(); err != nil {
		return &routers.SystemRoutes{}, errors.New("Error seeding database ==> " + err.Error())
	}

	practiceUsecase := practice.NewPracticeUseCase(practiceRepo)
	practiceController := controllers.NewPracticeController(practiceUsecase)
	systemRoutes := routers.NewSystemRoutes(practiceController)

	return systemRoutes, nil
}
