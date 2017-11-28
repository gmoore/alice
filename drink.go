package main
  
import (
  "fmt"
  "os"
  "flag"
  "bufio"
  "io"
  "strings"
  "path/filepath"
)

func writeHeader(writeBuffer []byte) []byte {
  //Our name
  writeBuffer = append(writeBuffer, []byte("ALICE")...)

  //The version. We're hosed if we go past 255
  writeBuffer = append(writeBuffer, byte(1))

  return writeBuffer
}

func write(writeBuffer []byte, thing byte, counter int) []byte {
  if counter == 1 {
    writeBuffer = append(writeBuffer, thing)
  } else {
    writeBuffer = append(writeBuffer, thing)
    writeBuffer = append(writeBuffer, thing)
    writeBuffer = append(writeBuffer, byte(counter))
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

  writeBuffer = writeHeader(writeBuffer)

  bytesRead, err := reader.Read(readBuffer)
  for {
    if (err != nil && err == io.EOF) {
      break
    }

    var lastByte byte;
    counter := 1
    data := readBuffer[:bytesRead]
    for idx, bufferByte := range data {
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

    writeBuffer = write(writeBuffer, lastByte, counter)

    bytesRead,err = reader.Read(readBuffer)
  }

  outFile.Write(writeBuffer)

  outFileSize, _ := outFile.Stat()
  fmt.Printf("Compresszed file size is: %v\n", outFileSize.Size())
}