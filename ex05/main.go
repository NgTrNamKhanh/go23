package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	routes:= gin.Default()

	routes.GET("", func(ctx *gin.Context){
		ctx.JSON(http.StatusOK, "Welcome home")
	})

	server := &http.Server{
		Addr : "8888",
		Handler: routes,
	}
	err := server.ListenAndServe()
	//hepler.ErrorPanic(err)

}