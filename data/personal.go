package data

import (
	"github.com/JFeng2048/jfeng_blog/models"
)

// GetPersonalInfo returns personal information
func GetPersonalInfo() models.PersonalInfo {
	return models.PersonalInfo{
		Name:     "JFeng",
		Title:    "全栈开发工程师",
		Bio:      "热爱编程，专注于Web开发和软件工程。喜欢学习新技术，分享技术经验。",
		Email:    "jfeng@example.com",
		Location: "中国",
		Avatar:   "/avatar.jpg",
		SocialLinks: map[string]string{
			"github":  "https://github.com/JFeng2048",
			"twitter": "https://twitter.com/jfeng",
			"blog":    "https://jfeng.blog",
		},
		Skills: []string{
			"Go",
			"JavaScript",
			"Python",
			"React",
			"Node.js",
			"Docker",
			"Git",
			"RESTful API",
			"数据库设计",
		},
		Education: []models.Education{
			{
				School:    "某某大学",
				Degree:    "学士",
				Field:     "计算机科学与技术",
				StartYear: 2016,
				EndYear:   2020,
			},
		},
		Experience: []models.Experience{
			{
				Company:     "科技公司",
				Position:    "全栈开发工程师",
				Description: "负责Web应用开发，包括前端和后端功能的设计与实现",
				StartYear:   2020,
				EndYear:     nil,
				Technologies: []string{"Go", "React", "PostgreSQL"},
			},
			{
				Company:     "创业公司",
				Position:    "后端开发实习生",
				Description: "参与后端API开发和数据库设计工作",
				StartYear:   2019,
				EndYear:     func() *int { y := 2020; return &y }(),
				Technologies: []string{"Python", "Flask", "MySQL"},
			},
		},
	}
}
