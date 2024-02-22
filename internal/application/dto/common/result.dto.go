package common_dto

type ResultDto struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func NewSuccessDataResult(data interface{}) *ResultDto {
	return &ResultDto{
		Success: true,
		Data:    data,
	}
}

func NewSuccessMessageResult(message string) *ResultDto {
	return &ResultDto{
		Success: true,
		Message: message,
	}
}

func NewErrorResult(message, err string) *ResultDto {
	return &ResultDto{
		Success: false,
		Message: message,
		Error:   err,
	}
}
