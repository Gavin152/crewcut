package cmd

import (
	"database/sql"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var newCmd = &cobra.Command{
	Use:   "new [name of new crew]",
	Short: "Create a new crew trip",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		Db, _ := sql.Open("sqlite", "data.db")
		defer Db.Close()

		newCrew, _ := Db.Prepare("INSERT INTO crews (name) VALUES (?)")
		_, err := newCrew.Exec(args[0])
		if err != nil {
			fmt.Printf("Error creating new crew: %s\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
