package main
import(
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("1.1  CLI Args with command")
	fmt.Println(strings.Join(os.Args, " "))
	fmt.Println("1.2  CLI Args with their index value")
	for i, arg := range os.Args[1:]{
		fmt.Println(i,"     ",arg)
	}
	fmt.Println("1.3  time comaprison")
	start := time.Now()
	s, sep := "", ""
	for _,arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	elapsed := time.Since(start).Seconds()
	fmt.Println("Time in loop   ", elapsed)
	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	elapsed = time.Since(start).Seconds()
	fmt.Println("Time in join   ", elapsed)
}