package cmd

import (
	"github.com/spf13/cobra"
)

func main() {
	db()
}

func db() {
	cmd := &cobra.Command{}

	cmd.AddCommand(DropTables())
	cmd.AddCommand(Migrate())
	cmd.AddCommand(Seed())
	cmd.AddCommand(Backup())
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
