package types

// HoldingFlags represents CLI flags for holding operations
type HoldingFlags struct {
	IsAdding     bool
	IsListing    bool
	IsDeleting   bool
	IsValue      bool
	OutputFormat string
	// CLI args for adding holdings (bypass wizard)
	Name      string
	Price     string
	Source    string
	SpotPrice string
	Units     string
	Weight    string
	Type      string
	// CLI args for deleting holdings (bypass wizard)
	DeleteID string
}

func (flags *HoldingFlags) HasAllAddFields() bool {
	return flags.Name != "" && flags.Price != "" && flags.Source != "" &&
		flags.SpotPrice != "" && flags.Units != "" && flags.Weight != "" && flags.Type != ""
}
