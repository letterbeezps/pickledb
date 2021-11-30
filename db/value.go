package db

// N --> noraml type like int, string
// L --> list
// D --> dict
type value struct {
	Data     interface{}
	DataType string
}

func (v *value) getType() string {
	return v.DataType
}

func newValue(data interface{}, dataType string) *value {
	return &value{
		Data:     data,
		DataType: dataType,
	}
}
