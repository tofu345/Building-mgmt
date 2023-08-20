package scripts

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	m "github.com/tofu345/Building-mgmt-backend/src/models"
	s "github.com/tofu345/Building-mgmt-backend/src/services"
)

type Script struct {
	name        string
	description string
	function    func()
}

var (
	r       = bufio.NewReader(os.Stdin)
	scripts = []Script{
		{"create_admin", "Create Admin User", createAdmin},
	}
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
}

func Shell(args ...string) {
	if len(args) > 0 {
		script, err := getScript(args[0])
		if err != nil {
			fatal(err)
		}

		script.function()
		return
	}

	fmt.Println("? 'help' to view list of commands")

	for {
		input := getUserInput("> ")

		switch input {
		case "":
			continue
		case "help":
			fmt.Println("list\tlist all commands")
			fmt.Println("exit\tquit")
		case "list":
			if len(scripts) == 0 {
				fmt.Println("! There are no scripts")
				return
			}

			for _, script := range scripts {
				fmt.Printf("%v\t%v\n", script.name, script.description)
			}
		case "ex", "exit":
			return
		default:
			script, err := getScript(input)
			if err != nil {
				fmt.Printf("! %v\n", err)
				continue
			}

			script.function()
		}
	}
}

func createAdmin() {
	admins, err := m.GetAdmins()
	if err != nil {
		fatal(err)
	}

	if len(admins) != 0 {
		_, err := adminLogin()
		if err != nil {
			fatal(err)
		}
	}

	fmt.Println(">> Create Admin user")

	first_name := getUserInput("> First Name: ")
	last_name := getUserInput("> Last Name: ")
	email := getUserInput("> Email: ")
	password := getAndComparePasswords()

	user := m.User{
		FirstName:   first_name,
		LastName:    last_name,
		Password:    password,
		Email:       email,
		IsSuperuser: true,
	}

	errs := s.ValidateModel(user)
	if errs != nil {
		formatValidationErrors(errs)
		_, ok := errs["Password"]
		if len(errs) == 1 && ok {
			proceed := getUserInput("> Save anyway? (y/n): ")
			if proceed != "y" {
				return
			}
		} else {
			return
		}

	}

	user.Password, err = s.HashPassword(user.Password)
	if err != nil {
		fatal(err)
	}

	err = m.CreateUser(&user)
	if err != nil {
		fatal(err)
	}

	fmt.Println("! Admin Created Successfully")
}
