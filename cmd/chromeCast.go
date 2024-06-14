/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/vishen/go-chromecast/application"
	"github.com/vishen/go-chromecast/dns"
)

// chromeCastCmd represents the chromeCast command
var chromeCastCmd = &cobra.Command{
	Use:   "chromeCast",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: displayPickachu,
}

func init() {
	rootCmd.AddCommand(chromeCastCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chromeCastCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	chromeCastCmd.Flags().StringP("name", "n", "", "Name of the device")
}

func displayPickachu(cmd *cobra.Command, args []string) {
	name := cmd.Flag("name").Value.String()
	if name == "" {
		discoverChromeCastDevices()
	} else {
		connectToDevice(name)
	}
}

func discoverChromeCastDevices() {
	devices, err := listOfDevices()
	if err != nil {
		fmt.Println("Error discovering devices: ", err)
		return
	}

	for _, device := range devices {
		fmt.Println(device.DeviceName, device.GetAddr(), device.GetPort(), device.GetUUID(), device.Device)
	}
}

func connectToDevice(name string) {
	var device *dns.CastEntry
	devices, err := listOfDevices()
	if err != nil {
		fmt.Println("Error discovering devices: ", err)
		return
	}

	for _, d := range devices {
		if strings.EqualFold(d.DeviceName, name) {
			device = &d
			break
		}
	}

	if device == nil {
		fmt.Println("Device not found")
		return
	}

	fmt.Println("Connecting to device", device.DeviceName)
	app := application.NewApplication()
	if err = app.Start(device.GetAddr(), device.GetPort()); err != nil {
		fmt.Println("Error starting application: ", err)
		return
	}

	err = app.Load(
		"https://media1.giphy.com/media/v1.Y2lkPTc5MGI3NjExYWx1bGxyY2piZmVzeTVkcDR0c2JrZW1td2c5ZjlmaGswcmF4bWhsZCZlcD12MV9naWZzX3NlYXJjaCZjdD1n/lOa0tPKiMLdqVdFiS8/giphy.webp",
		0,
		"",
		false,
		false,
		false,
	)
	if err != nil {
		fmt.Println("Error loading media: ", err)
		return
	}
}

func listOfDevices() ([]dns.CastEntry, error) {
	// Create a context with a timeout of 15 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(5))
	defer cancel()

	castEntryChan, err := dns.DiscoverCastDNSEntries(ctx, nil)
	if err != nil {
		return nil, err
	}
	devices := []dns.CastEntry{}
	for d := range castEntryChan {
		devices = append(devices, d)
	}

	return devices, nil
}
