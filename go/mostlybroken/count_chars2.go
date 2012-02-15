package main; import ("io/ioutil"; "flag")

func main() {
    const NUMCORES = 6
    bytes, _ := ioutil.ReadFile(flag.Arg(0))
    size := len(bytes) / NUMCORES  // chunk size
    str := string(bytes)

    chunks := []string{}
    var i int
    for i = 0; i < NUMCORES*size; i += size {
        chunks = append(chunks, str[i:i+size])
    }
    // Don't forget the (possible) partial chunk at the end! Stems
    // from integer division to get `size`; rounds down
    chunks = append(chunks, str[i:])

    intChan := make(chan int)

    // Launch a goroutine (cheap thread) for each paragraph
    for _, chunk := range chunks {
        go func(str string) {
            count := 0
            for _, char := range str {
                if char == 'x' { count += 1 }
            }
            intChan <- count
        }(chunk)
    }
    charCount := 0
    for i := 0; i < len(chunks); i++ {
        charCount += <-intChan
    }
    println(charCount)
}