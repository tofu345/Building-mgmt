package scripts

import (
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/tofu345/Building-mgmt-backend/internal"
	"golang.org/x/term"
)

func getScript(name string) (Script, error) {
	for _, v := range scripts {
		if v.name == name {
			return v, nil
		}
	}

	return Script{}, fmt.Errorf("Script '%v' not found", name)
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	text, _ := r.ReadString('\n')
	if text == "exit" {
		os.Exit(0)
	}

	return strings.TrimSpace(text)
}

func adminLogin() (internal.User, error) {
	fmt.Println("! Admin Login Required")

	email := getUserInput("> Admin Email: ")
	admin := internal.User{}
	err := db.Where("email = ?", email).Find(&admin).Error
	if err != nil {
		return internal.User{}, err
	}

	if admin.ID == 0 {
		fmt.Printf("! No user found with email '%v'\n", email)
		return adminLogin()
	}

	fmt.Print("> Admin Password: ")
	password := readPassword()
	if !internal.CheckPasswordHash(password, admin.Password) {
		fmt.Println("! Incorrect password")
		return adminLogin()
	}

	if !admin.IsSuperuser {
		fmt.Printf("! %v does not have admin permissions\n", admin.Name())
		return adminLogin()
	}

	return admin, nil
}

// fatal prints the given error message and exits the program.
//
// It takes an error as a parameter and does not return anything.
func fatal(err error) {
	fmt.Printf("! %v\n", err)
	os.Exit(0)
}

func readPassword() string {
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fatal(err)
	}
	fmt.Println()
	return string(bytePassword)
}

// getPassword returns the user's password after validating it.
//
// It prompts the user to enter a password and retype it for confirmation.
// If the passwords do not match, an error message is displayed and the function is called recursively.
// The entered password is then hashed using the internal.HashPassword function.
// If an error occurs during the hashing process, the function calls the fatal function with the error.
// Finally, the hashed password is returned.
func getPassword() string {
	fmt.Print("> Password: ")
	password := readPassword()

	fmt.Print("> Retype Password: ")
	password2 := readPassword()

	if password != password2 {
		fmt.Println("! Passwords do not match")
		return getPassword()
	}

	password, err := internal.HashPassword(password)
	if err != nil {
		fatal(err)
	}

	return password
}
