package encoding

func Encode(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func Decode(data []byte, out interface{}) {
	json.Unmarshal(data, &out)
}
