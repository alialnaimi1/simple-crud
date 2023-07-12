package main

import (
	"context"
	"fmt"
	"simple-crud/internal/controller/user"
	"simple-crud/internal/controller/user/repository"
	"simple-crud/internal/database"
)

func main() {

	ctx := context.Background()

	db, err := database.NewDatabase("root", "", "localhost", "simple_crud", 3306)
	if err != nil {
		panic(err)
	}
	userRepo := repository.NewMySQLRepository(db.DB)

	userService := user.NewDefaultService(userRepo)

	// use userService
	// C
	userService.Create(ctx, "username", "email", "password")

	// R
	user, err := userService.Read(ctx, 1)
	if err != nil {
		panic(err)
	}
	println(user.Username)

	// U
	userService.Update(ctx, "username2", "email2", "password2")

	// R
	user, err = userService.Read(ctx, 1)
	if err != nil {
		panic(err)
	}
	println(user.Username)

	// D
	userService.Delete(ctx, 1)

	fmt.Println("done")

}
