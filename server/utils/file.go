package utils

import "os"

func ReadFromFile(fname string) (string, error) {
	if fl, err := os.Open(fname); err != nil {
		return "", err
	} else {
		defer fl.Close()

		chunks := make([]byte, 1024, 1024)
		var buffer []byte
		var length int
		for {
			if n, err := fl.Read(chunks); n == 0 {
				break
			} else if err != nil {
				return "", err
			} else {
				length += n
				buffer = append(buffer, chunks...)
			}
		}
		return string(buffer[:length]), nil
	}
}