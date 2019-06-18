package FdfExtract

import (
	"errors"
	"strings"
)

// checkpath checks for correct extension
func checkpath(path string) error {
	if strings.Contains(path, ".fdf") {
		return nil
	}
	return errors.New("NOT_FDF")
}
