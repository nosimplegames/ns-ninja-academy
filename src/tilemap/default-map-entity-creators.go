package tilemap

const (
	MapLimitMapKey int = 1
)

func GetDefaultMapEntityCreators() MapEntityCreators {
	creators := MapEntityCreators{}

	creators[MapLimitMapKey] = &FloorFactory{}

	return creators
}
