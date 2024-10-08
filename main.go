package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("\n\nWelcome to your cli tool to generate mantine tables\n\n\n")
	table := Table{
		horizontalSpacing: "sm",
		verticalSpacing:   "sm",
		backgroundColor:   "",
		shadow:            "none",
	}

	size_spacing_shadow_Options := []string{"1. for small", "2. for medium", "3. for large", "4. for the largest \n"}
	bgColorOptions := []string{"1. for white", "2. for gray \n"}
	bgColorIntensityOptions := []string{"1. for light (closer to white)", "9. for intense (closer to black) \n"}
	// fontWeightOptions := []string{"1. for light", "2. for normal", "3. for medium", "4. for the for semibold", "5. for bold \n"}
	// fontColorOptions := []string{"1. for white", "2. for gray", "3. for black", "4. for blue \n"}
	stripeOptions := []string{"1. for even", "2. for odd \n"}
	stripedOptions := []string{"1. If yes", "Any other number for no \n"}
	emptyOptions := []string{"\n"}

	hSpacing := uint8(10)
	vSpacing := uint8(10)
	bgColor := uint8(10)
	bgColorIntensity := uint8(10)
	shadow := uint8(10)
	var columnsNumber int
	stripe := uint8(10)
	striped := false

	print("Let's get started...\n")

	// ! horizontal spacing
	hSpacing = handleInput("Choose the horizontal spacing", size_spacing_shadow_Options, 1, 4, false)

	// ! vertical spacing
	vSpacing = handleInput("Choose the vertical spacing", size_spacing_shadow_Options, 1, 4, false)

	// ! background color
	bgColor = handleInput("Choose the background color", bgColorOptions, 1, 2, true)

	// ! background color intensity
	bgColorIntensity = handleInput("Choose a background color intensity", bgColorIntensityOptions, 1, 9, false)

	// ! shadow
	shadow = handleInput("Choose the shadow", size_spacing_shadow_Options, 1, 4, false)

	// ! striped
	for true {
		inputPlaceholder, err := acceptIntegerInput("Would you like the rows of your table to be striped?", stripedOptions)
		if err != nil {
			println("Please enter an integer input...")
			continue
		}
		if inputPlaceholder == 1 {

			// ! stripe
			striped = true
			stripe = handleInput("Choose the stripe order", stripeOptions, 1, 2, true)
		}
		break
	}

	for true {
		inputPlaceholder, err := acceptIntegerInput("How many columns would you like to have?", emptyOptions)
		if err != nil {
			println("Please enter an integer input...")
			continue
		}
		columnsNumber = inputPlaceholder
		break
	}

	table.editTable(hSpacing, vSpacing, striped, stripe, bgColor, bgColorIntensity, shadow)

	// ! columns
	for i := 0; i < columnsNumber; i++ {
		fmt.Printf("What is your title for column #%v? \t", i+1)
		var columnTitle string
		fmt.Scanln(&columnTitle)
		column := Column{
			title: strings.TrimSpace(columnTitle),
		}

		// var fontSize uint8

		// for true {
		// 	inputPlaceholder, err := acceptIntegerInput("Choose a font size for this column", size_spacing_shadow_Options)
		// 	if err != nil {
		// 		println("Please enter an integer input...")
		// 		continue
		// 	}
		// 	min, max := 1, 4
		// 	uint8placeholder, err := validateIntegerInput(inputPlaceholder, min, max)
		// 	if err != nil {
		// 		fmt.Printf("Please enter a number between %d and %d. Both inclusive \n", min, max)
		// 		continue
		// 	}
		// 	fontSize = uint8placeholder
		// 	break
		// }

		// column.fontSize = fontSize

		// var fontWeight uint8

		// for true {
		// 	inputPlaceholder, err := acceptIntegerInput("Choose a font weight for this column", fontWeightOptions)
		// 	if err != nil {
		// 		println("Please enter an integer input...")
		// 		continue
		// 	}
		// 	min, max := 1, 5
		// 	uint8placeholder, err := validateIntegerInput(inputPlaceholder, min, max)
		// 	if err != nil {
		// 		fmt.Printf("Please enter a number between %d and %d. Both inclusive \n", min, max)
		// 		continue
		// 	}
		// 	fontWeight = uint8placeholder
		// 	break
		// }

		// column.fontWeight = fontWeight

		// var fontColor uint8

		// for true {
		// 	inputPlaceholder, err := acceptIntegerInput("Choose a font weight for this column", fontColorOptions)
		// 	if err != nil {
		// 		println("Please enter an integer input...")
		// 		continue
		// 	}
		// 	min, max := 1, 5
		// 	uint8placeholder, err := validateIntegerInput(inputPlaceholder, min, max)
		// 	if err != nil {
		// 		fmt.Printf("Please enter a number between %d and %d. Both inclusive \n", min, max)
		// 		continue
		// 	}
		// 	fontColor = uint8placeholder
		// 	break
		// }

		// column.fontColor = fontColor

		table.mountColumn(column)

	}

	table.buildTable()

}
