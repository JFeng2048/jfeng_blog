package data

import (
	"time"
	"github.com/JFeng2048/jfeng_blog/models"
)

// GetBlogs returns sample blog data
func GetBlogs() []models.Blog {
	return []models.Blog{
		{
			ID:          1,
			Title:       "开始我的个人博客之旅",
			Content:     "这是我的第一篇博客文章，记录了我开始写博客的初衷和计划。作为一名开发者，我希望通过博客分享我的技术学习经验和项目实践。",
			Author:      "JFeng",
			Tags:        []string{"博客", "个人成长", "技术"},
			CreatedAt:   time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
			UpdatedAt:   time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
			ViewCount:   120,
			Category:    "随笔",
			Description: "记录开始写博客的初衷和计划",
		},
		{
			ID:          2,
			Title:       "Go语言学习笔记",
			Content:     "Go语言是一门简洁高效的编程语言，本文总结了我在学习Go语言过程中的一些心得体会，包括并发编程、接口设计等核心概念。",
			Author:      "JFeng",
			Tags:        []string{"Go", "编程", "后端开发"},
			CreatedAt:   time.Date(2024, 2, 15, 14, 30, 0, 0, time.UTC),
			UpdatedAt:   time.Date(2024, 2, 16, 9, 0, 0, 0, time.UTC),
			ViewCount:   256,
			Category:    "技术",
			Description: "Go语言学习心得与经验分享",
		},
		{
			ID:          3,
			Title:       "RESTful API设计最佳实践",
			Content:     "本文介绍了RESTful API设计的最佳实践，包括URL设计、HTTP方法使用、状态码选择、错误处理等方面的内容。",
			Author:      "JFeng",
			Tags:        []string{"API", "RESTful", "后端开发"},
			CreatedAt:   time.Date(2024, 3, 10, 16, 0, 0, 0, time.UTC),
			UpdatedAt:   time.Date(2024, 3, 10, 16, 0, 0, 0, time.UTC),
			ViewCount:   189,
			Category:    "技术",
			Description: "RESTful API设计规范与实践指南",
		},
	}
}

// GetBlogByID returns a blog by ID
func GetBlogByID(id int) *models.Blog {
	blogs := GetBlogs()
	for _, blog := range blogs {
		if blog.ID == id {
			return &blog
		}
	}
	return nil
}
