/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path"

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
	Run: collectReceipts,
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

func collectAllReceipts(cmd *cobra.Command, args []string) {
	fmt.Println("collecting receipts")

	// Collect all files in data/receipts, and all subfolders
	folder := path.Join("data", "receipts")
	files, err := collectFolder(folder)
	if err != nil {
		fmt.Println("could not collect files:", err)
		return
	}

	// Load all the files
	fmt.Println("found", len(files), "files")
	receipts, err := loadFiles(files)
	if err != nil {
		fmt.Println("could not load files:", err)
		return
	}

	// Processing the receipts
	toBuy := make(map[string]int64)
	for _, receipt := range receipts {
		for _, item := range receipt.Items {
			if v, ok := toBuy[item.ID]; !ok {
				toBuy[item.ID] = item.Quantity
			} else {
				toBuy[item.ID] = v + item.Quantity
			}
		}
	}

	printMap("to buy", toBuy)
}

type Client struct {
	Receipts []*inventory.Receipt
	Id       string
}

func printMap[K comparable, T any](title string, m map[K]T) {
	fmt.Println("====", title, "====")
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("")
}

func loadFiles(files []string) ([]*inventory.Receipt, error) {
	receipts := make([]*inventory.Receipt, 0)
	for _, file := range files {
		r, err := loadFile(file)
		if err != nil {
			return nil, err
		}
		receipts = append(receipts, r)
	}
	return receipts, nil
}

func loadFile(file string) (*inventory.Receipt, error) {
	switch path.Ext(file) {
	case ".json":
		return inventory.FromJson(file)
	case ".csv":
		return inventory.FromCsv(file)
	}
	return nil, nil
}

func collectClient(folder string) ([]*Client, error) {
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
		receipts, err := collectReceipts(path.Join(folder, file.Name()))
		if err != nil {
			return nil, err
		}
		c.Receipts = receipts
	}

	return result, nil
}

func collectReceipts(folder string) ([]*inventory.Receipt, error) {
	files, err := os.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	result := make([]*inventory.Receipt, 0)

	for _, file := range files {
		if !file.IsDir() {
		} else {
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
