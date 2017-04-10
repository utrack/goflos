package archetype

import "time"

// Arch is the basic archetype data.
// It can be anything from this pkg.
type Arch interface {
	ID() uint64
	Nickname() string
}

// arch is the basic archetype data.
type arch struct {
	ArchID       uint64
	ArchNickname string
}

func (a arch) ID() uint64 {
	return a.ArchID
}

func (a arch) Nickname() string {
	return a.ArchNickname
}

// Equipment is the ship's equipment archetype.
type Equipment struct {
	arch
	// Lootable is true if this equipment can be shot out and looted.
	Lootable bool
	// UnitsPerContainer is the number of these units per one cargo pod.
	UnitsPerContainer uint64

	// Volume is this thing's occupied cargohold volume.
	Volume float32
}

// EquipMotor is the motor archetype. In official game it is used by missiles' engines
// only.
type EquipMotor struct {
	arch
	Acceleration float32
	// Delay is time after which engine starts.
	Delay   time.Duration
	AiRange float32

	Lifetime time.Duration
}

// EquipArmor is armor archetype.
type EquipArmor struct {
	arch
	// TODO check this. damage multiplier or hitpoints?
	HitPtsScale float64
}

// EquipPower is the powerplant archetype.
type EquipPower struct {
	arch
	// Capacity is total energy storage capacity of this powerplant.
	Capacity float32
	// ChargeRate is this powerplant's charge rate.
	// TODO per sec or per tick?
	ChargeRate float32

	// ThrustCapacity is this powerplant's energy capacity diverged to the
	// boosters/thrusters.
	ThrustCapacity float32

	// ThrustChargeRate is this powerplant's boost charge rate.
	ThrustChargeRate float32
}

// Projectile is the external projectile's archetype (i.e. mines, CMs, rockets, etc)
// see archprojectile package.
type Projectile struct {
	arch

	// TODO check this
	//ForceGunOri bool

	// Lifetime is this projectile's lifetime in seconds.
	Lifetime time.Duration

	// OwnerSafeTime is a time during which projectile can't explode and
	// physical contacts are not accounted.
	OwnerSafeTime time.Duration

	RequiresAmmo bool
}
