package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// decCmd represents the dec command
var decCmd = &cobra.Command{
	Use:   "dec [filepath]",
	Short: "Prompts for a password and decrypts the given file",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalln("Missing filepath")
			os.Exit(1)
		}

		path := filepath.Clean(args[0])

		// attempt to read file
		ciphertext, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("Error reading file: %s\n", err.Error())
			os.Exit(1)
		}

		fmt.Printf("Enter password: ")
		password, err := term.ReadPassword(int(os.Stdin.Fd()))
		fmt.Println() // feed line
		if err != nil {
			log.Fatalf("Error reading password: %s\n", err.Error())
			os.Exit(1)
		}

		cleartext, err := AesDecryptWithPassword(string(password), ciphertext)
		if err != nil {
			log.Fatalf("Error decrypting file: %s\n", err.Error())
			os.Exit(1)
		}

		var filename string
		base := filepath.Base(path)
		if ext := filepath.Ext(path); ext == ".enc" && strings.Count(base, ".") > 1 {
			filename = base[:len(base)-4]
		} else {
			filename = base
		}

		outfile := filepath.Join(filepath.Dir(path), filename)
		if err := os.WriteFile(outfile, cleartext, 0600); err != nil {
			log.Fatalf("Error writing decrypted output file: %s\n", err.Error())
			os.Exit(1)
		}

		if delete {
			os.Remove(path)
		}
	},
}

func init() {
	rootCmd.AddCommand(decCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	decCmd.Flags().BoolVarP(&delete, "delete", "d", false, "Delete the input file on success")
}
