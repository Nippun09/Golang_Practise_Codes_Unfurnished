package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for i := 1; i <= t; i++ {

		var n string
		var f string
		fmt.Scanf("%s %s", &n, &f)

		finalstring := operation2(n,f)
		fmt.Print("Case #", i, ": ")
		
		fmt.Print(finalstring)
		fmt.Println("")
		
	}
}

func operation2(phn string, format string) string {
strslice:=strings.Split(format,"-")
var intslice []int
intslice=append(intslice, 0)
for _,v:= range strslice{

	vint,_:=strconv.Atoi(v)
	intslice=append(intslice, vint)
}

phncharslice:=strings.Split(phn, "")

var finalstring []string
startindex:=intslice[0]

for j:=0; j<len(intslice)-2;j++{

	endindex:=startindex+intslice[j+1]-1
	fmt.Println("start index: ", startindex, " endindex: ",endindex+1)
	finalstring=append(finalstring, repeatreplace(phncharslice[startindex:endindex+1]))
	startindex=endindex+1
}
finalstring=append(finalstring, repeatreplace(phncharslice[startindex:]))
return strings.Join(finalstring, " ")
}

func repeatreplace(opslice []string) string {
	fmt.Println("slice received: ", opslice)
	fmt.Println("length of slice: ", len(opslice))
	var newslice []string
	tuple:=make(map[string]int)
	achar:=opslice[0]
	count:=1
	for i:=0; i<len(opslice)-1;i++ {
		if achar==opslice[i+1] {
			count,ok:=tuple[achar]
			if ok{
				count+=1
				tuple[achar]=count
			}else{
				count=2
				tuple[achar]=count
			}
			
		}else{
			count=1
			achar=opslice[i+1]
			tuple[achar]=count
		}
	}
	fmt.Println("last value: ",opslice[len(opslice)-1])
	_,ok:=tuple[opslice[len(opslice)-1]]
	if !ok{
		fmt.Println("i am coming inside to set value of last value")
		tuple[opslice[len(opslice)-1]]=1
	}
var timesword string
var numword string
	for i:=0;i<len(opslice);i++ {
		if tuple[opslice[i]]>1{
			timesword=repeatword(tuple[opslice[i]])
			numword=repeatnumber(opslice[i])
			tempstr:=fmt.Sprint(timesword," ", numword)
			newslice=append(newslice, tempstr)
			counter,_:=tuple[opslice[i]]
			i=i+counter-1
		}else{
			fmt.Println("value 1: ",opslice[i])
			numword=repeatnumber(opslice[i])
			newslice=append(newslice, numword)
		}
	}
return strings.Join(newslice, " ")	

}

func repeatword(num int) string{
	switch num {
	case 2: return "double"
	case 3: return "triple"
	case 4: return "quadruple"
	case 5: return "quintuple"
	case 6: return "sextuple"
	case 7: return "septuple"
	case 8: return "octuple"
	case 9: return "nonuple"
	case 10: return "decuple"
	default: return ""
	}
	return""
}

func repeatnumber(numstr string) string{
	switch numstr {
	case "0": return "zero"
	case "1": return "one"
	case "2": return "two"
	case "3": return "three"
	case "4": return "four"
	case "5": return "five"
	case "6": return "six"
	case "7": return "seven"
	case "8": return "eight"
	case "9": return "nine"
	default: return ""
	}
	return ""
}