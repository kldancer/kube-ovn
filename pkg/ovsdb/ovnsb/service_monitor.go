// Code generated by "libovsdb.modelgen"
// DO NOT EDIT.

package ovnsb

const ServiceMonitorTable = "Service_Monitor"

type (
	ServiceMonitorProtocol = string
	ServiceMonitorStatus   = string
)

var (
	ServiceMonitorProtocolTCP   ServiceMonitorProtocol = "tcp"
	ServiceMonitorProtocolUDP   ServiceMonitorProtocol = "udp"
	ServiceMonitorStatusOnline  ServiceMonitorStatus   = "online"
	ServiceMonitorStatusOffline ServiceMonitorStatus   = "offline"
	ServiceMonitorStatusError   ServiceMonitorStatus   = "error"
)

// ServiceMonitor defines an object in Service_Monitor table
type ServiceMonitor struct {
	UUID        string                  `ovsdb:"_uuid"`
	ChassisName string                  `ovsdb:"chassis_name"`
	ExternalIDs map[string]string       `ovsdb:"external_ids"`
	IP          string                  `ovsdb:"ip"`
	LogicalPort string                  `ovsdb:"logical_port"`
	Options     map[string]string       `ovsdb:"options"`
	Port        int                     `ovsdb:"port"`
	Protocol    *ServiceMonitorProtocol `ovsdb:"protocol"`
	SrcIP       string                  `ovsdb:"src_ip"`
	SrcMAC      string                  `ovsdb:"src_mac"`
	Status      *ServiceMonitorStatus   `ovsdb:"status"`
}
