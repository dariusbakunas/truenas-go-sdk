# VM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Vcpus** | Pointer to **int32** |  | [optional] 
**Memory** | Pointer to **int64** |  | [optional] 
**MinMemory** | Pointer to **int64** |  | [optional] 
**Autostart** | Pointer to **bool** |  | [optional] 
**Time** | Pointer to **string** |  | [optional] 
**Bootloader** | Pointer to **string** |  | [optional] 
**Cores** | Pointer to **int32** |  | [optional] 
**Threads** | Pointer to **int32** |  | [optional] 
**HypervEnlightenments** | Pointer to **bool** |  | [optional] 
**ShutdownTimeout** | Pointer to **int32** |  | [optional] 
**CpuMode** | Pointer to **string** |  | [optional] 
**CpuModel** | Pointer to **string** |  | [optional] 
**Cpuset** | Pointer to **string** |  | [optional] 
**Nodeset** | Pointer to **string** |  | [optional] 
**PinVcpus** | Pointer to **bool** |  | [optional] 
**HideFromMsr** | Pointer to **bool** |  | [optional] 
**SuspendOnSnapshot** | Pointer to **bool** |  | [optional] 
**EnsureDisplayDevice** | Pointer to **bool** |  | [optional] 
**ArchType** | Pointer to **string** |  | [optional] 
**MachineType** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**CommandLineArgs** | Pointer to **string** |  | [optional] 
**Devices** | Pointer to [**[]VMDevice**](VMDevice.md) |  | [optional] 
**Status** | Pointer to [**VMStatus**](VMStatus.md) |  | [optional] 

## Methods

### NewVM

`func NewVM(id int32, name string, ) *VM`

NewVM instantiates a new VM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMWithDefaults

`func NewVMWithDefaults() *VM`

NewVMWithDefaults instantiates a new VM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VM) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VM) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VM) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *VM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VM) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *VM) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VM) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VM) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VM) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVcpus

`func (o *VM) GetVcpus() int32`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *VM) GetVcpusOk() (*int32, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *VM) SetVcpus(v int32)`

SetVcpus sets Vcpus field to given value.

### HasVcpus

`func (o *VM) HasVcpus() bool`

HasVcpus returns a boolean if a field has been set.

### GetMemory

`func (o *VM) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *VM) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *VM) SetMemory(v int64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *VM) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetMinMemory

`func (o *VM) GetMinMemory() int64`

GetMinMemory returns the MinMemory field if non-nil, zero value otherwise.

### GetMinMemoryOk

`func (o *VM) GetMinMemoryOk() (*int64, bool)`

GetMinMemoryOk returns a tuple with the MinMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMemory

`func (o *VM) SetMinMemory(v int64)`

SetMinMemory sets MinMemory field to given value.

### HasMinMemory

`func (o *VM) HasMinMemory() bool`

HasMinMemory returns a boolean if a field has been set.

### GetAutostart

`func (o *VM) GetAutostart() bool`

GetAutostart returns the Autostart field if non-nil, zero value otherwise.

### GetAutostartOk

`func (o *VM) GetAutostartOk() (*bool, bool)`

GetAutostartOk returns a tuple with the Autostart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutostart

`func (o *VM) SetAutostart(v bool)`

SetAutostart sets Autostart field to given value.

### HasAutostart

`func (o *VM) HasAutostart() bool`

HasAutostart returns a boolean if a field has been set.

### GetTime

`func (o *VM) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *VM) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *VM) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *VM) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetBootloader

`func (o *VM) GetBootloader() string`

GetBootloader returns the Bootloader field if non-nil, zero value otherwise.

### GetBootloaderOk

`func (o *VM) GetBootloaderOk() (*string, bool)`

GetBootloaderOk returns a tuple with the Bootloader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootloader

`func (o *VM) SetBootloader(v string)`

SetBootloader sets Bootloader field to given value.

### HasBootloader

`func (o *VM) HasBootloader() bool`

HasBootloader returns a boolean if a field has been set.

### GetCores

`func (o *VM) GetCores() int32`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *VM) GetCoresOk() (*int32, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *VM) SetCores(v int32)`

SetCores sets Cores field to given value.

### HasCores

`func (o *VM) HasCores() bool`

HasCores returns a boolean if a field has been set.

### GetThreads

`func (o *VM) GetThreads() int32`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *VM) GetThreadsOk() (*int32, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *VM) SetThreads(v int32)`

SetThreads sets Threads field to given value.

### HasThreads

`func (o *VM) HasThreads() bool`

HasThreads returns a boolean if a field has been set.

### GetHypervEnlightenments

`func (o *VM) GetHypervEnlightenments() bool`

GetHypervEnlightenments returns the HypervEnlightenments field if non-nil, zero value otherwise.

### GetHypervEnlightenmentsOk

`func (o *VM) GetHypervEnlightenmentsOk() (*bool, bool)`

GetHypervEnlightenmentsOk returns a tuple with the HypervEnlightenments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervEnlightenments

`func (o *VM) SetHypervEnlightenments(v bool)`

SetHypervEnlightenments sets HypervEnlightenments field to given value.

### HasHypervEnlightenments

`func (o *VM) HasHypervEnlightenments() bool`

HasHypervEnlightenments returns a boolean if a field has been set.

### GetShutdownTimeout

`func (o *VM) GetShutdownTimeout() int32`

GetShutdownTimeout returns the ShutdownTimeout field if non-nil, zero value otherwise.

### GetShutdownTimeoutOk

`func (o *VM) GetShutdownTimeoutOk() (*int32, bool)`

GetShutdownTimeoutOk returns a tuple with the ShutdownTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownTimeout

`func (o *VM) SetShutdownTimeout(v int32)`

SetShutdownTimeout sets ShutdownTimeout field to given value.

### HasShutdownTimeout

`func (o *VM) HasShutdownTimeout() bool`

HasShutdownTimeout returns a boolean if a field has been set.

### GetCpuMode

`func (o *VM) GetCpuMode() string`

GetCpuMode returns the CpuMode field if non-nil, zero value otherwise.

### GetCpuModeOk

`func (o *VM) GetCpuModeOk() (*string, bool)`

GetCpuModeOk returns a tuple with the CpuMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuMode

`func (o *VM) SetCpuMode(v string)`

SetCpuMode sets CpuMode field to given value.

### HasCpuMode

`func (o *VM) HasCpuMode() bool`

HasCpuMode returns a boolean if a field has been set.

### GetCpuModel

`func (o *VM) GetCpuModel() string`

GetCpuModel returns the CpuModel field if non-nil, zero value otherwise.

### GetCpuModelOk

`func (o *VM) GetCpuModelOk() (*string, bool)`

GetCpuModelOk returns a tuple with the CpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuModel

`func (o *VM) SetCpuModel(v string)`

SetCpuModel sets CpuModel field to given value.

### HasCpuModel

`func (o *VM) HasCpuModel() bool`

HasCpuModel returns a boolean if a field has been set.

### GetCpuset

`func (o *VM) GetCpuset() string`

GetCpuset returns the Cpuset field if non-nil, zero value otherwise.

### GetCpusetOk

`func (o *VM) GetCpusetOk() (*string, bool)`

GetCpusetOk returns a tuple with the Cpuset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuset

`func (o *VM) SetCpuset(v string)`

SetCpuset sets Cpuset field to given value.

### HasCpuset

`func (o *VM) HasCpuset() bool`

HasCpuset returns a boolean if a field has been set.

### GetNodeset

`func (o *VM) GetNodeset() string`

GetNodeset returns the Nodeset field if non-nil, zero value otherwise.

### GetNodesetOk

`func (o *VM) GetNodesetOk() (*string, bool)`

GetNodesetOk returns a tuple with the Nodeset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeset

`func (o *VM) SetNodeset(v string)`

SetNodeset sets Nodeset field to given value.

### HasNodeset

`func (o *VM) HasNodeset() bool`

HasNodeset returns a boolean if a field has been set.

### GetPinVcpus

`func (o *VM) GetPinVcpus() bool`

GetPinVcpus returns the PinVcpus field if non-nil, zero value otherwise.

### GetPinVcpusOk

`func (o *VM) GetPinVcpusOk() (*bool, bool)`

GetPinVcpusOk returns a tuple with the PinVcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinVcpus

`func (o *VM) SetPinVcpus(v bool)`

SetPinVcpus sets PinVcpus field to given value.

### HasPinVcpus

`func (o *VM) HasPinVcpus() bool`

HasPinVcpus returns a boolean if a field has been set.

### GetHideFromMsr

`func (o *VM) GetHideFromMsr() bool`

GetHideFromMsr returns the HideFromMsr field if non-nil, zero value otherwise.

### GetHideFromMsrOk

`func (o *VM) GetHideFromMsrOk() (*bool, bool)`

GetHideFromMsrOk returns a tuple with the HideFromMsr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideFromMsr

`func (o *VM) SetHideFromMsr(v bool)`

SetHideFromMsr sets HideFromMsr field to given value.

### HasHideFromMsr

`func (o *VM) HasHideFromMsr() bool`

HasHideFromMsr returns a boolean if a field has been set.

### GetSuspendOnSnapshot

`func (o *VM) GetSuspendOnSnapshot() bool`

GetSuspendOnSnapshot returns the SuspendOnSnapshot field if non-nil, zero value otherwise.

### GetSuspendOnSnapshotOk

`func (o *VM) GetSuspendOnSnapshotOk() (*bool, bool)`

GetSuspendOnSnapshotOk returns a tuple with the SuspendOnSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendOnSnapshot

`func (o *VM) SetSuspendOnSnapshot(v bool)`

SetSuspendOnSnapshot sets SuspendOnSnapshot field to given value.

### HasSuspendOnSnapshot

`func (o *VM) HasSuspendOnSnapshot() bool`

HasSuspendOnSnapshot returns a boolean if a field has been set.

### GetEnsureDisplayDevice

`func (o *VM) GetEnsureDisplayDevice() bool`

GetEnsureDisplayDevice returns the EnsureDisplayDevice field if non-nil, zero value otherwise.

### GetEnsureDisplayDeviceOk

`func (o *VM) GetEnsureDisplayDeviceOk() (*bool, bool)`

GetEnsureDisplayDeviceOk returns a tuple with the EnsureDisplayDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnsureDisplayDevice

`func (o *VM) SetEnsureDisplayDevice(v bool)`

SetEnsureDisplayDevice sets EnsureDisplayDevice field to given value.

### HasEnsureDisplayDevice

`func (o *VM) HasEnsureDisplayDevice() bool`

HasEnsureDisplayDevice returns a boolean if a field has been set.

### GetArchType

`func (o *VM) GetArchType() string`

GetArchType returns the ArchType field if non-nil, zero value otherwise.

### GetArchTypeOk

`func (o *VM) GetArchTypeOk() (*string, bool)`

GetArchTypeOk returns a tuple with the ArchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchType

`func (o *VM) SetArchType(v string)`

SetArchType sets ArchType field to given value.

### HasArchType

`func (o *VM) HasArchType() bool`

HasArchType returns a boolean if a field has been set.

### GetMachineType

`func (o *VM) GetMachineType() string`

GetMachineType returns the MachineType field if non-nil, zero value otherwise.

### GetMachineTypeOk

`func (o *VM) GetMachineTypeOk() (*string, bool)`

GetMachineTypeOk returns a tuple with the MachineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineType

`func (o *VM) SetMachineType(v string)`

SetMachineType sets MachineType field to given value.

### HasMachineType

`func (o *VM) HasMachineType() bool`

HasMachineType returns a boolean if a field has been set.

### GetUuid

`func (o *VM) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *VM) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *VM) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *VM) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetCommandLineArgs

`func (o *VM) GetCommandLineArgs() string`

GetCommandLineArgs returns the CommandLineArgs field if non-nil, zero value otherwise.

### GetCommandLineArgsOk

`func (o *VM) GetCommandLineArgsOk() (*string, bool)`

GetCommandLineArgsOk returns a tuple with the CommandLineArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandLineArgs

`func (o *VM) SetCommandLineArgs(v string)`

SetCommandLineArgs sets CommandLineArgs field to given value.

### HasCommandLineArgs

`func (o *VM) HasCommandLineArgs() bool`

HasCommandLineArgs returns a boolean if a field has been set.

### GetDevices

`func (o *VM) GetDevices() []VMDevice`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *VM) GetDevicesOk() (*[]VMDevice, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *VM) SetDevices(v []VMDevice)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *VM) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetStatus

`func (o *VM) GetStatus() VMStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VM) GetStatusOk() (*VMStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VM) SetStatus(v VMStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VM) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


