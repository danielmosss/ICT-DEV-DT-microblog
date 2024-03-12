package tests

import (
	"github.com/che-ict/DEV-DT-Microblog/repositories"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateUser(t *testing.T) {

	// Given
	// 2 gebruikers
	// gebruiker 1 username TestDummy, gebruiker 2 username TestDummy
	err1 := repositories.CreateUser("TestDummy", "Test1", "TestDummy")
	err2 := repositories.CreateUser("TestDummy", "Test2", "TestDummy")

	// When
	// gebruiker 1 wilt TestDummy aanmaken
	// gebruiker 2 wilt TestDummy aanmaken

	// Then
	// gebruiker 1 wordt aangemaakt
	// gebruiker 2 wordt niet aangemaakt
	assert.Nil(t, err1)
	assert.NotNil(t, err2)
}
