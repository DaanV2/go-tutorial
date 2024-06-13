/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"path"

	"github.com/DaanV2/go-tutorial/pkg/inventory"
	"github.com/DaanV2/go-tutorial/pkg/rand"
	"github.com/spf13/cobra"
)

// generateCollectCmd represents the generateCollect command
var generateCollectCmd = &cobra.Command{
	Use:   "generate-collect",
	Aliases: []string{"gen-collect"},
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: GenerateCollect,
}

func init() {
	rootCmd.AddCommand(generateCollectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCollectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCollectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func GenerateCollect(cmd *cobra.Command, args []string) {
	fmt.Println("generating called")

	dataFolder := path.Join("data", "receipts")

	for range 7 {
		id := rand.RandomID()
		generateFolder(path.Join(dataFolder, id))
	}
}

func generateFolder(folder string) {
	fmt.Println("Generating folder", folder)

	amount := rand.Int64(5) + 3

	for range amount {
		receipt := inventory.Receipt{}
		items := rand.Int64(42) + 7

		for range items {
			receipt.AddProduct(inventory.GenerateProduct())
		}

		err := inventory.ToFile(&receipt, folder)
		if err != nil {
			fmt.Println("Error writing file", err)
		}
	}
}