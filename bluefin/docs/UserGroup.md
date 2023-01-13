# UserGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**BsdgrpGid** | Pointer to **int32** |  | [optional] 
**BsdgrpGroup** | Pointer to **string** |  | [optional] 
**BsdgrpBuiltin** | Pointer to **bool** |  | [optional] 
**BsdgrpSudo** | Pointer to **bool** |  | [optional] 
**BsdgrpSudoNopasswd** | Pointer to **bool** |  | [optional] 
**BsdgrpSudoCommands** | Pointer to **[]string** |  | [optional] 
**BsdgrpSmb** | Pointer to **bool** |  | [optional] 

## Methods

### NewUserGroup

`func NewUserGroup() *UserGroup`

NewUserGroup instantiates a new UserGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupWithDefaults

`func NewUserGroupWithDefaults() *UserGroup`

NewUserGroupWithDefaults instantiates a new UserGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserGroup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserGroup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserGroup) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBsdgrpGid

`func (o *UserGroup) GetBsdgrpGid() int32`

GetBsdgrpGid returns the BsdgrpGid field if non-nil, zero value otherwise.

### GetBsdgrpGidOk

`func (o *UserGroup) GetBsdgrpGidOk() (*int32, bool)`

GetBsdgrpGidOk returns a tuple with the BsdgrpGid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsdgrpGid

`func (o *UserGroup) SetBsdgrpGid(v int32)`

SetBsdgrpGid sets BsdgrpGid field to given value.

### HasBsdgrpGid

`func (o *UserGroup) HasBsdgrpGid() bool`

HasBsdgrpGid returns a boolean if a field has been set.

### GetBsdgrpGroup

`func (o *UserGroup) GetBsdgrpGroup() string`

GetBsdgrpGroup returns the BsdgrpGroup field if non-nil, zero value otherwise.

### GetBsdgrpGroupOk

`func (o *UserGroup) GetBsdgrpGroupOk() (*string, bool)`

GetBsdgrpGroupOk returns a tuple with the BsdgrpGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsdgrpGroup

`func (o *UserGroup) SetBsdgrpGroup(v string)`

SetBsdgrpGroup sets BsdgrpGroup field to given value.

### HasBsdgrpGroup

`func (o *UserGroup) HasBsdgrpGroup() bool`

HasBsdgrpGroup returns a boolean if a field has been set.

### GetBsdgrpBuiltin

`func (o *UserGroup) GetBsdgrpBuiltin() bool`

GetBsdgrpBuiltin returns the BsdgrpBuiltin field if non-nil, zero value otherwise.

### GetBsdgrpBuiltinOk

`func (o *UserGroup) GetBsdgrpBuiltinOk() (*bool, bool)`

GetBsdgrpBuiltinOk returns a tuple with the BsdgrpBuiltin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsdgrpBuiltin

`func (o *UserGroup) SetBsdgrpBuiltin(v bool)`

SetBsdgrpBuiltin sets BsdgrpBuiltin field to given value.

### HasBsdgrpBuiltin

`func (o *UserGroup) HasBsdgrpBuiltin() bool`

HasBsdgrpBuiltin returns a boolean if a field has been set.

### GetBsdgrpSudo

`func (o *UserGroup) GetBsdgrpSudo() bool`

GetBsdgrpSudo returns the BsdgrpSudo field if non-nil, zero value otherwise.

### GetBsdgrpSudoOk

`func (o *UserGroup) GetBsdgrpSudoOk() (*bool, bool)`

GetBsdgrpSudoOk returns a tuple with the BsdgrpSudo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsdgrpSudo

`func (o *UserGroup) SetBsdgrpSudo(v bool)`

SetBsdgrpSudo sets BsdgrpSudo field to given value.

### HasBsdgrpSudo

`func (o *UserGroup) HasBsdgrpSudo() bool`

HasBsdgrpSudo returns a boolean if a field has been set.

### GetBsdgrpSudoNopasswd

`func (o *UserGroup) GetBsdgrpSudoNopasswd() bool`

GetBsdgrpSudoNopasswd returns the BsdgrpSudoNopasswd field if non-nil, zero value otherwise.

### GetBsdgrpSudoNopasswdOk

`func (o *UserGroup) GetBsdgrpSudoNopasswdOk() (*bool, bool)`

GetBsdgrpSudoNopasswdOk returns a tuple with the BsdgrpSudoNopasswd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsdgrpSudoNopasswd

`func (o *UserGroup) SetBsdgrpSudoNopasswd(v bool)`

SetBsdgrpSudoNopasswd sets BsdgrpSudoNopasswd field to given value.

### HasBsdgrpSudoNopasswd

`func (o *UserGroup) HasBsdgrpSudoNopasswd() bool`

HasBsdgrpSudoNopasswd returns a boolean if a field has been set.

### GetBsdgrpSudoCommands

`func (o *UserGroup) GetBsdgrpSudoCommands() []string`

GetBsdgrpSudoCommands returns the BsdgrpSudoCommands field if non-nil, zero value otherwise.

### GetBsdgrpSudoCommandsOk

`func (o *UserGroup) GetBsdgrpSudoCommandsOk() (*[]string, bool)`

GetBsdgrpSudoCommandsOk returns a tuple with the BsdgrpSudoCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsdgrpSudoCommands

`func (o *UserGroup) SetBsdgrpSudoCommands(v []string)`

SetBsdgrpSudoCommands sets BsdgrpSudoCommands field to given value.

### HasBsdgrpSudoCommands

`func (o *UserGroup) HasBsdgrpSudoCommands() bool`

HasBsdgrpSudoCommands returns a boolean if a field has been set.

### GetBsdgrpSmb

`func (o *UserGroup) GetBsdgrpSmb() bool`

GetBsdgrpSmb returns the BsdgrpSmb field if non-nil, zero value otherwise.

### GetBsdgrpSmbOk

`func (o *UserGroup) GetBsdgrpSmbOk() (*bool, bool)`

GetBsdgrpSmbOk returns a tuple with the BsdgrpSmb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsdgrpSmb

`func (o *UserGroup) SetBsdgrpSmb(v bool)`

SetBsdgrpSmb sets BsdgrpSmb field to given value.

### HasBsdgrpSmb

`func (o *UserGroup) HasBsdgrpSmb() bool`

HasBsdgrpSmb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


