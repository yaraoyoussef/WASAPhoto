package api

import (
	"net/http"
	"regexp"
	"strings"
)

// function to validate username's format
func (p *User) IsValid(id string) bool {
	var trimmed = strings.TrimSpace(id)
	re := regexp.MustCompile(`^([a-z0-9._-])\w+$`)
	return re.MatchString(trimmed) && len(trimmed) >= 3 && len(trimmed) <= 16
}

// function to extract bearer from authentication header
func extractBearer(auth string) string {
	var tokens = strings.Split(auth, " ")
	if len(tokens) == 2 {
		return strings.TrimSpace(tokens[1])
	}
	return ""
}

// function to validate user, if it returns 0 then all is good
func validateUser(id string, token string) int {
	// if user is not logged, return error
	if token == "" {
		return http.StatusForbidden
	}
	// if user requesting to do operation is not logged user, return error
	if id != token {
		return http.StatusUnauthorized
	}
	return 0
}
