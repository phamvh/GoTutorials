package main

/**
See here: https://go.dev/tour/basics/3#:~:text=In%20Go%2C%20a%20name%20is,only%20to%20its%20exported%20names.
and here: https://stackoverflow.com/questions/37840981/in-go-is-it-convention-to-capitalize-type-names

Basically, any value or func that starts with a capital letter is exported outside the package and can be used
outside the package. Similar to public in Java.

However, if they start with a lower case letter, then they can be only used within the package.
This is similar to default access in Java.

No examples needed to clarify this.
*/

const Exported = "public things"
const nonExported = "private things"

// NOw assume we are in a different package
/**
import "PATH/a"

func main() {
  // This is fine
  println(a.Exported)

  // However, this won't compile. nonExported won't be recognized.
  // println(a.nonExported)
}
*/
