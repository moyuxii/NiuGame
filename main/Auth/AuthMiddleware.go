package Auth

import (
	"NiuGame/main/Config"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// 免登录接口列表
var notAuthArr = map[string]string{"/user/login": "1", "/api/user/get1": "1"}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		inWhite := notAuthArr[c.Request.URL.Path]
		if inWhite == "1" {
			return
		}

		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": -1, "msg": "请求未携带token，无权限访问"})
			c.Abort()
			return
		}
		if strings.HasPrefix(token, "Bearer") {
			token = strings.Split(token, " ")[1]
		}

		config := Config.GetConfig()
		// parse token, get the user and role info
		claims, err := ParseJwtToken(token, config.JwtConfig.SecretKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"status": -1, "msg": err.Error()})
			c.Abort()
			return
		}
		//token超时
		if time.Now().Unix() > claims.StandardClaims.ExpiresAt {
			c.JSON(http.StatusUnauthorized, gin.H{"status": -1, "msg": "token过期"})
			c.Abort() //阻止执行
			return
		}

		// 继续交由下一个路由处理,并将解析出的信息传递下去,可以在接口里获取：jwtUser, _ := context.Get("jwtUser")
		c.Set("jwtUser", claims)
		c.Next()
	}
}
