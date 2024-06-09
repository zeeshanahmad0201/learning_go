package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	return bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
}

func (b bill) format() string {
	fs := "Bill breakdown\n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...%v\n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%-25v ...%.2f\n", "tip:", b.tip)

	fs += fmt.Sprintf("%-25v ...%.2f", "total:", total+b.tip)

	return fs
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

func (b *bill) save() {
	data := []byte(b.format())

	os.WriteFile("bills/"+b.name+".txt", data, 0644)
}
