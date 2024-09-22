package response

func NewResponseEntity() *ResponseEntity {
	return &ResponseEntity{}
}

type ResponseEntity struct {
	HttpStatus int         `json:"-"`
	Content    interface{} `json:"data"`
}

func (r ResponseEntity) StatusCode(status int) *ResponseEntity {
	r.HttpStatus = status
	return &r
}

func (r ResponseEntity) Body(body interface{}) *ResponseEntity {
	r.Content = body
	return &r
}
