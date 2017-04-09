package archetype

import "github.com/golang/geo/r3"

// ShipMoorType sets which type of moors/docks this ship can use.
type ShipMoorType uint

const (
	// ShipMoorBerths indicates that this ship can use fighter/
	// small transport berths.
	ShipMoorBerths = 1
	// ShipMoorSmallMoors indicates that this ship can use small moors.
	ShipMoorSmallMoors = 4
	// ShipMoorMedMoors indicates that this ship can use medium-sized moors.
	ShipMoorMedMoors = 8
	// ShipMoorLargeMoors indicates that this ship can use large moors.
	ShipMoorLargeMoors = 16
)

type Ship struct {
	Arch

	// CargoHoldSize is this ship's cargo hold size (uh).
	CargoHoldSize float64

	// MoorType defines this ship's docking parameters.
	// Original FL files use xor and bit magic to count this; see
	// ShipMoorType values.
	// Originally MissionProperty.
	MoorType map[ShipMoorType]struct{}

	// NanobotLimit is max capacity of nonobot storage.
	NanobotLimit uint

	// ShieldBatteryLimit is max capacity of shield batteries' storage.
	ShieldBatteryLimit uint

	// PhysAngularDrag is the resistance to the steering torque.
	// max_rotational_velocity (radian/s) = steering_torque / angular_drag
	PhysAngularDrag r3.Vector

	// PhysLinearDrag is the force appllied to slow the ship.
	PhysLinearDrag float64

	// PhysNudgeForce is the force applied to move the ship side to side when
	// hitting (avoiding?) rocks.
	PhysNudgeForce float64

	// PhysRotationalIntertia is the amount of initial resistance to moving the
	// centerline on both the start and end of a turn.
	// Kind of like why a car plows down at the nose when you brake hard.
	// The Inertia controls how snappy or mushy the craft handles in flight.
	// rotational_acceleration (radian/s*s) = steering_torque / rotation_inertia;
	PhysRotationalIntertia r3.Vector

	// PhysSteeringTorqu is the amount of force applied to the centerline of the craft to make it turn;
	PhysSteeringTorque r3.Vector

	// PhysStrafeForce is the force applied to move the ship side to side
	// when strafe keys are pressed.
	PhysStrafeForce float64
}
