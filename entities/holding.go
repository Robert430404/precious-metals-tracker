package entities

type HoldingType = string

const (
	Silver HoldingType = "Silver"
	Gold   HoldingType = "Gold"
)

type Holding struct {
	Type HoldingType
}
