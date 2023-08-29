package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/stanely158831384/guluguluStorage/db/sqlc"
)

type ListCategoriesParams struct {
	Limit  int32 `form:"limit" binding:"required,min=5,max=10"`
	Offset int32 `form:"offset" binding:"required"`
}


func (server *Server) listCategoriesNoAuth(ctx *gin.Context){
	var req ListCategoriesParams

	if err := ctx.ShouldBindQuery(&req); err!= nil {
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	arg := db.ListCategoriesParams{
		Limit: req.Limit,
		Offset: req.Offset,
	}

	categories, err := server.store.ListCategories(ctx,arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	var rsp []categoryResponse
	for _, category := range categories {
		rsp = append(rsp, newCategoryResponse(category))
	}

	ctx.JSON(http.StatusOK,rsp)
}

type getCategoryRequest struct {
	ID int64 `json:"id" binding:"required"`
}

func (server *Server) getCategory(ctx *gin.Context){
	var req getCategoryRequest

	if err := ctx.Copy().ShouldBindUri(&req); err!= nil {
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	category, err := server.store.GetCategory(ctx,req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK,newCategoryResponse(category))
	return
}