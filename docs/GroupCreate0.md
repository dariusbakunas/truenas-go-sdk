# GroupCreate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gid** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Smb** | Pointer to **bool** |  | [optional] 
**Sudo** | Pointer to **bool** |  | [optional] 
**SudoNopasswd** | Pointer to **bool** |  | [optional] 
**SudoCommands** | Pointer to **[]string** |  | [optional] 
**AllowDuplicateGid** | Pointer to **bool** |  | [optional] 
**Users** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewGroupCreate0

`func NewGroupCreate0() *GroupCreate0`

NewGroupCreate0 instantiates a new GroupCreate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupCreate0WithDefaults

`func NewGroupCreate0WithDefaults() *GroupCreate0`

NewGroupCreate0WithDefaults instantiates a new GroupCreate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGid

`func (o *GroupCreate0) GetGid() int32`

GetGid returns the Gid field if non-nil, zero value otherwise.

### GetGidOk

`func (o *GroupCreate0) GetGidOk() (*int32, bool)`

GetGidOk returns a tuple with the Gid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGid

`func (o *GroupCreate0) SetGid(v int32)`

SetGid sets Gid field to given value.

### HasGid

`func (o *GroupCreate0) HasGid() bool`

HasGid returns a boolean if a field has been set.

### GetName

`func (o *GroupCreate0) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupCreate0) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupCreate0) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupCreate0) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSmb

`func (o *GroupCreate0) GetSmb() bool`

GetSmb returns the Smb field if non-nil, zero value otherwise.

### GetSmbOk

`func (o *GroupCreate0) GetSmbOk() (*bool, bool)`

GetSmbOk returns a tuple with the Smb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmb

`func (o *GroupCreate0) SetSmb(v bool)`

SetSmb sets Smb field to given value.

### HasSmb

`func (o *GroupCreate0) HasSmb() bool`

HasSmb returns a boolean if a field has been set.

### GetSudo

`func (o *GroupCreate0) GetSudo() bool`

GetSudo returns the Sudo field if non-nil, zero value otherwise.

### GetSudoOk

`func (o *GroupCreate0) GetSudoOk() (*bool, bool)`

GetSudoOk returns a tuple with the Sudo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudo

`func (o *GroupCreate0) SetSudo(v bool)`

SetSudo sets Sudo field to given value.

### HasSudo

`func (o *GroupCreate0) HasSudo() bool`

HasSudo returns a boolean if a field has been set.

### GetSudoNopasswd

`func (o *GroupCreate0) GetSudoNopasswd() bool`

GetSudoNopasswd returns the SudoNopasswd field if non-nil, zero value otherwise.

### GetSudoNopasswdOk

`func (o *GroupCreate0) GetSudoNopasswdOk() (*bool, bool)`

GetSudoNopasswdOk returns a tuple with the SudoNopasswd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudoNopasswd

`func (o *GroupCreate0) SetSudoNopasswd(v bool)`

SetSudoNopasswd sets SudoNopasswd field to given value.

### HasSudoNopasswd

`func (o *GroupCreate0) HasSudoNopasswd() bool`

HasSudoNopasswd returns a boolean if a field has been set.

### GetSudoCommands

`func (o *GroupCreate0) GetSudoCommands() []string`

GetSudoCommands returns the SudoCommands field if non-nil, zero value otherwise.

### GetSudoCommandsOk

`func (o *GroupCreate0) GetSudoCommandsOk() (*[]string, bool)`

GetSudoCommandsOk returns a tuple with the SudoCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudoCommands

`func (o *GroupCreate0) SetSudoCommands(v []string)`

SetSudoCommands sets SudoCommands field to given value.

### HasSudoCommands

`func (o *GroupCreate0) HasSudoCommands() bool`

HasSudoCommands returns a boolean if a field has been set.

### GetAllowDuplicateGid

`func (o *GroupCreate0) GetAllowDuplicateGid() bool`

GetAllowDuplicateGid returns the AllowDuplicateGid field if non-nil, zero value otherwise.

### GetAllowDuplicateGidOk

`func (o *GroupCreate0) GetAllowDuplicateGidOk() (*bool, bool)`

GetAllowDuplicateGidOk returns a tuple with the AllowDuplicateGid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDuplicateGid

`func (o *GroupCreate0) SetAllowDuplicateGid(v bool)`

SetAllowDuplicateGid sets AllowDuplicateGid field to given value.

### HasAllowDuplicateGid

`func (o *GroupCreate0) HasAllowDuplicateGid() bool`

HasAllowDuplicateGid returns a boolean if a field has been set.

### GetUsers

`func (o *GroupCreate0) GetUsers() []int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GroupCreate0) GetUsersOk() (*[]int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GroupCreate0) SetUsers(v []int32)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *GroupCreate0) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


