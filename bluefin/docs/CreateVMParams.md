# CreateVMParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandLineArgs** | Pointer to **string** |  | [optional] [default to ""]
**CpuMode** | Pointer to **string** |  | [optional] [default to "CUSTOM"]
**CpuModel** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Vcpus** | Pointer to **int32** | Maximum of 16 guest virtual CPUs are allowed. By default, every virtual CPU is configured as a separate package. Multiple cores can be configured per CPU by specifying &#x60;cores&#x60; attributes. &#x60;vcpus&#x60; specifies total number of CPU sockets. &#x60;cores&#x60; specifies number of cores per socket. &#x60;threads&#x60; specifies number of threads per core. | [optional] [default to 1]
**Cores** | Pointer to **int32** | Maximum of 16 guest virtual CPUs are allowed. By default, every virtual CPU is configured as a separate package. Multiple cores can be configured per CPU by specifying &#x60;cores&#x60; attributes. &#x60;vcpus&#x60; specifies total number of CPU sockets. &#x60;cores&#x60; specifies number of cores per socket. &#x60;threads&#x60; specifies number of threads per core. | [optional] [default to 1]
**Threads** | Pointer to **int32** | Maximum of 16 guest virtual CPUs are allowed. By default, every virtual CPU is configured as a separate package. Multiple cores can be configured per CPU by specifying &#x60;cores&#x60; attributes. &#x60;vcpus&#x60; specifies total number of CPU sockets. &#x60;cores&#x60; specifies number of cores per socket. &#x60;threads&#x60; specifies number of threads per core. | [optional] [default to 1]
**Cpuset** | Pointer to **NullableString** |  | [optional] 
**Nodeset** | Pointer to **NullableString** |  | [optional] 
**PinVcpus** | Pointer to **bool** |  | [optional] [default to false]
**SuspendOnSnapshot** | Pointer to **bool** |  | [optional] [default to false]
**Memory** | Pointer to **int64** |  | [optional] 
**MinMemory** | Pointer to **NullableInt64** |  | [optional] 
**HypervEnlightenments** | Pointer to **bool** | &#x60;hyperv_enlightenments&#x60; can be used to enable subset of predefined Hyper-V enlightenments for Windows guests. These enlightenments improve performance and enable otherwise missing features. | [optional] [default to false]
**Bootloader** | Pointer to **string** |  | [optional] [default to "UEFI"]
**Autostart** | Pointer to **bool** |  | [optional] [default to true]
**HideFromMsr** | Pointer to **bool** | &#x60;hide_from_msr&#x60; is a boolean which when set will hide the KVM hypervisor from standard MSR based discovery and is useful to enable when doing GPU passthrough. | [optional] [default to false]
**EnsureDisplayDevice** | Pointer to **bool** | &#x60;ensure_display_device&#x60; when set ( the default ) will ensure that the guest always has access to a video device. For headless installations like ubuntu server this is required for the guest to operate properly. However for cases where consumer would like to use GPU passthrough and does not want a display device added should set this to &#x60;false&#x60;. | [optional] [default to true]
**Time** | Pointer to **string** |  | [optional] [default to "LOCAL"]
**ShutdownTimeout** | Pointer to **int32** | &#x60;shutdown_timeout&#x60; indicates the time in seconds the system waits for the VM to cleanly shutdown. During system shutdown, if the VM hasn&#39;t exited after a hardware shutdown signal has been sent by the system within &#x60;shutdown_timeout&#x60; seconds, system initiates poweroff for the VM to stop it. | [optional] [default to 90]
**ArchType** | Pointer to **NullableString** | &#x60;arch_type&#x60; refers to architecture type and can be specified for the guest. By default the value is &#x60;null&#x60; and system in this case will choose a reasonable default based on host. &#x60;machine_type&#x60; refers to machine type of the guest based on the architecture type selected with &#x60;arch_type&#x60;. By default the value is &#x60;null&#x60; and system in this case will choose a reasonable default based on &#x60;arch_type&#x60; configuration. | [optional] 
**MachineType** | Pointer to **NullableString** | &#x60;machine_type&#x60; refers to machine type of the guest based on the architecture type selected with &#x60;arch_type&#x60;. By default the value is &#x60;null&#x60; and system in this case will choose a reasonable default based on &#x60;arch_type&#x60; configuration. | [optional] 
**Uuid** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateVMParams

`func NewCreateVMParams() *CreateVMParams`

NewCreateVMParams instantiates a new CreateVMParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVMParamsWithDefaults

`func NewCreateVMParamsWithDefaults() *CreateVMParams`

NewCreateVMParamsWithDefaults instantiates a new CreateVMParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandLineArgs

`func (o *CreateVMParams) GetCommandLineArgs() string`

GetCommandLineArgs returns the CommandLineArgs field if non-nil, zero value otherwise.

### GetCommandLineArgsOk

`func (o *CreateVMParams) GetCommandLineArgsOk() (*string, bool)`

GetCommandLineArgsOk returns a tuple with the CommandLineArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandLineArgs

`func (o *CreateVMParams) SetCommandLineArgs(v string)`

SetCommandLineArgs sets CommandLineArgs field to given value.

### HasCommandLineArgs

`func (o *CreateVMParams) HasCommandLineArgs() bool`

HasCommandLineArgs returns a boolean if a field has been set.

### GetCpuMode

`func (o *CreateVMParams) GetCpuMode() string`

GetCpuMode returns the CpuMode field if non-nil, zero value otherwise.

### GetCpuModeOk

`func (o *CreateVMParams) GetCpuModeOk() (*string, bool)`

GetCpuModeOk returns a tuple with the CpuMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuMode

`func (o *CreateVMParams) SetCpuMode(v string)`

SetCpuMode sets CpuMode field to given value.

### HasCpuMode

`func (o *CreateVMParams) HasCpuMode() bool`

HasCpuMode returns a boolean if a field has been set.

### GetCpuModel

`func (o *CreateVMParams) GetCpuModel() string`

GetCpuModel returns the CpuModel field if non-nil, zero value otherwise.

### GetCpuModelOk

`func (o *CreateVMParams) GetCpuModelOk() (*string, bool)`

GetCpuModelOk returns a tuple with the CpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuModel

`func (o *CreateVMParams) SetCpuModel(v string)`

SetCpuModel sets CpuModel field to given value.

### HasCpuModel

`func (o *CreateVMParams) HasCpuModel() bool`

HasCpuModel returns a boolean if a field has been set.

### SetCpuModelNil

`func (o *CreateVMParams) SetCpuModelNil(b bool)`

 SetCpuModelNil sets the value for CpuModel to be an explicit nil

### UnsetCpuModel
`func (o *CreateVMParams) UnsetCpuModel()`

UnsetCpuModel ensures that no value is present for CpuModel, not even an explicit nil
### GetName

`func (o *CreateVMParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateVMParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateVMParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateVMParams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateVMParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateVMParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateVMParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateVMParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVcpus

`func (o *CreateVMParams) GetVcpus() int32`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *CreateVMParams) GetVcpusOk() (*int32, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *CreateVMParams) SetVcpus(v int32)`

SetVcpus sets Vcpus field to given value.

### HasVcpus

`func (o *CreateVMParams) HasVcpus() bool`

HasVcpus returns a boolean if a field has been set.

### GetCores

`func (o *CreateVMParams) GetCores() int32`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *CreateVMParams) GetCoresOk() (*int32, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *CreateVMParams) SetCores(v int32)`

SetCores sets Cores field to given value.

### HasCores

`func (o *CreateVMParams) HasCores() bool`

HasCores returns a boolean if a field has been set.

### GetThreads

`func (o *CreateVMParams) GetThreads() int32`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *CreateVMParams) GetThreadsOk() (*int32, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *CreateVMParams) SetThreads(v int32)`

SetThreads sets Threads field to given value.

### HasThreads

`func (o *CreateVMParams) HasThreads() bool`

HasThreads returns a boolean if a field has been set.

### GetCpuset

`func (o *CreateVMParams) GetCpuset() string`

GetCpuset returns the Cpuset field if non-nil, zero value otherwise.

### GetCpusetOk

`func (o *CreateVMParams) GetCpusetOk() (*string, bool)`

GetCpusetOk returns a tuple with the Cpuset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuset

`func (o *CreateVMParams) SetCpuset(v string)`

SetCpuset sets Cpuset field to given value.

### HasCpuset

`func (o *CreateVMParams) HasCpuset() bool`

HasCpuset returns a boolean if a field has been set.

### SetCpusetNil

`func (o *CreateVMParams) SetCpusetNil(b bool)`

 SetCpusetNil sets the value for Cpuset to be an explicit nil

### UnsetCpuset
`func (o *CreateVMParams) UnsetCpuset()`

UnsetCpuset ensures that no value is present for Cpuset, not even an explicit nil
### GetNodeset

`func (o *CreateVMParams) GetNodeset() string`

GetNodeset returns the Nodeset field if non-nil, zero value otherwise.

### GetNodesetOk

`func (o *CreateVMParams) GetNodesetOk() (*string, bool)`

GetNodesetOk returns a tuple with the Nodeset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeset

`func (o *CreateVMParams) SetNodeset(v string)`

SetNodeset sets Nodeset field to given value.

### HasNodeset

`func (o *CreateVMParams) HasNodeset() bool`

HasNodeset returns a boolean if a field has been set.

### SetNodesetNil

`func (o *CreateVMParams) SetNodesetNil(b bool)`

 SetNodesetNil sets the value for Nodeset to be an explicit nil

### UnsetNodeset
`func (o *CreateVMParams) UnsetNodeset()`

UnsetNodeset ensures that no value is present for Nodeset, not even an explicit nil
### GetPinVcpus

`func (o *CreateVMParams) GetPinVcpus() bool`

GetPinVcpus returns the PinVcpus field if non-nil, zero value otherwise.

### GetPinVcpusOk

`func (o *CreateVMParams) GetPinVcpusOk() (*bool, bool)`

GetPinVcpusOk returns a tuple with the PinVcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinVcpus

`func (o *CreateVMParams) SetPinVcpus(v bool)`

SetPinVcpus sets PinVcpus field to given value.

### HasPinVcpus

`func (o *CreateVMParams) HasPinVcpus() bool`

HasPinVcpus returns a boolean if a field has been set.

### GetSuspendOnSnapshot

`func (o *CreateVMParams) GetSuspendOnSnapshot() bool`

GetSuspendOnSnapshot returns the SuspendOnSnapshot field if non-nil, zero value otherwise.

### GetSuspendOnSnapshotOk

`func (o *CreateVMParams) GetSuspendOnSnapshotOk() (*bool, bool)`

GetSuspendOnSnapshotOk returns a tuple with the SuspendOnSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendOnSnapshot

`func (o *CreateVMParams) SetSuspendOnSnapshot(v bool)`

SetSuspendOnSnapshot sets SuspendOnSnapshot field to given value.

### HasSuspendOnSnapshot

`func (o *CreateVMParams) HasSuspendOnSnapshot() bool`

HasSuspendOnSnapshot returns a boolean if a field has been set.

### GetMemory

`func (o *CreateVMParams) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *CreateVMParams) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *CreateVMParams) SetMemory(v int64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *CreateVMParams) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetMinMemory

`func (o *CreateVMParams) GetMinMemory() int64`

GetMinMemory returns the MinMemory field if non-nil, zero value otherwise.

### GetMinMemoryOk

`func (o *CreateVMParams) GetMinMemoryOk() (*int64, bool)`

GetMinMemoryOk returns a tuple with the MinMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMemory

`func (o *CreateVMParams) SetMinMemory(v int64)`

SetMinMemory sets MinMemory field to given value.

### HasMinMemory

`func (o *CreateVMParams) HasMinMemory() bool`

HasMinMemory returns a boolean if a field has been set.

### SetMinMemoryNil

`func (o *CreateVMParams) SetMinMemoryNil(b bool)`

 SetMinMemoryNil sets the value for MinMemory to be an explicit nil

### UnsetMinMemory
`func (o *CreateVMParams) UnsetMinMemory()`

UnsetMinMemory ensures that no value is present for MinMemory, not even an explicit nil
### GetHypervEnlightenments

`func (o *CreateVMParams) GetHypervEnlightenments() bool`

GetHypervEnlightenments returns the HypervEnlightenments field if non-nil, zero value otherwise.

### GetHypervEnlightenmentsOk

`func (o *CreateVMParams) GetHypervEnlightenmentsOk() (*bool, bool)`

GetHypervEnlightenmentsOk returns a tuple with the HypervEnlightenments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervEnlightenments

`func (o *CreateVMParams) SetHypervEnlightenments(v bool)`

SetHypervEnlightenments sets HypervEnlightenments field to given value.

### HasHypervEnlightenments

`func (o *CreateVMParams) HasHypervEnlightenments() bool`

HasHypervEnlightenments returns a boolean if a field has been set.

### GetBootloader

`func (o *CreateVMParams) GetBootloader() string`

GetBootloader returns the Bootloader field if non-nil, zero value otherwise.

### GetBootloaderOk

`func (o *CreateVMParams) GetBootloaderOk() (*string, bool)`

GetBootloaderOk returns a tuple with the Bootloader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootloader

`func (o *CreateVMParams) SetBootloader(v string)`

SetBootloader sets Bootloader field to given value.

### HasBootloader

`func (o *CreateVMParams) HasBootloader() bool`

HasBootloader returns a boolean if a field has been set.

### GetAutostart

`func (o *CreateVMParams) GetAutostart() bool`

GetAutostart returns the Autostart field if non-nil, zero value otherwise.

### GetAutostartOk

`func (o *CreateVMParams) GetAutostartOk() (*bool, bool)`

GetAutostartOk returns a tuple with the Autostart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutostart

`func (o *CreateVMParams) SetAutostart(v bool)`

SetAutostart sets Autostart field to given value.

### HasAutostart

`func (o *CreateVMParams) HasAutostart() bool`

HasAutostart returns a boolean if a field has been set.

### GetHideFromMsr

`func (o *CreateVMParams) GetHideFromMsr() bool`

GetHideFromMsr returns the HideFromMsr field if non-nil, zero value otherwise.

### GetHideFromMsrOk

`func (o *CreateVMParams) GetHideFromMsrOk() (*bool, bool)`

GetHideFromMsrOk returns a tuple with the HideFromMsr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideFromMsr

`func (o *CreateVMParams) SetHideFromMsr(v bool)`

SetHideFromMsr sets HideFromMsr field to given value.

### HasHideFromMsr

`func (o *CreateVMParams) HasHideFromMsr() bool`

HasHideFromMsr returns a boolean if a field has been set.

### GetEnsureDisplayDevice

`func (o *CreateVMParams) GetEnsureDisplayDevice() bool`

GetEnsureDisplayDevice returns the EnsureDisplayDevice field if non-nil, zero value otherwise.

### GetEnsureDisplayDeviceOk

`func (o *CreateVMParams) GetEnsureDisplayDeviceOk() (*bool, bool)`

GetEnsureDisplayDeviceOk returns a tuple with the EnsureDisplayDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnsureDisplayDevice

`func (o *CreateVMParams) SetEnsureDisplayDevice(v bool)`

SetEnsureDisplayDevice sets EnsureDisplayDevice field to given value.

### HasEnsureDisplayDevice

`func (o *CreateVMParams) HasEnsureDisplayDevice() bool`

HasEnsureDisplayDevice returns a boolean if a field has been set.

### GetTime

`func (o *CreateVMParams) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *CreateVMParams) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *CreateVMParams) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *CreateVMParams) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetShutdownTimeout

`func (o *CreateVMParams) GetShutdownTimeout() int32`

GetShutdownTimeout returns the ShutdownTimeout field if non-nil, zero value otherwise.

### GetShutdownTimeoutOk

`func (o *CreateVMParams) GetShutdownTimeoutOk() (*int32, bool)`

GetShutdownTimeoutOk returns a tuple with the ShutdownTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownTimeout

`func (o *CreateVMParams) SetShutdownTimeout(v int32)`

SetShutdownTimeout sets ShutdownTimeout field to given value.

### HasShutdownTimeout

`func (o *CreateVMParams) HasShutdownTimeout() bool`

HasShutdownTimeout returns a boolean if a field has been set.

### GetArchType

`func (o *CreateVMParams) GetArchType() string`

GetArchType returns the ArchType field if non-nil, zero value otherwise.

### GetArchTypeOk

`func (o *CreateVMParams) GetArchTypeOk() (*string, bool)`

GetArchTypeOk returns a tuple with the ArchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchType

`func (o *CreateVMParams) SetArchType(v string)`

SetArchType sets ArchType field to given value.

### HasArchType

`func (o *CreateVMParams) HasArchType() bool`

HasArchType returns a boolean if a field has been set.

### SetArchTypeNil

`func (o *CreateVMParams) SetArchTypeNil(b bool)`

 SetArchTypeNil sets the value for ArchType to be an explicit nil

### UnsetArchType
`func (o *CreateVMParams) UnsetArchType()`

UnsetArchType ensures that no value is present for ArchType, not even an explicit nil
### GetMachineType

`func (o *CreateVMParams) GetMachineType() string`

GetMachineType returns the MachineType field if non-nil, zero value otherwise.

### GetMachineTypeOk

`func (o *CreateVMParams) GetMachineTypeOk() (*string, bool)`

GetMachineTypeOk returns a tuple with the MachineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineType

`func (o *CreateVMParams) SetMachineType(v string)`

SetMachineType sets MachineType field to given value.

### HasMachineType

`func (o *CreateVMParams) HasMachineType() bool`

HasMachineType returns a boolean if a field has been set.

### SetMachineTypeNil

`func (o *CreateVMParams) SetMachineTypeNil(b bool)`

 SetMachineTypeNil sets the value for MachineType to be an explicit nil

### UnsetMachineType
`func (o *CreateVMParams) UnsetMachineType()`

UnsetMachineType ensures that no value is present for MachineType, not even an explicit nil
### GetUuid

`func (o *CreateVMParams) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CreateVMParams) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CreateVMParams) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CreateVMParams) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *CreateVMParams) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *CreateVMParams) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


