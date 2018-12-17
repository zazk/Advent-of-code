package main
import (
	"fmt"
	"strings"
	"io/ioutil"
	"strconv"
)

func check(err error){
	if (err != nil){
		panic(err)
	}
}
func main (){
	data, err := ioutil.ReadFile("input.txt")
	check(err)
	lines := strings.Split(string(data), "\n")
	result := 0
	for _,strOrder := range lines {
		order, err := strconv.Atoi(strOrder)
		check(err)
		result = result + order
	}
	fmt.Print(result)
}