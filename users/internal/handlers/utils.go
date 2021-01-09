package handlers

func errorResponse(message string) map[string]string {
	return map[string]string{
		"error": message,
	}
}
