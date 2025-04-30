package v2

import (
	"fmt"

	"github.com/unik0105/go-demo2/modules/model"
)

var v2Template string = "v2 - > name : %s age : %d "

func GetStudentName(u *model.User) string {
	return fmt.Sprintf(v2Template, u.Name, u.Age)
}
