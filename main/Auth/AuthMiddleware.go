package Auth

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"strings"
	"time"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		if MissAuth(c.Request.URL.Path) {
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

		jwtSecretKey := viper.GetString("jwt_config.secret_key")
		// parse token, get the user and role info
		claims, err := ParseJwtToken(token, jwtSecretKey)
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
		c.Set("LoginUser", claims.UserName)
		c.Next()
	}
}

func MissAuth(path string) bool {
	return strings.HasSuffix(path, "login")
}
