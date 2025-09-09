package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"blog/server/models"
	"blog/server/utils"
)

// BlogController 博客控制器
type BlogController struct {
	DB *gorm.DB
}

// NewBlogController 创建博客控制器实例
func NewBlogController(db *gorm.DB) *BlogController {
	return &BlogController{DB: db}
}

// GetAllBlogs 获取所有博客文章
// @Summary 获取所有博客文章
// @Description 获取所有博客文章列表，支持分页
// @Tags 博客文章
// @Accept json
// @Produce json
// @Param page query int false "页码，默认为1"
// @Param limit query int false "每页数量，默认为10"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /api/blogs [get]
func (bc *BlogController) GetAllBlogs(c *gin.Context) {
	// 获取查询参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	// 计算偏移量
	offset := (page - 1) * limit

	// 查询博客列表
	var blogs []models.Blog
	var total int64

	// 查询数据库
	if err := bc.DB.Model(&models.Blog{}).Count(&total).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "获取博客列表失败")
		return
	}

	if err := bc.DB.Order("created_at DESC").Offset(offset).Limit(limit).Find(&blogs).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "获取博客列表失败")
		return
	}

	// 返回结果
	utils.SuccessListResponse(c, http.StatusOK, blogs, total, page, limit)
}

// GetBlogByID 获取单篇博客文章
// @Summary 获取单篇博客文章
// @Description 获取指定ID的博客文章详情
// @Tags 博客文章
// @Accept json
// @Produce json
// @Param id path int true "博客文章ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/blogs/{id} [get]
func (bc *BlogController) GetBlogByID(c *gin.Context) {
	// 获取博客ID
	id := c.Param("id")

	// 查询博客
	var blog models.Blog
	if err := bc.DB.First(&blog, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "博客不存在")
		return
	}

	// 返回结果
	utils.SuccessResponse(c, http.StatusOK, blog)
}

// CreateBlog 创建新的博客文章
// @Summary 创建新的博客文章
// @Description 创建一篇新的博客文章
// @Tags 博客文章
// @Accept json
// @Produce json
// @Param blog body models.Blog true "博客文章数据"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/blogs [post]
func (bc *BlogController) CreateBlog(c *gin.Context) {
	// 绑定请求体
	var blog models.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "请求参数错误")
		return
	}

	// 验证参数
	if blog.Title == "" || blog.Content == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "标题和内容不能为空")
		return
	}

	// 创建博客
	if err := bc.DB.Create(&blog).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "创建博客失败")
		return
	}

	// 返回结果
	utils.SuccessResponse(c, http.StatusCreated, blog)
}

// UpdateBlog 更新博客文章
// @Summary 更新博客文章
// @Description 更新指定ID的博客文章
// @Tags 博客文章
// @Accept json
// @Produce json
// @Param id path int true "博客文章ID"
// @Param blog body models.Blog true "博客文章更新数据"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/blogs/{id} [put]
func (bc *BlogController) UpdateBlog(c *gin.Context) {
	// 获取博客ID
	id := c.Param("id")

	// 查询博客
	var blog models.Blog
	if err := bc.DB.First(&blog, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "博客不存在")
		return
	}

	// 绑定请求体
	var updateData models.Blog
	if err := c.ShouldBindJSON(&updateData); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "请求参数错误")
		return
	}

	// 验证参数
	if updateData.Title == "" || updateData.Content == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "标题和内容不能为空")
		return
	}

	// 更新博客
	blog.Title = updateData.Title
	blog.Content = updateData.Content
	if updateData.Author != "" {
		blog.Author = updateData.Author
	}

	if err := bc.DB.Save(&blog).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "更新博客失败")
		return
	}

	// 返回结果
	utils.SuccessResponse(c, http.StatusOK, blog)
}

// DeleteBlog 删除博客文章
// @Summary 删除博客文章
// @Description 删除指定ID的博客文章
// @Tags 博客文章
// @Accept json
// @Produce json
// @Param id path int true "博客文章ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/blogs/{id} [delete]
func (bc *BlogController) DeleteBlog(c *gin.Context) {
	// 获取博客ID
	id := c.Param("id")

	// 检查博客是否存在
	var blog models.Blog
	if err := bc.DB.First(&blog, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "博客不存在")
		return
	}

	// 删除博客
	if err := bc.DB.Delete(&models.Blog{}, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "删除博客失败")
		return
	}

	// 返回结果
	utils.SuccessResponse(c, http.StatusOK, gin.H{"message": "删除成功"})
}