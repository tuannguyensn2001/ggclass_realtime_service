package app

type HttpError struct {
	StatusCode int         `json:"statusCode" yaml:"statusCode"`
	Data       interface{} `json:"data"`
	Message    string      `json:"message" yaml:"message"`
	RootErr    error       `json:"-"`
	Code       string      `json:"code" yaml:"code"`
}

func (err *HttpError) Error() string {
	return err.Message
}
