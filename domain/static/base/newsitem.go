package base

// NewsItemIcon is an enum alias for NewsItem.Icon.
type NewsItemIcon uint

const (
	_ = iota
	NewsItemCritical
	NewsItemWorld
	NewsItemMission
	NewsItemSystem
	NewsItemFaction
	NewsItemUniverse
)

// NewsItem is the base data for news item shown in News section
// at bars of bases.
type NewsItem struct {
	HasAudio bool
	Category uint64
	Headline uint64
	Icon     NewsItemIcon
	Logo     string
	Text     uint64
}
