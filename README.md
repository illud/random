# random
Random string and int generator for Go programming

#Installation
  1. Get the package
  
    $ go get github.com/saturnavt/random
    
  2. Import it in your code:
  
    import "github.com/saturnavt/random"

#Quick start

Random String

    package main

    import (
      "fmt"
      "github.com/saturnavt/random"
    )

    func main() {
      fmt.Println(random.RandomString(12)) //Change 12 for your own string length, output example 'rlJlGUgEpLnE'
    }
    
random Integer

    package main

    import (
      "fmt"
      "github.com/saturnavt/random"
    )

    func main() {
      fmt.Println(random.RandomNumber(80, 0)) //80 is the maximun number and 0 the minimun number,change for your own max and minimun, output example '25'
    }
