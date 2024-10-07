package main

import (
	controllers "app1/internal/app/worker-receiver/controllers"
	internal "app1/internal/pkg/utils"
	utils "app1/pkg/utils"
	"context"
	"fmt"
	"net/http"
	"os"
)

// Instances
var (

	// Helpers
	environmentHelper = internal.NewEnvironmentHelper()
	applicationHelper = utils.NewApplicationHelper()
	httpHelper        = utils.NewHTTPHelper()

	// Services

	// Controllers
	livenessController  = controllers.NewLivenessController()
	readinessController = controllers.NewReadinessController()
	helloController     = controllers.NewHelloController()
)

func init() {
	ctx := utils.ContextWithTransaction(context.Background())
	environmentHelper.Check(ctx)
}

func main() {

	ctx, cancel := context.WithCancel(utils.ContextWithTransaction(context.Background()))

	fmt.Println(cancel)

	defer applicationHelper.RegisterGracefulShutdown(ctx, cancel)

	httpHelper.RegisterControllers(
		ctx,
		readinessController,
		livenessController,
		helloController,
	)

	port := os.Getenv("PORT")
	utils.GetLogger().Info(ctx).
		Str("port", port).
		Msg("starting app")

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err.Error())
	}
}
