package archprojectile

import "github.com/utrack/goflos/domain/static/archetype"

// Countermeasure is CM projectile archetype.
type Countermeasure struct {
	archetype.Projectile

	// DiversionChance is the chance that CM would divert the charge.
	DiversionChance float32
	LinearDrag      float32
	// Range is this CM's effective max range.
	Range float32
}

// Mine is the mine projectile archetype.
type Mine struct {
	archetype.Projectile

	// Acceleration is this mine's accel rate.
	// Mines are following the target in FL now.
	Acceleration float32
	// DetonationDist is the distance to the ship on which mine detonates.
	DetonationDist float32
	LinearDrag     float32
	SeekerDist     float32
	TopSpeed       float32
}

// Munition is the munition's archetype.
// Single launched rocket/CM/torpedo,etc.
type Munition struct {
	archetype.Projectile

	DisruptsCruise bool
	DetonationDist float32

	DmgHull   float32
	DmgEnergy float32

	MaxAngularVelocity float32
	SeekerFovDeg       float32
	SeekerRange        float32
	TimeToLock         float32

	weaponTypeHash uint64
	motorHash      uint64
	// TODO should remove?
	seeker string
}
