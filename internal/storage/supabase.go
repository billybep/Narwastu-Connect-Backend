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

// Member Avatar (Anggota Jemaat)
func (c *Client) UploadAvatar(file multipart.File, fileHeader *multipart.FileHeader, memberID uint) (string, error) {
	// ✅ Buat nama file unik
	ext := filepath.Ext(fileHeader.Filename)
	fileName := fmt.Sprintf("%d_%d%s", memberID, time.Now().Unix(), ext)
	filePath := fmt.Sprintf("avatars/%s", fileName)

	// ✅ Deteksi content type dari ekstensi
	contentType := "application/octet-stream"
	extLower := strings.ToLower(ext)

	switch extLower {
	case ".jpg", ".jpeg":
		contentType = "image/jpeg"
	case ".png":
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

func (c *Client) DeleteFile(publicURL string) error {
	// Contoh publicURL:
	// https://xxxx.supabase.co/storage/v1/object/public/DevelopNC/avatars/42_123.jpg

	parts := strings.Split(publicURL, fmt.Sprintf("/%s/", c.bucket))
	if len(parts) < 2 {
		return fmt.Errorf("invalid file URL")
	}
	filePath := parts[1]

	// 🗑️ Hapus file dari Supabase Storage
	_, err := c.storage.RemoveFile(c.bucket, []string{filePath})
	if err != nil {
		return fmt.Errorf("failed to delete old file: %v", err)
	}

	return nil
}

// Upload Warta Jemaat
func (c *Client) UploadWartaJemaat(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	ext := filepath.Ext(fileHeader.Filename)
	if ext == "" {
		ext = ".pdf" // default
	}

	now := time.Now()
	year := now.Year()
	month := int(now.Month())

	// Buat random suffix (pakai UnixNano biar unik)
	fileName := fmt.Sprintf("warta-%d-%02d-%d%s", year, month, time.Now().UnixNano(), ext)
	filePath := fmt.Sprintf("wartajemaat/%s", fileName)

	contentType := "application/pdf"
	extLower := strings.ToLower(ext)
	switch extLower {
	case ".jpg", ".jpeg":
		contentType = "image/jpeg"
	case ".png":
		contentType = "image/png"
	}

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, file); err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	// Upload ke Supabase
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

	// Public URL
	publicURL := fmt.Sprintf("%s/object/public/%s/%s", c.baseURL, c.bucket, filePath)
	return publicURL, nil
}
