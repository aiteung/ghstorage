package ghstorage

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

// UploadFileToGitHub mengunggah file ke repositori GitHub menggunakan GitHub API.
func UploadFileToGitHub(repoOwner, repoName, filePath, accessToken string) error {
	// Dapatkan path absolut dari direktori kerja saat ini
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return fmt.Errorf("gagal mendapatkan path absolut: %w", err)
	}

	// Baca file ke dalam []byte
	fileContent, err := ioutil.ReadFile(absPath)
	if err != nil {
		return fmt.Errorf("gagal membaca file: %w", err)
	}

	// Konversi konten file ke dalam format Base64
	fileContentBase64 := base64.StdEncoding.EncodeToString(fileContent)

	// Buat URL API untuk mengunggah file
	uploadURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s", repoOwner, repoName, absPath)

	// Buat payload JSON untuk permintaan API
	payload := []byte(fmt.Sprintf(`{
		"message": "Unggah file melalui API GitHub",
		"content": "%s"
	}`, fileContentBase64))

	// Buat permintaan HTTP POST ke GitHub API
	req, err := http.NewRequest("PUT", uploadURL, bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("gagal membuat permintaan: %w", err)
	}

	// Atur otorisasi menggunakan token akses
	req.Header.Set("Authorization", "token "+accessToken)

	// Atur tipe konten sebagai aplikasi JSON
	req.Header.Set("Content-Type", "application/json")

	// Buat klien HTTP dan kirim permintaan
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("gagal mengirim permintaan: %w", err)
	}
	defer resp.Body.Close()

	// Baca dan tampilkan respons dari API
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("gagal membaca respons: %w", err)
	}

	fmt.Println("Status Code:", resp.Status)
	fmt.Println("Respon Body:", string(respBody))

	return nil
}
