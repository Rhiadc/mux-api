package main

import (
	"course/controllers"
	router "course/http"
	"fmt"
	"net/http"
)

var (
	postController controllers.PostController = controllers.NewPostController()
	httpRouter     router.Router              = router.NewMuxRouter()
)

func main() {
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "up and running...")
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.CreatePost)

	const port string = ":8000"
	httpRouter.SERVE(port)
}
