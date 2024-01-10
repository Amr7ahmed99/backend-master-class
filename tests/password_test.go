package tests

import (
	"backend-master-class/util"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password1 := util.RandomString(6)
	hashedPassword1, err := util.HashPassword(password1)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)
	err = util.CheckPassword(password1, hashedPassword1)
	require.NoError(t, err)
	err = util.CheckPassword(util.RandomString(6), hashedPassword1)
	require.Error(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	password2 := util.RandomString(6)
	hashedPassword2, err := util.HashPassword(password2)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword2)
	err = util.CheckPassword(password2, hashedPassword2)
	require.NoError(t, err)
	err = util.CheckPassword(util.RandomString(6), hashedPassword2)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	require.NotEqual(t, hashedPassword1, hashedPassword2)

}
