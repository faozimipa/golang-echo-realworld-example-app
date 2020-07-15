package handler

import (
	"github.com/faozimipa/golang-echo-realworld-example-app/article"
	"github.com/faozimipa/golang-echo-realworld-example-app/user"

)

//Handler struct
type Handler struct {
	userStore    user.Store
	articleStore article.Store
}

//NewHandler func
func NewHandler(us user.Store, as article.Store) *Handler {
	return &Handler{
		userStore:    us,
		articleStore: as,
	}
}
