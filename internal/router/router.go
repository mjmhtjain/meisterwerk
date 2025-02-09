package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mjmhtjain/meisterwerk/internal/handlers"
	"github.com/mjmhtjain/meisterwerk/internal/services"
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
	createProductsRouter(api_v1)

	return r.engine
}

func createQuotesRouter(api_v1 *gin.RouterGroup) {
	quotes := api_v1.Group("/quote")
	quoteService := services.NewQuoteService()
	quoteHandler := handlers.NewQuoteHandler(quoteService)

	quotes.POST("", quoteHandler.CreateQuote)

	quotes.GET("/:id", quoteHandler.GetQuote)

	quotes.PUT("/:id/status", quoteHandler.UpdateQuoteStatus)
}

func createProductsRouter(api_v1 *gin.RouterGroup) {
	products := api_v1.Group("/all-products")
	productService := services.NewProductService()
	productHandler := handlers.NewProductHandler(productService)

	products.GET("", productHandler.GetAllProducts)
}
