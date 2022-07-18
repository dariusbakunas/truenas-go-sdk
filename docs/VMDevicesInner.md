# VMDevicesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Dtype** | **string** |  | 
**Order** | Pointer to **int32** |  | [optional] 
**Vm** | Pointer to **int32** |  | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewVMDevicesInner

`func NewVMDevicesInner(id int32, dtype string, ) *VMDevicesInner`

NewVMDevicesInner instantiates a new VMDevicesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMDevicesInnerWithDefaults

`func NewVMDevicesInnerWithDefaults() *VMDevicesInner`

NewVMDevicesInnerWithDefaults instantiates a new VMDevicesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VMDevicesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VMDevicesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VMDevicesInner) SetId(v int32)`

SetId sets Id field to given value.


### GetDtype

`func (o *VMDevicesInner) GetDtype() string`

GetDtype returns the Dtype field if non-nil, zero value otherwise.

### GetDtypeOk

`func (o *VMDevicesInner) GetDtypeOk() (*string, bool)`

GetDtypeOk returns a tuple with the Dtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtype

`func (o *VMDevicesInner) SetDtype(v string)`

SetDtype sets Dtype field to given value.


### GetOrder

`func (o *VMDevicesInner) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *VMDevicesInner) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *VMDevicesInner) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *VMDevicesInner) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetVm

`func (o *VMDevicesInner) GetVm() int32`

GetVm returns the Vm field if non-nil, zero value otherwise.

### GetVmOk

`func (o *VMDevicesInner) GetVmOk() (*int32, bool)`

GetVmOk returns a tuple with the Vm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVm

`func (o *VMDevicesInner) SetVm(v int32)`

SetVm sets Vm field to given value.

### HasVm

`func (o *VMDevicesInner) HasVm() bool`

HasVm returns a boolean if a field has been set.

### GetAttributes

`func (o *VMDevicesInner) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *VMDevicesInner) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *VMDevicesInner) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *VMDevicesInner) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


