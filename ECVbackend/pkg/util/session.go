package util
import(
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
    "ecvbackend/pkg"
)
var cookieSecret = []byte(setting.CookieSecret)

func EnableCookieSession() gin.HandlerFunc {
    store := cookie.NewStore([]byte(cookieSecret))
    return sessions.Sessions("CookieSession", store)
}

func SaveAuthSession(c *gin.Context, id string) {
    session := sessions.Default(c)
    session.Set("username", id)
    session.Save()
}

func ClearAuthSession(c *gin.Context) {
    session := sessions.Default(c)
    session.Clear()
    session.Save()
}

func HasSession(c *gin.Context) bool {
    session := sessions.Default(c)
    if sessionValue := session.Get("username"); sessionValue == nil {
        return false
    }
    return true
}

func GetSessionUserName(c *gin.Context) string {
    session := sessions.Default(c)
    sessionValue := session.Get("username")
    return sessionValue.(string)
}

func GetUserSession(c *gin.Context) map[string]interface{} {
    hasSession := HasSession(c)
    userName := ""
    if hasSession {
        userName = GetSessionUserName(c)
    }
    data := make(map[string]interface{})
    data["hasSession"] = hasSession
    data["userName"] = userName
    return data
}