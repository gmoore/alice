package main
  
import (
  "fmt"
  "os"
  "flag"
  "bufio"
  "io"
  "strings"
  "io/ioutil"
  "path/filepath"
)

func write(writeBuffer []byte, thing byte, counter int) []byte {
  fmt.Printf("Writing %v: %v\n", counter, string(thing))
  if counter == 1 {
    fmt.Printf("Writing one\n")
    writeBuffer = append(writeBuffer, thing)
  } else {
    fmt.Printf("Writing many\n")
    writeBuffer = append(writeBuffer, byte('|'))
    writeBuffer = append(writeBuffer, byte(counter))
    writeBuffer = append(writeBuffer, thing)
  }

  return writeBuffer
}

func main() {
  fileFlagValue  := flag.String("file", "", "The file to shrink.")
  flag.Parse()

  if(*fileFlagValue == "") {
    fmt.Println("Please specify an input file with the --file flag")
    os.Exit(1)
  }

  targetFileName := *fileFlagValue

  fmt.Printf("Drinking: %s\n", targetFileName)

  targetFile, _ := os.OpenFile(targetFileName, os.O_RDONLY, 0644)
  targetFileInfo, _ := targetFile.Stat()
  fmt.Printf("Target file size is %v bytes\n", targetFileInfo.Size())
  defer targetFile.Close()

  outFileName := strings.TrimSuffix(targetFileName, filepath.Ext(targetFileName))
  outFileName = outFileName + ".alice"
  outFile, err := os.OpenFile(outFileName, os.O_RDWR|os.O_CREATE, 0644)
  defer outFile.Close()

  reader      := bufio.NewReader(targetFile)
  readBuffer  := make([]byte, 10000)

  writeBuffer := make([]byte, 0)


  _,err = reader.Read(readBuffer)
  for {
    if (err != nil && err == io.EOF) {
      break
    }

    var lastByte byte;
    counter := 1
    for idx, bufferByte := range readBuffer {
      if bufferByte == 0 {
        writeBuffer = write(writeBuffer, lastByte, counter)
        break;
      }

      fmt.Printf("Read %v\n", bufferByte)

      if idx == 0 {
        lastByte = bufferByte
      } else {  
        if bufferByte == lastByte {
          counter += 1
        } else {
          writeBuffer = write(writeBuffer, lastByte, counter)
          lastByte = bufferByte
          counter = 1
        }
      }

    }

    _,err = reader.Read(readBuffer)
  }

  outFile.Write(writeBuffer)

  dat, _ := ioutil.ReadFile("sample.alice")
  fmt.Print(dat)

}