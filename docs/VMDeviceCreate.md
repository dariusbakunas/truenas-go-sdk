# VMDeviceCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dtype** | **string** |  | 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 

## Methods

### NewVMDeviceCreate

`func NewVMDeviceCreate(dtype string, ) *VMDeviceCreate`

NewVMDeviceCreate instantiates a new VMDeviceCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMDeviceCreateWithDefaults

`func NewVMDeviceCreateWithDefaults() *VMDeviceCreate`

NewVMDeviceCreateWithDefaults instantiates a new VMDeviceCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDtype

`func (o *VMDeviceCreate) GetDtype() string`

GetDtype returns the Dtype field if non-nil, zero value otherwise.

### GetDtypeOk

`func (o *VMDeviceCreate) GetDtypeOk() (*string, bool)`

GetDtypeOk returns a tuple with the Dtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtype

`func (o *VMDeviceCreate) SetDtype(v string)`

SetDtype sets Dtype field to given value.


### GetAttributes

`func (o *VMDeviceCreate) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *VMDeviceCreate) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *VMDeviceCreate) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *VMDeviceCreate) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetOrder

`func (o *VMDeviceCreate) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *VMDeviceCreate) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *VMDeviceCreate) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *VMDeviceCreate) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


