package main

import (
	"online-pathsaala/pkg/db"
	"online-pathsaala/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	engine := PrepareGinEngine()
	err := godotenv.Load(".env")
	if err != nil {
		return
	}
	con, err := db.ConnectDb()
	if err != nil {
		return
	}

	router.AddRoutes(con, engine)
	engine.Run(":8080") // listen and serve on 0.0.0.0:8080
}

func PrepareGinEngine() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	return r
}
