package errors

type URLIsDeletedError struct {
}

func (e URLIsDeletedError) Error() string {
	return "URL is deleted"
}

type URLNotFound struct {
}

func (e URLNotFound) Error() string {
	return "URL not found"
}
