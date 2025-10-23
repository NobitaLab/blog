package controllers

import (
	"blog/server/utils"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UploadController struct {
	db *gorm.DB
}

// NewUploadController 创建上传控制器
func NewUploadController(db *gorm.DB) *UploadController {
	return &UploadController{db: db}
}

// UploadImage 处理图片上传
func (uc *UploadController) UploadImage(c *gin.Context) {
	// 获取上传的文件
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "获取上传文件失败")
		return
	}
	defer file.Close()

	// 验证文件类型
	if !isValidImageType(header) {
		utils.ErrorResponse(c, http.StatusBadRequest, "不支持的文件类型，请上传图片文件")
		return
	}

	// 验证文件大小 (5MB)
	if header.Size > 5*1024*1024 {
		utils.ErrorResponse(c, http.StatusBadRequest, "文件大小不能超过5MB")
		return
	}

	// 创建上传目录
	uploadDir := "uploads/images"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "创建上传目录失败")
		return
	}

	// 生成唯一文件名
	ext := filepath.Ext(header.Filename)
	filename := fmt.Sprintf("%d_%s%s", time.Now().Unix(), generateRandomString(8), ext)
	filepath := filepath.Join(uploadDir, filename)

	// 保存文件
	if err := saveUploadedFile(file, filepath); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "保存文件失败")
		return
	}

	// 返回文件URL
	imageURL := fmt.Sprintf("/uploads/images/%s", filename)

	utils.SuccessResponse(c, http.StatusOK, gin.H{
		"url":      imageURL,
		"imageUrl": imageURL,
		"filename": filename,
		"size":     header.Size,
	})
}

// isValidImageType 验证是否为有效的图片类型
func isValidImageType(header *multipart.FileHeader) bool {
	allowedTypes := []string{
		"image/jpeg",
		"image/jpg",
		"image/png",
		"image/gif",
		"image/webp",
	}

	contentType := header.Header.Get("Content-Type")
	for _, allowedType := range allowedTypes {
		if contentType == allowedType {
			return true
		}
	}

	// 也检查文件扩展名
	ext := strings.ToLower(filepath.Ext(header.Filename))
	allowedExts := []string{".jpg", ".jpeg", ".png", ".gif", ".webp"}
	for _, allowedExt := range allowedExts {
		if ext == allowedExt {
			return true
		}
	}

	return false
}

// saveUploadedFile 保存上传的文件
func saveUploadedFile(file multipart.File, filepath string) error {
	// 创建目标文件
	dst, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer dst.Close()

	// 复制文件内容
	_, err = io.Copy(dst, file)
	return err
}

// generateRandomString 生成随机字符串
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[time.Now().UnixNano()%int64(len(charset))]
	}
	return string(b)
}

// DeleteImage 删除图片
func (uc *UploadController) DeleteImage(c *gin.Context) {
	// 获取要删除的图片文件名
	filename := c.PostForm("filename")
	if filename == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "图片文件名不能为空")
		return
	}

	// 构建文件路径
	filePath := filepath.Join("uploads/images", filename)

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		utils.ErrorResponse(c, http.StatusNotFound, "图片文件不存在")
		return
	}

	// 删除文件
	if err := os.Remove(filePath); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "删除图片文件失败")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, gin.H{
		"message": "图片文件删除成功",
		"filename": filename,
	})
}
