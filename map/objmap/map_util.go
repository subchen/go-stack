package strmap

type (
	Key   interface{}
	Value interface{}
)

func empty() map[Key]Value {
	return make(map[Key]Value)
}

func cp(dst map[Key]Value, src map[Key]Value) {
	for k, v := range src {
		dst[k] = v
	}
}

func dup(src map[Key]Value) map[Key]Value {
	dst := make(map[Key]Value, len(src))
	cp(dst, src)
	return dst
}
