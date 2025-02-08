package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mjmhtjain/meisterwerk/internal/handlers"
	"github.com/mjmhtjain/meisterwerk/internal/services"
	"gorm.io/gorm"
)

type Router struct {
	engine *gin.Engine
	db     *gorm.DB
}

func NewRouter(db *gorm.DB) *Router {
	return &Router{
		engine: gin.Default(),
		db:     db,
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
	quotes := api_v1.Group("/quote")
	quoteService := services.NewQuoteService()
	quoteHandler := handlers.NewQuoteHandler(quoteService)

	quotes.POST("", quoteHandler.CreateQuote())
}
