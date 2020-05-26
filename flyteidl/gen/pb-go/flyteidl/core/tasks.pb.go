// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/core/tasks.proto

package core

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

// Known resource names.
type Resources_ResourceName int32

const (
	Resources_UNKNOWN Resources_ResourceName = 0
	Resources_CPU     Resources_ResourceName = 1
	Resources_GPU     Resources_ResourceName = 2
	Resources_MEMORY  Resources_ResourceName = 3
	Resources_STORAGE Resources_ResourceName = 4
)

var Resources_ResourceName_name = map[int32]string{
	0: "UNKNOWN",
	1: "CPU",
	2: "GPU",
	3: "MEMORY",
	4: "STORAGE",
}

var Resources_ResourceName_value = map[string]int32{
	"UNKNOWN": 0,
	"CPU":     1,
	"GPU":     2,
	"MEMORY":  3,
	"STORAGE": 4,
}

func (x Resources_ResourceName) String() string {
	return proto.EnumName(Resources_ResourceName_name, int32(x))
}

func (Resources_ResourceName) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bd8423ba58d6ed80, []int{0, 0}
}

type RuntimeMetadata_RuntimeType int32

const (
	RuntimeMetadata_OTHER     RuntimeMetadata_RuntimeType = 0
	RuntimeMetadata_FLYTE_SDK RuntimeMetadata_RuntimeType = 1
)

var RuntimeMetadata_RuntimeType_name = map[int32]string{
	0: "OTHER",
	1: "FLYTE_SDK",
}

var RuntimeMetadata_RuntimeType_value = map[string]int32{
	"OTHER":     0,
	"FLYTE_SDK": 1,
}

func (x RuntimeMetadata_RuntimeType) String() string {
	return proto.EnumName(RuntimeMetadata_RuntimeType_name, int32(x))
}

func (RuntimeMetadata_RuntimeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bd8423ba58d6ed80, []int{1, 0}
}

// Mode to use for downloading
type IOStrategy_DownloadMode int32

const (
	// All data will be downloaded before the main container is executed
	IOStrategy_DOWNLOAD_EAGER IOStrategy_DownloadMode = 0
	// Data will be downloaded as a stream and an End-Of-Stream marker will be written to indicate all data has been downloaded. Refer to protocol for details
	IOStrategy_DOWNLOAD_STREAM IOStrategy_DownloadMode = 1
	// Large objects (offloaded) will not be downloaded
	IOStrategy_DO_NOT_DOWNLOAD IOStrategy_DownloadMode = 2
)

var IOStrategy_DownloadMode_name = map[int32]string{
	0: "DOWNLOAD_EAGER",
	1: "DOWNLOAD_STREAM",
	2: "DO_NOT_DOWNLOAD",
}

var IOStrategy_DownloadMode_value = map[string]int32{
	"DOWNLOAD_EAGER":  0,
	"DOWNLOAD_STREAM": 1,
	"DO_NOT_DOWNLOAD": 2,
}

func (x IOStrategy_DownloadMode) String() string {
	return proto.EnumName(IOStrategy_DownloadMode_name, int32(x))
}

func (IOStrategy_DownloadMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bd8423ba58d6ed80, []int{6, 0}
}

// Mode to use for uploading
type IOStrategy_UploadMode int32

const (
	// All data will be uploaded after the main container exits
	IOStrategy_UPLOAD_ON_EXIT IOStrategy_UploadMode = 0
	// Data will be uploaded as it appears. Refer to protocol specification for details
	IOStrategy_UPLOAD_EAGER IOStrategy_UploadMode = 1
	// Data will not be uploaded, only references will be written
	IOStrategy_DO_NOT_UPLOAD IOStrategy_UploadMode = 2
)

var IOStrategy_UploadMode_name = map[int32]string{
	0: "UPLOAD_ON_EXIT",
	1: "UPLOAD_EAGER",
	2: "DO_NOT_UPLOAD",
}

var IOStrategy_UploadMode_value = map[string]int32{
	"UPLOAD_ON_EXIT": 0,
	"UPLOAD_EAGER":   1,
	"DO_NOT_UPLOAD":  2,
}

func (x IOStrategy_UploadMode) String() string {
	return proto.EnumName(IOStrategy_UploadMode_name, int32(x))
}

func (IOStrategy_UploadMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bd8423ba58d6ed80, []int{6, 1}
}

// LiteralMapFormat decides the encoding format in which the input metadata should be made available to the containers.
// If the user has access to the protocol buffer definitions, it is recommended to use the PROTO format.
// JSON and YAML do not need any protobuf definitions to read it
// All remote references in core.LiteralMap are replaced with local filesystem references (the data is downloaded to local filesystem)
type DataLoadingConfig_LiteralMapFormat int32

const (
	// JSON / YAML for the metadata (which contains inlined primitive values). The representation is inline with the standard json specification as specified - https://www.json.org/json-en.html
	DataLoadingConfig_JSON DataLoadingConfig_LiteralMapFormat = 0
	DataLoadingConfig_YAML DataLoadingConfig_LiteralMapFormat = 1
	// Proto is a serialized binary of `core.LiteralMap` defined in flyteidl/core
	DataLoadingConfig_PROTO DataLoadingConfig_LiteralMapFormat = 2
)

var DataLoadingConfig_LiteralMapFormat_name = map[int32]string{
	0: "JSON",
	1: "YAML",
	2: "PROTO",
}

var DataLoadingConfig_LiteralMapFormat_value = map[string]int32{
	"JSON":  0,
	"YAML":  1,
	"PROTO": 2,
}

func (x DataLoadingConfig_LiteralMapFormat) String() string {
	return proto.EnumName(DataLoadingConfig_LiteralMapFormat_name, int32(x))
}

func (DataLoadingConfig_LiteralMapFormat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bd8423ba58d6ed80, []int{7, 0}
}

// A customizable interface to convey resources requested for a container. This can be interpretted differently for different
// container engines.
type Resources struct {
	// The desired set of resources requested. ResourceNames must be unique within the list.
	Requests []*Resources_ResourceEntry `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests,omitempty"`
	// Defines a set of bounds (e.g. min/max) within which the task can reliably run. ResourceNames must be unique
	// within the list.
	Limits               []*Resources_ResourceEntry `protobuf:"bytes,2,rep,name=limits,proto3" json:"limits,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Resources) Reset()         { *m = Resources{} }
func (m *Resources) String() string { return proto.CompactTextString(m) }
func (*Resources) ProtoMessage()    {}
func (*Resources) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd8423ba58d6ed80, []int{0}
}

func (m *Resources) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resources.Unmarshal(m, b)
}
func (m *Resources) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resources.Marshal(b, m, deterministic)
}
func (m *Resources) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resources.Merge(m, src)
}
func (m *Resources) XXX_Size() int {
	return xxx_messageInfo_Resources.Size(m)
}
func (m *Resources) XXX_DiscardUnknown() {
	xxx_messageInfo_Resources.DiscardUnknown(m)
}

var xxx_messageInfo_Resources proto.InternalMessageInfo

func (m *Resources) GetRequests() []*Resources_ResourceEntry {
	if m != nil {
		return m.Requests
	}
	return nil
}

func (m *Resources) GetLimits() []*Resources_ResourceEntry {
	if m != nil {
		return m.Limits
	}
	return nil
}

// Encapsulates a resource name and value.
type Resources_ResourceEntry struct {
	// Resource name.
	Name Resources_ResourceName `protobuf:"varint,1,opt,name=name,proto3,enum=flyteidl.core.Resources_ResourceName" json:"name,omitempty"`
	// Value must be a valid k8s quantity. See
	// https://github.com/kubernetes/apimachinery/blob/master/pkg/api/resource/quantity.go#L30-L80
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resources_ResourceEntry) Reset()         { *m = Resources_ResourceEntry{} }
func (m *Resources_ResourceEntry) String() string { return proto.CompactTextString(m) }
func (*Resources_ResourceEntry) ProtoMessage()    {}
func (*Resources_ResourceEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd8423ba58d6ed80, []int{0, 0}
}

func (m *Resources_ResourceEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resources_ResourceEntry.Unmarshal(m, b)
}
func (m *Resources_ResourceEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resources_ResourceEntry.Marshal(b, m, deterministic)
}
func (m *Resources_ResourceEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resources_ResourceEntry.Merge(m, src)
}
func (m *Resources_ResourceEntry) XXX_Size() int {
	return xxx_messageInfo_Resources_ResourceEntry.Size(m)
}
func (m *Resources_ResourceEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_Resources_ResourceEntry.DiscardUnknown(m)
}

var xxx_messageInfo_Resources_ResourceEntry proto.InternalMessageInfo

func (m *Resources_ResourceEntry) GetName() Resources_ResourceName {
	if m != nil {
		return m.Name
	}
	return Resources_UNKNOWN
}

func (m *Resources_ResourceEntry) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Runtime information. This is losely defined to allow for extensibility.
type RuntimeMetadata struct {
	// Type of runtime.
	Type RuntimeMetadata_RuntimeType `protobuf:"varint,1,opt,name=type,proto3,enum=flyteidl.core.RuntimeMetadata_RuntimeType" json:"type,omitempty"`
	// Version of the runtime. All versions should be backward compatible. However, certain cases call for version
	// checks to ensure tighter validation or setting expectations.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	//+optional It can be used to provide extra information about the runtime (e.g. python, golang... etc.).
	Flavor               string   `protobuf:"bytes,3,opt,name=flavor,proto3" json:"flavor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RuntimeMetadata) Reset()         { *m = RuntimeMetadata{} }
func (m *RuntimeMetadata) String() string { return proto.CompactTextString(m) }
func (*RuntimeMetadata) ProtoMessage()    {}
func (*RuntimeMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd8423ba58d6ed80, []int{1}
}

func (m *RuntimeMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RuntimeMetadata.Unmarshal(m, b)
}
func (m *RuntimeMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RuntimeMetadata.Marshal(b, m, deterministic)
}
func (m *RuntimeMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuntimeMetadata.Merge(m, src)
}
func (m *RuntimeMetadata) XXX_Size() int {
	return xxx_messageInfo_RuntimeMetadata.Size(m)
}
func (m *RuntimeMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_RuntimeMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_RuntimeMetadata proto.InternalMessageInfo

func (m *RuntimeMetadata) GetType() RuntimeMetadata_RuntimeType {
	if m != nil {
		return m.Type
	}
	return RuntimeMetadata_OTHER
}

func (m *RuntimeMetadata) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *RuntimeMetadata) GetFlavor() string {
	if m != nil {
		return m.Flavor
	}
	return ""
}

// Task Metadata
type TaskMetadata struct {
	// Indicates whether the system should attempt to lookup this task's output to avoid duplication of work.
	Discoverable bool `protobuf:"varint,1,opt,name=discoverable,proto3" json:"discoverable,omitempty"`
	// Runtime information about the task.
	Runtime *RuntimeMetadata `protobuf:"bytes,2,opt,name=runtime,proto3" json:"runtime,omitempty"`
	// The overall timeout of a task including user-triggered retries.
	Timeout *duration.Duration `protobuf:"bytes,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// Number of retries per task.
	Retries *RetryStrategy `protobuf:"bytes,5,opt,name=retries,proto3" json:"retries,omitempty"`
	// Indicates a logical version to apply to this task for the purpose of discovery.
	DiscoveryVersion string `protobuf:"bytes,6,opt,name=discovery_version,json=discoveryVersion,proto3" json:"discovery_version,omitempty"`
	// If set, this indicates that this task is deprecated.  This will enable owners of tasks to notify consumers
	// of the ending of support for a given task.
	DeprecatedErrorMessage string `protobuf:"bytes,7,opt,name=deprecated_error_message,json=deprecatedErrorMessage,proto3" json:"deprecated_error_message,omitempty"`
	// Identify whether task is interruptible
	//
	// Types that are valid to be assigned to InterruptibleValue:
	//	*TaskMetadata_Interruptible
	InterruptibleValue   isTaskMetadata_InterruptibleValue `protobuf_oneof:"interruptible_value"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *TaskMetadata) Reset()         { *m = TaskMetadata{} }
func (m *TaskMetadata) String() string { return proto.CompactTextString(m) }
func (*TaskMetadata) ProtoMessage()    {}
func (*TaskMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd8423ba58d6ed80, []int{2}
}

func (m *TaskMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskMetadata.Unmarshal(m, b)
}
func (m *TaskMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskMetadata.Marshal(b, m, deterministic)
}
func (m *TaskMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskMetadata.Merge(m, src)
}
func (m *TaskMetadata) XXX_Size() int {
	return xxx_messageInfo_TaskMetadata.Size(m)
}
func (m *TaskMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TaskMetadata proto.InternalMessageInfo

func (m *TaskMetadata) GetDiscoverable() bool {
	if m != nil {
		return m.Discoverable
	}
	return false
}

func (m *TaskMetadata) GetRuntime() *RuntimeMetadata {
	if m != nil {
		return m.Runtime
	}
	return nil
}

func (m *TaskMetadata) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *TaskMetadata) GetRetries() *RetryStrategy {
	if m != nil {
		return m.Retries
	}
	return nil
}

func (m *TaskMetadata) GetDiscoveryVersion() string {
	if m != nil {
		return m.DiscoveryVersion
	}
	return ""
}

func (m *TaskMetadata) GetDeprecatedErrorMessage() string {
	if m != nil {
		return m.DeprecatedErrorMessage
	}
	return ""
}

type isTaskMetadata_InterruptibleValue interface {
	isTaskMetadata_InterruptibleValue()
}

type TaskMetadata_Interruptible struct {
	Interruptible bool `protobuf:"varint,8,opt,name=interruptible,proto3,oneof"`
}

func (*TaskMetadata_Interruptible) isTaskMetadata_InterruptibleValue() {}

func (m *TaskMetadata) GetInterruptibleValue() isTaskMetadata_InterruptibleValue {
	if m != nil {
		return m.InterruptibleValue
	}
	return nil
}

func (m *TaskMetadata) GetInterruptible() bool {
	if x, ok := m.GetInterruptibleValue().(*TaskMetadata_Interruptible); ok {
		return x.Interruptible
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TaskMetadata) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TaskMetadata_Interruptible)(nil),
	}
}

// A Task structure that uniquely identifies a task in the system
// Tasks are registered as a first step in the system.
type TaskTemplate struct {
	// Auto generated taskId by the system. Task Id uniquely identifies this task globally.
	Id *Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// A predefined yet extensible Task type identifier. This can be used to customize any of the components. If no
	// extensions are provided in the system, Flyte will resolve the this task to its TaskCategory and default the
	// implementation registered for the TaskCategory.
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// Extra metadata about the task.
	Metadata *TaskMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// A strongly typed interface for the task. This enables others to use this task within a workflow and gauarantees
	// compile-time validation of the workflow to avoid costly runtime failures.
	Interface *TypedInterface `protobuf:"bytes,4,opt,name=interface,proto3" json:"interface,omitempty"`
	// Custom data about the task. This is extensible to allow various plugins in the system.
	Custom *_struct.Struct `protobuf:"bytes,5,opt,name=custom,proto3" json:"custom,omitempty"`
	// Known target types that the system will guarantee plugins for. Custom SDK plugins are allowed to set these if needed.
	// If no corresponding execution-layer plugins are found, the system will default to handling these using built-in
	// handlers.
	//
	// Types that are valid to be assigned to Target:
	//	*TaskTemplate_Container
	Target               isTaskTemplate_Target `protobuf_oneof:"target"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *TaskTemplate) Reset()         { *m = TaskTemplate{} }
func (m *TaskTemplate) String() string { return proto.CompactTextString(m) }
func (*TaskTemplate) ProtoMessage()    {}
func (*TaskTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd8423ba58d6ed80, []int{3}
}

func (m *TaskTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskTemplate.Unmarshal(m, b)
}
func (m *TaskTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskTemplate.Marshal(b, m, deterministic)
}
func (m *TaskTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskTemplate.Merge(m, src)
}
func (m *TaskTemplate) XXX_Size() int {
	return xxx_messageInfo_TaskTemplate.Size(m)
}
func (m *TaskTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_TaskTemplate proto.InternalMessageInfo

func (m *TaskTemplate) GetId() *Identifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *TaskTemplate) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *TaskTemplate) GetMetadata() *TaskMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *TaskTemplate) GetInterface() *TypedInterface {
	if m != nil {
		return m.Interface
	}
	return nil
}

func (m *TaskTemplate) GetCustom() *_struct.Struct {
	if m != nil {
		return m.Custom
	}
	return nil
}

type isTaskTemplate_Target interface {
	isTaskTemplate_Target()
}

type TaskTemplate_Container struct {
	Container *Container `protobuf:"bytes,6,opt,name=container,proto3,oneof"`
}

func (*TaskTemplate_Container) isTaskTemplate_Target() {}

func (m *TaskTemplate) GetTarget() isTaskTemplate_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *TaskTemplate) GetContainer() *Container {
	if x, ok := m.GetTarget().(*TaskTemplate_Container); ok {
		return x.Container
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TaskTemplate) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TaskTemplate_Container)(nil),
	}
}

// Defines port properties for a container.
type ContainerPort struct {
	// Number of port to expose on the pod's IP address.
	// This must be a valid port number, 0 < x < 65536.
	ContainerPort        uint32   `protobuf:"varint,1,opt,name=container_port,json=containerPort,proto3" json:"container_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContainerPort) Reset()         { *m = ContainerPort{} }
func (m *ContainerPort) String() string { return proto.CompactTextString(m) }
func (*ContainerPort) ProtoMessage()    {}
func (*ContainerPort) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd8423ba58d6ed80, []int{4}
}

func (m *ContainerPort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerPort.Unmarshal(m, b)
}
func (m *ContainerPort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerPort.Marshal(b, m, deterministic)
}
func (m *ContainerPort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerPort.Merge(m, src)
}
func (m *ContainerPort) XXX_Size() int {
	return xxx_messageInfo_ContainerPort.Size(m)
}
func (m *ContainerPort) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerPort.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerPort proto.InternalMessageInfo

func (m *ContainerPort) GetContainerPort() uint32 {
	if m != nil {
		return m.ContainerPort
	}
	return 0
}

type Container struct {
	// Container image url. Eg: docker/redis:latest
	Image string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	// Command to be executed, if not provided, the default entrypoint in the container image will be used.
	Command []string `protobuf:"bytes,2,rep,name=command,proto3" json:"command,omitempty"`
	// These will default to Flyte given paths. If provided, the system will not append known paths. If the task still
	// needs flyte's inputs and outputs path, add $(FLYTE_INPUT_FILE), $(FLYTE_OUTPUT_FILE) wherever makes sense and the
	// system will populate these before executing the container.
	Args []string `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty"`
	// Container resources requirement as specified by the container engine.
	Resources *Resources `protobuf:"bytes,4,opt,name=resources,proto3" json:"resources,omitempty"`
	// Environment variables will be set as the container is starting up.
	Env []*KeyValuePair `protobuf:"bytes,5,rep,name=env,proto3" json:"env,omitempty"`
	// Allows extra configs to be available for the container.
	// TODO: elaborate on how configs will become available.
	Config []*KeyValuePair `protobuf:"bytes,6,rep,name=config,proto3" json:"config,omitempty"`
	// Ports to open in the container. This feature is not supported by all execution engines. (e.g. supported on K8s but
	// not supported on AWS Batch)
	// Only K8s
	Ports []*ContainerPort `protobuf:"bytes,7,rep,name=ports,proto3" json:"ports,omitempty"`
	// BETA: Optional configuration for DataLoading. If not specified, then default values are used.
	// This makes it possible to to run a completely portable container, that uses inputs and outputs
	// only from the local file-system and without having any reference to flyteidl. This is supported only on K8s at the moment.
	// If data loading is enabled, then data will be mounted in accompanying directories specified in the DataLoadingConfig. If the directories
	// are not specified, inputs will be mounted onto and outputs will be uploaded from a pre-determined file-system path. Refer to the documentation
	// to understand the default paths.
	// Only K8s
	DataConfig           *DataLoadingConfig `protobuf:"bytes,9,opt,name=data_config,json=dataConfig,proto3" json:"data_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd8423ba58d6ed80, []int{5}
}

func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (m *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(m, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Container) GetCommand() []string {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *Container) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Container) GetResources() *Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *Container) GetEnv() []*KeyValuePair {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *Container) GetConfig() []*KeyValuePair {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *Container) GetPorts() []*ContainerPort {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *Container) GetDataConfig() *DataLoadingConfig {
	if m != nil {
		return m.DataConfig
	}
	return nil
}

// Strategy to use when dealing with Blob, Schema, or multipart blob data (large datasets)
type IOStrategy struct {
	// Mode to use to manage downloads
	DownloadMode IOStrategy_DownloadMode `protobuf:"varint,1,opt,name=download_mode,json=downloadMode,proto3,enum=flyteidl.core.IOStrategy_DownloadMode" json:"download_mode,omitempty"`
	// Mode to use to manage uploads
	UploadMode           IOStrategy_UploadMode `protobuf:"varint,2,opt,name=upload_mode,json=uploadMode,proto3,enum=flyteidl.core.IOStrategy_UploadMode" json:"upload_mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *IOStrategy) Reset()         { *m = IOStrategy{} }
func (m *IOStrategy) String() string { return proto.CompactTextString(m) }
func (*IOStrategy) ProtoMessage()    {}
func (*IOStrategy) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd8423ba58d6ed80, []int{6}
}

func (m *IOStrategy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IOStrategy.Unmarshal(m, b)
}
func (m *IOStrategy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IOStrategy.Marshal(b, m, deterministic)
}
func (m *IOStrategy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IOStrategy.Merge(m, src)
}
func (m *IOStrategy) XXX_Size() int {
	return xxx_messageInfo_IOStrategy.Size(m)
}
func (m *IOStrategy) XXX_DiscardUnknown() {
	xxx_messageInfo_IOStrategy.DiscardUnknown(m)
}

var xxx_messageInfo_IOStrategy proto.InternalMessageInfo

func (m *IOStrategy) GetDownloadMode() IOStrategy_DownloadMode {
	if m != nil {
		return m.DownloadMode
	}
	return IOStrategy_DOWNLOAD_EAGER
}

func (m *IOStrategy) GetUploadMode() IOStrategy_UploadMode {
	if m != nil {
		return m.UploadMode
	}
	return IOStrategy_UPLOAD_ON_EXIT
}

// This configuration allows executing raw containers in Flyte using the Flyte CoPilot system.
// Flyte CoPilot, eliminates the needs of flytekit or sdk inside the container. Any inputs required by the users container are side-loaded in the input_path
// Any outputs generated by the user container - within output_path are automatically uploaded.
type DataLoadingConfig struct {
	// Flag enables DataLoading Config. If this is not set, data loading will not be used!
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// File system path (start at root). This folder will contain all the inputs exploded to a separate file.
	// Example, if the input interface needs (x: int, y: blob, z: multipart_blob) and the input path is "/var/flyte/inputs", then the file system will look like
	// /var/flyte/inputs/inputs.<metadata format dependent -> .pb .json .yaml> -> Format as defined previously. The Blob and Multipart blob will reference local filesystem instead of remote locations
	// /var/flyte/inputs/x -> X is a file that contains the value of x (integer) in string format
	// /var/flyte/inputs/y -> Y is a file in Binary format
	// /var/flyte/inputs/z/... -> Note Z itself is a directory
	// More information about the protocol - refer to docs #TODO reference docs here
	InputPath string `protobuf:"bytes,2,opt,name=input_path,json=inputPath,proto3" json:"input_path,omitempty"`
	// File system path (start at root). This folder should contain all the outputs for the task as individual files and/or an error text file
	OutputPath string `protobuf:"bytes,3,opt,name=output_path,json=outputPath,proto3" json:"output_path,omitempty"`
	// In the inputs folder, there will be an additional summary/metadata file that contains references to all files or inlined primitive values.
	// This format decides the actual encoding for the data. Refer to the encoding to understand the specifics of the contents and the encoding
	Format               DataLoadingConfig_LiteralMapFormat `protobuf:"varint,4,opt,name=format,proto3,enum=flyteidl.core.DataLoadingConfig_LiteralMapFormat" json:"format,omitempty"`
	IoStrategy           *IOStrategy                        `protobuf:"bytes,5,opt,name=io_strategy,json=ioStrategy,proto3" json:"io_strategy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *DataLoadingConfig) Reset()         { *m = DataLoadingConfig{} }
func (m *DataLoadingConfig) String() string { return proto.CompactTextString(m) }
func (*DataLoadingConfig) ProtoMessage()    {}
func (*DataLoadingConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd8423ba58d6ed80, []int{7}
}

func (m *DataLoadingConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataLoadingConfig.Unmarshal(m, b)
}
func (m *DataLoadingConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataLoadingConfig.Marshal(b, m, deterministic)
}
func (m *DataLoadingConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataLoadingConfig.Merge(m, src)
}
func (m *DataLoadingConfig) XXX_Size() int {
	return xxx_messageInfo_DataLoadingConfig.Size(m)
}
func (m *DataLoadingConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DataLoadingConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DataLoadingConfig proto.InternalMessageInfo

func (m *DataLoadingConfig) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *DataLoadingConfig) GetInputPath() string {
	if m != nil {
		return m.InputPath
	}
	return ""
}

func (m *DataLoadingConfig) GetOutputPath() string {
	if m != nil {
		return m.OutputPath
	}
	return ""
}

func (m *DataLoadingConfig) GetFormat() DataLoadingConfig_LiteralMapFormat {
	if m != nil {
		return m.Format
	}
	return DataLoadingConfig_JSON
}

func (m *DataLoadingConfig) GetIoStrategy() *IOStrategy {
	if m != nil {
		return m.IoStrategy
	}
	return nil
}

func init() {
	proto.RegisterEnum("flyteidl.core.Resources_ResourceName", Resources_ResourceName_name, Resources_ResourceName_value)
	proto.RegisterEnum("flyteidl.core.RuntimeMetadata_RuntimeType", RuntimeMetadata_RuntimeType_name, RuntimeMetadata_RuntimeType_value)
	proto.RegisterEnum("flyteidl.core.IOStrategy_DownloadMode", IOStrategy_DownloadMode_name, IOStrategy_DownloadMode_value)
	proto.RegisterEnum("flyteidl.core.IOStrategy_UploadMode", IOStrategy_UploadMode_name, IOStrategy_UploadMode_value)
	proto.RegisterEnum("flyteidl.core.DataLoadingConfig_LiteralMapFormat", DataLoadingConfig_LiteralMapFormat_name, DataLoadingConfig_LiteralMapFormat_value)
	proto.RegisterType((*Resources)(nil), "flyteidl.core.Resources")
	proto.RegisterType((*Resources_ResourceEntry)(nil), "flyteidl.core.Resources.ResourceEntry")
	proto.RegisterType((*RuntimeMetadata)(nil), "flyteidl.core.RuntimeMetadata")
	proto.RegisterType((*TaskMetadata)(nil), "flyteidl.core.TaskMetadata")
	proto.RegisterType((*TaskTemplate)(nil), "flyteidl.core.TaskTemplate")
	proto.RegisterType((*ContainerPort)(nil), "flyteidl.core.ContainerPort")
	proto.RegisterType((*Container)(nil), "flyteidl.core.Container")
	proto.RegisterType((*IOStrategy)(nil), "flyteidl.core.IOStrategy")
	proto.RegisterType((*DataLoadingConfig)(nil), "flyteidl.core.DataLoadingConfig")
}

func init() { proto.RegisterFile("flyteidl/core/tasks.proto", fileDescriptor_bd8423ba58d6ed80) }

var fileDescriptor_bd8423ba58d6ed80 = []byte{
	// 1147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0x5f, 0x6f, 0xdb, 0xb6,
	0x17, 0x8d, 0x95, 0xc4, 0x7f, 0xae, 0xe3, 0x54, 0x65, 0x7f, 0xbf, 0x4e, 0xcd, 0xda, 0x2e, 0x10,
	0xd6, 0xae, 0xdb, 0x50, 0x7b, 0x75, 0x81, 0xae, 0xdb, 0x80, 0x02, 0x49, 0xed, 0xb6, 0x59, 0x6c,
	0xcb, 0xa0, 0x9d, 0x76, 0xdd, 0x8b, 0xc6, 0x48, 0xb4, 0x4b, 0x54, 0x12, 0x35, 0x8a, 0xca, 0xe0,
	0x4f, 0xb4, 0x97, 0xbd, 0x0e, 0xd8, 0xbe, 0xce, 0xbe, 0xc8, 0x06, 0x51, 0x94, 0xfc, 0xa7, 0xc8,
	0x82, 0x3d, 0x99, 0xbc, 0xe7, 0x9c, 0x4b, 0xea, 0xde, 0xeb, 0x23, 0xc1, 0xad, 0x59, 0xb0, 0x90,
	0x94, 0xf9, 0x41, 0xc7, 0xe3, 0x82, 0x76, 0x24, 0x49, 0xde, 0x27, 0xed, 0x58, 0x70, 0xc9, 0x51,
	0xab, 0x80, 0xda, 0x19, 0x74, 0x70, 0x77, 0x9d, 0xc9, 0x7c, 0x1a, 0x49, 0x36, 0x63, 0x54, 0xe4,
	0xf4, 0x83, 0x3b, 0x1b, 0x78, 0x24, 0xa9, 0x98, 0x11, 0x8f, 0x6a, 0xf8, 0xf6, 0x3a, 0x1c, 0x30,
	0x49, 0x05, 0x09, 0xf4, 0x59, 0x07, 0x77, 0xe7, 0x9c, 0xcf, 0x03, 0xda, 0x51, 0xbb, 0xf3, 0x74,
	0xd6, 0xf1, 0x53, 0x41, 0x24, 0xe3, 0x51, 0xa1, 0xde, 0xc4, 0x13, 0x29, 0x52, 0x4f, 0xe6, 0xa8,
	0xfd, 0xa7, 0x01, 0x0d, 0x4c, 0x13, 0x9e, 0x0a, 0x8f, 0x26, 0xe8, 0x18, 0xea, 0x82, 0xfe, 0x9c,
	0xd2, 0x44, 0x26, 0x56, 0xe5, 0x70, 0xfb, 0x41, 0xb3, 0x7b, 0xbf, 0xbd, 0xf6, 0x28, 0xed, 0x92,
	0x5b, 0xae, 0xfa, 0x91, 0x14, 0x0b, 0x5c, 0xea, 0xd0, 0x33, 0xa8, 0x06, 0x2c, 0x64, 0x32, 0xb1,
	0x8c, 0xff, 0x94, 0x41, 0xab, 0x0e, 0x7e, 0x82, 0xd6, 0x1a, 0x80, 0xbe, 0x81, 0x9d, 0x88, 0x84,
	0xd4, 0xaa, 0x1c, 0x56, 0x1e, 0xec, 0x77, 0xef, 0x5d, 0x99, 0x6e, 0x44, 0x42, 0x8a, 0x95, 0x04,
	0xfd, 0x0f, 0x76, 0x2f, 0x48, 0x90, 0x52, 0xcb, 0x38, 0xac, 0x3c, 0x68, 0xe0, 0x7c, 0x63, 0xbf,
	0x80, 0xbd, 0x55, 0x2e, 0x6a, 0x42, 0xed, 0x6c, 0x74, 0x3a, 0x72, 0xde, 0x8c, 0xcc, 0x2d, 0x54,
	0x83, 0xed, 0xe7, 0xe3, 0x33, 0xb3, 0x92, 0x2d, 0x5e, 0x8e, 0xcf, 0x4c, 0x03, 0x01, 0x54, 0x87,
	0xfd, 0xa1, 0x83, 0xdf, 0x9a, 0xdb, 0x19, 0x75, 0x32, 0x75, 0xf0, 0xd1, 0xcb, 0xbe, 0xb9, 0x63,
	0xff, 0x56, 0x81, 0x6b, 0x38, 0x8d, 0x24, 0x0b, 0xe9, 0x90, 0x4a, 0xe2, 0x13, 0x49, 0xd0, 0x33,
	0xd8, 0x91, 0x8b, 0xb8, 0xb8, 0xec, 0x17, 0x9b, 0x97, 0x5d, 0x67, 0x17, 0xfb, 0xe9, 0x22, 0xa6,
	0x58, 0xe9, 0x90, 0x05, 0xb5, 0x0b, 0x2a, 0x12, 0xc6, 0x23, 0x7d, 0xe7, 0x62, 0x8b, 0x6e, 0x42,
	0x75, 0x16, 0x90, 0x0b, 0x2e, 0xac, 0x6d, 0x05, 0xe8, 0x9d, 0xfd, 0x19, 0x34, 0x57, 0xd2, 0xa0,
	0x06, 0xec, 0x3a, 0xd3, 0x57, 0x7d, 0x6c, 0x6e, 0xa1, 0x16, 0x34, 0x5e, 0x0c, 0xde, 0x4e, 0xfb,
	0xee, 0xa4, 0x77, 0x6a, 0x56, 0xec, 0xbf, 0x0d, 0xd8, 0x9b, 0x92, 0xe4, 0x7d, 0x79, 0x57, 0x1b,
	0xf6, 0x7c, 0x96, 0x78, 0xfc, 0x82, 0x0a, 0x72, 0x1e, 0xe4, 0x77, 0xae, 0xe3, 0xb5, 0x18, 0x7a,
	0x0a, 0x35, 0x91, 0x67, 0x57, 0xf7, 0x69, 0x76, 0xef, 0xfe, 0xfb, 0x23, 0xe1, 0x82, 0x8e, 0x1e,
	0x43, 0x2d, 0xfb, 0xe5, 0xa9, 0xb4, 0x76, 0x94, 0xf2, 0x56, 0x3b, 0x9f, 0xc4, 0x76, 0x31, 0x89,
	0xed, 0x9e, 0x9e, 0x54, 0x5c, 0x30, 0xd1, 0x13, 0xa8, 0x09, 0x2a, 0x05, 0xa3, 0x89, 0xb5, 0xab,
	0x44, 0xb7, 0x3f, 0x68, 0xb7, 0x14, 0x8b, 0x89, 0x14, 0x44, 0xd2, 0xf9, 0x02, 0x17, 0x64, 0xf4,
	0x25, 0x5c, 0x2f, 0xae, 0xbd, 0x70, 0x8b, 0x02, 0x56, 0x55, 0x9d, 0xcc, 0x12, 0x78, 0xad, 0x2b,
	0xf9, 0x14, 0x2c, 0x9f, 0xc6, 0x82, 0x7a, 0x44, 0x52, 0xdf, 0xa5, 0x42, 0x70, 0xe1, 0x86, 0x34,
	0x49, 0xc8, 0x9c, 0x5a, 0x35, 0xa5, 0xb9, 0xb9, 0xc4, 0xfb, 0x19, 0x3c, 0xcc, 0x51, 0x74, 0x1f,
	0x5a, 0xea, 0xcf, 0x29, 0xd2, 0x58, 0xb2, 0xac, 0x64, 0xf5, 0xac, 0x64, 0xaf, 0xb6, 0xf0, 0x7a,
	0xf8, 0xf8, 0xff, 0x70, 0x63, 0x2d, 0xe0, 0xe6, 0x83, 0xf7, 0x87, 0xee, 0xc0, 0x94, 0x86, 0x71,
	0x40, 0x24, 0x45, 0x9f, 0x83, 0xc1, 0x7c, 0x55, 0xf7, 0xac, 0x3c, 0xeb, 0x4f, 0x7a, 0x52, 0xba,
	0x04, 0x36, 0x98, 0x8f, 0x90, 0x1e, 0xac, 0x7c, 0x2a, 0xf2, 0x61, 0xf9, 0x1a, 0xea, 0xa1, 0xae,
	0xbb, 0x1a, 0x8a, 0x66, 0xf7, 0xe3, 0x8d, 0x24, 0xab, 0xfd, 0xc6, 0x25, 0x19, 0x7d, 0x07, 0x8d,
	0xd2, 0x64, 0x74, 0x77, 0xee, 0x6c, 0x2a, 0x17, 0x31, 0xf5, 0x4f, 0x0a, 0x12, 0x5e, 0xf2, 0x51,
	0x07, 0xaa, 0x5e, 0x9a, 0x48, 0x1e, 0xea, 0x16, 0x7d, 0xf4, 0x41, 0x5f, 0x27, 0xca, 0x61, 0xb0,
	0xa6, 0xa1, 0xa7, 0xd0, 0xf0, 0x78, 0x24, 0x09, 0x8b, 0xa8, 0x50, 0x4d, 0x69, 0x76, 0xad, 0x8d,
	0xd3, 0x9e, 0x17, 0xf8, 0xab, 0x2d, 0xbc, 0x24, 0x1f, 0xd7, 0xa1, 0x2a, 0x89, 0x98, 0x53, 0x69,
	0x3f, 0x81, 0x56, 0xc9, 0x19, 0x73, 0x21, 0xd1, 0x3d, 0xd8, 0x2f, 0x79, 0x6e, 0xcc, 0x85, 0x54,
	0x65, 0x6c, 0xe1, 0x96, 0xb7, 0x4a, 0xb3, 0xff, 0x32, 0xa0, 0x51, 0x0a, 0x33, 0x3f, 0x60, 0x61,
	0xd6, 0xe6, 0x4a, 0xee, 0x07, 0x6a, 0x93, 0xfd, 0xe7, 0x3c, 0x1e, 0x86, 0x24, 0xf2, 0x95, 0x65,
	0x35, 0x70, 0xb1, 0xcd, 0x8a, 0x4e, 0xc4, 0x3c, 0xb1, 0xb6, 0x55, 0x58, 0xad, 0xd1, 0x13, 0x68,
	0x88, 0xc2, 0x73, 0x74, 0xed, 0xac, 0xcb, 0x3c, 0x09, 0x2f, 0xa9, 0xe8, 0x21, 0x6c, 0xd3, 0xe8,
	0xc2, 0xda, 0x55, 0xa6, 0xb8, 0xd9, 0xa7, 0x53, 0xba, 0x78, 0x9d, 0x8d, 0xc8, 0x98, 0x30, 0x81,
	0x33, 0x1e, 0x7a, 0x0c, 0x55, 0x8f, 0x47, 0x33, 0x36, 0xb7, 0xaa, 0x57, 0x2b, 0x34, 0x15, 0x75,
	0x61, 0x37, 0x2b, 0x45, 0x62, 0xd5, 0x94, 0xe6, 0xf6, 0x65, 0x55, 0xce, 0x4a, 0x83, 0x73, 0x2a,
	0x3a, 0x82, 0x66, 0x36, 0x13, 0xae, 0x3e, 0xad, 0xa1, 0x9e, 0xe8, 0x70, 0x43, 0xd9, 0x23, 0x92,
	0x0c, 0x38, 0xf1, 0x59, 0x34, 0x7f, 0xae, 0x78, 0x18, 0x32, 0x51, 0xbe, 0xb6, 0x7f, 0x37, 0x00,
	0x4e, 0x9c, 0xe2, 0x5f, 0x89, 0x4e, 0xa1, 0xe5, 0xf3, 0x5f, 0xa2, 0x80, 0x13, 0xdf, 0x0d, 0xb9,
	0x5f, 0x98, 0xe1, 0xe6, 0x8b, 0x60, 0xa9, 0x68, 0xf7, 0x34, 0x7d, 0xc8, 0x7d, 0x8a, 0xf7, 0xfc,
	0x95, 0x1d, 0xea, 0x43, 0x33, 0x8d, 0x97, 0xa9, 0x0c, 0x95, 0xea, 0xd3, 0xcb, 0x53, 0x9d, 0xc5,
	0x65, 0x22, 0x48, 0xcb, 0xb5, 0x3d, 0x80, 0xbd, 0xd5, 0x43, 0x10, 0x82, 0xfd, 0x9e, 0xf3, 0x66,
	0x34, 0x70, 0x8e, 0x7a, 0x6e, 0xff, 0xe8, 0xa5, 0xf2, 0xcb, 0x1b, 0x70, 0xad, 0x8c, 0x4d, 0xa6,
	0xb8, 0x7f, 0x34, 0x34, 0x2b, 0x79, 0xd0, 0x1d, 0x39, 0x53, 0xb7, 0xc0, 0x4c, 0xc3, 0xee, 0x03,
	0x2c, 0xcf, 0xc9, 0x72, 0x9d, 0x8d, 0x95, 0xca, 0x19, 0xb9, 0xfd, 0x1f, 0x4e, 0xa6, 0xe6, 0x16,
	0x32, 0x61, 0x4f, 0xc7, 0xf2, 0xec, 0x15, 0x74, 0x1d, 0x5a, 0x3a, 0x51, 0x0e, 0x98, 0x86, 0xfd,
	0xab, 0x01, 0xd7, 0x3f, 0xa8, 0x6c, 0x36, 0x8e, 0x34, 0xca, 0xcc, 0xd7, 0xd7, 0x8e, 0x5c, 0x6c,
	0xd1, 0x1d, 0x00, 0x16, 0xc5, 0xa9, 0x74, 0x63, 0x22, 0xdf, 0x69, 0x27, 0x68, 0xa8, 0xc8, 0x98,
	0xc8, 0x77, 0xe8, 0x13, 0x68, 0xf2, 0x54, 0x96, 0x78, 0xfe, 0x9a, 0x80, 0x3c, 0xa4, 0x08, 0x27,
	0x50, 0x9d, 0x71, 0x11, 0x92, 0xdc, 0x91, 0xf7, 0xbb, 0x8f, 0xae, 0xea, 0x72, 0x7b, 0x90, 0x7f,
	0x6b, 0x0c, 0x49, 0xfc, 0x42, 0x09, 0xb1, 0x4e, 0x80, 0xbe, 0x85, 0x26, 0xe3, 0x6e, 0xa2, 0xab,
	0xae, 0x9d, 0xe0, 0xd6, 0xa5, 0x6d, 0xc1, 0xc0, 0x78, 0xb1, 0xb6, 0x1f, 0x81, 0xb9, 0x99, 0x17,
	0xd5, 0x61, 0xe7, 0xfb, 0x89, 0x93, 0xbd, 0x80, 0xeb, 0xb0, 0xf3, 0xf6, 0x68, 0x38, 0x30, 0x2b,
	0xd9, 0xab, 0x6c, 0x8c, 0x9d, 0xa9, 0x63, 0x1a, 0xc7, 0xdd, 0x1f, 0xbf, 0x9a, 0x33, 0xf9, 0x2e,
	0x3d, 0x6f, 0x7b, 0x3c, 0xec, 0x04, 0x8b, 0x99, 0xec, 0x94, 0x1f, 0x45, 0x73, 0x1a, 0x75, 0xe2,
	0xf3, 0x87, 0x73, 0xde, 0x59, 0xfb, 0x4e, 0x3a, 0xaf, 0x2a, 0x3f, 0x7a, 0xfc, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x87, 0x5a, 0xf4, 0x00, 0xa8, 0x09, 0x00, 0x00,
}
