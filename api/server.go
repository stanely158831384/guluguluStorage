package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/stanely158831384/guluguluStorage/db/sqlc"
	"github.com/stanely158831384/guluguluStorage/token"
	"github.com/stanely158831384/guluguluStorage/util"
)

//Server serves HTTP requests for our banking service.
type Server struct {
	config util.Config
	store db.Store
	tokenMaker token.Maker
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config: config,
		tokenMaker: tokenMaker,
		store: store,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok{
		v.RegisterValidation("currency",validCurrency)
		v.RegisterValidation("category",validCategory)
	}



	server.setupRouter()

	return server, nil
}

// NewServer creates a new HTTP server and setup routing.
func NewServer2(store db.Store) (*Server, error) {
	server := &Server{
		store: store,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok{
		v.RegisterValidation("currency",validCurrency)
		v.RegisterValidation("category",validCategory)
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter(){
	router := gin.Default()
	router.GET("/listCategories/noAuth", server.listCategoriesNoAuth)
	router.GET("/getCategory/noAuth/:id", server.getCategory)
	router.GET("listFeelingsByProductId/noAuth", server.ListFeelingsByProductIdNoAuth)
	router.GET("listFeelingsByUserId/noAuth", server.ListFeelingsByUserIdNoAuth)
	router.GET("getFeeling/noAuth/:id", server.getFeeling)
	router.GET("listIngredients/noAuth", server.listIngredients)
	router.GET("GetIngredient/noAuth/:id", server.GetIngredient)
	router.GET("GetPicture/noAuth/:id", server.GetPicture)
	router.GET("ListProductByUsername/noAuth/:username",server.ListProductByUsername)
	router.GET("ListProductsByCategory/noAuth/:category_id",server.ListProductsByCategoryId)
	router.GET("listProducts/noAuth",server.listProducts)
	router.GET("getProduct/noAuth/:id",server.getProduct)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.POST("/categories", server.createCategory)
	authRoutes.DELETE("/categories/delete", server.deleteCategory)
	authRoutes.POST("/createFeeling", server.createFeeling)
	authRoutes.POST("/ListFeelingsByUserID", server.ListFeelingsByUserID)
	authRoutes.POST("/ListFeelingsByProductID", server.ListFeelingsByProductID)
	authRoutes.DELETE("/deleteFeeling", server.deleteFeeling)
	authRoutes.POST("/createIngredient", server.createIngredient)
	authRoutes.DELETE("/deleteIngredient", server.deleteIngredient)
	authRoutes.POST("/createPicture", server.createPicture)
	authRoutes.DELETE("/deletePicture", server.deletePicture)
	authRoutes.POST("/createProduct", server.createProduct)
	authRoutes.DELETE("/deleteProduct", server.DeleteProduct)

	server.router = router

}

//start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error":err.Error()}
}