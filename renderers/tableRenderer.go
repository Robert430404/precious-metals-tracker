package renderers

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/robert430404/precious-metals-tracker/db/entities"
)

const RightBottomCorner = "\U00002518" // ┘
const RightBreak = "\U00002524"        // ┤
const RightTopCorner = "\U00002510"    // ┐

const LeftBottomCorner = "\U00002514" // └
const LeftBreak = "\U0000251C"        // ├
const LeftTopCorner = "\U0000250C"    // ┌

const Cross = "\U0000253C" // ┼

const HorizontalLine = "\U00002500" // ─
const VerticalLine = "\U00002502"   // │

const BottomBreak = "\U00002534" // ┴
const TopBreak = "\U0000252C"    // ┬

type TableRenderer struct{}

func (self *TableRenderer) renderTable(headers []string, data [][]string) {
	colLengths := []int{}

	for _, entry := range data {
		for index, composedString := range entry {
			valueLen := utf8.RuneCountInString(composedString)

			if len(colLengths) < len(entry) {
				colLengths = append(colLengths, valueLen)
			}

			if valueLen > colLengths[index] {
				colLengths[index] = valueLen
			}
		}
	}

	for index, header := range headers {
		headerLen := utf8.RuneCountInString(header)
		if headerLen > colLengths[index] {
			colLengths[index] = headerLen
		}

		headers[index] = fmt.Sprintf("%-*s", colLengths[index], header)
	}

	joinedHeaders := strings.Join(headers, " "+VerticalLine+" ")
	headerLine := VerticalLine + " " + joinedHeaders + " " + VerticalLine

	// We are subtracting 2 because rules are typically braced with a rune on the left and right
	horizontalRule := strings.Repeat(HorizontalLine, utf8.RuneCountInString(headerLine)-2)

	topLine := LeftTopCorner + horizontalRule + RightTopCorner
	bottomLine := LeftBreak + horizontalRule + RightBreak

	fmt.Print(topLine + "\n")
	fmt.Print(headerLine + "\n")
	fmt.Print(bottomLine + "\n")

	for index, entry := range data {
		for index, composedString := range entry {
			entry[index] = fmt.Sprintf("%-*s", colLengths[index], composedString)
		}

		if index > 0 {
			fmt.Print(LeftBreak + horizontalRule + RightBreak + "\n")
		}

		fmt.Print(VerticalLine, " ", strings.Join(entry, " "+VerticalLine+" "), " ", VerticalLine, "\n")
	}

	fmt.Print(LeftBottomCorner + horizontalRule + RightBottomCorner + "\n")
}

func (self *TableRenderer) RenderHoldingList(holdings []entities.Holding) {
	headers := []string{"ID", "Name", "Purchase Spot Price", "Total Units", "Unit Weight", "Type"}
	data := [][]string{}

	for _, holding := range holdings {
		data = append(data, []string{
			fmt.Sprintf("%v", holding.ID),
			holding.Name,
			holding.PurchaseSpotPrice,
			holding.TotalUnits,
			holding.UnitWeight,
			holding.Type,
		})
	}

	self.renderTable(headers, data)
}

func (self *TableRenderer) RenderValueTable(value string, spotPrice string) {
	headers := []string{"Current Value", "Current Spot Price"}
	data := [][]string{{value, spotPrice}}

	self.renderTable(headers, data)
}
