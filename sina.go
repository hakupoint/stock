package stock

import (
  "net/http"
  "io/ioutil"
  "strings"
)

const (
  SINA_URL = "http://hq.sinajs.cn/list="
)

//股票数据
type Stocks struct {
  Number string
  Name string
  Today float64
  Tomorrow float64
  Current float64
  TodayMax float64
  TodayMin float64
  BuyOneBidPrice float64 //买一竞买价
  SellOneBidPrice float64 //卖一竞买价
  DealNumber float64
  DealMoney float64
  BuyOneList []ThighOrMoney //买一申请多少股,需要除以100得到手
  SellOneThigh []ThighOrMoney //卖一申请多少股,需要除以100得到手
  Cal string
  Time string
}

//每手的数据
type ThighOrMoney struct{
  Thigh float64
  Money float64
}

//历史数据
type History struct {
  Date string
  open string
  High string
  Low string
  Close string
  Volume string
  AdjClose string
}

//需要处理的数据
func Read(s string) ([]*Stocks, error){
  list := strings.Split(s, ",")
  stocksList := []*Stocks{}
  for _,val := range list{
    stocks,err := getData(val)
    if err != nil{
      return stocksList, err
    }
    stocksList = append(stocksList, stocks)
  }
  return stocksList, nil
}
//抓取数据
func getData(s string)(*Stocks, error){
  var stos *Stocks
  ch := make(chan *Stocks, 0)
  res,err := http.Get(SINA_URL + s)
  if err != nil{
    return stos, err
  }

  defer res.Body.Close()

  dst, err := ioutil.ReadAll(res.Body)
  if err != nil{
    return stos, err
  }
  go dataProcessing(dst, s, ch)
  stos = <-ch
  return stos, nil
}
//数据处理
func dataProcessing(s []byte, number string, ch chan *Stocks){
  udst := unicode(string(s))
  chars := splitString([]rune(udst))
  floats := stringListToFloatList(chars[:], 1, len(chars) - 3)
  chars = append(chars[:1], chars[30:]...)
  stocks := Stocks{
      number,
      chars[0],
      floats[0],
      floats[1],
      floats[2],
      floats[3],
      floats[4],
      floats[5],
      floats[6],
      floats[7],
      floats[8],
      []ThighOrMoney{
        {floats[9],floats[10]},
        {floats[12],floats[12]},
        {floats[13],floats[14]},
        {floats[15],floats[16]},
        {floats[17],floats[18]},
      },
      []ThighOrMoney{
        {floats[19],floats[20]},
        {floats[21],floats[22]},
        {floats[23],floats[24]},
        {floats[25],floats[26]},
        {floats[27],floats[28]},
      },
    chars[1],
    chars[2],
    }
  ch <- &stocks
}


//Stocks结构体方法
func (s *Stocks) History(){
  yahoo_url := "http://table.finance.yahoo.com/table.csv?s="
  stock_number := s.Number
  if ok := strings.HasPrefix(stock_number,"sh");ok{
    new_number := strings.Replace(stock_number, "sh", "", -1)
    yahoo_url = yahoo_url + new_number + ".ss"
  }else{
    new_number := strings.Replace(stock_number, "sz", "", -1)
    yahoo_url = yahoo_url + new_number + ".sz"
  }
  res, _ := http.Get(yahoo_url)
  defer res.Body.Close();
}