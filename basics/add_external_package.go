package main

/**
Here we import an external package
We can let IntelliJ sync the dependency,
or we can use go command to download and add the dependency to the project. For knowledge, let's do this way:
   - in terminal, in the project root, run: `go get rsc.io/quote`
That's all you need to do. Go will download the package, and add it to the go.mod file automatically.
You can explore the external libraries of the current project in IntelliJ int he Project view.
Note that since I set the GOPATH env to ~/go, go actually downloaded the package there, under
     ~/go/pkg/mod/rsc.io
I found out about this, by right-clicking on the external library in IntelliJ, and chose open in Finder.
Note that in the newer versions of Go, we can use Module, like this project, and we are not dependent on GOPATH like in
the past. GOPATH now is merely used for go to store downloaded libraries. Previously, under GOPATH dir, there must be
a dir named "src" and all source code went there. This is no longer the case now. Projects now are outside GOPATH.

In the project root, you can run `go list -m all` to see all the dependencies.
*/
import (
	"fmt"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())
}
