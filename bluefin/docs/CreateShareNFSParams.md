# CreateShareNFSParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paths** | **[]string** |  | 
**Comment** | Pointer to **string** |  | [optional] 
**Networks** | Pointer to **[]string** |  | [optional] 
**Hosts** | Pointer to **[]string** |  | [optional] 
**Alldirs** | Pointer to **bool** |  | [optional] 
**Ro** | Pointer to **bool** |  | [optional] 
**Quiet** | Pointer to **bool** |  | [optional] 
**MaprootUser** | Pointer to **string** |  | [optional] 
**MaprootGroup** | Pointer to **string** |  | [optional] 
**MapallUser** | Pointer to **string** |  | [optional] 
**MapallGroup** | Pointer to **string** |  | [optional] 
**Security** | Pointer to **[]string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateShareNFSParams

`func NewCreateShareNFSParams(paths []string, ) *CreateShareNFSParams`

NewCreateShareNFSParams instantiates a new CreateShareNFSParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateShareNFSParamsWithDefaults

`func NewCreateShareNFSParamsWithDefaults() *CreateShareNFSParams`

NewCreateShareNFSParamsWithDefaults instantiates a new CreateShareNFSParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaths

`func (o *CreateShareNFSParams) GetPaths() []string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *CreateShareNFSParams) GetPathsOk() (*[]string, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *CreateShareNFSParams) SetPaths(v []string)`

SetPaths sets Paths field to given value.


### GetComment

`func (o *CreateShareNFSParams) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateShareNFSParams) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateShareNFSParams) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateShareNFSParams) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetNetworks

`func (o *CreateShareNFSParams) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *CreateShareNFSParams) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *CreateShareNFSParams) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *CreateShareNFSParams) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetHosts

`func (o *CreateShareNFSParams) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *CreateShareNFSParams) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *CreateShareNFSParams) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *CreateShareNFSParams) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetAlldirs

`func (o *CreateShareNFSParams) GetAlldirs() bool`

GetAlldirs returns the Alldirs field if non-nil, zero value otherwise.

### GetAlldirsOk

`func (o *CreateShareNFSParams) GetAlldirsOk() (*bool, bool)`

GetAlldirsOk returns a tuple with the Alldirs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlldirs

`func (o *CreateShareNFSParams) SetAlldirs(v bool)`

SetAlldirs sets Alldirs field to given value.

### HasAlldirs

`func (o *CreateShareNFSParams) HasAlldirs() bool`

HasAlldirs returns a boolean if a field has been set.

### GetRo

`func (o *CreateShareNFSParams) GetRo() bool`

GetRo returns the Ro field if non-nil, zero value otherwise.

### GetRoOk

`func (o *CreateShareNFSParams) GetRoOk() (*bool, bool)`

GetRoOk returns a tuple with the Ro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRo

`func (o *CreateShareNFSParams) SetRo(v bool)`

SetRo sets Ro field to given value.

### HasRo

`func (o *CreateShareNFSParams) HasRo() bool`

HasRo returns a boolean if a field has been set.

### GetQuiet

`func (o *CreateShareNFSParams) GetQuiet() bool`

GetQuiet returns the Quiet field if non-nil, zero value otherwise.

### GetQuietOk

`func (o *CreateShareNFSParams) GetQuietOk() (*bool, bool)`

GetQuietOk returns a tuple with the Quiet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuiet

`func (o *CreateShareNFSParams) SetQuiet(v bool)`

SetQuiet sets Quiet field to given value.

### HasQuiet

`func (o *CreateShareNFSParams) HasQuiet() bool`

HasQuiet returns a boolean if a field has been set.

### GetMaprootUser

`func (o *CreateShareNFSParams) GetMaprootUser() string`

GetMaprootUser returns the MaprootUser field if non-nil, zero value otherwise.

### GetMaprootUserOk

`func (o *CreateShareNFSParams) GetMaprootUserOk() (*string, bool)`

GetMaprootUserOk returns a tuple with the MaprootUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaprootUser

`func (o *CreateShareNFSParams) SetMaprootUser(v string)`

SetMaprootUser sets MaprootUser field to given value.

### HasMaprootUser

`func (o *CreateShareNFSParams) HasMaprootUser() bool`

HasMaprootUser returns a boolean if a field has been set.

### GetMaprootGroup

`func (o *CreateShareNFSParams) GetMaprootGroup() string`

GetMaprootGroup returns the MaprootGroup field if non-nil, zero value otherwise.

### GetMaprootGroupOk

`func (o *CreateShareNFSParams) GetMaprootGroupOk() (*string, bool)`

GetMaprootGroupOk returns a tuple with the MaprootGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaprootGroup

`func (o *CreateShareNFSParams) SetMaprootGroup(v string)`

SetMaprootGroup sets MaprootGroup field to given value.

### HasMaprootGroup

`func (o *CreateShareNFSParams) HasMaprootGroup() bool`

HasMaprootGroup returns a boolean if a field has been set.

### GetMapallUser

`func (o *CreateShareNFSParams) GetMapallUser() string`

GetMapallUser returns the MapallUser field if non-nil, zero value otherwise.

### GetMapallUserOk

`func (o *CreateShareNFSParams) GetMapallUserOk() (*string, bool)`

GetMapallUserOk returns a tuple with the MapallUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapallUser

`func (o *CreateShareNFSParams) SetMapallUser(v string)`

SetMapallUser sets MapallUser field to given value.

### HasMapallUser

`func (o *CreateShareNFSParams) HasMapallUser() bool`

HasMapallUser returns a boolean if a field has been set.

### GetMapallGroup

`func (o *CreateShareNFSParams) GetMapallGroup() string`

GetMapallGroup returns the MapallGroup field if non-nil, zero value otherwise.

### GetMapallGroupOk

`func (o *CreateShareNFSParams) GetMapallGroupOk() (*string, bool)`

GetMapallGroupOk returns a tuple with the MapallGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapallGroup

`func (o *CreateShareNFSParams) SetMapallGroup(v string)`

SetMapallGroup sets MapallGroup field to given value.

### HasMapallGroup

`func (o *CreateShareNFSParams) HasMapallGroup() bool`

HasMapallGroup returns a boolean if a field has been set.

### GetSecurity

`func (o *CreateShareNFSParams) GetSecurity() []string`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *CreateShareNFSParams) GetSecurityOk() (*[]string, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *CreateShareNFSParams) SetSecurity(v []string)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *CreateShareNFSParams) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetEnabled

`func (o *CreateShareNFSParams) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateShareNFSParams) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateShareNFSParams) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateShareNFSParams) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


