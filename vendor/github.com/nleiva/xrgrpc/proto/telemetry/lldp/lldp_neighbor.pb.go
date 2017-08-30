// Code generated by protoc-gen-go.
// source: lldp_neighbor.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ethernet_lldp_oper_lldp_nodes_node_neighbors_details_detail is a generated protocol buffer package.

It is generated from these files:
	lldp_neighbor.proto

It has these top-level messages:
	LldpNeighbor_KEYS
	LldpNeighbor
	LldpNeighborItem
	In6AddrTd
	LldpL3Addr
	LldpAddrEntry
	LldpAddrEntryItem
	LldpUnknownTlvEntry
	LldpUnknownTlvEntryItem
	LldpOrgDefTlvEntry
	LldpOrgDefTlvEntryItem
	LldpNeighborDetail
	LldpNeighborMib
*/
package cisco_ios_xr_ethernet_lldp_oper_lldp_nodes_node_neighbors_details_detail

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// LLDP neighbor info
type LldpNeighbor_KEYS struct {
	NodeName      string `protobuf:"bytes,1,opt,name=node_name,json=nodeName" json:"node_name,omitempty"`
	InterfaceName string `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
	DeviceId      string `protobuf:"bytes,3,opt,name=device_id,json=deviceId" json:"device_id,omitempty"`
}

func (m *LldpNeighbor_KEYS) Reset()                    { *m = LldpNeighbor_KEYS{} }
func (m *LldpNeighbor_KEYS) String() string            { return proto.CompactTextString(m) }
func (*LldpNeighbor_KEYS) ProtoMessage()               {}
func (*LldpNeighbor_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LldpNeighbor_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *LldpNeighbor_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *LldpNeighbor_KEYS) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

type LldpNeighbor struct {
	// Next neighbor in the list
	LldpNeighbor []*LldpNeighborItem `protobuf:"bytes,50,rep,name=lldp_neighbor,json=lldpNeighbor" json:"lldp_neighbor,omitempty"`
}

func (m *LldpNeighbor) Reset()                    { *m = LldpNeighbor{} }
func (m *LldpNeighbor) String() string            { return proto.CompactTextString(m) }
func (*LldpNeighbor) ProtoMessage()               {}
func (*LldpNeighbor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LldpNeighbor) GetLldpNeighbor() []*LldpNeighborItem {
	if m != nil {
		return m.LldpNeighbor
	}
	return nil
}

type LldpNeighborItem struct {
	// Interface the neighbor entry was received on
	ReceivingInterfaceName string `protobuf:"bytes,1,opt,name=receiving_interface_name,json=receivingInterfaceName" json:"receiving_interface_name,omitempty"`
	// Parent Interface the neighbor entry was received on
	ReceivingParentInterfaceName string `protobuf:"bytes,2,opt,name=receiving_parent_interface_name,json=receivingParentInterfaceName" json:"receiving_parent_interface_name,omitempty"`
	// Device identifier
	DeviceId string `protobuf:"bytes,3,opt,name=device_id,json=deviceId" json:"device_id,omitempty"`
	// Chassis id
	ChassisId string `protobuf:"bytes,4,opt,name=chassis_id,json=chassisId" json:"chassis_id,omitempty"`
	// Outgoing port identifier
	PortIdDetail string `protobuf:"bytes,5,opt,name=port_id_detail,json=portIdDetail" json:"port_id_detail,omitempty"`
	// Version number
	HeaderVersion uint32 `protobuf:"varint,6,opt,name=header_version,json=headerVersion" json:"header_version,omitempty"`
	// Remaining hold time
	HoldTime uint32 `protobuf:"varint,7,opt,name=hold_time,json=holdTime" json:"hold_time,omitempty"`
	// Enabled Capabilities
	EnabledCapabilities string `protobuf:"bytes,8,opt,name=enabled_capabilities,json=enabledCapabilities" json:"enabled_capabilities,omitempty"`
	// Platform type
	Platform string `protobuf:"bytes,9,opt,name=platform" json:"platform,omitempty"`
	// Detailed neighbor info
	Detail *LldpNeighborDetail `protobuf:"bytes,10,opt,name=detail" json:"detail,omitempty"`
	// MIB nieghbor info
	Mib *LldpNeighborMib `protobuf:"bytes,11,opt,name=mib" json:"mib,omitempty"`
}

func (m *LldpNeighborItem) Reset()                    { *m = LldpNeighborItem{} }
func (m *LldpNeighborItem) String() string            { return proto.CompactTextString(m) }
func (*LldpNeighborItem) ProtoMessage()               {}
func (*LldpNeighborItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LldpNeighborItem) GetReceivingInterfaceName() string {
	if m != nil {
		return m.ReceivingInterfaceName
	}
	return ""
}

func (m *LldpNeighborItem) GetReceivingParentInterfaceName() string {
	if m != nil {
		return m.ReceivingParentInterfaceName
	}
	return ""
}

func (m *LldpNeighborItem) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *LldpNeighborItem) GetChassisId() string {
	if m != nil {
		return m.ChassisId
	}
	return ""
}

func (m *LldpNeighborItem) GetPortIdDetail() string {
	if m != nil {
		return m.PortIdDetail
	}
	return ""
}

func (m *LldpNeighborItem) GetHeaderVersion() uint32 {
	if m != nil {
		return m.HeaderVersion
	}
	return 0
}

func (m *LldpNeighborItem) GetHoldTime() uint32 {
	if m != nil {
		return m.HoldTime
	}
	return 0
}

func (m *LldpNeighborItem) GetEnabledCapabilities() string {
	if m != nil {
		return m.EnabledCapabilities
	}
	return ""
}

func (m *LldpNeighborItem) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *LldpNeighborItem) GetDetail() *LldpNeighborDetail {
	if m != nil {
		return m.Detail
	}
	return nil
}

func (m *LldpNeighborItem) GetMib() *LldpNeighborMib {
	if m != nil {
		return m.Mib
	}
	return nil
}

type In6AddrTd struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *In6AddrTd) Reset()                    { *m = In6AddrTd{} }
func (m *In6AddrTd) String() string            { return proto.CompactTextString(m) }
func (*In6AddrTd) ProtoMessage()               {}
func (*In6AddrTd) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *In6AddrTd) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type LldpL3Addr struct {
	AddressType string `protobuf:"bytes,1,opt,name=address_type,json=addressType" json:"address_type,omitempty"`
	// IPv4 address
	Ipv4Address string `protobuf:"bytes,2,opt,name=ipv4_address,json=ipv4Address" json:"ipv4_address,omitempty"`
	// IPv6 address
	Ipv6Address *In6AddrTd `protobuf:"bytes,3,opt,name=ipv6_address,json=ipv6Address" json:"ipv6_address,omitempty"`
}

func (m *LldpL3Addr) Reset()                    { *m = LldpL3Addr{} }
func (m *LldpL3Addr) String() string            { return proto.CompactTextString(m) }
func (*LldpL3Addr) ProtoMessage()               {}
func (*LldpL3Addr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *LldpL3Addr) GetAddressType() string {
	if m != nil {
		return m.AddressType
	}
	return ""
}

func (m *LldpL3Addr) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *LldpL3Addr) GetIpv6Address() *In6AddrTd {
	if m != nil {
		return m.Ipv6Address
	}
	return nil
}

type LldpAddrEntry struct {
	// Next address entry in list
	LldpAddrEntry []*LldpAddrEntryItem `protobuf:"bytes,1,rep,name=lldp_addr_entry,json=lldpAddrEntry" json:"lldp_addr_entry,omitempty"`
}

func (m *LldpAddrEntry) Reset()                    { *m = LldpAddrEntry{} }
func (m *LldpAddrEntry) String() string            { return proto.CompactTextString(m) }
func (*LldpAddrEntry) ProtoMessage()               {}
func (*LldpAddrEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *LldpAddrEntry) GetLldpAddrEntry() []*LldpAddrEntryItem {
	if m != nil {
		return m.LldpAddrEntry
	}
	return nil
}

type LldpAddrEntryItem struct {
	// Network layer address
	Address *LldpL3Addr `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	// MA sub type
	MaSubtype uint32 `protobuf:"varint,2,opt,name=ma_subtype,json=maSubtype" json:"ma_subtype,omitempty"`
	// Interface num
	IfNum uint32 `protobuf:"varint,3,opt,name=if_num,json=ifNum" json:"if_num,omitempty"`
}

func (m *LldpAddrEntryItem) Reset()                    { *m = LldpAddrEntryItem{} }
func (m *LldpAddrEntryItem) String() string            { return proto.CompactTextString(m) }
func (*LldpAddrEntryItem) ProtoMessage()               {}
func (*LldpAddrEntryItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *LldpAddrEntryItem) GetAddress() *LldpL3Addr {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *LldpAddrEntryItem) GetMaSubtype() uint32 {
	if m != nil {
		return m.MaSubtype
	}
	return 0
}

func (m *LldpAddrEntryItem) GetIfNum() uint32 {
	if m != nil {
		return m.IfNum
	}
	return 0
}

type LldpUnknownTlvEntry struct {
	// Next unknown TLV entry in list
	LldpUnknownTlvEntry []*LldpUnknownTlvEntryItem `protobuf:"bytes,1,rep,name=lldp_unknown_tlv_entry,json=lldpUnknownTlvEntry" json:"lldp_unknown_tlv_entry,omitempty"`
}

func (m *LldpUnknownTlvEntry) Reset()                    { *m = LldpUnknownTlvEntry{} }
func (m *LldpUnknownTlvEntry) String() string            { return proto.CompactTextString(m) }
func (*LldpUnknownTlvEntry) ProtoMessage()               {}
func (*LldpUnknownTlvEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *LldpUnknownTlvEntry) GetLldpUnknownTlvEntry() []*LldpUnknownTlvEntryItem {
	if m != nil {
		return m.LldpUnknownTlvEntry
	}
	return nil
}

type LldpUnknownTlvEntryItem struct {
	// Unknown TLV type
	TlvType uint32 `protobuf:"varint,1,opt,name=tlv_type,json=tlvType" json:"tlv_type,omitempty"`
	// Unknown TLV payload
	TlvValue []byte `protobuf:"bytes,2,opt,name=tlv_value,json=tlvValue,proto3" json:"tlv_value,omitempty"`
}

func (m *LldpUnknownTlvEntryItem) Reset()                    { *m = LldpUnknownTlvEntryItem{} }
func (m *LldpUnknownTlvEntryItem) String() string            { return proto.CompactTextString(m) }
func (*LldpUnknownTlvEntryItem) ProtoMessage()               {}
func (*LldpUnknownTlvEntryItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *LldpUnknownTlvEntryItem) GetTlvType() uint32 {
	if m != nil {
		return m.TlvType
	}
	return 0
}

func (m *LldpUnknownTlvEntryItem) GetTlvValue() []byte {
	if m != nil {
		return m.TlvValue
	}
	return nil
}

type LldpOrgDefTlvEntry struct {
	// Next Org Def TLV entry in list
	LldpOrgDefTlvEntry []*LldpOrgDefTlvEntryItem `protobuf:"bytes,1,rep,name=lldp_org_def_tlv_entry,json=lldpOrgDefTlvEntry" json:"lldp_org_def_tlv_entry,omitempty"`
}

func (m *LldpOrgDefTlvEntry) Reset()                    { *m = LldpOrgDefTlvEntry{} }
func (m *LldpOrgDefTlvEntry) String() string            { return proto.CompactTextString(m) }
func (*LldpOrgDefTlvEntry) ProtoMessage()               {}
func (*LldpOrgDefTlvEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *LldpOrgDefTlvEntry) GetLldpOrgDefTlvEntry() []*LldpOrgDefTlvEntryItem {
	if m != nil {
		return m.LldpOrgDefTlvEntry
	}
	return nil
}

type LldpOrgDefTlvEntryItem struct {
	// Organizationally Unique Identifier
	Oui uint32 `protobuf:"varint,1,opt,name=oui" json:"oui,omitempty"`
	// Org Def TLV subtype
	TlvSubtype uint32 `protobuf:"varint,2,opt,name=tlv_subtype,json=tlvSubtype" json:"tlv_subtype,omitempty"`
	// lldpRemOrgDefInfoIndex
	TlvInfoIndes uint32 `protobuf:"varint,3,opt,name=tlv_info_indes,json=tlvInfoIndes" json:"tlv_info_indes,omitempty"`
	// Org Def TLV payload
	TlvValue []byte `protobuf:"bytes,4,opt,name=tlv_value,json=tlvValue,proto3" json:"tlv_value,omitempty"`
}

func (m *LldpOrgDefTlvEntryItem) Reset()                    { *m = LldpOrgDefTlvEntryItem{} }
func (m *LldpOrgDefTlvEntryItem) String() string            { return proto.CompactTextString(m) }
func (*LldpOrgDefTlvEntryItem) ProtoMessage()               {}
func (*LldpOrgDefTlvEntryItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *LldpOrgDefTlvEntryItem) GetOui() uint32 {
	if m != nil {
		return m.Oui
	}
	return 0
}

func (m *LldpOrgDefTlvEntryItem) GetTlvSubtype() uint32 {
	if m != nil {
		return m.TlvSubtype
	}
	return 0
}

func (m *LldpOrgDefTlvEntryItem) GetTlvInfoIndes() uint32 {
	if m != nil {
		return m.TlvInfoIndes
	}
	return 0
}

func (m *LldpOrgDefTlvEntryItem) GetTlvValue() []byte {
	if m != nil {
		return m.TlvValue
	}
	return nil
}

type LldpNeighborDetail struct {
	// Port Description
	PortDescription string `protobuf:"bytes,1,opt,name=port_description,json=portDescription" json:"port_description,omitempty"`
	// System Name
	SystemName string `protobuf:"bytes,2,opt,name=system_name,json=systemName" json:"system_name,omitempty"`
	// System Description
	SystemDescription string `protobuf:"bytes,3,opt,name=system_description,json=systemDescription" json:"system_description,omitempty"`
	// Time remaining
	TimeRemaining uint32 `protobuf:"varint,4,opt,name=time_remaining,json=timeRemaining" json:"time_remaining,omitempty"`
	// System Capabilities
	SystemCapabilities string `protobuf:"bytes,5,opt,name=system_capabilities,json=systemCapabilities" json:"system_capabilities,omitempty"`
	// Enabled Capabilities
	EnabledCapabilities string `protobuf:"bytes,6,opt,name=enabled_capabilities,json=enabledCapabilities" json:"enabled_capabilities,omitempty"`
	// Management Addresses
	NetworkAddresses *LldpAddrEntry `protobuf:"bytes,7,opt,name=network_addresses,json=networkAddresses" json:"network_addresses,omitempty"`
	// Auto Negotiation
	AutoNegotiation string `protobuf:"bytes,8,opt,name=auto_negotiation,json=autoNegotiation" json:"auto_negotiation,omitempty"`
	// Physical media capabilities
	PhysicalMediaCapabilities string `protobuf:"bytes,9,opt,name=physical_media_capabilities,json=physicalMediaCapabilities" json:"physical_media_capabilities,omitempty"`
	// Media Attachment Unit type
	MediaAttachmentUnitType uint32 `protobuf:"varint,10,opt,name=media_attachment_unit_type,json=mediaAttachmentUnitType" json:"media_attachment_unit_type,omitempty"`
	// Vlan ID
	PortVlanId uint32 `protobuf:"varint,11,opt,name=port_vlan_id,json=portVlanId" json:"port_vlan_id,omitempty"`
}

func (m *LldpNeighborDetail) Reset()                    { *m = LldpNeighborDetail{} }
func (m *LldpNeighborDetail) String() string            { return proto.CompactTextString(m) }
func (*LldpNeighborDetail) ProtoMessage()               {}
func (*LldpNeighborDetail) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *LldpNeighborDetail) GetPortDescription() string {
	if m != nil {
		return m.PortDescription
	}
	return ""
}

func (m *LldpNeighborDetail) GetSystemName() string {
	if m != nil {
		return m.SystemName
	}
	return ""
}

func (m *LldpNeighborDetail) GetSystemDescription() string {
	if m != nil {
		return m.SystemDescription
	}
	return ""
}

func (m *LldpNeighborDetail) GetTimeRemaining() uint32 {
	if m != nil {
		return m.TimeRemaining
	}
	return 0
}

func (m *LldpNeighborDetail) GetSystemCapabilities() string {
	if m != nil {
		return m.SystemCapabilities
	}
	return ""
}

func (m *LldpNeighborDetail) GetEnabledCapabilities() string {
	if m != nil {
		return m.EnabledCapabilities
	}
	return ""
}

func (m *LldpNeighborDetail) GetNetworkAddresses() *LldpAddrEntry {
	if m != nil {
		return m.NetworkAddresses
	}
	return nil
}

func (m *LldpNeighborDetail) GetAutoNegotiation() string {
	if m != nil {
		return m.AutoNegotiation
	}
	return ""
}

func (m *LldpNeighborDetail) GetPhysicalMediaCapabilities() string {
	if m != nil {
		return m.PhysicalMediaCapabilities
	}
	return ""
}

func (m *LldpNeighborDetail) GetMediaAttachmentUnitType() uint32 {
	if m != nil {
		return m.MediaAttachmentUnitType
	}
	return 0
}

func (m *LldpNeighborDetail) GetPortVlanId() uint32 {
	if m != nil {
		return m.PortVlanId
	}
	return 0
}

type LldpNeighborMib struct {
	// TimeFilter
	RemTimeMark uint32 `protobuf:"varint,1,opt,name=rem_time_mark,json=remTimeMark" json:"rem_time_mark,omitempty"`
	// LldpPortNumber
	RemLocalPortNum uint32 `protobuf:"varint,2,opt,name=rem_local_port_num,json=remLocalPortNum" json:"rem_local_port_num,omitempty"`
	// lldpRemIndex
	RemIndex uint32 `protobuf:"varint,3,opt,name=rem_index,json=remIndex" json:"rem_index,omitempty"`
	// Chassis ID sub type
	ChassisIdSubType uint32 `protobuf:"varint,4,opt,name=chassis_id_sub_type,json=chassisIdSubType" json:"chassis_id_sub_type,omitempty"`
	// Chassis ID length
	ChassisIdLen uint32 `protobuf:"varint,5,opt,name=chassis_id_len,json=chassisIdLen" json:"chassis_id_len,omitempty"`
	// Port ID sub type
	PortIdSubType uint32 `protobuf:"varint,6,opt,name=port_id_sub_type,json=portIdSubType" json:"port_id_sub_type,omitempty"`
	// Port ID length
	PortIdLen uint32 `protobuf:"varint,7,opt,name=port_id_len,json=portIdLen" json:"port_id_len,omitempty"`
	// Supported and combined cpabilities
	CombinedCapabilities uint32 `protobuf:"varint,8,opt,name=combined_capabilities,json=combinedCapabilities" json:"combined_capabilities,omitempty"`
	// Unknown TLV list
	UnknownTlvList *LldpUnknownTlvEntry `protobuf:"bytes,9,opt,name=unknown_tlv_list,json=unknownTlvList" json:"unknown_tlv_list,omitempty"`
	// Org Def TLV list
	OrgDefTlvList *LldpOrgDefTlvEntry `protobuf:"bytes,10,opt,name=org_def_tlv_list,json=orgDefTlvList" json:"org_def_tlv_list,omitempty"`
}

func (m *LldpNeighborMib) Reset()                    { *m = LldpNeighborMib{} }
func (m *LldpNeighborMib) String() string            { return proto.CompactTextString(m) }
func (*LldpNeighborMib) ProtoMessage()               {}
func (*LldpNeighborMib) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *LldpNeighborMib) GetRemTimeMark() uint32 {
	if m != nil {
		return m.RemTimeMark
	}
	return 0
}

func (m *LldpNeighborMib) GetRemLocalPortNum() uint32 {
	if m != nil {
		return m.RemLocalPortNum
	}
	return 0
}

func (m *LldpNeighborMib) GetRemIndex() uint32 {
	if m != nil {
		return m.RemIndex
	}
	return 0
}

func (m *LldpNeighborMib) GetChassisIdSubType() uint32 {
	if m != nil {
		return m.ChassisIdSubType
	}
	return 0
}

func (m *LldpNeighborMib) GetChassisIdLen() uint32 {
	if m != nil {
		return m.ChassisIdLen
	}
	return 0
}

func (m *LldpNeighborMib) GetPortIdSubType() uint32 {
	if m != nil {
		return m.PortIdSubType
	}
	return 0
}

func (m *LldpNeighborMib) GetPortIdLen() uint32 {
	if m != nil {
		return m.PortIdLen
	}
	return 0
}

func (m *LldpNeighborMib) GetCombinedCapabilities() uint32 {
	if m != nil {
		return m.CombinedCapabilities
	}
	return 0
}

func (m *LldpNeighborMib) GetUnknownTlvList() *LldpUnknownTlvEntry {
	if m != nil {
		return m.UnknownTlvList
	}
	return nil
}

func (m *LldpNeighborMib) GetOrgDefTlvList() *LldpOrgDefTlvEntry {
	if m != nil {
		return m.OrgDefTlvList
	}
	return nil
}

func init() {
	proto.RegisterType((*LldpNeighbor_KEYS)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.details.detail.lldp_neighbor_KEYS")
	proto.RegisterType((*LldpNeighbor)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.details.detail.lldp_neighbor")
	proto.RegisterType((*LldpNeighborItem)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.details.detail.lldp_neighbor_item")
	proto.RegisterType((*In6AddrTd)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.details.detail.in6_addr_td")
	proto.RegisterType((*LldpL3Addr)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.details.detail.lldp_l3_addr")
	proto.RegisterType((*LldpAddrEntry)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.details.detail.lldp_addr_entry")
	proto.RegisterType((*LldpAddrEntryItem)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.details.detail.lldp_addr_entry_item")
	proto.RegisterType((*LldpUnknownTlvEntry)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.details.detail.lldp_unknown_tlv_entry")
	proto.RegisterType((*LldpUnknownTlvEntryItem)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.details.detail.lldp_unknown_tlv_entry_item")
	proto.RegisterType((*LldpOrgDefTlvEntry)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.details.detail.lldp_org_def_tlv_entry")
	proto.RegisterType((*LldpOrgDefTlvEntryItem)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.details.detail.lldp_org_def_tlv_entry_item")
	proto.RegisterType((*LldpNeighborDetail)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.details.detail.lldp_neighbor_detail")
	proto.RegisterType((*LldpNeighborMib)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.details.detail.lldp_neighbor_mib")
}

func init() { proto.RegisterFile("lldp_neighbor.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x57, 0xdf, 0x6e, 0x1b, 0xc5,
	0x17, 0xd6, 0x36, 0x69, 0x6a, 0x9f, 0xf5, 0x26, 0xee, 0x24, 0xed, 0xcf, 0x6d, 0x7e, 0xd0, 0x60,
	0x5a, 0x11, 0x84, 0x6a, 0x44, 0x8a, 0x2a, 0x24, 0x24, 0xa4, 0x8a, 0x56, 0xc2, 0xa2, 0x0d, 0xd5,
	0x36, 0x89, 0x54, 0x81, 0x18, 0xc6, 0xde, 0xb1, 0x3d, 0xca, 0xce, 0xec, 0x6a, 0x76, 0x76, 0x9b,
	0x5c, 0x72, 0x53, 0x09, 0x6e, 0x7b, 0x45, 0x9f, 0x03, 0x6e, 0x79, 0x04, 0xde, 0x04, 0x21, 0xde,
	0x00, 0x9d, 0x99, 0x59, 0xc7, 0xce, 0x1f, 0xc4, 0x45, 0x7c, 0x13, 0xdb, 0xdf, 0xf9, 0x66, 0xce,
	0x9c, 0x6f, 0xce, 0x9f, 0x09, 0xac, 0xa7, 0x69, 0x92, 0x53, 0xc5, 0xc5, 0x78, 0x32, 0xc8, 0x74,
	0x2f, 0xd7, 0x99, 0xc9, 0xc8, 0x57, 0x43, 0x51, 0x0c, 0x33, 0x2a, 0xb2, 0x82, 0x1e, 0x69, 0xca,
	0xcd, 0x84, 0x6b, 0xc5, 0x0d, 0xb5, 0xd4, 0x2c, 0xe7, 0xba, 0x87, 0xdf, 0x7a, 0x2a, 0x4b, 0x78,
	0x61, 0xff, 0xf6, 0xea, 0xf5, 0x45, 0x2f, 0xe1, 0x86, 0x89, 0xb4, 0xfe, 0xec, 0x96, 0x40, 0xe6,
	0x1c, 0xd0, 0xaf, 0x9f, 0xbc, 0x7c, 0x41, 0x36, 0xa1, 0x89, 0xcb, 0xa8, 0x62, 0x92, 0x77, 0x82,
	0xad, 0x60, 0xbb, 0x19, 0x37, 0x10, 0xd8, 0x65, 0x92, 0x93, 0x7b, 0xb0, 0x2a, 0x94, 0xe1, 0x7a,
	0xc4, 0x86, 0x9e, 0x71, 0xc5, 0x32, 0xa2, 0x29, 0x6a, 0x69, 0x9b, 0xd0, 0x4c, 0x78, 0x25, 0x86,
	0x9c, 0x8a, 0xa4, 0xb3, 0xe4, 0xf6, 0x70, 0x40, 0x3f, 0xe9, 0xbe, 0x09, 0x20, 0x9a, 0xf3, 0x4b,
	0x7e, 0x3c, 0x8d, 0x74, 0x76, 0xb6, 0x96, 0xb6, 0xc3, 0x9d, 0xef, 0x7a, 0x97, 0x15, 0x6b, 0x6f,
	0x3e, 0x50, 0x61, 0xb8, 0x8c, 0x5b, 0x88, 0xed, 0x7a, 0xa8, 0xfb, 0xf7, 0xf2, 0x69, 0x35, 0x90,
	0x44, 0x3e, 0x83, 0x8e, 0xe6, 0x43, 0x2e, 0x2a, 0xa1, 0xc6, 0xf4, 0x54, 0xe8, 0x4e, 0x9c, 0x9b,
	0x53, 0x7b, 0x7f, 0x4e, 0x83, 0x27, 0x70, 0xe7, 0x64, 0x65, 0xce, 0x34, 0x57, 0x86, 0x9e, 0xab,
	0xdd, 0xff, 0xa7, 0xb4, 0xe7, 0x96, 0xd5, 0xff, 0xcf, 0x52, 0x92, 0x77, 0x00, 0x86, 0x13, 0x56,
	0x14, 0xa2, 0x40, 0xeb, 0xb2, 0xb5, 0x36, 0x3d, 0xd2, 0x4f, 0xc8, 0x5d, 0x58, 0xcd, 0x33, 0x6d,
	0xa8, 0x48, 0xa8, 0x93, 0xa1, 0x73, 0xd5, 0x52, 0x5a, 0x88, 0xf6, 0x93, 0xc7, 0x16, 0xc3, 0x3b,
	0x9d, 0x70, 0x96, 0x70, 0x4d, 0x2b, 0xae, 0x0b, 0x91, 0xa9, 0xce, 0xca, 0x56, 0xb0, 0x1d, 0xc5,
	0x91, 0x43, 0x0f, 0x1c, 0x88, 0x07, 0x99, 0x64, 0x69, 0x42, 0x8d, 0x90, 0xbc, 0x73, 0xcd, 0x32,
	0x1a, 0x08, 0xec, 0x09, 0xc9, 0xc9, 0x27, 0xb0, 0xc1, 0x15, 0x1b, 0xa4, 0x3c, 0xa1, 0x43, 0x96,
	0xb3, 0x81, 0x48, 0x85, 0x11, 0xbc, 0xe8, 0x34, 0xac, 0xbf, 0x75, 0x6f, 0xfb, 0x72, 0xc6, 0x44,
	0x6e, 0x43, 0x23, 0x4f, 0x99, 0x19, 0x65, 0x5a, 0x76, 0x9a, 0x2e, 0xae, 0xfa, 0x37, 0xa9, 0x60,
	0xc5, 0x1f, 0x18, 0xb6, 0x82, 0xed, 0x70, 0xe7, 0xfb, 0x45, 0x25, 0x82, 0x03, 0x63, 0xef, 0x8d,
	0x48, 0x58, 0x92, 0x62, 0xd0, 0x09, 0xad, 0xd3, 0x6f, 0x17, 0xe5, 0x54, 0x8a, 0x41, 0x8c, 0x7e,
	0xba, 0xef, 0x43, 0x28, 0xd4, 0x43, 0xca, 0x92, 0x44, 0x53, 0x93, 0x90, 0x0d, 0xb8, 0x5a, 0xb1,
	0xb4, 0xac, 0x13, 0xcb, 0xfd, 0xe8, 0xfe, 0x11, 0x80, 0xcd, 0x54, 0x9a, 0x3e, 0xb0, 0x4c, 0xf2,
	0x1e, 0xb4, 0xf0, 0x93, 0x17, 0x05, 0x35, 0xc7, 0x79, 0xcd, 0x0e, 0x3d, 0xb6, 0x77, 0x9c, 0x73,
	0xa4, 0x88, 0xbc, 0xfa, 0x94, 0x7a, 0xcc, 0x27, 0x5a, 0x88, 0xd8, 0x23, 0x07, 0x91, 0x23, 0x4b,
	0x79, 0x38, 0xa5, 0x2c, 0xd9, 0x98, 0xf7, 0x2f, 0x2f, 0xe6, 0x99, 0xc8, 0xac, 0xe7, 0x87, 0xde,
	0x73, 0xf7, 0x6d, 0x00, 0x6b, 0x76, 0x3f, 0x6b, 0xe5, 0xca, 0xe8, 0x63, 0xf2, 0xfa, 0x2c, 0xd6,
	0x09, 0x6c, 0x0f, 0xb8, 0xec, 0xab, 0x3f, 0x71, 0xe0, 0xba, 0x80, 0x6d, 0x3c, 0x78, 0xb4, 0x27,
	0x88, 0x75, 0x7f, 0x0f, 0x60, 0xe3, 0x3c, 0x1e, 0xc9, 0xe1, 0x5a, 0x2d, 0x55, 0x60, 0xa5, 0x3a,
	0xb8, 0xe4, 0x83, 0xf9, 0xeb, 0x8d, 0x6b, 0x37, 0x58, 0xdc, 0x92, 0xd1, 0xa2, 0x1c, 0xd8, 0x5b,
	0xbe, 0x62, 0x2b, 0xae, 0x29, 0xd9, 0x0b, 0x07, 0x90, 0x1b, 0xb0, 0x22, 0x46, 0x54, 0x95, 0xd2,
	0x5e, 0x5d, 0x14, 0x5f, 0x15, 0xa3, 0xdd, 0x52, 0x76, 0x7f, 0x0b, 0xe0, 0xa6, 0xdd, 0xaf, 0x54,
	0x87, 0x2a, 0x7b, 0xa5, 0xa8, 0x49, 0x2b, 0x2f, 0xf2, 0xdb, 0x0b, 0x4d, 0x5e, 0x6b, 0x7e, 0xc9,
	0x21, 0x9d, 0xf1, 0xe3, 0x24, 0xb7, 0x63, 0x6d, 0xdf, 0xd9, 0xf6, 0xd2, 0xca, 0x09, 0xbf, 0x0f,
	0x9b, 0xff, 0xb2, 0x86, 0xdc, 0x82, 0x06, 0x22, 0xd3, 0x84, 0x8f, 0xe2, 0x6b, 0x26, 0xad, 0x6c,
	0xb2, 0x6f, 0x42, 0x13, 0x4d, 0xae, 0x74, 0x50, 0xa6, 0x56, 0x8c, 0xdc, 0x03, 0x5b, 0x3d, 0xbf,
	0xd6, 0x31, 0x67, 0x7a, 0x4c, 0x13, 0x3e, 0x9a, 0x91, 0xe3, 0x97, 0x0b, 0x4d, 0x0b, 0x92, 0xe3,
	0x8c, 0x1f, 0x27, 0x87, 0x1d, 0x3b, 0xdf, 0xe8, 0xf1, 0x63, 0x3e, 0x9a, 0xaa, 0xf1, 0x26, 0xf0,
	0x72, 0x9c, 0xbf, 0x86, 0xb4, 0x61, 0x29, 0x2b, 0x85, 0x57, 0x02, 0xbf, 0x92, 0x3b, 0x10, 0x22,
	0x67, 0x3e, 0x5d, 0xc0, 0xa4, 0x55, 0x9d, 0x2f, 0x77, 0x61, 0x15, 0x09, 0x42, 0x8d, 0x32, 0x2a,
	0x54, 0xc2, 0x0b, 0x9f, 0x37, 0x2d, 0x93, 0x56, 0x7d, 0x35, 0xca, 0xfa, 0x88, 0xcd, 0x8b, 0xb9,
	0x7c, 0x4a, 0xcc, 0xbf, 0x96, 0x7d, 0x71, 0x9c, 0xea, 0x9f, 0xe4, 0x43, 0x68, 0xdb, 0x41, 0x93,
	0xf0, 0x62, 0xa8, 0x45, 0x6e, 0x70, 0x88, 0xb8, 0xb6, 0xb4, 0x86, 0xf8, 0xe3, 0x13, 0x18, 0xcf,
	0x59, 0x1c, 0x17, 0x86, 0xcb, 0xd9, 0x11, 0x08, 0x0e, 0xb2, 0x03, 0xef, 0x3e, 0x10, 0x4f, 0x98,
	0xdd, 0xcd, 0x4d, 0xbe, 0xeb, 0xce, 0x32, 0xbb, 0xdf, 0x3d, 0x58, 0xc5, 0x89, 0x44, 0x35, 0x97,
	0x4c, 0x28, 0xa1, 0xc6, 0xf6, 0xd4, 0x51, 0x1c, 0x21, 0x1a, 0xd7, 0x20, 0xf9, 0x18, 0xd6, 0xfd,
	0xae, 0x73, 0xf3, 0xc9, 0xcd, 0x43, 0xef, 0x70, 0x6e, 0x3c, 0x5d, 0x34, 0xd1, 0x56, 0x2e, 0x9e,
	0x68, 0xaf, 0x03, 0xb8, 0xae, 0xb8, 0x79, 0x95, 0xe9, 0xc3, 0xba, 0xad, 0xf2, 0xc2, 0x8e, 0xca,
	0x70, 0xe7, 0xe5, 0xc2, 0xda, 0x58, 0xdc, 0xf6, 0x3e, 0x1f, 0xd5, 0x2e, 0xf1, 0x3a, 0x58, 0x69,
	0x32, 0xaa, 0xf8, 0x38, 0x33, 0x82, 0x59, 0x01, 0xdd, 0x24, 0x5e, 0x43, 0x7c, 0xf7, 0x04, 0x26,
	0x5f, 0xc0, 0x66, 0x3e, 0x39, 0x2e, 0xc4, 0x90, 0xa5, 0x54, 0xf2, 0x44, 0xb0, 0xf9, 0x68, 0xdd,
	0x60, 0xbe, 0x55, 0x53, 0x9e, 0x21, 0x63, 0x2e, 0xe6, 0xcf, 0xe1, 0xb6, 0x5b, 0xc6, 0x8c, 0x61,
	0xc3, 0x89, 0xc4, 0x57, 0x4e, 0xa9, 0x84, 0x71, 0x95, 0x0a, 0xf6, 0x2a, 0xfe, 0x67, 0x19, 0x8f,
	0xa6, 0x84, 0x7d, 0x25, 0x8c, 0xad, 0xdc, 0x2d, 0xb0, 0x2f, 0x11, 0x5a, 0xa5, 0x4c, 0xe1, 0x03,
	0x26, 0x74, 0x49, 0x8b, 0xd8, 0x41, 0xca, 0x54, 0x3f, 0xe9, 0xfe, 0xb9, 0x0c, 0xd7, 0xcf, 0x0c,
	0x4f, 0xd2, 0x85, 0x48, 0x73, 0x69, 0x5f, 0x22, 0x54, 0x32, 0x7d, 0xe8, 0xeb, 0x20, 0xd4, 0x5c,
	0xe2, 0x6b, 0xe4, 0x19, 0xd3, 0x87, 0xe4, 0x23, 0x20, 0xc8, 0x49, 0x33, 0x8c, 0xcc, 0x7a, 0xc1,
	0x56, 0xe9, 0xca, 0x62, 0x4d, 0x73, 0xf9, 0x14, 0x0d, 0xcf, 0x33, 0x6d, 0x76, 0x4b, 0x89, 0x59,
	0x8f, 0x64, 0x2c, 0x8b, 0x23, 0x5f, 0x16, 0x0d, 0xcd, 0x25, 0x96, 0xc4, 0x11, 0xb9, 0x0f, 0xeb,
	0x27, 0x8f, 0x2c, 0x2c, 0x30, 0x17, 0x9b, 0x4b, 0xb3, 0xf6, 0xf4, 0xb5, 0xf5, 0xa2, 0x1c, 0xec,
	0xf9, 0x3a, 0x9b, 0xa1, 0xa7, 0x5c, 0xd9, 0x24, 0x8b, 0xe2, 0xd6, 0x94, 0xf9, 0x94, 0x2b, 0xf2,
	0x81, 0xaf, 0x98, 0xd9, 0x1d, 0xfd, 0xb3, 0xcb, 0x3d, 0xce, 0xea, 0xed, 0xde, 0x85, 0xb0, 0x26,
	0xe2, 0x5e, 0xee, 0xe1, 0xd5, 0x74, 0x1c, 0xdc, 0xe8, 0x01, 0xdc, 0x18, 0x66, 0x72, 0x20, 0xd4,
	0x79, 0x4f, 0xaf, 0x28, 0xde, 0xa8, 0x8d, 0x73, 0xb7, 0xf6, 0x73, 0x00, 0xed, 0xd9, 0x46, 0x9b,
	0x8a, 0xc2, 0xd8, 0xbb, 0x0e, 0x77, 0x7e, 0x58, 0xf4, 0x0c, 0x88, 0x57, 0xcb, 0x69, 0xeb, 0x7f,
	0x2a, 0x0a, 0x43, 0x7e, 0x0a, 0xa0, 0x3d, 0xdb, 0xe6, 0xec, 0x61, 0x60, 0x21, 0x87, 0x39, 0xd3,
	0x4d, 0xe3, 0x28, 0xab, 0x1b, 0x2f, 0x9e, 0x65, 0xb0, 0x62, 0xff, 0xc7, 0x7a, 0xf0, 0x4f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xb6, 0x90, 0xec, 0xe6, 0x7a, 0x0d, 0x00, 0x00,
}
