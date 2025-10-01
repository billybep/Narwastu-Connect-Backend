package supabase

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	storage_go "github.com/supabase-community/storage-go"
)

type Client struct {
	storage *storage_go.Client
	bucket  string
	baseURL string
}

func NewClient(url, key, bucket string) *Client {
	client := storage_go.NewClient(url, key, nil)
	return &Client{
		storage: client,
		bucket:  bucket,
		baseURL: url,
	}
}

func (c *Client) UploadAvatar(file multipart.File, fileHeader *multipart.FileHeader, memberID uint) (string, error) {
	// ✅ Buat nama file unik
	ext := filepath.Ext(fileHeader.Filename)
	fileName := fmt.Sprintf("%d_%d%s", memberID, time.Now().Unix(), ext)
	filePath := fmt.Sprintf("avatars/%s", fileName)

	// ✅ Deteksi content type dari ekstensi
	contentType := "application/octet-stream"
	extLower := strings.ToLower(ext)
	if extLower == ".jpg" || extLower == ".jpeg" {
		contentType = "image/jpeg"
	} else if extLower == ".png" {
		contentType = "image/png"
	}

	// ✅ Baca isi file ke buffer
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, file); err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	// ✅ Upload ke Supabase Storage dengan ContentType & Upsert
	_, err := c.storage.UploadFile(
		c.bucket,
		filePath,
		&buf,
		storage_go.FileOptions{
			ContentType: &contentType,
			Upsert:      func(b bool) *bool { return &b }(true),
		},
	)
	if err != nil {
		return "", fmt.Errorf("upload to supabase failed: %v", err)
	}

	// ✅ Buat URL publik
	publicURL := fmt.Sprintf("%s/object/public/%s/%s",
		c.baseURL,
		c.bucket,
		filePath,
	)

	return publicURL, nil
}
