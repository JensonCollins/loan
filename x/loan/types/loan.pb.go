// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: loan/loan.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type Loan struct {
	Id         uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount     string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Fee        string `protobuf:"bytes,3,opt,name=fee,proto3" json:"fee,omitempty"`
	Collateral string `protobuf:"bytes,4,opt,name=collateral,proto3" json:"collateral,omitempty"`
	Deadline   string `protobuf:"bytes,5,opt,name=deadline,proto3" json:"deadline,omitempty"`
	State      string `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty"`
	Borrower   string `protobuf:"bytes,7,opt,name=borrower,proto3" json:"borrower,omitempty"`
	Lender     string `protobuf:"bytes,8,opt,name=lender,proto3" json:"lender,omitempty"`
}

func (m *Loan) Reset()         { *m = Loan{} }
func (m *Loan) String() string { return proto.CompactTextString(m) }
func (*Loan) ProtoMessage()    {}
func (*Loan) Descriptor() ([]byte, []int) {
	return fileDescriptor_f667ff5036a9b094, []int{0}
}
func (m *Loan) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Loan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Loan.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Loan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Loan.Merge(m, src)
}
func (m *Loan) XXX_Size() int {
	return m.Size()
}
func (m *Loan) XXX_DiscardUnknown() {
	xxx_messageInfo_Loan.DiscardUnknown(m)
}

var xxx_messageInfo_Loan proto.InternalMessageInfo

func (m *Loan) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Loan) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *Loan) GetFee() string {
	if m != nil {
		return m.Fee
	}
	return ""
}

func (m *Loan) GetCollateral() string {
	if m != nil {
		return m.Collateral
	}
	return ""
}

func (m *Loan) GetDeadline() string {
	if m != nil {
		return m.Deadline
	}
	return ""
}

func (m *Loan) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Loan) GetBorrower() string {
	if m != nil {
		return m.Borrower
	}
	return ""
}

func (m *Loan) GetLender() string {
	if m != nil {
		return m.Lender
	}
	return ""
}

func init() {
	proto.RegisterType((*Loan)(nil), "jensoncollins.loan.loan.Loan")
}

func init() { proto.RegisterFile("loan/loan.proto", fileDescriptor_f667ff5036a9b094) }

var fileDescriptor_f667ff5036a9b094 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0xe3, 0x34, 0x0d, 0xe5, 0x06, 0x40, 0x16, 0x02, 0x8b, 0xc1, 0xaa, 0x98, 0x2a, 0x86,
	0x64, 0xe0, 0x0d, 0x10, 0x23, 0x53, 0x47, 0x36, 0xa7, 0x3e, 0xc0, 0xc8, 0xf5, 0x55, 0x8e, 0x2b,
	0xe0, 0x2d, 0x78, 0x2a, 0xc4, 0xd8, 0x91, 0x11, 0x25, 0x2f, 0x82, 0x6c, 0x03, 0x82, 0xe5, 0xec,
	0xef, 0x3e, 0xdd, 0xe9, 0xf4, 0xc3, 0xa1, 0x25, 0xe5, 0xda, 0x58, 0x9a, 0x8d, 0xa7, 0x40, 0xfc,
	0xf4, 0x11, 0x5d, 0x4f, 0x6e, 0x45, 0xd6, 0x1a, 0xd7, 0x37, 0xc9, 0xc4, 0x72, 0xfe, 0xc6, 0xa0,
	0xba, 0x21, 0xe5, 0xf8, 0x01, 0x94, 0x46, 0x0b, 0x36, 0x67, 0x8b, 0x6a, 0x59, 0x1a, 0xcd, 0x4f,
	0xa0, 0x56, 0x6b, 0xda, 0xba, 0x20, 0xca, 0x39, 0x5b, 0xec, 0x2f, 0xbf, 0x89, 0x1f, 0xc1, 0xe4,
	0x0e, 0x51, 0x4c, 0x52, 0x33, 0x7e, 0xb9, 0x04, 0x88, 0x7b, 0x55, 0x40, 0xaf, 0xac, 0xa8, 0x92,
	0xf8, 0xd3, 0xe1, 0x67, 0x30, 0xd3, 0xa8, 0xb4, 0x35, 0x0e, 0xc5, 0x34, 0xd9, 0x5f, 0xe6, 0xc7,
	0x30, 0xed, 0x83, 0x0a, 0x28, 0xea, 0x24, 0x32, 0xc4, 0x89, 0x8e, 0xbc, 0xa7, 0x27, 0xf4, 0x62,
	0x2f, 0x4f, 0xfc, 0x70, 0xbc, 0xcb, 0xa2, 0xd3, 0xe8, 0xc5, 0x2c, 0xdf, 0x95, 0xe9, 0xea, 0xfa,
	0x7d, 0x90, 0x6c, 0x37, 0x48, 0xf6, 0x39, 0x48, 0xf6, 0x3a, 0xca, 0x62, 0x37, 0xca, 0xe2, 0x63,
	0x94, 0xc5, 0xed, 0xc5, 0xbd, 0x09, 0x0f, 0xdb, 0xae, 0x59, 0xd1, 0xba, 0xfd, 0x17, 0x43, 0x0a,
	0xa8, 0x7d, 0xce, 0x4f, 0x78, 0xd9, 0x60, 0xdf, 0xd5, 0x29, 0xae, 0xcb, 0xaf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xf0, 0x06, 0x91, 0xa4, 0x41, 0x01, 0x00, 0x00,
}

func (m *Loan) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Loan) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Loan) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Lender) > 0 {
		i -= len(m.Lender)
		copy(dAtA[i:], m.Lender)
		i = encodeVarintLoan(dAtA, i, uint64(len(m.Lender)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Borrower) > 0 {
		i -= len(m.Borrower)
		copy(dAtA[i:], m.Borrower)
		i = encodeVarintLoan(dAtA, i, uint64(len(m.Borrower)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.State) > 0 {
		i -= len(m.State)
		copy(dAtA[i:], m.State)
		i = encodeVarintLoan(dAtA, i, uint64(len(m.State)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Deadline) > 0 {
		i -= len(m.Deadline)
		copy(dAtA[i:], m.Deadline)
		i = encodeVarintLoan(dAtA, i, uint64(len(m.Deadline)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Collateral) > 0 {
		i -= len(m.Collateral)
		copy(dAtA[i:], m.Collateral)
		i = encodeVarintLoan(dAtA, i, uint64(len(m.Collateral)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Fee) > 0 {
		i -= len(m.Fee)
		copy(dAtA[i:], m.Fee)
		i = encodeVarintLoan(dAtA, i, uint64(len(m.Fee)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintLoan(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintLoan(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintLoan(dAtA []byte, offset int, v uint64) int {
	offset -= sovLoan(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Loan) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovLoan(uint64(m.Id))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovLoan(uint64(l))
	}
	l = len(m.Fee)
	if l > 0 {
		n += 1 + l + sovLoan(uint64(l))
	}
	l = len(m.Collateral)
	if l > 0 {
		n += 1 + l + sovLoan(uint64(l))
	}
	l = len(m.Deadline)
	if l > 0 {
		n += 1 + l + sovLoan(uint64(l))
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovLoan(uint64(l))
	}
	l = len(m.Borrower)
	if l > 0 {
		n += 1 + l + sovLoan(uint64(l))
	}
	l = len(m.Lender)
	if l > 0 {
		n += 1 + l + sovLoan(uint64(l))
	}
	return n
}

func sovLoan(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLoan(x uint64) (n int) {
	return sovLoan(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Loan) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLoan
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
			return fmt.Errorf("proto: Loan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Loan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLoan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLoan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Collateral", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLoan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Collateral = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deadline", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLoan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Deadline = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLoan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.State = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Borrower", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLoan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Borrower = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLoan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Lender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLoan(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLoan
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
func skipLoan(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLoan
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
					return 0, ErrIntOverflowLoan
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
					return 0, ErrIntOverflowLoan
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
				return 0, ErrInvalidLengthLoan
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLoan
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLoan
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLoan        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLoan          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLoan = fmt.Errorf("proto: unexpected end of group")
)
