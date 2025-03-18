package database

import (
	"astrogo-cli/internal/db"
	"log"

	"github.com/spf13/cobra"
)

// NewDatabaseCmd creates the database command
func NewDatabaseCmd() *cobra.Command {
	dbCmd := &cobra.Command{
		Use:   "db",
		Short: "Database service commands",
	}

	var port int
	startDBCmd := &cobra.Command{
		Use:   "start",
		Short: "Start the database service",
		Run: func(cmd *cobra.Command, args []string) {
			if err := db.StartServer(port); err != nil {
				log.Fatalf("Failed to start database service: %v", err)
			}
		},
	}

	startDBCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to run the database service on")

	// Add subcommands to db command
	dbCmd.AddCommand(startDBCmd, NewTestDBCmd())

	return dbCmd
}
