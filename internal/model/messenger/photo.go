package messenger

type Photo struct {
	ID    int    `json:"id"`
	Date  int    `json:"date"`
	Sizes []Size `json:"sizes"`
}

type Size struct {
	Type   SizeType `json:"type"`
	Url    string   `json:"url"`
	Width  int      `json:"width"`
	Height int      `json:"height"`
}

type SizeType string

func (st SizeType) Priority() int {
	return SizePriority[st]
}

const (
	SizeTypeS SizeType = "s"
	SizeTypeM SizeType = "m"
	SizeTypeX SizeType = "x"
	SizeTypeO SizeType = "o"
	SizeTypeP SizeType = "p"
	SizeTypeQ SizeType = "q"
	SizeTypeR SizeType = "r"
	SizeTypeY SizeType = "y"
	SizeTypeZ SizeType = "z"
	SizeTypeW SizeType = "w"
)

var SizePriority = map[SizeType]int{
	SizeTypeS: 1,
	SizeTypeM: 2,
	SizeTypeX: 3,
	SizeTypeO: 4,
	SizeTypeP: 5,
	SizeTypeQ: 6,
	SizeTypeR: 7,
	SizeTypeY: 8,
	SizeTypeZ: 9,
	SizeTypeW: 10,
}
