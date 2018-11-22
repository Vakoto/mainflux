// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/brocaar/loraserver/api/common"
import gw "github.com/brocaar/loraserver/api/gw"
import duration "github.com/golang/protobuf/ptypes/duration"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RXWindow int32

const (
	RXWindow_RX1 RXWindow = 0
	RXWindow_RX2 RXWindow = 1
)

var RXWindow_name = map[int32]string{
	0: "RX1",
	1: "RX2",
}

var RXWindow_value = map[string]int32{
	"RX1": 0,
	"RX2": 1,
}

func (x RXWindow) String() string {
	return proto.EnumName(RXWindow_name, int32(x))
}

func (RXWindow) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

type UplinkFrameLog struct {
	// TX information of the uplink.
	TxInfo *gw.UplinkTXInfo `protobuf:"bytes,1,opt,name=tx_info,json=txInfo,proto3" json:"tx_info,omitempty"`
	// RX information of the uplink.
	RxInfo []*UplinkRXInfo `protobuf:"bytes,2,rep,name=rx_info,json=rxInfo,proto3" json:"rx_info,omitempty"`
	// LoRaWAN PHYPayload.
	PhyPayloadJson       string   `protobuf:"bytes,3,opt,name=phy_payload_json,json=phyPayloadJSON,proto3" json:"phy_payload_json,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UplinkFrameLog) Reset()         { *m = UplinkFrameLog{} }
func (m *UplinkFrameLog) String() string { return proto.CompactTextString(m) }
func (*UplinkFrameLog) ProtoMessage()    {}
func (*UplinkFrameLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}
func (m *UplinkFrameLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UplinkFrameLog.Unmarshal(m, b)
}
func (m *UplinkFrameLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UplinkFrameLog.Marshal(b, m, deterministic)
}
func (dst *UplinkFrameLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UplinkFrameLog.Merge(dst, src)
}
func (m *UplinkFrameLog) XXX_Size() int {
	return xxx_messageInfo_UplinkFrameLog.Size(m)
}
func (m *UplinkFrameLog) XXX_DiscardUnknown() {
	xxx_messageInfo_UplinkFrameLog.DiscardUnknown(m)
}

var xxx_messageInfo_UplinkFrameLog proto.InternalMessageInfo

func (m *UplinkFrameLog) GetTxInfo() *gw.UplinkTXInfo {
	if m != nil {
		return m.TxInfo
	}
	return nil
}

func (m *UplinkFrameLog) GetRxInfo() []*UplinkRXInfo {
	if m != nil {
		return m.RxInfo
	}
	return nil
}

func (m *UplinkFrameLog) GetPhyPayloadJson() string {
	if m != nil {
		return m.PhyPayloadJson
	}
	return ""
}

type DownlinkFrameLog struct {
	// TX information of the downlink.
	TxInfo *DownlinkTXInfo `protobuf:"bytes,1,opt,name=tx_info,json=txInfo,proto3" json:"tx_info,omitempty"`
	// LoRaWAN PHYPayload.
	PhyPayloadJson       string   `protobuf:"bytes,2,opt,name=phy_payload_json,json=phyPayloadJSON,proto3" json:"phy_payload_json,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownlinkFrameLog) Reset()         { *m = DownlinkFrameLog{} }
func (m *DownlinkFrameLog) String() string { return proto.CompactTextString(m) }
func (*DownlinkFrameLog) ProtoMessage()    {}
func (*DownlinkFrameLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}
func (m *DownlinkFrameLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownlinkFrameLog.Unmarshal(m, b)
}
func (m *DownlinkFrameLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownlinkFrameLog.Marshal(b, m, deterministic)
}
func (dst *DownlinkFrameLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownlinkFrameLog.Merge(dst, src)
}
func (m *DownlinkFrameLog) XXX_Size() int {
	return xxx_messageInfo_DownlinkFrameLog.Size(m)
}
func (m *DownlinkFrameLog) XXX_DiscardUnknown() {
	xxx_messageInfo_DownlinkFrameLog.DiscardUnknown(m)
}

var xxx_messageInfo_DownlinkFrameLog proto.InternalMessageInfo

func (m *DownlinkFrameLog) GetTxInfo() *DownlinkTXInfo {
	if m != nil {
		return m.TxInfo
	}
	return nil
}

func (m *DownlinkFrameLog) GetPhyPayloadJson() string {
	if m != nil {
		return m.PhyPayloadJson
	}
	return ""
}

// This is a copy of gw.UplinkRXInfo with the only change that the
// gateway_id is of type string so that we can return it as HEX encoded
// instead of base64.
type UplinkRXInfo struct {
	// Gateway ID.
	GatewayId string `protobuf:"bytes,1,opt,name=gateway_id,json=gatewayId,proto3" json:"gateway_id,omitempty"`
	// RX time (only set when the gateway has a GPS module).
	Time *timestamp.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	// RX time since GPS epoch (only set when the gateway has a GPS module).
	TimeSinceGpsEpoch *duration.Duration `protobuf:"bytes,3,opt,name=time_since_gps_epoch,json=timeSinceGpsEpoch,proto3" json:"time_since_gps_epoch,omitempty"`
	// Gateway internal timestamp.
	Timestamp uint32 `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// RSSI.
	Rssi int32 `protobuf:"varint,5,opt,name=rssi,proto3" json:"rssi,omitempty"`
	// LoRa SNR.
	LoraSnr float64 `protobuf:"fixed64,6,opt,name=lora_snr,json=loraSnr,proto3" json:"lora_snr,omitempty"`
	// Channel.
	Channel uint32 `protobuf:"varint,7,opt,name=channel,proto3" json:"channel,omitempty"`
	// RF Chain.
	RfChain uint32 `protobuf:"varint,8,opt,name=rf_chain,json=rfChain,proto3" json:"rf_chain,omitempty"`
	// Board.
	Board uint32 `protobuf:"varint,9,opt,name=board,proto3" json:"board,omitempty"`
	// Antenna.
	Antenna uint32 `protobuf:"varint,10,opt,name=antenna,proto3" json:"antenna,omitempty"`
	// Location.
	Location *common.Location `protobuf:"bytes,11,opt,name=location,proto3" json:"location,omitempty"`
	// Fine-timestamp type.
	FineTimestampType gw.FineTimestampType `protobuf:"varint,12,opt,name=fine_timestamp_type,json=fineTimestampType,proto3,enum=gw.FineTimestampType" json:"fine_timestamp_type,omitempty"`
	// Fine-timestamp data.
	//
	// Types that are valid to be assigned to FineTimestamp:
	//	*UplinkRXInfo_EncryptedFineTimestamp
	//	*UplinkRXInfo_PlainFineTimestamp
	FineTimestamp        isUplinkRXInfo_FineTimestamp `protobuf_oneof:"fine_timestamp"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *UplinkRXInfo) Reset()         { *m = UplinkRXInfo{} }
func (m *UplinkRXInfo) String() string { return proto.CompactTextString(m) }
func (*UplinkRXInfo) ProtoMessage()    {}
func (*UplinkRXInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}
func (m *UplinkRXInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UplinkRXInfo.Unmarshal(m, b)
}
func (m *UplinkRXInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UplinkRXInfo.Marshal(b, m, deterministic)
}
func (dst *UplinkRXInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UplinkRXInfo.Merge(dst, src)
}
func (m *UplinkRXInfo) XXX_Size() int {
	return xxx_messageInfo_UplinkRXInfo.Size(m)
}
func (m *UplinkRXInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UplinkRXInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UplinkRXInfo proto.InternalMessageInfo

func (m *UplinkRXInfo) GetGatewayId() string {
	if m != nil {
		return m.GatewayId
	}
	return ""
}

func (m *UplinkRXInfo) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *UplinkRXInfo) GetTimeSinceGpsEpoch() *duration.Duration {
	if m != nil {
		return m.TimeSinceGpsEpoch
	}
	return nil
}

func (m *UplinkRXInfo) GetTimestamp() uint32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *UplinkRXInfo) GetRssi() int32 {
	if m != nil {
		return m.Rssi
	}
	return 0
}

func (m *UplinkRXInfo) GetLoraSnr() float64 {
	if m != nil {
		return m.LoraSnr
	}
	return 0
}

func (m *UplinkRXInfo) GetChannel() uint32 {
	if m != nil {
		return m.Channel
	}
	return 0
}

func (m *UplinkRXInfo) GetRfChain() uint32 {
	if m != nil {
		return m.RfChain
	}
	return 0
}

func (m *UplinkRXInfo) GetBoard() uint32 {
	if m != nil {
		return m.Board
	}
	return 0
}

func (m *UplinkRXInfo) GetAntenna() uint32 {
	if m != nil {
		return m.Antenna
	}
	return 0
}

func (m *UplinkRXInfo) GetLocation() *common.Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *UplinkRXInfo) GetFineTimestampType() gw.FineTimestampType {
	if m != nil {
		return m.FineTimestampType
	}
	return gw.FineTimestampType_NONE
}

type isUplinkRXInfo_FineTimestamp interface {
	isUplinkRXInfo_FineTimestamp()
}

type UplinkRXInfo_EncryptedFineTimestamp struct {
	EncryptedFineTimestamp *EncryptedFineTimestamp `protobuf:"bytes,13,opt,name=encrypted_fine_timestamp,json=encryptedFineTimestamp,proto3,oneof"`
}

type UplinkRXInfo_PlainFineTimestamp struct {
	PlainFineTimestamp *gw.PlainFineTimestamp `protobuf:"bytes,14,opt,name=plain_fine_timestamp,json=plainFineTimestamp,proto3,oneof"`
}

func (*UplinkRXInfo_EncryptedFineTimestamp) isUplinkRXInfo_FineTimestamp() {}

func (*UplinkRXInfo_PlainFineTimestamp) isUplinkRXInfo_FineTimestamp() {}

func (m *UplinkRXInfo) GetFineTimestamp() isUplinkRXInfo_FineTimestamp {
	if m != nil {
		return m.FineTimestamp
	}
	return nil
}

func (m *UplinkRXInfo) GetEncryptedFineTimestamp() *EncryptedFineTimestamp {
	if x, ok := m.GetFineTimestamp().(*UplinkRXInfo_EncryptedFineTimestamp); ok {
		return x.EncryptedFineTimestamp
	}
	return nil
}

func (m *UplinkRXInfo) GetPlainFineTimestamp() *gw.PlainFineTimestamp {
	if x, ok := m.GetFineTimestamp().(*UplinkRXInfo_PlainFineTimestamp); ok {
		return x.PlainFineTimestamp
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*UplinkRXInfo) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _UplinkRXInfo_OneofMarshaler, _UplinkRXInfo_OneofUnmarshaler, _UplinkRXInfo_OneofSizer, []interface{}{
		(*UplinkRXInfo_EncryptedFineTimestamp)(nil),
		(*UplinkRXInfo_PlainFineTimestamp)(nil),
	}
}

func _UplinkRXInfo_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*UplinkRXInfo)
	// fine_timestamp
	switch x := m.FineTimestamp.(type) {
	case *UplinkRXInfo_EncryptedFineTimestamp:
		b.EncodeVarint(13<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.EncryptedFineTimestamp); err != nil {
			return err
		}
	case *UplinkRXInfo_PlainFineTimestamp:
		b.EncodeVarint(14<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PlainFineTimestamp); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("UplinkRXInfo.FineTimestamp has unexpected type %T", x)
	}
	return nil
}

func _UplinkRXInfo_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*UplinkRXInfo)
	switch tag {
	case 13: // fine_timestamp.encrypted_fine_timestamp
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EncryptedFineTimestamp)
		err := b.DecodeMessage(msg)
		m.FineTimestamp = &UplinkRXInfo_EncryptedFineTimestamp{msg}
		return true, err
	case 14: // fine_timestamp.plain_fine_timestamp
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(gw.PlainFineTimestamp)
		err := b.DecodeMessage(msg)
		m.FineTimestamp = &UplinkRXInfo_PlainFineTimestamp{msg}
		return true, err
	default:
		return false, nil
	}
}

func _UplinkRXInfo_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*UplinkRXInfo)
	// fine_timestamp
	switch x := m.FineTimestamp.(type) {
	case *UplinkRXInfo_EncryptedFineTimestamp:
		s := proto.Size(x.EncryptedFineTimestamp)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UplinkRXInfo_PlainFineTimestamp:
		s := proto.Size(x.PlainFineTimestamp)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// this s a copy of gw.EncryptedFineTimestamp which the only change that
// the fpga_id is of type string so that it can be returned in HEX format
// instead of base64.
type EncryptedFineTimestamp struct {
	// AES key index used for encrypting the fine timestamp.
	AesKeyIndex uint32 `protobuf:"varint,1,opt,name=aes_key_index,json=aesKeyIndex,proto3" json:"aes_key_index,omitempty"`
	// Encrypted 'main' fine-timestamp (ns precision part of the timestamp).
	EncryptedNs []byte `protobuf:"bytes,2,opt,name=encrypted_ns,json=encryptedNS,proto3" json:"encrypted_ns,omitempty"`
	// FPGA ID.
	FpgaId               string   `protobuf:"bytes,3,opt,name=fpga_id,json=fpgaID,proto3" json:"fpga_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EncryptedFineTimestamp) Reset()         { *m = EncryptedFineTimestamp{} }
func (m *EncryptedFineTimestamp) String() string { return proto.CompactTextString(m) }
func (*EncryptedFineTimestamp) ProtoMessage()    {}
func (*EncryptedFineTimestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{3}
}
func (m *EncryptedFineTimestamp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncryptedFineTimestamp.Unmarshal(m, b)
}
func (m *EncryptedFineTimestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncryptedFineTimestamp.Marshal(b, m, deterministic)
}
func (dst *EncryptedFineTimestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptedFineTimestamp.Merge(dst, src)
}
func (m *EncryptedFineTimestamp) XXX_Size() int {
	return xxx_messageInfo_EncryptedFineTimestamp.Size(m)
}
func (m *EncryptedFineTimestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptedFineTimestamp.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptedFineTimestamp proto.InternalMessageInfo

func (m *EncryptedFineTimestamp) GetAesKeyIndex() uint32 {
	if m != nil {
		return m.AesKeyIndex
	}
	return 0
}

func (m *EncryptedFineTimestamp) GetEncryptedNs() []byte {
	if m != nil {
		return m.EncryptedNs
	}
	return nil
}

func (m *EncryptedFineTimestamp) GetFpgaId() string {
	if m != nil {
		return m.FpgaId
	}
	return ""
}

// Same comment as above applies to this message.
type DownlinkTXInfo struct {
	// Gateway ID.
	GatewayId string `protobuf:"bytes,1,opt,name=gateway_id,json=gatewayId,proto3" json:"gateway_id,omitempty"`
	// Frame must be sent immediately.
	Immediately bool `protobuf:"varint,2,opt,name=immediately,proto3" json:"immediately,omitempty"`
	// Emit frame at the given time since GPS epoch.
	TimeSinceGpsEpoch *duration.Duration `protobuf:"bytes,3,opt,name=time_since_gps_epoch,json=timeSinceGpsEpoch,proto3" json:"time_since_gps_epoch,omitempty"`
	// Emit the frame at the given gateway internal timestamp.
	Timestamp uint32 `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// TX frequency (in Hz).
	Frequency uint32 `protobuf:"varint,5,opt,name=frequency,proto3" json:"frequency,omitempty"`
	// TX power (in dBm).
	Power int32 `protobuf:"varint,6,opt,name=power,proto3" json:"power,omitempty"`
	// Modulation.
	Modulation common.Modulation `protobuf:"varint,7,opt,name=modulation,proto3,enum=common.Modulation" json:"modulation,omitempty"`
	// Types that are valid to be assigned to ModulationInfo:
	//	*DownlinkTXInfo_LoraModulationInfo
	//	*DownlinkTXInfo_FskModulationInfo
	ModulationInfo isDownlinkTXInfo_ModulationInfo `protobuf_oneof:"modulation_info"`
	// The board identifier for emitting the frame.
	Board uint32 `protobuf:"varint,10,opt,name=board,proto3" json:"board,omitempty"`
	// The antenna identifier for emitting the frame.
	Antenna              uint32   `protobuf:"varint,11,opt,name=antenna,proto3" json:"antenna,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownlinkTXInfo) Reset()         { *m = DownlinkTXInfo{} }
func (m *DownlinkTXInfo) String() string { return proto.CompactTextString(m) }
func (*DownlinkTXInfo) ProtoMessage()    {}
func (*DownlinkTXInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{4}
}
func (m *DownlinkTXInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownlinkTXInfo.Unmarshal(m, b)
}
func (m *DownlinkTXInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownlinkTXInfo.Marshal(b, m, deterministic)
}
func (dst *DownlinkTXInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownlinkTXInfo.Merge(dst, src)
}
func (m *DownlinkTXInfo) XXX_Size() int {
	return xxx_messageInfo_DownlinkTXInfo.Size(m)
}
func (m *DownlinkTXInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DownlinkTXInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DownlinkTXInfo proto.InternalMessageInfo

func (m *DownlinkTXInfo) GetGatewayId() string {
	if m != nil {
		return m.GatewayId
	}
	return ""
}

func (m *DownlinkTXInfo) GetImmediately() bool {
	if m != nil {
		return m.Immediately
	}
	return false
}

func (m *DownlinkTXInfo) GetTimeSinceGpsEpoch() *duration.Duration {
	if m != nil {
		return m.TimeSinceGpsEpoch
	}
	return nil
}

func (m *DownlinkTXInfo) GetTimestamp() uint32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *DownlinkTXInfo) GetFrequency() uint32 {
	if m != nil {
		return m.Frequency
	}
	return 0
}

func (m *DownlinkTXInfo) GetPower() int32 {
	if m != nil {
		return m.Power
	}
	return 0
}

func (m *DownlinkTXInfo) GetModulation() common.Modulation {
	if m != nil {
		return m.Modulation
	}
	return common.Modulation_LORA
}

type isDownlinkTXInfo_ModulationInfo interface {
	isDownlinkTXInfo_ModulationInfo()
}

type DownlinkTXInfo_LoraModulationInfo struct {
	LoraModulationInfo *gw.LoRaModulationInfo `protobuf:"bytes,8,opt,name=lora_modulation_info,json=loraModulationInfo,proto3,oneof"`
}

type DownlinkTXInfo_FskModulationInfo struct {
	FskModulationInfo *gw.FSKModulationInfo `protobuf:"bytes,9,opt,name=fsk_modulation_info,json=fskModulationInfo,proto3,oneof"`
}

func (*DownlinkTXInfo_LoraModulationInfo) isDownlinkTXInfo_ModulationInfo() {}

func (*DownlinkTXInfo_FskModulationInfo) isDownlinkTXInfo_ModulationInfo() {}

func (m *DownlinkTXInfo) GetModulationInfo() isDownlinkTXInfo_ModulationInfo {
	if m != nil {
		return m.ModulationInfo
	}
	return nil
}

func (m *DownlinkTXInfo) GetLoraModulationInfo() *gw.LoRaModulationInfo {
	if x, ok := m.GetModulationInfo().(*DownlinkTXInfo_LoraModulationInfo); ok {
		return x.LoraModulationInfo
	}
	return nil
}

func (m *DownlinkTXInfo) GetFskModulationInfo() *gw.FSKModulationInfo {
	if x, ok := m.GetModulationInfo().(*DownlinkTXInfo_FskModulationInfo); ok {
		return x.FskModulationInfo
	}
	return nil
}

func (m *DownlinkTXInfo) GetBoard() uint32 {
	if m != nil {
		return m.Board
	}
	return 0
}

func (m *DownlinkTXInfo) GetAntenna() uint32 {
	if m != nil {
		return m.Antenna
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DownlinkTXInfo) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _DownlinkTXInfo_OneofMarshaler, _DownlinkTXInfo_OneofUnmarshaler, _DownlinkTXInfo_OneofSizer, []interface{}{
		(*DownlinkTXInfo_LoraModulationInfo)(nil),
		(*DownlinkTXInfo_FskModulationInfo)(nil),
	}
}

func _DownlinkTXInfo_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DownlinkTXInfo)
	// modulation_info
	switch x := m.ModulationInfo.(type) {
	case *DownlinkTXInfo_LoraModulationInfo:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LoraModulationInfo); err != nil {
			return err
		}
	case *DownlinkTXInfo_FskModulationInfo:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FskModulationInfo); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("DownlinkTXInfo.ModulationInfo has unexpected type %T", x)
	}
	return nil
}

func _DownlinkTXInfo_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DownlinkTXInfo)
	switch tag {
	case 8: // modulation_info.lora_modulation_info
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(gw.LoRaModulationInfo)
		err := b.DecodeMessage(msg)
		m.ModulationInfo = &DownlinkTXInfo_LoraModulationInfo{msg}
		return true, err
	case 9: // modulation_info.fsk_modulation_info
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(gw.FSKModulationInfo)
		err := b.DecodeMessage(msg)
		m.ModulationInfo = &DownlinkTXInfo_FskModulationInfo{msg}
		return true, err
	default:
		return false, nil
	}
}

func _DownlinkTXInfo_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*DownlinkTXInfo)
	// modulation_info
	switch x := m.ModulationInfo.(type) {
	case *DownlinkTXInfo_LoraModulationInfo:
		s := proto.Size(x.LoraModulationInfo)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *DownlinkTXInfo_FskModulationInfo:
		s := proto.Size(x.FskModulationInfo)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*UplinkFrameLog)(nil), "api.UplinkFrameLog")
	proto.RegisterType((*DownlinkFrameLog)(nil), "api.DownlinkFrameLog")
	proto.RegisterType((*UplinkRXInfo)(nil), "api.UplinkRXInfo")
	proto.RegisterType((*EncryptedFineTimestamp)(nil), "api.EncryptedFineTimestamp")
	proto.RegisterType((*DownlinkTXInfo)(nil), "api.DownlinkTXInfo")
	proto.RegisterEnum("api.RXWindow", RXWindow_name, RXWindow_value)
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 795 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x51, 0x73, 0xda, 0x46,
	0x10, 0x8e, 0x82, 0x31, 0xb0, 0x02, 0x0a, 0xe7, 0xd4, 0x55, 0xa8, 0xdb, 0x52, 0x9e, 0x68, 0x26,
	0x15, 0x53, 0x3a, 0xfd, 0x03, 0xa9, 0x9d, 0xd4, 0x89, 0x9b, 0x66, 0x0e, 0x77, 0xe2, 0x37, 0xcd,
	0x21, 0x9d, 0xc4, 0xc5, 0xe8, 0xee, 0x7a, 0x12, 0xc5, 0xfa, 0x05, 0x7d, 0xee, 0xef, 0xed, 0x4b,
	0xe7, 0xee, 0x04, 0x18, 0x99, 0x8e, 0xdf, 0xfa, 0x24, 0xed, 0x77, 0xdf, 0x7e, 0xbb, 0xa7, 0xfd,
	0x56, 0xd0, 0x0e, 0x45, 0x9a, 0x0a, 0xee, 0x4b, 0x25, 0x72, 0x81, 0x6a, 0x44, 0xb2, 0xc1, 0x37,
	0x89, 0x10, 0xc9, 0x92, 0x4e, 0x0c, 0x34, 0x5f, 0xc5, 0x93, 0x9c, 0xa5, 0x34, 0xcb, 0x49, 0x2a,
	0x2d, 0x6b, 0xf0, 0x75, 0x95, 0x10, 0xad, 0x14, 0xc9, 0xd9, 0x46, 0x65, 0xf0, 0x53, 0xc2, 0xf2,
	0xc5, 0x6a, 0xee, 0x87, 0x22, 0x9d, 0xcc, 0x95, 0x08, 0x09, 0x51, 0x93, 0xa5, 0x50, 0x24, 0xa3,
	0xea, 0x4f, 0xaa, 0x26, 0x44, 0xb2, 0x89, 0xad, 0x3a, 0xb9, 0x5f, 0x7c, 0xf0, 0xfd, 0xe3, 0x69,
	0xc9, 0x7a, 0x92, 0xac, 0x2d, 0x7d, 0xf4, 0xb7, 0x03, 0xdd, 0xdf, 0xe5, 0x92, 0xf1, 0xdb, 0xd7,
	0x8a, 0xa4, 0xf4, 0x4a, 0x24, 0xe8, 0x3b, 0x68, 0xe4, 0x77, 0x01, 0xe3, 0xb1, 0xf0, 0x9c, 0xa1,
	0x33, 0x76, 0xa7, 0x3d, 0x3f, 0x59, 0xfb, 0x96, 0x74, 0x7d, 0x73, 0xc9, 0x63, 0x81, 0x8f, 0xf3,
	0x3b, 0xfd, 0x44, 0x2f, 0xa0, 0xa1, 0x4a, 0xea, 0xd3, 0x61, 0x6d, 0xec, 0x4e, 0xfb, 0x3e, 0x91,
	0xac, 0xe4, 0xe2, 0x92, 0xab, 0x2c, 0x77, 0x0c, 0x3d, 0xb9, 0x28, 0x02, 0x49, 0x8a, 0xa5, 0x20,
	0x51, 0xf0, 0x29, 0x13, 0xdc, 0xab, 0x0d, 0x9d, 0x71, 0x0b, 0x77, 0xe5, 0xa2, 0xf8, 0x60, 0xe1,
	0xb7, 0xb3, 0xdf, 0xde, 0x8f, 0x3e, 0x41, 0xef, 0x5c, 0xac, 0xf9, 0x5e, 0x53, 0x2f, 0xab, 0x4d,
	0x9d, 0x98, 0x4a, 0x1b, 0x5e, 0xa5, 0xaf, 0x43, 0xb5, 0x9e, 0x1e, 0xac, 0xf5, 0x57, 0x1d, 0xda,
	0xf7, 0xdb, 0x45, 0x5f, 0x01, 0x24, 0x24, 0xa7, 0x6b, 0x52, 0x04, 0x2c, 0x32, 0xb5, 0x5a, 0xb8,
	0x55, 0x22, 0x97, 0x11, 0xf2, 0xe1, 0x48, 0x0f, 0xd2, 0xa8, 0xb9, 0xd3, 0x81, 0x6f, 0x87, 0xe8,
	0x6f, 0x86, 0xe8, 0x5f, 0x6f, 0xa6, 0x8c, 0x0d, 0x0f, 0xbd, 0x85, 0x67, 0xfa, 0x19, 0x64, 0x8c,
	0x87, 0x34, 0x48, 0x64, 0x16, 0x50, 0x29, 0xc2, 0x85, 0xb9, 0xb9, 0x3b, 0x7d, 0xfe, 0x20, 0xff,
	0xbc, 0x34, 0x01, 0xee, 0xeb, 0xb4, 0x99, 0xce, 0x7a, 0x23, 0xb3, 0x0b, 0x9d, 0x83, 0xce, 0xa0,
	0xb5, 0x35, 0x91, 0x77, 0x34, 0x74, 0xc6, 0x1d, 0xbc, 0x03, 0x10, 0x82, 0x23, 0x95, 0x65, 0xcc,
	0xab, 0x0f, 0x9d, 0x71, 0x1d, 0x9b, 0x77, 0xf4, 0x1c, 0x9a, 0x7a, 0xf6, 0x41, 0xc6, 0x95, 0x77,
	0x3c, 0x74, 0xc6, 0x0e, 0x6e, 0xe8, 0x78, 0xc6, 0x15, 0xf2, 0xa0, 0x11, 0x2e, 0x08, 0xe7, 0x74,
	0xe9, 0x35, 0x8c, 0xd4, 0x26, 0xd4, 0x49, 0x2a, 0x0e, 0xc2, 0x05, 0x61, 0xdc, 0x6b, 0xda, 0x23,
	0x15, 0xff, 0xac, 0x43, 0xf4, 0x0c, 0xea, 0x73, 0x41, 0x54, 0xe4, 0xb5, 0x0c, 0x6e, 0x03, 0x2d,
	0x45, 0x78, 0x4e, 0x39, 0x27, 0x1e, 0x58, 0x7e, 0x19, 0xa2, 0x97, 0xba, 0x7e, 0x68, 0x2e, 0xe4,
	0xb9, 0xa5, 0x97, 0x4a, 0xb7, 0x5e, 0x95, 0x38, 0xde, 0x32, 0xd0, 0x05, 0x9c, 0xc4, 0x8c, 0xd3,
	0x60, 0x7b, 0xa7, 0x20, 0x2f, 0x24, 0xf5, 0xda, 0x43, 0x67, 0xdc, 0x9d, 0x7e, 0xae, 0x4d, 0xf8,
	0x9a, 0x71, 0xba, 0xfd, 0xc2, 0xd7, 0x85, 0xa4, 0xb8, 0x1f, 0x57, 0x21, 0xf4, 0x11, 0x3c, 0xca,
	0x43, 0x55, 0xc8, 0x9c, 0x46, 0xc1, 0xbe, 0xa0, 0xd7, 0x31, 0x4d, 0x7c, 0x69, 0xbc, 0x73, 0xb1,
	0x21, 0xed, 0xa9, 0xfe, 0xf2, 0x04, 0x9f, 0xd2, 0x83, 0x27, 0x7a, 0x96, 0x72, 0x49, 0x18, 0xaf,
	0x8a, 0x76, 0x8d, 0xe8, 0xa9, 0x6e, 0xf0, 0x83, 0x3e, 0xaf, 0xea, 0x21, 0xf9, 0x00, 0x7d, 0xd5,
	0x83, 0xee, 0xbe, 0xca, 0xe8, 0x0e, 0x4e, 0x0f, 0x77, 0x84, 0x46, 0xd0, 0x21, 0x34, 0x0b, 0x6e,
	0x69, 0x11, 0x30, 0x1e, 0xd1, 0x3b, 0xe3, 0xca, 0x0e, 0x76, 0x09, 0xcd, 0xde, 0xd1, 0xe2, 0x52,
	0x43, 0xe8, 0x5b, 0x68, 0xef, 0x2e, 0xcd, 0x33, 0xe3, 0xcf, 0x36, 0x76, 0xb7, 0xd8, 0xfb, 0x19,
	0xfa, 0x02, 0x1a, 0xb1, 0x4c, 0x88, 0xb6, 0xb5, 0xdd, 0xbb, 0x63, 0x1d, 0x5e, 0x9e, 0x8f, 0xfe,
	0xa9, 0x41, 0x77, 0x7f, 0x91, 0x1e, 0xdb, 0x82, 0x21, 0xb8, 0x2c, 0x4d, 0x69, 0xc4, 0x48, 0x4e,
	0x97, 0x85, 0x29, 0xd6, 0xc4, 0xf7, 0xa1, 0xff, 0xd1, 0xf7, 0x67, 0xd0, 0x8a, 0x15, 0xfd, 0x63,
	0x45, 0x79, 0x58, 0x18, 0xf3, 0x77, 0xf0, 0x0e, 0xd0, 0x8e, 0x95, 0x62, 0x4d, 0xad, 0xfd, 0xeb,
	0xd8, 0x06, 0x68, 0x0a, 0x90, 0x8a, 0x68, 0xb5, 0xb4, 0xce, 0x6c, 0x18, 0x83, 0xa1, 0x8d, 0x33,
	0x7f, 0xdd, 0x9e, 0xe0, 0x7b, 0x2c, 0x7d, 0x23, 0xb3, 0x4b, 0x3b, 0xc8, 0xfe, 0x8e, 0x9a, 0xbb,
	0xe9, 0x5f, 0x09, 0x4c, 0x76, 0xd9, 0xfa, 0x43, 0xea, 0xe9, 0xeb, 0xac, 0x7d, 0x14, 0xbd, 0x81,
	0x93, 0x38, 0xbb, 0x7d, 0x20, 0xd5, 0x32, 0x52, 0xd6, 0xe9, 0xb3, 0x77, 0x0f, 0x94, 0xfa, 0x71,
	0x76, 0x5b, 0x11, 0xda, 0x2e, 0x24, 0xfc, 0xc7, 0x42, 0xba, 0x7b, 0x0b, 0xf9, 0xaa, 0x0f, 0x9f,
	0x55, 0x8a, 0xbe, 0x38, 0x83, 0x26, 0xbe, 0xf9, 0xc8, 0x78, 0x24, 0xd6, 0xa8, 0x01, 0x35, 0x7c,
	0xf3, 0x43, 0xef, 0x89, 0x7d, 0x99, 0xf6, 0x9c, 0xf9, 0xb1, 0x99, 0xd0, 0x8f, 0xff, 0x06, 0x00,
	0x00, 0xff, 0xff, 0x1b, 0x68, 0xe9, 0x51, 0xe2, 0x06, 0x00, 0x00,
}