package packageone

import "fmt"

// public variable from a package
var PackageVar = "Public variable from packageone"

// public function from a package
func PrintMe(myVar, blockVar, PackageVar string){
	fmt.Println(myVar, blockVar, PackageVar)
}
