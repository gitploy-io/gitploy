package interactor

type (
	// ListCursorOptions specifies the optional parameters that
	// support cursor pagination.
	ListOptions struct {
		Page    int
		PerPage int
	}
)
