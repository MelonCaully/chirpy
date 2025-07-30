package auth

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "testpassword123"
	hashed, err := HashPassword(password)

	if err != nil {
		t.Fatalf("HashPassword returned an error: %v", err)
	}

	if hashed == "" {
		t.Fatal("Expected hashed password, got empty string")
	}

	if hashed == password {
		t.Fatal("Hashed password should not match the original password")
	}
}

func TestCheckPasswordHash_CorrectPassword(t *testing.T) {
	password := "correct-password"
	hashed, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword returned an error: %v", err)
	}

	err = CheckPasswordHash(password, hashed)
	if err != nil {
		t.Fatal("Hashed password does not match the original")
	}
}

func TestCheckPasswordHash_WrongPassword(t *testing.T) {
	password := "correct-password"
	wrongPassword := "wrong-password"

	hashed, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword returned an error: %v", err)
	}

	err = CheckPasswordHash(wrongPassword, hashed)
	if err == nil {
		t.Fatal("Expected CheckPassword to fail for wrong password, but got no error")
	}
}
