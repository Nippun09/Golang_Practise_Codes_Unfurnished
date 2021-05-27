//this is the dog package for exercise of documentation
//i want to see if this second line is also shown as documentation
package dog

import "fmt"

//YearstoDogYears convert human years to dog years
func YearstoDogYears(y int) int{
  fmt.Println("hi i am inside dog package")
  return y*7
}