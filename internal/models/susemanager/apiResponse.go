// Package sumamodels - structs needed for SUSE Manager API Calls
package sumamodels

// RespAPISuccess - return of api call
type RespAPISuccess struct {
	Success bool        `json:"success"`
	Result  interface{} `json:"result"`
}

// RespAPI - return of api call
type RespAPI struct {
	Success  bool        `json:"success"`
	Result   interface{} `json:"result,omitempty"`
	Message  string      `json:"message,omitempty"`
	Messages []string    `json:"messages,omitempty"`
}
