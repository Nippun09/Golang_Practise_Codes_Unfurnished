package main

//import "fmt"

type Custmap struct{
  Key string
  Value int
}

type Newstruct struct{
  Sliceofmap []Custmap
}

func CreateMap() Newstruct {
Newsliceofmap:=make([]Custmap, 0)
  return Newstruct{Newsliceofmap}
}


func (c Newstruct) AddNewEntry(key string, value int) Newstruct{

var Mapvar []Custmap
  Mapvar=append(Mapvar, c.Sliceofmap...)
  Mapvar=append(Mapvar, Custmap{key, value})
  return Newstruct{Mapvar}
}

// func (c Newstruct) Test(key string) {
//   fmt.Printf("%T", c.Sliceofmap)
// }

func (c Newstruct) GetValue(key string) (int, bool) {
  loopslice:=c.Sliceofmap //[]Custmap
  for _,singlemap:= range loopslice{
    if singlemap.Key==key{
      return singlemap.Value, true
    }
  }
  return 0,false

}

func (c Newstruct) RemoveKey(key string) Newstruct{
loopslice:=c.Sliceofmap //[]Custmap
var Mapvar []Custmap
var removeindex int
  for index,singlemap:= range loopslice{
    if singlemap.Key==key{
      removeindex=index
    }
  }
 Mapvar=append(Mapvar, loopslice[0:removeindex]...)
 Mapvar=append(Mapvar, loopslice[removeindex+1:]...)
 return Newstruct{Mapvar}

}