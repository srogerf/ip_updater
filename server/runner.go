// using https://github.com/gin-gonic/gin

package server

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	_ "net/http"
)

func Runner() {
	log.Println("starting a server")
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("./public_html", true)))
	// Setup route group for the API
	api := r.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	//	r.GET("/", func(c *gin.Context) {
	//		c.JSON(200, gin.H{
	//			"message": "pong",
	//		})
	//	})

	//	r.NoRoute(func(c *gin.Context) {
	//		c.HTML(http.StatusOK, "index.html", nil)
	//	})

	r.Run(":8081")
	//	http.Handle("/", http.FileServer(http.Dir("./public_html")))

	//	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
	//		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	//		//panic("were done")
	//	})
	//	log.Fatal(http.ListenAndServe(":8081", nil))
}
