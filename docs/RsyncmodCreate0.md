# RsyncmodCreate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**Maxconn** | Pointer to **int32** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**Hostsallow** | Pointer to **[]string** |  | [optional] 
**Hostsdeny** | Pointer to **[]string** |  | [optional] 
**Auxiliary** | Pointer to **string** |  | [optional] 

## Methods

### NewRsyncmodCreate0

`func NewRsyncmodCreate0() *RsyncmodCreate0`

NewRsyncmodCreate0 instantiates a new RsyncmodCreate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRsyncmodCreate0WithDefaults

`func NewRsyncmodCreate0WithDefaults() *RsyncmodCreate0`

NewRsyncmodCreate0WithDefaults instantiates a new RsyncmodCreate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *RsyncmodCreate0) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RsyncmodCreate0) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RsyncmodCreate0) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RsyncmodCreate0) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *RsyncmodCreate0) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RsyncmodCreate0) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RsyncmodCreate0) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RsyncmodCreate0) HasName() bool`

HasName returns a boolean if a field has been set.

### GetComment

`func (o *RsyncmodCreate0) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RsyncmodCreate0) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RsyncmodCreate0) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RsyncmodCreate0) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetPath

`func (o *RsyncmodCreate0) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RsyncmodCreate0) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RsyncmodCreate0) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *RsyncmodCreate0) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetMode

`func (o *RsyncmodCreate0) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *RsyncmodCreate0) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *RsyncmodCreate0) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *RsyncmodCreate0) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetMaxconn

`func (o *RsyncmodCreate0) GetMaxconn() int32`

GetMaxconn returns the Maxconn field if non-nil, zero value otherwise.

### GetMaxconnOk

`func (o *RsyncmodCreate0) GetMaxconnOk() (*int32, bool)`

GetMaxconnOk returns a tuple with the Maxconn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxconn

`func (o *RsyncmodCreate0) SetMaxconn(v int32)`

SetMaxconn sets Maxconn field to given value.

### HasMaxconn

`func (o *RsyncmodCreate0) HasMaxconn() bool`

HasMaxconn returns a boolean if a field has been set.

### GetUser

`func (o *RsyncmodCreate0) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RsyncmodCreate0) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RsyncmodCreate0) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *RsyncmodCreate0) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetGroup

`func (o *RsyncmodCreate0) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *RsyncmodCreate0) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *RsyncmodCreate0) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *RsyncmodCreate0) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetHostsallow

`func (o *RsyncmodCreate0) GetHostsallow() []string`

GetHostsallow returns the Hostsallow field if non-nil, zero value otherwise.

### GetHostsallowOk

`func (o *RsyncmodCreate0) GetHostsallowOk() (*[]string, bool)`

GetHostsallowOk returns a tuple with the Hostsallow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsallow

`func (o *RsyncmodCreate0) SetHostsallow(v []string)`

SetHostsallow sets Hostsallow field to given value.

### HasHostsallow

`func (o *RsyncmodCreate0) HasHostsallow() bool`

HasHostsallow returns a boolean if a field has been set.

### GetHostsdeny

`func (o *RsyncmodCreate0) GetHostsdeny() []string`

GetHostsdeny returns the Hostsdeny field if non-nil, zero value otherwise.

### GetHostsdenyOk

`func (o *RsyncmodCreate0) GetHostsdenyOk() (*[]string, bool)`

GetHostsdenyOk returns a tuple with the Hostsdeny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsdeny

`func (o *RsyncmodCreate0) SetHostsdeny(v []string)`

SetHostsdeny sets Hostsdeny field to given value.

### HasHostsdeny

`func (o *RsyncmodCreate0) HasHostsdeny() bool`

HasHostsdeny returns a boolean if a field has been set.

### GetAuxiliary

`func (o *RsyncmodCreate0) GetAuxiliary() string`

GetAuxiliary returns the Auxiliary field if non-nil, zero value otherwise.

### GetAuxiliaryOk

`func (o *RsyncmodCreate0) GetAuxiliaryOk() (*string, bool)`

GetAuxiliaryOk returns a tuple with the Auxiliary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxiliary

`func (o *RsyncmodCreate0) SetAuxiliary(v string)`

SetAuxiliary sets Auxiliary field to given value.

### HasAuxiliary

`func (o *RsyncmodCreate0) HasAuxiliary() bool`

HasAuxiliary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


