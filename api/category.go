package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/stanely158831384/guluguluStorage/db/sqlc"
)


type createCategoryRequest struct {
	Name string `json:"name" binding:"required,category"`
}

type deleteCategoryRequest struct {
	ID int64 `json:"id" binding:"required"`
}

type categoryResponse struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
}

func newCategoryResponse(category db.Category) categoryResponse {
	return categoryResponse{
		ID: category.ID,
		Name: category.Name,
	}
}

func (server *Server) createCategory(ctx *gin.Context){
	var req createCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err!= nil {
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	category, err := server.store.CreateCategory(ctx,req.Name)
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

	rsp := newCategoryResponse(category)
	ctx.JSON(http.StatusOK,rsp)
}

func (server *Server) deleteCategory(ctx *gin.Context){
	var req deleteCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err!= nil {
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	err := server.store.DeleteCategory(ctx,req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK,gin.H{"status":"ok"})
	return
}



