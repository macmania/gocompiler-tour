//Temporary go file that parsers a piet image then classifies the pixels by using the
//rules specified on this site: http://www.dangermouse.net/esoteric/piet.html

//A lex parser in Go for the language Cool
package lexer //figure out how to run a go file using another package name

import (
	"fmt"
	_ "go/build"
	//"go/token"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	//"io/ioutil"
	//"log"
	_ "encoding/hex"
	"os"
	_ "reflect"
	_ "strconv"
	"strings"
)

//enums in golang, represents the dp's direction
const (
	DPRight = iota // ->
	DPDown  = iota
	DPLeft  = iota //<-
	DPUp    = iota
)

const (
	CCLeft  = iota
	CCRight = iota
)

//lists all of the colors that will be represented in this language
//props to gitpan/Piet-Interpreter, just following his
//interpreter example to get started with this compiler
var tableHexColors = map[string]string{
	"FFC0C0": "light red", "FF0000": "red",
	"FFFFC0": "light yellow", "FFFF00": "yellow",
	"C0FFC0": "light green", "00FF00": "green",
	"C0FFFF": "light cyan", "00FFFF": "cyan",
	"C0C0FF": "light blue", "0000FF": "blue",
	"FFC0FF": "light magenta", "FF00FF": "magenta",
	"C00000": "dark red", "FFFFFF": "white",
	"C0C000": "dark yellow", "000000": "black",
	"00C000": "dark green",
	"00C0C0": "dark cyan",
	"0000C0": "dark blue",
	"C000C0": "dark magenta",
}

var tableHexCycle = map[string]int{
	"FFC0C0": 0, "FF0000": 0,
	"FFFFC0": 1, "FFFF00": 1,
	"C0FFC0": 2, "00FF00": 2,
	"C0FFFF": 3, "00FFFF": 3,
	"C0C0FF": 4, "0000FF": 4,
	"FFC0FF": 5, "FF00FF": 5,
	"C00000": 0, "FFFFFF": -1,
	"C0C000": 1, "000000": -1,
	"00C000": 2,
	"00C0C0": 3,
	"0000C0": 4,
	"C000C0": 5,
}

var tableHexLight = map[string]int{
	"FFC0C0": 0, "FF0000": 1, "C00000": 2,
	"FFFFC0": 0, "FFFF00": 1, "C0C000": 2,
	"C0FFC0": 0, "00FF00": 1, "00C000": 2,
	"C0FFFF": 0, "00FFFF": 1, "00C0C0": 2,
	"C0C0FF": 0, "0000FF": 1, "0000C0": 2,
	"FFC0FF": 0, "FF00FF": 1, "C000C0": 2,
	"FFFFFF": -1, "000000": -1,
}

var colorCommands = [...]string{"push", "pop", "noop",
	"add", "subtract", "multiply",
	"divide", "mod", "not",
	"greater", "pointer", "switch",
	"duplicate", "roll", "in_char",
	"out_char", "in_num", "out_num",
}

type StackOps interface {
	push()
	pop() string
	peek() string
}

//make an interface for this
type Stack struct {
	command string
	next    *Stack
}

//contains all of the fields needed to run the interpreter
type PietInterpreter struct {
	DP        int
	CC        int
	stackProg Stack //maintains the stack operation details
}

func push(stack *Stack, command string) {
	tempComm := *stack
	stack.command = command
	stack.next = &tempComm
}

func peek(stack *Stack) string {
	return stack.command
}

func pop(stack *Stack) {
	if stack != nil {
		fmt.Println(" here", stack)
		stack.command = stack.next.command
		fmt.Println(stack.command)
		for ; stack != nil; stack = stack.next {
			fmt.Print(stack.command)
		}
		stack = stack.next
		fmt.Println(stack.command, stack.next)
	}

}

//goes to the entire file and parser each of the pixels and classifies the change of colors,
//hues and light and outputs the changes to a new text
func parser(fileName string) {
	infile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer infile.Close() //puts this function call to a list, it ensures that this call is called

	src, _, err := image.Decode(infile)
	if err != nil {
		fmt.Println(err)
	}

	//Need to be redone
	bounds := src.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r, b, g, _ := src.At(x, y).RGBA()
			colorInt := []byte{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8)}
			hexStr := strings.Replace(fmt.Sprintf("% x", colorInt), " ", "", -1)
			hexStr = strings.ToUpper(hexStr)
			fmt.Print(tableHexColors[hexStr], " ")
		}
		fmt.Print("\n")
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please enter the file name") //to-do, print based on the error specified on the top of the file
	} else {
		fmt.Println(os.Args[1])
		parser(os.Args[1])
	}
}
