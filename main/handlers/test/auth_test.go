package test

import (
	"hibernum-backend/main/handlers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateToken(t *testing.T) {
	sut := handlers.NewAuthentication()
	token, err := sut.CreateToken("testuser")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestParseToken(t *testing.T) {
	sut := handlers.NewAuthentication()
	tokenString, _ := sut.CreateToken("testuser")
	token, err := sut.ParseToken(tokenString)
	assert.NoError(t, err)
	assert.NotNil(t, token)
}

// TODO: FAIL
func TestLoginHandler(t *testing.T) {
	sut := handlers.NewAuthentication()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _ := sut.CreateToken("testuser")
		w.Header().Set("Authorization", token)
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	// req, _ := http.NewRequest("GET", ts.URL, nil)
	res := httptest.NewRecorder()

	sut.LoginHandler("testuser", "testpassword", res)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.NotEmpty(t, res.Header().Get("Authorization"))
}
