package main

import (
	"fmt"
	"net/http"

	controller "github.com/egnptr/dating-app/delivery/http"
	router "github.com/egnptr/dating-app/pkg/http"
	"github.com/egnptr/dating-app/repository/cache"
	"github.com/egnptr/dating-app/repository/db"
	"github.com/egnptr/dating-app/usecase"
)

func main() {
	var (
		dbRepo     = db.NewSQLiteRepository()
		cacheRepo  = cache.NewRedisCache("localhost:6379", 1, 10)
		service    = usecase.NewUsecase(dbRepo, cacheRepo)
		delivery   = controller.NewPostController(service)
		httpRouter = router.NewMuxRouter()
	)

	const port string = ":8080"
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})
	httpRouter.POST("/user/signup", delivery.SignUp)
	httpRouter.POST("/user/login", delivery.LoginUser)

	httpRouter.POST("/subscribe-premium", delivery.UpdateSubscription)
	httpRouter.POST("/unsubscribe-premium", delivery.UpdateSubscription)
	httpRouter.GET("/related-profiles", delivery.GetProfiles)
	httpRouter.GET("/swipe", delivery.Swipe)

	httpRouter.SERVE(port)
}
