package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"blog/server/models"
	"blog/server/utils"
)

// UserController 用户控制器
type UserController struct {
	DB *gorm.DB
}

// NewUserController 创建用户控制器实例
func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

// Register 用户注册
// @Summary 用户注册
// @Description 创建新用户账号
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param user body models.User true "用户注册信息"（不含role字段）
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/users/register [post]
func (uc *UserController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "请求参数错误")
		return
	}

	// 处理Email字段，如果为空字符串则设置为nil
	if user.Email != nil && *user.Email == "" {
		user.Email = nil
	}

	// 验证参数
	if user.Username == "" || user.Password == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "用户名和密码不能为空")
		return
	}

	// 检查用户名是否已存在
	var existingUser models.User
	if err := uc.DB.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "用户名已存在")
		return
	}

	// 加密密码
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "注册失败")
		return
	}

	// 设置用户角色为普通用户
	user.Password = hashedPassword
	user.Role = "user"

	// 创建用户
	if err := uc.DB.Create(&user).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "注册失败")
		return
	}

	// 返回结果（不包含密码）
	user.Password = ""
	utils.SuccessResponse(c, http.StatusCreated, user)
}

// Login 用户登录
// @Summary 用户登录
// @Description 用户登录获取JWT令牌
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param credentials body object true "登录凭证" {"username":"string", "password":"string"}
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/users/login [post]
func (uc *UserController) Login(c *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "请求参数错误")
		return
	}

	// 查找用户
	var user models.User
	if err := uc.DB.Where("username = ?", credentials.Username).First(&user).Error; err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "用户名或密码错误")
		return
	}

	// 验证密码
	if !utils.CheckPassword(credentials.Password, user.Password) {
		utils.ErrorResponse(c, http.StatusUnauthorized, "用户名或密码错误")
		return
	}

	// 生成JWT令牌
	token, err := utils.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "登录失败")
		return
	}

	// 返回令牌和用户信息
	user.Password = ""
	utils.SuccessResponse(c, http.StatusOK, gin.H{
		"token": token,
		"user":  user,
	})
}

// Logout 用户注销
// @Summary 用户注销
// @Description 用户注销（客户端应删除本地存储的令牌）
// @Tags 用户管理
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Router /api/users/logout [post]
func (uc *UserController) Logout(c *gin.Context) {
	// 在服务器端，我们只需要返回成功响应
	// 客户端应该负责删除本地存储的JWT令牌
	utils.SuccessResponse(c, http.StatusOK, gin.H{"message": "注销成功"})
}

// GetCurrentUser 获取当前用户信息
// @Summary 获取当前用户信息
// @Description 通过JWT令牌获取当前登录用户的信息
// @Tags 用户管理
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]string
// @Router /api/users/me [get]
func (uc *UserController) GetCurrentUser(c *gin.Context) {
	// 从上下文中获取用户ID
	userID, exists := c.Get("userID")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "未授权访问")
		return
	}

	// 获取用户信息
	var user models.User
	if err := uc.DB.First(&user, userID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "用户不存在")
		return
	}

	// 返回用户信息（不包含密码）
	user.Password = ""
	utils.SuccessResponse(c, http.StatusOK, user)
}

// GetAllUsers 获取所有用户信息（仅管理员可访问）
// @Summary 获取所有用户信息
// @Description 获取所有用户的详细信息
// @Tags 用户管理
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/admin/users [get]
func (uc *UserController) GetAllUsers(c *gin.Context) {
	var users []models.User
	if err := uc.DB.Find(&users).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "获取用户列表失败")
		return
	}

	// 移除密码信息
	for i := range users {
		users[i].Password = ""
	}

	utils.SuccessResponse(c, http.StatusOK, users)
}

// UpdateUserInfo 更新用户信息（仅管理员可访问）
// @Summary 更新用户信息
// @Description 更新指定用户的信息，包括角色
// @Tags 用户管理
// @Security ApiKeyAuth
// @Param id path int true "用户ID"
// @Param user body models.User true "用户更新信息"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/admin/users/{id} [put]
func (uc *UserController) UpdateUserInfo(c *gin.Context) {
	// 获取用户ID
	id := c.Param("id")

	// 查找用户
	var user models.User
	if err := uc.DB.First(&user, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "用户不存在")
		return
	}

	// 绑定请求体
	var updateData struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Role     string `json:"role"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "请求参数错误")
		return
	}

	// 检查并更新用户名
	if updateData.Username != "" {
		// 检查新用户名是否已存在
		var existingUser models.User
		if err := uc.DB.Where("username = ? AND id != ?", updateData.Username, id).First(&existingUser).Error; err == nil {
			utils.ErrorResponse(c, http.StatusBadRequest, "用户名已存在")
			return
		}
		user.Username = updateData.Username
	}

	// 更新其他信息
	if updateData.Email != "" {
		user.Email = &updateData.Email
	} else {
		// 如果Email为空字符串，将其设置为nil
		user.Email = nil
	}

	// 只有管理员可以更改角色
	if updateData.Role != "" {
		// 确保角色只能是admin或user
		if updateData.Role != "admin" && updateData.Role != "user" {
			utils.ErrorResponse(c, http.StatusBadRequest, "角色只能是admin或user")
			return
		}
		user.Role = updateData.Role
	}

	// 保存更新
	if err := uc.DB.Save(&user).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "更新用户信息失败")
		return
	}

	// 返回结果（不包含密码）
	user.Password = ""
	utils.SuccessResponse(c, http.StatusOK, user)
}