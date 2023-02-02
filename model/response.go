package model

type SuccessResponse struct {
	Data any
}

type FailResponse struct {
	Data    any
	Message string
	Error   error
}
