package ports

type JWTservice interface {
	Sign(claims map[string]interface{}) (string, error)
	Verify(token string) (map[string]interface{}, error)
}
