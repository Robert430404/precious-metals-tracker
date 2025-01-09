package renderers

import (
	"fmt"
	"regexp"
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

const BoldOpen = "\033[1m"
const BoldClose = "\033[0m"

const RedOpen = "\x1b[1;31m"
const RedClose = "\x1b[1;0m"

const GreenOpen = "\x1b[1;92m"
const GreenClose = "\x1b[1;0m"

const BlackOpen = "\x1b[1;90m"
const BlackClose = "\x1b[1;0m"

const VerticalDivider = BlackOpen + VerticalLine + BlackClose
const HorizontalDivider = BlackOpen + HorizontalLine + BlackClose
const LeftJoiner = BlackOpen + LeftBreak + BlackClose
const RightJoiner = BlackOpen + RightBreak + BlackClose
const LeftTopJoiner = BlackOpen + LeftTopCorner + BlackClose
const RightTopJoiner = BlackOpen + RightTopCorner + BlackClose
const LeftBottomJoiner = BlackOpen + LeftBottomCorner + BlackClose
const RightBottomJoiner = BlackOpen + RightBottomCorner + BlackClose

type TableRenderer struct{}

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

func (self *TableRenderer) RenderValueTable(value string, spotPrice string, totalWeight string) {
	headers := []string{"Current Value", "Current Spot Price", "Total Holding Weight"}
	data := [][]string{{value, spotPrice, totalWeight}}

	self.renderTable(headers, data)
}

func (self *TableRenderer) renderTable(headers []string, data [][]string) {
	colLengths := []int{}

	for _, entry := range data {
		for index, composedString := range entry {
			valueLen := self.getCharacterCount(composedString)

			if len(colLengths) < len(entry) {
				colLengths = append(colLengths, valueLen)
			}

			if valueLen > colLengths[index] {
				colLengths[index] = valueLen
			}
		}
	}

	for index, header := range headers {
		headerLen := self.getCharacterCount(header)
		if headerLen > colLengths[index] {
			colLengths[index] = headerLen
		}

		paddedHeader := fmt.Sprintf("%-*s", colLengths[index], header)
		boldedHeader := fmt.Sprintf(BoldOpen+"%v"+BoldClose, paddedHeader)
		headers[index] = fmt.Sprintf(GreenOpen+"%v"+GreenClose, boldedHeader)
	}

	joinedHeaders := strings.Join(headers, " "+VerticalDivider+" ")
	headerLine := VerticalDivider + " " + joinedHeaders + " " + VerticalDivider

	horizontalRule := strings.Repeat(HorizontalDivider, self.getCharacterCount(headerLine)-2)

	topLine := LeftTopJoiner + horizontalRule + RightTopJoiner
	bottomLine := LeftJoiner + horizontalRule + RightJoiner

	fmt.Print(topLine + "\n")
	fmt.Print(headerLine + "\n")
	fmt.Print(bottomLine + "\n")

	for index, entry := range data {
		for index, composedString := range entry {
			entry[index] = fmt.Sprintf("%-*s", colLengths[index], composedString)
		}

		if index > 0 {
			fmt.Print(LeftJoiner + horizontalRule + RightJoiner + "\n")
		}

		fmt.Print(VerticalDivider, " ", strings.Join(entry, " "+VerticalDivider+" "), " ", VerticalDivider, "\n")
	}

	fmt.Print(LeftBottomJoiner + horizontalRule + RightBottomJoiner + "\n")
}

func (self *TableRenderer) getCharacterCount(input string) int {
	prepared := self.stripAnsiCodes(input)

	return utf8.RuneCountInString(prepared)
}

func (self *TableRenderer) stripAnsiCodes(input string) string {
	const ansiCharacters = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"

	var expression = regexp.MustCompile(ansiCharacters)

	return expression.ReplaceAllString(input, "")
}
