// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: v1/UserRank.proto

package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

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

type UserRankGetUserRankReq struct {
	// 房间号
	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid"`
	// 类型
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type"`
}

func (m *UserRankGetUserRankReq) Reset()         { *m = UserRankGetUserRankReq{} }
func (m *UserRankGetUserRankReq) String() string { return proto.CompactTextString(m) }
func (*UserRankGetUserRankReq) ProtoMessage()    {}
func (*UserRankGetUserRankReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_UserRank_7150ae0e89c4f10c, []int{0}
}
func (m *UserRankGetUserRankReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserRankGetUserRankReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserRankGetUserRankReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *UserRankGetUserRankReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRankGetUserRankReq.Merge(dst, src)
}
func (m *UserRankGetUserRankReq) XXX_Size() int {
	return m.Size()
}
func (m *UserRankGetUserRankReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRankGetUserRankReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserRankGetUserRankReq proto.InternalMessageInfo

func (m *UserRankGetUserRankReq) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *UserRankGetUserRankReq) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type UserRankGetUserRankResp struct {
	// code
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	// msg
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`
	//
	Data *UserRankGetUserRankResp_Data `protobuf:"bytes,3,opt,name=data" json:"data"`
}

func (m *UserRankGetUserRankResp) Reset()         { *m = UserRankGetUserRankResp{} }
func (m *UserRankGetUserRankResp) String() string { return proto.CompactTextString(m) }
func (*UserRankGetUserRankResp) ProtoMessage()    {}
func (*UserRankGetUserRankResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_UserRank_7150ae0e89c4f10c, []int{1}
}
func (m *UserRankGetUserRankResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserRankGetUserRankResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserRankGetUserRankResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *UserRankGetUserRankResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRankGetUserRankResp.Merge(dst, src)
}
func (m *UserRankGetUserRankResp) XXX_Size() int {
	return m.Size()
}
func (m *UserRankGetUserRankResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRankGetUserRankResp.DiscardUnknown(m)
}

var xxx_messageInfo_UserRankGetUserRankResp proto.InternalMessageInfo

func (m *UserRankGetUserRankResp) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UserRankGetUserRankResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *UserRankGetUserRankResp) GetData() *UserRankGetUserRankResp_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type UserRankGetUserRankResp_Data struct {
	//
	Rank int64 `protobuf:"varint,1,opt,name=rank,proto3" json:"rank"`
	//
	Score int64 `protobuf:"varint,2,opt,name=score,proto3" json:"score"`
}

func (m *UserRankGetUserRankResp_Data) Reset()         { *m = UserRankGetUserRankResp_Data{} }
func (m *UserRankGetUserRankResp_Data) String() string { return proto.CompactTextString(m) }
func (*UserRankGetUserRankResp_Data) ProtoMessage()    {}
func (*UserRankGetUserRankResp_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_UserRank_7150ae0e89c4f10c, []int{1, 0}
}
func (m *UserRankGetUserRankResp_Data) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserRankGetUserRankResp_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserRankGetUserRankResp_Data.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *UserRankGetUserRankResp_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRankGetUserRankResp_Data.Merge(dst, src)
}
func (m *UserRankGetUserRankResp_Data) XXX_Size() int {
	return m.Size()
}
func (m *UserRankGetUserRankResp_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRankGetUserRankResp_Data.DiscardUnknown(m)
}

var xxx_messageInfo_UserRankGetUserRankResp_Data proto.InternalMessageInfo

func (m *UserRankGetUserRankResp_Data) GetRank() int64 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *UserRankGetUserRankResp_Data) GetScore() int64 {
	if m != nil {
		return m.Score
	}
	return 0
}

func init() {
	proto.RegisterType((*UserRankGetUserRankReq)(nil), "rankdb.v1.UserRankGetUserRankReq")
	proto.RegisterType((*UserRankGetUserRankResp)(nil), "rankdb.v1.UserRankGetUserRankResp")
	proto.RegisterType((*UserRankGetUserRankResp_Data)(nil), "rankdb.v1.UserRankGetUserRankResp.Data")
}
func (m *UserRankGetUserRankReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserRankGetUserRankReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Uid != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintUserRank(dAtA, i, uint64(m.Uid))
	}
	if len(m.Type) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintUserRank(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	return i, nil
}

func (m *UserRankGetUserRankResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserRankGetUserRankResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintUserRank(dAtA, i, uint64(m.Code))
	}
	if len(m.Msg) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintUserRank(dAtA, i, uint64(len(m.Msg)))
		i += copy(dAtA[i:], m.Msg)
	}
	if m.Data != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintUserRank(dAtA, i, uint64(m.Data.Size()))
		n1, err := m.Data.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *UserRankGetUserRankResp_Data) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserRankGetUserRankResp_Data) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Rank != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintUserRank(dAtA, i, uint64(m.Rank))
	}
	if m.Score != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintUserRank(dAtA, i, uint64(m.Score))
	}
	return i, nil
}

func encodeVarintUserRank(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *UserRankGetUserRankReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Uid != 0 {
		n += 1 + sovUserRank(uint64(m.Uid))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovUserRank(uint64(l))
	}
	return n
}

func (m *UserRankGetUserRankResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovUserRank(uint64(m.Code))
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovUserRank(uint64(l))
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovUserRank(uint64(l))
	}
	return n
}

func (m *UserRankGetUserRankResp_Data) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Rank != 0 {
		n += 1 + sovUserRank(uint64(m.Rank))
	}
	if m.Score != 0 {
		n += 1 + sovUserRank(uint64(m.Score))
	}
	return n
}

func sovUserRank(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozUserRank(x uint64) (n int) {
	return sovUserRank(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UserRankGetUserRankReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUserRank
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
			return fmt.Errorf("proto: UserRankGetUserRankReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserRankGetUserRankReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserRank
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserRank
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
				return ErrInvalidLengthUserRank
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUserRank(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUserRank
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
func (m *UserRankGetUserRankResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUserRank
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
			return fmt.Errorf("proto: UserRankGetUserRankResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserRankGetUserRankResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserRank
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserRank
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
				return ErrInvalidLengthUserRank
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserRank
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
				return ErrInvalidLengthUserRank
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &UserRankGetUserRankResp_Data{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUserRank(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUserRank
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
func (m *UserRankGetUserRankResp_Data) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUserRank
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
			return fmt.Errorf("proto: Data: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Data: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rank", wireType)
			}
			m.Rank = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserRank
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Rank |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Score", wireType)
			}
			m.Score = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserRank
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Score |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUserRank(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUserRank
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
func skipUserRank(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUserRank
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
					return 0, ErrIntOverflowUserRank
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
					return 0, ErrIntOverflowUserRank
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
				return 0, ErrInvalidLengthUserRank
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowUserRank
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
				next, err := skipUserRank(dAtA[start:])
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
	ErrInvalidLengthUserRank = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUserRank   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("v1/UserRank.proto", fileDescriptor_UserRank_7150ae0e89c4f10c) }

var fileDescriptor_UserRank_7150ae0e89c4f10c = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0xbf, 0x4e, 0xc2, 0x40,
	0x18, 0xe7, 0x28, 0x2a, 0x1c, 0x93, 0x37, 0x28, 0x12, 0x72, 0x45, 0x16, 0x59, 0x3c, 0x02, 0xbe,
	0x41, 0x23, 0x71, 0xf6, 0xa2, 0x8b, 0x93, 0x57, 0x5a, 0x4f, 0x42, 0xca, 0x95, 0xf6, 0xda, 0xc4,
	0xb7, 0xf0, 0xb1, 0x1c, 0x19, 0x9d, 0x1a, 0xd2, 0x6e, 0x7d, 0x0a, 0x73, 0x5f, 0xa9, 0x3a, 0x68,
	0x58, 0x7e, 0xfd, 0x7e, 0x5f, 0xfa, 0xfb, 0xf3, 0xe5, 0xf0, 0x69, 0x3a, 0x9d, 0x3c, 0xc6, 0x7e,
	0xc4, 0xc5, 0x7a, 0xc5, 0xc2, 0x48, 0x69, 0x45, 0x3a, 0x91, 0x58, 0xaf, 0x3c, 0x97, 0xa5, 0xd3,
	0xfe, 0xb5, 0x5c, 0xea, 0xd7, 0xc4, 0x65, 0x0b, 0x15, 0x4c, 0xa4, 0x92, 0x6a, 0x02, 0x7f, 0xb8,
	0xc9, 0x0b, 0x30, 0x20, 0x30, 0x55, 0xca, 0xd1, 0x3d, 0x3e, 0xab, 0xbd, 0xee, 0x7c, 0x5d, 0x8f,
	0xdc, 0xdf, 0x90, 0x0b, 0x6c, 0x25, 0x4b, 0xaf, 0x87, 0x86, 0x68, 0x6c, 0x39, 0x27, 0x65, 0x66,
	0x1b, 0xca, 0x0d, 0x90, 0x01, 0x6e, 0xe9, 0xb7, 0xd0, 0xef, 0x35, 0x87, 0x68, 0xdc, 0x71, 0xda,
	0x65, 0x66, 0x03, 0xe7, 0x80, 0xa3, 0x1d, 0xc2, 0xe7, 0x7f, 0x7a, 0xc6, 0xa1, 0x51, 0x2e, 0x94,
	0xe7, 0xef, 0x5d, 0x41, 0x69, 0x38, 0x07, 0x34, 0x91, 0x41, 0x2c, 0xf7, 0xb6, 0x10, 0x19, 0xc4,
	0x92, 0x1b, 0x20, 0x73, 0xdc, 0xf2, 0x84, 0x16, 0x3d, 0x6b, 0x88, 0xc6, 0xdd, 0xd9, 0x15, 0xfb,
	0x3e, 0x98, 0xfd, 0x13, 0xc5, 0x6e, 0x85, 0x16, 0x55, 0x82, 0x11, 0x72, 0xc0, 0xfe, 0x1c, 0xb7,
	0xcc, 0xde, 0xf4, 0x30, 0x0e, 0xbf, 0x7b, 0x18, 0xce, 0x01, 0x89, 0x8d, 0x8f, 0xe2, 0x85, 0x8a,
	0xaa, 0x03, 0x2d, 0xa7, 0x53, 0x66, 0x76, 0xb5, 0xe0, 0xd5, 0x67, 0xf6, 0x8c, 0xdb, 0x75, 0x16,
	0x79, 0xc0, 0x5d, 0xf9, 0x13, 0x4d, 0x2e, 0x0f, 0x55, 0xdb, 0xf4, 0x47, 0x87, 0xdb, 0x3b, 0x83,
	0x8f, 0x9c, 0xa2, 0x6d, 0x4e, 0xd1, 0x2e, 0xa7, 0xe8, 0xbd, 0xa0, 0x8d, 0x6d, 0x41, 0x1b, 0x9f,
	0x05, 0x6d, 0x3c, 0x35, 0xd3, 0xa9, 0x7b, 0x0c, 0x8f, 0x77, 0xf3, 0x15, 0x00, 0x00, 0xff, 0xff,
	0xec, 0xff, 0x41, 0x5d, 0x0b, 0x02, 0x00, 0x00,
}