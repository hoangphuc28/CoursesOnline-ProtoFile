// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: Cart/cart.proto

package Cart

import (
	Error "github.com/hoangphuc28/CoursesOnline-ProtoFile/Error"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url    string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Width  string `protobuf:"bytes,2,opt,name=width,proto3" json:"width,omitempty"`
	Height string `protobuf:"bytes,3,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Cart_cart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_Cart_cart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_Cart_cart_proto_rawDescGZIP(), []int{0}
}

func (x *Image) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Image) GetWidth() string {
	if x != nil {
		return x.Width
	}
	return ""
}

func (x *Image) GetHeight() string {
	if x != nil {
		return x.Height
	}
	return ""
}

type Course struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title        string  `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description  string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Level        string  `protobuf:"bytes,4,opt,name=level,proto3" json:"level,omitempty"`
	Price        string  `protobuf:"bytes,6,opt,name=price,proto3" json:"price,omitempty"`
	Discount     float32 `protobuf:"fixed32,7,opt,name=discount,proto3" json:"discount,omitempty"`
	Currency     string  `protobuf:"bytes,8,opt,name=currency,proto3" json:"currency,omitempty"`
	Duration     string  `protobuf:"bytes,9,opt,name=duration,proto3" json:"duration,omitempty"`
	Status       string  `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	Rating       float32 `protobuf:"fixed32,11,opt,name=rating,proto3" json:"rating,omitempty"`
	Thumbnail    *Image  `protobuf:"bytes,12,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	InstructorId string  `protobuf:"bytes,14,opt,name=instructorId,proto3" json:"instructorId,omitempty"`
}

func (x *Course) Reset() {
	*x = Course{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Cart_cart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Course) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Course) ProtoMessage() {}

func (x *Course) ProtoReflect() protoreflect.Message {
	mi := &file_Cart_cart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Course.ProtoReflect.Descriptor instead.
func (*Course) Descriptor() ([]byte, []int) {
	return file_Cart_cart_proto_rawDescGZIP(), []int{1}
}

func (x *Course) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Course) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Course) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Course) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *Course) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *Course) GetDiscount() float32 {
	if x != nil {
		return x.Discount
	}
	return 0
}

func (x *Course) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Course) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

func (x *Course) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Course) GetRating() float32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Course) GetThumbnail() *Image {
	if x != nil {
		return x.Thumbnail
	}
	return nil
}

func (x *Course) GetInstructorId() string {
	if x != nil {
		return x.InstructorId
	}
	return ""
}

type Cart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Courses    []*Course `protobuf:"bytes,2,rep,name=courses,proto3" json:"courses,omitempty"`
	TotalPrice string    `protobuf:"bytes,3,opt,name=totalPrice,proto3" json:"totalPrice,omitempty"`
	Currency   string    `protobuf:"bytes,4,opt,name=currency,proto3" json:"currency,omitempty"`
}

func (x *Cart) Reset() {
	*x = Cart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Cart_cart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cart) ProtoMessage() {}

func (x *Cart) ProtoReflect() protoreflect.Message {
	mi := &file_Cart_cart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cart.ProtoReflect.Descriptor instead.
func (*Cart) Descriptor() ([]byte, []int) {
	return file_Cart_cart_proto_rawDescGZIP(), []int{2}
}

func (x *Cart) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Cart) GetCourses() []*Course {
	if x != nil {
		return x.Courses
	}
	return nil
}

func (x *Cart) GetTotalPrice() string {
	if x != nil {
		return x.TotalPrice
	}
	return ""
}

func (x *Cart) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

type GetCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCartRequest) Reset() {
	*x = GetCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Cart_cart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartRequest) ProtoMessage() {}

func (x *GetCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Cart_cart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartRequest.ProtoReflect.Descriptor instead.
func (*GetCartRequest) Descriptor() ([]byte, []int) {
	return file_Cart_cart_proto_rawDescGZIP(), []int{3}
}

func (x *GetCartRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cart        *Cart                `protobuf:"bytes,1,opt,name=cart,proto3" json:"cart,omitempty"`
	TotalCourse string               `protobuf:"bytes,2,opt,name=totalCourse,proto3" json:"totalCourse,omitempty"`
	Error       *Error.ErrorResponse `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GetCartResponse) Reset() {
	*x = GetCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Cart_cart_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartResponse) ProtoMessage() {}

func (x *GetCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Cart_cart_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartResponse.ProtoReflect.Descriptor instead.
func (*GetCartResponse) Descriptor() ([]byte, []int) {
	return file_Cart_cart_proto_rawDescGZIP(), []int{4}
}

func (x *GetCartResponse) GetCart() *Cart {
	if x != nil {
		return x.Cart
	}
	return nil
}

func (x *GetCartResponse) GetTotalCourse() string {
	if x != nil {
		return x.TotalCourse
	}
	return ""
}

func (x *GetCartResponse) GetError() *Error.ErrorResponse {
	if x != nil {
		return x.Error
	}
	return nil
}

type CartItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartId   string `protobuf:"bytes,1,opt,name=cartId,proto3" json:"cartId,omitempty"`
	CourseId string `protobuf:"bytes,2,opt,name=courseId,proto3" json:"courseId,omitempty"`
}

func (x *CartItemRequest) Reset() {
	*x = CartItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Cart_cart_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItemRequest) ProtoMessage() {}

func (x *CartItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Cart_cart_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItemRequest.ProtoReflect.Descriptor instead.
func (*CartItemRequest) Descriptor() ([]byte, []int) {
	return file_Cart_cart_proto_rawDescGZIP(), []int{5}
}

func (x *CartItemRequest) GetCartId() string {
	if x != nil {
		return x.CartId
	}
	return ""
}

func (x *CartItemRequest) GetCourseId() string {
	if x != nil {
		return x.CourseId
	}
	return ""
}

type CartItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *Error.ErrorResponse `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CartItemResponse) Reset() {
	*x = CartItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Cart_cart_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItemResponse) ProtoMessage() {}

func (x *CartItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Cart_cart_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItemResponse.ProtoReflect.Descriptor instead.
func (*CartItemResponse) Descriptor() ([]byte, []int) {
	return file_Cart_cart_proto_rawDescGZIP(), []int{6}
}

func (x *CartItemResponse) GetError() *Error.ErrorResponse {
	if x != nil {
		return x.Error
	}
	return nil
}

type ErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ErrorResponse) Reset() {
	*x = ErrorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Cart_cart_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorResponse) ProtoMessage() {}

func (x *ErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Cart_cart_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorResponse.ProtoReflect.Descriptor instead.
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return file_Cart_cart_proto_rawDescGZIP(), []int{7}
}

func (x *ErrorResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ErrorResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ResetCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartId string `protobuf:"bytes,1,opt,name=cartId,proto3" json:"cartId,omitempty"`
}

func (x *ResetCartRequest) Reset() {
	*x = ResetCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Cart_cart_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetCartRequest) ProtoMessage() {}

func (x *ResetCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Cart_cart_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetCartRequest.ProtoReflect.Descriptor instead.
func (*ResetCartRequest) Descriptor() ([]byte, []int) {
	return file_Cart_cart_proto_rawDescGZIP(), []int{8}
}

func (x *ResetCartRequest) GetCartId() string {
	if x != nil {
		return x.CartId
	}
	return ""
}

type ResetCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *Error.ErrorResponse `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ResetCartResponse) Reset() {
	*x = ResetCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Cart_cart_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetCartResponse) ProtoMessage() {}

func (x *ResetCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Cart_cart_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetCartResponse.ProtoReflect.Descriptor instead.
func (*ResetCartResponse) Descriptor() ([]byte, []int) {
	return file_Cart_cart_proto_rawDescGZIP(), []int{9}
}

func (x *ResetCartResponse) GetError() *Error.ErrorResponse {
	if x != nil {
		return x.Error
	}
	return nil
}

type CreateCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *CreateCartRequest) Reset() {
	*x = CreateCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Cart_cart_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCartRequest) ProtoMessage() {}

func (x *CreateCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Cart_cart_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCartRequest.ProtoReflect.Descriptor instead.
func (*CreateCartRequest) Descriptor() ([]byte, []int) {
	return file_Cart_cart_proto_rawDescGZIP(), []int{10}
}

func (x *CreateCartRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type CreateCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *Error.ErrorResponse `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CreateCartResponse) Reset() {
	*x = CreateCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Cart_cart_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCartResponse) ProtoMessage() {}

func (x *CreateCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Cart_cart_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCartResponse.ProtoReflect.Descriptor instead.
func (*CreateCartResponse) Descriptor() ([]byte, []int) {
	return file_Cart_cart_proto_rawDescGZIP(), []int{11}
}

func (x *CreateCartResponse) GetError() *Error.ErrorResponse {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_Cart_cart_proto protoreflect.FileDescriptor

var file_Cart_cart_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x43, 0x61, 0x72, 0x74, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x63, 0x61, 0x72, 0x74, 0x1a, 0x11, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x2f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x05, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x22, 0xcf, 0x02, 0x0a, 0x06, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x29, 0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e,
	0x61, 0x69, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x63, 0x61, 0x72, 0x74,
	0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69,
	0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x49,
	0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x7a, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a,
	0x07, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x07, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x7d, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x74,
	0x52, 0x04, 0x63, 0x61, 0x72, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x72, 0x72, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x45, 0x0a, 0x0f, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x72, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x10, 0x43, 0x61, 0x72,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65,
	0x72, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x3d, 0x0a, 0x0d, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2a, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x65, 0x74, 0x43,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61,
	0x72, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x72, 0x74,
	0x49, 0x64, 0x22, 0x3d, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x72, 0x72, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x22, 0x2b, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3e,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x72, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xbb,
	0x02, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x12, 0x14, 0x2e, 0x63, 0x61, 0x72, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x15, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e,
	0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3b, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x15,
	0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x61, 0x72,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a,
	0x09, 0x52, 0x65, 0x73, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x12, 0x16, 0x2e, 0x63, 0x61, 0x72,
	0x74, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x43,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x12, 0x17, 0x2e, 0x63, 0x61, 0x72, 0x74,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x35, 0x5a, 0x33,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x6f, 0x61, 0x6e, 0x67,
	0x70, 0x68, 0x75, 0x63, 0x32, 0x38, 0x2f, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x4f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x2d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x2f, 0x43,
	0x61, 0x72, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Cart_cart_proto_rawDescOnce sync.Once
	file_Cart_cart_proto_rawDescData = file_Cart_cart_proto_rawDesc
)

func file_Cart_cart_proto_rawDescGZIP() []byte {
	file_Cart_cart_proto_rawDescOnce.Do(func() {
		file_Cart_cart_proto_rawDescData = protoimpl.X.CompressGZIP(file_Cart_cart_proto_rawDescData)
	})
	return file_Cart_cart_proto_rawDescData
}

var file_Cart_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_Cart_cart_proto_goTypes = []interface{}{
	(*Image)(nil),               // 0: cart.Image
	(*Course)(nil),              // 1: cart.Course
	(*Cart)(nil),                // 2: cart.Cart
	(*GetCartRequest)(nil),      // 3: cart.GetCartRequest
	(*GetCartResponse)(nil),     // 4: cart.GetCartResponse
	(*CartItemRequest)(nil),     // 5: cart.CartItemRequest
	(*CartItemResponse)(nil),    // 6: cart.CartItemResponse
	(*ErrorResponse)(nil),       // 7: cart.ErrorResponse
	(*ResetCartRequest)(nil),    // 8: cart.ResetCartRequest
	(*ResetCartResponse)(nil),   // 9: cart.ResetCartResponse
	(*CreateCartRequest)(nil),   // 10: cart.CreateCartRequest
	(*CreateCartResponse)(nil),  // 11: cart.CreateCartResponse
	(*Error.ErrorResponse)(nil), // 12: err.ErrorResponse
}
var file_Cart_cart_proto_depIdxs = []int32{
	0,  // 0: cart.Course.thumbnail:type_name -> cart.Image
	1,  // 1: cart.Cart.courses:type_name -> cart.Course
	2,  // 2: cart.GetCartResponse.cart:type_name -> cart.Cart
	12, // 3: cart.GetCartResponse.error:type_name -> err.ErrorResponse
	12, // 4: cart.CartItemResponse.error:type_name -> err.ErrorResponse
	12, // 5: cart.ResetCartResponse.error:type_name -> err.ErrorResponse
	12, // 6: cart.CreateCartResponse.error:type_name -> err.ErrorResponse
	3,  // 7: cart.CartService.GetCart:input_type -> cart.GetCartRequest
	5,  // 8: cart.CartService.AddItem:input_type -> cart.CartItemRequest
	5,  // 9: cart.CartService.RemoveItem:input_type -> cart.CartItemRequest
	8,  // 10: cart.CartService.ResetCart:input_type -> cart.ResetCartRequest
	10, // 11: cart.CartService.CreateCart:input_type -> cart.CreateCartRequest
	4,  // 12: cart.CartService.GetCart:output_type -> cart.GetCartResponse
	6,  // 13: cart.CartService.AddItem:output_type -> cart.CartItemResponse
	6,  // 14: cart.CartService.RemoveItem:output_type -> cart.CartItemResponse
	9,  // 15: cart.CartService.ResetCart:output_type -> cart.ResetCartResponse
	11, // 16: cart.CartService.CreateCart:output_type -> cart.CreateCartResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_Cart_cart_proto_init() }
func file_Cart_cart_proto_init() {
	if File_Cart_cart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Cart_cart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Cart_cart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Course); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Cart_cart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cart); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Cart_cart_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Cart_cart_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Cart_cart_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItemRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Cart_cart_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItemResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Cart_cart_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Cart_cart_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetCartRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Cart_cart_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetCartResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Cart_cart_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCartRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Cart_cart_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCartResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Cart_cart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Cart_cart_proto_goTypes,
		DependencyIndexes: file_Cart_cart_proto_depIdxs,
		MessageInfos:      file_Cart_cart_proto_msgTypes,
	}.Build()
	File_Cart_cart_proto = out.File
	file_Cart_cart_proto_rawDesc = nil
	file_Cart_cart_proto_goTypes = nil
	file_Cart_cart_proto_depIdxs = nil
}
