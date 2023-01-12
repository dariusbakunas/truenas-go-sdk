# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Gid** | Pointer to **int32** |  | [optional] 
**Group** | **string** |  | 
**Builtin** | Pointer to **bool** |  | [optional] 
**Sudo** | Pointer to **bool** |  | [optional] 
**SudoNopasswd** | Pointer to **bool** |  | [optional] 
**SudoCommands** | Pointer to **[]string** |  | [optional] 
**Smb** | Pointer to **bool** |  | [optional] 
**Users** | Pointer to **[]int32** |  | [optional] 
**Local** | Pointer to **bool** |  | [optional] 
**IdTypeBoth** | Pointer to **bool** |  | [optional] 

## Methods

### NewGroup

`func NewGroup(id int32, group string, ) *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Group) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Group) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Group) SetId(v int32)`

SetId sets Id field to given value.


### GetGid

`func (o *Group) GetGid() int32`

GetGid returns the Gid field if non-nil, zero value otherwise.

### GetGidOk

`func (o *Group) GetGidOk() (*int32, bool)`

GetGidOk returns a tuple with the Gid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGid

`func (o *Group) SetGid(v int32)`

SetGid sets Gid field to given value.

### HasGid

`func (o *Group) HasGid() bool`

HasGid returns a boolean if a field has been set.

### GetGroup

`func (o *Group) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Group) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Group) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetBuiltin

`func (o *Group) GetBuiltin() bool`

GetBuiltin returns the Builtin field if non-nil, zero value otherwise.

### GetBuiltinOk

`func (o *Group) GetBuiltinOk() (*bool, bool)`

GetBuiltinOk returns a tuple with the Builtin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltin

`func (o *Group) SetBuiltin(v bool)`

SetBuiltin sets Builtin field to given value.

### HasBuiltin

`func (o *Group) HasBuiltin() bool`

HasBuiltin returns a boolean if a field has been set.

### GetSudo

`func (o *Group) GetSudo() bool`

GetSudo returns the Sudo field if non-nil, zero value otherwise.

### GetSudoOk

`func (o *Group) GetSudoOk() (*bool, bool)`

GetSudoOk returns a tuple with the Sudo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudo

`func (o *Group) SetSudo(v bool)`

SetSudo sets Sudo field to given value.

### HasSudo

`func (o *Group) HasSudo() bool`

HasSudo returns a boolean if a field has been set.

### GetSudoNopasswd

`func (o *Group) GetSudoNopasswd() bool`

GetSudoNopasswd returns the SudoNopasswd field if non-nil, zero value otherwise.

### GetSudoNopasswdOk

`func (o *Group) GetSudoNopasswdOk() (*bool, bool)`

GetSudoNopasswdOk returns a tuple with the SudoNopasswd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudoNopasswd

`func (o *Group) SetSudoNopasswd(v bool)`

SetSudoNopasswd sets SudoNopasswd field to given value.

### HasSudoNopasswd

`func (o *Group) HasSudoNopasswd() bool`

HasSudoNopasswd returns a boolean if a field has been set.

### GetSudoCommands

`func (o *Group) GetSudoCommands() []string`

GetSudoCommands returns the SudoCommands field if non-nil, zero value otherwise.

### GetSudoCommandsOk

`func (o *Group) GetSudoCommandsOk() (*[]string, bool)`

GetSudoCommandsOk returns a tuple with the SudoCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudoCommands

`func (o *Group) SetSudoCommands(v []string)`

SetSudoCommands sets SudoCommands field to given value.

### HasSudoCommands

`func (o *Group) HasSudoCommands() bool`

HasSudoCommands returns a boolean if a field has been set.

### GetSmb

`func (o *Group) GetSmb() bool`

GetSmb returns the Smb field if non-nil, zero value otherwise.

### GetSmbOk

`func (o *Group) GetSmbOk() (*bool, bool)`

GetSmbOk returns a tuple with the Smb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmb

`func (o *Group) SetSmb(v bool)`

SetSmb sets Smb field to given value.

### HasSmb

`func (o *Group) HasSmb() bool`

HasSmb returns a boolean if a field has been set.

### GetUsers

`func (o *Group) GetUsers() []int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *Group) GetUsersOk() (*[]int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *Group) SetUsers(v []int32)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *Group) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetLocal

`func (o *Group) GetLocal() bool`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *Group) GetLocalOk() (*bool, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *Group) SetLocal(v bool)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *Group) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### GetIdTypeBoth

`func (o *Group) GetIdTypeBoth() bool`

GetIdTypeBoth returns the IdTypeBoth field if non-nil, zero value otherwise.

### GetIdTypeBothOk

`func (o *Group) GetIdTypeBothOk() (*bool, bool)`

GetIdTypeBothOk returns a tuple with the IdTypeBoth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTypeBoth

`func (o *Group) SetIdTypeBoth(v bool)`

SetIdTypeBoth sets IdTypeBoth field to given value.

### HasIdTypeBoth

`func (o *Group) HasIdTypeBoth() bool`

HasIdTypeBoth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


