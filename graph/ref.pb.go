// Code generated by protoc-gen-gogo.
// source: ref.proto
// DO NOT EDIT!

package graph

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ref represents a reference from source code to a Def.
type Ref struct {
	// DefRepo is the repository URI of the Def that this Ref refers
	// to.
	DefRepo string `protobuf:"bytes,1,opt,name=def_repo,proto3" json:"DefRepo,omitempty"`
	// DefUnitType is the source unit type of the Def that this Ref refers to.
	DefUnitType string `protobuf:"bytes,3,opt,name=def_unit_type,proto3" json:"DefUnitType,omitempty"`
	// DefUnit is the name of the source unit that this ref exists in.
	DefUnit string `protobuf:"bytes,4,opt,name=def_unit,proto3" json:"DefUnit,omitempty"`
	// Path is the path of the Def that this ref refers to.
	DefPath string `protobuf:"bytes,5,opt,name=def_path,proto3" json:"DefPath"`
	// Repo is the VCS repository in which this ref exists.
	Repo string `protobuf:"bytes,6,opt,name=repo,proto3" json:"Repo,omitempty"`
	// CommitID is the ID of the VCS commit that this ref exists
	// in. The CommitID is always a full commit ID (40 hexadecimal
	// characters for git and hg), never a branch or tag name.
	CommitID string `protobuf:"bytes,7,opt,name=commit_id,proto3" json:"CommitID,omitempty"`
	// UnitType is the type name of the source unit that this ref
	// exists in.
	UnitType string `protobuf:"bytes,8,opt,name=unit_type,proto3" json:"UnitType,omitempty"`
	// Unit is the name of the source unit that this ref exists in.
	Unit string `protobuf:"bytes,9,opt,name=unit,proto3" json:"Unit,omitempty"`
	// Def is true if this Ref spans the name of the Def it points to.
	Def bool `protobuf:"varint,17,opt,name=def,proto3" json:"Def,omitempty"`
	// File is the filename in which this Ref exists.
	File string `protobuf:"bytes,10,opt,name=file,proto3" json:"File,omitempty"`
	// Start is the byte offset of this ref's first byte in File.
	Start uint32 `protobuf:"varint,11,opt,name=start,proto3" json:"Start"`
	// End is the byte offset of this ref's last byte in File.
	End uint32 `protobuf:"varint,12,opt,name=end,proto3" json:"End"`
}

func (m *Ref) Reset()         { *m = Ref{} }
func (m *Ref) String() string { return proto.CompactTextString(m) }
func (*Ref) ProtoMessage()    {}

type RefDefKey struct {
	DefRepo     string `protobuf:"bytes,1,opt,name=def_repo,proto3" json:"DefRepo,omitempty"`
	DefUnitType string `protobuf:"bytes,3,opt,name=def_unit_type,proto3" json:"DefUnitType,omitempty"`
	DefUnit     string `protobuf:"bytes,4,opt,name=def_unit,proto3" json:"DefUnit,omitempty"`
	DefPath     string `protobuf:"bytes,5,opt,name=def_path,proto3" json:"DefPath"`
}

func (m *RefDefKey) Reset()         { *m = RefDefKey{} }
func (m *RefDefKey) String() string { return proto.CompactTextString(m) }
func (*RefDefKey) ProtoMessage()    {}

func (m *Ref) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Ref) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.DefRepo) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintRef(data, i, uint64(len(m.DefRepo)))
		i += copy(data[i:], m.DefRepo)
	}
	if len(m.DefUnitType) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintRef(data, i, uint64(len(m.DefUnitType)))
		i += copy(data[i:], m.DefUnitType)
	}
	if len(m.DefUnit) > 0 {
		data[i] = 0x22
		i++
		i = encodeVarintRef(data, i, uint64(len(m.DefUnit)))
		i += copy(data[i:], m.DefUnit)
	}
	if len(m.DefPath) > 0 {
		data[i] = 0x2a
		i++
		i = encodeVarintRef(data, i, uint64(len(m.DefPath)))
		i += copy(data[i:], m.DefPath)
	}
	if len(m.Repo) > 0 {
		data[i] = 0x32
		i++
		i = encodeVarintRef(data, i, uint64(len(m.Repo)))
		i += copy(data[i:], m.Repo)
	}
	if len(m.CommitID) > 0 {
		data[i] = 0x3a
		i++
		i = encodeVarintRef(data, i, uint64(len(m.CommitID)))
		i += copy(data[i:], m.CommitID)
	}
	if len(m.UnitType) > 0 {
		data[i] = 0x42
		i++
		i = encodeVarintRef(data, i, uint64(len(m.UnitType)))
		i += copy(data[i:], m.UnitType)
	}
	if len(m.Unit) > 0 {
		data[i] = 0x4a
		i++
		i = encodeVarintRef(data, i, uint64(len(m.Unit)))
		i += copy(data[i:], m.Unit)
	}
	if len(m.File) > 0 {
		data[i] = 0x52
		i++
		i = encodeVarintRef(data, i, uint64(len(m.File)))
		i += copy(data[i:], m.File)
	}
	if m.Start != 0 {
		data[i] = 0x58
		i++
		i = encodeVarintRef(data, i, uint64(m.Start))
	}
	if m.End != 0 {
		data[i] = 0x60
		i++
		i = encodeVarintRef(data, i, uint64(m.End))
	}
	if m.Def {
		data[i] = 0x88
		i++
		data[i] = 0x1
		i++
		if m.Def {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *RefDefKey) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RefDefKey) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.DefRepo) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintRef(data, i, uint64(len(m.DefRepo)))
		i += copy(data[i:], m.DefRepo)
	}
	if len(m.DefUnitType) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintRef(data, i, uint64(len(m.DefUnitType)))
		i += copy(data[i:], m.DefUnitType)
	}
	if len(m.DefUnit) > 0 {
		data[i] = 0x22
		i++
		i = encodeVarintRef(data, i, uint64(len(m.DefUnit)))
		i += copy(data[i:], m.DefUnit)
	}
	if len(m.DefPath) > 0 {
		data[i] = 0x2a
		i++
		i = encodeVarintRef(data, i, uint64(len(m.DefPath)))
		i += copy(data[i:], m.DefPath)
	}
	return i, nil
}

func encodeFixed64Ref(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Ref(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintRef(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Ref) Size() (n int) {
	var l int
	_ = l
	l = len(m.DefRepo)
	if l > 0 {
		n += 1 + l + sovRef(uint64(l))
	}
	l = len(m.DefUnitType)
	if l > 0 {
		n += 1 + l + sovRef(uint64(l))
	}
	l = len(m.DefUnit)
	if l > 0 {
		n += 1 + l + sovRef(uint64(l))
	}
	l = len(m.DefPath)
	if l > 0 {
		n += 1 + l + sovRef(uint64(l))
	}
	l = len(m.Repo)
	if l > 0 {
		n += 1 + l + sovRef(uint64(l))
	}
	l = len(m.CommitID)
	if l > 0 {
		n += 1 + l + sovRef(uint64(l))
	}
	l = len(m.UnitType)
	if l > 0 {
		n += 1 + l + sovRef(uint64(l))
	}
	l = len(m.Unit)
	if l > 0 {
		n += 1 + l + sovRef(uint64(l))
	}
	l = len(m.File)
	if l > 0 {
		n += 1 + l + sovRef(uint64(l))
	}
	if m.Start != 0 {
		n += 1 + sovRef(uint64(m.Start))
	}
	if m.End != 0 {
		n += 1 + sovRef(uint64(m.End))
	}
	if m.Def {
		n += 3
	}
	return n
}

func (m *RefDefKey) Size() (n int) {
	var l int
	_ = l
	l = len(m.DefRepo)
	if l > 0 {
		n += 1 + l + sovRef(uint64(l))
	}
	l = len(m.DefUnitType)
	if l > 0 {
		n += 1 + l + sovRef(uint64(l))
	}
	l = len(m.DefUnit)
	if l > 0 {
		n += 1 + l + sovRef(uint64(l))
	}
	l = len(m.DefPath)
	if l > 0 {
		n += 1 + l + sovRef(uint64(l))
	}
	return n
}

func sovRef(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRef(x uint64) (n int) {
	return sovRef(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Ref) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRef
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Ref: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Ref: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefRepo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRef
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRef
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DefRepo = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefUnitType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRef
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRef
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DefUnitType = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefUnit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRef
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRef
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DefUnit = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRef
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRef
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DefPath = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Repo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRef
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRef
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Repo = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommitID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRef
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRef
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CommitID = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnitType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRef
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRef
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UnitType = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRef
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRef
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Unit = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field File", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRef
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRef
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.File = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			m.Start = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRef
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Start |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field End", wireType)
			}
			m.End = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRef
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.End |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 17:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Def", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRef
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Def = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRef(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRef
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
func (m *RefDefKey) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRef
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RefDefKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RefDefKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefRepo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRef
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRef
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DefRepo = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefUnitType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRef
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRef
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DefUnitType = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefUnit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRef
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRef
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DefUnit = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRef
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRef
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DefPath = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRef(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRef
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
func skipRef(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRef
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowRef
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowRef
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRef
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRef
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipRef(data[start:])
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
	ErrInvalidLengthRef = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRef   = fmt.Errorf("proto: integer overflow")
)
