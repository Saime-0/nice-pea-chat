package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"reflect"

	"github.com/saime-0/cute-chat-backend/internal/usecases"
	"github.com/sirupsen/logrus"
)

type Auth struct {
	authUc usecases.AuthUsecase
}

func (h *Auth) Endpoint() string {
	return "/auth"
}

func (h *Auth) Method() string {
	return http.MethodPost
}

func (h *Auth) Fn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			logrus.Debug("[Auth] read body: %v", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		requestBody := _AuthRequestBody{}
		err = json.Unmarshal(b, &requestBody)
		if err != nil {
			logrus.Debug("[Auth] Failed unmarshal request body: %v", err)
			http.Error(w, "Failed unmarshal request body", http.StatusBadRequest)
			return
		}
		out, err := h.authUc.Auth(usecases.AuthIn{
			Login: requestBody.Login,
		})
		w.Write([]byte("ok"))
		w.WriteHeader(http.StatusOK)
	}
}

type _AuthRequestBody struct {
	Login string `json: "login"`
}

type _AuthResponse struct {
	AccessToken string `json: "access_token"`
}
