package handlers

import (
	"fmt"
	"image"
	"net/http"

	_ "image/gif"
	_ "image/jpeg"
	"image/png"
)

func UploadImage(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Error uploading file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Open the uploaded image
	uploadedImage, _, err := image.Decode(file)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error decoding image", http.StatusInternalServerError)
		return
	}

	// Provide download link to the user
	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, uploadedImage)
}
