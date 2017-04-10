package response

type ResponseWithCollection struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type ResponseWithSuccess struct {
	Status  int         `json:"status"`
	Success interface{} `json:"success"`
}

type ResponseWithError struct {
	Status int         `json:"status"`
	Error  interface{} `json:"error"`
}
