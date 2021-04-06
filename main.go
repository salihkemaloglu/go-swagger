package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/testarea/api_endpoint_doc/docs"
)

// @title Backend API Service
// @version 1.0.0
// @host localhost:8080

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @securityDefinitions.apikey ApiKeyAuth (JWT)
// @in header
// @name AccessToken
// @BasePath /api/v1

func main() {
	router := gin.Default()

	// Simple group: v1
	h := NewHandler()
	v1 := router.Group("/api/v1")
	{
		health := v1.Group("/todo")
		{
			health.GET("get", h.SayHello)
			health.POST("post", h.TodoPost)
			health.PUT("put", h.TodoPut)
			health.DELETE("delete", h.TodoDelete)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}

// @Summary Todo GET Example
// @Description check gateway and account service health :)
// @Success 200 {object} Todo
// @Failure 400 {object} HTTPError
// @Failure 401 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /todo/get [get]
func (h *Handler) SayHello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "api-gateway is working ...")
}

// @Summary Todo POST Example
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param data body Todo true "The input todo struct"
// @Success 200 {object} Todo
// @Failure 400 {object} HTTPError
// @Failure 401 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /todo/post [post]
func (h *Handler) TodoPost(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Todo{UserID: 1})
}

// @Summary Todo PUT Example
// @Description put
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param data body Todo true "The input todo struct"
// @Success 200 {object} Todo
// @Failure 400 {object} HTTPError
// @Failure 401 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /todo/put [put]
func (h *Handler) TodoPut(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "TodoPut is working ...")
}

// @Summary Todo DELETE Example
// @Description delete
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param user_id  path   int  true  "user's ID"
// @Success 200 {object} Todo
// @Failure 400 {object} HTTPError
// @Failure 401 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /todo/delete [delete]
func (h *Handler) TodoDelete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "TodoDelete is working ...")
}

type Todo struct {
	UserID   int    `json:"userID"`
	Username string `json:"username"`
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
