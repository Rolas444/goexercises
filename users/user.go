package users

import (
	"fmt"
	"time"

	"github.com/rrq/golang/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(10, "Pablo", time.Now(), true)
	fmt.Println(u)
}
