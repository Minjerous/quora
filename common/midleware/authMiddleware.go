package midleware

import (
	"net/http"
	"sync"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

var wg sync.WaitGroup

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return nil
}
