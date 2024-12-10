package cmd

import (
	"errors"
	"fmt"

	"github.com/robert430404/precious-metals-tracker/services"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type OutputFormat = string

const (
	Table OutputFormat = "table"
	Json  OutputFormat = "json"
)

type HoldingFlags struct {
	IsAdding     bool
	IsListing    bool
	IsDeleting   bool
	OutputFormat OutputFormat
}

func NewHoldingFlags(flags *pflag.FlagSet) (*HoldingFlags, error) {
	isAdding, err := flags.GetBool("add")
	isListing, err2 := flags.GetBool("list")
	isDeleting, err3 := flags.GetBool("delete")
	outputFormat, err4 := flags.GetString("format")

	if err != nil || err2 != nil || err3 != nil || err4 != nil {
		return nil, errors.New("could not parse flags")
	}

	hydratedFlags := &HoldingFlags{
		IsAdding:     isAdding,
		IsListing:    isListing,
		IsDeleting:   isDeleting,
		OutputFormat: outputFormat,
	}

	return hydratedFlags, nil
}

func (self *HoldingFlags) IsValid() error {
	trues := 0
	for _, value := range []bool{self.IsAdding, self.IsListing, self.IsDeleting} {
		if value {
			trues += 1
		}
	}

	if trues > 1 {
		return errors.New("invalid flags passed")
	}

	if self.OutputFormat != Json && self.OutputFormat != Table {
		return errors.New("invalid flags passed")
	}

	return nil
}

func handleHolding(cmd *cobra.Command, args []string) {
	flags, err := NewHoldingFlags(cmd.Flags())
	if err != nil || flags.IsValid() != nil {
		fmt.Println("please provide a valid signiture, run `holding --help` for more information")
		return
	}

	holdingService := services.GetHoldingService()

	if flags.IsDeleting {
		holdingService.Delete()
		return
	}

	if flags.IsListing {
		holdingService.List()
		return
	}

	holdingService.Add()
}

var holdingCmd = &cobra.Command{
	Use:   "holding [--add | -a | --delete | -d | --list | -l] [--format output | -f output]",
	Short: "Interacts with holdings in the system",
	Long:  `This command allows you to interact with all holdings in the system. This allows you to add new holdings, view your holding, and do other operations on them as you use the flags this command provides.`,
	Run:   handleHolding,
}

func init() {
	holdingCmd.Flags().BoolP("add", "a", false, "tells the command you want to add a holding")
	holdingCmd.Flags().BoolP("delete", "d", false, "tells the command you want to delete a holding")
	holdingCmd.Flags().BoolP("list", "l", false, "tells the command you want to list your holdings")
	holdingCmd.Flags().StringP("format", "f", Table, fmt.Sprintf("decides the output format, supports: [\"%v\", \"%v\"]", Json, Table))

	rootCmd.AddCommand(holdingCmd)
}
