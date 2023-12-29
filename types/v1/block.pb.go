// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types/v1/block.proto

package v1

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_prysmaticlabs_prysm_v4_consensus_types_primitives "github.com/prysmaticlabs/prysm/v4/consensus-types/primitives"
	github_com_prysmaticlabs_prysm_v4_math "github.com/prysmaticlabs/prysm/v4/math"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenericBeaconKitBlock defines the data of a beacon block that is injected by the proposer.
type BaseBeaconKitBlock struct {
	Slot     github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Slot `protobuf:"varint,1,opt,name=slot,proto3,casttype=github.com/prysmaticlabs/prysm/v4/consensus-types/primitives.Slot" json:"slot,omitempty"`
	Time     uint64                                                            `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	Value    github_com_prysmaticlabs_prysm_v4_math.Gwei                       `protobuf:"varint,3,opt,name=value,proto3,casttype=github.com/prysmaticlabs/prysm/v4/math.Gwei" json:"value,omitempty"`
	Version  uint64                                                            `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	ExecData []byte                                                            `protobuf:"bytes,5,opt,name=exec_data,json=execData,proto3" json:"exec_data,omitempty"`
}

func (m *BaseBeaconKitBlock) Reset()         { *m = BaseBeaconKitBlock{} }
func (m *BaseBeaconKitBlock) String() string { return proto.CompactTextString(m) }
func (*BaseBeaconKitBlock) ProtoMessage()    {}
func (*BaseBeaconKitBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c7c34efa65ddc20, []int{0}
}
func (m *BaseBeaconKitBlock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BaseBeaconKitBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BaseBeaconKitBlock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BaseBeaconKitBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseBeaconKitBlock.Merge(m, src)
}
func (m *BaseBeaconKitBlock) XXX_Size() int {
	return m.Size()
}
func (m *BaseBeaconKitBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseBeaconKitBlock.DiscardUnknown(m)
}

var xxx_messageInfo_BaseBeaconKitBlock proto.InternalMessageInfo

func (m *BaseBeaconKitBlock) GetSlot() github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Slot {
	if m != nil {
		return m.Slot
	}
	return 0
}

func (m *BaseBeaconKitBlock) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *BaseBeaconKitBlock) GetValue() github_com_prysmaticlabs_prysm_v4_math.Gwei {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *BaseBeaconKitBlock) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *BaseBeaconKitBlock) GetExecData() []byte {
	if m != nil {
		return m.ExecData
	}
	return nil
}

func init() {
	proto.RegisterType((*BaseBeaconKitBlock)(nil), "types.v1.BaseBeaconKitBlock")
}

func init() { proto.RegisterFile("types/v1/block.proto", fileDescriptor_4c7c34efa65ddc20) }

var fileDescriptor_4c7c34efa65ddc20 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x50, 0xcd, 0x4a, 0xf3, 0x40,
	0x14, 0x6d, 0xfa, 0xa5, 0x9f, 0x75, 0x70, 0x35, 0x74, 0x11, 0x14, 0xd2, 0xe2, 0x42, 0x0a, 0x62,
	0x86, 0xa2, 0x0f, 0xa0, 0xc1, 0xe2, 0xc2, 0x5d, 0x5c, 0xe9, 0x46, 0x26, 0xe9, 0xa5, 0x1d, 0x4c,
	0x32, 0x61, 0xee, 0xed, 0x68, 0x97, 0xbe, 0x81, 0x8f, 0xe5, 0xb2, 0x4b, 0x57, 0x45, 0x9a, 0xb7,
	0xe8, 0x4a, 0x32, 0xa1, 0xe0, 0xce, 0xdd, 0x3d, 0xe7, 0x72, 0x7e, 0x38, 0x6c, 0x40, 0xab, 0x0a,
	0x50, 0xd8, 0x89, 0x48, 0x73, 0x9d, 0xbd, 0x44, 0x95, 0xd1, 0xa4, 0x79, 0xdf, 0xb1, 0x91, 0x9d,
	0x1c, 0x0f, 0xe6, 0x7a, 0xae, 0x1d, 0x29, 0x9a, 0xab, 0xfd, 0x9f, 0xbe, 0x77, 0x19, 0x8f, 0x25,
	0x42, 0x0c, 0x32, 0xd3, 0xe5, 0xbd, 0xa2, 0xb8, 0x11, 0xf3, 0x47, 0xe6, 0x63, 0xae, 0x29, 0xf0,
	0x46, 0xde, 0xd8, 0x8f, 0xa7, 0xbb, 0xcd, 0xf0, 0x66, 0xae, 0x68, 0xb1, 0x4c, 0xa3, 0x4c, 0x17,
	0xa2, 0x32, 0x2b, 0x2c, 0x24, 0xa9, 0x2c, 0x97, 0x29, 0xb6, 0x48, 0xd8, 0x2b, 0x91, 0xe9, 0x12,
	0xa1, 0xc4, 0x25, 0x5e, 0xb4, 0x55, 0x2a, 0xa3, 0x0a, 0x45, 0xca, 0x02, 0x46, 0x0f, 0xb9, 0xa6,
	0xc4, 0x59, 0x72, 0xce, 0x7c, 0x52, 0x05, 0x04, 0xdd, 0xc6, 0x3a, 0x71, 0x37, 0x9f, 0xb2, 0x9e,
	0x95, 0xf9, 0x12, 0x82, 0x7f, 0x2e, 0x4f, 0xec, 0x36, 0xc3, 0xf3, 0xbf, 0xf3, 0x0a, 0x49, 0x8b,
	0xe8, 0xee, 0x15, 0x54, 0xd2, 0xaa, 0x79, 0xc0, 0x0e, 0x2c, 0x18, 0x54, 0xba, 0x0c, 0x7c, 0xe7,
	0xbe, 0x87, 0xfc, 0x84, 0x1d, 0xc2, 0x1b, 0x64, 0xcf, 0x33, 0x49, 0x32, 0xe8, 0x8d, 0xbc, 0xf1,
	0x51, 0xd2, 0x6f, 0x88, 0x5b, 0x49, 0x32, 0xbe, 0xfe, 0xdc, 0x86, 0xde, 0x7a, 0x1b, 0x7a, 0xdf,
	0xdb, 0xd0, 0xfb, 0xa8, 0xc3, 0xce, 0xba, 0x0e, 0x3b, 0x5f, 0x75, 0xd8, 0x79, 0x3a, 0xfb, 0x55,
	0x42, 0x11, 0xce, 0xc0, 0xa6, 0x20, 0x8d, 0x48, 0x75, 0x2e, 0x8d, 0x42, 0xb1, 0x5f, 0x3c, 0xfd,
	0xef, 0xc6, 0xbc, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x2c, 0x0f, 0xd1, 0x84, 0x01, 0x00,
	0x00,
}

func (m *BaseBeaconKitBlock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BaseBeaconKitBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BaseBeaconKitBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ExecData) > 0 {
		i -= len(m.ExecData)
		copy(dAtA[i:], m.ExecData)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.ExecData)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Version != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x20
	}
	if m.Value != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x18
	}
	if m.Time != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.Time))
		i--
		dAtA[i] = 0x10
	}
	if m.Slot != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.Slot))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBlock(dAtA []byte, offset int, v uint64) int {
	offset -= sovBlock(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BaseBeaconKitBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Slot != 0 {
		n += 1 + sovBlock(uint64(m.Slot))
	}
	if m.Time != 0 {
		n += 1 + sovBlock(uint64(m.Time))
	}
	if m.Value != 0 {
		n += 1 + sovBlock(uint64(m.Value))
	}
	if m.Version != 0 {
		n += 1 + sovBlock(uint64(m.Version))
	}
	l = len(m.ExecData)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	return n
}

func sovBlock(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBlock(x uint64) (n int) {
	return sovBlock(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BaseBeaconKitBlock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BaseBeaconKitBlock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BaseBeaconKitBlock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slot", wireType)
			}
			m.Slot = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Slot |= github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Slot(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Time |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= github_com_prysmaticlabs_prysm_v4_math.Gwei(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExecData = append(m.ExecData[:0], dAtA[iNdEx:postIndex]...)
			if m.ExecData == nil {
				m.ExecData = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlock
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
func skipBlock(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlock
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
					return 0, ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBlock
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
			if length < 0 {
				return 0, ErrInvalidLengthBlock
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBlock
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBlock
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBlock        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlock          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBlock = fmt.Errorf("proto: unexpected end of group")
)
