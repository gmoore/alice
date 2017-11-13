package main
  
import (
  "fmt"
  "os"
  "flag"
  "bufio"
  "io"
  // "strings"
  // "path/filepath"
)

func main() {
  fileFlagValue  := flag.String("file", "", "The file to expand.")
  flag.Parse()

  if(*fileFlagValue == "") {
    fmt.Println("Please specify an input file with the --file flag")
    os.Exit(1)
  }

  targetFileName := *fileFlagValue

  fmt.Printf("Eating: %s\n", targetFileName)

  targetFile, _ := os.OpenFile(targetFileName, os.O_RDONLY, 0644)

  reader      := bufio.NewReader(targetFile)
  readBuffer  := make([]byte, 10000)

  _,err := reader.Read(readBuffer)
  for {
    if (err != nil && err == io.EOF) {
      break
    }

    for i:=0; i<len(readBuffer); i++ {
      if readBuffer[i] == 0 {
        break
      }

      bufferByte := readBuffer[i]

      if bufferByte == byte('|') {
        count   := int(readBuffer[i+1])
        //fmt.Printf("%v\n", count)
        for j:=0; j<count; j++ {
          fmt.Printf("%v", string(readBuffer[i+2]))
        }
      } else {
        fmt.Printf("%v", string(readBuffer[i]))
      }
    }

    _,err = reader.Read(readBuffer)
  }
  fmt.Printf("\n")
}