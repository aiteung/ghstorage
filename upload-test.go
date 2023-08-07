package ghstorage

import (
	"fmt"
	"testing"
)

func TestUpload(t *testing.T) {
	repoOwner := "repoulbi"
	repoName := "data_simpelbi"
	filePath := "test.txt"
	accessToken := "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"

	err := UploadFileToGitHub(repoOwner, repoName, filePath, accessToken)
	if err != nil {
		fmt.Println("Gagal mengunggah file:", err)
		return
	}

	fmt.Println("File berhasil diunggah ke GitHub!")
}
