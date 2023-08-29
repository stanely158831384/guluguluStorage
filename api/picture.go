package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/stanely158831384/guluguluStorage/db/sqlc"
)

type createPictureRequest struct {
	link string `json:"link" binding:"required"`
	username string `json:"username" binding:"required"`
}

type deletePictureRequest struct {
	id int64 `json:"id" binding:"required"`
}

type pictureResponse struct {
	id int64 `json:"id"`
	link string `json:"link"`
	username string `json:"username"`
	createdAt time.Time `json:"created_at"`
}

func newPictureResponse(picture db.Picture) pictureResponse {
	return pictureResponse{
		id: picture.ID,
		link: picture.Link,
		username: picture.Username,
		createdAt: picture.CreatedAt,
	}
}

func (server *Server) createPicture(ctx *gin.Context){
	var req createPictureRequest
	if err := ctx.ShouldBindJSON(&req); err!= nil {
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	arg := db.CreatePictureParams{
		Link: req.link,
		Username: req.username,
	}



	picture, err := server.store.CreatePicture(ctx,arg)
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

	rsp := newPictureResponse(picture)
	ctx.JSON(http.StatusOK,rsp)
	return
}

func (server *Server) deletePicture(ctx *gin.Context){
	var req deletePictureRequest
	if err := ctx.ShouldBindJSON(&req); err!= nil {
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	err := server.store.DeletePicture(ctx,req.id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK,gin.H{"status":"ok"})
	return
}

