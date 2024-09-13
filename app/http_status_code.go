package main

// HTTPStatus represents standard HTTP status codes
type HTTPStatus int

const (
	// 1xx Informational
	StatusContinue           HTTPStatus = 100
	StatusSwitchingProtocols HTTPStatus = 101
	StatusProcessing         HTTPStatus = 102

	// 2xx Success
	StatusOK               HTTPStatus = 200
	StatusCreated          HTTPStatus = 201
	StatusAccepted         HTTPStatus = 202
	StatusNonAuthoritative HTTPStatus = 203
	StatusNoContent        HTTPStatus = 204
	StatusResetContent     HTTPStatus = 205
	StatusPartialContent   HTTPStatus = 206

	// 3xx Redirection
	StatusMultipleChoices   HTTPStatus = 300
	StatusMovedPermanently  HTTPStatus = 301
	StatusFound             HTTPStatus = 302
	StatusSeeOther          HTTPStatus = 303
	StatusNotModified       HTTPStatus = 304
	StatusUseProxy          HTTPStatus = 305
	StatusTemporaryRedirect HTTPStatus = 307
	StatusPermanentRedirect HTTPStatus = 308

	// 4xx Client Errors
	StatusBadRequest                   HTTPStatus = 400
	StatusUnauthorized                 HTTPStatus = 401
	StatusPaymentRequired              HTTPStatus = 402
	StatusForbidden                    HTTPStatus = 403
	StatusNotFound                     HTTPStatus = 404
	StatusMethodNotAllowed             HTTPStatus = 405
	StatusNotAcceptable                HTTPStatus = 406
	StatusProxyAuthRequired            HTTPStatus = 407
	StatusRequestTimeout               HTTPStatus = 408
	StatusConflict                     HTTPStatus = 409
	StatusGone                         HTTPStatus = 410
	StatusLengthRequired               HTTPStatus = 411
	StatusPreconditionFailed           HTTPStatus = 412
	StatusRequestEntityTooLarge        HTTPStatus = 413
	StatusRequestURITooLong            HTTPStatus = 414
	StatusUnsupportedMediaType         HTTPStatus = 415
	StatusRequestedRangeNotSatisfiable HTTPStatus = 416
	StatusExpectationFailed            HTTPStatus = 417
	StatusTeapot                       HTTPStatus = 418 // RFC 2324: "I'm a teapot"

	// 5xx Server Errors
	StatusInternalServerError     HTTPStatus = 500
	StatusNotImplemented          HTTPStatus = 501
	StatusBadGateway              HTTPStatus = 502
	StatusServiceUnavailable      HTTPStatus = 503
	StatusGatewayTimeout          HTTPStatus = 504
	StatusHTTPVersionNotSupported HTTPStatus = 505
)
