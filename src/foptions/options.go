package foptions

import "fmt"

type House struct {
	Material     string
	HasFireplace bool
	Floors       int
}

func NewHouse(opts ...HouseOption) *House {
	const (
		defaultFloors    = 2
		defaultMaterial  = "straw"
		defaultFireplace = false
	)

	h := &House{
		Material:     defaultMaterial,
		HasFireplace: defaultFireplace,
		Floors:       defaultFloors,
	}

	for _, opt := range opts {
		opt(h)
	}

	return h
}

func (h *House) Present() {
	pattern := "This house has %d floors, has no fireplace and it is made of %s\n"
	if h.HasFireplace {
		pattern = "This house has %d floors, has a fireplace and it is made of %s\n"
	}

	fmt.Printf(pattern, h.Floors, h.Material)
}

type HouseOption func(*House)

func WithConcrete() HouseOption {
	return func(h *House) {
		h.Material = "concrete"
	}
}

func WithWood() HouseOption {
	return func(h *House) {
		h.Material = "wood"
	}
}

func WithFireplace() HouseOption {
	return func(h *House) {
		h.HasFireplace = true
	}
}

func WithoutFireplace() HouseOption {
	return func(h *House) {
		h.HasFireplace = false
	}
}

func SetFloors(f int) HouseOption {
	return func(h *House) {
		h.Floors = f
	}
}
