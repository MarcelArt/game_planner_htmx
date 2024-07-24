package handlers

import "github.com/MarcelArt/game_planner_htmx/repositories"

type UserHandler struct {
	userRepo repositories.IUserRepo
}

func NewUserHandler(userRepo repositories.IUserRepo) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}

// func (h *UserHandler)
