package helpers

import(
	"errors"
	"github.com/gin-gonic/gin"
)


func CheckUserType(c *gin.Context,role string) (e error){
	userType := c.GetString("user_type")
	e = nil
	if userType != role{
		e = errors.New("Unauthorize to access this resource")
		return e
	}
	return e
}


func MatchUserTypetoUid(c *gin.Context,userId string) (e error){
	userType :=c.GetString("user_type")
	uid := c.GetString("uid")
	e = nil

	if userType=="USER" && uid!=userId{
		e =errors.New("Unauthorize to access this resource")
		return e
	}
	e = CheckUserType(c,userType)
	return e
}


