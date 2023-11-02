package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/BearTS/go-echo/dbapp/pkg"
	"github.com/spf13/cobra"
)

func DropTables() *cobra.Command {
	return &cobra.Command{
		Use: "droptables",
		RunE: func(cmd *cobra.Command, args []string) error {

			dbConnection, sqlConnection := pkg.Connection()
			defer sqlConnection.Close()

			var tableNames []string
			if err := dbConnection.Table("information_schema.tables").
				Where("table_schema = ?", "public").Pluck("table_name", &tableNames).Error; err != nil {
				panic(err)
			}

			if os.Getenv("ENIRONMENT") != "development" {
				tablesToDrop := make(map[string]bool)
				// tablesToDrop["events"] = true
				// tablesToDrop["event_details"] = true
				// tablesToDrop["admins"] = true
				// tablesToDrop["merch"] = true

				fmt.Println("Warning: Environment is not development. Only Certain tables will be dropped")
				if len(tableNames) > 0 {
					for i, tableName := range tableNames {
						if tablesToDrop[tableName] {
							if err := dbConnection.Migrator().DropTable(tableName); err != nil {
								return errors.New("Error: While dropping tables:" + tableName)
							}
							fmt.Println("[", i, "]: ", "dropped table: ", tableName)
						}
					}

				}
				return nil
			} else {
				fmt.Println("App env is development")

				if len(tableNames) > 0 {
					for i, tableName := range tableNames {
						if err := dbConnection.Migrator().DropTable(tableName); err != nil {
							return errors.New("Error: While dropping tables:" + tableName)
						}
						fmt.Println("[", i, "]: ", "dropped table: ", tableName)
					}
				}
				fmt.Println("Dropped all tables sucessfully")
				return nil
			}
		},
	}
}
