package handlers

import (
	"canvas-backend/util"
	"encoding/json"
	"net/http"
)

func (a *APIState) GenerateBackgrounds(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Product string   `json:"product"`
		Brand   string   `json:"brand"`
		Colors  []string `json:"colors"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	// 1. Build prompt
	prompt := util.GenerateBackgroundPrompt(
		req.Product,
		req.Brand,
		req.Colors,
	)

	// 2. Generate AI Backgrounds
	imgBytesList, err := util.GenerateBackgroundImages(prompt)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 3. Upload to Cloudinary
	var urls []string
	for _, img := range imgBytesList {
		u := util.UploadToCloudinary(img)
		urls = append(urls, u)
	}

	// 4. Return URLs
	json.NewEncoder(w).Encode(map[string]interface{}{
		"background_urls": urls,
	})
}
