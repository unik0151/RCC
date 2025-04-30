package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/unik0105/go-demo2/modules/model"
)

var v1Template string = "v1 - > name : %s age : %d "

func GetStudentName(c *gin.Context) string {
	var u model.User
	c.Bind(&u)

	return fmt.Sprintf(v1Template, u.Name, u.Age)
}
