package fixtures

// RouteInt type
type RouteInt int

const (
	// RouteTestResponse enum
	RouteTestResponse RouteInt = iota
	// RouteTestResponseError enum
	RouteTestResponseError
	// RouteTestResponseNotFound enum
	RouteTestResponseNotFound
)

var routeTestResponseText = map[RouteInt]string{
	RouteTestResponse:         "/test-response",
	RouteTestResponseError:    "/test-response-error",
	RouteTestResponseNotFound: "/test-response-notfound",
}

// RouteText value
func RouteText(code RouteInt) string {
	return routeTestResponseText[code]
}
