package error

type APIErrors struct {
	Errors []*APIError `json:"errors"`
}

func (errors *APIErrors) Status() int {
	return errors.Errors[0].Status
}

type APIError struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Title   string `json:"title"`
	Details string `json:"details"`
	Href    string `json:"href"`
}

func newAPIError(status int, code string, title string, details string, href string) *APIError {
	return &APIError{
		Status:  status,
		Code:    code,
		Title:   title,
		Details: details,
		Href:    href,
	}
}

var (
	errDatabase     = newAPIError(500, "database_error", "Database Error", "An unknown error occurred.", "")
	errInvalidSet   = newAPIError(404, "invalid_set", "Invalid Set", "The set you requested does not exist.", "")
	errInvalidGroup = newAPIError(404, "invalid_group", "Invalid Group", "The group you requested does not exist.", "")
)
