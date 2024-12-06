package response

import "github.com/gin-gonic/gin"

// Response contains common fields and methods for all responses
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// SetCode sets the response code.
func (b *Response) SetCode(code int) *Response {

	b.Code = code
	
	return b
}

// SetMessage sets the response message
func (b *Response) SetMessage(message string) *Response {

	b.Message = message

	return b
}

// SetData sets the response data
func (b *Response) SetData(data interface{}) *Response {

	b.Data = data

	return b
}

// Send sends the response to the client
func (b *Response) Json(ctx *gin.Context) {

	ctx.JSON(200, b)
}

// Success represents a successful response
type Success struct {
	Response
}

// NewSuccess creates a new success response
func NewSuccess() *Success {

	return &Success{
		Response: Response{
			Code:    10200,
			Message: "success",
		},
	}
}

// Error represents an error response
type Error struct {
	Response
}

// NewError creates a new error response
func NewError() *Error {

	return &Error{
		Response: Response{
			Code:    10500,
			Message: "error",
		},
	}
}

// Send is a method that all responses must implement to send the response
func (s *Success) Json(ctx *gin.Context) {
	s.Response.Json(ctx)
}

func (e *Error) Json(ctx *gin.Context) {
	e.Response.Json(ctx)
}
