// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.3
// source: server/proto/ingredient_service/v1/ingredient_service_v1.proto

package ingredientpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Ingredient message to represent the ingredient data.
type Ingredient struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`     // Unique identifier for the ingredient.
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` // Name of the ingredient (e.g. Butter).
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Ingredient) Reset() {
	*x = Ingredient{}
	mi := &file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Ingredient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ingredient) ProtoMessage() {}

func (x *Ingredient) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ingredient.ProtoReflect.Descriptor instead.
func (*Ingredient) Descriptor() ([]byte, []int) {
	return file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_rawDescGZIP(), []int{0}
}

func (x *Ingredient) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Ingredient) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request message for creating an ingredient.
type CreateIngredientRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ingredient    *Ingredient            `protobuf:"bytes,1,opt,name=ingredient,proto3" json:"ingredient,omitempty"` // The ingredient to create.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateIngredientRequest) Reset() {
	*x = CreateIngredientRequest{}
	mi := &file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateIngredientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateIngredientRequest) ProtoMessage() {}

func (x *CreateIngredientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateIngredientRequest.ProtoReflect.Descriptor instead.
func (*CreateIngredientRequest) Descriptor() ([]byte, []int) {
	return file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_rawDescGZIP(), []int{1}
}

func (x *CreateIngredientRequest) GetIngredient() *Ingredient {
	if x != nil {
		return x.Ingredient
	}
	return nil
}

// Request message for fetching an ingredient by its ID.
type GetIngredientRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // ID of the ingredient to fetch.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetIngredientRequest) Reset() {
	*x = GetIngredientRequest{}
	mi := &file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetIngredientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIngredientRequest) ProtoMessage() {}

func (x *GetIngredientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIngredientRequest.ProtoReflect.Descriptor instead.
func (*GetIngredientRequest) Descriptor() ([]byte, []int) {
	return file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_rawDescGZIP(), []int{2}
}

func (x *GetIngredientRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Request message for listing ingredients by its ID.
type ListIngredientsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Limit         int32                  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`   // Optional limit for pagination.
	Offset        int32                  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"` // Optional offset for pagination.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListIngredientsRequest) Reset() {
	*x = ListIngredientsRequest{}
	mi := &file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListIngredientsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListIngredientsRequest) ProtoMessage() {}

func (x *ListIngredientsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListIngredientsRequest.ProtoReflect.Descriptor instead.
func (*ListIngredientsRequest) Descriptor() ([]byte, []int) {
	return file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_rawDescGZIP(), []int{3}
}

func (x *ListIngredientsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListIngredientsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

// Response message for returning the ingredients list.
type ListIngredientsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ingredients   []*Ingredient          `protobuf:"bytes,1,rep,name=ingredients,proto3" json:"ingredients,omitempty"` // List of ingredients.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListIngredientsResponse) Reset() {
	*x = ListIngredientsResponse{}
	mi := &file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListIngredientsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListIngredientsResponse) ProtoMessage() {}

func (x *ListIngredientsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListIngredientsResponse.ProtoReflect.Descriptor instead.
func (*ListIngredientsResponse) Descriptor() ([]byte, []int) {
	return file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_rawDescGZIP(), []int{4}
}

func (x *ListIngredientsResponse) GetIngredients() []*Ingredient {
	if x != nil {
		return x.Ingredients
	}
	return nil
}

// Request message for updating an ingredient by its ID.
type UpdateIngredientRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`     // Ingredient ID to update.
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` // New name for the ingredient.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateIngredientRequest) Reset() {
	*x = UpdateIngredientRequest{}
	mi := &file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateIngredientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateIngredientRequest) ProtoMessage() {}

func (x *UpdateIngredientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateIngredientRequest.ProtoReflect.Descriptor instead.
func (*UpdateIngredientRequest) Descriptor() ([]byte, []int) {
	return file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateIngredientRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateIngredientRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request message for deleting an ingredient by its ID.
type DeleteIngredientRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // Ingredient ID to delete.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteIngredientRequest) Reset() {
	*x = DeleteIngredientRequest{}
	mi := &file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteIngredientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteIngredientRequest) ProtoMessage() {}

func (x *DeleteIngredientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteIngredientRequest.ProtoReflect.Descriptor instead.
func (*DeleteIngredientRequest) Descriptor() ([]byte, []int) {
	return file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteIngredientRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Response message for an ingredient.
type IngredientResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ingredient    *Ingredient            `protobuf:"bytes,1,opt,name=ingredient,proto3" json:"ingredient,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IngredientResponse) Reset() {
	*x = IngredientResponse{}
	mi := &file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IngredientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngredientResponse) ProtoMessage() {}

func (x *IngredientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngredientResponse.ProtoReflect.Descriptor instead.
func (*IngredientResponse) Descriptor() ([]byte, []int) {
	return file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_rawDescGZIP(), []int{7}
}

func (x *IngredientResponse) GetIngredient() *Ingredient {
	if x != nil {
		return x.Ingredient
	}
	return nil
}

// Empty response message
type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_rawDescGZIP(), []int{8}
}

var File_server_proto_ingredient_service_v1_ingredient_service_v1_proto protoreflect.FileDescriptor

var file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_rawDesc = string([]byte{
	0x0a, 0x3e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69,
	0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x15, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x30, 0x0a, 0x0a, 0x49, 0x6e, 0x67, 0x72, 0x65,
	0x64, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5c, 0x0a, 0x17, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x0a, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65,
	0x64, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x69, 0x6e, 0x67,
	0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x26, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x49, 0x6e,
	0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x46, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x5e, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x43, 0x0a, 0x0b, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x69, 0x6e, 0x67, 0x72,
	0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x3d, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x57, 0x0a, 0x12, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x69, 0x6e, 0x67, 0x72, 0x65,
	0x64, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x69, 0x6e,
	0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x0a,
	0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x32, 0xae, 0x04, 0x0a, 0x11, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6d, 0x0a, 0x10, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x2e,
	0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x67, 0x72,
	0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x67, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x2b, 0x2e, 0x69, 0x6e, 0x67, 0x72,
	0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x70, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x2d, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x6d, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x67,
	0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x60, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x67, 0x72,
	0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x42, 0x3b, 0x5a, 0x39, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69,
	0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_rawDescOnce sync.Once
	file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_rawDescData []byte
)

func file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_rawDescGZIP() []byte {
	file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_rawDescOnce.Do(func() {
		file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_rawDesc), len(file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_rawDesc)))
	})
	return file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_rawDescData
}

var file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_goTypes = []any{
	(*Ingredient)(nil),              // 0: ingredient_service.v1.Ingredient
	(*CreateIngredientRequest)(nil), // 1: ingredient_service.v1.CreateIngredientRequest
	(*GetIngredientRequest)(nil),    // 2: ingredient_service.v1.GetIngredientRequest
	(*ListIngredientsRequest)(nil),  // 3: ingredient_service.v1.ListIngredientsRequest
	(*ListIngredientsResponse)(nil), // 4: ingredient_service.v1.ListIngredientsResponse
	(*UpdateIngredientRequest)(nil), // 5: ingredient_service.v1.UpdateIngredientRequest
	(*DeleteIngredientRequest)(nil), // 6: ingredient_service.v1.DeleteIngredientRequest
	(*IngredientResponse)(nil),      // 7: ingredient_service.v1.IngredientResponse
	(*Empty)(nil),                   // 8: ingredient_service.v1.Empty
}
var file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_depIdxs = []int32{
	0, // 0: ingredient_service.v1.CreateIngredientRequest.ingredient:type_name -> ingredient_service.v1.Ingredient
	0, // 1: ingredient_service.v1.ListIngredientsResponse.ingredients:type_name -> ingredient_service.v1.Ingredient
	0, // 2: ingredient_service.v1.IngredientResponse.ingredient:type_name -> ingredient_service.v1.Ingredient
	1, // 3: ingredient_service.v1.IngredientService.CreateIngredient:input_type -> ingredient_service.v1.CreateIngredientRequest
	2, // 4: ingredient_service.v1.IngredientService.GetIngredient:input_type -> ingredient_service.v1.GetIngredientRequest
	3, // 5: ingredient_service.v1.IngredientService.ListIngredients:input_type -> ingredient_service.v1.ListIngredientsRequest
	5, // 6: ingredient_service.v1.IngredientService.UpdateIngredient:input_type -> ingredient_service.v1.UpdateIngredientRequest
	6, // 7: ingredient_service.v1.IngredientService.DeleteIngredient:input_type -> ingredient_service.v1.DeleteIngredientRequest
	7, // 8: ingredient_service.v1.IngredientService.CreateIngredient:output_type -> ingredient_service.v1.IngredientResponse
	7, // 9: ingredient_service.v1.IngredientService.GetIngredient:output_type -> ingredient_service.v1.IngredientResponse
	4, // 10: ingredient_service.v1.IngredientService.ListIngredients:output_type -> ingredient_service.v1.ListIngredientsResponse
	7, // 11: ingredient_service.v1.IngredientService.UpdateIngredient:output_type -> ingredient_service.v1.IngredientResponse
	8, // 12: ingredient_service.v1.IngredientService.DeleteIngredient:output_type -> ingredient_service.v1.Empty
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_init() }
func file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_init() {
	if File_server_proto_ingredient_service_v1_ingredient_service_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_rawDesc), len(file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_goTypes,
		DependencyIndexes: file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_depIdxs,
		MessageInfos:      file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_msgTypes,
	}.Build()
	File_server_proto_ingredient_service_v1_ingredient_service_v1_proto = out.File
	file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_goTypes = nil
	file_server_proto_ingredient_service_v1_ingredient_service_v1_proto_depIdxs = nil
}
