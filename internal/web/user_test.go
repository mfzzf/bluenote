package web

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	pwd := "123456"
	password, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	fmt.Println(string(password))
	err = bcrypt.CompareHashAndPassword(password, []byte(pwd))
	assert.NoError(t, err)
}
