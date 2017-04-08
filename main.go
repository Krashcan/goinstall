package main

import(
	"log"
	"os"
	"os/exec"
	"fmt"
	"strings"
)

func main() {
	pwd,err := os.Getwd()
	if err!=nil{
		log.Printf("Unable to get working directory %s\n",err)
	}
	pwd = ReplaceSlash(pwd)
	gopath := ReplaceSlash(os.Getenv("GOPATH"))

	// +5 because /src/ is not included in gopath
	fmt.Print("go install ",pwd[len(gopath)+5:]) 
	
	out,err := exec.Command("go","install",pwd[len(gopath)+5:]).CombinedOutput()
	if err!=nil{
		log.Print(err)
	}
	fmt.Println(string(out))
}

func ReplaceSlash(x string) string{
	return strings.Replace(x,"\\","/",-1)
}
