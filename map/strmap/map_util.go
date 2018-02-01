package strmap

func empty() map[string]interface{} {
	return make(map[string]interface{})
}

func cp(dst map[string]interface{}, src map[string]interface{}) {
	for k, v := range src {
		dst[k] = v
	}
}

func dup(src map[string]interface{}) map[string]interface{} {
	dst := make(map[string]interface{}, len(src))
	cp(dst, src)
	return dst
}
