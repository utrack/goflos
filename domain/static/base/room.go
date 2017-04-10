package base

// Room is one room of a base.
type Room struct {
	ID       uint64
	Nickname string

	// CharacterDensity is the (max) number of NPC characters to have
	// in a room.
	CharacterDensity uint
}
