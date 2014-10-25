package main

import "fmt"
import "bufio"
import "strings"
import "os"
import "bytes"

import "github.com/hermanschaaf/cedict"
import "github.com/tealeg/xlsx"

type Entry struct {
	Zw  string
	Py  string
	Def string
}

type Stats struct {
	Exist int
	New   int
}

func formatStickyStudy(e *Entry, existing string, stats *Stats) string {
	if strings.HasPrefix(existing, e.Zw) {
		stats.Exist = stats.Exist + 1
		return fmt.Sprintf("%s\n", existing)
	}
	withTones := cedict.ToPinyinTonemarks(strings.TrimSpace(e.Py))
	stats.New = stats.New + 1
	return fmt.Sprintf("%s\t\t%s\t%s\t%s\n", e.Zw, e.Py, withTones, e.Def)
}

func processCell(row *xlsx.Row) (*Entry, bool) {
	//fmt.Printf("%s", row.Cells[0].Value)
	if len(row.Cells) == 3 {
		return &Entry{row.Cells[0].Value, row.Cells[1].Value, row.Cells[2].Value}, true
	}
	return &Entry{}, false
}

func processSheet(sheet *xlsx.Sheet, scanner *bufio.Scanner) error {

	var out bytes.Buffer
	var stats Stats
	for rowNum, row := range sheet.Rows {
		if rowNum == 0 {
			continue
		}
		scanner.Scan()
		entry, cont := processCell(row)
		if cont == true {
			s := formatStickyStudy(entry, scanner.Text(), &stats)

			out.WriteString(s)
		}
	}
	f, err := os.Create(targetFileName(sheet.Name))
	if err != nil {
		panic("Could not open file to write final values")
	}
	defer f.Close()
	_, err = f.WriteString(out.String())
	if err != nil {
		panic("Error while writing output to file.")
	}
	fmt.Printf("Processed sheet: %s\nExisting entries: %d\nNew entries: %d\n\n", sheet.Name, stats.Exist, stats.New)

	return fmt.Errorf("end of sheet")
}

func targetFileName(sheetName string) string {
	home := os.Getenv("HOME")
	return home + "/Dropbox/Apps/StickyStudyChinese/" + sheetName + ".txt"
}

func openExistingStickyRecord(sheetName string) *bufio.Scanner {
	fName := targetFileName(sheetName)
	f, err := os.Open(fName)
	if err != nil {
		fmt.Errorf("%s", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for i := 0; i < 2; i++ {
		scanner.Scan()
		//fmt.Println(scanner.Text())

	}
	return scanner
}

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Error: No file to process.\nUsage: stickstudy <file>\n")
		os.Exit(1)
	}
	xlsxFileName := os.Args[1]
	xlFile, err := xlsx.OpenFile(xlsxFileName)
	if err != nil {
		panic("Input file not found")
	}

	for _, sheet := range xlFile.Sheets {
		scanner := openExistingStickyRecord(sheet.Name)
		err := processSheet(sheet, scanner)
		if err != nil {
			continue
		}

	}

}
