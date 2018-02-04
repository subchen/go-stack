package conv

type Data struct {
	Value interface{}
}

func NewData(value interface{}) *Data {
	return &Data{value}
}

func (d *Data) IsNil() bool {
	return d.Value == nil
}

func (d *Data) AsString() string {
	return AsString(d.Value)
}

func (d *Data) ToString() (string, error) {
	return ToString(d.Value)
}

func (d *Data) AsBool() bool {
	return AsBool(d.Value)
}

func (d *Data) ToBool() (bool, error) {
	return ToBool(d.Value)
}

func (d *Data) AsInt() int {
	return AsInt(d.Value)
}

func (d *Data) ToInt() (int, error) {
	return ToInt(d.Value)
}

func (d *Data) AsInt8() int8 {
	return AsInt8(d.Value)
}

func (d *Data) ToInt8() (int8, error) {
	return ToInt8(d.Value)
}

func (d *Data) AsInt16() int16 {
	return AsInt16(d.Value)
}

func (d *Data) ToInt16() (int16, error) {
	return ToInt16(d.Value)
}

func (d *Data) AsInt32() int32 {
	return AsInt32(d.Value)
}

func (d *Data) ToInt32() (int32, error) {
	return ToInt32(d.Value)
}

func (d *Data) AsInt64() int64 {
	return AsInt64(d.Value)
}

func (d *Data) ToInt64() (int64, error) {
	return ToInt64(d.Value)
}

func (d *Data) AsUint() uint {
	return AsUint(d.Value)
}

func (d *Data) ToUint() (uint, error) {
	return ToUint(d.Value)
}

func (d *Data) AsUint8() uint8 {
	return AsUint8(d.Value)
}

func (d *Data) ToUint8() (uint8, error) {
	return ToUint8(d.Value)
}

func (d *Data) AsUint16() uint16 {
	return AsUint16(d.Value)
}

func (d *Data) ToUint16() (uint16, error) {
	return ToUint16(d.Value)
}

func (d *Data) AsUint32() uint32 {
	return AsUint32(d.Value)
}

func (d *Data) ToUint32() (uint32, error) {
	return ToUint32(d.Value)
}

func (d *Data) AsUint64() uint64 {
	return AsUint64(d.Value)
}

func (d *Data) ToUint64() (uint64, error) {
	return ToUint64(d.Value)
}

func (d *Data) AsFloat32() float32 {
	return AsFloat32(d.Value)
}

func (d *Data) ToFloat32() (float32, error) {
	return ToFloat32(d.Value)
}

func (d *Data) AsFloat64() float64 {
	return AsFloat64(d.Value)
}

func (d *Data) ToFloat64() (float64, error) {
	return ToFloat64(d.Value)
}

func (d *Data) AsStringSlice() []string {
	return AsStringSlice(d.Value)
}

func (d *Data) ToStringSlice() ([]string, error) {
	return ToStringSlice(d.Value)
}

func (d *Data) AsBoolSlice() []bool {
	return AsBoolSlice(d.Value)
}

func (d *Data) ToBoolSlice() ([]bool, error) {
	return ToBoolSlice(d.Value)
}

func (d *Data) AsIntSlice() []int {
	return AsIntSlice(d.Value)
}

func (d *Data) ToIntSlice() ([]int, error) {
	return ToIntSlice(d.Value)
}

func (d *Data) AsInt8Slice() []int8 {
	return AsInt8Slice(d.Value)
}

func (d *Data) ToInt8Slice() ([]int8, error) {
	return ToInt8Slice(d.Value)
}

func (d *Data) AsInt16Slice() []int16 {
	return AsInt16Slice(d.Value)
}

func (d *Data) ToInt16Slice() ([]int16, error) {
	return ToInt16Slice(d.Value)
}

func (d *Data) AsInt32Slice() []int32 {
	return AsInt32Slice(d.Value)
}

func (d *Data) ToInt32Slice() ([]int32, error) {
	return ToInt32Slice(d.Value)
}

func (d *Data) AsInt64Slice() []int64 {
	return AsInt64Slice(d.Value)
}

func (d *Data) ToInt64Slice() ([]int64, error) {
	return ToInt64Slice(d.Value)
}

func (d *Data) AsUintSlice() []uint {
	return AsUintSlice(d.Value)
}

func (d *Data) ToUintSlice() ([]uint, error) {
	return ToUintSlice(d.Value)
}

func (d *Data) AsUint8Slice() []uint8 {
	return AsUint8Slice(d.Value)
}

func (d *Data) ToUint8Slice() ([]uint8, error) {
	return ToUint8Slice(d.Value)
}

func (d *Data) AsUint16Slice() []uint16 {
	return AsUint16Slice(d.Value)
}

func (d *Data) ToUint16Slice() ([]uint16, error) {
	return ToUint16Slice(d.Value)
}

func (d *Data) AsUint32Slice() []uint32 {
	return AsUint32Slice(d.Value)
}

func (d *Data) ToUint32Slice() ([]uint32, error) {
	return ToUint32Slice(d.Value)
}

func (d *Data) AsUint64Slice() []uint64 {
	return AsUint64Slice(d.Value)
}

func (d *Data) ToUint64Slice() ([]uint64, error) {
	return ToUint64Slice(d.Value)
}

func (d *Data) AsFloat32Slice() []float32 {
	return AsFloat32Slice(d.Value)
}

func (d *Data) ToFloat32Slice() ([]float32, error) {
	return ToFloat32Slice(d.Value)
}

func (d *Data) AsFloat64Slice() []float64 {
	return AsFloat64Slice(d.Value)
}

func (d *Data) ToFloat64Slice() ([]float64, error) {
	return ToFloat64Slice(d.Value)
}
