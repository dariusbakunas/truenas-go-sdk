# IscsiGlobalUpdate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Basename** | Pointer to **string** |  | [optional] 
**IsnsServers** | Pointer to **[]string** |  | [optional] 
**PoolAvailThreshold** | Pointer to **NullableInt32** |  | [optional] 
**Alua** | Pointer to **bool** |  | [optional] 

## Methods

### NewIscsiGlobalUpdate0

`func NewIscsiGlobalUpdate0() *IscsiGlobalUpdate0`

NewIscsiGlobalUpdate0 instantiates a new IscsiGlobalUpdate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIscsiGlobalUpdate0WithDefaults

`func NewIscsiGlobalUpdate0WithDefaults() *IscsiGlobalUpdate0`

NewIscsiGlobalUpdate0WithDefaults instantiates a new IscsiGlobalUpdate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasename

`func (o *IscsiGlobalUpdate0) GetBasename() string`

GetBasename returns the Basename field if non-nil, zero value otherwise.

### GetBasenameOk

`func (o *IscsiGlobalUpdate0) GetBasenameOk() (*string, bool)`

GetBasenameOk returns a tuple with the Basename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasename

`func (o *IscsiGlobalUpdate0) SetBasename(v string)`

SetBasename sets Basename field to given value.

### HasBasename

`func (o *IscsiGlobalUpdate0) HasBasename() bool`

HasBasename returns a boolean if a field has been set.

### GetIsnsServers

`func (o *IscsiGlobalUpdate0) GetIsnsServers() []string`

GetIsnsServers returns the IsnsServers field if non-nil, zero value otherwise.

### GetIsnsServersOk

`func (o *IscsiGlobalUpdate0) GetIsnsServersOk() (*[]string, bool)`

GetIsnsServersOk returns a tuple with the IsnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsnsServers

`func (o *IscsiGlobalUpdate0) SetIsnsServers(v []string)`

SetIsnsServers sets IsnsServers field to given value.

### HasIsnsServers

`func (o *IscsiGlobalUpdate0) HasIsnsServers() bool`

HasIsnsServers returns a boolean if a field has been set.

### GetPoolAvailThreshold

`func (o *IscsiGlobalUpdate0) GetPoolAvailThreshold() int32`

GetPoolAvailThreshold returns the PoolAvailThreshold field if non-nil, zero value otherwise.

### GetPoolAvailThresholdOk

`func (o *IscsiGlobalUpdate0) GetPoolAvailThresholdOk() (*int32, bool)`

GetPoolAvailThresholdOk returns a tuple with the PoolAvailThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAvailThreshold

`func (o *IscsiGlobalUpdate0) SetPoolAvailThreshold(v int32)`

SetPoolAvailThreshold sets PoolAvailThreshold field to given value.

### HasPoolAvailThreshold

`func (o *IscsiGlobalUpdate0) HasPoolAvailThreshold() bool`

HasPoolAvailThreshold returns a boolean if a field has been set.

### SetPoolAvailThresholdNil

`func (o *IscsiGlobalUpdate0) SetPoolAvailThresholdNil(b bool)`

 SetPoolAvailThresholdNil sets the value for PoolAvailThreshold to be an explicit nil

### UnsetPoolAvailThreshold
`func (o *IscsiGlobalUpdate0) UnsetPoolAvailThreshold()`

UnsetPoolAvailThreshold ensures that no value is present for PoolAvailThreshold, not even an explicit nil
### GetAlua

`func (o *IscsiGlobalUpdate0) GetAlua() bool`

GetAlua returns the Alua field if non-nil, zero value otherwise.

### GetAluaOk

`func (o *IscsiGlobalUpdate0) GetAluaOk() (*bool, bool)`

GetAluaOk returns a tuple with the Alua field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlua

`func (o *IscsiGlobalUpdate0) SetAlua(v bool)`

SetAlua sets Alua field to given value.

### HasAlua

`func (o *IscsiGlobalUpdate0) HasAlua() bool`

HasAlua returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


