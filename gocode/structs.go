package gocode

type Addressing struct {
	Feature []AddressFeatures `json:"features"`
}
type AddressFeatures struct {
	Type       string          `json:"type"`
	Properties AddressProperty `json:"properties"`
	Geometry   AddressGeometry `json:"geometry"`
}
type AddressProperty struct {
	ID                    float64     `json:"ID"`
	MIN_SYM     string    `json:"MIN_SYM"`
	POINT_X            float64   `json:"POINT_X"`
	POINT_Y     float64    `json:"POINT_Y"`
	MINERAL        string     `json:"MINERAL"`
	PRIMARY_DI        string    `json:"PRIMARY_DI"`
	NEW_ABBREV   string  `json:"NEW_ABBREV"`
	LABEL_1_SP        string  `json:"LABEL_1_SP"`
	LABEL_2_VA      string  `json:"LABEL_2_VA"`
	REMARKS                string    `json:"REMARKS"`
	
}
type AddressGeometry struct {
	Type        string      `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}
type CombinedCountyMinerals struct {
	Address   AddressFeatures `json:"features"`
	COUNTY_NAM string                 `json:"COUNTY_NAM"`
}
type Combined struct{
	Addresses  []CombinedCountyMinerals `json:"addresses"`
}

type AllZoning struct {
}
type Zoning struct {
	Feature []Feature `json:"features"`
}
type Feature struct {
	ID      int `json:"ID"`
	Type       string   `json:"type"`
	
	Properties Property `json:"properties"`
	Geometry   Geometry `json:"geometry"`
}
type Property struct {
	OBJECTID         int     `json:"OBJECTID"`
	ID_   int     `json:"ID_"`
	COUNTY_NAM   string  `json:"COUNTY_NAM"`
	CONST_CODE int  `json:"CONST_CODE"`
	CONSTITUEN string     `json:"CONSTITUEN"`
	COUNTY_COD   int `json:"COUNTY_COD"`
	Shape_Leng    float64    `json:"Shape_Leng"`
	Shape_Area     float64    `json:"Shape_Area"`
	
	
	
}
type Geometry struct {
	Type        string          `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
}
