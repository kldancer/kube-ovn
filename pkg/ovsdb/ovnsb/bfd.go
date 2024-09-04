// Code generated by "libovsdb.modelgen"
// DO NOT EDIT.

package ovnsb

const BFDTable = "BFD"

type (
	BFDStatus = string
)

var (
	BFDStatusDown      BFDStatus = "down"
	BFDStatusInit      BFDStatus = "init"
	BFDStatusUp        BFDStatus = "up"
	BFDStatusAdminDown BFDStatus = "admin_down"
)

// BFD defines an object in BFD table
type BFD struct {
	UUID        string            `ovsdb:"_uuid"`
	ChassisName string            `ovsdb:"chassis_name"`
	DetectMult  int               `ovsdb:"detect_mult"`
	Disc        int               `ovsdb:"disc"`
	DstIP       string            `ovsdb:"dst_ip"`
	ExternalIDs map[string]string `ovsdb:"external_ids"`
	LogicalPort string            `ovsdb:"logical_port"`
	MinRx       int               `ovsdb:"min_rx"`
	MinTx       int               `ovsdb:"min_tx"`
	Options     map[string]string `ovsdb:"options"`
	SrcPort     int               `ovsdb:"src_port"`
	Status      BFDStatus         `ovsdb:"status"`
}
