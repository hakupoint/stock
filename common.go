package stock

import (
  "github.com/axgle/mahonia"
  "strconv"
  "fmt"
  "strings"
  "bufio"
  "io"
)

func unicode(s string) string{
  enc := mahonia.NewDecoder("gbk")
  return enc.ConvertString(s)
}

func stringListToFloatList(s []string, start int, end int) []float64{
  tmp := []float64{}
  for _, v := range s[start : end]{
    dst, err := strconv.ParseFloat(v, 64)
    if err != nil{
      fmt.Print(err)
    }
    tmp = append(tmp, dst)
  }
  return tmp
}

func splitString(r []rune) []string{
  r = r[21: len(r) - 3]
  return strings.Split(string(r), ",")
}

func floatToString(f float64) string{
  return strconv.FormatFloat(f, 'f', 6, 64)
}

func csvToString(s io.Reader) []string{
  chars := []string{}
  reader := bufio.NewReader(s)
  for{
    char, err := reader.ReadBytes('\n')
    if err == io.EOF{
      break
    }
    chars = append(chars,string(char))
  }
  return chars[1:]
}

func parseCsv(s []string){

}