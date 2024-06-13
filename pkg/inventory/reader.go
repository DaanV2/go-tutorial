package inventory

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"io"
	"os"
	"path"
	"slices"
	"strconv"
)

func ReadFile(filepath string) (*Receipt, error) {
	ext := path.Ext(filepath)

	if ext == ".csv" {
		return FromCsv(filepath)
	}

	if ext == ".json" {
		return FromJson(filepath)
	}

	return nil, errors.New("unsupported file type")
}

func FromJson(filepath string) (*Receipt, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a new JSON decoder
	decoder := json.NewDecoder(file)
	var receipt Receipt

	if err := decoder.Decode(&receipt); err != nil {
		return nil, err
	}

	return &receipt, nil
}

func FromCsv(filepath string) (*Receipt, error) {
	var (
		errs error
		err  error
	)

	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)
	receipt := &Receipt{
		Items: make([]Product, 0),
	}


	for {
		line, err := reader.Read()
		// End Of File?
		if errors.Is(err, io.EOF) {
			return receipt, nil
		}
		if err != nil {
			return receipt, err
		}
		// Skip the header
		if slices.Equal(line, header) {
			continue
		}
		if len(line) < len(header) {
			return receipt, errors.New("invalid CSV")
		}

		item := Product{
			ID:   line[0],
			Note: line[2],
			Time: line[5],
		}

		item.Quantity, err = strconv.ParseInt(line[1], 10, 64)
		errs = errors.Join(errs, err)
		item.Price, err = strconv.ParseFloat(line[3], 64)
		errs = errors.Join(errs, err)
		item.Tax, err = strconv.ParseFloat(line[4], 64)
		errs = errors.Join(errs, err)
		if errs != nil {
			return receipt, errs
		}

		receipt.Items = append(receipt.Items, item)
	}
}