package main
  
import (
  "fmt"
  "os"
  "flag"
  "bufio"
  "io"
  "encoding/binary"
)

func readHeader(reader io.Reader) {

  aliceHeader := make([]byte, 5)
  var aliceVersion uint8

  err := binary.Read(reader, binary.LittleEndian, &aliceHeader)
  err  = binary.Read(reader, binary.LittleEndian, &aliceVersion)

  if (err != nil) {
    fmt.Printf("This does not appear to be a file encoded by Alice\n")
    os.Exit(1)
  }

  aliceHeaderText := string(aliceHeader)

  if aliceHeaderText == "ALICE" && aliceVersion == 1 {
    fmt.Printf("This is an Alice file encoded by Alice version %v\n", int(aliceVersion))
  } else {
    fmt.Println("This does not appear to be a file encoded by Alice")
    os.Exit(1)
  }
}

func main() {
  fileFlagValue  := flag.String("file", "", "The file to expand.")
  outFlagValue  := flag.String("out", "", "The destination file. If one is not specified, we will decompress to stdout.")
  flag.Parse()

  if(*fileFlagValue == "") {
    fmt.Println("Please specify an input file with the --file flag")
    os.Exit(1)
  }

  targetFileName  := *fileFlagValue
  outFileName     := *outFlagValue

  fmt.Printf("Eating: %s\n", targetFileName)

  var outFile *os.File
  if outFileName != "" {
    fmt.Printf("Outputting to: %s\n", outFileName)
    outFile, _ = os.OpenFile(outFileName, os.O_RDWR|os.O_CREATE, 0644)

    defer outFile.Close()
  }

  targetFile, _ := os.OpenFile(targetFileName, os.O_RDONLY, 0644)

  reader      := bufio.NewReader(targetFile)
  readBuffer  := make([]byte, 10000)
  writeBuffer  := make([]byte, 0)

  readHeader(reader)

  bytesRead, err := reader.Read(readBuffer)

  var lastByte byte
  for {
    if (err != nil && err == io.EOF) {
      break
    }

    readBufferPosition := 0
    data := readBuffer[:bytesRead]
    for readBufferPosition < len(data) {

      bufferByte := data[readBufferPosition]
      readBufferPosition++

      if bufferByte == lastByte {
        count   := int(data[readBufferPosition])
        readBufferPosition++

        for j:=1; j<count; j++ {
          writeBuffer = append(writeBuffer, bufferByte)
        }
      } else {
        writeBuffer = append(writeBuffer, bufferByte)
      }

      lastByte = bufferByte
    }

    bytesRead, err = reader.Read(readBuffer)

    if (err != nil && err == io.EOF) {
      break
    }

    data = readBuffer[:bytesRead]
    readBufferPosition = 0
  }


  if outFileName != "" {
    fmt.Println(len(writeBuffer))
    outFile.Write(writeBuffer)

    outFileSize, _ := outFile.Stat()
    fmt.Printf("Uncompresszed file size is: %v\n", outFileSize.Size())
  } else {
    fmt.Println(len(writeBuffer))
    fmt.Printf("%v\n", string(writeBuffer))
  }

}