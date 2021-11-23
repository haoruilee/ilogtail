// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sls_logs.proto

/*
	Package sls_logs is a generated protocol buffer package.

	It is generated from these files:
		sls_logs.proto

	It has these top-level messages:
		Log
		LogTag
		LogGroup
		SlsLogPackage
		SlsLogPackageList
		LogGroupList
*/
package protocol

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Log struct {
	Time     uint32         `protobuf:"varint,1,req,name=Time" json:"Time"`
	Contents []*Log_Content `protobuf:"bytes,2,rep,name=Contents" json:"Contents,omitempty"`
}

func (m *Log) Reset()                    { *m = Log{} }
func (m *Log) String() string            { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()               {}
func (*Log) Descriptor() ([]byte, []int) { return fileDescriptorSlsLogs, []int{0} }

func (m *Log) GetTime() uint32 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Log) GetContents() []*Log_Content {
	if m != nil {
		return m.Contents
	}
	return nil
}

type Log_Content struct {
	Key   string `protobuf:"bytes,1,req,name=Key" json:"Key"`
	Value string `protobuf:"bytes,2,req,name=Value" json:"Value"`
}

func (m *Log_Content) Reset()                    { *m = Log_Content{} }
func (m *Log_Content) String() string            { return proto.CompactTextString(m) }
func (*Log_Content) ProtoMessage()               {}
func (*Log_Content) Descriptor() ([]byte, []int) { return fileDescriptorSlsLogs, []int{0, 0} }

func (m *Log_Content) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Log_Content) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type LogTag struct {
	Key   string `protobuf:"bytes,1,req,name=Key" json:"Key"`
	Value string `protobuf:"bytes,2,req,name=Value" json:"Value"`
}

func (m *LogTag) Reset()                    { *m = LogTag{} }
func (m *LogTag) String() string            { return proto.CompactTextString(m) }
func (*LogTag) ProtoMessage()               {}
func (*LogTag) Descriptor() ([]byte, []int) { return fileDescriptorSlsLogs, []int{1} }

func (m *LogTag) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *LogTag) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type LogGroup struct {
	Logs        []*Log    `protobuf:"bytes,1,rep,name=Logs" json:"Logs,omitempty"`
	Category    string    `protobuf:"bytes,2,opt,name=Category" json:"Category"`
	Topic       string    `protobuf:"bytes,3,opt,name=Topic" json:"Topic"`
	Source      string    `protobuf:"bytes,4,opt,name=Source" json:"Source"`
	MachineUUID string    `protobuf:"bytes,5,opt,name=MachineUUID" json:"MachineUUID"`
	LogTags     []*LogTag `protobuf:"bytes,6,rep,name=LogTags" json:"LogTags,omitempty"`
}

func (m *LogGroup) Reset()                    { *m = LogGroup{} }
func (m *LogGroup) String() string            { return proto.CompactTextString(m) }
func (*LogGroup) ProtoMessage()               {}
func (*LogGroup) Descriptor() ([]byte, []int) { return fileDescriptorSlsLogs, []int{2} }

func (m *LogGroup) GetLogs() []*Log {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *LogGroup) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *LogGroup) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *LogGroup) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *LogGroup) GetMachineUUID() string {
	if m != nil {
		return m.MachineUUID
	}
	return ""
}

func (m *LogGroup) GetLogTags() []*LogTag {
	if m != nil {
		return m.LogTags
	}
	return nil
}

type SlsLogPackage struct {
	Data           []byte `protobuf:"bytes,1,req,name=data" json:"data"`
	UncompressSize int32  `protobuf:"varint,2,opt,name=uncompress_size,json=uncompressSize" json:"uncompress_size"`
}

func (m *SlsLogPackage) Reset()                    { *m = SlsLogPackage{} }
func (m *SlsLogPackage) String() string            { return proto.CompactTextString(m) }
func (*SlsLogPackage) ProtoMessage()               {}
func (*SlsLogPackage) Descriptor() ([]byte, []int) { return fileDescriptorSlsLogs, []int{3} }

func (m *SlsLogPackage) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *SlsLogPackage) GetUncompressSize() int32 {
	if m != nil {
		return m.UncompressSize
	}
	return 0
}

type SlsLogPackageList struct {
	Packages []*SlsLogPackage `protobuf:"bytes,1,rep,name=packages" json:"packages,omitempty"`
}

func (m *SlsLogPackageList) Reset()                    { *m = SlsLogPackageList{} }
func (m *SlsLogPackageList) String() string            { return proto.CompactTextString(m) }
func (*SlsLogPackageList) ProtoMessage()               {}
func (*SlsLogPackageList) Descriptor() ([]byte, []int) { return fileDescriptorSlsLogs, []int{4} }

func (m *SlsLogPackageList) GetPackages() []*SlsLogPackage {
	if m != nil {
		return m.Packages
	}
	return nil
}

type LogGroupList struct {
	LogGroupList []*LogGroup `protobuf:"bytes,1,rep,name=logGroupList" json:"logGroupList,omitempty"`
}

func (m *LogGroupList) Reset()                    { *m = LogGroupList{} }
func (m *LogGroupList) String() string            { return proto.CompactTextString(m) }
func (*LogGroupList) ProtoMessage()               {}
func (*LogGroupList) Descriptor() ([]byte, []int) { return fileDescriptorSlsLogs, []int{5} }

func (m *LogGroupList) GetLogGroupList() []*LogGroup {
	if m != nil {
		return m.LogGroupList
	}
	return nil
}

func init() {
	proto.RegisterType((*Log)(nil), "sls_logs.Log")
	proto.RegisterType((*Log_Content)(nil), "sls_logs.Log.Content")
	proto.RegisterType((*LogTag)(nil), "sls_logs.LogTag")
	proto.RegisterType((*LogGroup)(nil), "sls_logs.LogGroup")
	proto.RegisterType((*SlsLogPackage)(nil), "sls_logs.SlsLogPackage")
	proto.RegisterType((*SlsLogPackageList)(nil), "sls_logs.SlsLogPackageList")
	proto.RegisterType((*LogGroupList)(nil), "sls_logs.LogGroupList")
}
func (m *Log) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Log) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintSlsLogs(dAtA, i, uint64(m.Time))
	if len(m.Contents) > 0 {
		for _, msg := range m.Contents {
			dAtA[i] = 0x12
			i++
			i = encodeVarintSlsLogs(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Log_Content) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Log_Content) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintSlsLogs(dAtA, i, uint64(len(m.Key)))
	i += copy(dAtA[i:], m.Key)
	dAtA[i] = 0x12
	i++
	i = encodeVarintSlsLogs(dAtA, i, uint64(len(m.Value)))
	i += copy(dAtA[i:], m.Value)
	return i, nil
}

func (m *LogTag) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LogTag) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintSlsLogs(dAtA, i, uint64(len(m.Key)))
	i += copy(dAtA[i:], m.Key)
	dAtA[i] = 0x12
	i++
	i = encodeVarintSlsLogs(dAtA, i, uint64(len(m.Value)))
	i += copy(dAtA[i:], m.Value)
	return i, nil
}

func (m *LogGroup) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LogGroup) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Logs) > 0 {
		for _, msg := range m.Logs {
			dAtA[i] = 0xa
			i++
			i = encodeVarintSlsLogs(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintSlsLogs(dAtA, i, uint64(len(m.Category)))
	i += copy(dAtA[i:], m.Category)
	dAtA[i] = 0x1a
	i++
	i = encodeVarintSlsLogs(dAtA, i, uint64(len(m.Topic)))
	i += copy(dAtA[i:], m.Topic)
	dAtA[i] = 0x22
	i++
	i = encodeVarintSlsLogs(dAtA, i, uint64(len(m.Source)))
	i += copy(dAtA[i:], m.Source)
	dAtA[i] = 0x2a
	i++
	i = encodeVarintSlsLogs(dAtA, i, uint64(len(m.MachineUUID)))
	i += copy(dAtA[i:], m.MachineUUID)
	if len(m.LogTags) > 0 {
		for _, msg := range m.LogTags {
			dAtA[i] = 0x32
			i++
			i = encodeVarintSlsLogs(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *SlsLogPackage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SlsLogPackage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Data != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSlsLogs(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	dAtA[i] = 0x10
	i++
	i = encodeVarintSlsLogs(dAtA, i, uint64(m.UncompressSize))
	return i, nil
}

func (m *SlsLogPackageList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SlsLogPackageList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Packages) > 0 {
		for _, msg := range m.Packages {
			dAtA[i] = 0xa
			i++
			i = encodeVarintSlsLogs(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *LogGroupList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LogGroupList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.LogGroupList) > 0 {
		for _, msg := range m.LogGroupList {
			dAtA[i] = 0xa
			i++
			i = encodeVarintSlsLogs(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintSlsLogs(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Log) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovSlsLogs(uint64(m.Time))
	if len(m.Contents) > 0 {
		for _, e := range m.Contents {
			l = e.Size()
			n += 1 + l + sovSlsLogs(uint64(l))
		}
	}
	return n
}

func (m *Log_Content) Size() (n int) {
	var l int
	_ = l
	l = len(m.Key)
	n += 1 + l + sovSlsLogs(uint64(l))
	l = len(m.Value)
	n += 1 + l + sovSlsLogs(uint64(l))
	return n
}

func (m *LogTag) Size() (n int) {
	var l int
	_ = l
	l = len(m.Key)
	n += 1 + l + sovSlsLogs(uint64(l))
	l = len(m.Value)
	n += 1 + l + sovSlsLogs(uint64(l))
	return n
}

func (m *LogGroup) Size() (n int) {
	var l int
	_ = l
	if len(m.Logs) > 0 {
		for _, e := range m.Logs {
			l = e.Size()
			n += 1 + l + sovSlsLogs(uint64(l))
		}
	}
	l = len(m.Category)
	n += 1 + l + sovSlsLogs(uint64(l))
	l = len(m.Topic)
	n += 1 + l + sovSlsLogs(uint64(l))
	l = len(m.Source)
	n += 1 + l + sovSlsLogs(uint64(l))
	l = len(m.MachineUUID)
	n += 1 + l + sovSlsLogs(uint64(l))
	if len(m.LogTags) > 0 {
		for _, e := range m.LogTags {
			l = e.Size()
			n += 1 + l + sovSlsLogs(uint64(l))
		}
	}
	return n
}

func (m *SlsLogPackage) Size() (n int) {
	var l int
	_ = l
	if m.Data != nil {
		l = len(m.Data)
		n += 1 + l + sovSlsLogs(uint64(l))
	}
	n += 1 + sovSlsLogs(uint64(m.UncompressSize))
	return n
}

func (m *SlsLogPackageList) Size() (n int) {
	var l int
	_ = l
	if len(m.Packages) > 0 {
		for _, e := range m.Packages {
			l = e.Size()
			n += 1 + l + sovSlsLogs(uint64(l))
		}
	}
	return n
}

func (m *LogGroupList) Size() (n int) {
	var l int
	_ = l
	if len(m.LogGroupList) > 0 {
		for _, e := range m.LogGroupList {
			l = e.Size()
			n += 1 + l + sovSlsLogs(uint64(l))
		}
	}
	return n
}

func sovSlsLogs(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSlsLogs(x uint64) (n int) {
	return sovSlsLogs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Log) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSlsLogs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Log: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Log: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlsLogs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Time |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contents", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlsLogs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSlsLogs
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contents = append(m.Contents, &Log_Content{})
			if err := m.Contents[len(m.Contents)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSlsLogs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSlsLogs
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return proto.NewRequiredNotSetError("Time")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Log_Content) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSlsLogs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Content: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Content: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlsLogs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSlsLogs
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlsLogs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSlsLogs
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000002)
		default:
			iNdEx = preIndex
			skippy, err := skipSlsLogs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSlsLogs
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return proto.NewRequiredNotSetError("Key")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return proto.NewRequiredNotSetError("Value")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LogTag) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSlsLogs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LogTag: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LogTag: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlsLogs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSlsLogs
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlsLogs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSlsLogs
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000002)
		default:
			iNdEx = preIndex
			skippy, err := skipSlsLogs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSlsLogs
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return proto.NewRequiredNotSetError("Key")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return proto.NewRequiredNotSetError("Value")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LogGroup) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSlsLogs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LogGroup: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LogGroup: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Logs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlsLogs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSlsLogs
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Logs = append(m.Logs, &Log{})
			if err := m.Logs[len(m.Logs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Category", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlsLogs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSlsLogs
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Category = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Topic", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlsLogs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSlsLogs
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Topic = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Source", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlsLogs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSlsLogs
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Source = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MachineUUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlsLogs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSlsLogs
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MachineUUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogTags", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlsLogs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSlsLogs
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LogTags = append(m.LogTags, &LogTag{})
			if err := m.LogTags[len(m.LogTags)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSlsLogs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSlsLogs
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SlsLogPackage) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSlsLogs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SlsLogPackage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SlsLogPackage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlsLogs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSlsLogs
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UncompressSize", wireType)
			}
			m.UncompressSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlsLogs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UncompressSize |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSlsLogs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSlsLogs
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return proto.NewRequiredNotSetError("data")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SlsLogPackageList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSlsLogs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SlsLogPackageList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SlsLogPackageList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Packages", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlsLogs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSlsLogs
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Packages = append(m.Packages, &SlsLogPackage{})
			if err := m.Packages[len(m.Packages)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSlsLogs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSlsLogs
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LogGroupList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSlsLogs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LogGroupList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LogGroupList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogGroupList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlsLogs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSlsLogs
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LogGroupList = append(m.LogGroupList, &LogGroup{})
			if err := m.LogGroupList[len(m.LogGroupList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSlsLogs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSlsLogs
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSlsLogs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSlsLogs
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSlsLogs
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSlsLogs
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSlsLogs
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSlsLogs
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSlsLogs(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSlsLogs = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSlsLogs   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("sls_logs.proto", fileDescriptorSlsLogs) }

var fileDescriptorSlsLogs = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xcf, 0xaa, 0xd3, 0x40,
	0x14, 0xc6, 0x3b, 0x49, 0xda, 0xc6, 0xd3, 0x3f, 0xd6, 0x01, 0x75, 0x28, 0x12, 0x63, 0x16, 0x52,
	0x04, 0x0b, 0x2a, 0xb8, 0xd2, 0x4d, 0x2b, 0xfe, 0xc1, 0x08, 0x92, 0xa6, 0xe2, 0xae, 0x0c, 0x71,
	0x18, 0x83, 0x69, 0x26, 0x64, 0x92, 0x45, 0xfb, 0x14, 0xfa, 0x56, 0x5d, 0xfa, 0x02, 0xca, 0xa5,
	0xf7, 0x45, 0x2e, 0x99, 0xfc, 0x69, 0xb2, 0xbd, 0xbb, 0x9e, 0xef, 0xf7, 0xf5, 0xcb, 0x77, 0xe6,
	0xc0, 0x54, 0x46, 0x72, 0x17, 0x09, 0x2e, 0x97, 0x49, 0x2a, 0x32, 0x81, 0xcd, 0x7a, 0x76, 0xfe,
	0x20, 0xd0, 0x5d, 0xc1, 0x31, 0x01, 0xc3, 0x0f, 0xf7, 0x8c, 0x20, 0x5b, 0x5b, 0x4c, 0x56, 0xc6,
	0xe9, 0xff, 0xe3, 0x9e, 0xa7, 0x14, 0xfc, 0x02, 0xcc, 0xb5, 0x88, 0x33, 0x16, 0x67, 0x92, 0x68,
	0xb6, 0xbe, 0x18, 0xbd, 0xbc, 0xbf, 0x6c, 0xe2, 0x5c, 0xc1, 0x97, 0x15, 0xf5, 0x1a, 0xdb, 0xfc,
	0x2d, 0x0c, 0xab, 0xdf, 0xf8, 0x01, 0xe8, 0x9f, 0xd9, 0x41, 0xc5, 0xde, 0xa9, 0x62, 0x0b, 0x01,
	0xcf, 0xa1, 0xff, 0x8d, 0x46, 0x39, 0x23, 0x5a, 0x8b, 0x94, 0x92, 0xf3, 0x06, 0x06, 0xae, 0xe0,
	0x3e, 0xe5, 0xb7, 0xfa, 0xf7, 0x3f, 0x04, 0xa6, 0x2b, 0xf8, 0x87, 0x54, 0xe4, 0x09, 0x7e, 0x02,
	0x86, 0x2b, 0xb8, 0x24, 0x48, 0x15, 0x9f, 0x74, 0x8a, 0x7b, 0x0a, 0x61, 0x1b, 0xcc, 0x35, 0xcd,
	0x18, 0x17, 0xe9, 0x81, 0x68, 0x36, 0x6a, 0xe2, 0x1a, 0xb5, 0xf8, 0x9a, 0x2f, 0x92, 0x30, 0x20,
	0x7a, 0x0b, 0x97, 0x12, 0x7e, 0x04, 0x83, 0x8d, 0xc8, 0xd3, 0x80, 0x11, 0xa3, 0x05, 0x2b, 0x0d,
	0x3f, 0x85, 0xd1, 0x17, 0x1a, 0xfc, 0x0c, 0x63, 0xb6, 0xdd, 0x7e, 0x7a, 0x47, 0xfa, 0x2d, 0x4b,
	0x1b, 0xe0, 0x67, 0x30, 0x2c, 0x37, 0x96, 0x64, 0xa0, 0x9a, 0xce, 0x3a, 0x4d, 0x7d, 0xca, 0xbd,
	0xda, 0xe0, 0x7c, 0x87, 0xc9, 0x26, 0x92, 0xae, 0xe0, 0x5f, 0x69, 0xf0, 0x8b, 0x72, 0x56, 0x9c,
	0xee, 0x07, 0xcd, 0xa8, 0x7a, 0xa5, 0x71, 0x7d, 0xba, 0x42, 0xc1, 0xcf, 0xe1, 0x6e, 0x1e, 0x07,
	0x62, 0x9f, 0xa4, 0x4c, 0xca, 0x9d, 0x0c, 0x8f, 0x4c, 0x6d, 0xd8, 0xaf, 0x4c, 0xd3, 0x0b, 0xdc,
	0x84, 0x47, 0xe6, 0x7c, 0x84, 0x7b, 0x9d, 0x64, 0x37, 0x94, 0x19, 0x7e, 0x05, 0x66, 0x52, 0x8e,
	0xf5, 0x2b, 0x3e, 0xbc, 0x74, 0xeb, 0xd8, 0xbd, 0xc6, 0xe8, 0xbc, 0x87, 0x71, 0x7d, 0x02, 0x15,
	0xf2, 0x1a, 0xc6, 0x51, 0x6b, 0xae, 0x82, 0x70, 0x67, 0x49, 0x45, 0xbd, 0x8e, 0x6f, 0x35, 0x3b,
	0x9d, 0x2d, 0xf4, 0xf7, 0x6c, 0xa1, 0xab, 0xb3, 0x85, 0x7e, 0x5f, 0x5b, 0xbd, 0x9b, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x25, 0x11, 0x72, 0x20, 0xca, 0x02, 0x00, 0x00,
}
