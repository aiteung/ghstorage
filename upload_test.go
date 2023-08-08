package ghstorage

import (
	"testing"
)

func TestUploadFileToGitHub(t *testing.T) {
	// Ganti dengan informasi Anda
	owner := "valenrio66"
	accessToken := "your_access_token"

	// Ganti dengan informasi file yang ingin diuji
	filePath := "test.txt"

	err := UploadFileToGitHub(owner, filePath, accessToken)
	if err != nil {
		t.Errorf("Error while uploading file: %v", err)
	}
}
