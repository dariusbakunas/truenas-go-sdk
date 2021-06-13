# NfsUpdate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Servers** | Pointer to **int32** |  | [optional] 
**Udp** | Pointer to **bool** |  | [optional] 
**AllowNonroot** | Pointer to **bool** |  | [optional] 
**V4** | Pointer to **bool** |  | [optional] 
**V4V3owner** | Pointer to **bool** |  | [optional] 
**V4Krb** | Pointer to **bool** |  | [optional] 
**V4Domain** | Pointer to **string** |  | [optional] 
**Bindip** | Pointer to **[]string** |  | [optional] 
**MountdPort** | Pointer to **NullableInt32** |  | [optional] 
**RpcstatdPort** | Pointer to **NullableInt32** |  | [optional] 
**RpclockdPort** | Pointer to **NullableInt32** |  | [optional] 
**UserdManageGids** | Pointer to **bool** |  | [optional] 
**MountdLog** | Pointer to **bool** |  | [optional] 
**StatdLockdLog** | Pointer to **bool** |  | [optional] 

## Methods

### NewNfsUpdate0

`func NewNfsUpdate0() *NfsUpdate0`

NewNfsUpdate0 instantiates a new NfsUpdate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNfsUpdate0WithDefaults

`func NewNfsUpdate0WithDefaults() *NfsUpdate0`

NewNfsUpdate0WithDefaults instantiates a new NfsUpdate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServers

`func (o *NfsUpdate0) GetServers() int32`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *NfsUpdate0) GetServersOk() (*int32, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *NfsUpdate0) SetServers(v int32)`

SetServers sets Servers field to given value.

### HasServers

`func (o *NfsUpdate0) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetUdp

`func (o *NfsUpdate0) GetUdp() bool`

GetUdp returns the Udp field if non-nil, zero value otherwise.

### GetUdpOk

`func (o *NfsUpdate0) GetUdpOk() (*bool, bool)`

GetUdpOk returns a tuple with the Udp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdp

`func (o *NfsUpdate0) SetUdp(v bool)`

SetUdp sets Udp field to given value.

### HasUdp

`func (o *NfsUpdate0) HasUdp() bool`

HasUdp returns a boolean if a field has been set.

### GetAllowNonroot

`func (o *NfsUpdate0) GetAllowNonroot() bool`

GetAllowNonroot returns the AllowNonroot field if non-nil, zero value otherwise.

### GetAllowNonrootOk

`func (o *NfsUpdate0) GetAllowNonrootOk() (*bool, bool)`

GetAllowNonrootOk returns a tuple with the AllowNonroot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowNonroot

`func (o *NfsUpdate0) SetAllowNonroot(v bool)`

SetAllowNonroot sets AllowNonroot field to given value.

### HasAllowNonroot

`func (o *NfsUpdate0) HasAllowNonroot() bool`

HasAllowNonroot returns a boolean if a field has been set.

### GetV4

`func (o *NfsUpdate0) GetV4() bool`

GetV4 returns the V4 field if non-nil, zero value otherwise.

### GetV4Ok

`func (o *NfsUpdate0) GetV4Ok() (*bool, bool)`

GetV4Ok returns a tuple with the V4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV4

`func (o *NfsUpdate0) SetV4(v bool)`

SetV4 sets V4 field to given value.

### HasV4

`func (o *NfsUpdate0) HasV4() bool`

HasV4 returns a boolean if a field has been set.

### GetV4V3owner

`func (o *NfsUpdate0) GetV4V3owner() bool`

GetV4V3owner returns the V4V3owner field if non-nil, zero value otherwise.

### GetV4V3ownerOk

`func (o *NfsUpdate0) GetV4V3ownerOk() (*bool, bool)`

GetV4V3ownerOk returns a tuple with the V4V3owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV4V3owner

`func (o *NfsUpdate0) SetV4V3owner(v bool)`

SetV4V3owner sets V4V3owner field to given value.

### HasV4V3owner

`func (o *NfsUpdate0) HasV4V3owner() bool`

HasV4V3owner returns a boolean if a field has been set.

### GetV4Krb

`func (o *NfsUpdate0) GetV4Krb() bool`

GetV4Krb returns the V4Krb field if non-nil, zero value otherwise.

### GetV4KrbOk

`func (o *NfsUpdate0) GetV4KrbOk() (*bool, bool)`

GetV4KrbOk returns a tuple with the V4Krb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV4Krb

`func (o *NfsUpdate0) SetV4Krb(v bool)`

SetV4Krb sets V4Krb field to given value.

### HasV4Krb

`func (o *NfsUpdate0) HasV4Krb() bool`

HasV4Krb returns a boolean if a field has been set.

### GetV4Domain

`func (o *NfsUpdate0) GetV4Domain() string`

GetV4Domain returns the V4Domain field if non-nil, zero value otherwise.

### GetV4DomainOk

`func (o *NfsUpdate0) GetV4DomainOk() (*string, bool)`

GetV4DomainOk returns a tuple with the V4Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV4Domain

`func (o *NfsUpdate0) SetV4Domain(v string)`

SetV4Domain sets V4Domain field to given value.

### HasV4Domain

`func (o *NfsUpdate0) HasV4Domain() bool`

HasV4Domain returns a boolean if a field has been set.

### GetBindip

`func (o *NfsUpdate0) GetBindip() []string`

GetBindip returns the Bindip field if non-nil, zero value otherwise.

### GetBindipOk

`func (o *NfsUpdate0) GetBindipOk() (*[]string, bool)`

GetBindipOk returns a tuple with the Bindip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindip

`func (o *NfsUpdate0) SetBindip(v []string)`

SetBindip sets Bindip field to given value.

### HasBindip

`func (o *NfsUpdate0) HasBindip() bool`

HasBindip returns a boolean if a field has been set.

### GetMountdPort

`func (o *NfsUpdate0) GetMountdPort() int32`

GetMountdPort returns the MountdPort field if non-nil, zero value otherwise.

### GetMountdPortOk

`func (o *NfsUpdate0) GetMountdPortOk() (*int32, bool)`

GetMountdPortOk returns a tuple with the MountdPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountdPort

`func (o *NfsUpdate0) SetMountdPort(v int32)`

SetMountdPort sets MountdPort field to given value.

### HasMountdPort

`func (o *NfsUpdate0) HasMountdPort() bool`

HasMountdPort returns a boolean if a field has been set.

### SetMountdPortNil

`func (o *NfsUpdate0) SetMountdPortNil(b bool)`

 SetMountdPortNil sets the value for MountdPort to be an explicit nil

### UnsetMountdPort
`func (o *NfsUpdate0) UnsetMountdPort()`

UnsetMountdPort ensures that no value is present for MountdPort, not even an explicit nil
### GetRpcstatdPort

`func (o *NfsUpdate0) GetRpcstatdPort() int32`

GetRpcstatdPort returns the RpcstatdPort field if non-nil, zero value otherwise.

### GetRpcstatdPortOk

`func (o *NfsUpdate0) GetRpcstatdPortOk() (*int32, bool)`

GetRpcstatdPortOk returns a tuple with the RpcstatdPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcstatdPort

`func (o *NfsUpdate0) SetRpcstatdPort(v int32)`

SetRpcstatdPort sets RpcstatdPort field to given value.

### HasRpcstatdPort

`func (o *NfsUpdate0) HasRpcstatdPort() bool`

HasRpcstatdPort returns a boolean if a field has been set.

### SetRpcstatdPortNil

`func (o *NfsUpdate0) SetRpcstatdPortNil(b bool)`

 SetRpcstatdPortNil sets the value for RpcstatdPort to be an explicit nil

### UnsetRpcstatdPort
`func (o *NfsUpdate0) UnsetRpcstatdPort()`

UnsetRpcstatdPort ensures that no value is present for RpcstatdPort, not even an explicit nil
### GetRpclockdPort

`func (o *NfsUpdate0) GetRpclockdPort() int32`

GetRpclockdPort returns the RpclockdPort field if non-nil, zero value otherwise.

### GetRpclockdPortOk

`func (o *NfsUpdate0) GetRpclockdPortOk() (*int32, bool)`

GetRpclockdPortOk returns a tuple with the RpclockdPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpclockdPort

`func (o *NfsUpdate0) SetRpclockdPort(v int32)`

SetRpclockdPort sets RpclockdPort field to given value.

### HasRpclockdPort

`func (o *NfsUpdate0) HasRpclockdPort() bool`

HasRpclockdPort returns a boolean if a field has been set.

### SetRpclockdPortNil

`func (o *NfsUpdate0) SetRpclockdPortNil(b bool)`

 SetRpclockdPortNil sets the value for RpclockdPort to be an explicit nil

### UnsetRpclockdPort
`func (o *NfsUpdate0) UnsetRpclockdPort()`

UnsetRpclockdPort ensures that no value is present for RpclockdPort, not even an explicit nil
### GetUserdManageGids

`func (o *NfsUpdate0) GetUserdManageGids() bool`

GetUserdManageGids returns the UserdManageGids field if non-nil, zero value otherwise.

### GetUserdManageGidsOk

`func (o *NfsUpdate0) GetUserdManageGidsOk() (*bool, bool)`

GetUserdManageGidsOk returns a tuple with the UserdManageGids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserdManageGids

`func (o *NfsUpdate0) SetUserdManageGids(v bool)`

SetUserdManageGids sets UserdManageGids field to given value.

### HasUserdManageGids

`func (o *NfsUpdate0) HasUserdManageGids() bool`

HasUserdManageGids returns a boolean if a field has been set.

### GetMountdLog

`func (o *NfsUpdate0) GetMountdLog() bool`

GetMountdLog returns the MountdLog field if non-nil, zero value otherwise.

### GetMountdLogOk

`func (o *NfsUpdate0) GetMountdLogOk() (*bool, bool)`

GetMountdLogOk returns a tuple with the MountdLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountdLog

`func (o *NfsUpdate0) SetMountdLog(v bool)`

SetMountdLog sets MountdLog field to given value.

### HasMountdLog

`func (o *NfsUpdate0) HasMountdLog() bool`

HasMountdLog returns a boolean if a field has been set.

### GetStatdLockdLog

`func (o *NfsUpdate0) GetStatdLockdLog() bool`

GetStatdLockdLog returns the StatdLockdLog field if non-nil, zero value otherwise.

### GetStatdLockdLogOk

`func (o *NfsUpdate0) GetStatdLockdLogOk() (*bool, bool)`

GetStatdLockdLogOk returns a tuple with the StatdLockdLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatdLockdLog

`func (o *NfsUpdate0) SetStatdLockdLog(v bool)`

SetStatdLockdLog sets StatdLockdLog field to given value.

### HasStatdLockdLog

`func (o *NfsUpdate0) HasStatdLockdLog() bool`

HasStatdLockdLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


