# VMDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Dtype** | **string** |  | 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**Vm** | Pointer to **int32** |  | [optional] 

## Methods

### NewVMDevice

`func NewVMDevice(dtype string, ) *VMDevice`

NewVMDevice instantiates a new VMDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMDeviceWithDefaults

`func NewVMDeviceWithDefaults() *VMDevice`

NewVMDeviceWithDefaults instantiates a new VMDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VMDevice) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VMDevice) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VMDevice) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VMDevice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDtype

`func (o *VMDevice) GetDtype() string`

GetDtype returns the Dtype field if non-nil, zero value otherwise.

### GetDtypeOk

`func (o *VMDevice) GetDtypeOk() (*string, bool)`

GetDtypeOk returns a tuple with the Dtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtype

`func (o *VMDevice) SetDtype(v string)`

SetDtype sets Dtype field to given value.


### GetAttributes

`func (o *VMDevice) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *VMDevice) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *VMDevice) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *VMDevice) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetOrder

`func (o *VMDevice) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *VMDevice) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *VMDevice) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *VMDevice) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetVm

`func (o *VMDevice) GetVm() int32`

GetVm returns the Vm field if non-nil, zero value otherwise.

### GetVmOk

`func (o *VMDevice) GetVmOk() (*int32, bool)`

GetVmOk returns a tuple with the Vm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVm

`func (o *VMDevice) SetVm(v int32)`

SetVm sets Vm field to given value.

### HasVm

`func (o *VMDevice) HasVm() bool`

HasVm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


