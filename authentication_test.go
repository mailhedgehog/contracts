package contracts

import (
	"github.com/mailhedgehog/gounit"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestCreatePasswordHash_matchPassword(t *testing.T) {
	pass := "foo-bar"
	hashedPass, err := CreatePasswordHash(pass)

	(*gounit.T)(t).AssertNotError(err)

	err = bcrypt.CompareHashAndPassword(hashedPass, []byte(pass))

	(*gounit.T)(t).AssertNotError(err)
}

func TestCreatePasswordHashError_noMatchPassword(t *testing.T) {
	pass := "foo-bar"
	hashedPass, err := CreatePasswordHash(pass)

	(*gounit.T)(t).AssertNotError(err)

	err = bcrypt.CompareHashAndPassword(hashedPass, []byte("baz"))

	(*gounit.T)(t).ExpectError(err)
	(*gounit.T)(t).AssertEqualsString(
		"crypto/bcrypt: hashedPassword is not the hash of the given password",
		err.Error(),
	)
}
