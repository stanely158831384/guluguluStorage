package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/lib/pq"
	db "github.com/stanely158831384/guluguluStorage/db/sqlc"
)

type createIngredientRequest struct {
	PictureID int64 `json:"name" binding:"required,ingredient"`
	Ingredient []string `json:"ingredient" binding:"required"`
}

type ingredientResponse struct {
	ID int64 `json:"id"`
	Ingredient []string `json:"ingredient"`
	PictureID pgtype.Int8 `json:"picture_id"`
	CreatedAt time.Time `json:"created_at"`
}

type getIngredientRequest struct {
	ID int64 `json:"id" binding:"required"`
}

type deleteIngredientRequest struct {
	ID int64 `json:"id" binding:"required"`
}

func newIngredientResponse (ingredient db.Ingredient) ingredientResponse {
	return ingredientResponse{
		ID: ingredient.ID,
		Ingredient: ingredient.Ingredient,
		PictureID: ingredient.PictureID,
		CreatedAt: ingredient.CreatedAt,
	}
}

func (server *Server) createIngredient(ctx *gin.Context){
	var req createIngredientRequest
	if err := ctx.ShouldBindJSON(&req); err!= nil {
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	arg := db.CreateIngredientParams{
		Ingredient: req.Ingredient,
		PictureID: pgtype.Int8{Int64: req.PictureID, Valid: true},
	}

	ingredient, err := server.store.CreateIngredient(ctx,arg)
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
}



func (server *Server) deleteIngredient(ctx *gin.Context){
	var req deleteIngredientRequest
	if err := ctx.ShouldBindJSON(&req); err!= nil {
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	err := server.store.DeleteIngredient(ctx,req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK,gin.H{"status":"ok"})
	return
}