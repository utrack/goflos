/*Package base provides base-related domain structs.*/
package base

import "github.com/utrack/goflos/domain/static/market"

const (
	defaultPriceMod     = float32(1)
	defaultSellPriceMod = float32(0.2)
)

// MarketGood is a wrapper around Good, unique for the base.
// It contains this good's price and buying limit.
type MarketGood struct {
	id uint64

	Good market.Good

	minLevelToBuy float32
	minRepToBuy   float32

	// SellDo is true if this MarketGood is sold at the base.
	SellDo bool

	// SellMax sets how many things you can buy at once (during one
	// dock-undock sequence).
	SellMax uint64

	// BuyPriceMod sets the percent that the base pays for this item
	// when buying from the player.
	// Ex.: item of $500 with BPM 0.2 can be sold to the base for $100.
	BuyPriceMod float32
}

// ID returns an ID of this MarketGood.
func (mg MarketGood) ID() uint64 {
	return mg.id
}
