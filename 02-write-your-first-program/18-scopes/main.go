package main

import "fmt" // File scoped (Only visible in this package)

const ok = true // package scope (visible to all the files belong to the package)

// package scope (visible to all the files belong to the package)
func main() {
	var hello = "Hello!" // block scope declaration
	fmt.Println(hello, ok)
	hey()
	bye()
}
