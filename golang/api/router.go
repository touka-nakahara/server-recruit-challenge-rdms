package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pulse227/server-recruit-challenge-sample/api/middleware"
	"github.com/pulse227/server-recruit-challenge-sample/controller"
	"github.com/pulse227/server-recruit-challenge-sample/infra/db/mysql"
	"github.com/pulse227/server-recruit-challenge-sample/service"
)

func NewRouter() *mux.Router {
	//TODO 具体DBとの機能の切り替えを行うために, Composeを行う
	// DB作成
	connection := mysql.NewMySQLDB()
	db := mysql.NewSingerDB(connection)

	// RepoはDBとのIFを担当
	// singerRepo := memorydb.NewSingerRepository()
	// サービスはHTTPとDBの境界を担当
	singerService := service.NewSingerService(db)
	// コントローラーはHTTPリクエストの処理を担当
	singerController := controller.NewSingerController(singerService)

	r := mux.NewRouter()

	// GET
	r.HandleFunc("/singers", singerController.GetSingerListHandler).Methods(http.MethodGet)
	// GET/ID
	r.HandleFunc("/singers/{id:[0-9]+}", singerController.GetSingerDetailHandler).Methods(http.MethodGet)
	// POST
	r.HandleFunc("/singers", singerController.PostSingerHandler).Methods(http.MethodPost)
	// DELETE
	r.HandleFunc("/singers/{id:[0-9]+}", singerController.DeleteSingerHandler).Methods(http.MethodDelete)

	// LOGGING
	r.Use(middleware.LoggingMiddleware)

	return r
}
