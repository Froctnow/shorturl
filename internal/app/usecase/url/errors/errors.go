package errors

import (
	"fmt"
)

type URLIsDeletedError struct {
}

func (e URLIsDeletedError) Error() string {
	return fmt.Sprintf("URL is deleted")
}

type URLNotFound struct {
}

func (e URLNotFound) Error() string {
	return fmt.Sprintf("URL not found")
}
