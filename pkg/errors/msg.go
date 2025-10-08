package e

const (
	SUCCESS                      string = "00000"
	ERROR_BAD_REQUEST            string = "00001" // 400
	ERROR_UNAUTHORIZED           string = "00002" // 401
	ERROR_FORBIDDEN              string = "00003" // 403
	ERROR_NOT_FOUND              string = "00004" // 404
	ERROR_METHOD_NOT_ALLOWED     string = "00005" // 405
	ERROR_CONFLICT               string = "00006" // 409
	ERROR_UNSUPPORTED_MEDIA_TYPE string = "00007" // 415
	ERROR_UNPROCESSABLE_ENTITY   string = "00008" // 422
	ERROR_TOO_MANY_REQUESTS      string = "00009" // 429
	INVALID_PARAMS               string = "00010"

	ERROR_INTERNAL_SERVER     string = "00011" // 500
	ERROR_NOT_IMPLEMENTED     string = "00012" // 501
	ERROR_BAD_GATEWAY         string = "00013" // 502
	ERROR_SERVICE_UNAVAILABLE string = "00014" // 503
	ERROR_GATEWAY_TIMEOUT     string = "00015" // 504
	ERROR_DB                  string = "00016"
	ERROR_EXTERNAL_SERVICE    string = "00017"

	// More error here
)

var MsgFlags = map[string]string{
	SUCCESS:                      "Success.",
	ERROR_BAD_REQUEST:            "Bad request.",
	ERROR_UNAUTHORIZED:           "Unauthorized access.",
	ERROR_FORBIDDEN:              "Forbidden.",
	ERROR_NOT_FOUND:              "Resource not found.",
	ERROR_METHOD_NOT_ALLOWED:     "Method not allowed.",
	ERROR_CONFLICT:               "Conflict detected.",
	ERROR_UNSUPPORTED_MEDIA_TYPE: "Unsupported media type.",
	ERROR_UNPROCESSABLE_ENTITY:   "Unprocessable entity.",
	ERROR_TOO_MANY_REQUESTS:      "Too many requests.",
	INVALID_PARAMS:               "Invalid parameters.",

	ERROR_INTERNAL_SERVER:     "Internal server error.",
	ERROR_NOT_IMPLEMENTED:     "Not implemented.",
	ERROR_BAD_GATEWAY:         "Bad gateway.",
	ERROR_SERVICE_UNAVAILABLE: "Service unavailable.",
	ERROR_GATEWAY_TIMEOUT:     "Gateway timeout.",
	ERROR_DB:                  "Database error.",
	ERROR_EXTERNAL_SERVICE:    "External service error.",

	// More error messages here
}

func GetMsg(code string) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR_INTERNAL_SERVER]
}
