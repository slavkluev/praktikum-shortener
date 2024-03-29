// Package middlewares содержит Middlewares
package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net/http"

	"github.com/google/uuid"
)

// Authenticator является Middleware для аутентификации пользователя
type Authenticator struct {
	secret []byte
}

// NewAuthenticator создание нового Authenticator
func NewAuthenticator(secret []byte) *Authenticator {
	return &Authenticator{secret: secret}
}

// Handle обработка Middleware
func (a Authenticator) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userCookie, userErr := r.Cookie("user_id")
		signCookie, signErr := r.Cookie("sign")

		if userErr != nil || signErr != nil {
			newUserID, sign, _ := a.generateUserID()
			a.setCookies(w, r, newUserID, sign)
		} else {
			h := hmac.New(sha256.New, a.secret)
			h.Write([]byte(userCookie.Value))
			calculatedSign := h.Sum(nil)
			sign, err := hex.DecodeString(signCookie.Value)
			if err != nil {
				io.WriteString(w, err.Error())
				return
			}

			if !hmac.Equal(calculatedSign, sign) {
				newUserID, sign, _ := a.generateUserID()
				a.setCookies(w, r, newUserID, sign)
			}
		}

		next.ServeHTTP(w, r)
	})
}

func (a Authenticator) generateUserID() (string, string, error) {
	newUserID := uuid.New().String()

	h := hmac.New(sha256.New, a.secret)
	_, err := h.Write([]byte(newUserID))

	if err != nil {
		return "", "", err
	}

	sign := h.Sum(nil)

	return newUserID, hex.EncodeToString(sign), nil
}

func (a Authenticator) setCookies(w http.ResponseWriter, r *http.Request, userID, sign string) {
	userIDCookie := &http.Cookie{
		Name:  "user_id",
		Value: userID,
	}
	signCookie := &http.Cookie{
		Name:  "sign",
		Value: sign,
	}

	http.SetCookie(w, userIDCookie)
	http.SetCookie(w, signCookie)

	_, err := r.Cookie("user_id")
	if err != nil {
		r.AddCookie(userIDCookie)
	}

	_, err = r.Cookie("sign")
	if err != nil {
		r.AddCookie(signCookie)
	}
}
