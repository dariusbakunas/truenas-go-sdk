# VMStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] 
**Pid** | Pointer to **int32** |  | [optional] 
**DomainState** | Pointer to **string** |  | [optional] 

## Methods

### NewVMStatus

`func NewVMStatus() *VMStatus`

NewVMStatus instantiates a new VMStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMStatusWithDefaults

`func NewVMStatusWithDefaults() *VMStatus`

NewVMStatusWithDefaults instantiates a new VMStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *VMStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VMStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VMStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *VMStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPid

`func (o *VMStatus) GetPid() int32`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *VMStatus) GetPidOk() (*int32, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *VMStatus) SetPid(v int32)`

SetPid sets Pid field to given value.

### HasPid

`func (o *VMStatus) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetDomainState

`func (o *VMStatus) GetDomainState() string`

GetDomainState returns the DomainState field if non-nil, zero value otherwise.

### GetDomainStateOk

`func (o *VMStatus) GetDomainStateOk() (*string, bool)`

GetDomainStateOk returns a tuple with the DomainState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainState

`func (o *VMStatus) SetDomainState(v string)`

SetDomainState sets DomainState field to given value.

### HasDomainState

`func (o *VMStatus) HasDomainState() bool`

HasDomainState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


