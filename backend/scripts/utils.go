package scripts

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"syscall"

	m "github.com/tofu345/Building-mgmt-backend/src/models"
	s "github.com/tofu345/Building-mgmt-backend/src/services"
	"golang.org/x/term"
	"gorm.io/gorm"
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

func adminLogin() (m.User, error) {
	if loggedInAdmin.ID != 0 {
		return loggedInAdmin, nil
	}

	fmt.Println("! Admin Login Required")

	email := getUserInput("> Admin Email: ")
	admin, err := m.GetUserByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("! No user found with email '%v'\n", email)
		} else {
			fatal(err)
		}
		return adminLogin()
	}

	fmt.Print("> Admin Password: ")
	password := readPassword()
	if !s.CheckPasswordHash(password, admin.Password) {
		fmt.Println("! Incorrect password")
		return adminLogin()
	}

	if !admin.IsSuperuser {
		fmt.Printf("! %v does not have admin permissions\n", admin.Name())
		return adminLogin()
	}

	loggedInAdmin = admin

	return admin, nil
}

// fatal prints the given error message and exits the program.
//
// It takes an error as a parameter and does not return anything.
func fatal(err error) {
	fmt.Printf("! %v\n", s.ParseError(err))
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

func getAndComparePasswords() string {
	fmt.Print("> Password: ")
	password := readPassword()

	fmt.Print("> Retype Password: ")
	password2 := readPassword()

	if password != password2 {
		fmt.Println("! Passwords do not match")
		return getAndComparePasswords()
	}

	return password
}

func printValidationErrors(errs map[string]string) {
	longest := 0
	for k := range errs {
		if len(k) > longest {
			longest = len(k)
		}
	}

	for k, v := range errs {
		fmt.Printf("! %v", k)
		if len(k) < longest {
			fmt.Printf(strings.Repeat(" ", longest-len(k)))
		}
		fmt.Printf("\t%v\n", v)
	}
}
