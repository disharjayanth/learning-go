package benchmark

import "os"

func FileLen(f string, buffsize int) (int, error) {
	file, err := os.Open(f)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	count := 0
	buff := make([]byte, buffsize)
	for {
		n, err := file.Read(buff)
		if err != nil {
			break
		}
		count += n
	}

	return count, nil
}
