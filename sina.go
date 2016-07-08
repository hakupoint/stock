package stock

import (
  "net/http"
  "fmt"
  "io/ioutil"
  "strings"
)

const URL = "http://hq.sinajs.cn/list="

type Sina struct {
  Name string
  Today string
  Tomorrow string
  Current string
  TodayMax string
  TodayMin string
  BuyOneBidPrice string //买一竞买价
  SellOneBidPrice string //卖一竞买价
  DealNumber string
  DealMoney string
  BuyOneThigh string //买一申请多少股,需要除以100得到手
  BuyOneMoney string //买一价格
  BuyTwoThigh string
  BuyTwoMoney string
  BuyThreeThigh string
  BuyThreeMoney string
  BuyFourThigh string
  BuyFourMoney string
  BuyFiveThigh string
  BuyFiveMoney string
  SellOneThigh string //卖一申请多少股,需要除以100得到手
  SellOneMoney string //卖一价格
  SellTwoThigh string
  SellTwoMoney string
  SellThreeThigh string
  SellThreeMoney string
  SellFourThigh string
  SellFourMoney string
  SellFiveThigh string
  SellFiveMoney string
  Cal string
  Time string
}

func Read(num string) (*Sina){
  var sin Sina
  res,err := http.Get(URL + num)
  if err != nil{
    fmt.Println(err)
  }
  defer res.Body.Close()

  dst, err := ioutil.ReadAll(res.Body)
  if err != nil{
    fmt.Print(err)
  }
  udst := unicode(string(dst))
  chars := split([]rune(udst))
  sin = Sina{
    chars[0],
    chars[1],
    chars[2],
    chars[3],
    chars[4],
    chars[5],
    chars[6],
    chars[7],
    chars[8],
    chars[9],
    chars[10],
    chars[11],
    chars[12],
    chars[13],
    chars[14],
    chars[15],
    chars[16],
    chars[17],
    chars[18],
    chars[19],
    chars[20],
    chars[21],
    chars[22],
    chars[23],
    chars[24],
    chars[25],
    chars[26],
    chars[27],
    chars[28],
    chars[29],
    chars[30],
    chars[31],
  }
  return &sin
}

func split(r []rune) []string{
  r = r[21: len(r) - 3]
  return strings.Split(string(r), ",")
}

