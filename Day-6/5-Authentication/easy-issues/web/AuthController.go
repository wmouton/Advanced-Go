package web

import (
	"encoding/json"
	"github.com/Advanced-Go/Day-6/5-Authentication/easy-issues"
	"github.com/Advanced-Go/Day-6/5-Authentication/easy-issues/domain"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

// Controller for Issue model
type AuthController struct {
	AuthService domain.AuthService
	Secret      string
}

type LoginResponse struct {
	Token string `json:"token"`
}

type JWTData struct {
	jwt.StandardClaims
	CustomClaims map[string]string `json:"custom,omitempty"`
}

func (c AuthController) Verify(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (c AuthController) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusBadRequest)
	}

	email, password, ok := r.BasicAuth()
	if !ok {
		http.Error(w, "authorization failed", http.StatusUnauthorized)
	}

	userRegistration, err := c.AuthService.GetRegistrationByEmail(email)
	if err != nil {
		http.Error(w, "authorization failed", http.StatusUnauthorized)
	}

	if userRegistration.Status == domain.RegistrationStatusDeleted {
		http.Error(w, "authorization failed", http.StatusUnauthorized)
	}

	ok = easy_issues.CheckPasswordHash(password, userRegistration.PasswordHash)
	if !ok {
		http.Error(w, "authorization failed", http.StatusUnauthorized)
	}

	claims := JWTData{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			Issuer:    "auth.service",
		},
		CustomClaims: map[string]string{
			"userId": userRegistration.Uuid,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(c.Secret))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response := LoginResponse{
		Token: tokenString,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
