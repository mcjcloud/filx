package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// encCmd represents the enc command
var encCmd = &cobra.Command{
	Use:   "enc [filepath]",
	Short: "Prompts for a password and encrypts the given file using AES",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalln("Missing filepath")
		}

		path := filepath.Clean(args[0])

		// attempt to read file
		cleartext, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("Error reading file: %s\n", err.Error())
		}

		fmt.Printf("Enter a password: ")
		password, err := term.ReadPassword(int(os.Stdin.Fd()))
		fmt.Println() // feed line
		if err != nil {
			log.Fatalf("Error reading password: %s\n", err.Error())
		}

		fmt.Printf("Repeat password: ")
		repeatedPw, err := term.ReadPassword(int(os.Stdin.Fd()))
		fmt.Println()
		if err != nil {
			log.Fatalf("Error reading password: %s\n", err.Error())
		}

		if string(password) != string(repeatedPw) {
			log.Fatalf("Passwords do not match.")
		}

		ciphertext, err := AesEncryptWithPassword(string(password), cleartext)
		if err != nil {
			log.Fatalf("Error encrypting file: %s\n", err.Error())
		}

		outfile := filepath.Join(filepath.Dir(path), fmt.Sprintf("%s.enc", filepath.Base(path)))
		if err := os.WriteFile(outfile, ciphertext, 0600); err != nil {
			log.Fatalf("Error writing encrypted output file: %s\n", err.Error())
		}

		if delete {
			os.Remove(path)
		}
	},
}

func init() {
	rootCmd.AddCommand(encCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	encCmd.Flags().BoolVarP(&delete, "delete", "d", false, "Delete the input file on success")
}
