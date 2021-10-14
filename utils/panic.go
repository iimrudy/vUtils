package utils

func PanicIfNull(t interface{}) {
	if t == nil {
		panic("Given items is null")
	}
}
