package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"mymachine707/clients"
	"mymachine707/config"
	docs "mymachine707/docs" // docs is generated by Swag CLI, you have to import it.
	"mymachine707/handlars"

	_ "github.com/lib/pq"
)

// @license.name	Apache 2.0
func main() {

	cfg := config.Load()
	if cfg.Environment != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	docs.SwaggerInfo.Title = cfg.App
	docs.SwaggerInfo.Version = cfg.AppVersion

	fmt.Println("----->>")
	fmt.Printf("%+v\n", cfg)
	fmt.Println("---->>")

	r := gin.New()

	if cfg.Environment != "production" {
		r.Use(gin.Logger(), gin.Recovery()) // Later they will be replaced by custom Logger and Recovery
	}

	r.GET("/ping", MyCORSMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	grpcClients, err := clients.NewGrpcClients(cfg)
	if err != nil {
		panic(err)
	}

	defer grpcClients.Close()

	h := handlars.NewHandler(cfg, grpcClients)
	// Gruppirovka qilindi
	v1 := r.Group("v2")
	{
		v1.Use(MyCORSMiddleware())
		v1.POST("/login", h.Login)

		v1.POST("/article", h.AuthMiddleware("*"), h.CreatArticle)
		v1.GET("/article/:id", h.AuthMiddleware("*"), h.GetArticleByID)
		v1.GET("/article", h.AuthMiddleware("*"), h.GetArticleList)
		v1.PUT("/article", h.AuthMiddleware("*"), h.ArticleUpdate)
		v1.DELETE("/article/:id", h.AuthMiddleware("ADMIN"), h.DeleteArticle)
		v1.GET("/my-article/:id", h.AuthMiddleware("*"), h.SearchArticleByMyUsername)

		v1.POST("/author", h.AuthMiddleware("*"), h.CreatAuthor)
		v1.GET("/author/:id", h.AuthMiddleware("*"), h.GetAuthorByID)
		v1.GET("/author", h.AuthMiddleware("*"), h.GetAuthorList)
		v1.PUT("/author", h.AuthMiddleware("*"), h.AuthorUpdate)
		v1.DELETE("/author/:id", h.AuthMiddleware("ADMIN"), h.DeleteAuthor)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(cfg.HTTPPort) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func MyCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("MyCORSMiddleware...")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()

	}
}
