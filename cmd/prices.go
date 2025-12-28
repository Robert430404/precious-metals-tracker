package cmd

import (
	"errors"
	"fmt"

	"github.com/robert430404/precious-metals-tracker/http/pricing"
	"github.com/robert430404/precious-metals-tracker/renderers"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type PricesFlags struct {
	OutputFormat OutputFormat
}

func newPricesFlags(flags *pflag.FlagSet) (*PricesFlags, error) {
	outputFormat, err := flags.GetString("format")

	if err != nil {
		return nil, errors.New("could not parse flags")
	}

	hydratedFlags := &PricesFlags{
		OutputFormat: outputFormat,
	}

	return hydratedFlags, nil
}

func (self *PricesFlags) isPricesFlagsValid() error {
	if self.OutputFormat != Json && self.OutputFormat != Table {
		return errors.New("invalid flags passed")
	}

	return nil
}

func handlePrices(cmd *cobra.Command, args []string) {
	flags, err := newPricesFlags(cmd.Flags())
	if err != nil || flags.isPricesFlagsValid() != nil {
		fmt.Println("please provide a valid signiture, run `prices --help` for more information")
		return
	}

	repository, err := pricing.GetPricingRepository()
	if err != nil {
		fmt.Printf("could not resolve pricing repository: %v", err)
		return
	}

	outputType := flags.OutputFormat

	var renderer renderers.Renderer = nil
	if outputType == "json" {
		renderer = &renderers.JsonRenderer{}
	} else {
		renderer = &renderers.TableRenderer{}
	}

	silverSpotPrice := repository.GetSilverSpot()
	goldSpotPrice := repository.GetGoldSpot()

	renderer.RenderSpotPricing(fmt.Sprintf("$%.2f", silverSpotPrice), fmt.Sprintf("$%.2f", goldSpotPrice))
}

var pricesCmd = &cobra.Command{
	Use:   "prices [--format output | -f output]",
	Short: "Pulls the prices using Gold API",
	Long:  `This command allows you to pull metal prices from Gold API io.`,
	Run:   handlePrices,
}

func init() {
	pricesCmd.Flags().StringP("format", "f", Table, fmt.Sprintf("decides the output format, supports: [\"%v\", \"%v\"]", Json, Table))

	rootCmd.AddCommand(pricesCmd)
}
