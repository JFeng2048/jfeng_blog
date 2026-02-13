package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/JFeng2048/jfeng_blog/data"
)

// GetProjects handles GET /api/projects - returns all projects
func GetProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	projects := data.GetProjects()
	json.NewEncoder(w).Encode(projects)
}

// GetProjectByID handles GET /api/projects/{id} - returns a specific project
func GetProjectByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Extract ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/api/projects/")
	id, err := strconv.Atoi(path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid project ID"})
		return
	}

	project := data.GetProjectByID(id)
	if project == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Project not found"})
		return
	}

	json.NewEncoder(w).Encode(project)
}
