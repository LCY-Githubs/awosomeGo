package data_base_save

import "os"

func SaveData1(path string, data []byte) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err1 := file.Write(data)
	return err1
}
