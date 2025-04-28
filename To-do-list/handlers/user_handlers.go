package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"www.github.com/2021-5-20/t_d_list/config"
	"www.github.com/2021-5-20/t_d_list/models"
)

func GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{ // 有效负载
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(), // 1天后过期
		"iat":     time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.JwtSecret)
}

func HelloGetfunc(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}
func LoginHandler(c *gin.Context) {
	var loginUser models.User
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		fmt.Println("绑定失败：", err.Error())
		c.String(http.StatusBadRequest, "请输入账号和密码")
		return
	}

	db := config.GetDB()
	var foundUser models.User

	if err := db.Where("username = ?", loginUser.Username).First(&foundUser).Error; err != nil {
		c.String(http.StatusUnauthorized, "用户名或密码错误")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(loginUser.Password)); err != nil {
		c.String(http.StatusUnauthorized, "用户名或密码错误")
		return
	}
	token, err := GenerateToken(foundUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成 Token 失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"恭喜你":   "登录成功",
	})

}
func ProfileHandler(c *gin.Context) {
	userID := c.GetUint("user_id") // 取出中间件设置的 user_id
	c.JSON(http.StatusOK, gin.H{
		"message": "这是受保护的个人中心",
		"user_id": userID,
	})
}
func HandlePostFormStruct(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()
	db := config.GetDB()
	result := db.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "用户名重复"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "用户创建成功",
		"ID":       user.ID,
		"Username": user.Username,
		"Created":  user.CreatedAt,
	})
}
