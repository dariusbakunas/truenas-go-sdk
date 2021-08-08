# VMDevices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Dtype** | **string** |  | 
**Order** | Pointer to **int32** |  | [optional] 
**Vm** | Pointer to **int32** |  | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewVMDevices

`func NewVMDevices(id int32, dtype string, ) *VMDevices`

NewVMDevices instantiates a new VMDevices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMDevicesWithDefaults

`func NewVMDevicesWithDefaults() *VMDevices`

NewVMDevicesWithDefaults instantiates a new VMDevices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VMDevices) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VMDevices) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VMDevices) SetId(v int32)`

SetId sets Id field to given value.


### GetDtype

`func (o *VMDevices) GetDtype() string`

GetDtype returns the Dtype field if non-nil, zero value otherwise.

### GetDtypeOk

`func (o *VMDevices) GetDtypeOk() (*string, bool)`

GetDtypeOk returns a tuple with the Dtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtype

`func (o *VMDevices) SetDtype(v string)`

SetDtype sets Dtype field to given value.


### GetOrder

`func (o *VMDevices) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *VMDevices) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *VMDevices) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *VMDevices) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetVm

`func (o *VMDevices) GetVm() int32`

GetVm returns the Vm field if non-nil, zero value otherwise.

### GetVmOk

`func (o *VMDevices) GetVmOk() (*int32, bool)`

GetVmOk returns a tuple with the Vm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVm

`func (o *VMDevices) SetVm(v int32)`

SetVm sets Vm field to given value.

### HasVm

`func (o *VMDevices) HasVm() bool`

HasVm returns a boolean if a field has been set.

### GetAttributes

`func (o *VMDevices) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *VMDevices) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *VMDevices) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *VMDevices) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


