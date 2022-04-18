# CreateGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gid** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**Smb** | Pointer to **bool** |  | [optional] 
**Sudo** | Pointer to **bool** |  | [optional] 
**SudoNopasswd** | Pointer to **bool** |  | [optional] 
**SudoCommands** | Pointer to **[]string** |  | [optional] 
**AllowDuplicateGid** | Pointer to **bool** |  | [optional] 
**Users** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewCreateGroupParams

`func NewCreateGroupParams(name string, ) *CreateGroupParams`

NewCreateGroupParams instantiates a new CreateGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGroupParamsWithDefaults

`func NewCreateGroupParamsWithDefaults() *CreateGroupParams`

NewCreateGroupParamsWithDefaults instantiates a new CreateGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGid

`func (o *CreateGroupParams) GetGid() int32`

GetGid returns the Gid field if non-nil, zero value otherwise.

### GetGidOk

`func (o *CreateGroupParams) GetGidOk() (*int32, bool)`

GetGidOk returns a tuple with the Gid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGid

`func (o *CreateGroupParams) SetGid(v int32)`

SetGid sets Gid field to given value.

### HasGid

`func (o *CreateGroupParams) HasGid() bool`

HasGid returns a boolean if a field has been set.

### GetName

`func (o *CreateGroupParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGroupParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGroupParams) SetName(v string)`

SetName sets Name field to given value.


### GetSmb

`func (o *CreateGroupParams) GetSmb() bool`

GetSmb returns the Smb field if non-nil, zero value otherwise.

### GetSmbOk

`func (o *CreateGroupParams) GetSmbOk() (*bool, bool)`

GetSmbOk returns a tuple with the Smb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmb

`func (o *CreateGroupParams) SetSmb(v bool)`

SetSmb sets Smb field to given value.

### HasSmb

`func (o *CreateGroupParams) HasSmb() bool`

HasSmb returns a boolean if a field has been set.

### GetSudo

`func (o *CreateGroupParams) GetSudo() bool`

GetSudo returns the Sudo field if non-nil, zero value otherwise.

### GetSudoOk

`func (o *CreateGroupParams) GetSudoOk() (*bool, bool)`

GetSudoOk returns a tuple with the Sudo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudo

`func (o *CreateGroupParams) SetSudo(v bool)`

SetSudo sets Sudo field to given value.

### HasSudo

`func (o *CreateGroupParams) HasSudo() bool`

HasSudo returns a boolean if a field has been set.

### GetSudoNopasswd

`func (o *CreateGroupParams) GetSudoNopasswd() bool`

GetSudoNopasswd returns the SudoNopasswd field if non-nil, zero value otherwise.

### GetSudoNopasswdOk

`func (o *CreateGroupParams) GetSudoNopasswdOk() (*bool, bool)`

GetSudoNopasswdOk returns a tuple with the SudoNopasswd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudoNopasswd

`func (o *CreateGroupParams) SetSudoNopasswd(v bool)`

SetSudoNopasswd sets SudoNopasswd field to given value.

### HasSudoNopasswd

`func (o *CreateGroupParams) HasSudoNopasswd() bool`

HasSudoNopasswd returns a boolean if a field has been set.

### GetSudoCommands

`func (o *CreateGroupParams) GetSudoCommands() []string`

GetSudoCommands returns the SudoCommands field if non-nil, zero value otherwise.

### GetSudoCommandsOk

`func (o *CreateGroupParams) GetSudoCommandsOk() (*[]string, bool)`

GetSudoCommandsOk returns a tuple with the SudoCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudoCommands

`func (o *CreateGroupParams) SetSudoCommands(v []string)`

SetSudoCommands sets SudoCommands field to given value.

### HasSudoCommands

`func (o *CreateGroupParams) HasSudoCommands() bool`

HasSudoCommands returns a boolean if a field has been set.

### GetAllowDuplicateGid

`func (o *CreateGroupParams) GetAllowDuplicateGid() bool`

GetAllowDuplicateGid returns the AllowDuplicateGid field if non-nil, zero value otherwise.

### GetAllowDuplicateGidOk

`func (o *CreateGroupParams) GetAllowDuplicateGidOk() (*bool, bool)`

GetAllowDuplicateGidOk returns a tuple with the AllowDuplicateGid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDuplicateGid

`func (o *CreateGroupParams) SetAllowDuplicateGid(v bool)`

SetAllowDuplicateGid sets AllowDuplicateGid field to given value.

### HasAllowDuplicateGid

`func (o *CreateGroupParams) HasAllowDuplicateGid() bool`

HasAllowDuplicateGid returns a boolean if a field has been set.

### GetUsers

`func (o *CreateGroupParams) GetUsers() []int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *CreateGroupParams) GetUsersOk() (*[]int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *CreateGroupParams) SetUsers(v []int32)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *CreateGroupParams) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


