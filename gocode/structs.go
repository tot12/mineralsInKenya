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
	ID                    int     `json:"_id"`
	ADDRESS_POINT_ID      int     `json:"ADDRESS_POINT_ID"`
	ADDRESS_ID            int     `json:"ADDRESS_ID"`
	ADDRESS_STRING_ID     int     `json:"ADDRESS_STRING_ID"`
	LINEAR_NAME_ID        int     `json:"LINEAR_NAME_ID"`
	CENTRELINE_ID         int     `json:"CENTRELINE_ID"`
	MAINT_STAGE           string  `json:"MAINT_STAGE"`
	ADDRESS_NUMBER        string  `json:"ADDRESS_NUMBER"`
	LINEAR_NAME_FULL      string  `json:"LINEAR_NAME_FULL"`
	LO_NUM                int     `json:"LO_NUM"`
	LO_NUM_SUF            string  `json:"LO_NUM_SUF"`
	HI_NUM                int     `json:"HI_NUM"`
	HI_NUM_SUF            string  `json:"HI_NUM_SUF"`
	LINEAR_NAME           string  `json:"LINEAR_NAME"`
	LINEAR_NAME_TYPE      string  `json:"LINEAR_NAME_TYPE"`
	LINEAR_NAME_DIR       string  `json:"LINEAR_NAME_DIR"`
	LINEAR_NAME_DESC      string  `json:"LINEAR_NAME_DESC"`
	CENTRELINE_SIDE       string  `json:"CENTRELINE_SIDE"`
	CENTRELINE_MEASURE    float64 `json:"CENTRELINE_MEASURE"`
	CENTRELINE_OFFSET     float64 `json:"CENTRELINE_OFFSET"`
	GENERAL_USE_CODE      int     `json:"GENERAL_USE_CODE"`
	GENERAL_USE           string  `json:"GENERAL_USE"`
	CLASS_FAMILY          int     `json:"CLASS_FAMILY"`
	CLASS_FAMILY_DESC     string  `json:"CLASS_FAMILY_DESC"`
	ADDRESS_CLASS         string  `json:"ADDRESS_CLASS"`
	ADDRESS_CLASS_DESC    string  `json:"ADDRESS_CLASS_DESC"`
	ADDRESS_POINT_ID_LINK int     `json:"ADDRESS_POINT_ID_LINK"`
	ADDRESS_ID_LINK       int     `json:"ADDRESS_ID_LINK"`
	PLACE_NAME            string  `json:"PLACE_NAME"`
	PLACE_NAME_ALL        string  `json:"PLACE_NAME_ALL"`
	ADDRESS_STATUS        string  `json:"ADDRESS_STATUS"`
	OBJECTID              string  `json:"OBJECTID"`
	MUNICIPALITY          string  `json:"MUNICIPALITY"`
	MUNICIPALITY_NAME     string  `json:"MUNICIPALITY_NAME"`
	WARD                  string  `json:"WARD"`
	WARD_NAME             string  `json:"WARD_NAME"`
	ADDRESS_FULL          string  `json:"ADDRESS_FULL"`
}
type AddressGeometry struct {
	Type        string      `json:"type"`
	Coordinates [][]float64 `json:"coordinates"`
}
type CombinedAddressZN_STRING struct {
	Address   AddressFeatures `json:"features"`
	ZN_STRING string                 `json:"ZN_STRING"`
}
type Combined struct{
	Addresses  []CombinedAddressZN_STRING `json:"addresses"`
}
