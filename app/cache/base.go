package cache

// Cache stores arbitrary data keyed by strings
type Cache interface {
	Get(string) (interface{}, error)
	Set(string, interface{}) error
}
