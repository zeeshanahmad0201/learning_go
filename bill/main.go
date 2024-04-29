package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(reader *bufio.Reader, question string) (string, error) {
	fmt.Print(question)

	input, err := reader.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	input, _ := getInput(reader, "Please enter the bill name: ")

	fmt.Println("Created the bill - ", input)

	return newBill(input)
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	input, _ := getInput(reader, "Choose option a - add item, s - save bill, t - tip: ")

	switch input {
	case "a":
		itemName, _ := getInput(reader, "Please enter name: ")
		price, _ := getInput(reader, "Please enter price: ")

		parsedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Please enter a valid number!")
			promptOptions(b)
		}
		b.addItem(itemName, parsedPrice)
		fmt.Println("Item added - ", itemName, parsedPrice)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("bill saved as ", b.name)
	case "t":
		tip, _ := getInput(reader, "Please enter the tip($): ")
		parsedTip, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("Please enter a valid number!")
			promptOptions(b)
		}
		b.updateTip(parsedTip)
		fmt.Println("Tip added -", parsedTip)
		promptOptions(b)
	default:
		fmt.Println("Invalid option!")
		promptOptions(b)

	}
}

func main() {
	myBill := createBill()

	// fmt.Println(myBill)
	promptOptions(myBill)
}
