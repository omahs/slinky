package json

import "encoding/json"

// IsValid checks if the given byte array is valid JSON.
// If the byte array is 0 length, this is a valid empty JSON object.
func IsValid(jsonBz []byte) error {
	if len(jsonBz) == 0 {
		return nil
	}

	var checkStruct map[string]interface{}
	return json.Unmarshal(jsonBz, &checkStruct)
}

// RPCRequest is a JSON-RPC request object, defined in accordance with the 
// [spec](See: http://www.jsonrpc.org/specification#request_object)
// 
// Data:
//  - JSONRPC: the version of the JSON-RPC protocol
//  - Method: the method to be invoked
//  - Params: the parameters to be passed to the method
//  - ID: the request ID
type RPCRequest struct {
	JSONRPC string            `json:"jsonrpc"`
	Method  string            `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int64             `json:"id"`
}

// DefaultRPCRequest returns a new RPCRequest with the given method and params.
func DefaultRPCRequest(method string, params []interface{}) RPCRequest {
	return RPCRequest{
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
		ID:      1,
	}
}


// RPCResponse is a JSON-RPC response object, defined in accordance with the
// [spec](See: http://www.jsonrpc.org/specification#response_object)
//
// Data:
//  - JSONRPC: the version of the JSON-RPC protocol
//  - Result: the result of the method invocation
//  - Error: the error that occurred during the method invocation
//  - ID: the request ID
type RPCResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *RPCError       `json:"error,omitempty"`
	ID      int64           `json:"id"`
}

// RPCError is a JSON-RPC error object, defined in accordance with the
// [spec](See: http://www.jsonrpc.org/specification#error_object)
//
// Data:
//  - Code: the error code
//  - Message: the error message
//  - Data: the error data
type RPCError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (e *RPCError) Error() string {
	return e.Message
}
