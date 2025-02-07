package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mjmhtjain/meisterwerk/internal/handlers"
)

type Router struct {
	engine *gin.Engine
}

func NewRouter() *Router {
	return &Router{
		engine: gin.Default(),
	}
}

func (r *Router) Setup() *gin.Engine {

	// Health endpoint
	r.engine.GET("/health", handlers.NewHealthHandler().Handle())

	api_v1 := r.engine.Group("/api/v1")

	createQuotesRouter(api_v1)

	return r.engine
}

func createQuotesRouter(api_v1 *gin.RouterGroup) {
	quotes := api_v1.Group("/quotes")

	quoteHandler := handlers.NewQuoteHandler()
	quotes.POST("", quoteHandler.CreateQuote())

}
