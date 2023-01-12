# UpdateVMParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**CpuMode** | Pointer to **string** |  | [optional] [default to "CUSTOM"]
**CpuModel** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Vcpus** | Pointer to **int32** | Maximum of 16 guest virtual CPUs are allowed. By default, every virtual CPU is configured as a separate package. Multiple cores can be configured per CPU by specifying &#x60;cores&#x60; attributes. &#x60;vcpus&#x60; specifies total number of CPU sockets. &#x60;cores&#x60; specifies number of cores per socket. &#x60;threads&#x60; specifies number of threads per core. | [optional] [default to 1]
**Cores** | Pointer to **int32** | Maximum of 16 guest virtual CPUs are allowed. By default, every virtual CPU is configured as a separate package. Multiple cores can be configured per CPU by specifying &#x60;cores&#x60; attributes. &#x60;vcpus&#x60; specifies total number of CPU sockets. &#x60;cores&#x60; specifies number of cores per socket. &#x60;threads&#x60; specifies number of threads per core. | [optional] [default to 1]
**Threads** | Pointer to **int32** | Maximum of 16 guest virtual CPUs are allowed. By default, every virtual CPU is configured as a separate package. Multiple cores can be configured per CPU by specifying &#x60;cores&#x60; attributes. &#x60;vcpus&#x60; specifies total number of CPU sockets. &#x60;cores&#x60; specifies number of cores per socket. &#x60;threads&#x60; specifies number of threads per core. | [optional] [default to 1]
**Memory** | Pointer to **int64** |  | [optional] 
**Bootloader** | Pointer to **string** |  | [optional] [default to "UEFI"]
**Devices** | Pointer to [**[]VMDevice**](VMDevice.md) | &#x60;devices&#x60; is a list of virtualized hardware to add to the newly created Virtual Machine. Failure to attach a device destroys the VM and any resources allocated by the VM devices. | [optional] 
**Autostart** | Pointer to **bool** |  | [optional] [default to true]
**HideFromMsr** | Pointer to **bool** | &#x60;hide_from_msr&#x60; is a boolean which when set will hide the KVM hypervisor from standard MSR based discovery and is useful to enable when doing GPU passthrough. | [optional] [default to false]
**EnsureDisplayDevice** | Pointer to **bool** | &#x60;ensure_display_device&#x60; when set ( the default ) will ensure that the guest always has access to a video device. For headless installations like ubuntu server this is required for the guest to operate properly. However for cases where consumer would like to use GPU passthrough and does not want a display device added should set this to &#x60;false&#x60;. | [optional] [default to true]
**Time** | Pointer to **string** |  | [optional] [default to "LOCAL"]
**ShutdownTimeout** | Pointer to **int32** | &#x60;shutdown_timeout&#x60; indicates the time in seconds the system waits for the VM to cleanly shutdown. During system shutdown, if the VM hasn&#39;t exited after a hardware shutdown signal has been sent by the system within &#x60;shutdown_timeout&#x60; seconds, system initiates poweroff for the VM to stop it. | [optional] [default to 90]
**ArchType** | Pointer to **NullableString** | &#x60;arch_type&#x60; refers to architecture type and can be specified for the guest. By default the value is &#x60;null&#x60; and system in this case will choose a reasonable default based on host. &#x60;machine_type&#x60; refers to machine type of the guest based on the architecture type selected with &#x60;arch_type&#x60;. By default the value is &#x60;null&#x60; and system in this case will choose a reasonable default based on &#x60;arch_type&#x60; configuration. | [optional] 
**MachineType** | Pointer to **NullableString** | &#x60;machine_type&#x60; refers to machine type of the guest based on the architecture type selected with &#x60;arch_type&#x60;. By default the value is &#x60;null&#x60; and system in this case will choose a reasonable default based on &#x60;arch_type&#x60; configuration. | [optional] 
**Uuid** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateVMParams

`func NewUpdateVMParams() *UpdateVMParams`

NewUpdateVMParams instantiates a new UpdateVMParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVMParamsWithDefaults

`func NewUpdateVMParamsWithDefaults() *UpdateVMParams`

NewUpdateVMParamsWithDefaults instantiates a new UpdateVMParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateVMParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateVMParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateVMParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateVMParams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCpuMode

`func (o *UpdateVMParams) GetCpuMode() string`

GetCpuMode returns the CpuMode field if non-nil, zero value otherwise.

### GetCpuModeOk

`func (o *UpdateVMParams) GetCpuModeOk() (*string, bool)`

GetCpuModeOk returns a tuple with the CpuMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuMode

`func (o *UpdateVMParams) SetCpuMode(v string)`

SetCpuMode sets CpuMode field to given value.

### HasCpuMode

`func (o *UpdateVMParams) HasCpuMode() bool`

HasCpuMode returns a boolean if a field has been set.

### GetCpuModel

`func (o *UpdateVMParams) GetCpuModel() string`

GetCpuModel returns the CpuModel field if non-nil, zero value otherwise.

### GetCpuModelOk

`func (o *UpdateVMParams) GetCpuModelOk() (*string, bool)`

GetCpuModelOk returns a tuple with the CpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuModel

`func (o *UpdateVMParams) SetCpuModel(v string)`

SetCpuModel sets CpuModel field to given value.

### HasCpuModel

`func (o *UpdateVMParams) HasCpuModel() bool`

HasCpuModel returns a boolean if a field has been set.

### SetCpuModelNil

`func (o *UpdateVMParams) SetCpuModelNil(b bool)`

 SetCpuModelNil sets the value for CpuModel to be an explicit nil

### UnsetCpuModel
`func (o *UpdateVMParams) UnsetCpuModel()`

UnsetCpuModel ensures that no value is present for CpuModel, not even an explicit nil
### GetDescription

`func (o *UpdateVMParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateVMParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateVMParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateVMParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVcpus

`func (o *UpdateVMParams) GetVcpus() int32`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *UpdateVMParams) GetVcpusOk() (*int32, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *UpdateVMParams) SetVcpus(v int32)`

SetVcpus sets Vcpus field to given value.

### HasVcpus

`func (o *UpdateVMParams) HasVcpus() bool`

HasVcpus returns a boolean if a field has been set.

### GetCores

`func (o *UpdateVMParams) GetCores() int32`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *UpdateVMParams) GetCoresOk() (*int32, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *UpdateVMParams) SetCores(v int32)`

SetCores sets Cores field to given value.

### HasCores

`func (o *UpdateVMParams) HasCores() bool`

HasCores returns a boolean if a field has been set.

### GetThreads

`func (o *UpdateVMParams) GetThreads() int32`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *UpdateVMParams) GetThreadsOk() (*int32, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *UpdateVMParams) SetThreads(v int32)`

SetThreads sets Threads field to given value.

### HasThreads

`func (o *UpdateVMParams) HasThreads() bool`

HasThreads returns a boolean if a field has been set.

### GetMemory

`func (o *UpdateVMParams) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *UpdateVMParams) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *UpdateVMParams) SetMemory(v int64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *UpdateVMParams) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetBootloader

`func (o *UpdateVMParams) GetBootloader() string`

GetBootloader returns the Bootloader field if non-nil, zero value otherwise.

### GetBootloaderOk

`func (o *UpdateVMParams) GetBootloaderOk() (*string, bool)`

GetBootloaderOk returns a tuple with the Bootloader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootloader

`func (o *UpdateVMParams) SetBootloader(v string)`

SetBootloader sets Bootloader field to given value.

### HasBootloader

`func (o *UpdateVMParams) HasBootloader() bool`

HasBootloader returns a boolean if a field has been set.

### GetDevices

`func (o *UpdateVMParams) GetDevices() []VMDevice`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *UpdateVMParams) GetDevicesOk() (*[]VMDevice, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *UpdateVMParams) SetDevices(v []VMDevice)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *UpdateVMParams) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetAutostart

`func (o *UpdateVMParams) GetAutostart() bool`

GetAutostart returns the Autostart field if non-nil, zero value otherwise.

### GetAutostartOk

`func (o *UpdateVMParams) GetAutostartOk() (*bool, bool)`

GetAutostartOk returns a tuple with the Autostart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutostart

`func (o *UpdateVMParams) SetAutostart(v bool)`

SetAutostart sets Autostart field to given value.

### HasAutostart

`func (o *UpdateVMParams) HasAutostart() bool`

HasAutostart returns a boolean if a field has been set.

### GetHideFromMsr

`func (o *UpdateVMParams) GetHideFromMsr() bool`

GetHideFromMsr returns the HideFromMsr field if non-nil, zero value otherwise.

### GetHideFromMsrOk

`func (o *UpdateVMParams) GetHideFromMsrOk() (*bool, bool)`

GetHideFromMsrOk returns a tuple with the HideFromMsr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideFromMsr

`func (o *UpdateVMParams) SetHideFromMsr(v bool)`

SetHideFromMsr sets HideFromMsr field to given value.

### HasHideFromMsr

`func (o *UpdateVMParams) HasHideFromMsr() bool`

HasHideFromMsr returns a boolean if a field has been set.

### GetEnsureDisplayDevice

`func (o *UpdateVMParams) GetEnsureDisplayDevice() bool`

GetEnsureDisplayDevice returns the EnsureDisplayDevice field if non-nil, zero value otherwise.

### GetEnsureDisplayDeviceOk

`func (o *UpdateVMParams) GetEnsureDisplayDeviceOk() (*bool, bool)`

GetEnsureDisplayDeviceOk returns a tuple with the EnsureDisplayDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnsureDisplayDevice

`func (o *UpdateVMParams) SetEnsureDisplayDevice(v bool)`

SetEnsureDisplayDevice sets EnsureDisplayDevice field to given value.

### HasEnsureDisplayDevice

`func (o *UpdateVMParams) HasEnsureDisplayDevice() bool`

HasEnsureDisplayDevice returns a boolean if a field has been set.

### GetTime

`func (o *UpdateVMParams) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *UpdateVMParams) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *UpdateVMParams) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *UpdateVMParams) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetShutdownTimeout

`func (o *UpdateVMParams) GetShutdownTimeout() int32`

GetShutdownTimeout returns the ShutdownTimeout field if non-nil, zero value otherwise.

### GetShutdownTimeoutOk

`func (o *UpdateVMParams) GetShutdownTimeoutOk() (*int32, bool)`

GetShutdownTimeoutOk returns a tuple with the ShutdownTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownTimeout

`func (o *UpdateVMParams) SetShutdownTimeout(v int32)`

SetShutdownTimeout sets ShutdownTimeout field to given value.

### HasShutdownTimeout

`func (o *UpdateVMParams) HasShutdownTimeout() bool`

HasShutdownTimeout returns a boolean if a field has been set.

### GetArchType

`func (o *UpdateVMParams) GetArchType() string`

GetArchType returns the ArchType field if non-nil, zero value otherwise.

### GetArchTypeOk

`func (o *UpdateVMParams) GetArchTypeOk() (*string, bool)`

GetArchTypeOk returns a tuple with the ArchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchType

`func (o *UpdateVMParams) SetArchType(v string)`

SetArchType sets ArchType field to given value.

### HasArchType

`func (o *UpdateVMParams) HasArchType() bool`

HasArchType returns a boolean if a field has been set.

### SetArchTypeNil

`func (o *UpdateVMParams) SetArchTypeNil(b bool)`

 SetArchTypeNil sets the value for ArchType to be an explicit nil

### UnsetArchType
`func (o *UpdateVMParams) UnsetArchType()`

UnsetArchType ensures that no value is present for ArchType, not even an explicit nil
### GetMachineType

`func (o *UpdateVMParams) GetMachineType() string`

GetMachineType returns the MachineType field if non-nil, zero value otherwise.

### GetMachineTypeOk

`func (o *UpdateVMParams) GetMachineTypeOk() (*string, bool)`

GetMachineTypeOk returns a tuple with the MachineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineType

`func (o *UpdateVMParams) SetMachineType(v string)`

SetMachineType sets MachineType field to given value.

### HasMachineType

`func (o *UpdateVMParams) HasMachineType() bool`

HasMachineType returns a boolean if a field has been set.

### SetMachineTypeNil

`func (o *UpdateVMParams) SetMachineTypeNil(b bool)`

 SetMachineTypeNil sets the value for MachineType to be an explicit nil

### UnsetMachineType
`func (o *UpdateVMParams) UnsetMachineType()`

UnsetMachineType ensures that no value is present for MachineType, not even an explicit nil
### GetUuid

`func (o *UpdateVMParams) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *UpdateVMParams) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *UpdateVMParams) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *UpdateVMParams) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *UpdateVMParams) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *UpdateVMParams) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


