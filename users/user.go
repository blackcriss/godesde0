package users

import (
	"fmt"
	"time"

	"github.com/blackcriss/godesde0/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(10, "Cristian", time.Now(), true)
	fmt.Println(u)
}
