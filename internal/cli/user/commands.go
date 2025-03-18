package user

import (
	"astrogo-cli/internal/db"
	"fmt"

	"github.com/spf13/cobra"
)

// NewUserCmd creates the user command and its subcommands
func NewUserCmd() *cobra.Command {
	userCmd := &cobra.Command{
		Use:   "user",
		Short: "User management commands",
	}

	var username, email string
	createUserCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new user",
		Run: func(cmd *cobra.Command, args []string) {
			if username == "" || email == "" {
				fmt.Println("Error: Both username and email are required")
				return
			}

			user, err := db.CreateUser(username, email)
			if err != nil {
				fmt.Printf("Error creating user: %v\n", err)
				return
			}

			fmt.Printf("User created successfully:\nID: %d\nUsername: %s\nEmail: %s\n",
				user.ID, user.Username, user.Email)
		},
	}

	listUsersCmd := &cobra.Command{
		Use:   "list",
		Short: "List all users",
		Run: func(cmd *cobra.Command, args []string) {
			users, err := db.GetAllUsers()
			if err != nil {
				fmt.Printf("Error listing users: %v\n", err)
				return
			}

			fmt.Println("Users:")
			for _, user := range users {
				fmt.Printf("- ID: %d, Username: %s, Email: %s\n",
					user.ID, user.Username, user.Email)
			}
		},
	}

	// Add flags to the create user command
	createUserCmd.Flags().StringVarP(&username, "username", "u", "", "Username")
	createUserCmd.Flags().StringVarP(&email, "email", "e", "", "Email address")

	// Add subcommands to user command
	userCmd.AddCommand(createUserCmd, listUsersCmd)

	return userCmd
}
