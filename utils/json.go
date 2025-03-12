package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

// saves JSON to a file
func SaveToFile(data interface{}, filename string) error {

	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	return err
}

// generic function to load data from a JSON file into a slice of a specified type.
func LoadFromFile(filename string, resultSlice interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// get file info, later used to determine the size of the file
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	// read the file into a byte slice
	byteValue := make([]byte, fileInfo.Size())
	_, err = file.Read(byteValue)
	if err != nil {
		return err
	}

	// we create a new instance of the result type
	resultType := reflect.TypeOf(resultSlice)
	if resultType.Kind() != reflect.Ptr || resultType.Elem().Kind() != reflect.Slice {
		return fmt.Errorf("resultSlice must be a pointer to a slice")
	}

	// unmarshal JSON data into the result slice
	err = json.Unmarshal(byteValue, resultSlice)
	if err != nil {
		return err
	}

	return nil
}
