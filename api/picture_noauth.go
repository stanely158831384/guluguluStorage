package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/stanely158831384/guluguluStorage/db/sqlc"
)

type listPicturesByUsernameRequest struct {
	Username string `json:"username" binding:"required"`
	Limit  int32 `form:"limit" binding:"required,min=5,max=10"`
	Offset int32 `form:"offset" binding:"required"`
}

func (server *Server) ListPicturesByUsername(ctx *gin.Context){
	var req listPicturesByUsernameRequest
	if err := ctx.ShouldBindQuery(&req); err!= nil {
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	arg := db.ListPicturesByUsernameParams{
		Username: req.Username,
		Limit: 10,
		Offset: 0,
	}

	pictures, err := server.store.ListPicturesByUsername(ctx,arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	var rsp []pictureResponse
	for _, picture := range pictures{
		rsp = append(rsp,newPictureResponse(picture))
	}

	ctx.JSON(http.StatusOK,rsp)
	return
}

type listPictures struct {
	Limit  int32 `form:"limit" binding:"required,min=5,max=10"`
	Offset int32 `form:"offset" binding:"required"`
}

func (server *Server) listPictures(ctx *gin.Context){
	var req listPictures
	if err := ctx.ShouldBindQuery(&req); err!= nil {
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	arg := db.ListPicturesParams{
		Limit: 10,
		Offset: 0,
	}

	pictures, err := server.store.ListPictures(ctx,arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	var rsp []pictureResponse
	for _, picture := range pictures{
		rsp = append(rsp,newPictureResponse(picture))
	}

	ctx.JSON(http.StatusOK,rsp)
	return
}

type getPictureRequest struct {
	ID int64 `json:"id" binding:"required"`
}

func (server *Server) GetPicture(ctx *gin.Context){
	var req getPictureRequest

	if err := ctx.Copy().ShouldBindUri(&req); err!= nil {
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	picture, err := server.store.GetPicture(ctx,req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK,newPictureResponse(picture))
	return
}