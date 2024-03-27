package routers

import (
	"fmt"
	"marketplace/app"
	"marketplace/configs"
	"marketplace/internal/repository"
	"marketplace/internal/rest/middleware"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
)

func Run() (err error) {
	db, err := app.GetPostgres()
	if err != nil {
		return
	}
	defer db.Close()

	userStorage := repository.NewUserPg(db)

	rootRouter := mux.NewRouter()
	MountAuthRouter(rootRouter, userStorage)

	rootRouter.Use(middleware.PanicRecoverMiddleware)

	fmt.Printf("\tstarting server at %d\n", configs.PORT)
	err = http.ListenAndServe(fmt.Sprintf(":%d", configs.PORT), rootRouter)

	return
}
