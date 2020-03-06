package lib

import (
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
	hashPassword := "$2a$10$YryYXVCk5KE2IoZR1e41WOEE6RMSxduBVb6XgpqqMojsduWTEmH02"
	if auth := AuthenticatePassword(mypwd, hashPassword); !auth {
		t.Error("Error Bcrypt algothym")
	}
}
