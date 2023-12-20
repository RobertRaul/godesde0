package users

import (
	"fmt"
	"time"

	"github.com/robertraul/godesde0/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(1, "Robertraul", time.Now(), true)
	fmt.Println(u)
}
