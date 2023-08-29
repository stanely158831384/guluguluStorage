package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/stanely158831384/guluguluStorage/db/sqlc"
)

type ListFeelingsByProductIdParams struct {
	Limit  int32 `form:"limit" binding:"required,min=5,max=10"`
	Offset int32 `form:"offset" binding:"required"`
	ProductID int64 `form:"product_id" binding:"required"`
}

func (server *Server) ListFeelingsByProductIdNoAuth(ctx *gin.Context){
	var req ListFeelingsByProductIdParams

	if err := ctx.ShouldBindQuery(&req); err!= nil {
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	arg := db.ListFeelingsByProductIdParams{
		Limit: req.Limit,
		Offset: req.Offset,
		ProductID: req.ProductID,
	}

	feelings, err := server.store.ListFeelingsByProductId(ctx,arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	var rsp []FeelingResponse
	for _, feeling := range feelings {
		rsp = append(rsp, newFeelingResponse(feeling))
	}

	ctx.JSON(http.StatusOK,rsp)
	return
}

type ListFeelingsByUserIDParams struct {
	Limit  int32 `form:"limit" binding:"required,min=5,max=10"`
	Offset int32 `form:"offset" binding:"required"`
	UserID int64 `form:"user_id" binding:"required"`
}

func (server *Server) ListFeelingsByUserIdNoAuth(ctx *gin.Context){
	var req ListFeelingsByUserIDParams

	if err := ctx.ShouldBindQuery(&req); err!= nil {
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	arg := db.ListFeelingsByUserIdParams{
		Limit: req.Limit,
		Offset: req.Offset,
		UserID: req.UserID,
	}

	feelings, err := server.store.ListFeelingsByUserId(ctx,arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	var rsp []FeelingResponse
	for _, feeling := range feelings {
		rsp = append(rsp, newFeelingResponse(feeling))
	}

	ctx.JSON(http.StatusOK,rsp)
	return
}

type getFeelingRequest struct {
	ID int64 `json:"id" binding:"required"`
}

func (server *Server) getFeeling(ctx *gin.Context){
	var req getFeelingRequest
	if err := ctx.ShouldBindUri(&req); err!= nil {
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	feeling, err := server.store.GetFeeling(ctx,req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK,newFeelingResponse(feeling))
	return
}