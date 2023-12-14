package http

import "net/http"

func getUser(r *http.Request) (string, error) {
	userCookie, err := r.Cookie("user_id")
	if err != nil {
		return "", err
	}

	return userCookie.Value, nil
}
