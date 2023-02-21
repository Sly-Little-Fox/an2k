package main

const (
	LeftFree   = "слева свободно"
	RightFree  = "справа свободно"
	TopFree    = "сверху свободно"
	BottomFree = "снизу свободно"
)

const (
	LeftOccupied   = "слева стена"
	RightOccupied  = "справа стена"
	TopOccupied    = "сверху стена"
	BottomOccupied = "справа стена"
)

const (
	MoveLeft  = "влево"
	MoveRight = "вправо"
	MoveUp    = "вверх"
	MoveDown  = "вниз"
)

const (
	Fill = "закрасить"
)

var DirReplaceMap = map[string]string{
	"bottomFree":     BottomFree,
	"bottomOccupied": BottomOccupied,
	"leftFree":       LeftFree,
	"leftOccupied":   LeftOccupied,
	"rightFree":      RightFree,
	"rightOccupied":  RightOccupied,
	"topFree":        TopFree,
	"topOccupied":    TopOccupied,
}

var MoveReplaceMap = map[string]string{
	"left":  MoveLeft,
	"right": MoveRight,
	"up":    MoveUp,
	"down":  MoveDown,
	"fill":  Fill,
}
