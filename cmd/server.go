/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: serverWorkload,
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	serverCmd.Flags().StringP("address", "a", "localhost:8080", "Server address")
}

func serverWorkload(cmd *cobra.Command, args []string) {
	address := cmd.Flag("address").Value.String()
	apiRouter := http.NewServeMux()
	router := http.NewServeMux()
	router.Handle("/api/", http.StripPrefix("/api", apiRouter))

	// Serve API
	apiRouter.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		d := time.Now().Format(time.RFC3339)

		_, err := w.Write([]byte(d))
		if err != nil {
			fmt.Println("Error while writing response: ", err)
		}
	})

	// Serve file in ./static folder
	router.Handle("/", http.FileServer(http.Dir("./static")))

	server := http.Server{
		Addr:    address,
		Handler: &Logger{
			Next: router,
		},
	}

	go func() {
		fmt.Println("Server started at ", address)
		defer fmt.Println("Server closed")

		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				fmt.Println("Server error: ", err)
			}
		}
	}()

	// Await for signal to close server
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-signals
	fmt.Println("Server is shutting down")

	// Graceful shutdown
	err := server.Shutdown(context.Background())
	if err != nil {
		fmt.Println("Error while shutting down server: ", err)
	}
	if err := server.Close(); err != nil {
		fmt.Println("Error while closing server: ", err)
	}
}

type Logger struct {
	Next http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[%s] %s %s\n", time.Now().Format(time.DateTime), r.Method, r.URL.Path)
	l.Next.ServeHTTP(w, r)
}