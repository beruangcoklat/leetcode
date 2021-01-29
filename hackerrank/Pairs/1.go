package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func abs(a,b int32)int32{
    x := a - b
    if x < 0{
        x *= -1
    }
    return x
}

// Complete the pairs function below.
func pairs(k int32, arr []int32) int32 {
    dict := map[int32]struct{}{}
    used := make(map[int32]map[int32]struct{})

    arrLen := len(arr)
    ret := int32(0)

    for i := 0; i < arrLen; i++ {
        dict[arr[i]] = struct{}{}
    }

    for i := 0; i < arrLen; i++ {
        x := abs(arr[i], k)
        if _, ok := used[x][arr[i]]; ok {
            continue
        }

        if _, ok := dict[x]; ok && abs(x, arr[i]) == k {
            if used[x] == nil {
                used[x] = map[int32]struct{}{}
            }
            used[x][arr[i]] = struct{}{}
            ret++
        }
    }

    return ret
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nk := strings.Split(readLine(reader), " ")

    nTemp, err := strconv.ParseInt(nk[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    kTemp, err := strconv.ParseInt(nk[1], 10, 64)
    checkError(err)
    k := int32(kTemp)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    result := pairs(k, arr)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
