# random
Random string and int generator for Go programming

EXAMPLE

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
