package main

import (
	"encoding/json"
	"fmt"

	"github.com/leekchan/accounting"
)

var (
	Nominals = []int{
		100000,
		50000,
		20000,
		10000,
		5000,
		2000,
		1000,
		500,
		200,
		100,
	}
)

func IDRCurrencyFractions(num int) string {
	mp := make(map[string]int)

	ac := accounting.Accounting{Symbol: "Rp. ", Precision: 0, Thousand: ".", Decimal: ","}
	total := num

	for _, Nominal := range Nominals {
		count := 0
		for total >= Nominal {
			count += 1
			total -= Nominal
		}
		mp[ac.FormatMoney(Nominal)] = count
	}

	if total >= 1 && total >= 99 {
		mp[ac.FormatMoney(Nominals[9])]++
	}

	for k, v := range mp {
		if v == 0 {
			delete(mp, k)
		}
	}

	bs, _ := json.Marshal(mp)

	return string(bs)
}

func main() {
	fmt.Println("INPUT:", 145000)
	fmt.Println("OUTPUT", IDRCurrencyFractions(145000))

	fmt.Println("")

	fmt.Println("INPUT:", 2050)
	fmt.Println("OUTPUT", IDRCurrencyFractions(2050))
}
