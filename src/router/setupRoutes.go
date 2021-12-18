package router

import (
	"fmt"
	"net/http"

	"github.com/ranon-rat/sharemyscreen/src/controllers"
)

func SetupRoutes() {

	http.HandleFunc("/", controllers.StreamingImage)
	fmt.Println("Listening on port 8080")

	// Start server
	http.ListenAndServe(":8080", nil)
}
