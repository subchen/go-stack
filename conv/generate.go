package conv

//go:generate frep --overwrite generate_to_number.gotmpl:to_int.go     -e name=int   -e size=""
//go:generate frep --overwrite generate_to_number.gotmpl:to_int8.go    -e name=int   -e size=8
//go:generate frep --overwrite generate_to_number.gotmpl:to_int16.go   -e name=int   -e size=16
//go:generate frep --overwrite generate_to_number.gotmpl:to_int32.go   -e name=int   -e size=32
//go:generate frep --overwrite generate_to_number.gotmpl:to_int64.go   -e name=int   -e size=64
//go:generate frep --overwrite generate_to_number.gotmpl:to_uint.go    -e name=uint  -e size=""
//go:generate frep --overwrite generate_to_number.gotmpl:to_uint8.go   -e name=uint  -e size=8
//go:generate frep --overwrite generate_to_number.gotmpl:to_uint16.go  -e name=uint  -e size=16
//go:generate frep --overwrite generate_to_number.gotmpl:to_uint32.go  -e name=uint  -e size=32
//go:generate frep --overwrite generate_to_number.gotmpl:to_uint64.go  -e name=uint  -e size=64
//go:generate frep --overwrite generate_to_number.gotmpl:to_float32.go -e name=float -e size=32
//go:generate frep --overwrite generate_to_number.gotmpl:to_float64.go -e name=float -e size=64

//go:generate frep --overwrite generate_to_slice.gotmpl:to_string_slice.go  -e name=string
//go:generate frep --overwrite generate_to_slice.gotmpl:to_bool_slice.go    -e name=bool
//go:generate frep --overwrite generate_to_slice.gotmpl:to_int_slice.go     -e name=int
//go:generate frep --overwrite generate_to_slice.gotmpl:to_int8_slice.go    -e name=int8
//go:generate frep --overwrite generate_to_slice.gotmpl:to_int16_slice.go   -e name=int16
//go:generate frep --overwrite generate_to_slice.gotmpl:to_int32_slice.go   -e name=int32
//go:generate frep --overwrite generate_to_slice.gotmpl:to_int64_slice.go   -e name=int64
//go:generate frep --overwrite generate_to_slice.gotmpl:to_uint_slice.go    -e name=uint
//go:generate frep --overwrite generate_to_slice.gotmpl:to_uint8_slice.go   -e name=uint8
//go:generate frep --overwrite generate_to_slice.gotmpl:to_uint16_slice.go  -e name=uint16
//go:generate frep --overwrite generate_to_slice.gotmpl:to_uint32_slice.go  -e name=uint32
//go:generate frep --overwrite generate_to_slice.gotmpl:to_uint64_slice.go  -e name=uint64
//go:generate frep --overwrite generate_to_slice.gotmpl:to_float32_slice.go -e name=float32
//go:generate frep --overwrite generate_to_slice.gotmpl:to_float64_slice.go -e name=float64
