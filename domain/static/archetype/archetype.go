package archetype

// Arch is the basic archetype data.
type Arch struct {
	ArchID   uint64
	Nickname string
}

// Equipment is the ship's equipment archetype.
type Equipment struct {
	Arch
	// Lootable is true if this equipment can be shot out and looted.
	Lootable bool
	// UnitsPerContainer is the number of these units per one cargo pod.
	UnitsPerContainer uint64

	// TODO dunno. src/GameDB/Arch/EquipmentArchetype
	// maybe weight
	Volume float32
}

// Motor is the motor archetype. TODO missiles only?
type Motor struct {
	Arch
	Acceleration float32
	// TODO check this. motor delay used in missiles
	Delay   float32
	AiRange float32

	Lifetime float32
}

// EquipArmor is armor archetype.
type EquipArmor struct {
	Arch
	// TODO check this. damage multiplier or hitpoints?
	HitPtsScale float64
}

// EquipPower is the powerplant archetype.
type EquipPower struct {
	Arch
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
	Arch

	// TODO check this
	//ForceGunOri bool

	// Lifetime is this projectile's lifetime in seconds.
	Lifetime float32

	// OwnerSafeTime is a time during which projectile can't explode and
	// physical contacts are not accounted.
	OwnerSafeTime float32

	RequiresAmmo bool
}
