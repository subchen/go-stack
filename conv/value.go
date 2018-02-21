package conv

type Value struct {
	Data interface{}
}

func NewValue(v interface{}) *Value {
	return &Value{v}
}

func (v *Value) IsNil() bool {
	return v.Data == nil
}

func (v *Value) AsString() string {
	return AsString(v.Data)
}

func (v *Value) ToString() (string, error) {
	return ToString(v.Data)
}

func (v *Value) AsBool() bool {
	return AsBool(v.Data)
}

func (v *Value) ToBool() (bool, error) {
	return ToBool(v.Data)
}

func (v *Value) AsInt() int {
	return AsInt(v.Data)
}

func (v *Value) ToInt() (int, error) {
	return ToInt(v.Data)
}

func (v *Value) AsInt8() int8 {
	return AsInt8(v.Data)
}

func (v *Value) ToInt8() (int8, error) {
	return ToInt8(v.Data)
}

func (v *Value) AsInt16() int16 {
	return AsInt16(v.Data)
}

func (v *Value) ToInt16() (int16, error) {
	return ToInt16(v.Data)
}

func (v *Value) AsInt32() int32 {
	return AsInt32(v.Data)
}

func (v *Value) ToInt32() (int32, error) {
	return ToInt32(v.Data)
}

func (v *Value) AsInt64() int64 {
	return AsInt64(v.Data)
}

func (v *Value) ToInt64() (int64, error) {
	return ToInt64(v.Data)
}

func (v *Value) AsUint() uint {
	return AsUint(v.Data)
}

func (v *Value) ToUint() (uint, error) {
	return ToUint(v.Data)
}

func (v *Value) AsUint8() uint8 {
	return AsUint8(v.Data)
}

func (v *Value) ToUint8() (uint8, error) {
	return ToUint8(v.Data)
}

func (v *Value) AsUint16() uint16 {
	return AsUint16(v.Data)
}

func (v *Value) ToUint16() (uint16, error) {
	return ToUint16(v.Data)
}

func (v *Value) AsUint32() uint32 {
	return AsUint32(v.Data)
}

func (v *Value) ToUint32() (uint32, error) {
	return ToUint32(v.Data)
}

func (v *Value) AsUint64() uint64 {
	return AsUint64(v.Data)
}

func (v *Value) ToUint64() (uint64, error) {
	return ToUint64(v.Data)
}

func (v *Value) AsFloat32() float32 {
	return AsFloat32(v.Data)
}

func (v *Value) ToFloat32() (float32, error) {
	return ToFloat32(v.Data)
}

func (v *Value) AsFloat64() float64 {
	return AsFloat64(v.Data)
}

func (v *Value) ToFloat64() (float64, error) {
	return ToFloat64(v.Data)
}

func (v *Value) AsStringSlice() []string {
	return AsStringSlice(v.Data)
}

func (v *Value) ToStringSlice() ([]string, error) {
	return ToStringSlice(v.Data)
}

func (v *Value) AsBoolSlice() []bool {
	return AsBoolSlice(v.Data)
}

func (v *Value) ToBoolSlice() ([]bool, error) {
	return ToBoolSlice(v.Data)
}

func (v *Value) AsIntSlice() []int {
	return AsIntSlice(v.Data)
}

func (v *Value) ToIntSlice() ([]int, error) {
	return ToIntSlice(v.Data)
}

func (v *Value) AsInt8Slice() []int8 {
	return AsInt8Slice(v.Data)
}

func (v *Value) ToInt8Slice() ([]int8, error) {
	return ToInt8Slice(v.Data)
}

func (v *Value) AsInt16Slice() []int16 {
	return AsInt16Slice(v.Data)
}

func (v *Value) ToInt16Slice() ([]int16, error) {
	return ToInt16Slice(v.Data)
}

func (v *Value) AsInt32Slice() []int32 {
	return AsInt32Slice(v.Data)
}

func (v *Value) ToInt32Slice() ([]int32, error) {
	return ToInt32Slice(v.Data)
}

func (v *Value) AsInt64Slice() []int64 {
	return AsInt64Slice(v.Data)
}

func (v *Value) ToInt64Slice() ([]int64, error) {
	return ToInt64Slice(v.Data)
}

func (v *Value) AsUintSlice() []uint {
	return AsUintSlice(v.Data)
}

func (v *Value) ToUintSlice() ([]uint, error) {
	return ToUintSlice(v.Data)
}

func (v *Value) AsUint8Slice() []uint8 {
	return AsUint8Slice(v.Data)
}

func (v *Value) ToUint8Slice() ([]uint8, error) {
	return ToUint8Slice(v.Data)
}

func (v *Value) AsUint16Slice() []uint16 {
	return AsUint16Slice(v.Data)
}

func (v *Value) ToUint16Slice() ([]uint16, error) {
	return ToUint16Slice(v.Data)
}

func (v *Value) AsUint32Slice() []uint32 {
	return AsUint32Slice(v.Data)
}

func (v *Value) ToUint32Slice() ([]uint32, error) {
	return ToUint32Slice(v.Data)
}

func (v *Value) AsUint64Slice() []uint64 {
	return AsUint64Slice(v.Data)
}

func (v *Value) ToUint64Slice() ([]uint64, error) {
	return ToUint64Slice(v.Data)
}

func (v *Value) AsFloat32Slice() []float32 {
	return AsFloat32Slice(v.Data)
}

func (v *Value) ToFloat32Slice() ([]float32, error) {
	return ToFloat32Slice(v.Data)
}

func (v *Value) AsFloat64Slice() []float64 {
	return AsFloat64Slice(v.Data)
}

func (v *Value) ToFloat64Slice() ([]float64, error) {
	return ToFloat64Slice(v.Data)
}

func (v *Value) AsStringStringMap() map[string]string {
	return AsStringStringMap(v.Data)
}

func (v *Value) ToStringStringMap() (map[string]string, error) {
	return ToStringStringMap(v.Data)
}

func (v *Value) AsStringInterfaceMap() map[string]interface{} {
	return AsStringInterfaceMap(v.Data)
}

func (v *Value) ToStringInterfaceMap() (map[string]interface{}, error) {
	return ToStringInterfaceMap(v.Data)
}

func (v *Value) AsInterfaceInterfaceMap() map[interface{}]interface{} {
	return AsInterfaceInterfaceMap(v.Data)
}

func (v *Value) ToInterfaceInterfaceMap() (map[interface{}]interface{}, error) {
	return ToInterfaceInterfaceMap(v.Data)
}
