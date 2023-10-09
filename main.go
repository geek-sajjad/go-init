package main

func main() {

	// variables

	// Explicit define
	// var age int //Declare and integer variable

	// age = 30

	// Type Inference
	// name := "my name"

	// isValid := true

	// var count int // 'count' is assigned the zero value for integers, which is 0

	// var isActive bool // 'isActive' is assigned the zero value for booleans, which is false

	// Multiple Variable Declaration

	// var a, b, c int

	// a, b, c = 1, 2, 3

	// Constants

	// const pi = 3.14

	// if statements

	// age2 := 25

	// if age2 >= 18 {
	// 	fmt.Println("You are an adult")
	// } else {
	// 	fmt.Println("You are a minor")
	// }

	// For loops

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// WHile loop
	// count2 := 0
	// for count2 < 5 {
	// 	fmt.Println(count2)
	// 	count2++
	// }

	// for index, value := range collection {
	// 	// Code to execute for each item in the collection
	// }

	// names := []string{"Alice", "Bob", "Charlie"}

	// for index, name := range names {
	// 	fmt.Println(index, name)
	// }

	// Arrays

	// var myArray [size]Type

	// var myArray [3]int
	// for i := 0; i < len(myArray); i++ {
	// 	fmt.Println(myArray[i])
	// }

	// for i, value := range myArray {
	// 	fmt.Println(i, value)
	// }

	// When you pass an array to a function, it's passed by value, meaning a copy of the array is made. To avoid copying large arrays, you can pass a pointer to the array or use slices instead

	// Slices

	// var mySlice []int

	// mySlice := []int{1, 2, 3, 4}

	// println(mySlice)

	// mySlice = append(mySlice, 5, 6)

	// When you pass a slice to a function, you're passing a reference to the underlying array. Any modifications made to the slice within the function will be reflected in the original slice

	// var nilSlice []int
	// fmt.Println(nilSlice)

	// Functions

	// result := add(3, 5)
	// a, b := swap(10, 20)

	//Function Types and Interfaces

	// Anonymous Functions (Closures):
	// increment := func(x, y int) int {
	// 	return x + y
	// }

	// increment(2, 3)

	// Variadic Functions
	// fmt.Println(sum(4, 6, 10))

	// Named return values

	// result := add(2, 3)
	// fmt.Println(result)

	// Interfaces
	// c := Circle{Radius: 5.0}
	// PrintShapeInfo(c)

	// Structs
	// person1 := Person{
	// 	firstName: "Sajad",
	// 	lastName:  "Sohrabi",
	// 	age:       24,
	// }

	// fmt.Println(person1)
	// fmt.Println(person1.fullName())
}

// Function Types and Interfaces
// type BinaryOperation func(int, int) int

// func add(x int, y int) int {
// 	return x + y
// }
// func swap(x, y int) (int, int) {
// 	return y, x
// }

// func calculate(operation BinaryOperation, x, y int) int {
// 	return operation(x, y)
// }

// Variadic Functions
// func sum(numbers ...int) int {
// 	// var total int
// 	total := 0

// 	for _, num := range numbers {
// 		total += num
// 	}

// 	return total
// }

// Named return values
// func add(x, y int) (result int) {
// 	result = x + y

// 	return
// }

// Interfaces
// type Shape interface {
// 	Area() float64
// 	Perimeter() float64
// }

// type Circle struct {
// 	Radius float64
// }

// func (c Circle) Area() float64 {
// 	Pi := 3.14
// 	return Pi * c.Radius * c.Radius
// }

// func (c Circle) Perimeter() float64 {
// 	Pi := 3.14
// 	return 2 * Pi * c.Radius
// }
// func PrintShapeInfo(s Shape) {
// 	fmt.Printf("Area: %f\n", s.Area())
// 	fmt.Printf("Perimeter: %f\n", s.Perimeter())
// }

// Structs
// type Person struct {
// 	firstName string
// 	lastName  string
// 	age       int
// }

// func (p Person) fullName() string {
// 	return p.firstName + " " + p.lastName
// }
