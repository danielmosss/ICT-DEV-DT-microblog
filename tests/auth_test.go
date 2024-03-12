package tests

import (
	"context"
	"github.com/che-ict/DEV-DT-Microblog/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mariadb"
	"log"
	"os"
	"testing"
)

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	mariadbContainer, err := mariadb.RunContainer(ctx,
		testcontainers.WithImage("mariadb:11.0.3"),
		mariadb.WithDatabase("microblog-test"),
		mariadb.WithUsername("test"),
		mariadb.WithPassword("test1234"),
	)
	if err != nil {
		log.Fatalf("failed to start container: %s", err)
	}

	os.Setenv("APP_MODE", "test")
	connectionString, err := mariadbContainer.ConnectionString(ctx)
	if err != nil {
		log.Fatalf("failed to get connection string: %s", err)
	}
	os.Setenv("DB_CONNECTION_STRING", connectionString)

	err1 := repositories.CreateUser("TestDummy", "Test1", "TestDummy")
	err2 := repositories.CreateUser("TestDummy", "Test2", "TestDummy")

	assert.Nil(t, err1)
	assert.NotNil(t, err2)

	defer func() {
		if err := mariadbContainer.Terminate(ctx); err != nil {
			log.Fatalf("failed to terminate container: %s", err)
		}
	}()
}
