package geonames

import (
	"context"
	"fmt"
)

type srtm1Resp struct {
	Srtm1 int     `json:"srtm1"`
	Lng   float64 `json:"lng"`
	Lat   float64 `json:"lat"`
}

type srtm3Resp struct {
	Srtm3 int     `json:"srtm3"`
	Lng   float64 `json:"lng"`
	Lat   float64 `json:"lat"`
}

type asterResp struct {
	Astergdem int     `json:"astergdem"`
	Lng       float64 `json:"lng"`
	Lat       float64 `json:"lat"`
}

type gtopoResp struct {
	Gtopo30 int     `json:"gtopo30"`
	Lng     float64 `json:"lng"`
	Lat     float64 `json:"lat"`
}

// Elevation holds elevation data expressed in meters npm.
type Elevation struct {
	Type  string
	Lat   float64
	Lng   float64
	Value int
}

// GetElevationSRTM1 takes two float numbers representing latitude and longitude
// and returns elevation in meters according to SRMT1. The sample area is ca 30m x 30m.
// Ocean areas returns "no data", and have assigned a value of -32768.
func (c *Client) GetElevationSRTM1(ctx context.Context, lat, lng float64) (Elevation, error) {
	path := fmt.Sprintf("/srtm1JSON?lat=%.3f&lng=%.3f&username=%s", lat, lng, c.UserName)
	var er srtm1Resp
	err := c.get(ctx, c.BaseURL+path, &er)
	if err != nil {
		return Elevation{}, err
	}
	e := Elevation{
		Type:  "srtm1",
		Lat:   er.Lat,
		Lng:   er.Lng,
		Value: er.Srtm1,
	}
	return e, nil
}

// GetElevationSRTM3 takes two float numbers representing latitude and longitude
// and returns elevation in meters according to SRMT3.
//
// SRTM data consisted of a specially modified radar system that flew
// onboard the Space Shuttle Endeavour during an 11-day mission in February of 2000.
// The dataset covers land areas between 60 degrees north and 56 degrees south.
// SRTM3 data are data points located every 3-arc-second (approximately 90 meters) on a latitude/longitude grid.
func (c *Client) GetElevationSRTM3(ctx context.Context, lat, lng float64) (Elevation, error) {
	path := fmt.Sprintf("/srtm3JSON?lat=%.3f&lng=%.3f&username=%s", lat, lng, c.UserName)
	var er srtm3Resp
	err := c.get(ctx, c.BaseURL+path, &er)
	if err != nil {
		return Elevation{}, err
	}
	e := Elevation{
		Type:  "srtm3",
		Lat:   er.Lat,
		Lng:   er.Lng,
		Value: er.Srtm3,
	}
	return e, nil
}

// GetElevationAstergdem returns elevation in meters according to aster gdem.
//
// Sample are: ca 30m x 30m, between 83N and 65S latitude. Ocean areas have been assigned a value of -32768
func (c *Client) GetElevationAstergdem(ctx context.Context, lat, lng float64) (Elevation, error) {
	path := fmt.Sprintf("/astergdemJSON?lat=%.3f&lng=%.3f&username=%s", lat, lng, c.UserName)
	var er asterResp
	err := c.get(ctx, c.BaseURL+path, &er)
	if err != nil {
		return Elevation{}, err
	}
	e := Elevation{
		Type:  "astergdem",
		Lat:   er.Lat,
		Lng:   er.Lng,
		Value: er.Astergdem,
	}
	return e, nil
}

// GetElevationGTOPO30 returns elevation data sampled for the area of 1km x 1km.
//
// Ocean areas have assigned value of -9999 indicating no data for the requested lat and lng.
// GTOPO30 is a global digital elevation model (DEM) with a horizontal grid spacing
// of 30 arc seconds (approximately 1 kilometer).
// GTOPO30 is derived from several raster and vector sources of topographic information.
//
// Documentation: http://eros.usgs.gov/#/Find_Data/Products_and_Data_Available/gtopo30_info
func (c *Client) GetElevationGTOPO30(ctx context.Context, lat, lng float64) (Elevation, error) {
	path := fmt.Sprintf("/gtopo30JSON?lat=%.3f&lng=%.3f&username=%s", lat, lng, c.UserName)
	var er gtopoResp
	err := c.get(ctx, c.BaseURL+path, &er)
	if err != nil {
		return Elevation{}, err
	}
	e := Elevation{
		Type:  "gtopo30",
		Lat:   er.Lat,
		Lng:   er.Lng,
		Value: er.Gtopo30,
	}
	return e, nil
}
