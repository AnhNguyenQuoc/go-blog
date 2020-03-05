package lib

import (
	"log"
	"os"
	"testing"
)

func TestGetPort(t *testing.T) {
	os.Setenv("PORT", "3000")
	value := GetPort()
	want := ":3000"

	if value != want {
		t.Errorf("PORT error: want %v but take %v", want, value)
	}
}

func TestBcrypt(t *testing.T) {
	mypwd := "123456"
	hashPassword, _ := HashPassword(mypwd)
	if auth := AuthenticatePassword(mypwd, hashPassword); !auth {
		t.Error("Error Bcrypt algothym")
	}
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}