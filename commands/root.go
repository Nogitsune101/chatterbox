package commands

import (
	"log"

	"github.com/spf13/cobra"

	"chatterbox/server"
)

// ServerIP Specifies the root address of the api server
var serverIP string
var serverPort string

var rootCmd = &cobra.Command{
	Use:     "chatterbox",
	Short:   "Secure Portable Mini Chat Server",
	Version: "1.0.0b",
	Run: func(cmd *cobra.Command, args []string) {
		server.StartServer(serverIP + ":" + serverPort)
	},
}

// Execute is main entry point and launches the chat server
func Execute() {

	rootCmd.Flags().StringVarP(&serverIP, "ipaddr", "a", "0.0.0.0", "IP Address of the webserver (default: 0.0.0.0)")
	rootCmd.Flags().StringVarP(&serverPort, "port", "p", "8080", "Port of the webserver (default: 8080)")

	if err := rootCmd.Execute(); err != nil {
		log.Println("Chatterbox has encountered an error:", err)
	}
}
