package data

import (
	"github.com/JFeng2048/jfeng_blog/models"
	"time"
)

// GetProjects returns sample project data
func GetProjects() []models.Project {
	return []models.Project{
		{
			ID:           1,
			Name:         "个人博客系统",
			Description:  "基于Go语言开发的个人博客后端API系统，提供博客、项目、个人信息等RESTful接口",
			Technologies: []string{"Go", "RESTful API", "JSON"},
			GithubURL:    "https://github.com/JFeng2048/jfeng_blog",
			LiveURL:      "",
			ImageURL:     "",
			StartDate:    time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			EndDate:      nil,
			Status:       "active",
		},
		{
			ID:           2,
			Name:         "任务管理应用",
			Description:  "一个简单的任务管理应用，支持任务的创建、编辑、删除和状态管理",
			Technologies: []string{"React", "Node.js", "MongoDB"},
			GithubURL:    "https://github.com/JFeng2048/task-manager",
			LiveURL:      "https://task-manager-demo.example.com",
			ImageURL:     "",
			StartDate:    time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC),
			EndDate:      func() *time.Time { t := time.Date(2023, 8, 31, 0, 0, 0, 0, time.UTC); return &t }(),
			Status:       "completed",
		},
		{
			ID:           3,
			Name:         "天气查询工具",
			Description:  "通过API获取天气信息的命令行工具，支持多城市查询",
			Technologies: []string{"Python", "API集成", "CLI"},
			GithubURL:    "https://github.com/JFeng2048/weather-cli",
			LiveURL:      "",
			ImageURL:     "",
			StartDate:    time.Date(2023, 3, 15, 0, 0, 0, 0, time.UTC),
			EndDate:      func() *time.Time { t := time.Date(2023, 4, 20, 0, 0, 0, 0, time.UTC); return &t }(),
			Status:       "completed",
		},
	}
}

// GetProjectByID returns a project by ID
func GetProjectByID(id int) *models.Project {
	projects := GetProjects()
	for _, project := range projects {
		if project.ID == id {
			return &project
		}
	}
	return nil
}
