# SmbUpdate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Netbiosname** | Pointer to **string** |  | [optional] 
**NetbiosnameB** | Pointer to **string** |  | [optional] 
**Netbiosalias** | Pointer to **[]string** |  | [optional] 
**Workgroup** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EnableSmb1** | Pointer to **bool** |  | [optional] 
**Unixcharset** | Pointer to **string** |  | [optional] 
**Loglevel** | Pointer to **string** |  | [optional] 
**Syslog** | Pointer to **bool** |  | [optional] 
**AaplExtensions** | Pointer to **bool** |  | [optional] 
**Localmaster** | Pointer to **bool** |  | [optional] 
**Guest** | Pointer to **string** |  | [optional] 
**AdminGroup** | Pointer to **NullableString** |  | [optional] 
**Filemask** | Pointer to **string** |  | [optional] 
**Dirmask** | Pointer to **string** |  | [optional] 
**Ntlmv1Auth** | Pointer to **bool** |  | [optional] 
**Bindip** | Pointer to **[]string** |  | [optional] 
**SmbOptions** | Pointer to **string** |  | [optional] 

## Methods

### NewSmbUpdate0

`func NewSmbUpdate0() *SmbUpdate0`

NewSmbUpdate0 instantiates a new SmbUpdate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmbUpdate0WithDefaults

`func NewSmbUpdate0WithDefaults() *SmbUpdate0`

NewSmbUpdate0WithDefaults instantiates a new SmbUpdate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetbiosname

`func (o *SmbUpdate0) GetNetbiosname() string`

GetNetbiosname returns the Netbiosname field if non-nil, zero value otherwise.

### GetNetbiosnameOk

`func (o *SmbUpdate0) GetNetbiosnameOk() (*string, bool)`

GetNetbiosnameOk returns a tuple with the Netbiosname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbiosname

`func (o *SmbUpdate0) SetNetbiosname(v string)`

SetNetbiosname sets Netbiosname field to given value.

### HasNetbiosname

`func (o *SmbUpdate0) HasNetbiosname() bool`

HasNetbiosname returns a boolean if a field has been set.

### GetNetbiosnameB

`func (o *SmbUpdate0) GetNetbiosnameB() string`

GetNetbiosnameB returns the NetbiosnameB field if non-nil, zero value otherwise.

### GetNetbiosnameBOk

`func (o *SmbUpdate0) GetNetbiosnameBOk() (*string, bool)`

GetNetbiosnameBOk returns a tuple with the NetbiosnameB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbiosnameB

`func (o *SmbUpdate0) SetNetbiosnameB(v string)`

SetNetbiosnameB sets NetbiosnameB field to given value.

### HasNetbiosnameB

`func (o *SmbUpdate0) HasNetbiosnameB() bool`

HasNetbiosnameB returns a boolean if a field has been set.

### GetNetbiosalias

`func (o *SmbUpdate0) GetNetbiosalias() []string`

GetNetbiosalias returns the Netbiosalias field if non-nil, zero value otherwise.

### GetNetbiosaliasOk

`func (o *SmbUpdate0) GetNetbiosaliasOk() (*[]string, bool)`

GetNetbiosaliasOk returns a tuple with the Netbiosalias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbiosalias

`func (o *SmbUpdate0) SetNetbiosalias(v []string)`

SetNetbiosalias sets Netbiosalias field to given value.

### HasNetbiosalias

`func (o *SmbUpdate0) HasNetbiosalias() bool`

HasNetbiosalias returns a boolean if a field has been set.

### GetWorkgroup

`func (o *SmbUpdate0) GetWorkgroup() string`

GetWorkgroup returns the Workgroup field if non-nil, zero value otherwise.

### GetWorkgroupOk

`func (o *SmbUpdate0) GetWorkgroupOk() (*string, bool)`

GetWorkgroupOk returns a tuple with the Workgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkgroup

`func (o *SmbUpdate0) SetWorkgroup(v string)`

SetWorkgroup sets Workgroup field to given value.

### HasWorkgroup

`func (o *SmbUpdate0) HasWorkgroup() bool`

HasWorkgroup returns a boolean if a field has been set.

### GetDescription

`func (o *SmbUpdate0) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SmbUpdate0) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SmbUpdate0) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SmbUpdate0) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnableSmb1

`func (o *SmbUpdate0) GetEnableSmb1() bool`

GetEnableSmb1 returns the EnableSmb1 field if non-nil, zero value otherwise.

### GetEnableSmb1Ok

`func (o *SmbUpdate0) GetEnableSmb1Ok() (*bool, bool)`

GetEnableSmb1Ok returns a tuple with the EnableSmb1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSmb1

`func (o *SmbUpdate0) SetEnableSmb1(v bool)`

SetEnableSmb1 sets EnableSmb1 field to given value.

### HasEnableSmb1

`func (o *SmbUpdate0) HasEnableSmb1() bool`

HasEnableSmb1 returns a boolean if a field has been set.

### GetUnixcharset

`func (o *SmbUpdate0) GetUnixcharset() string`

GetUnixcharset returns the Unixcharset field if non-nil, zero value otherwise.

### GetUnixcharsetOk

`func (o *SmbUpdate0) GetUnixcharsetOk() (*string, bool)`

GetUnixcharsetOk returns a tuple with the Unixcharset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixcharset

`func (o *SmbUpdate0) SetUnixcharset(v string)`

SetUnixcharset sets Unixcharset field to given value.

### HasUnixcharset

`func (o *SmbUpdate0) HasUnixcharset() bool`

HasUnixcharset returns a boolean if a field has been set.

### GetLoglevel

`func (o *SmbUpdate0) GetLoglevel() string`

GetLoglevel returns the Loglevel field if non-nil, zero value otherwise.

### GetLoglevelOk

`func (o *SmbUpdate0) GetLoglevelOk() (*string, bool)`

GetLoglevelOk returns a tuple with the Loglevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoglevel

`func (o *SmbUpdate0) SetLoglevel(v string)`

SetLoglevel sets Loglevel field to given value.

### HasLoglevel

`func (o *SmbUpdate0) HasLoglevel() bool`

HasLoglevel returns a boolean if a field has been set.

### GetSyslog

`func (o *SmbUpdate0) GetSyslog() bool`

GetSyslog returns the Syslog field if non-nil, zero value otherwise.

### GetSyslogOk

`func (o *SmbUpdate0) GetSyslogOk() (*bool, bool)`

GetSyslogOk returns a tuple with the Syslog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslog

`func (o *SmbUpdate0) SetSyslog(v bool)`

SetSyslog sets Syslog field to given value.

### HasSyslog

`func (o *SmbUpdate0) HasSyslog() bool`

HasSyslog returns a boolean if a field has been set.

### GetAaplExtensions

`func (o *SmbUpdate0) GetAaplExtensions() bool`

GetAaplExtensions returns the AaplExtensions field if non-nil, zero value otherwise.

### GetAaplExtensionsOk

`func (o *SmbUpdate0) GetAaplExtensionsOk() (*bool, bool)`

GetAaplExtensionsOk returns a tuple with the AaplExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAaplExtensions

`func (o *SmbUpdate0) SetAaplExtensions(v bool)`

SetAaplExtensions sets AaplExtensions field to given value.

### HasAaplExtensions

`func (o *SmbUpdate0) HasAaplExtensions() bool`

HasAaplExtensions returns a boolean if a field has been set.

### GetLocalmaster

`func (o *SmbUpdate0) GetLocalmaster() bool`

GetLocalmaster returns the Localmaster field if non-nil, zero value otherwise.

### GetLocalmasterOk

`func (o *SmbUpdate0) GetLocalmasterOk() (*bool, bool)`

GetLocalmasterOk returns a tuple with the Localmaster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalmaster

`func (o *SmbUpdate0) SetLocalmaster(v bool)`

SetLocalmaster sets Localmaster field to given value.

### HasLocalmaster

`func (o *SmbUpdate0) HasLocalmaster() bool`

HasLocalmaster returns a boolean if a field has been set.

### GetGuest

`func (o *SmbUpdate0) GetGuest() string`

GetGuest returns the Guest field if non-nil, zero value otherwise.

### GetGuestOk

`func (o *SmbUpdate0) GetGuestOk() (*string, bool)`

GetGuestOk returns a tuple with the Guest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuest

`func (o *SmbUpdate0) SetGuest(v string)`

SetGuest sets Guest field to given value.

### HasGuest

`func (o *SmbUpdate0) HasGuest() bool`

HasGuest returns a boolean if a field has been set.

### GetAdminGroup

`func (o *SmbUpdate0) GetAdminGroup() string`

GetAdminGroup returns the AdminGroup field if non-nil, zero value otherwise.

### GetAdminGroupOk

`func (o *SmbUpdate0) GetAdminGroupOk() (*string, bool)`

GetAdminGroupOk returns a tuple with the AdminGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminGroup

`func (o *SmbUpdate0) SetAdminGroup(v string)`

SetAdminGroup sets AdminGroup field to given value.

### HasAdminGroup

`func (o *SmbUpdate0) HasAdminGroup() bool`

HasAdminGroup returns a boolean if a field has been set.

### SetAdminGroupNil

`func (o *SmbUpdate0) SetAdminGroupNil(b bool)`

 SetAdminGroupNil sets the value for AdminGroup to be an explicit nil

### UnsetAdminGroup
`func (o *SmbUpdate0) UnsetAdminGroup()`

UnsetAdminGroup ensures that no value is present for AdminGroup, not even an explicit nil
### GetFilemask

`func (o *SmbUpdate0) GetFilemask() string`

GetFilemask returns the Filemask field if non-nil, zero value otherwise.

### GetFilemaskOk

`func (o *SmbUpdate0) GetFilemaskOk() (*string, bool)`

GetFilemaskOk returns a tuple with the Filemask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilemask

`func (o *SmbUpdate0) SetFilemask(v string)`

SetFilemask sets Filemask field to given value.

### HasFilemask

`func (o *SmbUpdate0) HasFilemask() bool`

HasFilemask returns a boolean if a field has been set.

### GetDirmask

`func (o *SmbUpdate0) GetDirmask() string`

GetDirmask returns the Dirmask field if non-nil, zero value otherwise.

### GetDirmaskOk

`func (o *SmbUpdate0) GetDirmaskOk() (*string, bool)`

GetDirmaskOk returns a tuple with the Dirmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirmask

`func (o *SmbUpdate0) SetDirmask(v string)`

SetDirmask sets Dirmask field to given value.

### HasDirmask

`func (o *SmbUpdate0) HasDirmask() bool`

HasDirmask returns a boolean if a field has been set.

### GetNtlmv1Auth

`func (o *SmbUpdate0) GetNtlmv1Auth() bool`

GetNtlmv1Auth returns the Ntlmv1Auth field if non-nil, zero value otherwise.

### GetNtlmv1AuthOk

`func (o *SmbUpdate0) GetNtlmv1AuthOk() (*bool, bool)`

GetNtlmv1AuthOk returns a tuple with the Ntlmv1Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtlmv1Auth

`func (o *SmbUpdate0) SetNtlmv1Auth(v bool)`

SetNtlmv1Auth sets Ntlmv1Auth field to given value.

### HasNtlmv1Auth

`func (o *SmbUpdate0) HasNtlmv1Auth() bool`

HasNtlmv1Auth returns a boolean if a field has been set.

### GetBindip

`func (o *SmbUpdate0) GetBindip() []string`

GetBindip returns the Bindip field if non-nil, zero value otherwise.

### GetBindipOk

`func (o *SmbUpdate0) GetBindipOk() (*[]string, bool)`

GetBindipOk returns a tuple with the Bindip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindip

`func (o *SmbUpdate0) SetBindip(v []string)`

SetBindip sets Bindip field to given value.

### HasBindip

`func (o *SmbUpdate0) HasBindip() bool`

HasBindip returns a boolean if a field has been set.

### GetSmbOptions

`func (o *SmbUpdate0) GetSmbOptions() string`

GetSmbOptions returns the SmbOptions field if non-nil, zero value otherwise.

### GetSmbOptionsOk

`func (o *SmbUpdate0) GetSmbOptionsOk() (*string, bool)`

GetSmbOptionsOk returns a tuple with the SmbOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbOptions

`func (o *SmbUpdate0) SetSmbOptions(v string)`

SetSmbOptions sets SmbOptions field to given value.

### HasSmbOptions

`func (o *SmbUpdate0) HasSmbOptions() bool`

HasSmbOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


