package main

import (
	"context"
	"fmt"

	"github.com/iitung-dev/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	
	return func(s *state, cmd command) error {
		user_name := s.config.Current_user_name
		current_user, err := s.db.GetUser(context.Background(), user_name)
		if err != nil {
			return fmt.Errorf("%v handler failed to get user: %v", cmd.Name, err)
		}

		return handler(s, cmd, current_user)
	}
}