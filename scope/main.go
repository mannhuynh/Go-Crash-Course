package main

// import module/packagename
import "myapp/packageone"
// package variable
var myVar = "Package variable"

func main(){
	var blockVar = "Block variable in main"

	packageone.PrintMe(myVar, blockVar, packageone.PackageVar)

}
