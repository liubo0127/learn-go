package main

import (
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"os"
	"strconv"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

type excel struct {
	path string
	file *excelize.File
	state int // 1: new file; 0: exist file
}

func (e *excel) open() error {
	_, err := os.Stat(e.path)
	if err == nil {
		var err2 error
		e.file, err2 = excelize.OpenFile(e.path)
		if err2 != nil {
			return err2
		}
		e.state = 0
	} else {
		if os.IsNotExist(err) {
			e.file = excelize.NewFile()
			e.state = 1
		} else {
			return err
		}
	}
	return nil
}

func (e *excel) GetSheetList() []string {
	return e.file.GetSheetList()
}

func (e *excel) SheetNotExist(name string) bool {
	for _, sheet := range e.GetSheetList() {
		if sheet == name {
			return false
		}
	}
	return true
}

func (e *excel) AddSheet(name string) {
	if e.SheetNotExist(name) {
		e.file.NewSheet(name)
	} else {
		fmt.Printf("Sheet: %v already exist\n", name)
	}
}

func (e *excel) Write(sheet string, data [][]string) error {
	FetColPrefix := func(idx int) string {
		if idx > 702 {
			panic(errors.New("the sum of columns can't more than 702"))
		}

		var r string
		if idx <= 25 {
			r = string(65 + idx)
		} else {
			a := idx / 26
			b := idx % 26
			r = string(65 + a - 1) + string(65 + b)
		}

		return r
	}

	if e.SheetNotExist(sheet) {
		fmt.Printf("Sheet: %v not exist, add it...\n", sheet)
		e.AddSheet(sheet)
	}

	var err error

	for i, v := range data {
		colNum := i + 2 // start write from second line
		for ii, vv := range v {

			err = e.file.SetCellValue(sheet, FetColPrefix(ii)+strconv.Itoa(colNum), vv)
		}
	}
	return err
}

func (e *excel) Close() error {
	var err error
	switch e.state{
	case 0:
		err = e.file.Save()
	case 1:
		err = e.file.SaveAs(e.path)
	}
	return err
}

func main() {
	data := [][]string{
		{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "aa"},
		{"b", "c", "d", "e"},
	}

	addr := "./1.xlsx"

	xls := excel{
		path:  addr,
		file:  nil,
		state: 0,
	}
	if err := xls.open(); err != nil {
		panic(err)
	}
	defer xls.Close()

	xls.AddSheet("test")
	if err := xls.Write("test", data); err != nil {
		panic(err)
	}

	fmt.Println(xls.GetSheetList())
}
