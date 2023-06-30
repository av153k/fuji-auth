package utils

import (
	"fmt"
	"fuji-auth/pkg/constants"
)

func VerifyRole(role string) (string, error) {
	// Switch given role.
	switch role {
	case constants.Admin:
		// Nothing to do, verified successfully.
	case constants.Moderator:
		// Nothing to do, verified successfully.
	case constants.User:
		// Nothing to do, verified successfully.
	default:
		// Return error message.
		return "", fmt.Errorf("role '%v' does not exist", role)
	}

	return role, nil
}
