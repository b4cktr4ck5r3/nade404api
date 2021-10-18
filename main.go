package main

import (
	"log"

	"github.com/b4cktr4ck5r3/nade404api/database"
	"github.com/b4cktr4ck5r3/nade404api/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// if len(os.Args) < 6 {
	// 	log.Fatal("Not enought args, need to run this programme like this : ./script {UserDB} {PasswordDB} {HostDB} {PortDB} {NameDB}")
	// }

	// userDB := os.Args[1]
	// pwdDB := os.Args[2]
	// hostDB := os.Args[3]
	// portDB := os.Args[4]
	// nameDB := os.Args[5]

	// if err := database.ConnectWithArgs(userDB, pwdDB, hostDB, portDB, nameDB); err != nil {
	// 	log.Fatal(err)
	// }

	if err := database.ConnectWithEnv(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	router.SetupRoutes(app)
	app.Listen(":3000")
}
