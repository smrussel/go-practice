package fileops

import (
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(filename string) (float64, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return 1000, fmt.Errorf("failed to find file")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 1000, fmt.Errorf("failed to parse stored value")
	}
	return value, nil
}

// upper case function can be only import

func WriteFloatToFile(value float64, filename string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(filename, []byte(valueText), 0644)
}
