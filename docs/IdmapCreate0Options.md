# IdmapCreate0Options

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaMode** | Pointer to **string** |  | [optional] 
**UnixPrimaryGroup** | Pointer to **bool** |  | [optional] 
**UnixNssInfo** | Pointer to **bool** |  | [optional] 
**Rangesize** | Pointer to **int32** |  | [optional] 
**Readonly** | Pointer to **bool** |  | [optional] 
**IgnoreBuiltin** | Pointer to **bool** |  | [optional] 
**LdapBaseDn** | Pointer to **string** |  | [optional] 
**LdapUserDn** | Pointer to **string** |  | [optional] 
**LdapUserDnPassword** | Pointer to **string** |  | [optional] 
**LdapUrl** | Pointer to **string** |  | [optional] 
**Ssl** | Pointer to **string** |  | [optional] 
**LinkedService** | Pointer to **string** |  | [optional] 
**LdapServer** | Pointer to **string** |  | [optional] 
**LdapRealm** | Pointer to **bool** |  | [optional] 
**BindPathUser** | Pointer to **string** |  | [optional] 
**BindPathGroup** | Pointer to **string** |  | [optional] 
**UserCn** | Pointer to **bool** |  | [optional] 
**CnRealm** | Pointer to **string** |  | [optional] 
**LdapDomain** | Pointer to **string** |  | [optional] 
**SssdCompat** | Pointer to **bool** |  | [optional] 

## Methods

### NewIdmapCreate0Options

`func NewIdmapCreate0Options() *IdmapCreate0Options`

NewIdmapCreate0Options instantiates a new IdmapCreate0Options object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdmapCreate0OptionsWithDefaults

`func NewIdmapCreate0OptionsWithDefaults() *IdmapCreate0Options`

NewIdmapCreate0OptionsWithDefaults instantiates a new IdmapCreate0Options object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaMode

`func (o *IdmapCreate0Options) GetSchemaMode() string`

GetSchemaMode returns the SchemaMode field if non-nil, zero value otherwise.

### GetSchemaModeOk

`func (o *IdmapCreate0Options) GetSchemaModeOk() (*string, bool)`

GetSchemaModeOk returns a tuple with the SchemaMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaMode

`func (o *IdmapCreate0Options) SetSchemaMode(v string)`

SetSchemaMode sets SchemaMode field to given value.

### HasSchemaMode

`func (o *IdmapCreate0Options) HasSchemaMode() bool`

HasSchemaMode returns a boolean if a field has been set.

### GetUnixPrimaryGroup

`func (o *IdmapCreate0Options) GetUnixPrimaryGroup() bool`

GetUnixPrimaryGroup returns the UnixPrimaryGroup field if non-nil, zero value otherwise.

### GetUnixPrimaryGroupOk

`func (o *IdmapCreate0Options) GetUnixPrimaryGroupOk() (*bool, bool)`

GetUnixPrimaryGroupOk returns a tuple with the UnixPrimaryGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixPrimaryGroup

`func (o *IdmapCreate0Options) SetUnixPrimaryGroup(v bool)`

SetUnixPrimaryGroup sets UnixPrimaryGroup field to given value.

### HasUnixPrimaryGroup

`func (o *IdmapCreate0Options) HasUnixPrimaryGroup() bool`

HasUnixPrimaryGroup returns a boolean if a field has been set.

### GetUnixNssInfo

`func (o *IdmapCreate0Options) GetUnixNssInfo() bool`

GetUnixNssInfo returns the UnixNssInfo field if non-nil, zero value otherwise.

### GetUnixNssInfoOk

`func (o *IdmapCreate0Options) GetUnixNssInfoOk() (*bool, bool)`

GetUnixNssInfoOk returns a tuple with the UnixNssInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixNssInfo

`func (o *IdmapCreate0Options) SetUnixNssInfo(v bool)`

SetUnixNssInfo sets UnixNssInfo field to given value.

### HasUnixNssInfo

`func (o *IdmapCreate0Options) HasUnixNssInfo() bool`

HasUnixNssInfo returns a boolean if a field has been set.

### GetRangesize

`func (o *IdmapCreate0Options) GetRangesize() int32`

GetRangesize returns the Rangesize field if non-nil, zero value otherwise.

### GetRangesizeOk

`func (o *IdmapCreate0Options) GetRangesizeOk() (*int32, bool)`

GetRangesizeOk returns a tuple with the Rangesize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangesize

`func (o *IdmapCreate0Options) SetRangesize(v int32)`

SetRangesize sets Rangesize field to given value.

### HasRangesize

`func (o *IdmapCreate0Options) HasRangesize() bool`

HasRangesize returns a boolean if a field has been set.

### GetReadonly

`func (o *IdmapCreate0Options) GetReadonly() bool`

GetReadonly returns the Readonly field if non-nil, zero value otherwise.

### GetReadonlyOk

`func (o *IdmapCreate0Options) GetReadonlyOk() (*bool, bool)`

GetReadonlyOk returns a tuple with the Readonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonly

`func (o *IdmapCreate0Options) SetReadonly(v bool)`

SetReadonly sets Readonly field to given value.

### HasReadonly

`func (o *IdmapCreate0Options) HasReadonly() bool`

HasReadonly returns a boolean if a field has been set.

### GetIgnoreBuiltin

`func (o *IdmapCreate0Options) GetIgnoreBuiltin() bool`

GetIgnoreBuiltin returns the IgnoreBuiltin field if non-nil, zero value otherwise.

### GetIgnoreBuiltinOk

`func (o *IdmapCreate0Options) GetIgnoreBuiltinOk() (*bool, bool)`

GetIgnoreBuiltinOk returns a tuple with the IgnoreBuiltin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreBuiltin

`func (o *IdmapCreate0Options) SetIgnoreBuiltin(v bool)`

SetIgnoreBuiltin sets IgnoreBuiltin field to given value.

### HasIgnoreBuiltin

`func (o *IdmapCreate0Options) HasIgnoreBuiltin() bool`

HasIgnoreBuiltin returns a boolean if a field has been set.

### GetLdapBaseDn

`func (o *IdmapCreate0Options) GetLdapBaseDn() string`

GetLdapBaseDn returns the LdapBaseDn field if non-nil, zero value otherwise.

### GetLdapBaseDnOk

`func (o *IdmapCreate0Options) GetLdapBaseDnOk() (*string, bool)`

GetLdapBaseDnOk returns a tuple with the LdapBaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapBaseDn

`func (o *IdmapCreate0Options) SetLdapBaseDn(v string)`

SetLdapBaseDn sets LdapBaseDn field to given value.

### HasLdapBaseDn

`func (o *IdmapCreate0Options) HasLdapBaseDn() bool`

HasLdapBaseDn returns a boolean if a field has been set.

### GetLdapUserDn

`func (o *IdmapCreate0Options) GetLdapUserDn() string`

GetLdapUserDn returns the LdapUserDn field if non-nil, zero value otherwise.

### GetLdapUserDnOk

`func (o *IdmapCreate0Options) GetLdapUserDnOk() (*string, bool)`

GetLdapUserDnOk returns a tuple with the LdapUserDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUserDn

`func (o *IdmapCreate0Options) SetLdapUserDn(v string)`

SetLdapUserDn sets LdapUserDn field to given value.

### HasLdapUserDn

`func (o *IdmapCreate0Options) HasLdapUserDn() bool`

HasLdapUserDn returns a boolean if a field has been set.

### GetLdapUserDnPassword

`func (o *IdmapCreate0Options) GetLdapUserDnPassword() string`

GetLdapUserDnPassword returns the LdapUserDnPassword field if non-nil, zero value otherwise.

### GetLdapUserDnPasswordOk

`func (o *IdmapCreate0Options) GetLdapUserDnPasswordOk() (*string, bool)`

GetLdapUserDnPasswordOk returns a tuple with the LdapUserDnPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUserDnPassword

`func (o *IdmapCreate0Options) SetLdapUserDnPassword(v string)`

SetLdapUserDnPassword sets LdapUserDnPassword field to given value.

### HasLdapUserDnPassword

`func (o *IdmapCreate0Options) HasLdapUserDnPassword() bool`

HasLdapUserDnPassword returns a boolean if a field has been set.

### GetLdapUrl

`func (o *IdmapCreate0Options) GetLdapUrl() string`

GetLdapUrl returns the LdapUrl field if non-nil, zero value otherwise.

### GetLdapUrlOk

`func (o *IdmapCreate0Options) GetLdapUrlOk() (*string, bool)`

GetLdapUrlOk returns a tuple with the LdapUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUrl

`func (o *IdmapCreate0Options) SetLdapUrl(v string)`

SetLdapUrl sets LdapUrl field to given value.

### HasLdapUrl

`func (o *IdmapCreate0Options) HasLdapUrl() bool`

HasLdapUrl returns a boolean if a field has been set.

### GetSsl

`func (o *IdmapCreate0Options) GetSsl() string`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *IdmapCreate0Options) GetSslOk() (*string, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *IdmapCreate0Options) SetSsl(v string)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *IdmapCreate0Options) HasSsl() bool`

HasSsl returns a boolean if a field has been set.

### GetLinkedService

`func (o *IdmapCreate0Options) GetLinkedService() string`

GetLinkedService returns the LinkedService field if non-nil, zero value otherwise.

### GetLinkedServiceOk

`func (o *IdmapCreate0Options) GetLinkedServiceOk() (*string, bool)`

GetLinkedServiceOk returns a tuple with the LinkedService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedService

`func (o *IdmapCreate0Options) SetLinkedService(v string)`

SetLinkedService sets LinkedService field to given value.

### HasLinkedService

`func (o *IdmapCreate0Options) HasLinkedService() bool`

HasLinkedService returns a boolean if a field has been set.

### GetLdapServer

`func (o *IdmapCreate0Options) GetLdapServer() string`

GetLdapServer returns the LdapServer field if non-nil, zero value otherwise.

### GetLdapServerOk

`func (o *IdmapCreate0Options) GetLdapServerOk() (*string, bool)`

GetLdapServerOk returns a tuple with the LdapServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapServer

`func (o *IdmapCreate0Options) SetLdapServer(v string)`

SetLdapServer sets LdapServer field to given value.

### HasLdapServer

`func (o *IdmapCreate0Options) HasLdapServer() bool`

HasLdapServer returns a boolean if a field has been set.

### GetLdapRealm

`func (o *IdmapCreate0Options) GetLdapRealm() bool`

GetLdapRealm returns the LdapRealm field if non-nil, zero value otherwise.

### GetLdapRealmOk

`func (o *IdmapCreate0Options) GetLdapRealmOk() (*bool, bool)`

GetLdapRealmOk returns a tuple with the LdapRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapRealm

`func (o *IdmapCreate0Options) SetLdapRealm(v bool)`

SetLdapRealm sets LdapRealm field to given value.

### HasLdapRealm

`func (o *IdmapCreate0Options) HasLdapRealm() bool`

HasLdapRealm returns a boolean if a field has been set.

### GetBindPathUser

`func (o *IdmapCreate0Options) GetBindPathUser() string`

GetBindPathUser returns the BindPathUser field if non-nil, zero value otherwise.

### GetBindPathUserOk

`func (o *IdmapCreate0Options) GetBindPathUserOk() (*string, bool)`

GetBindPathUserOk returns a tuple with the BindPathUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindPathUser

`func (o *IdmapCreate0Options) SetBindPathUser(v string)`

SetBindPathUser sets BindPathUser field to given value.

### HasBindPathUser

`func (o *IdmapCreate0Options) HasBindPathUser() bool`

HasBindPathUser returns a boolean if a field has been set.

### GetBindPathGroup

`func (o *IdmapCreate0Options) GetBindPathGroup() string`

GetBindPathGroup returns the BindPathGroup field if non-nil, zero value otherwise.

### GetBindPathGroupOk

`func (o *IdmapCreate0Options) GetBindPathGroupOk() (*string, bool)`

GetBindPathGroupOk returns a tuple with the BindPathGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindPathGroup

`func (o *IdmapCreate0Options) SetBindPathGroup(v string)`

SetBindPathGroup sets BindPathGroup field to given value.

### HasBindPathGroup

`func (o *IdmapCreate0Options) HasBindPathGroup() bool`

HasBindPathGroup returns a boolean if a field has been set.

### GetUserCn

`func (o *IdmapCreate0Options) GetUserCn() bool`

GetUserCn returns the UserCn field if non-nil, zero value otherwise.

### GetUserCnOk

`func (o *IdmapCreate0Options) GetUserCnOk() (*bool, bool)`

GetUserCnOk returns a tuple with the UserCn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCn

`func (o *IdmapCreate0Options) SetUserCn(v bool)`

SetUserCn sets UserCn field to given value.

### HasUserCn

`func (o *IdmapCreate0Options) HasUserCn() bool`

HasUserCn returns a boolean if a field has been set.

### GetCnRealm

`func (o *IdmapCreate0Options) GetCnRealm() string`

GetCnRealm returns the CnRealm field if non-nil, zero value otherwise.

### GetCnRealmOk

`func (o *IdmapCreate0Options) GetCnRealmOk() (*string, bool)`

GetCnRealmOk returns a tuple with the CnRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnRealm

`func (o *IdmapCreate0Options) SetCnRealm(v string)`

SetCnRealm sets CnRealm field to given value.

### HasCnRealm

`func (o *IdmapCreate0Options) HasCnRealm() bool`

HasCnRealm returns a boolean if a field has been set.

### GetLdapDomain

`func (o *IdmapCreate0Options) GetLdapDomain() string`

GetLdapDomain returns the LdapDomain field if non-nil, zero value otherwise.

### GetLdapDomainOk

`func (o *IdmapCreate0Options) GetLdapDomainOk() (*string, bool)`

GetLdapDomainOk returns a tuple with the LdapDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapDomain

`func (o *IdmapCreate0Options) SetLdapDomain(v string)`

SetLdapDomain sets LdapDomain field to given value.

### HasLdapDomain

`func (o *IdmapCreate0Options) HasLdapDomain() bool`

HasLdapDomain returns a boolean if a field has been set.

### GetSssdCompat

`func (o *IdmapCreate0Options) GetSssdCompat() bool`

GetSssdCompat returns the SssdCompat field if non-nil, zero value otherwise.

### GetSssdCompatOk

`func (o *IdmapCreate0Options) GetSssdCompatOk() (*bool, bool)`

GetSssdCompatOk returns a tuple with the SssdCompat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSssdCompat

`func (o *IdmapCreate0Options) SetSssdCompat(v bool)`

SetSssdCompat sets SssdCompat field to given value.

### HasSssdCompat

`func (o *IdmapCreate0Options) HasSssdCompat() bool`

HasSssdCompat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


