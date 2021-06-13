# VmDeviceUpdate1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dtype** | Pointer to **string** |  | [optional] 
**Vm** | Pointer to **int32** |  | [optional] 
**Attributes** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Order** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewVmDeviceUpdate1

`func NewVmDeviceUpdate1() *VmDeviceUpdate1`

NewVmDeviceUpdate1 instantiates a new VmDeviceUpdate1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmDeviceUpdate1WithDefaults

`func NewVmDeviceUpdate1WithDefaults() *VmDeviceUpdate1`

NewVmDeviceUpdate1WithDefaults instantiates a new VmDeviceUpdate1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDtype

`func (o *VmDeviceUpdate1) GetDtype() string`

GetDtype returns the Dtype field if non-nil, zero value otherwise.

### GetDtypeOk

`func (o *VmDeviceUpdate1) GetDtypeOk() (*string, bool)`

GetDtypeOk returns a tuple with the Dtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtype

`func (o *VmDeviceUpdate1) SetDtype(v string)`

SetDtype sets Dtype field to given value.

### HasDtype

`func (o *VmDeviceUpdate1) HasDtype() bool`

HasDtype returns a boolean if a field has been set.

### GetVm

`func (o *VmDeviceUpdate1) GetVm() int32`

GetVm returns the Vm field if non-nil, zero value otherwise.

### GetVmOk

`func (o *VmDeviceUpdate1) GetVmOk() (*int32, bool)`

GetVmOk returns a tuple with the Vm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVm

`func (o *VmDeviceUpdate1) SetVm(v int32)`

SetVm sets Vm field to given value.

### HasVm

`func (o *VmDeviceUpdate1) HasVm() bool`

HasVm returns a boolean if a field has been set.

### GetAttributes

`func (o *VmDeviceUpdate1) GetAttributes() map[string]map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *VmDeviceUpdate1) GetAttributesOk() (*map[string]map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *VmDeviceUpdate1) SetAttributes(v map[string]map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *VmDeviceUpdate1) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetOrder

`func (o *VmDeviceUpdate1) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *VmDeviceUpdate1) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *VmDeviceUpdate1) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *VmDeviceUpdate1) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrderNil

`func (o *VmDeviceUpdate1) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *VmDeviceUpdate1) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


