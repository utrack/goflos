package market

import "github.com/utrack/goflos/domain/static/archetype"

type GoodCategory uint

const (
	_ = iota
	GoodShipHull
	GoodShipPackage
	GoodCommodity
	GoodEquipment
)

// Good describes one thing that can be sold or bought over universe.
// UniGood in FLOS.
type Good struct {
	Arch archetype.Arch

	// Category is this good's category.
	Category GoodCategory

	// BasePrice is this Good's base price without any modifiers applied.
	BasePrice float64

	// TODO well now what is that. Unused
	// src/GameDB/BaseDB.cs
	Combinable bool
}
