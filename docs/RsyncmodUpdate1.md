# RsyncmodUpdate1

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

### NewRsyncmodUpdate1

`func NewRsyncmodUpdate1() *RsyncmodUpdate1`

NewRsyncmodUpdate1 instantiates a new RsyncmodUpdate1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRsyncmodUpdate1WithDefaults

`func NewRsyncmodUpdate1WithDefaults() *RsyncmodUpdate1`

NewRsyncmodUpdate1WithDefaults instantiates a new RsyncmodUpdate1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *RsyncmodUpdate1) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RsyncmodUpdate1) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RsyncmodUpdate1) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RsyncmodUpdate1) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *RsyncmodUpdate1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RsyncmodUpdate1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RsyncmodUpdate1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RsyncmodUpdate1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetComment

`func (o *RsyncmodUpdate1) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RsyncmodUpdate1) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RsyncmodUpdate1) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RsyncmodUpdate1) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetPath

`func (o *RsyncmodUpdate1) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RsyncmodUpdate1) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RsyncmodUpdate1) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *RsyncmodUpdate1) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetMode

`func (o *RsyncmodUpdate1) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *RsyncmodUpdate1) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *RsyncmodUpdate1) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *RsyncmodUpdate1) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetMaxconn

`func (o *RsyncmodUpdate1) GetMaxconn() int32`

GetMaxconn returns the Maxconn field if non-nil, zero value otherwise.

### GetMaxconnOk

`func (o *RsyncmodUpdate1) GetMaxconnOk() (*int32, bool)`

GetMaxconnOk returns a tuple with the Maxconn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxconn

`func (o *RsyncmodUpdate1) SetMaxconn(v int32)`

SetMaxconn sets Maxconn field to given value.

### HasMaxconn

`func (o *RsyncmodUpdate1) HasMaxconn() bool`

HasMaxconn returns a boolean if a field has been set.

### GetUser

`func (o *RsyncmodUpdate1) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RsyncmodUpdate1) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RsyncmodUpdate1) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *RsyncmodUpdate1) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetGroup

`func (o *RsyncmodUpdate1) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *RsyncmodUpdate1) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *RsyncmodUpdate1) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *RsyncmodUpdate1) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetHostsallow

`func (o *RsyncmodUpdate1) GetHostsallow() []string`

GetHostsallow returns the Hostsallow field if non-nil, zero value otherwise.

### GetHostsallowOk

`func (o *RsyncmodUpdate1) GetHostsallowOk() (*[]string, bool)`

GetHostsallowOk returns a tuple with the Hostsallow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsallow

`func (o *RsyncmodUpdate1) SetHostsallow(v []string)`

SetHostsallow sets Hostsallow field to given value.

### HasHostsallow

`func (o *RsyncmodUpdate1) HasHostsallow() bool`

HasHostsallow returns a boolean if a field has been set.

### GetHostsdeny

`func (o *RsyncmodUpdate1) GetHostsdeny() []string`

GetHostsdeny returns the Hostsdeny field if non-nil, zero value otherwise.

### GetHostsdenyOk

`func (o *RsyncmodUpdate1) GetHostsdenyOk() (*[]string, bool)`

GetHostsdenyOk returns a tuple with the Hostsdeny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsdeny

`func (o *RsyncmodUpdate1) SetHostsdeny(v []string)`

SetHostsdeny sets Hostsdeny field to given value.

### HasHostsdeny

`func (o *RsyncmodUpdate1) HasHostsdeny() bool`

HasHostsdeny returns a boolean if a field has been set.

### GetAuxiliary

`func (o *RsyncmodUpdate1) GetAuxiliary() string`

GetAuxiliary returns the Auxiliary field if non-nil, zero value otherwise.

### GetAuxiliaryOk

`func (o *RsyncmodUpdate1) GetAuxiliaryOk() (*string, bool)`

GetAuxiliaryOk returns a tuple with the Auxiliary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxiliary

`func (o *RsyncmodUpdate1) SetAuxiliary(v string)`

SetAuxiliary sets Auxiliary field to given value.

### HasAuxiliary

`func (o *RsyncmodUpdate1) HasAuxiliary() bool`

HasAuxiliary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


