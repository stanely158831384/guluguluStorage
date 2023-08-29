package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/stanely158831384/guluguluStorage/db/sqlc"
)


type getProductByUsernameRequest struct {
	Username string `json:"username" binding:"required"`
}

func (server *Server) ListProductByUsername(ctx *gin.Context){
	var req getProductByUsernameRequest
	if err := ctx.ShouldBindQuery(&req); err!= nil {
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	arg := db.ListProductsByUserIDParams{
		Username: req.Username,
	}

	products, err := server.store.ListProductsByUserID(ctx,arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	var rsp []productResponse
	for _, product := range products {
		rsp = append(rsp, newProductResponse(product))
	}

	ctx.JSON(http.StatusOK,rsp)
	return
}

type getProductByCategoryRequest struct {
	CategoryID int64 `json:"category_id" binding:"required"`
}

func (server *Server) ListProductsByCategoryId(ctx *gin.Context){
	var req getProductByCategoryRequest
	if err := ctx.ShouldBindQuery(&req); err!= nil {
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	arg := db.ListProductsByCategoryParams{
		CategoryID: req.CategoryID,
	}

	products, err := server.store.ListProductsByCategory(ctx,arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	var rsp []productResponse
	for _, product := range products {
		rsp = append(rsp, newProductResponse(product))
	}

	ctx.JSON(http.StatusOK,rsp)
	return
}

type ListProductsParams struct {
	Limit  int32 `form:"limit" binding:"required,min=5,max=10"`
	Offset int32 `form:"offset" binding:"required"`
}

func (server *Server) listProducts(ctx *gin.Context){
	var req listProductRequest
	if err := ctx.ShouldBindQuery(&req); err!= nil {
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	ListProductsParams := db.ListProductsParams{
		Limit: req.Limit,
		Offset: req.Offset,
	}

	products, err := server.store.ListProducts(ctx,	ListProductsParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	var rsp []productResponse
	for _, product := range products {
		rsp = append(rsp, newProductResponse(product))
	}

	ctx.JSON(http.StatusOK,rsp)
	return
}

func (server *Server) getProduct(ctx *gin.Context){
	var req getProductRequest
	if err := ctx.ShouldBindUri(&req); err!= nil {
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	product, err := server.store.GetProduct(ctx,req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	rsp := newProductResponse(product)
	ctx.JSON(http.StatusOK,rsp)
	return
}