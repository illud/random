# random
Random string and int generator for Go programming

#Installation
  1. Get the package
  
    $ go get github.com/illud/random
    
  2. Import it in your code:
  
    import "github.com/illud/random"

#Quick start

Random String

    package main

    import (
      "fmt"
      "github.com/illud/random"
    )

    func main() {
      fmt.Println(random.RandomString(12)) //Change 12 for your own string length, output example 'rlJlGUgEpLnE'
    }
    
random Integer

    package main

    import (
      "fmt"
      "github.com/illud/random"
    )

    func main() {
      fmt.Println(random.RandomNumber(0, 80)) //0 is the minimun number and 80(CAN BE ANY NUMBER) is the maximun number, change for your own minimun and maximun numbers, output example '25'
    }
