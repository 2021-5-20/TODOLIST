package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"www.github.com/2021-5-20/t_d_list/config"
	"www.github.com/2021-5-20/t_d_list/models"
)

func CreateMemoHandler(c *gin.Context) {
	type request struct {
		Content string `json:"content" binding:"required"`
	}

	var input request
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//从jwt中取user_id
	userID := c.GetUint("user_id")
	memo := models.Memo{
		Content: input.Content,
		UserID:  userID,
		// CreatedAt/UpdatedAt GORM 会自动维护
	}

	db := config.GetDB()
	if err := db.Create(&memo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建备忘录失败"})
		return
	}
	//db.Preload("User").First(&memo, memo.ID) 可以选择是否需要看到user
	c.JSON(http.StatusCreated, gin.H{
		"message": "备忘录创建成功",
		"memo":    memo,
	})
}

func GetAllMemoHandler(c *gin.Context) {
	// 1. 从数据库中获取所有 Memo
	userID := c.GetUint("user_id")
	var memos []models.Memo
	db := config.GetDB()
	db.Where("user_id = ?", userID).Order("created_at desc").Find(&memos)
	// 2. 返回所有 Memo
	c.JSON(http.StatusOK, memos)

}

// GetMemoHandler —— 根据 ID 查询一条 Memo
func GetMemoHandler(c *gin.Context) {
	// 1. 解析 ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "非法的备忘录 ID"})
		return
	}
	userID := c.GetUint("user_id")
	// 2. 查询
	var memo models.Memo
	db := config.GetDB()
	if err := db.Preload("User").First(&memo, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "备忘录不存在"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
		}
		return
	}
	if userID != memo.UserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "权限不足"})
		return
	}
	// 3. 返回
	c.JSON(http.StatusOK, memo)
}

// UpdateMemoHandler —— 更新一条 Memo 的内容
func UpdateMemoHandler(c *gin.Context) {
	// 1. 解析 ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "非法的备忘录 ID"})
		return
	}
	userID := c.GetUint("user_id")
	db := config.GetDB()
	// 2. 取出已有记录
	var memo models.Memo
	if err := db.First(&memo, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "备忘录不存在"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
		}
		return
	}
	if userID != memo.UserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "权限不足"})
		return
	}
	// 3. 绑定更新字段
	type request struct {
		Content string `json:"content" binding:"required"`
	}
	var input request
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 4. 应用修改并保存
	memo.Content = input.Content
	if err := db.Save(&memo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失败"})
		return
	}

	// 5. 返回最新记录
	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
		"memo":    memo,
	})
}

// DeleteMemoHandler —— 删除一条 Memo
func DeleteMemoHandler(c *gin.Context) {
	// 1. 解析 ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "非法的备忘录 ID"})
		return
	}
	userID := c.GetUint("user_id")

	db := config.GetDB()
	// 2. 确认存在
	var memo models.Memo
	if err := db.First(&memo, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "备忘录不存在"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
		}
		return
	}
	if userID != memo.UserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "权限不足"})
		return
	}
	// 3. 执行删除
	if err := db.Delete(&memo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除失败"})
		return
	}

	// 4. 返回结果
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}
