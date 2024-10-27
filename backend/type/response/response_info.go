package response

type InfoResponse struct {
	Success bool   `json:"success"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type GenericInfoResponse[T any] struct {
	Success bool   `json:"success"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
}

func Success(args1 any, args2 ...any) *InfoResponse {
	if message, ok := args1.(string); ok {
		if len(args2) == 0 {
			return &InfoResponse{
				Success: true,
				Message: message,
			}
		}
		if message2, ok := args2[0].(string); ok {
			return &InfoResponse{
				Success: true,
				Code:    message,
				Message: message2,
			}
		} else {
			return &InfoResponse{
				Success: true,
				Code:    message,
				Data:    message2,
			}
		}
	}
	return &InfoResponse{
		Success: true,
		Data:    args1,
	}
}
