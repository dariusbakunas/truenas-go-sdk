# VmUpdate1

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

### NewVmUpdate1

`func NewVmUpdate1() *VmUpdate1`

NewVmUpdate1 instantiates a new VmUpdate1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmUpdate1WithDefaults

`func NewVmUpdate1WithDefaults() *VmUpdate1`

NewVmUpdate1WithDefaults instantiates a new VmUpdate1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VmUpdate1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VmUpdate1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VmUpdate1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VmUpdate1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *VmUpdate1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VmUpdate1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VmUpdate1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VmUpdate1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVcpus

`func (o *VmUpdate1) GetVcpus() int32`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *VmUpdate1) GetVcpusOk() (*int32, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *VmUpdate1) SetVcpus(v int32)`

SetVcpus sets Vcpus field to given value.

### HasVcpus

`func (o *VmUpdate1) HasVcpus() bool`

HasVcpus returns a boolean if a field has been set.

### GetCores

`func (o *VmUpdate1) GetCores() int32`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *VmUpdate1) GetCoresOk() (*int32, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *VmUpdate1) SetCores(v int32)`

SetCores sets Cores field to given value.

### HasCores

`func (o *VmUpdate1) HasCores() bool`

HasCores returns a boolean if a field has been set.

### GetThreads

`func (o *VmUpdate1) GetThreads() int32`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *VmUpdate1) GetThreadsOk() (*int32, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *VmUpdate1) SetThreads(v int32)`

SetThreads sets Threads field to given value.

### HasThreads

`func (o *VmUpdate1) HasThreads() bool`

HasThreads returns a boolean if a field has been set.

### GetMemory

`func (o *VmUpdate1) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *VmUpdate1) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *VmUpdate1) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *VmUpdate1) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetBootloader

`func (o *VmUpdate1) GetBootloader() string`

GetBootloader returns the Bootloader field if non-nil, zero value otherwise.

### GetBootloaderOk

`func (o *VmUpdate1) GetBootloaderOk() (*string, bool)`

GetBootloaderOk returns a tuple with the Bootloader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootloader

`func (o *VmUpdate1) SetBootloader(v string)`

SetBootloader sets Bootloader field to given value.

### HasBootloader

`func (o *VmUpdate1) HasBootloader() bool`

HasBootloader returns a boolean if a field has been set.

### GetGrubconfig

`func (o *VmUpdate1) GetGrubconfig() string`

GetGrubconfig returns the Grubconfig field if non-nil, zero value otherwise.

### GetGrubconfigOk

`func (o *VmUpdate1) GetGrubconfigOk() (*string, bool)`

GetGrubconfigOk returns a tuple with the Grubconfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrubconfig

`func (o *VmUpdate1) SetGrubconfig(v string)`

SetGrubconfig sets Grubconfig field to given value.

### HasGrubconfig

`func (o *VmUpdate1) HasGrubconfig() bool`

HasGrubconfig returns a boolean if a field has been set.

### SetGrubconfigNil

`func (o *VmUpdate1) SetGrubconfigNil(b bool)`

 SetGrubconfigNil sets the value for Grubconfig to be an explicit nil

### UnsetGrubconfig
`func (o *VmUpdate1) UnsetGrubconfig()`

UnsetGrubconfig ensures that no value is present for Grubconfig, not even an explicit nil
### GetDevices

`func (o *VmUpdate1) GetDevices() []map[string]interface{}`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *VmUpdate1) GetDevicesOk() (*[]map[string]interface{}, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *VmUpdate1) SetDevices(v []map[string]interface{})`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *VmUpdate1) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetAutostart

`func (o *VmUpdate1) GetAutostart() bool`

GetAutostart returns the Autostart field if non-nil, zero value otherwise.

### GetAutostartOk

`func (o *VmUpdate1) GetAutostartOk() (*bool, bool)`

GetAutostartOk returns a tuple with the Autostart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutostart

`func (o *VmUpdate1) SetAutostart(v bool)`

SetAutostart sets Autostart field to given value.

### HasAutostart

`func (o *VmUpdate1) HasAutostart() bool`

HasAutostart returns a boolean if a field has been set.

### GetTime

`func (o *VmUpdate1) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *VmUpdate1) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *VmUpdate1) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *VmUpdate1) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetShutdownTimeout

`func (o *VmUpdate1) GetShutdownTimeout() int32`

GetShutdownTimeout returns the ShutdownTimeout field if non-nil, zero value otherwise.

### GetShutdownTimeoutOk

`func (o *VmUpdate1) GetShutdownTimeoutOk() (*int32, bool)`

GetShutdownTimeoutOk returns a tuple with the ShutdownTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownTimeout

`func (o *VmUpdate1) SetShutdownTimeout(v int32)`

SetShutdownTimeout sets ShutdownTimeout field to given value.

### HasShutdownTimeout

`func (o *VmUpdate1) HasShutdownTimeout() bool`

HasShutdownTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


