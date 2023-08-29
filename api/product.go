package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/stanely158831384/guluguluStorage/db/sqlc"
)

type createProductRequest struct {
	Name string `json:"name" binding:"required,product"`
	CategoryID int64 `json:"category_id" binding:"required"`
	IngredientsID int64 `json:"ingredients_id" binding:"required"`
}

type productResponse struct {
	Name string `json:"name"`
	CategoryID int64 `json:"category_id"`
	IngredientsID int64 `json:"ingredients_id"`
	CreatedAt         time.Time `json:"created_at"`
}

type getProductRequest struct {
	ID int64 `json:"id" binding:"required"`
}





type deleteProductRequest struct {
	ID int64 `json:"id" binding:"required"`
}

type listProductRequest struct {
	Limit  int32 `form:"limit" binding:"required,min=5,max=10"`
	Offset int32 `form:"offset" binding:"required"`
}

func newProductResponse(product db.Product) productResponse {
	return productResponse{
		Name: product.Name,
		CategoryID: product.CategoryID,
		IngredientsID: product.IngredientsID,
		CreatedAt: product.CreatedAt,
	}
}

func (server *Server) createProduct(ctx *gin.Context){
	var req createProductRequest
	if err := ctx.ShouldBindJSON(&req); err!= nil {
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	arg := db.CreateProductParams{
		Name: req.Name,
		CategoryID: req.CategoryID,
		IngredientsID: req.IngredientsID,
	}

	product, err := server.store.CreateProduct(ctx,arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok{
			fmt.Println("here is the err:", pqErr.Code.Name())

			switch pqErr.Code.Name(){
			case "foreign_key_violation","unique_violation":
				ctx.JSON(http.StatusForbidden,errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newProductResponse(product)
	ctx.JSON(http.StatusOK,rsp)
}







func (server *Server) DeleteProduct(ctx *gin.Context){
	var req deleteProductRequest
	if err := ctx.ShouldBindJSON(&req); err!= nil {
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	if err := server.store.DeleteProduct(ctx,req.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK,gin.H{"status": "ok"})
	return
}

