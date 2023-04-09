/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package password

import (
	"encoding/json"
	"fmt"
	"os"
	"password_manager/cmd/utils"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete file with password",
	Long: `Delete file with password`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
	},
}

func init() {
	PasswordCmd.AddCommand(deleteCmd)

	searchCmd.Flags().StringVarP(&utils.Domain, "domain", "d", "", "user domain")
	searchCmd.Flags().StringVarP(&utils.Login, "login", "l", "", "user login")
	searchCmd.Flags().StringVarP(&utils.Password, "password", "p", "", "user password")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func deleteCredential() {
	var creds utils.Storage
	
	data, err1 := os.ReadFile("credentials.json")
	if err1 != nil {
		panic(err1)
	}

	err := json.Unmarshal(data, &creds)
	if err != nil {
		panic(err)
	}



	for _, v := range creds.Credentials {
		if v.Domain == utils.Domain || v.Login == utils.Login || v.Password == utils.Password {
		}
	}

}

