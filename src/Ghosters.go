package main

func Ceaser(bytes []byte) []byte {
	for index, currbyte := range bytes {
		bytes[index] = byte((int(currbyte) + index) % 256)
	}
	return bytes
}

func ReverseCeaser(bytes []byte) []byte {
	for index, currbyte := range bytes {
		bytes[index] = byte((int(currbyte) + (256 - index)) % 256)
	}
	return bytes
}
