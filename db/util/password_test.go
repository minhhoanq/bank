package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)

	//create password hasd
	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	// match password and password hash is correct
	err = ComparePassword(password, hashedPassword)
	require.NoError(t, err)

	// match password and password hash is not correct
	passwordWrong := RandomString(6)
	err = ComparePassword(passwordWrong, hashedPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
