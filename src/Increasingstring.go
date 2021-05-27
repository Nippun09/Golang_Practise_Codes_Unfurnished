package main

import (
	"fmt"
	"strings"
)

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for i := 1; i <= t; i++ {

		var n int
		fmt.Scanf("%d", &n)

		var str string
		fmt.Scanf("%s",&str)

		reslice := operation(n, str)

		fmt.Print("Case #", i, ": ")
		var restring []string
		for _,v:=range reslice{
			restring=append(restring, fmt.Sprint(v))
		}
		fmt.Print(strings.Join(restring," "))
		fmt.Println("")
		
	}
}

func operation(length int, str string) []int {
strslice:=strings.Split(str, "")

var slint []int

for i:=0;i<length-1;i++{
	slint=append(slint, strings.Compare(strslice[i], strslice[i+1]))
}
slint=append(slint, 1)
count:=1
for j:=0;j<length; j++{
	if slint[j]==-1{
		slint[j]=count
		count+=1
	}else{
		slint[j]=count
		count=1
	}
}

return slint
}