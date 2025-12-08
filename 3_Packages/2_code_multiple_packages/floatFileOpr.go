package fileOprs

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func writeFloatToFile(fileName string, Float_val float64) {

	val_string := fmt.Sprint(Float_val)
	os.WriteFile(fileName, []byte(val_string), 0644)

}

func readFloatFromFile(fileName string) (float64, error) {

	byte_data, err := os.ReadFile(fileName)

	if err != nil {
		return 0, errors.New("no Such File Found")
	}
	balance_text := string(byte_data)                       //byte --> String
	floatValue, err := strconv.ParseFloat(balance_text, 64) //string --> float

	if err != nil {
		return 0, errors.New("text Can't be Converted to Float")
	}
	return floatValue, nil
}
