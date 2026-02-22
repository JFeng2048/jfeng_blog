package handler

import (
	"strconv"

	"jfeng_blog/api/dto"
	"jfeng_blog/api/dto/response"
	"jfeng_blog/internal/service/user"

	"github.com/gin-gonic/gin"
)

// @Summary 根据用户ID获取用户信息
// @Description 根据用户ID获取用户详细信息
// @Tags 用户
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} response.SuccessResponse "获取用户信息成功"
// @Failure 400 {object} response.ErrorResponse "请求参数错误"
// @Failure 500 {object} response.ErrorResponse "内部服务器错误"
// @Router /user/{id} [get]
func GetUserInfoByID(c *gin.Context) {
	// 处理获取用户信息逻辑

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, response.ErrorResponse{
			Code: 400,
			Msg:  "用户ID格式错误",
			Data: nil,
		})
		return
	}
	user, err := user.GetUserInfoByID(ID)
	if err != nil {
		c.JSON(500, response.ErrorResponse{
			Code: 500,
			Msg:  "获取用户信息失败",
			Data: nil,
		})
		return
	}
	c.JSON(200, response.SuccessResponse{
		Code: 200,
		Msg:  "获取用户信息成功",
		Data: dto.UserResponse{
			ID:          int(user.ID),
			Username:    user.Username,
			DisplayName: user.DisplayName,
			AvatarURL:   user.AvatarURL,
			Bio:         user.Bio,
			Role:        string(user.Role),
			Email:       user.Email,
			SocialLinks: user.SocialLinks,
		}})
}
