package iota

type IsAaa int

const (
	_ IsAaa = iota
	Status1
	Status2
	Status3
)

var aaMap = map[IsAaa]string{
	Status1: "呵1",
	Status2: "呵2",
	Status3: "呵3",
}

func (aa IsAaa) String() string {
	return aaMap[aa]
}
