package apiserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-next-pizza/internal/app/model"
	"github.com/go-next-pizza/internal/app/storage/test_storage"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestServer_AuthenticateUser(t *testing.T) {
    storage := test_storage.NewSQLStorage()
    u := model.TestUser(t)
    storage.User().CreateUser(u)

    secretKey := []byte("secret_key")
    s := newServer(storage, secretKey, 3600, 86400)
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
    })

    // valid token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"sub": u.Id})
    tokenStr, _ := token.SignedString(secretKey)

    t.Run("authenticated", func(t *testing.T) {
        rec := httptest.NewRecorder()
        req, _ := http.NewRequest(http.MethodGet, "/", nil)
        req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tokenStr))
        s.authenticateUser(handler).ServeHTTP(rec, req)
        assert.Equal(t, http.StatusOK, rec.Code)
    })

    t.Run("not authenticated", func(t *testing.T) {
        rec := httptest.NewRecorder()
        req, _ := http.NewRequest(http.MethodGet, "/", nil)
        s.authenticateUser(handler).ServeHTTP(rec, req)
        assert.Equal(t, http.StatusUnauthorized, rec.Code)
    })
}

func TestServer_HandleUsersCreate(t *testing.T) {
    s := newServer(test_storage.NewSQLStorage(), []byte("secret_key"), 3600, 86400)
	testCases := []struct{
		name string
		payload interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email": "test_email@test.com",
				"password": "2007Coolo_ww!",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "invalid payload",
			payload: "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid request body",
			payload: map[string]string{
				"email": "invalid email",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}

			json.NewEncoder(b).Encode(tc.payload)

			req, _ := http.NewRequest(http.MethodPost, "/users", b)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_HandleSessionCreate_And_Refresh(t *testing.T) {
	u := model.TestUser(t)
	storage := test_storage.NewSQLStorage()

    s := newServer(storage, []byte("secret_key"), 3600, 86400)

	storage.User().CreateUser(u)

	// login
	loginBody := map[string]string{
		"email": u.Email,
		"password": u.Password,
	}
	
	b := &bytes.Buffer{}
	json.NewEncoder(b).Encode(loginBody)
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/sessions", b)
	s.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)

	var tokens map[string]string
	json.Unmarshal(rec.Body.Bytes(), &tokens)
	assert.NotEmpty(t, tokens["accessToken"])
	assert.NotEmpty(t, tokens["refreshToken"])

	// access private with accessToken
	rec2 := httptest.NewRecorder()
	req2, _ := http.NewRequest(http.MethodGet, "/private/whoami", nil)
	req2.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tokens["accessToken"]))
	s.ServeHTTP(rec2, req2)
	assert.Equal(t, http.StatusOK, rec2.Code)

	// refresh access token
	refreshBody := map[string]string{"refreshToken": tokens["refreshToken"]}
	b2 := &bytes.Buffer{}
	json.NewEncoder(b2).Encode(refreshBody)
	rec3 := httptest.NewRecorder()
	req3, _ := http.NewRequest(http.MethodPost, "/sessions/refresh", b2)
	s.ServeHTTP(rec3, req3)
	assert.Equal(t, http.StatusOK, rec3.Code)

	var refreshed map[string]string
	json.Unmarshal(rec3.Body.Bytes(), &refreshed)
	assert.NotEmpty(t, refreshed["accessToken"])
}

func TestServer_HandleSessionRefresh_Negative(t *testing.T) {
    u := model.TestUser(t)
    storage := test_storage.NewSQLStorage()
    secretKey := []byte("secret_key")
    s := newServer(storage, secretKey, 3600, 10)

    storage.User().CreateUser(u)

    // helper to call refresh
    callRefresh := func(tokenStr string) *httptest.ResponseRecorder {
        rec := httptest.NewRecorder()
        body := map[string]string{"refreshToken": tokenStr}
        b := &bytes.Buffer{}
        json.NewEncoder(b).Encode(body)
        req, _ := http.NewRequest(http.MethodPost, "/sessions/refresh", b)
        s.ServeHTTP(rec, req)
        return rec
    }

    t.Run("wrong typ claim", func(t *testing.T) {
        // typ omitted or set to access
        bad := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
            "sub": u.Id,
            "exp": time.Now().Add(10 * time.Second).Unix(),
            "iat": time.Now().Unix(),
            "typ": "access",
        })
        badStr, _ := bad.SignedString(secretKey)
        rec := callRefresh(badStr)
        assert.Equal(t, http.StatusUnauthorized, rec.Code)
    })

    t.Run("expired refresh", func(t *testing.T) {
        exp := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
            "sub": u.Id,
            "exp": time.Now().Add(-1 * time.Minute).Unix(),
            "iat": time.Now().Add(-2 * time.Minute).Unix(),
            "typ": "refresh",
        })
        expStr, _ := exp.SignedString(secretKey)
        rec := callRefresh(expStr)
        assert.Equal(t, http.StatusUnauthorized, rec.Code)
    })

    t.Run("signed with different secret", func(t *testing.T) {
        otherKey := []byte("other_secret")
        tok := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
            "sub": u.Id,
            "exp": time.Now().Add(10 * time.Second).Unix(),
            "iat": time.Now().Unix(),
            "typ": "refresh",
        })
        tokStr, _ := tok.SignedString(otherKey)
        rec := callRefresh(tokStr)
        assert.Equal(t, http.StatusUnauthorized, rec.Code)
    })
}