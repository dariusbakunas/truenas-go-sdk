# VmCreate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Vcpus** | Pointer to **int32** |  | [optional] 
**Cores** | Pointer to **int32** |  | [optional] 
**Threads** | Pointer to **int32** |  | [optional] 
**Memory** | Pointer to **int32** |  | [optional] 
**Bootloader** | Pointer to **string** |  | [optional] 
**Grubconfig** | Pointer to **NullableString** |  | [optional] 
**Devices** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Autostart** | Pointer to **bool** |  | [optional] 
**Time** | Pointer to **string** |  | [optional] 
**ShutdownTimeout** | Pointer to **int32** |  | [optional] 

## Methods

### NewVmCreate0

`func NewVmCreate0() *VmCreate0`

NewVmCreate0 instantiates a new VmCreate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmCreate0WithDefaults

`func NewVmCreate0WithDefaults() *VmCreate0`

NewVmCreate0WithDefaults instantiates a new VmCreate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VmCreate0) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VmCreate0) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VmCreate0) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VmCreate0) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *VmCreate0) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VmCreate0) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VmCreate0) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VmCreate0) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVcpus

`func (o *VmCreate0) GetVcpus() int32`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *VmCreate0) GetVcpusOk() (*int32, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *VmCreate0) SetVcpus(v int32)`

SetVcpus sets Vcpus field to given value.

### HasVcpus

`func (o *VmCreate0) HasVcpus() bool`

HasVcpus returns a boolean if a field has been set.

### GetCores

`func (o *VmCreate0) GetCores() int32`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *VmCreate0) GetCoresOk() (*int32, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *VmCreate0) SetCores(v int32)`

SetCores sets Cores field to given value.

### HasCores

`func (o *VmCreate0) HasCores() bool`

HasCores returns a boolean if a field has been set.

### GetThreads

`func (o *VmCreate0) GetThreads() int32`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *VmCreate0) GetThreadsOk() (*int32, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *VmCreate0) SetThreads(v int32)`

SetThreads sets Threads field to given value.

### HasThreads

`func (o *VmCreate0) HasThreads() bool`

HasThreads returns a boolean if a field has been set.

### GetMemory

`func (o *VmCreate0) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *VmCreate0) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *VmCreate0) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *VmCreate0) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetBootloader

`func (o *VmCreate0) GetBootloader() string`

GetBootloader returns the Bootloader field if non-nil, zero value otherwise.

### GetBootloaderOk

`func (o *VmCreate0) GetBootloaderOk() (*string, bool)`

GetBootloaderOk returns a tuple with the Bootloader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootloader

`func (o *VmCreate0) SetBootloader(v string)`

SetBootloader sets Bootloader field to given value.

### HasBootloader

`func (o *VmCreate0) HasBootloader() bool`

HasBootloader returns a boolean if a field has been set.

### GetGrubconfig

`func (o *VmCreate0) GetGrubconfig() string`

GetGrubconfig returns the Grubconfig field if non-nil, zero value otherwise.

### GetGrubconfigOk

`func (o *VmCreate0) GetGrubconfigOk() (*string, bool)`

GetGrubconfigOk returns a tuple with the Grubconfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrubconfig

`func (o *VmCreate0) SetGrubconfig(v string)`

SetGrubconfig sets Grubconfig field to given value.

### HasGrubconfig

`func (o *VmCreate0) HasGrubconfig() bool`

HasGrubconfig returns a boolean if a field has been set.

### SetGrubconfigNil

`func (o *VmCreate0) SetGrubconfigNil(b bool)`

 SetGrubconfigNil sets the value for Grubconfig to be an explicit nil

### UnsetGrubconfig
`func (o *VmCreate0) UnsetGrubconfig()`

UnsetGrubconfig ensures that no value is present for Grubconfig, not even an explicit nil
### GetDevices

`func (o *VmCreate0) GetDevices() []map[string]interface{}`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *VmCreate0) GetDevicesOk() (*[]map[string]interface{}, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *VmCreate0) SetDevices(v []map[string]interface{})`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *VmCreate0) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetAutostart

`func (o *VmCreate0) GetAutostart() bool`

GetAutostart returns the Autostart field if non-nil, zero value otherwise.

### GetAutostartOk

`func (o *VmCreate0) GetAutostartOk() (*bool, bool)`

GetAutostartOk returns a tuple with the Autostart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutostart

`func (o *VmCreate0) SetAutostart(v bool)`

SetAutostart sets Autostart field to given value.

### HasAutostart

`func (o *VmCreate0) HasAutostart() bool`

HasAutostart returns a boolean if a field has been set.

### GetTime

`func (o *VmCreate0) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *VmCreate0) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *VmCreate0) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *VmCreate0) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetShutdownTimeout

`func (o *VmCreate0) GetShutdownTimeout() int32`

GetShutdownTimeout returns the ShutdownTimeout field if non-nil, zero value otherwise.

### GetShutdownTimeoutOk

`func (o *VmCreate0) GetShutdownTimeoutOk() (*int32, bool)`

GetShutdownTimeoutOk returns a tuple with the ShutdownTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownTimeout

`func (o *VmCreate0) SetShutdownTimeout(v int32)`

SetShutdownTimeout sets ShutdownTimeout field to given value.

### HasShutdownTimeout

`func (o *VmCreate0) HasShutdownTimeout() bool`

HasShutdownTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


