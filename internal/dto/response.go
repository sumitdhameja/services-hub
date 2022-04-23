package dto

// DataResponse is response object
type DataResponse struct {
	Data interface{} `json:"data"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}
type ErrorReply struct {
	Error ErrorMessage `json:"error"`
}

func ReplyError(message string) ErrorReply {
	return ErrorReply{
		Error: ErrorMessage{
			Message: message,
		},
	}
}
