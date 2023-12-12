package cmd

import (
	"fmt"
	"math/rand"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a password",
	Long: `Generate a password using the following flags:
	
	for exaample:

	passwordGen generate -l 12 -d -s`,
	Run: generatePassword,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().IntP("length", "l", 8, "Length of the generated password")
	generateCmd.Flags().BoolP("digits", "d", false, "Include digits in the generated password")
	generateCmd.Flags().BoolP("special-chars", "s", false, "Include special characters in the generated password")
}

func generatePassword(cmd *cobra.Command, args []string) {
	length, _ := cmd.Flags().GetInt("length")
	isDigits, _ := cmd.Flags().GetBool("digits")
	isSpecialChars, _ := cmd.Flags().GetBool("special-chars")

	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	if isDigits {
		charset += "0123456789"
	}

	if isSpecialChars {
		charset += "!@#$%^&*()_+{}[]|;:,.<>?-="
	}

	password := make([]byte, length)

	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	fmt.Println("Generating password...")
	fmt.Println(string(password))

	err := clipboard.WriteAll(string(password))
	if err != nil {
		fmt.Println("Error copying to clipboard:", err)
	} else {
		fmt.Println("Password copied to clipboard!")
	}
}
