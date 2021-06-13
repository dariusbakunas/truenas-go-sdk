# SharingNfsUpdate1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paths** | Pointer to **[]string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Networks** | Pointer to **[]string** |  | [optional] 
**Hosts** | Pointer to **[]string** |  | [optional] 
**Alldirs** | Pointer to **bool** |  | [optional] 
**Ro** | Pointer to **bool** |  | [optional] 
**Quiet** | Pointer to **bool** |  | [optional] 
**MaprootUser** | Pointer to **NullableString** |  | [optional] 
**MaprootGroup** | Pointer to **NullableString** |  | [optional] 
**MapallUser** | Pointer to **NullableString** |  | [optional] 
**MapallGroup** | Pointer to **NullableString** |  | [optional] 
**Security** | Pointer to **[]string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewSharingNfsUpdate1

`func NewSharingNfsUpdate1() *SharingNfsUpdate1`

NewSharingNfsUpdate1 instantiates a new SharingNfsUpdate1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharingNfsUpdate1WithDefaults

`func NewSharingNfsUpdate1WithDefaults() *SharingNfsUpdate1`

NewSharingNfsUpdate1WithDefaults instantiates a new SharingNfsUpdate1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaths

`func (o *SharingNfsUpdate1) GetPaths() []string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *SharingNfsUpdate1) GetPathsOk() (*[]string, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *SharingNfsUpdate1) SetPaths(v []string)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *SharingNfsUpdate1) HasPaths() bool`

HasPaths returns a boolean if a field has been set.

### GetComment

`func (o *SharingNfsUpdate1) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SharingNfsUpdate1) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SharingNfsUpdate1) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SharingNfsUpdate1) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetNetworks

`func (o *SharingNfsUpdate1) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *SharingNfsUpdate1) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *SharingNfsUpdate1) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *SharingNfsUpdate1) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetHosts

`func (o *SharingNfsUpdate1) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *SharingNfsUpdate1) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *SharingNfsUpdate1) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *SharingNfsUpdate1) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetAlldirs

`func (o *SharingNfsUpdate1) GetAlldirs() bool`

GetAlldirs returns the Alldirs field if non-nil, zero value otherwise.

### GetAlldirsOk

`func (o *SharingNfsUpdate1) GetAlldirsOk() (*bool, bool)`

GetAlldirsOk returns a tuple with the Alldirs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlldirs

`func (o *SharingNfsUpdate1) SetAlldirs(v bool)`

SetAlldirs sets Alldirs field to given value.

### HasAlldirs

`func (o *SharingNfsUpdate1) HasAlldirs() bool`

HasAlldirs returns a boolean if a field has been set.

### GetRo

`func (o *SharingNfsUpdate1) GetRo() bool`

GetRo returns the Ro field if non-nil, zero value otherwise.

### GetRoOk

`func (o *SharingNfsUpdate1) GetRoOk() (*bool, bool)`

GetRoOk returns a tuple with the Ro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRo

`func (o *SharingNfsUpdate1) SetRo(v bool)`

SetRo sets Ro field to given value.

### HasRo

`func (o *SharingNfsUpdate1) HasRo() bool`

HasRo returns a boolean if a field has been set.

### GetQuiet

`func (o *SharingNfsUpdate1) GetQuiet() bool`

GetQuiet returns the Quiet field if non-nil, zero value otherwise.

### GetQuietOk

`func (o *SharingNfsUpdate1) GetQuietOk() (*bool, bool)`

GetQuietOk returns a tuple with the Quiet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuiet

`func (o *SharingNfsUpdate1) SetQuiet(v bool)`

SetQuiet sets Quiet field to given value.

### HasQuiet

`func (o *SharingNfsUpdate1) HasQuiet() bool`

HasQuiet returns a boolean if a field has been set.

### GetMaprootUser

`func (o *SharingNfsUpdate1) GetMaprootUser() string`

GetMaprootUser returns the MaprootUser field if non-nil, zero value otherwise.

### GetMaprootUserOk

`func (o *SharingNfsUpdate1) GetMaprootUserOk() (*string, bool)`

GetMaprootUserOk returns a tuple with the MaprootUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaprootUser

`func (o *SharingNfsUpdate1) SetMaprootUser(v string)`

SetMaprootUser sets MaprootUser field to given value.

### HasMaprootUser

`func (o *SharingNfsUpdate1) HasMaprootUser() bool`

HasMaprootUser returns a boolean if a field has been set.

### SetMaprootUserNil

`func (o *SharingNfsUpdate1) SetMaprootUserNil(b bool)`

 SetMaprootUserNil sets the value for MaprootUser to be an explicit nil

### UnsetMaprootUser
`func (o *SharingNfsUpdate1) UnsetMaprootUser()`

UnsetMaprootUser ensures that no value is present for MaprootUser, not even an explicit nil
### GetMaprootGroup

`func (o *SharingNfsUpdate1) GetMaprootGroup() string`

GetMaprootGroup returns the MaprootGroup field if non-nil, zero value otherwise.

### GetMaprootGroupOk

`func (o *SharingNfsUpdate1) GetMaprootGroupOk() (*string, bool)`

GetMaprootGroupOk returns a tuple with the MaprootGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaprootGroup

`func (o *SharingNfsUpdate1) SetMaprootGroup(v string)`

SetMaprootGroup sets MaprootGroup field to given value.

### HasMaprootGroup

`func (o *SharingNfsUpdate1) HasMaprootGroup() bool`

HasMaprootGroup returns a boolean if a field has been set.

### SetMaprootGroupNil

`func (o *SharingNfsUpdate1) SetMaprootGroupNil(b bool)`

 SetMaprootGroupNil sets the value for MaprootGroup to be an explicit nil

### UnsetMaprootGroup
`func (o *SharingNfsUpdate1) UnsetMaprootGroup()`

UnsetMaprootGroup ensures that no value is present for MaprootGroup, not even an explicit nil
### GetMapallUser

`func (o *SharingNfsUpdate1) GetMapallUser() string`

GetMapallUser returns the MapallUser field if non-nil, zero value otherwise.

### GetMapallUserOk

`func (o *SharingNfsUpdate1) GetMapallUserOk() (*string, bool)`

GetMapallUserOk returns a tuple with the MapallUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapallUser

`func (o *SharingNfsUpdate1) SetMapallUser(v string)`

SetMapallUser sets MapallUser field to given value.

### HasMapallUser

`func (o *SharingNfsUpdate1) HasMapallUser() bool`

HasMapallUser returns a boolean if a field has been set.

### SetMapallUserNil

`func (o *SharingNfsUpdate1) SetMapallUserNil(b bool)`

 SetMapallUserNil sets the value for MapallUser to be an explicit nil

### UnsetMapallUser
`func (o *SharingNfsUpdate1) UnsetMapallUser()`

UnsetMapallUser ensures that no value is present for MapallUser, not even an explicit nil
### GetMapallGroup

`func (o *SharingNfsUpdate1) GetMapallGroup() string`

GetMapallGroup returns the MapallGroup field if non-nil, zero value otherwise.

### GetMapallGroupOk

`func (o *SharingNfsUpdate1) GetMapallGroupOk() (*string, bool)`

GetMapallGroupOk returns a tuple with the MapallGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapallGroup

`func (o *SharingNfsUpdate1) SetMapallGroup(v string)`

SetMapallGroup sets MapallGroup field to given value.

### HasMapallGroup

`func (o *SharingNfsUpdate1) HasMapallGroup() bool`

HasMapallGroup returns a boolean if a field has been set.

### SetMapallGroupNil

`func (o *SharingNfsUpdate1) SetMapallGroupNil(b bool)`

 SetMapallGroupNil sets the value for MapallGroup to be an explicit nil

### UnsetMapallGroup
`func (o *SharingNfsUpdate1) UnsetMapallGroup()`

UnsetMapallGroup ensures that no value is present for MapallGroup, not even an explicit nil
### GetSecurity

`func (o *SharingNfsUpdate1) GetSecurity() []string`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *SharingNfsUpdate1) GetSecurityOk() (*[]string, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *SharingNfsUpdate1) SetSecurity(v []string)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *SharingNfsUpdate1) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetEnabled

`func (o *SharingNfsUpdate1) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SharingNfsUpdate1) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SharingNfsUpdate1) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SharingNfsUpdate1) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


