package util

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

type MLImageRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Steps  int    `json:"steps"`
	N      int    `json:"num_images"`
}

type MLImageResponse struct {
	Status string   `json:"status"`
	Images []string `json:"images"` // base64 strings
}

// Add utility functions here

// UploadToCloudinary uploads image bytes to Cloudinary and returns the URL.
// You need to set your Cloudinary credentials as environment variables.
func UploadToCloudinary(imgBytes []byte) string {
	cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")
	if cloudName == "" || apiKey == "" || apiSecret == "" {
		return ""
	}

	url := fmt.Sprintf("https://api.cloudinary.com/v1_1/%s/image/upload", cloudName)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_ = writer.WriteField("upload_preset", "ml_default")
	part, _ := writer.CreateFormFile("file", "background.png")
	_, _ = io.Copy(part, bytes.NewReader(imgBytes))
	writer.Close()

	req, _ := http.NewRequest("POST", url, body)
	req.SetBasicAuth(apiKey, apiSecret)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	var result struct {
		SecureURL string `json:"secure_url"`
	}
	json.NewDecoder(resp.Body).Decode(&result)
	return result.SecureURL
}

func GenerateBackgroundImages(prompt string) ([][]byte, error) {

	reqBody := MLImageRequest{
		Model:  "flux-2-dev",
		Prompt: prompt,
		Width:  1024,
		Height: 1024,
		Steps:  28,
		N:      4,
	}

	body, _ := json.Marshal(reqBody)

	httpReq, _ := http.NewRequest("POST",
		"https://api.models.lab/v1/image/generate",
		bytes.NewBuffer(body),
	)

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("x-api-key", os.Getenv("MODELSLAB_API_KEY"))

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, _ := ioutil.ReadAll(resp.Body)

	var parsed MLImageResponse
	json.Unmarshal(respBytes, &parsed)

	images := make([][]byte, 0)
	for _, b64 := range parsed.Images {
		decoded, _ := base64.StdEncoding.DecodeString(b64)
		images = append(images, decoded)
	}

	return images, nil
}
