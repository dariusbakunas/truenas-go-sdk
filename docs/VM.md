# VM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Vcpus** | Pointer to **int32** |  | [optional] 
**Memory** | Pointer to **int32** |  | [optional] 
**Autostart** | Pointer to **bool** |  | [optional] 
**Time** | Pointer to **string** |  | [optional] 
**Bootloader** | Pointer to **string** |  | [optional] 
**Cores** | Pointer to **int32** |  | [optional] 
**Threads** | Pointer to **int32** |  | [optional] 
**ShutdownTimeout** | Pointer to **int32** |  | [optional] 
**Devices** | Pointer to [**[]VMDevicesInner**](VMDevicesInner.md) |  | [optional] 
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

`func (o *VM) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *VM) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *VM) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *VM) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

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

### GetDevices

`func (o *VM) GetDevices() []VMDevicesInner`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *VM) GetDevicesOk() (*[]VMDevicesInner, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *VM) SetDevices(v []VMDevicesInner)`

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


