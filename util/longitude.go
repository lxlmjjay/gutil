package util

import (
	"math"
)

/**
116.34008339  39.9997921
SELECT
	*
FROM
	`bh_group`
WHERE
	(
	ACOS(
	SIN( ( 39.9997921 * 3.1415 ) / 180 ) * SIN( ( latitude * 3.1415 ) / 180 ) + COS( ( 39.9997921 * 3.1415 ) / 180 ) * COS( ( latitude * 3.1415 ) / 180 ) * COS( ( 116.34008339 * 3.1415 ) / 180 - ( longitude * 3.1415 ) / 180 )
	) * 6370.996
	) <= 3;
*/
//转化为弧度(rad)
func rad(d float64) (r float64) {
	r = d * math.Pi / 180.0
	return
}
//经传度转化为距离
func LonLatToDistance(lon1, lat1, lon2, lat2 float64) (distance float64) {
	//赤道半径(单位m)
	const EARTH_RADIUS = 6378137
	rad_lat1 := rad(lat1)
	rad_lon1 := rad(lon1)
	rad_lat2 := rad(lat2)
	rad_lon2 := rad(lon2)
	if rad_lat1 < 0 {
		rad_lat1 = math.Pi/2 + math.Abs(rad_lat1)
	}
	if rad_lat1 > 0 {
		rad_lat1 = math.Pi/2 - math.Abs(rad_lat1)
	}
	if rad_lon1 < 0 {
		rad_lon1 = math.Pi*2 - math.Abs(rad_lon1)
	}
	if rad_lat2 < 0 {
		rad_lat2 = math.Pi/2 + math.Abs(rad_lat2)
	}
	if rad_lat2 > 0 {
		rad_lat2 = math.Pi/2 - math.Abs(rad_lat2)
	}
	if rad_lon2 < 0 {
		rad_lon2 = math.Pi*2 - math.Abs(rad_lon2)
	}
	x1 := EARTH_RADIUS * math.Cos(rad_lon1) * math.Sin(rad_lat1)
	y1 := EARTH_RADIUS * math.Sin(rad_lon1) * math.Sin(rad_lat1)
	z1 := EARTH_RADIUS * math.Cos(rad_lat1)

	x2 := EARTH_RADIUS * math.Cos(rad_lon2) * math.Sin(rad_lat2)
	y2 := EARTH_RADIUS * math.Sin(rad_lon2) * math.Sin(rad_lat2)
	z2 := EARTH_RADIUS * math.Cos(rad_lat2)
	d := math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2) + (z1-z2)*(z1-z2))
	theta := math.Acos((EARTH_RADIUS*EARTH_RADIUS + EARTH_RADIUS*EARTH_RADIUS - d*d) / (2 * EARTH_RADIUS * EARTH_RADIUS))
	distance = theta * EARTH_RADIUS
	return
}