/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/DaanV2/go-tutorial/pkg/inventory"
	"github.com/spf13/cobra"
)

// collectCmd represents the collect command
var collectCmd = &cobra.Command{
	Use:   "collect",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: collectAllReceipts,
}

func init() {
	rootCmd.AddCommand(collectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// collectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// collectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Client represents a client with multiple receipts
type Client struct {
	Receipts []*inventory.Receipt // Receipts is a list of receipts for the client
	Id       string               // Id is the unique identifier of the client
}

type Overview struct {
	ToBuy        map[string]int64 // ToBuy is a map of product IDs to the number of items to buy
	maxKeyLength int              // maxKeyLength is the length of the longest key in the map, used for printing
}

// NewOverview creates a new Overview value
func NewOverview() *Overview {
	return &Overview{
		ToBuy: make(map[string]int64),
	}
}

// AddReceipts adds multiple receipts to the overview
func (o *Overview) AddReceipts(receipts []*inventory.Receipt) {
	for _, r := range receipts {
		o.AddReceipt(r)
	}
}

// AddReceipt adds a single receipt to the overview
func (o *Overview) AddReceipt(r *inventory.Receipt) {
	for _, item := range r.Items {
		id := item.ID
		o.maxKeyLength = max(o.maxKeyLength, len(id))

		// If the item is not in the map, add it, otherwise increment the quantity
		if v, ok := o.ToBuy[id]; !ok {
			o.ToBuy[id] = item.Quantity
		} else {
			o.ToBuy[id] = v + item.Quantity
		}
	}
}

// Print prints the overview to the console
func (o *Overview) Print(title string) {
	fmt.Printf("==== %s ====\n", title)
	for k, v := range o.ToBuy {
		padding := o.maxKeyLength - len(k)
		fmt.Printf("%s:%s    %d\n", k, strings.Repeat(" ", padding), v)
	}

	fmt.Println()
}

// collectAllReceipts collects all receipts in the data/receipts folder
func collectAllReceipts(cmd *cobra.Command, args []string) {
	fmt.Println("collecting receipts")

	// Collect all files in data/receipts, and all subfolders
	folder := path.Join("data", "receipts")
	client, err := collectClients(folder)
	if err != nil {
		fmt.Println("could not collect clients:", err)
		return
	}
	fmt.Println("collected clients", len(client))

	// Processing the receipts
	total := NewOverview()
	clientTotal := make(map[string]*Overview)

	for _, c := range client {
		overview := NewOverview()
		overview.AddReceipts(c.Receipts)
		total.AddReceipts(c.Receipts)
		clientTotal[c.Id] = overview
	}

	// Print the results
	total.Print("Total")
	for k, v := range clientTotal {
		v.Print("client: " + k)
	}
}

// collectClients collects all clients in a folder, excepts the folders to be client ids
func collectClients(folder string) ([]*Client, error) {
	files, err := os.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	result := make([]*Client, 0)

	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		c := &Client{
			Id: file.Name(),
		}
		result = append(result, c)
		receipts, err := collectReceipts(path.Join(folder, file.Name()))
		if err != nil {
			return nil, err
		}
		c.Receipts = receipts
	}

	return result, nil
}

// collectReceipts collects all receipts in a folder
func collectReceipts(folder string) ([]*inventory.Receipt, error) {
	files, err := os.ReadDir(folder)
	if err != nil {
		return nil, err
	}
	result := make([]*inventory.Receipt, 0)

	for _, file := range files {
		if !file.IsDir() {
			filepath := path.Join(folder, file.Name())
			receipt, err := loadFile(filepath)
			if err != nil {
				return nil, err
			}
			result = append(result, receipt)
		}
	}

	return result, nil
}

// loadFile loads a file into a receipt, based on the file extension
func loadFile(file string) (*inventory.Receipt, error) {
	switch path.Ext(file) {
	case ".json":
		return inventory.FromJson(file)
	case ".csv":
		return inventory.FromCsv(file)
	}
	return nil, nil
}