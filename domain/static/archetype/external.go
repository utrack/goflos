package archetype

import (
	"time"

	"github.com/golang/geo/r3"
)

// EquipLauncher is the ExternalArchetype for launchers mounted on ships.
type EquipLauncher struct {
	arch
	DamagePerFire  float32
	MuzzleVelocity r3.Vector
	PowerUsage     float64
	RefireDelay    time.Duration

	ProjectileArch *Projectile
}

// EquipShieldGenerator is an ExternalArchetype for shield generators.
type EquipShieldGenerator struct {
	arch
	ConstantPowerDraw float64
	MaxCapacity       float64

	// OfflineRebuildTime is the time after which shield gets back up.
	OfflineRebuildTime time.Duration
	// OfflineThreshold is the penalty to usable MaxCapacity.
	// It seems that in original game, usable capacity = maxC - (ot*100)%
	// Usually it is set to 0.15, so 15% of capacity is really unusable
	// and shield should get dropped after this 15% amount.
	OfflineThreshold float32

	// RebuildPowerDraw is the amount of power drained(+Constant) when
	// shield is regenerating.
	RebuildPowerDraw float64
	// RegenerationRate is amount of capacity restored per second.
	RegenerationRate float64

	// Type is this shield's type.
	Type ShieldType
}

// ShieldType is an enum alias for ShieldGenerator.Type.
type ShieldType uint

const (
	_ = iota
	// ShieldTypeGraviton is graviton shield.
	ShieldTypeGraviton
	// ShieldTypePositron is positron shield.
	ShieldTypePositron
	// ShieldTypeMolecular is molecular shield.
	ShieldTypeMolecular
)

// EquipThruster is an ExternalArchetype for thrusters.
type EquipThruster struct {
	arch
	// MaxForce of this thruster.
	MaxForce float64
	// PowerUsage is this thruster's power usage when boosting.
	PowerUsage float64
}
