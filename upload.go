package ghstorage

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	Organization = "repoulbi"
	Repository   = "data_simpelbi"
	Branch       = "main"
)

// Fungsi ini melakukan otentikasi OAuth dan mengembalikan token akses yang valid.
func AuthenticateOAuth(clientID, clientSecret, code, redirectURL string) (string, error) {
	// Menyusun payload untuk mendapatkan token akses.
	payload := fmt.Sprintf("client_id=%s&client_secret=%s&code=%s&redirect_uri=%s",
		clientID, clientSecret, code, redirectURL)

	url := "https://github.com/login/oauth/access_token"

	req, err := http.NewRequest("POST", url, bytes.NewBufferString(payload))
	if err != nil {
		return "", fmt.Errorf("gagal membuat HTTP request: %s", err.Error())
	}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("gagal melakukan HTTP request: %s", err.Error())
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("gagal mendapatkan token akses")
	}

	var responseBody struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		return "", fmt.Errorf("gagal membaca respons JSON: %s", err.Error())
	}

	return responseBody.AccessToken, nil
}

// Mengecek apakah file sudah ada atau belum
func IsFileExist(owner, Organization, Repository, Branch, filePath, accessToken string) (bool, error) {
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
		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return false, fmt.Errorf("gagal membaca respons HTTP: %s", err.Error())
		}
		return false, fmt.Errorf("gagal memeriksa file: %s", string(responseBody))
	}
}

// UploadFileToGitHub mengunggah file ke repositori GitHub menggunakan GitHub API.
func UploadFileToGitHub(owner, filePath, accessToken string) error {
	fileExists, err := IsFileExist(owner, Organization, Repository, Branch, filePath, accessToken)
	if err != nil {
		return err
	}

	if fileExists {
		return fmt.Errorf("File sudah ada")
	}

	fileContent, err := ioutil.ReadFile(filePath)
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
		return fmt.Errorf("gagal mengunggah file")
	}

	return nil
}
