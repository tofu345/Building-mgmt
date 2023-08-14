package scripts

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tofu345/Building-mgmt-backend/internal"
	"gorm.io/gorm"
)

type Script struct {
	name        string
	description string
	function    func()
}

var (
	db      *gorm.DB
	r       = bufio.NewReader(os.Stdin)
	scripts = []Script{
		{"create_admin", "Create Admin User", createAdmin},
	}
)

func init() {
	db = internal.GetDB()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
}

func Shell() {
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

			for _, v := range scripts {
				fmt.Printf("%v\t%v\n", v.name, v.description)
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
	admins := []internal.User{}
	err := db.Where("is_superuser <> ?", true).Find(&admins).Error
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
	password := getPassword()

	user := internal.User{
		FirstName:   first_name,
		LastName:    last_name,
		Password:    password,
		Email:       email,
		IsSuperuser: true,
	}

	err = db.Create(&user).Error
	if err != nil {
		fatal(err)
		return
	}

	fmt.Println("! Admin Created Successfully")
}
