// Steve Phillips / elimisteve
// 2011.06.13

package main
import ("io/ioutil"; "flag")

func main() {
    filename := flag.Arg(0)
    contents, _ := ioutil.ReadFile(filename)
    count := 0
    for _, char := range contents {
        if char == 'x' {
            count += 1
        }
    }
    println(count)
}