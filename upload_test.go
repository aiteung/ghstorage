package ghstorage

import (
	"fmt"
	"os"
	"testing"
)

func TestUploadFileToGitHub(t *testing.T) {
	key := "ACCESSTOKEN"
	// Ganti dengan informasi Anda
	accessToken := os.Getenv(key)
	fmt.Println(accessToken)

	// Ganti dengan informasi file yang ingin diuji
	filePath := "presensi.txt"

	err := UploadFileToGitHub(Owner, Organization, Repository, Branch, filePath, accessToken)
	if err != nil {
		t.Errorf("Error while uploading file: %v", err)
	}
}
