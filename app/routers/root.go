package routers

import (
	"fmt"
	"marketplace/app"
	"marketplace/configs"
	"marketplace/internal/repository"
	"marketplace/internal/rest/middleware"
	"marketplace/pkg/logging"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
)

func Run() (err error) {
	db, err := app.GetPostgres()
	if err != nil {
		return
	}
	defer db.Close()

	logFile, err := os.OpenFile(configs.LOGS_DIR+configs.LOGFILE_NAME, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	logging.InitLogger(logFile)

	userStorage := repository.NewUserPg(db)
	adStorage := repository.NewAdPg(db)

	rootRouter := mux.NewRouter()
	MountAuthRouter(rootRouter, userStorage)
	MountAdRouter(rootRouter, adStorage)

	rootRouter.Use(middleware.RequestID)
	rootRouter.Use(middleware.AccessLogMiddleware)
	rootRouter.Use(middleware.PanicRecoverMiddleware(logging.Logger))

	fmt.Printf("\tstarting server at %d\n", configs.PORT)
	err = http.ListenAndServe(fmt.Sprintf(":%d", configs.PORT), rootRouter)

	return
}
