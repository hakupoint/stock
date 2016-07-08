package stock

import "github.com/axgle/mahonia"

func unicode(s string) string{
  enc := mahonia.NewDecoder("gbk")
  return enc.ConvertString(s)
}
