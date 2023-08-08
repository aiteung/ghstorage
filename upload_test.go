package ghstorage

import (
	"fmt"
	"testing"
)

func TestUpload(t *testing.T) {
	// Ganti dengan informasi Anda
	owner := "valenrio66"
	org := "repoulbi"
	repo := "data_simpelbi"
	branch := "main"
	filePath := "presensi.txt" // Ganti dengan path file yang ingin Anda unggah
	accessToken := "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"

	err := UploadFileToGitHub(owner, org, repo, branch, filePath, accessToken)
	if err != nil {
		fmt.Println(err.Error())
	}
}
