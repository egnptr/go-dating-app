package main

import (
	"fmt"
	"net/http"
	"os"

	controller "github.com/egnptr/dating-app/delivery/http"
	router "github.com/egnptr/dating-app/pkg/http"
	"github.com/egnptr/dating-app/repository/cache"
	"github.com/egnptr/dating-app/repository/db"
	"github.com/egnptr/dating-app/usecase"
)

func main() {
	redisURL := "127.0.0.1:6379"
	if os.Getenv("REDIS_URL") != "" {
		redisURL = os.Getenv("REDIS_URL")
	}

	var (
		dbRepo     = db.NewSQLiteRepository()
		cacheRepo  = cache.NewRedisCache(redisURL, 1)
		service    = usecase.NewUsecase(dbRepo, cacheRepo)
		delivery   = controller.NewPostController(service)
		httpRouter = router.NewMuxRouter()
	)

	const port string = ":8080"
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})
	httpRouter.POST("/user/sign-up", delivery.SignUp)
	httpRouter.POST("/user/login", delivery.LoginUser)

	httpRouter.POST("/subscribe-premium", delivery.UpdateSubscription)
	httpRouter.POST("/unsubscribe-premium", delivery.UpdateSubscription)
	httpRouter.GET("/related-profiles", delivery.GetProfiles)
	httpRouter.POST("/swipe", delivery.Swipe)

	httpRouter.SERVE(port)
}
