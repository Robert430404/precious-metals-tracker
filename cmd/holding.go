package cmd

import (
	"errors"
	"fmt"

	"github.com/robert430404/precious-metals-tracker/services"
	"github.com/robert430404/precious-metals-tracker/types"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type OutputFormat = string

const (
	Table OutputFormat = "table"
	Json  OutputFormat = "json"
)

func newHoldingFlags(flags *pflag.FlagSet) (*types.HoldingFlags, error) {
	isAdding, err := flags.GetBool("add")
	isListing, err2 := flags.GetBool("list")
	isDeleting, err3 := flags.GetBool("delete")
	isValue, err4 := flags.GetBool("value")
	outputFormat, err5 := flags.GetString("format")

	// Get new CLI args for adding holdings
	name, err6 := flags.GetString("name")
	price, err7 := flags.GetString("price")
	source, err8 := flags.GetString("source")
	spotPrice, err9 := flags.GetString("spot-price")
	units, err10 := flags.GetString("units")
	weight, err11 := flags.GetString("weight")
	holdingType, err12 := flags.GetString("type")

	// Get CLI args for deleting holdings
	deleteID, err13 := flags.GetString("id")

	if err != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil ||
		err6 != nil || err7 != nil || err8 != nil || err9 != nil || err10 != nil ||
		err11 != nil || err12 != nil || err13 != nil {
		return nil, errors.New("could not parse flags")
	}

	hydratedFlags := &types.HoldingFlags{
		IsAdding:     isAdding,
		IsListing:    isListing,
		IsDeleting:   isDeleting,
		IsValue:      isValue,
		OutputFormat: outputFormat,
		Name:         name,
		Price:        price,
		Source:       source,
		SpotPrice:    spotPrice,
		Units:        units,
		Weight:       weight,
		Type:         holdingType,
		DeleteID:     deleteID,
	}

	return hydratedFlags, nil
}

func isValidFlags(flags *types.HoldingFlags) error {
	trues := 0
	for _, value := range []bool{flags.IsAdding, flags.IsListing, flags.IsDeleting} {
		if value {
			trues += 1
		}
	}

	if trues > 1 {
		return errors.New("invalid flags passed")
	}

	if flags.OutputFormat != Json && flags.OutputFormat != Table {
		return errors.New("invalid flags passed")
	}

	// Validate holding type if provided
	if flags.Type != "" && flags.Type != "Silver" && flags.Type != "Gold" && flags.Type != "Platinum" {
		return errors.New("invalid holding type, must be Silver, Gold, or Platinum")
	}

	return nil
}

func handleHolding(cmd *cobra.Command, args []string) {
	flags, err := newHoldingFlags(cmd.Flags())
	if err != nil || isValidFlags(flags) != nil {
		fmt.Println("please provide a valid signiture, run `holding --help` for more information")
		return
	}

	holdingService, err := services.GetHoldingService(flags.OutputFormat)
	if err != nil {
		fmt.Printf("there was a problem resolving the holding service: %v", err)
		return
	}

	if flags.IsDeleting {
		holdingService.Delete(flags.DeleteID)
		return
	}

	if flags.IsListing {
		holdingService.List()
		return
	}

	if flags.IsValue {
		holdingService.GetValue()
		return
	}

	holdingService.Add(flags)
}

var holdingCmd = &cobra.Command{
	Use:   "holding [--add | -a | --delete | -d | --list | -l | --value | -v] [--format output | -f output]",
	Short: "Interacts with holdings in the system",
	Long:  `This command allows you to interact with all holdings in the system. This allows you to add new holdings, view your holding, and do other operations on them as you use the flags this command provides.`,
	Run:   handleHolding,
}

func init() {
	holdingCmd.Flags().BoolP("add", "a", false, "tells the command you want to add a holding")
	holdingCmd.Flags().BoolP("delete", "d", false, "tells the command you want to delete a holding")
	holdingCmd.Flags().BoolP("list", "l", false, "tells the command you want to list your holdings")
	holdingCmd.Flags().BoolP("value", "v", false, "tells the command you want to get your holdings value")
	holdingCmd.Flags().StringP("format", "f", Table, fmt.Sprintf("decides the output format, supports: [\"%v\", \"%v\"]", Json, Table))

	// CLI args for adding holdings (to bypass wizard)
	holdingCmd.Flags().StringP("name", "n", "", "product name for the holding")
	holdingCmd.Flags().StringP("price", "p", "", "purchase price of the holding")
	holdingCmd.Flags().StringP("source", "s", "", "purchase source of the holding")
	holdingCmd.Flags().String("spot-price", "", "spot price at time of purchase")
	holdingCmd.Flags().StringP("units", "u", "", "total number of units")
	holdingCmd.Flags().StringP("weight", "w", "", "weight of a single unit (in toz)")
	holdingCmd.Flags().StringP("type", "t", "", "holding type (Silver, Gold, or Platinum)")

	// CLI args for deleting holdings (to bypass wizard)
	holdingCmd.Flags().String("id", "", "holding ID to delete")

	rootCmd.AddCommand(holdingCmd)
}
