// Code generated by protoc-gen-go. DO NOT EDIT.
// source: machines/machinespb/machines.proto

package machinespb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Socket struct {
	SocketId string `protobuf:"bytes,1,opt,name=socket_id,json=socketId,proto3" json:"socket_id,omitempty"`
	//string machine_id = 2;
	NumCores int32 `protobuf:"varint,3,opt,name=num_cores,json=numCores,proto3" json:"num_cores,omitempty"`
	//repeated Core cores = 4;
	L3CacheSize          int64    `protobuf:"varint,5,opt,name=l3_cache_size,json=l3CacheSize,proto3" json:"l3_cache_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Socket) Reset()         { *m = Socket{} }
func (m *Socket) String() string { return proto.CompactTextString(m) }
func (*Socket) ProtoMessage()    {}
func (*Socket) Descriptor() ([]byte, []int) {
	return fileDescriptor_e909cb39bc20a1cf, []int{0}
}

func (m *Socket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Socket.Unmarshal(m, b)
}
func (m *Socket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Socket.Marshal(b, m, deterministic)
}
func (m *Socket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Socket.Merge(m, src)
}
func (m *Socket) XXX_Size() int {
	return xxx_messageInfo_Socket.Size(m)
}
func (m *Socket) XXX_DiscardUnknown() {
	xxx_messageInfo_Socket.DiscardUnknown(m)
}

var xxx_messageInfo_Socket proto.InternalMessageInfo

func (m *Socket) GetSocketId() string {
	if m != nil {
		return m.SocketId
	}
	return ""
}

func (m *Socket) GetNumCores() int32 {
	if m != nil {
		return m.NumCores
	}
	return 0
}

func (m *Socket) GetL3CacheSize() int64 {
	if m != nil {
		return m.L3CacheSize
	}
	return 0
}

type HardwareThread struct {
	HardwareThreadId     string   `protobuf:"bytes,1,opt,name=hardware_thread_id,json=hardwareThreadId,proto3" json:"hardware_thread_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HardwareThread) Reset()         { *m = HardwareThread{} }
func (m *HardwareThread) String() string { return proto.CompactTextString(m) }
func (*HardwareThread) ProtoMessage()    {}
func (*HardwareThread) Descriptor() ([]byte, []int) {
	return fileDescriptor_e909cb39bc20a1cf, []int{1}
}

func (m *HardwareThread) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HardwareThread.Unmarshal(m, b)
}
func (m *HardwareThread) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HardwareThread.Marshal(b, m, deterministic)
}
func (m *HardwareThread) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HardwareThread.Merge(m, src)
}
func (m *HardwareThread) XXX_Size() int {
	return xxx_messageInfo_HardwareThread.Size(m)
}
func (m *HardwareThread) XXX_DiscardUnknown() {
	xxx_messageInfo_HardwareThread.DiscardUnknown(m)
}

var xxx_messageInfo_HardwareThread proto.InternalMessageInfo

func (m *HardwareThread) GetHardwareThreadId() string {
	if m != nil {
		return m.HardwareThreadId
	}
	return ""
}

type Core struct {
	CoreId               string            `protobuf:"bytes,1,opt,name=core_id,json=coreId,proto3" json:"core_id,omitempty"`
	IsHyperThread        bool              `protobuf:"varint,2,opt,name=is_hyper_thread,json=isHyperThread,proto3" json:"is_hyper_thread,omitempty"`
	Hws                  []*HardwareThread `protobuf:"bytes,3,rep,name=hws,proto3" json:"hws,omitempty"`
	L2CacheSize          int64             `protobuf:"varint,4,opt,name=l2_cache_size,json=l2CacheSize,proto3" json:"l2_cache_size,omitempty"`
	L1CacheSize          int64             `protobuf:"varint,5,opt,name=l1_cache_size,json=l1CacheSize,proto3" json:"l1_cache_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Core) Reset()         { *m = Core{} }
func (m *Core) String() string { return proto.CompactTextString(m) }
func (*Core) ProtoMessage()    {}
func (*Core) Descriptor() ([]byte, []int) {
	return fileDescriptor_e909cb39bc20a1cf, []int{2}
}

func (m *Core) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Core.Unmarshal(m, b)
}
func (m *Core) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Core.Marshal(b, m, deterministic)
}
func (m *Core) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Core.Merge(m, src)
}
func (m *Core) XXX_Size() int {
	return xxx_messageInfo_Core.Size(m)
}
func (m *Core) XXX_DiscardUnknown() {
	xxx_messageInfo_Core.DiscardUnknown(m)
}

var xxx_messageInfo_Core proto.InternalMessageInfo

func (m *Core) GetCoreId() string {
	if m != nil {
		return m.CoreId
	}
	return ""
}

func (m *Core) GetIsHyperThread() bool {
	if m != nil {
		return m.IsHyperThread
	}
	return false
}

func (m *Core) GetHws() []*HardwareThread {
	if m != nil {
		return m.Hws
	}
	return nil
}

func (m *Core) GetL2CacheSize() int64 {
	if m != nil {
		return m.L2CacheSize
	}
	return 0
}

func (m *Core) GetL1CacheSize() int64 {
	if m != nil {
		return m.L1CacheSize
	}
	return 0
}

type Machine struct {
	MachineId            string    `protobuf:"bytes,1,opt,name=machine_id,json=machineId,proto3" json:"machine_id,omitempty"`
	Processor            string    `protobuf:"bytes,2,opt,name=processor,proto3" json:"processor,omitempty"`
	NumCores             int32     `protobuf:"varint,3,opt,name=num_cores,json=numCores,proto3" json:"num_cores,omitempty"`
	NumSockets           int32     `protobuf:"varint,4,opt,name=num_sockets,json=numSockets,proto3" json:"num_sockets,omitempty"`
	Sockets              []*Socket `protobuf:"bytes,5,rep,name=sockets,proto3" json:"sockets,omitempty"`
	MemorySize           int64     `protobuf:"varint,6,opt,name=memory_size,json=memorySize,proto3" json:"memory_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Machine) Reset()         { *m = Machine{} }
func (m *Machine) String() string { return proto.CompactTextString(m) }
func (*Machine) ProtoMessage()    {}
func (*Machine) Descriptor() ([]byte, []int) {
	return fileDescriptor_e909cb39bc20a1cf, []int{3}
}

func (m *Machine) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Machine.Unmarshal(m, b)
}
func (m *Machine) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Machine.Marshal(b, m, deterministic)
}
func (m *Machine) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Machine.Merge(m, src)
}
func (m *Machine) XXX_Size() int {
	return xxx_messageInfo_Machine.Size(m)
}
func (m *Machine) XXX_DiscardUnknown() {
	xxx_messageInfo_Machine.DiscardUnknown(m)
}

var xxx_messageInfo_Machine proto.InternalMessageInfo

func (m *Machine) GetMachineId() string {
	if m != nil {
		return m.MachineId
	}
	return ""
}

func (m *Machine) GetProcessor() string {
	if m != nil {
		return m.Processor
	}
	return ""
}

func (m *Machine) GetNumCores() int32 {
	if m != nil {
		return m.NumCores
	}
	return 0
}

func (m *Machine) GetNumSockets() int32 {
	if m != nil {
		return m.NumSockets
	}
	return 0
}

func (m *Machine) GetSockets() []*Socket {
	if m != nil {
		return m.Sockets
	}
	return nil
}

func (m *Machine) GetMemorySize() int64 {
	if m != nil {
		return m.MemorySize
	}
	return 0
}

type CreateMachineRequest struct {
	Machine              *Machine `protobuf:"bytes,1,opt,name=machine,proto3" json:"machine,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMachineRequest) Reset()         { *m = CreateMachineRequest{} }
func (m *CreateMachineRequest) String() string { return proto.CompactTextString(m) }
func (*CreateMachineRequest) ProtoMessage()    {}
func (*CreateMachineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e909cb39bc20a1cf, []int{4}
}

func (m *CreateMachineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMachineRequest.Unmarshal(m, b)
}
func (m *CreateMachineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMachineRequest.Marshal(b, m, deterministic)
}
func (m *CreateMachineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMachineRequest.Merge(m, src)
}
func (m *CreateMachineRequest) XXX_Size() int {
	return xxx_messageInfo_CreateMachineRequest.Size(m)
}
func (m *CreateMachineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMachineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMachineRequest proto.InternalMessageInfo

func (m *CreateMachineRequest) GetMachine() *Machine {
	if m != nil {
		return m.Machine
	}
	return nil
}

type CreateMachineResponse struct {
	Machine              *Machine `protobuf:"bytes,1,opt,name=machine,proto3" json:"machine,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMachineResponse) Reset()         { *m = CreateMachineResponse{} }
func (m *CreateMachineResponse) String() string { return proto.CompactTextString(m) }
func (*CreateMachineResponse) ProtoMessage()    {}
func (*CreateMachineResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e909cb39bc20a1cf, []int{5}
}

func (m *CreateMachineResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMachineResponse.Unmarshal(m, b)
}
func (m *CreateMachineResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMachineResponse.Marshal(b, m, deterministic)
}
func (m *CreateMachineResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMachineResponse.Merge(m, src)
}
func (m *CreateMachineResponse) XXX_Size() int {
	return xxx_messageInfo_CreateMachineResponse.Size(m)
}
func (m *CreateMachineResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMachineResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMachineResponse proto.InternalMessageInfo

func (m *CreateMachineResponse) GetMachine() *Machine {
	if m != nil {
		return m.Machine
	}
	return nil
}

type ReadMachineRequest struct {
	MachineId            string   `protobuf:"bytes,1,opt,name=machine_id,json=machineId,proto3" json:"machine_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadMachineRequest) Reset()         { *m = ReadMachineRequest{} }
func (m *ReadMachineRequest) String() string { return proto.CompactTextString(m) }
func (*ReadMachineRequest) ProtoMessage()    {}
func (*ReadMachineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e909cb39bc20a1cf, []int{6}
}

func (m *ReadMachineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadMachineRequest.Unmarshal(m, b)
}
func (m *ReadMachineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadMachineRequest.Marshal(b, m, deterministic)
}
func (m *ReadMachineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadMachineRequest.Merge(m, src)
}
func (m *ReadMachineRequest) XXX_Size() int {
	return xxx_messageInfo_ReadMachineRequest.Size(m)
}
func (m *ReadMachineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadMachineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadMachineRequest proto.InternalMessageInfo

func (m *ReadMachineRequest) GetMachineId() string {
	if m != nil {
		return m.MachineId
	}
	return ""
}

type ReadMachineResponse struct {
	Machine              *Machine `protobuf:"bytes,1,opt,name=machine,proto3" json:"machine,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadMachineResponse) Reset()         { *m = ReadMachineResponse{} }
func (m *ReadMachineResponse) String() string { return proto.CompactTextString(m) }
func (*ReadMachineResponse) ProtoMessage()    {}
func (*ReadMachineResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e909cb39bc20a1cf, []int{7}
}

func (m *ReadMachineResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadMachineResponse.Unmarshal(m, b)
}
func (m *ReadMachineResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadMachineResponse.Marshal(b, m, deterministic)
}
func (m *ReadMachineResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadMachineResponse.Merge(m, src)
}
func (m *ReadMachineResponse) XXX_Size() int {
	return xxx_messageInfo_ReadMachineResponse.Size(m)
}
func (m *ReadMachineResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadMachineResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadMachineResponse proto.InternalMessageInfo

func (m *ReadMachineResponse) GetMachine() *Machine {
	if m != nil {
		return m.Machine
	}
	return nil
}

type UpdateMachineRequest struct {
	Machine              *Machine `protobuf:"bytes,1,opt,name=machine,proto3" json:"machine,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateMachineRequest) Reset()         { *m = UpdateMachineRequest{} }
func (m *UpdateMachineRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateMachineRequest) ProtoMessage()    {}
func (*UpdateMachineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e909cb39bc20a1cf, []int{8}
}

func (m *UpdateMachineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateMachineRequest.Unmarshal(m, b)
}
func (m *UpdateMachineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateMachineRequest.Marshal(b, m, deterministic)
}
func (m *UpdateMachineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateMachineRequest.Merge(m, src)
}
func (m *UpdateMachineRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateMachineRequest.Size(m)
}
func (m *UpdateMachineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateMachineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateMachineRequest proto.InternalMessageInfo

func (m *UpdateMachineRequest) GetMachine() *Machine {
	if m != nil {
		return m.Machine
	}
	return nil
}

type UpdateMachineResponse struct {
	Machine              *Machine `protobuf:"bytes,1,opt,name=machine,proto3" json:"machine,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateMachineResponse) Reset()         { *m = UpdateMachineResponse{} }
func (m *UpdateMachineResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateMachineResponse) ProtoMessage()    {}
func (*UpdateMachineResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e909cb39bc20a1cf, []int{9}
}

func (m *UpdateMachineResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateMachineResponse.Unmarshal(m, b)
}
func (m *UpdateMachineResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateMachineResponse.Marshal(b, m, deterministic)
}
func (m *UpdateMachineResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateMachineResponse.Merge(m, src)
}
func (m *UpdateMachineResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateMachineResponse.Size(m)
}
func (m *UpdateMachineResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateMachineResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateMachineResponse proto.InternalMessageInfo

func (m *UpdateMachineResponse) GetMachine() *Machine {
	if m != nil {
		return m.Machine
	}
	return nil
}

type DeleteMachineRequest struct {
	MachineId            string   `protobuf:"bytes,1,opt,name=machine_id,json=machineId,proto3" json:"machine_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteMachineRequest) Reset()         { *m = DeleteMachineRequest{} }
func (m *DeleteMachineRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteMachineRequest) ProtoMessage()    {}
func (*DeleteMachineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e909cb39bc20a1cf, []int{10}
}

func (m *DeleteMachineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteMachineRequest.Unmarshal(m, b)
}
func (m *DeleteMachineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteMachineRequest.Marshal(b, m, deterministic)
}
func (m *DeleteMachineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteMachineRequest.Merge(m, src)
}
func (m *DeleteMachineRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteMachineRequest.Size(m)
}
func (m *DeleteMachineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteMachineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteMachineRequest proto.InternalMessageInfo

func (m *DeleteMachineRequest) GetMachineId() string {
	if m != nil {
		return m.MachineId
	}
	return ""
}

type DeleteMachineResponse struct {
	MachineId            string   `protobuf:"bytes,1,opt,name=machine_id,json=machineId,proto3" json:"machine_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteMachineResponse) Reset()         { *m = DeleteMachineResponse{} }
func (m *DeleteMachineResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteMachineResponse) ProtoMessage()    {}
func (*DeleteMachineResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e909cb39bc20a1cf, []int{11}
}

func (m *DeleteMachineResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteMachineResponse.Unmarshal(m, b)
}
func (m *DeleteMachineResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteMachineResponse.Marshal(b, m, deterministic)
}
func (m *DeleteMachineResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteMachineResponse.Merge(m, src)
}
func (m *DeleteMachineResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteMachineResponse.Size(m)
}
func (m *DeleteMachineResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteMachineResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteMachineResponse proto.InternalMessageInfo

func (m *DeleteMachineResponse) GetMachineId() string {
	if m != nil {
		return m.MachineId
	}
	return ""
}

type ListMachineRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListMachineRequest) Reset()         { *m = ListMachineRequest{} }
func (m *ListMachineRequest) String() string { return proto.CompactTextString(m) }
func (*ListMachineRequest) ProtoMessage()    {}
func (*ListMachineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e909cb39bc20a1cf, []int{12}
}

func (m *ListMachineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMachineRequest.Unmarshal(m, b)
}
func (m *ListMachineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMachineRequest.Marshal(b, m, deterministic)
}
func (m *ListMachineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMachineRequest.Merge(m, src)
}
func (m *ListMachineRequest) XXX_Size() int {
	return xxx_messageInfo_ListMachineRequest.Size(m)
}
func (m *ListMachineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMachineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListMachineRequest proto.InternalMessageInfo

type ListMachineResponse struct {
	Machine              *Machine `protobuf:"bytes,1,opt,name=machine,proto3" json:"machine,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListMachineResponse) Reset()         { *m = ListMachineResponse{} }
func (m *ListMachineResponse) String() string { return proto.CompactTextString(m) }
func (*ListMachineResponse) ProtoMessage()    {}
func (*ListMachineResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e909cb39bc20a1cf, []int{13}
}

func (m *ListMachineResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMachineResponse.Unmarshal(m, b)
}
func (m *ListMachineResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMachineResponse.Marshal(b, m, deterministic)
}
func (m *ListMachineResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMachineResponse.Merge(m, src)
}
func (m *ListMachineResponse) XXX_Size() int {
	return xxx_messageInfo_ListMachineResponse.Size(m)
}
func (m *ListMachineResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMachineResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListMachineResponse proto.InternalMessageInfo

func (m *ListMachineResponse) GetMachine() *Machine {
	if m != nil {
		return m.Machine
	}
	return nil
}

func init() {
	proto.RegisterType((*Socket)(nil), "machines.Socket")
	proto.RegisterType((*HardwareThread)(nil), "machines.HardwareThread")
	proto.RegisterType((*Core)(nil), "machines.Core")
	proto.RegisterType((*Machine)(nil), "machines.Machine")
	proto.RegisterType((*CreateMachineRequest)(nil), "machines.CreateMachineRequest")
	proto.RegisterType((*CreateMachineResponse)(nil), "machines.CreateMachineResponse")
	proto.RegisterType((*ReadMachineRequest)(nil), "machines.ReadMachineRequest")
	proto.RegisterType((*ReadMachineResponse)(nil), "machines.ReadMachineResponse")
	proto.RegisterType((*UpdateMachineRequest)(nil), "machines.UpdateMachineRequest")
	proto.RegisterType((*UpdateMachineResponse)(nil), "machines.UpdateMachineResponse")
	proto.RegisterType((*DeleteMachineRequest)(nil), "machines.DeleteMachineRequest")
	proto.RegisterType((*DeleteMachineResponse)(nil), "machines.DeleteMachineResponse")
	proto.RegisterType((*ListMachineRequest)(nil), "machines.ListMachineRequest")
	proto.RegisterType((*ListMachineResponse)(nil), "machines.ListMachineResponse")
}

func init() { proto.RegisterFile("machines/machinespb/machines.proto", fileDescriptor_e909cb39bc20a1cf) }

var fileDescriptor_e909cb39bc20a1cf = []byte{
	// 565 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0x65, 0xd2, 0x7c, 0x8d, 0x49, 0x29, 0x4b, 0x2a, 0xa2, 0xd2, 0xd2, 0xc8, 0x07, 0x14,
	0x15, 0x54, 0x68, 0x22, 0x38, 0x72, 0x68, 0x7a, 0x68, 0x50, 0x91, 0x22, 0x07, 0x2e, 0x5c, 0x2c,
	0xd7, 0x1e, 0x64, 0x8b, 0xf8, 0x83, 0x5d, 0x87, 0xaa, 0x7d, 0x1a, 0xde, 0x83, 0xa7, 0xe0, 0x8d,
	0xd0, 0xae, 0xd7, 0xd9, 0x75, 0xe2, 0xb6, 0x22, 0xbd, 0x79, 0xff, 0x33, 0x3b, 0x33, 0xbf, 0xf9,
	0xaf, 0x64, 0xb0, 0x22, 0xd7, 0x0b, 0xc2, 0x18, 0xd9, 0xdb, 0xe2, 0x23, 0xbd, 0x5c, 0x7e, 0x1e,
	0xa7, 0x34, 0xc9, 0x12, 0xd2, 0x2a, 0xce, 0xd6, 0x77, 0x68, 0xcc, 0x12, 0xef, 0x07, 0x66, 0xe4,
	0x05, 0xb4, 0x99, 0xf8, 0x72, 0x42, 0xbf, 0x67, 0xf4, 0x8d, 0x41, 0xdb, 0x6e, 0xe5, 0xc2, 0xc4,
	0xe7, 0xc1, 0x78, 0x11, 0x39, 0x5e, 0x42, 0x91, 0xf5, 0x6a, 0x7d, 0x63, 0x50, 0xb7, 0x5b, 0xf1,
	0x22, 0x1a, 0xf3, 0x33, 0xb1, 0xa0, 0x33, 0x1f, 0x39, 0x9e, 0xeb, 0x05, 0xe8, 0xb0, 0xf0, 0x06,
	0x7b, 0xf5, 0xbe, 0x31, 0xa8, 0xd9, 0xe6, 0x7c, 0x34, 0xe6, 0xda, 0x2c, 0xbc, 0x41, 0xeb, 0x23,
	0x6c, 0x9f, 0xbb, 0xd4, 0xbf, 0x72, 0x29, 0x7e, 0x09, 0x28, 0xba, 0x3e, 0x79, 0x03, 0x24, 0x90,
	0x8a, 0x93, 0x09, 0x49, 0x35, 0xde, 0x09, 0x4a, 0xb9, 0x13, 0xdf, 0xfa, 0x63, 0xc0, 0x16, 0xef,
	0x46, 0x9e, 0x43, 0x93, 0x4f, 0xa1, 0x72, 0x1b, 0xfc, 0x38, 0xf1, 0xc9, 0x2b, 0x78, 0x12, 0x32,
	0x27, 0xb8, 0x4e, 0x91, 0xca, 0x7a, 0xbd, 0x47, 0x7d, 0x63, 0xd0, 0xb2, 0x3b, 0x21, 0x3b, 0xe7,
	0xaa, 0xec, 0x7b, 0x04, 0xb5, 0xe0, 0x8a, 0x43, 0xd4, 0x06, 0xe6, 0xb0, 0x77, 0xbc, 0xdc, 0x4c,
	0x79, 0x3c, 0x9b, 0x27, 0x09, 0xb2, 0xa1, 0x4e, 0xb6, 0x25, 0xc9, 0x86, 0x4b, 0x32, 0x91, 0x73,
	0x52, 0x45, 0x7f, 0xa2, 0xe8, 0xff, 0x1a, 0xd0, 0xfc, 0x9c, 0x37, 0x22, 0x07, 0x00, 0xb2, 0xa7,
	0x62, 0x68, 0x4b, 0x65, 0xe2, 0x93, 0x7d, 0x68, 0xa7, 0x34, 0xf1, 0x90, 0xb1, 0x84, 0x0a, 0x80,
	0xb6, 0xad, 0x84, 0xbb, 0x7d, 0x38, 0x04, 0x93, 0x07, 0x73, 0xd3, 0x98, 0x98, 0xb5, 0x6e, 0x43,
	0xbc, 0x88, 0x72, 0x87, 0x19, 0x39, 0x82, 0x66, 0x11, 0xac, 0x0b, 0xfc, 0x1d, 0x85, 0x9f, 0xe7,
	0xd8, 0x45, 0x02, 0x2f, 0x16, 0x61, 0x94, 0xd0, 0xeb, 0x1c, 0xaa, 0x21, 0xa0, 0x20, 0x97, 0x04,
	0xd3, 0x18, 0xba, 0x63, 0x8a, 0x6e, 0x86, 0x12, 0xcc, 0xc6, 0x9f, 0x0b, 0x64, 0x19, 0x79, 0x0d,
	0x4d, 0x59, 0x54, 0xc0, 0x99, 0xc3, 0xa7, 0xaa, 0x49, 0x91, 0x5a, 0x64, 0x58, 0x67, 0xb0, 0xbb,
	0x52, 0x84, 0xa5, 0x49, 0xcc, 0xf0, 0xff, 0xaa, 0x8c, 0x80, 0xd8, 0xe8, 0xfa, 0x2b, 0x83, 0xdc,
	0xbd, 0x68, 0xeb, 0x14, 0x9e, 0x95, 0x2e, 0x6d, 0xd2, 0x78, 0x0c, 0xdd, 0xaf, 0xa9, 0xff, 0xf0,
	0x1d, 0xac, 0x14, 0xd9, 0x64, 0x94, 0xf7, 0xd0, 0x3d, 0xc3, 0x39, 0xae, 0x8d, 0x72, 0xcf, 0x16,
	0x3e, 0xc0, 0xee, 0xca, 0x35, 0xd9, 0xfc, 0x9e, 0x7b, 0x5d, 0x20, 0x17, 0x21, 0xcb, 0xca, 0xcd,
	0xf8, 0x4e, 0x4b, 0xea, 0x06, 0x20, 0xc3, 0xdf, 0x35, 0xd8, 0x96, 0xe2, 0x0c, 0xe9, 0xaf, 0xd0,
	0x43, 0x32, 0x85, 0x8e, 0xa7, 0xbf, 0x12, 0xf2, 0x52, 0xdd, 0xaf, 0x7a, 0x83, 0x7b, 0x87, 0xb7,
	0xc6, 0xe5, 0x44, 0x9f, 0xc0, 0xd4, 0xcc, 0x27, 0xfb, 0x2a, 0x7f, 0xfd, 0x21, 0xed, 0x1d, 0xdc,
	0x12, 0x95, 0xb5, 0xa6, 0xd0, 0x29, 0xf9, 0xa7, 0x4f, 0x57, 0xf5, 0x3a, 0xf4, 0xe9, 0xaa, 0x8d,
	0x9f, 0x42, 0xa7, 0x64, 0x8a, 0x5e, 0xb1, 0xca, 0x64, 0xbd, 0x62, 0xb5, 0x9b, 0x17, 0x60, 0x6a,
	0xc6, 0xe8, 0xbc, 0xeb, 0x2e, 0xea, 0xbc, 0x15, 0x6e, 0xbe, 0x33, 0x4e, 0x1f, 0x7f, 0x03, 0xf5,
	0x6f, 0xb9, 0x6c, 0x88, 0x7f, 0xca, 0xe8, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x67, 0x9d, 0xd6,
	0xe9, 0x79, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MachineServiceClient is the client API for MachineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MachineServiceClient interface {
	CreateMachine(ctx context.Context, in *CreateMachineRequest, opts ...grpc.CallOption) (*CreateMachineResponse, error)
	ReadMachine(ctx context.Context, in *ReadMachineRequest, opts ...grpc.CallOption) (*ReadMachineResponse, error)
	UpdateMachine(ctx context.Context, in *UpdateMachineRequest, opts ...grpc.CallOption) (*UpdateMachineResponse, error)
	DeleteMachine(ctx context.Context, in *DeleteMachineRequest, opts ...grpc.CallOption) (*DeleteMachineResponse, error)
	ListMachine(ctx context.Context, in *ListMachineRequest, opts ...grpc.CallOption) (MachineService_ListMachineClient, error)
}

type machineServiceClient struct {
	cc *grpc.ClientConn
}

func NewMachineServiceClient(cc *grpc.ClientConn) MachineServiceClient {
	return &machineServiceClient{cc}
}

func (c *machineServiceClient) CreateMachine(ctx context.Context, in *CreateMachineRequest, opts ...grpc.CallOption) (*CreateMachineResponse, error) {
	out := new(CreateMachineResponse)
	err := c.cc.Invoke(ctx, "/machines.MachineService/createMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineServiceClient) ReadMachine(ctx context.Context, in *ReadMachineRequest, opts ...grpc.CallOption) (*ReadMachineResponse, error) {
	out := new(ReadMachineResponse)
	err := c.cc.Invoke(ctx, "/machines.MachineService/ReadMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineServiceClient) UpdateMachine(ctx context.Context, in *UpdateMachineRequest, opts ...grpc.CallOption) (*UpdateMachineResponse, error) {
	out := new(UpdateMachineResponse)
	err := c.cc.Invoke(ctx, "/machines.MachineService/UpdateMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineServiceClient) DeleteMachine(ctx context.Context, in *DeleteMachineRequest, opts ...grpc.CallOption) (*DeleteMachineResponse, error) {
	out := new(DeleteMachineResponse)
	err := c.cc.Invoke(ctx, "/machines.MachineService/DeleteMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineServiceClient) ListMachine(ctx context.Context, in *ListMachineRequest, opts ...grpc.CallOption) (MachineService_ListMachineClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MachineService_serviceDesc.Streams[0], "/machines.MachineService/ListMachine", opts...)
	if err != nil {
		return nil, err
	}
	x := &machineServiceListMachineClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MachineService_ListMachineClient interface {
	Recv() (*ListMachineResponse, error)
	grpc.ClientStream
}

type machineServiceListMachineClient struct {
	grpc.ClientStream
}

func (x *machineServiceListMachineClient) Recv() (*ListMachineResponse, error) {
	m := new(ListMachineResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MachineServiceServer is the server API for MachineService service.
type MachineServiceServer interface {
	CreateMachine(context.Context, *CreateMachineRequest) (*CreateMachineResponse, error)
	ReadMachine(context.Context, *ReadMachineRequest) (*ReadMachineResponse, error)
	UpdateMachine(context.Context, *UpdateMachineRequest) (*UpdateMachineResponse, error)
	DeleteMachine(context.Context, *DeleteMachineRequest) (*DeleteMachineResponse, error)
	ListMachine(*ListMachineRequest, MachineService_ListMachineServer) error
}

// UnimplementedMachineServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMachineServiceServer struct {
}

func (*UnimplementedMachineServiceServer) CreateMachine(ctx context.Context, req *CreateMachineRequest) (*CreateMachineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMachine not implemented")
}
func (*UnimplementedMachineServiceServer) ReadMachine(ctx context.Context, req *ReadMachineRequest) (*ReadMachineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadMachine not implemented")
}
func (*UnimplementedMachineServiceServer) UpdateMachine(ctx context.Context, req *UpdateMachineRequest) (*UpdateMachineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMachine not implemented")
}
func (*UnimplementedMachineServiceServer) DeleteMachine(ctx context.Context, req *DeleteMachineRequest) (*DeleteMachineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMachine not implemented")
}
func (*UnimplementedMachineServiceServer) ListMachine(req *ListMachineRequest, srv MachineService_ListMachineServer) error {
	return status.Errorf(codes.Unimplemented, "method ListMachine not implemented")
}

func RegisterMachineServiceServer(s *grpc.Server, srv MachineServiceServer) {
	s.RegisterService(&_MachineService_serviceDesc, srv)
}

func _MachineService_CreateMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMachineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServiceServer).CreateMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/machines.MachineService/CreateMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServiceServer).CreateMachine(ctx, req.(*CreateMachineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MachineService_ReadMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadMachineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServiceServer).ReadMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/machines.MachineService/ReadMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServiceServer).ReadMachine(ctx, req.(*ReadMachineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MachineService_UpdateMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMachineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServiceServer).UpdateMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/machines.MachineService/UpdateMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServiceServer).UpdateMachine(ctx, req.(*UpdateMachineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MachineService_DeleteMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMachineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServiceServer).DeleteMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/machines.MachineService/DeleteMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServiceServer).DeleteMachine(ctx, req.(*DeleteMachineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MachineService_ListMachine_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListMachineRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MachineServiceServer).ListMachine(m, &machineServiceListMachineServer{stream})
}

type MachineService_ListMachineServer interface {
	Send(*ListMachineResponse) error
	grpc.ServerStream
}

type machineServiceListMachineServer struct {
	grpc.ServerStream
}

func (x *machineServiceListMachineServer) Send(m *ListMachineResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _MachineService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "machines.MachineService",
	HandlerType: (*MachineServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createMachine",
			Handler:    _MachineService_CreateMachine_Handler,
		},
		{
			MethodName: "ReadMachine",
			Handler:    _MachineService_ReadMachine_Handler,
		},
		{
			MethodName: "UpdateMachine",
			Handler:    _MachineService_UpdateMachine_Handler,
		},
		{
			MethodName: "DeleteMachine",
			Handler:    _MachineService_DeleteMachine_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListMachine",
			Handler:       _MachineService_ListMachine_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "machines/machinespb/machines.proto",
}