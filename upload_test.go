package ghstorage

import (
	"fmt"
	"testing"
)

func TestUpload(t *testing.T) {
	// Ganti dengan informasi Anda
	org := "repoulbi"
	repo := "data_simpelbi"
	branch := "main"
	filePath := "coba.txt" // Ganti dengan path file yang ingin Anda unggah
	accessToken := "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"

	err := UploadFileToGitHub("valenrio66", org, repo, branch, filePath, accessToken)
	if err != nil {
		fmt.Println(err.Error())
	}
}
