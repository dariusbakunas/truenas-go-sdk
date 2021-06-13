# VmDeviceCreate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dtype** | Pointer to **string** |  | [optional] 
**Vm** | Pointer to **int32** |  | [optional] 
**Attributes** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Order** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewVmDeviceCreate0

`func NewVmDeviceCreate0() *VmDeviceCreate0`

NewVmDeviceCreate0 instantiates a new VmDeviceCreate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmDeviceCreate0WithDefaults

`func NewVmDeviceCreate0WithDefaults() *VmDeviceCreate0`

NewVmDeviceCreate0WithDefaults instantiates a new VmDeviceCreate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDtype

`func (o *VmDeviceCreate0) GetDtype() string`

GetDtype returns the Dtype field if non-nil, zero value otherwise.

### GetDtypeOk

`func (o *VmDeviceCreate0) GetDtypeOk() (*string, bool)`

GetDtypeOk returns a tuple with the Dtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtype

`func (o *VmDeviceCreate0) SetDtype(v string)`

SetDtype sets Dtype field to given value.

### HasDtype

`func (o *VmDeviceCreate0) HasDtype() bool`

HasDtype returns a boolean if a field has been set.

### GetVm

`func (o *VmDeviceCreate0) GetVm() int32`

GetVm returns the Vm field if non-nil, zero value otherwise.

### GetVmOk

`func (o *VmDeviceCreate0) GetVmOk() (*int32, bool)`

GetVmOk returns a tuple with the Vm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVm

`func (o *VmDeviceCreate0) SetVm(v int32)`

SetVm sets Vm field to given value.

### HasVm

`func (o *VmDeviceCreate0) HasVm() bool`

HasVm returns a boolean if a field has been set.

### GetAttributes

`func (o *VmDeviceCreate0) GetAttributes() map[string]map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *VmDeviceCreate0) GetAttributesOk() (*map[string]map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *VmDeviceCreate0) SetAttributes(v map[string]map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *VmDeviceCreate0) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetOrder

`func (o *VmDeviceCreate0) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *VmDeviceCreate0) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *VmDeviceCreate0) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *VmDeviceCreate0) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrderNil

`func (o *VmDeviceCreate0) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *VmDeviceCreate0) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


