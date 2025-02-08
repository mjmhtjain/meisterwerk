package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mjmhtjain/meisterwerk/internal/handlers"
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

	createQuotesRouter(api_v1, r.db)

	return r.engine
}

func createQuotesRouter(api_v1 *gin.RouterGroup, db *gorm.DB) {
	quotes := api_v1.Group("/quotes")
	quoteHandler := handlers.NewQuoteHandler(db)
	quotes.POST("", quoteHandler.CreateQuote())
}
