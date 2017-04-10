package base

// Base is the Base's type.
// It contains info about rooms, goods for sale, news,...
type Base struct {
	ID       uint64
	Nickname string

	Goods []MarketGood
	News  []NewsItem

	// Rooms is a set of rooms for this base.
	// map is of a nickname -> room
	Rooms map[string]Room
	// RoomStart is a ptr to the starting room.
	RoomStart *Room
}
