
package main

import "regexp"
import "fmt"

func main() {
	pattern := `([A-Za-z0-9/\.\?\-_]+)`
  str1 := "centodds/9-87/my_image.img"

  reg1,_ := regexp.Compile(pattern)
 
  fmt.Print(reg1.FindStringSubmatch(str1))


}
