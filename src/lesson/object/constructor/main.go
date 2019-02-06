package main

import (
	"fmt"
	"log"
)

// Money型を定義
type Money struct {
	amount   uint
	currency string
}

/* constructor: ポインタ型Moneyを定義するfunc()
1行目：第一引数にuint型をとるNewMoney関数(コンストラクタ)を生成、*Moneyを返す
2行目：第一フィールドに引数のamount,第二フィールドに"yen"をとる&Moneyを返す
*/
func NewMoney(amount uint) *Money {
	return &Money{amount, "yen"}
}

// ポインタ型Moneyだけが使えるメソッド。レシーバーのamountとcurrencyを出力
func (this *Money) Format() string {
	return fmt.Sprintf("%d %s", this.amount, this.currency)
}

// Add()メソッド。*Money型のレシーバ(this)に、第一引数である*Money型のamountを足し合わせる。
// 返り値は渡さない。
func (this *Money) Add(that *Money) {
	this.amount += that.amount
}

func main() {
	// コンストラクタにて*Moneyを生成。NewMoneyの返り値&Moneyは「アドレス演算子」
	// moneyに*Money型を、フォーマットにて内容を出力
	money := NewMoney(120)
	fmt.Println(&money)
	log.Print(money.Format())

	money.Add(NewMoney(180))
	log.Print(money.Format())
	log.Println(money)

	money.amount += 180
	log.Println(money)
}
