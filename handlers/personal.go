package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/JFeng2048/jfeng_blog/data"
)

// GetPersonalInfo handles GET /api/personal - returns personal information
func GetPersonalInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	info := data.GetPersonalInfo()
	json.NewEncoder(w).Encode(info)
}
