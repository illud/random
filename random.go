package random

import (
  "math/rand"
  "time"
)

func RandomString(count int) string{
  var arr = []string{"a","b","c","d","e","f","g","h","i","j","l","m","n","o","p","q","r","s","t","w","x","y","z","A","B","C","D","E","F","G","H","I","J","L","M","N","O","P","Q","R","S","T","W","X","Y","Z"}
  var randomString string

  rand.Seed(time.Now().UnixNano())

  for x := 0; x < count; x++{
    randomString += arr[rand.Intn(42 - 0 + 1)]
  }

  return randomString
}

func RandomNumber(max int, min int) int{
  rand.Seed(time.Now().UnixNano())
  
  return rand.Intn(max - min + 1)
}
