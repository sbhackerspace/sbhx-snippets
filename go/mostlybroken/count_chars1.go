// Steve Phillips / elimisteve
// 2011.06.13

package main
import "io/ioutil"

func main() {
    contents, _ := ioutil.ReadFile("temp.html")
    count := 0
    for _, char := range contents {
        if char == 'x' {
            count += 1
        }
    }
    println(count)
}