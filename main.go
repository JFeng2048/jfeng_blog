package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/JFeng2048/jfeng_blog/handlers"
)

func main() {
	// API routes
	http.HandleFunc("/api/blogs", handlers.GetBlogs)
	http.HandleFunc("/api/blogs/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/blogs/") && r.URL.Path != "/api/blogs/" {
			handlers.GetBlogByID(w, r)
		} else {
			handlers.GetBlogs(w, r)
		}
	})
	
	http.HandleFunc("/api/projects", handlers.GetProjects)
	http.HandleFunc("/api/projects/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/projects/") && r.URL.Path != "/api/projects/" {
			handlers.GetProjectByID(w, r)
		} else {
			handlers.GetProjects(w, r)
		}
	})
	
	http.HandleFunc("/api/personal", handlers.GetPersonalInfo)

	// Root endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"message":"欢迎访问JFeng个人博客API","endpoints":{"/api/blogs":"获取所有博客","/api/blogs/{id}":"获取指定博客","/api/projects":"获取所有项目","/api/projects/{id}":"获取指定项目","/api/personal":"获取个人信息"}}`)
	})

	port := "8080"
	fmt.Printf("服务器启动在端口 %s\n", port)
	fmt.Println("访问 http://localhost:8080 查看API信息")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
