package ghstorage

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	Owner        = "iteungchatbot"
	Organization = "repoulbi"
	Repository   = "data_simpelbi"
	Branch       = "main"
)

// Mengecek apakah file sudah ada atau belum
func IsFileExist(Owner, Organization, Repository, Branch, filePath, accessToken string) (bool, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s", Organization, Repository, filePath)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, fmt.Errorf("gagal membuat HTTP request: %s", err.Error())
	}

	req.Header.Set("Authorization", fmt.Sprintf("token %s", accessToken))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, fmt.Errorf("gagal melakukan HTTP request: %s", err.Error())
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		// File belum ada, bisa diunggah
		return false, nil
	} else if resp.StatusCode == http.StatusOK {
		// File sudah ada
		return true, nil
	} else {
		// Respon lainnya, ada kesalahan
		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return false, fmt.Errorf("gagal membaca respons HTTP: %s", err.Error())
		}
		return false, fmt.Errorf("gagal memeriksa file: %s", string(responseBody))
	}
}

// UploadFileToGitHub mengunggah file ke repositori GitHub menggunakan GitHub API.
func UploadFileToGitHub(Owner, Organization, Repository, Branch, filePath, accessToken string) error {
	fileExists, err := IsFileExist(Owner, Organization, Repository, Branch, filePath, accessToken)
	if err != nil {
		return err
	}

	if fileExists {
		return fmt.Errorf("File sudah ada")
	}

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("gagal membuka file: %s", err.Error())
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("gagal membaca file: %s", err.Error())
	}

	// Encode file content ke base64
	encodedContent := base64.StdEncoding.EncodeToString(fileContent)

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s", Organization, Repository, filePath)
	payload := []byte(fmt.Sprintf(`{
		"message": "Upload file %s",
		"content": "%s",
		"branch": "%s"
	}`, filePath, encodedContent, Branch))

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("gagal membuat HTTP request: %s", err.Error())
	}

	req.Header.Set("Authorization", fmt.Sprintf("token %s", accessToken))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("gagal melakukan HTTP request: %s", err.Error())
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated {
		// Tidak perlu mencetak log di sini
	} else {
		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("gagal membaca respons HTTP: %s", err.Error())
		}
		return fmt.Errorf("gagal mengunggah file: %s", string(responseBody))
	}

	return nil
}
