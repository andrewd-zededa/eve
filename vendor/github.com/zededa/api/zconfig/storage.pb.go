// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage.proto

package zconfig

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DsType int32

const (
	DsType_DsUnknown DsType = 0
	DsType_DsHttp    DsType = 1
	DsType_DsHttps   DsType = 2
	DsType_DsS3      DsType = 3
	DsType_DsSFTP    DsType = 4
)

var DsType_name = map[int32]string{
	0: "DsUnknown",
	1: "DsHttp",
	2: "DsHttps",
	3: "DsS3",
	4: "DsSFTP",
}
var DsType_value = map[string]int32{
	"DsUnknown": 0,
	"DsHttp":    1,
	"DsHttps":   2,
	"DsS3":      3,
	"DsSFTP":    4,
}

func (x DsType) String() string {
	return proto.EnumName(DsType_name, int32(x))
}
func (DsType) EnumDescriptor() ([]byte, []int) { return fileDescriptorStorage, []int{0} }

type Format int32

const (
	Format_FmtUnknown Format = 0
	Format_Raw        Format = 1
	Format_QCOW       Format = 2
	Format_QCOW2      Format = 3
	Format_VHD        Format = 4
	Format_VMDK       Format = 5
	Format_OVA        Format = 6
	Format_VHDX       Format = 7
)

var Format_name = map[int32]string{
	0: "FmtUnknown",
	1: "Raw",
	2: "QCOW",
	3: "QCOW2",
	4: "VHD",
	5: "VMDK",
	6: "OVA",
	7: "VHDX",
}
var Format_value = map[string]int32{
	"FmtUnknown": 0,
	"Raw":        1,
	"QCOW":       2,
	"QCOW2":      3,
	"VHD":        4,
	"VMDK":       5,
	"OVA":        6,
	"VHDX":       7,
}

func (x Format) String() string {
	return proto.EnumName(Format_name, int32(x))
}
func (Format) EnumDescriptor() ([]byte, []int) { return fileDescriptorStorage, []int{1} }

type Target int32

const (
	Target_Disk    Target = 0
	Target_Kernel  Target = 1
	Target_Initrd  Target = 2
	Target_RamDisk Target = 3
)

var Target_name = map[int32]string{
	0: "Disk",
	1: "Kernel",
	2: "Initrd",
	3: "RamDisk",
}
var Target_value = map[string]int32{
	"Disk":    0,
	"Kernel":  1,
	"Initrd":  2,
	"RamDisk": 3,
}

func (x Target) String() string {
	return proto.EnumName(Target_name, int32(x))
}
func (Target) EnumDescriptor() ([]byte, []int) { return fileDescriptorStorage, []int{2} }

type DriveType int32

const (
	DriveType_Unclassified DriveType = 0
	DriveType_CDROM        DriveType = 1
	DriveType_HDD          DriveType = 2
	DriveType_NET          DriveType = 3
)

var DriveType_name = map[int32]string{
	0: "Unclassified",
	1: "CDROM",
	2: "HDD",
	3: "NET",
}
var DriveType_value = map[string]int32{
	"Unclassified": 0,
	"CDROM":        1,
	"HDD":          2,
	"NET":          3,
}

func (x DriveType) String() string {
	return proto.EnumName(DriveType_name, int32(x))
}
func (DriveType) EnumDescriptor() ([]byte, []int) { return fileDescriptorStorage, []int{3} }

type SignatureInfo struct {
	Intercertsurl string `protobuf:"bytes,1,opt,name=intercertsurl,proto3" json:"intercertsurl,omitempty"`
	Signercerturl string `protobuf:"bytes,2,opt,name=signercerturl,proto3" json:"signercerturl,omitempty"`
	Signature     []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *SignatureInfo) Reset()                    { *m = SignatureInfo{} }
func (m *SignatureInfo) String() string            { return proto.CompactTextString(m) }
func (*SignatureInfo) ProtoMessage()               {}
func (*SignatureInfo) Descriptor() ([]byte, []int) { return fileDescriptorStorage, []int{0} }

func (m *SignatureInfo) GetIntercertsurl() string {
	if m != nil {
		return m.Intercertsurl
	}
	return ""
}

func (m *SignatureInfo) GetSignercerturl() string {
	if m != nil {
		return m.Signercerturl
	}
	return ""
}

func (m *SignatureInfo) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type DatastoreConfig struct {
	Id       string `protobuf:"bytes,100,opt,name=id,proto3" json:"id,omitempty"`
	DType    DsType `protobuf:"varint,1,opt,name=dType,proto3,enum=DsType" json:"dType,omitempty"`
	Fqdn     string `protobuf:"bytes,2,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	ApiKey   string `protobuf:"bytes,3,opt,name=apiKey,proto3" json:"apiKey,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	// depending on datastore types, it could be bucket or path
	Dpath string `protobuf:"bytes,5,opt,name=dpath,proto3" json:"dpath,omitempty"`
}

func (m *DatastoreConfig) Reset()                    { *m = DatastoreConfig{} }
func (m *DatastoreConfig) String() string            { return proto.CompactTextString(m) }
func (*DatastoreConfig) ProtoMessage()               {}
func (*DatastoreConfig) Descriptor() ([]byte, []int) { return fileDescriptorStorage, []int{1} }

func (m *DatastoreConfig) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DatastoreConfig) GetDType() DsType {
	if m != nil {
		return m.DType
	}
	return DsType_DsUnknown
}

func (m *DatastoreConfig) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

func (m *DatastoreConfig) GetApiKey() string {
	if m != nil {
		return m.ApiKey
	}
	return ""
}

func (m *DatastoreConfig) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *DatastoreConfig) GetDpath() string {
	if m != nil {
		return m.Dpath
	}
	return ""
}

type Image struct {
	Id string `protobuf:"bytes,100,opt,name=id,proto3" json:"id,omitempty"`
	// it could be relative path/name as well
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Sha256  string `protobuf:"bytes,3,opt,name=sha256,proto3" json:"sha256,omitempty"`
	Iformat Format `protobuf:"varint,4,opt,name=iformat,proto3,enum=Format" json:"iformat,omitempty"`
	// if its signed image
	Siginfo *SignatureInfo `protobuf:"bytes,5,opt,name=siginfo" json:"siginfo,omitempty"`
	DsId    string         `protobuf:"bytes,6,opt,name=dsId,proto3" json:"dsId,omitempty"`
}

func (m *Image) Reset()                    { *m = Image{} }
func (m *Image) String() string            { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()               {}
func (*Image) Descriptor() ([]byte, []int) { return fileDescriptorStorage, []int{2} }

func (m *Image) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Image) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Image) GetSha256() string {
	if m != nil {
		return m.Sha256
	}
	return ""
}

func (m *Image) GetIformat() Format {
	if m != nil {
		return m.Iformat
	}
	return Format_FmtUnknown
}

func (m *Image) GetSiginfo() *SignatureInfo {
	if m != nil {
		return m.Siginfo
	}
	return nil
}

func (m *Image) GetDsId() string {
	if m != nil {
		return m.DsId
	}
	return ""
}

type Drive struct {
	Image    *Image    `protobuf:"bytes,1,opt,name=image" json:"image,omitempty"`
	Maxsize  int64     `protobuf:"varint,2,opt,name=maxsize,proto3" json:"maxsize,omitempty"`
	Readonly bool      `protobuf:"varint,5,opt,name=readonly,proto3" json:"readonly,omitempty"`
	Preserve bool      `protobuf:"varint,6,opt,name=preserve,proto3" json:"preserve,omitempty"`
	Drvtype  DriveType `protobuf:"varint,8,opt,name=drvtype,proto3,enum=DriveType" json:"drvtype,omitempty"`
	Target   Target    `protobuf:"varint,9,opt,name=target,proto3,enum=Target" json:"target,omitempty"`
}

func (m *Drive) Reset()                    { *m = Drive{} }
func (m *Drive) String() string            { return proto.CompactTextString(m) }
func (*Drive) ProtoMessage()               {}
func (*Drive) Descriptor() ([]byte, []int) { return fileDescriptorStorage, []int{3} }

func (m *Drive) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *Drive) GetMaxsize() int64 {
	if m != nil {
		return m.Maxsize
	}
	return 0
}

func (m *Drive) GetReadonly() bool {
	if m != nil {
		return m.Readonly
	}
	return false
}

func (m *Drive) GetPreserve() bool {
	if m != nil {
		return m.Preserve
	}
	return false
}

func (m *Drive) GetDrvtype() DriveType {
	if m != nil {
		return m.Drvtype
	}
	return DriveType_Unclassified
}

func (m *Drive) GetTarget() Target {
	if m != nil {
		return m.Target
	}
	return Target_Disk
}

func init() {
	proto.RegisterType((*SignatureInfo)(nil), "SignatureInfo")
	proto.RegisterType((*DatastoreConfig)(nil), "DatastoreConfig")
	proto.RegisterType((*Image)(nil), "Image")
	proto.RegisterType((*Drive)(nil), "Drive")
	proto.RegisterEnum("DsType", DsType_name, DsType_value)
	proto.RegisterEnum("Format", Format_name, Format_value)
	proto.RegisterEnum("Target", Target_name, Target_value)
	proto.RegisterEnum("DriveType", DriveType_name, DriveType_value)
}
func (m *SignatureInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignatureInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Intercertsurl) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.Intercertsurl)))
		i += copy(dAtA[i:], m.Intercertsurl)
	}
	if len(m.Signercerturl) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.Signercerturl)))
		i += copy(dAtA[i:], m.Signercerturl)
	}
	if len(m.Signature) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.Signature)))
		i += copy(dAtA[i:], m.Signature)
	}
	return i, nil
}

func (m *DatastoreConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DatastoreConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.DType != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintStorage(dAtA, i, uint64(m.DType))
	}
	if len(m.Fqdn) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.Fqdn)))
		i += copy(dAtA[i:], m.Fqdn)
	}
	if len(m.ApiKey) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.ApiKey)))
		i += copy(dAtA[i:], m.ApiKey)
	}
	if len(m.Password) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.Password)))
		i += copy(dAtA[i:], m.Password)
	}
	if len(m.Dpath) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.Dpath)))
		i += copy(dAtA[i:], m.Dpath)
	}
	if len(m.Id) > 0 {
		dAtA[i] = 0xa2
		i++
		dAtA[i] = 0x6
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	return i, nil
}

func (m *Image) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Image) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Sha256) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.Sha256)))
		i += copy(dAtA[i:], m.Sha256)
	}
	if m.Iformat != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintStorage(dAtA, i, uint64(m.Iformat))
	}
	if m.Siginfo != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintStorage(dAtA, i, uint64(m.Siginfo.Size()))
		n1, err := m.Siginfo.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.DsId) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.DsId)))
		i += copy(dAtA[i:], m.DsId)
	}
	if len(m.Id) > 0 {
		dAtA[i] = 0xa2
		i++
		dAtA[i] = 0x6
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	return i, nil
}

func (m *Drive) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Drive) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Image != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintStorage(dAtA, i, uint64(m.Image.Size()))
		n2, err := m.Image.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Maxsize != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintStorage(dAtA, i, uint64(m.Maxsize))
	}
	if m.Readonly {
		dAtA[i] = 0x28
		i++
		if m.Readonly {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Preserve {
		dAtA[i] = 0x30
		i++
		if m.Preserve {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Drvtype != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintStorage(dAtA, i, uint64(m.Drvtype))
	}
	if m.Target != 0 {
		dAtA[i] = 0x48
		i++
		i = encodeVarintStorage(dAtA, i, uint64(m.Target))
	}
	return i, nil
}

func encodeVarintStorage(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SignatureInfo) Size() (n int) {
	var l int
	_ = l
	l = len(m.Intercertsurl)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	l = len(m.Signercerturl)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	return n
}

func (m *DatastoreConfig) Size() (n int) {
	var l int
	_ = l
	if m.DType != 0 {
		n += 1 + sovStorage(uint64(m.DType))
	}
	l = len(m.Fqdn)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	l = len(m.ApiKey)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	l = len(m.Dpath)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 2 + l + sovStorage(uint64(l))
	}
	return n
}

func (m *Image) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	l = len(m.Sha256)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	if m.Iformat != 0 {
		n += 1 + sovStorage(uint64(m.Iformat))
	}
	if m.Siginfo != nil {
		l = m.Siginfo.Size()
		n += 1 + l + sovStorage(uint64(l))
	}
	l = len(m.DsId)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 2 + l + sovStorage(uint64(l))
	}
	return n
}

func (m *Drive) Size() (n int) {
	var l int
	_ = l
	if m.Image != nil {
		l = m.Image.Size()
		n += 1 + l + sovStorage(uint64(l))
	}
	if m.Maxsize != 0 {
		n += 1 + sovStorage(uint64(m.Maxsize))
	}
	if m.Readonly {
		n += 2
	}
	if m.Preserve {
		n += 2
	}
	if m.Drvtype != 0 {
		n += 1 + sovStorage(uint64(m.Drvtype))
	}
	if m.Target != 0 {
		n += 1 + sovStorage(uint64(m.Target))
	}
	return n
}

func sovStorage(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozStorage(x uint64) (n int) {
	return sovStorage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SignatureInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorage
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
			return fmt.Errorf("proto: SignatureInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignatureInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Intercertsurl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
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
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Intercertsurl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signercerturl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
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
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signercerturl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
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
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStorage
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
func (m *DatastoreConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorage
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
			return fmt.Errorf("proto: DatastoreConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DatastoreConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DType", wireType)
			}
			m.DType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DType |= (DsType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fqdn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
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
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fqdn = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
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
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApiKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
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
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dpath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
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
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Dpath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 100:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
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
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStorage
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
func (m *Image) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorage
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
			return fmt.Errorf("proto: Image: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Image: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
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
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sha256", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
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
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sha256 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Iformat", wireType)
			}
			m.Iformat = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Iformat |= (Format(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Siginfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
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
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Siginfo == nil {
				m.Siginfo = &SignatureInfo{}
			}
			if err := m.Siginfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DsId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
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
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DsId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 100:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
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
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStorage
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
func (m *Drive) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorage
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
			return fmt.Errorf("proto: Drive: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Drive: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Image", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
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
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Image == nil {
				m.Image = &Image{}
			}
			if err := m.Image.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Maxsize", wireType)
			}
			m.Maxsize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Maxsize |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Readonly", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Readonly = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Preserve", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Preserve = bool(v != 0)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Drvtype", wireType)
			}
			m.Drvtype = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Drvtype |= (DriveType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Target", wireType)
			}
			m.Target = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Target |= (Target(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStorage
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
func skipStorage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStorage
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
					return 0, ErrIntOverflowStorage
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
					return 0, ErrIntOverflowStorage
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
				return 0, ErrInvalidLengthStorage
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowStorage
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
				next, err := skipStorage(dAtA[start:])
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
	ErrInvalidLengthStorage = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStorage   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("storage.proto", fileDescriptorStorage) }

var fileDescriptorStorage = []byte{
	// 643 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x53, 0x41, 0x4f, 0x13, 0x41,
	0x14, 0x66, 0xdb, 0xee, 0x6e, 0xfb, 0xa0, 0x75, 0x32, 0x31, 0x66, 0x63, 0xa0, 0x20, 0xe1, 0x40,
	0x7a, 0x58, 0x92, 0x12, 0x35, 0xf1, 0xa6, 0xac, 0x4d, 0x1b, 0x82, 0xe0, 0x52, 0xd0, 0x18, 0x2f,
	0x43, 0x67, 0xba, 0x4c, 0xe8, 0xce, 0xd6, 0x99, 0x29, 0x58, 0x7e, 0x89, 0x27, 0x4f, 0xfe, 0x0e,
	0xcf, 0x1e, 0xfd, 0x09, 0x06, 0xff, 0x88, 0x99, 0x99, 0x6e, 0x09, 0xf1, 0xf6, 0xbe, 0xef, 0x7d,
	0xbb, 0xdf, 0xf7, 0xde, 0xcb, 0x40, 0x53, 0xe9, 0x42, 0x92, 0x8c, 0xc5, 0x53, 0x59, 0xe8, 0x62,
	0x7b, 0x0e, 0xcd, 0x53, 0x9e, 0x09, 0xa2, 0x67, 0x92, 0x0d, 0xc4, 0xb8, 0xc0, 0x3b, 0xd0, 0xe4,
	0x42, 0x33, 0x39, 0x62, 0x52, 0xab, 0x99, 0x9c, 0x44, 0xde, 0x96, 0xb7, 0xdb, 0x48, 0x1f, 0x92,
	0x46, 0xa5, 0x78, 0x26, 0x1c, 0x63, 0x54, 0x15, 0xa7, 0x7a, 0x40, 0xe2, 0x75, 0x68, 0xa8, 0xf2,
	0xe7, 0x51, 0x75, 0xcb, 0xdb, 0x5d, 0x4b, 0xef, 0x89, 0xed, 0xef, 0x1e, 0x3c, 0x4a, 0x88, 0x26,
	0x26, 0x10, 0x3b, 0x28, 0xc4, 0x98, 0x67, 0x78, 0x03, 0x7c, 0x3a, 0x9c, 0x4f, 0x99, 0x75, 0x6d,
	0x75, 0xc3, 0x38, 0x51, 0x06, 0xa6, 0x8e, 0xc5, 0x18, 0x6a, 0xe3, 0x2f, 0x54, 0x2c, 0xdc, 0x6c,
	0x8d, 0x9f, 0x40, 0x40, 0xa6, 0xfc, 0x90, 0xcd, 0xad, 0x43, 0x23, 0x5d, 0x20, 0xfc, 0x14, 0xea,
	0x53, 0xa2, 0xd4, 0x4d, 0x21, 0x69, 0x54, 0xb3, 0x9d, 0x25, 0xc6, 0x8f, 0xc1, 0xa7, 0x53, 0xa2,
	0x2f, 0x23, 0xdf, 0x36, 0x1c, 0xc0, 0x2d, 0xa8, 0x70, 0x1a, 0x51, 0x4b, 0x55, 0x38, 0xdd, 0xfe,
	0xe1, 0x81, 0x3f, 0xc8, 0x49, 0x66, 0x7d, 0x05, 0xc9, 0x59, 0xe9, 0x6b, 0x6a, 0xe3, 0xab, 0x2e,
	0x49, 0xf7, 0xf9, 0x8b, 0xd2, 0xd7, 0x21, 0xfc, 0x0c, 0x42, 0x3e, 0x2e, 0x64, 0x4e, 0xb4, 0xb5,
	0x35, 0x43, 0xf4, 0x2c, 0x4c, 0x4b, 0x1e, 0xef, 0x42, 0xa8, 0x78, 0xc6, 0xc5, 0xb8, 0xb0, 0x01,
	0x56, 0xbb, 0xad, 0xf8, 0xc1, 0x11, 0xd2, 0xb2, 0x6d, 0x8c, 0xa9, 0x1a, 0xd0, 0x28, 0x70, 0xc6,
	0xa6, 0xfe, 0x2f, 0xe6, 0x4f, 0x0f, 0xfc, 0x44, 0xf2, 0x6b, 0x86, 0xd7, 0xc1, 0xe7, 0x26, 0xaf,
	0xdd, 0xde, 0x6a, 0x37, 0x88, 0x6d, 0xfa, 0xd4, 0x91, 0x38, 0x82, 0x30, 0x27, 0x5f, 0x15, 0xbf,
	0x75, 0x73, 0x54, 0xd3, 0x12, 0x9a, 0x55, 0x49, 0x46, 0x68, 0x21, 0x26, 0x73, 0x1b, 0xa8, 0x9e,
	0x2e, 0xb1, 0x5d, 0xa3, 0x64, 0x8a, 0xc9, 0x6b, 0x66, 0x53, 0xd4, 0xd3, 0x25, 0xc6, 0x3b, 0x10,
	0x52, 0x79, 0xad, 0xcd, 0xbd, 0xea, 0x76, 0x54, 0x88, 0x6d, 0x10, 0x7b, 0xb2, 0xb2, 0x85, 0x37,
	0x21, 0xd0, 0x44, 0x66, 0x4c, 0x47, 0x8d, 0xc5, 0x3e, 0x86, 0x16, 0xa6, 0x0b, 0xba, 0xd3, 0x83,
	0xc0, 0x9d, 0x19, 0x37, 0xa1, 0x91, 0xa8, 0x33, 0x71, 0x25, 0x8a, 0x1b, 0x81, 0x56, 0x30, 0x98,
	0x46, 0x5f, 0xeb, 0x29, 0xf2, 0xf0, 0x2a, 0x84, 0xae, 0x56, 0xa8, 0x82, 0xeb, 0x50, 0x4b, 0xd4,
	0xe9, 0x3e, 0xaa, 0x3a, 0xc9, 0x69, 0x6f, 0x78, 0x82, 0x6a, 0x9d, 0xcf, 0x10, 0xb8, 0x4d, 0xe3,
	0x16, 0x40, 0x2f, 0xd7, 0xf7, 0x3f, 0x0a, 0xa1, 0x9a, 0x92, 0x1b, 0xe4, 0x99, 0x0f, 0xdf, 0x1f,
	0x1c, 0x7f, 0x40, 0x15, 0xdc, 0x00, 0xdf, 0x54, 0x5d, 0x54, 0x35, 0xdd, 0xf3, 0x7e, 0x82, 0x6a,
	0xa6, 0x7b, 0x7e, 0x94, 0x1c, 0x22, 0xdf, 0x50, 0xc7, 0xe7, 0xaf, 0x51, 0x60, 0xa9, 0x7e, 0xf2,
	0x11, 0x85, 0x9d, 0x97, 0x10, 0xb8, 0xdc, 0xd6, 0x9d, 0xab, 0x2b, 0x17, 0xf0, 0x90, 0x49, 0xc1,
	0x26, 0xc8, 0x33, 0xf5, 0x40, 0x70, 0x2d, 0x29, 0xaa, 0x98, 0xb0, 0x29, 0xc9, 0xad, 0xa8, 0xda,
	0x79, 0x05, 0x8d, 0xe5, 0x56, 0x30, 0x82, 0xb5, 0x33, 0x31, 0x9a, 0x10, 0xa5, 0xf8, 0x98, 0x33,
	0x8a, 0x56, 0x4c, 0x90, 0x83, 0x24, 0x3d, 0x3e, 0x42, 0x9e, 0x71, 0xed, 0x27, 0x09, 0xaa, 0x98,
	0xe2, 0xdd, 0xdb, 0x21, 0xaa, 0xbe, 0x39, 0xf9, 0x75, 0xd7, 0xf6, 0x7e, 0xdf, 0xb5, 0xbd, 0x3f,
	0x77, 0x6d, 0xef, 0xdb, 0xdf, 0xf6, 0x0a, 0x6c, 0x8e, 0x8a, 0x3c, 0xbe, 0x65, 0x94, 0x51, 0x12,
	0x8f, 0x26, 0xc5, 0x8c, 0xc6, 0x33, 0x73, 0x0c, 0x3e, 0x5a, 0xbc, 0xe8, 0x4f, 0x1b, 0x19, 0xd7,
	0x97, 0xb3, 0x8b, 0x78, 0x54, 0xe4, 0x7b, 0x4e, 0xb7, 0x47, 0xa6, 0x7c, 0xef, 0x76, 0x64, 0x5f,
	0xd8, 0x45, 0x60, 0x55, 0xfb, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x4d, 0xee, 0x38, 0x08,
	0x04, 0x00, 0x00,
}
