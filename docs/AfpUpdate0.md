# AfpUpdate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Guest** | Pointer to **bool** |  | [optional] 
**GuestUser** | Pointer to **string** |  | [optional] 
**Bindip** | Pointer to **[]string** |  | [optional] 
**ConnectionsLimit** | Pointer to **int32** |  | [optional] 
**Dbpath** | Pointer to **string** |  | [optional] 
**GlobalAux** | Pointer to **string** |  | [optional] 
**MapAcls** | Pointer to **string** |  | [optional] 
**ChmodRequest** | Pointer to **string** |  | [optional] 
**Loglevel** | Pointer to **string** |  | [optional] 

## Methods

### NewAfpUpdate0

`func NewAfpUpdate0() *AfpUpdate0`

NewAfpUpdate0 instantiates a new AfpUpdate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAfpUpdate0WithDefaults

`func NewAfpUpdate0WithDefaults() *AfpUpdate0`

NewAfpUpdate0WithDefaults instantiates a new AfpUpdate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuest

`func (o *AfpUpdate0) GetGuest() bool`

GetGuest returns the Guest field if non-nil, zero value otherwise.

### GetGuestOk

`func (o *AfpUpdate0) GetGuestOk() (*bool, bool)`

GetGuestOk returns a tuple with the Guest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuest

`func (o *AfpUpdate0) SetGuest(v bool)`

SetGuest sets Guest field to given value.

### HasGuest

`func (o *AfpUpdate0) HasGuest() bool`

HasGuest returns a boolean if a field has been set.

### GetGuestUser

`func (o *AfpUpdate0) GetGuestUser() string`

GetGuestUser returns the GuestUser field if non-nil, zero value otherwise.

### GetGuestUserOk

`func (o *AfpUpdate0) GetGuestUserOk() (*string, bool)`

GetGuestUserOk returns a tuple with the GuestUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestUser

`func (o *AfpUpdate0) SetGuestUser(v string)`

SetGuestUser sets GuestUser field to given value.

### HasGuestUser

`func (o *AfpUpdate0) HasGuestUser() bool`

HasGuestUser returns a boolean if a field has been set.

### GetBindip

`func (o *AfpUpdate0) GetBindip() []string`

GetBindip returns the Bindip field if non-nil, zero value otherwise.

### GetBindipOk

`func (o *AfpUpdate0) GetBindipOk() (*[]string, bool)`

GetBindipOk returns a tuple with the Bindip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindip

`func (o *AfpUpdate0) SetBindip(v []string)`

SetBindip sets Bindip field to given value.

### HasBindip

`func (o *AfpUpdate0) HasBindip() bool`

HasBindip returns a boolean if a field has been set.

### GetConnectionsLimit

`func (o *AfpUpdate0) GetConnectionsLimit() int32`

GetConnectionsLimit returns the ConnectionsLimit field if non-nil, zero value otherwise.

### GetConnectionsLimitOk

`func (o *AfpUpdate0) GetConnectionsLimitOk() (*int32, bool)`

GetConnectionsLimitOk returns a tuple with the ConnectionsLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionsLimit

`func (o *AfpUpdate0) SetConnectionsLimit(v int32)`

SetConnectionsLimit sets ConnectionsLimit field to given value.

### HasConnectionsLimit

`func (o *AfpUpdate0) HasConnectionsLimit() bool`

HasConnectionsLimit returns a boolean if a field has been set.

### GetDbpath

`func (o *AfpUpdate0) GetDbpath() string`

GetDbpath returns the Dbpath field if non-nil, zero value otherwise.

### GetDbpathOk

`func (o *AfpUpdate0) GetDbpathOk() (*string, bool)`

GetDbpathOk returns a tuple with the Dbpath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbpath

`func (o *AfpUpdate0) SetDbpath(v string)`

SetDbpath sets Dbpath field to given value.

### HasDbpath

`func (o *AfpUpdate0) HasDbpath() bool`

HasDbpath returns a boolean if a field has been set.

### GetGlobalAux

`func (o *AfpUpdate0) GetGlobalAux() string`

GetGlobalAux returns the GlobalAux field if non-nil, zero value otherwise.

### GetGlobalAuxOk

`func (o *AfpUpdate0) GetGlobalAuxOk() (*string, bool)`

GetGlobalAuxOk returns a tuple with the GlobalAux field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalAux

`func (o *AfpUpdate0) SetGlobalAux(v string)`

SetGlobalAux sets GlobalAux field to given value.

### HasGlobalAux

`func (o *AfpUpdate0) HasGlobalAux() bool`

HasGlobalAux returns a boolean if a field has been set.

### GetMapAcls

`func (o *AfpUpdate0) GetMapAcls() string`

GetMapAcls returns the MapAcls field if non-nil, zero value otherwise.

### GetMapAclsOk

`func (o *AfpUpdate0) GetMapAclsOk() (*string, bool)`

GetMapAclsOk returns a tuple with the MapAcls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapAcls

`func (o *AfpUpdate0) SetMapAcls(v string)`

SetMapAcls sets MapAcls field to given value.

### HasMapAcls

`func (o *AfpUpdate0) HasMapAcls() bool`

HasMapAcls returns a boolean if a field has been set.

### GetChmodRequest

`func (o *AfpUpdate0) GetChmodRequest() string`

GetChmodRequest returns the ChmodRequest field if non-nil, zero value otherwise.

### GetChmodRequestOk

`func (o *AfpUpdate0) GetChmodRequestOk() (*string, bool)`

GetChmodRequestOk returns a tuple with the ChmodRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChmodRequest

`func (o *AfpUpdate0) SetChmodRequest(v string)`

SetChmodRequest sets ChmodRequest field to given value.

### HasChmodRequest

`func (o *AfpUpdate0) HasChmodRequest() bool`

HasChmodRequest returns a boolean if a field has been set.

### GetLoglevel

`func (o *AfpUpdate0) GetLoglevel() string`

GetLoglevel returns the Loglevel field if non-nil, zero value otherwise.

### GetLoglevelOk

`func (o *AfpUpdate0) GetLoglevelOk() (*string, bool)`

GetLoglevelOk returns a tuple with the Loglevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoglevel

`func (o *AfpUpdate0) SetLoglevel(v string)`

SetLoglevel sets Loglevel field to given value.

### HasLoglevel

`func (o *AfpUpdate0) HasLoglevel() bool`

HasLoglevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


