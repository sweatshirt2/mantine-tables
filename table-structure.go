package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Column struct {
	title string
}

type Table struct {
	horizontalSpacing string
	verticalSpacing   string
	striped           bool
	stripe            string
	backgroundColor   string
	shadow            string
	columns           []Column
}

func newTable() Table {
	return Table{}
}

func newColumn() Column {
	return Column{}
}

func (table *Table) editTable(hSpacing uint8, vSpacing uint8, striped bool, stripe uint8, bgColor uint8, bgColorIntensity uint8, shadow uint8) {
	if striped {
		table.striped = striped
		if stripe == 1 {
			table.stripe = "even"
		} else {
			table.stripe = "odd"
		}
	}

	if bgColor != 10 {
		backgroundColor := "white"
		if bgColor == 2 {
			backgroundColor = "gray"
		}
		table.backgroundColor = "bg-" + backgroundColor

		if bgColor == 2 {
			table.backgroundColor += "-" + strconv.Itoa(int(bgColorIntensity)) + "00"
		}
	}

	// horizontal spacing
	switch hSpacing {
	case 1:
		table.horizontalSpacing = "sm"
	case 2:
		table.horizontalSpacing = "md"
	case 3:
		table.horizontalSpacing = "lg"
	case 4:
		table.horizontalSpacing = "xl"
	}

	// vertical spacing
	switch vSpacing {
	case 1:
		table.verticalSpacing = "sm"
	case 2:
		table.verticalSpacing = "md"
	case 3:
		table.verticalSpacing = "lg"
	case 4:
		table.verticalSpacing = "xl"
	}

	// shadow
	switch shadow {
	case 1:
		table.shadow = "sm"
	case 2:
		table.shadow = "md"
	case 3:
		table.shadow = "lg"
	case 4:
		table.shadow = "xl"
	}
}

// func (column *Column) editColumn(fSize uint8, fWeight uint8, fColor uint8) {
// 	column.fontColor = fColor
// 	column.fontSize = fSize
// 	column.fontWeight = fWeight
// }

func (table *Table) mountColumn(column Column) {
	table.columns = append(table.columns, column)
}

func (table *Table) buildTable() {
	tableFormat := fmt.Sprintf("<Table horizontalSpacing={%q} verticalSpacing={%q} ", table.horizontalSpacing, table.verticalSpacing)
	conditionals := ""
	if table.striped {
		conditionals += fmt.Sprintf("striped stripe={%q} ", table.stripe)
	}
	conditionals += fmt.Sprintf("class=%q ", table.backgroundColor)
	tableFormat += conditionals + "> \n"
	tableFormat += ""

	// Add the Thead
	tableFormat += fmt.Sprintf("<Thead> \n <Th> \n")
	// Todo -> Th styling
	for _, column := range table.columns {
		tableFormat += fmt.Sprintf("<Td> %v </Td> \n", column.title)
	}
	tableFormat += "</Th> \n </Thead> \n <Tbody> \n"

	tableFormat += "// add your looping method of your favorite frontend framework \n"

	for _, column := range table.columns {
		tableFormat += "<Td>{ youritem." + column.title + " }</Td> \n"
	}

	tableFormat += "</Tbody> \n </Table>"

	var tableName string
	print("\n A name for your table: ")

	for true {
		fmt.Scanln(&tableName)
		strings.TrimSpace(tableName)

		rs := saveTableStructure(tableFormat, "tables/"+tableName)
		if rs == 2 {
			println("A file with the same name exists. Please choose another name.")
		}
		fmt.Printf(tableFormat)
		println("Saved!")
		break
	}
}

func saveTableStructure(tableStructure string, name string) uint8 {
	_, err := os.Stat(name)
	if err == nil {
		return 2
	}
	// if !os.IsNotExist(err) {
	// 	return 2
	// }
	file, err := os.Create(fmt.Sprintf("%v.txt", name))
	if err != nil {
		panic("Error creating file.")
	}
	defer file.Close()
	_, err = file.WriteString(tableStructure)
	if err != nil {
		panic("Error saving text to the file")
	}
	return 1
}

func acceptIntegerInput(promptMessage string, options []string) (int, error) {
	println(promptMessage)
	for _, option := range options {
		print(option, "\t")
	}
	var input string
	fmt.Scanln(&input)
	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func validateIntegerInput(num int, min int, max int) (uint8, error) {

	if num < min || num > max {
		return 0, errors.New(fmt.Sprintf("Please enter an integer between %v and %v", min, max))
	}
	return uint8(num), nil
}
