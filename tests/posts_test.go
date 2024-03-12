package tests

import (
	"github.com/che-ict/DEV-DT-Microblog/handlers"
	"github.com/che-ict/DEV-DT-Microblog/repositories"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeletePost(t *testing.T) {

	// GIVEN
	// 2 users
	// 2 post van user 1
	err := repositories.CreateUser("user1", "password1", "displayname1")
	err2 := repositories.CreateUser("user2", "password2", "displayname2")
	if err != nil || err2 != nil {
		t.Error("Error creating user")
	}
	repositories.CreatePost("post1", "user1")
	repositories.CreatePost("post2", "user1")

	// WHEN
	// 1 post verwijdern door user 1 -> goed gaan
	// 1 post verwijderen door user 2 -> fout gaan
	err = repositories.DeletePost(1)
	err2 = repositories.DeletePost(2)

	// THEN -- use assert
	// delete post 1 -> no error -> post 1 should be deleted
	// delete post 2 -> error -> post 2 should not be deleted
	assert.Nil(t, err)
	assert.NotNil(t, err2)
}

func TestDeleteHTTP(t *testing.T) {
	// GIVEN
	// 2 users
	// 2 post van user 1
	err := repositories.CreateUser("user1", "password1", "displayname1")
	err2 := repositories.CreateUser("user2", "password2", "displayname2")
	if err != nil || err2 != nil {
		t.Error("Error creating user")
	}
	repositories.CreatePost("post1", "user1")
	repositories.CreatePost("post2", "user1")

	// WHEN
	// creat cookie voor user1
	// delete post 1 -> goed gaan

	cookie := &http.Cookie{
		Name:  "user",
		Value: "user1",
	}
	req := httptest.NewRequest(http.MethodDelete, "/post/1", nil)
	rec := httptest.NewRecorder()
	req.AddCookie(cookie)
	c := echo.New().NewContext(req, rec)
	handlers.DeletePostHandler(c)

	cookie = &http.Cookie{
		Name:  "user",
		Value: "user2",
	}
	req2 := httptest.NewRequest(http.MethodDelete, "/post/2", nil)
	rec2 := httptest.NewRecorder()
	req2.AddCookie(cookie)
	c2 := echo.New().NewContext(req2, rec2)
	handlers.DeletePostHandler(c2)

	// THEN -- use assert
	// delete post 1 -> no error -> post 1 should be deleted
	// delete post 2 -> error -> post 2 should not be deleted
	assert.Equal(t, http.StatusFound, rec.Code)
	assert.Equal(t, http.StatusForbidden, rec2.Code)
}
