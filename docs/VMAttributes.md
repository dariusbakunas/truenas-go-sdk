# VMAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Mac** | Pointer to **string** |  | [optional] 
**NicAttach** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Wait** | Pointer to **bool** |  | [optional] 
**VncPort** | Pointer to **int32** |  | [optional] 
**VncResolution** | Pointer to **string** |  | [optional] 
**VncBind** | Pointer to **string** |  | [optional] 
**VncPassword** | Pointer to **string** |  | [optional] 
**VncWeb** | Pointer to **bool** |  | [optional] 

## Methods

### NewVMAttributes

`func NewVMAttributes() *VMAttributes`

NewVMAttributes instantiates a new VMAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMAttributesWithDefaults

`func NewVMAttributesWithDefaults() *VMAttributes`

NewVMAttributesWithDefaults instantiates a new VMAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VMAttributes) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VMAttributes) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VMAttributes) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VMAttributes) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMac

`func (o *VMAttributes) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *VMAttributes) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *VMAttributes) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *VMAttributes) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetNicAttach

`func (o *VMAttributes) GetNicAttach() string`

GetNicAttach returns the NicAttach field if non-nil, zero value otherwise.

### GetNicAttachOk

`func (o *VMAttributes) GetNicAttachOk() (*string, bool)`

GetNicAttachOk returns a tuple with the NicAttach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicAttach

`func (o *VMAttributes) SetNicAttach(v string)`

SetNicAttach sets NicAttach field to given value.

### HasNicAttach

`func (o *VMAttributes) HasNicAttach() bool`

HasNicAttach returns a boolean if a field has been set.

### GetPath

`func (o *VMAttributes) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *VMAttributes) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *VMAttributes) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *VMAttributes) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetWait

`func (o *VMAttributes) GetWait() bool`

GetWait returns the Wait field if non-nil, zero value otherwise.

### GetWaitOk

`func (o *VMAttributes) GetWaitOk() (*bool, bool)`

GetWaitOk returns a tuple with the Wait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWait

`func (o *VMAttributes) SetWait(v bool)`

SetWait sets Wait field to given value.

### HasWait

`func (o *VMAttributes) HasWait() bool`

HasWait returns a boolean if a field has been set.

### GetVncPort

`func (o *VMAttributes) GetVncPort() int32`

GetVncPort returns the VncPort field if non-nil, zero value otherwise.

### GetVncPortOk

`func (o *VMAttributes) GetVncPortOk() (*int32, bool)`

GetVncPortOk returns a tuple with the VncPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVncPort

`func (o *VMAttributes) SetVncPort(v int32)`

SetVncPort sets VncPort field to given value.

### HasVncPort

`func (o *VMAttributes) HasVncPort() bool`

HasVncPort returns a boolean if a field has been set.

### GetVncResolution

`func (o *VMAttributes) GetVncResolution() string`

GetVncResolution returns the VncResolution field if non-nil, zero value otherwise.

### GetVncResolutionOk

`func (o *VMAttributes) GetVncResolutionOk() (*string, bool)`

GetVncResolutionOk returns a tuple with the VncResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVncResolution

`func (o *VMAttributes) SetVncResolution(v string)`

SetVncResolution sets VncResolution field to given value.

### HasVncResolution

`func (o *VMAttributes) HasVncResolution() bool`

HasVncResolution returns a boolean if a field has been set.

### GetVncBind

`func (o *VMAttributes) GetVncBind() string`

GetVncBind returns the VncBind field if non-nil, zero value otherwise.

### GetVncBindOk

`func (o *VMAttributes) GetVncBindOk() (*string, bool)`

GetVncBindOk returns a tuple with the VncBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVncBind

`func (o *VMAttributes) SetVncBind(v string)`

SetVncBind sets VncBind field to given value.

### HasVncBind

`func (o *VMAttributes) HasVncBind() bool`

HasVncBind returns a boolean if a field has been set.

### GetVncPassword

`func (o *VMAttributes) GetVncPassword() string`

GetVncPassword returns the VncPassword field if non-nil, zero value otherwise.

### GetVncPasswordOk

`func (o *VMAttributes) GetVncPasswordOk() (*string, bool)`

GetVncPasswordOk returns a tuple with the VncPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVncPassword

`func (o *VMAttributes) SetVncPassword(v string)`

SetVncPassword sets VncPassword field to given value.

### HasVncPassword

`func (o *VMAttributes) HasVncPassword() bool`

HasVncPassword returns a boolean if a field has been set.

### GetVncWeb

`func (o *VMAttributes) GetVncWeb() bool`

GetVncWeb returns the VncWeb field if non-nil, zero value otherwise.

### GetVncWebOk

`func (o *VMAttributes) GetVncWebOk() (*bool, bool)`

GetVncWebOk returns a tuple with the VncWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVncWeb

`func (o *VMAttributes) SetVncWeb(v bool)`

SetVncWeb sets VncWeb field to given value.

### HasVncWeb

`func (o *VMAttributes) HasVncWeb() bool`

HasVncWeb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


