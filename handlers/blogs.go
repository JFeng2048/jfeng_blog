package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/JFeng2048/jfeng_blog/data"
)

// GetBlogs handles GET /api/blogs - returns all blogs
func GetBlogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	blogs := data.GetBlogs()
	json.NewEncoder(w).Encode(blogs)
}

// GetBlogByID handles GET /api/blogs/{id} - returns a specific blog
func GetBlogByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/api/blogs/")
	id, err := strconv.Atoi(path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid blog ID"})
		return
	}

	blog := data.GetBlogByID(id)
	if blog == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Blog not found"})
		return
	}

	json.NewEncoder(w).Encode(blog)
}
