package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, _ command) error {
	err := s.db.ClearUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Failed to clear user table: %v", err)
	}

	fmt.Println("User table has been cleared.")

	return nil
}