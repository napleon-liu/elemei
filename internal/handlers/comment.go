package handlers

import (
	"FoodDelivery/internal/service"
	"FoodDelivery/internal/typ/req"
	"FoodDelivery/internal/typ/resp"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// CreateComment customer create one comment
func CreateComment(ctx *gin.Context) {
	var createCommentReq req.CreateCommentReq
	if err := ctx.ShouldBindJSON(&createCommentReq); err != nil {
		log.Println(fmt.Errorf("[CreateComment] param bind error, %v", err))
		ctx.JSON(http.StatusBadRequest, &resp.T{
			MetaData: &resp.MetaData{
				Code: http.StatusBadRequest,
				Msg:  "An error message generated by ShouldBindJson: " + err.Error(),
			},
		})
	}
	if err := service.CreateComment(createCommentReq); err != nil {
		log.Println(fmt.Sprintf("[CreateComment] wrong: %v", err))
		ctx.JSON(http.StatusBadRequest, &resp.T{
			MetaData: &resp.MetaData{
				Code: http.StatusBadRequest,
				Msg:  "An error message generated by service.CreateComment: " + err.Error(),
			},
		})
		return
	}
	ctx.JSON(http.StatusOK, &resp.T{
		MetaData: &resp.MetaData{
			Code: http.StatusOK,
			Msg:  "create comment successfully",
		},
	})
	return
}

// GetCommentList customer view one dish's all comments
func GetCommentList(ctx *gin.Context) {
	var getCommentReq req.GetCommentReq
	if err := ctx.ShouldBindJSON(&getCommentReq); err != nil {
		log.Println(fmt.Errorf("[GetComment] param bind error, %v", err))
		ctx.JSON(http.StatusBadRequest, &resp.T{
			MetaData: &resp.MetaData{
				Code: http.StatusBadRequest,
				Msg:  "An error message generated by ShouldBindJson: " + err.Error(),
			},
		})
	}
	commentList, err := service.GetCommentList(getCommentReq)
	if err != nil {
		log.Println(fmt.Sprintf("[GetComment] wrong: %v", err))
		ctx.JSON(http.StatusBadRequest, &resp.T{
			MetaData: &resp.MetaData{
				Code: http.StatusBadRequest,
				Msg:  "An error message generated by service.GetComment: " + err.Error(),
			},
		})
		return
	}
	ctx.JSON(http.StatusOK, &resp.T{
		MetaData: &resp.MetaData{
			Code: http.StatusOK,
			Msg:  "get comment list successfully",
		},
		Result: commentList,
	})
}
