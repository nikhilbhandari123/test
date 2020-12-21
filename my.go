package main

import (
    "fmt"
    "strings"
)

const max = "1234567890123456"
func main() {
   arr := [] string{"\"Engineering (ACME)\"", "\"QA (ACME)\""}
   var parts, list2 []string
   for _, x := range arr {
        parts = strings.Split(x, " ")
        list2 = append(list2, parts[0])  // note the = instead of :=
   }

   name := "sixteen123456789"
   fmt.Printf("%s", name[:16])
   if i := len(name); i < 16 {
     name += max[i:]
   }
}
