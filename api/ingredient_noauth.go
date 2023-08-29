package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/stanely158831384/guluguluStorage/db/sqlc"
)

type listIngredientsParams struct {
	Limit  int32 `form:"limit" binding:"required,min=5,max=10"`
	Offset int32 `form:"offset" binding:"required"`
}

func (server *Server) listIngredients(ctx *gin.Context){
	var req listIngredientsParams

	if err := ctx.ShouldBindQuery(&req); err!= nil {
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	arg := db.ListIngredientsParams{
		Limit: req.Limit,
		Offset: req.Offset,
	}

	ingredients, err := server.store.ListIngredients(ctx,arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	var rsp []ingredientResponse
	for _, ingredient := range ingredients {
		rsp = append(rsp, newIngredientResponse(ingredient))
	}

	ctx.JSON(http.StatusOK,rsp)
	return
}

func (server *Server) GetIngredient(ctx *gin.Context){
	var req getIngredientRequest
	if err := ctx.ShouldBindUri(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ingredient, err := server.store.GetIngredient(ctx,req.ID)
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

	rsp := newIngredientResponse(ingredient)
	ctx.JSON(http.StatusOK,rsp)

	return

}
