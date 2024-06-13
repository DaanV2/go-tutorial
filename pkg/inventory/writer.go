package inventory

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"os"
	"path"
	"strconv"

	"github.com/DaanV2/go-tutorial/pkg/rand"
)

func ToFile(p *Receipt, folder string) error {
	if os.MkdirAll(folder, 0755) != nil {
		return errors.New("could not create folder")
	}

	filename := rand.RandomID()
	filepath := path.Join(folder, filename)

	// 50/50 csv or json
	if rand.RandomBool() {
		return ToCsv(p, filepath+".csv")
	}

	return ToJson(p, filepath+".json")
}

var header = []string{"id", "quantity", "note", "price", "tax", "time"}

func ToCsv(r *Receipt, path string) error {
	// Open the file for writing
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a new CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()
	if err := writer.Write(header); err != nil {
		return err
	}

	for _, item := range r.Items {
		line := []string{
			item.ID,
			strconv.FormatInt(item.Quantity, 10),
			item.Note,
			strconv.FormatFloat(item.Price, 'f', -1, 64),
			strconv.FormatFloat(item.Tax, 'f', -1, 64),
			item.Time,
		}
		if err := writer.Write(line); err != nil {
			return err
		}
	}

	return nil
}

func ToJson(r *Receipt, path string) error {
	// Open the file for writing
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(r)
}
