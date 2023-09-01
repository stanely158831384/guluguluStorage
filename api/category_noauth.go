package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/stanely158831384/guluguluStorage/db/sqlc"
)

type ListCategoriesParams struct {
	Limit  int32 `form:"limit" binding:"required,min=5,max=10"`
	Offset *int32 `form:"offset" binding:"required"`
}


func (server *Server) listCategoriesNoAuth(ctx *gin.Context){
	fmt.Printf("listCategoriesNoAuth is called\n")
	var req ListCategoriesParams
	if err := ctx.ShouldBindQuery(&req); err!= nil {
		fmt.Printf("error\n")
		fmt.Printf("%v\n", req)
		fmt.Printf("%v\n", err)
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	arg := db.ListCategoriesParams{
		Limit: req.Limit,
		Offset: *req.Offset,
	}
	fmt.Printf("listCategoriesNoAuth is called2\n")

	categories, err := server.store.ListCategories(ctx,arg)
	if err != nil {
		fmt.Printf("error2\n")
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}
	fmt.Printf("listCategoriesNoAuth is called3\n")
	var rsp []categoryResponse
	for _, category := range categories {
		rsp = append(rsp, newCategoryResponse(category))
	}
	fmt.Printf("listCategoriesNoAuth is called4\n")

	ctx.JSON(http.StatusOK,rsp)
	fmt.Printf("listCategoriesNoAuth is called5\n")

}

type getCategoryRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

func (server *Server) getCategory(ctx *gin.Context){
	fmt.Printf("getCategory is called\n")
	var req getCategoryRequest

	if err := ctx.ShouldBindUri(&req); err!= nil {
		fmt.Printf("getCategory2 error\n")
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}
	fmt.Printf("getCategory2 is called\n")
	category, err := server.store.GetCategory(ctx,req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}
	fmt.Printf("getCategory3 is called\n")



	ctx.JSON(http.StatusOK,newCategoryResponse(category))
}