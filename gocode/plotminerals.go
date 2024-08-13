
// try and do zoning per county 
//check if coordinates are in a certain county and add to file 
package gocode

import (
	"encoding/json"
	"fmt"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"github.com/paulmach/orb/planar"
	"io/ioutil"
	
)

func Poly()(countyminerals Combined) {

	const (
		GEO_FILE = "geojson/kenya.geojson"
	)

	b, _ := ioutil.ReadFile(GEO_FILE)
	featureCollection, _ := geojson.UnmarshalFeatureCollection(b)

	CallAddressAPI()

	CallZoningAPI()

	zoning := Zone
	Address1 := Address

	

	
	var allslice []CombinedCountyMinerals
	for _, a := range Address1.Feature {
		//for _, l := range a.Geometry.Coordinates {
			b1, zn := isPointInsidePolygon(featureCollection, zoning, orb.Point{a.Geometry.Coordinates[0], a.Geometry.Coordinates[1]})
			if b1 == true {
				newstruct := CombinedCountyMinerals{Address: a,COUNTY_NAM : zn}
				allslice = append(allslice, newstruct)
				//fmt.Printf("found :%v \n",allslice)
			} else {
				//fmt.Println("Point 1 is not found inside Polygon")
			}
		//}
	}
	fmt.Println("len:", len(allslice))
	//type Combined struct {
	//	Addresses []CombinedCountyMinerals `json:"addresses"`
	//}
	combined := Combined{Addresses: allslice}
	file, _ := json.MarshalIndent(combined, "", " ")
	_ = ioutil.WriteFile("CombinedAddresses.json", file, 0644)

	// isPointInsidePolygon runs through the MultiPolygon and Polygons within a
	// feature collection and checks if a point (long/lat) lies within it.

	return combined
	

}
func isPointInsidePolygon(fc *geojson.FeatureCollection, feat Zoning, point orb.Point) (bool, string) {
	for i, feature := range fc.Features {
		//if feat.Feature[i].Properties.ZN_ZONE == "CR" && strings.Contains(feat.Feature[i].Properties.ZN_STRING, "r0.0") == false {
			multiPoly, isMulti := feature.Geometry.(orb.MultiPolygon)
			if isMulti {
				if planar.MultiPolygonContains(multiPoly, point) {
					return true, feat.Feature[i].Properties.COUNTY_NAM
				}
			} else {
				//nothing
			}
		//}
	}

	return false, ""
}
