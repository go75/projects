package main

import loadbalance "proxy/load_balance"

func main() {
	b := loadbalance.WeightRoundBalance{}
	b.Add("127.0.0.1", "6")
	b.Add("localhost", "4")
	b.Add("0.0.0.0", "2")
	for i := 0; i < 6; i++ {
		println(b.Next())
	}
}