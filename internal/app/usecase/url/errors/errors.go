package errors

type URLIsDeletedError struct{}

func (e URLIsDeletedError) Error() string {
	return "URL is deleted"
}

type URLNotFoundError struct{}

func (e URLNotFoundError) Error() string {
	return "URL not found"
}
