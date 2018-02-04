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
	return AsString(d.Value)
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
	return AsInt(d.Value)
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
	return AsUint(d.Value)
}

func (d *Data) ToUint16() (uint16, error) {
	return ToUint16(d.Value)
}

func (d *Data) AsUint32() uuint32 {
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
