package main

import (
	"fmt"
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
}

type User struct {
	gorm.Model
	Name  string
	Email string
	Age   int
}

func createUser(name string, email string, age int) {
	user := User{Name: name, Email: email, Age: age}
	result := DB.Create(&user)
	if result.Error != nil {
		log.Fatalf("failed to create user: %v", result.Error)
	}
	fmt.Println("User created successfully:", user)
}

func getUserByID(id uint) {
	var user User
	result := DB.First(&user, id)
	if result.Error != nil {
		log.Fatalf("failed to get user: %v", result.Error)
	}
	fmt.Println("User retrieved successfully:", user)
}

func updateUser(id uint, name string, email string, age int) {
	var user User
	DB.First(&user, id)
	user.Name = name
	user.Email = email
	user.Age = age
	DB.Save(&user)
	fmt.Println("User updated successfully:", user)
}

func deleteUser(id uint) {
	var user User
	DB.Delete(&user, id)
	fmt.Println("User deleted successfully")
}

func main() {
	// Migrate the schema
	DB.AutoMigrate(&User{})

	// Perform CRUD operations
	createUser("John Doe", "john@example.com", 30)
	getUserByID(1)
	updateUser(1, "John Doe Updated", "john_updated@example.com", 31)
	deleteUser(1)
}
