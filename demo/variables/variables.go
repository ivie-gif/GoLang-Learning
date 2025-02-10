package main

import "fmt"

func main() {
	var myName string = "Ivie"
	fmt.Println("This is my Name:", myName, "Maryann")

	userName := "Emwiongbon"
	fmt.Println("My surname is:", userName)

	partOne, partTwo := "james", "Admin"
	fmt.Println(partOne, partTwo)

	partTwo, Default := "mary", "Guest"
	fmt.Println(partTwo, Default)

	var (
		lessonName string = "Variables"
		lessonClass int = 12
		lessonDuration float32 = 3.5
		fee float32 = 1000
	)

	fmt.Println( "Lesson Name:", lessonName, "Class Age Grade:", lessonClass, "Duration:", lessonDuration, "Fee:", fee)
}
