package middlewares

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/google/uuid"
	"io"
	"net/http"
)

type Authenticator struct {
	secret []byte
}

func NewAuthenticator(secret []byte) *Authenticator {
	return &Authenticator{secret: secret}
}

func (a Authenticator) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userCookie, userErr := r.Cookie("user_id")
		signCookie, signErr := r.Cookie("sign")

		if userErr != nil || signErr != nil {
			newUserId, sign, _ := a.generateUserId()
			a.setCookies(w, r, newUserId, sign)
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
				newUserId, sign, _ := a.generateUserId()
				a.setCookies(w, r, newUserId, sign)
			}
		}

		next.ServeHTTP(w, r)
	}
}

func (a Authenticator) generateUserId() (string, string, error) {
	newUserId := uuid.New().String()

	h := hmac.New(sha256.New, a.secret)
	_, err := h.Write([]byte(newUserId))

	if err != nil {
		return "", "", err
	}

	sign := h.Sum(nil)

	return newUserId, hex.EncodeToString(sign), nil
}

func (a Authenticator) setCookies(w http.ResponseWriter, r *http.Request, userId, sign string) {
	userIdCookie := &http.Cookie{
		Name:  "user_id",
		Value: userId,
	}
	signCookie := &http.Cookie{
		Name:  "sign",
		Value: sign,
	}

	http.SetCookie(w, userIdCookie)
	http.SetCookie(w, signCookie)

	_, err := r.Cookie("user_id")
	if err != nil {
		r.AddCookie(userIdCookie)
	}

	_, err = r.Cookie("sign")
	if err != nil {
		r.AddCookie(signCookie)
	}
}
