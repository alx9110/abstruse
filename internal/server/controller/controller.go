package controller

import (
	"github.com/google/wire"
	"github.com/jkuri/abstruse/internal/pkg/http"
)

func CreateInitControllersFn(
	uc *UserController,
	vc *VersionController,
	wc *WorkerController,
	bc *BuildController,
) http.InitControllers {
	return func(r *http.Router) {
		r.POST("/api/user/login", uc.Login)
		r.GET("/api/user/:id", uc.Find)
		r.GET("/api/version", vc.GetInfo)
		r.GET("/api/worker", wc.GetWorkers)
		r.POST("/api/build/start", bc.StartJob)
	}
}

var ProviderSet = wire.NewSet(
	NewUserController,
	NewVersionController,
	NewWorkerController,
	NewBuildController,
	CreateInitControllersFn,
)
