package geonames

type ElevationSRTM1 struct {
	Srtm1 int     `json:"srtm1"`
	Lng   float64 `json:"lng"`
	Lat   float64 `json:"lat"`
}

type Elevation struct {
	Type string
	Lat  float64
	Lng  float64
}

func GetElevationSRTM1() Elevation {

	return Elevation{}
}
