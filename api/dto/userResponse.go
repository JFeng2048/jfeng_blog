package dto

import (
	"gorm.io/datatypes"
)

// @Description 用户响应结构体
// @Property ID int "用户ID"
// @Property Username string "用户名"
// @Property DisplayName string "显示名"
// @Property AvatarURL string "头像URL"
// @Property Bio string "用户简介"
// @Property Role string "用户角色"
// @Property Email string "用户邮箱"
// @Property SocialLinks datatypes.JSON "社交媒体链接"
type UserResponse struct {
	ID int `json:"id"`
	Username string `json:"username"`
	DisplayName string `json:"display_name"`
	AvatarURL string `json:"avatar_url"`
	Bio string `json:"bio"`
	Role string `json:"role"`
	Email string `json:"email"`
	SocialLinks datatypes.JSON `json:"social_links"`
}