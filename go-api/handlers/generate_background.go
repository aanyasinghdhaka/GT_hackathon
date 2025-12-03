package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"canvas-backend/util"
)

type BackgroundRequest struct {
	ProductDescription string   `json:"productDescription"`
	BrandDescription   string   `json:"brandDescription"`
	Colors             []string `json:"colors"`
}

type BackgroundResponse struct {
	GeneratedURLs []string `json:"generatedURLs"`
}

func GenerateBackgroundHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var req BackgroundRequest
	json.NewDecoder(r.Body).Decode(&req)

	// 1. Create the prompt
	bgPrompt := util.GenerateBackgroundPrompt(
		req.ProductDescription,
		req.BrandDescription,
		req.Colors,
	)

	// 2. Generate background images (DALL·E)
	imgBytesList := util.GenerateBackgroundImages(bgPrompt)

	// 3. Upload to Cloudinary
	var urls []string
	for _, imgBytes := range imgBytesList {
		u := util.UploadToCloudinary(imgBytes)
		urls = append(urls, u)
	}

	// 4. Return URLs
	json.NewEncoder(w).Encode(BackgroundResponse{
		GeneratedURLs: urls,
	})
}
