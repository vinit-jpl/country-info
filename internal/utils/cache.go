package utils

type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
}
