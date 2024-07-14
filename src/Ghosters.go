package main

func Ceaser(bytes []byte) []byte {
	for index, currbyte := range bytes {
		bytes[index] = byte((int64(currbyte) + 6) % 256)
	}
	return bytes
}
