package Auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type UserInfo struct {
	Level      int         `json:"level"`
	NickName   string      `json:"nickName"`
	Sex        interface{} `json:"sex"`
	Mobile     interface{} `json:"mobile"`
	Pic        interface{} `json:"pic"`
	BirthDate  interface{} `json:"birthDate"`
	UserID     string      `json:"userId"`
	Score      int         `json:"score"`
	Balance    float64     `json:"balance"`
	UserMobile string      `json:"userMobile"`
	LevelType  int         `json:"levelType"`
	Growth     int         `json:"growth"`
	Status     int         `json:"status"`
	Username   string      `json:"username"`
}

// VerifyAuth
// @Description: 鉴权中间件
// @return gin.HandlerFunc
func VerifyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		RequestURI := c.Request.RequestURI

		fmt.Println(RequestURI)

		// 免验证接口
		allowInterface := []string{
			// "/",
		}

		for _, v := range allowInterface {
			if RequestURI == v {
				return
			}
		}

		authorization := c.Request.Header.Get("Authorization")

		if authorization == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "身份验证失败",
				"code":    1,
				"data":    "",
			})
			c.Abort()
			return
		}

		b2bPcPath := viper.GetString("b2b-pc-path")
		// 验证token

		client := &http.Client{}
		req, _ := http.NewRequest("GET", b2bPcPath, nil)
		req.Header.Add("Authorization", authorization)
		resp, _ := client.Do(req)

		if resp.StatusCode != 200 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "token error",
				"code":    1,
				"data":    "",
			})
			c.Abort()
			return
		}

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "token error",
				"code":    1,
				"data":    "",
			})
			c.Abort()
			return
		}

		if body == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "token error",
				"code":    1,
				"data":    "",
			})
			c.Abort()
			return
		}

		var res UserInfo
		if err := json.Unmarshal(body, &res); err != nil {
			// 返回错误信息
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
				"code":    1,
				"data":    "",
			})
			c.Abort()
			return
		}

		// 存储上下文
		c.Set("user_id", res.UserID)

		c.Next()
	}
}
