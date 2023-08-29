package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/stanely158831384/guluguluStorage/db/sqlc"
)


type createFeelingRequest struct {
	productID int64 `json:"product_id" binding:"required"`
	userID int64 `json:"user_id" binding:"required"`
	userName string `json:"user_name" binding:"required"`
	comment string `json:"comment" binding:"required"`
	recommend bool `json:"recommend" binding:"required"`
}

type FeelingResponse struct {
	productID int64 `json:"product_id"`
	userID int64 `json:"user_id"`
	userName string `json:"user_name"`
	comment string `json:"comment"`
	recommend bool `json:"recommend"`
	createdAt time.Time `json:"created_at"`
}

func newFeelingResponse(feeling db.Feeling) FeelingResponse {
	return FeelingResponse{
		productID: feeling.ProductID,
		userID: feeling.UserID,
		userName: feeling.Username,
		comment: feeling.Comment,
		recommend: feeling.Recommend,
		createdAt: feeling.CreatedAt,
	}
}

func (server *Server) createFeeling(ctx *gin.Context){
	var req createFeelingRequest
	if err := ctx.ShouldBindJSON(&req); err!= nil {
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	arg := db.CreateFeelingParams{
		ProductID: req.productID,
		UserID: req.userID,
		Username: req.userName,
		Comment: req.comment,
		Recommend: req.recommend,
	}

	feeling, err := server.store.CreateFeeling(ctx,arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK,newFeelingResponse(feeling))
	return
}


type getFeelingByUserIDRequest struct {
	UserID int64 `json:"user_id" binding:"required"`
}

func (server *Server) ListFeelingsByUserID(ctx *gin.Context){
	var req getFeelingByUserIDRequest
	if err := ctx.ShouldBindQuery(&req); err!= nil {
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	arg := db.ListFeelingsByProductIdParams{
		ProductID: req.UserID,
		Limit: 10,
		Offset: 0,
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

type getFeelingByProductIDRequest struct {
	ProductID int64 `json:"product_id" binding:"required"`
}

func (server *Server) ListFeelingsByProductID(ctx *gin.Context){
	var req getFeelingByProductIDRequest
	if err := ctx.ShouldBindQuery(&req); err!= nil {
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	arg := db.ListFeelingsByProductIdParams{
		ProductID: req.ProductID,
		Limit: 10,
		Offset: 0,
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

type deleteFeelingRequest struct {
	ID int64 `json:"id" binding:"required"`
}

func (server *Server) deleteFeeling(ctx *gin.Context){
	var req getFeelingRequest
	if err := ctx.ShouldBindJSON(&req); err!= nil {
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	err := server.store.DeleteFeeling(ctx,req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK,gin.H{"message":"deleted"})
	return
}
