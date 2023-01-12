# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Uid** | Pointer to **int32** |  | [optional] 
**Username** | **string** |  | 
**Unixhash** | Pointer to **string** |  | [optional] 
**Smbhash** | Pointer to **string** |  | [optional] 
**Home** | Pointer to **string** |  | [optional] 
**Shell** | Pointer to **string** |  | [optional] 
**FullName** | **string** |  | 
**Builtin** | Pointer to **bool** |  | [optional] 
**Smb** | Pointer to **bool** |  | [optional] 
**PasswordDisabled** | Pointer to **bool** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**Sudo** | Pointer to **bool** |  | [optional] 
**SudoNopasswd** | Pointer to **bool** |  | [optional] 
**SudoCommands** | Pointer to **[]string** |  | [optional] 
**MicrosoftAccount** | Pointer to **bool** |  | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**Group** | Pointer to [**UserGroup**](UserGroup.md) |  | [optional] 
**Groups** | Pointer to **[]int32** |  | [optional] 
**Sshpubkey** | Pointer to **NullableString** |  | [optional] 
**Local** | Pointer to **bool** |  | [optional] 
**IdTypeBoth** | Pointer to **bool** |  | [optional] 

## Methods

### NewUser

`func NewUser(id int32, username string, fullName string, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v int32)`

SetId sets Id field to given value.


### GetUid

`func (o *User) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *User) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *User) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *User) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetUnixhash

`func (o *User) GetUnixhash() string`

GetUnixhash returns the Unixhash field if non-nil, zero value otherwise.

### GetUnixhashOk

`func (o *User) GetUnixhashOk() (*string, bool)`

GetUnixhashOk returns a tuple with the Unixhash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnixhash

`func (o *User) SetUnixhash(v string)`

SetUnixhash sets Unixhash field to given value.

### HasUnixhash

`func (o *User) HasUnixhash() bool`

HasUnixhash returns a boolean if a field has been set.

### GetSmbhash

`func (o *User) GetSmbhash() string`

GetSmbhash returns the Smbhash field if non-nil, zero value otherwise.

### GetSmbhashOk

`func (o *User) GetSmbhashOk() (*string, bool)`

GetSmbhashOk returns a tuple with the Smbhash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbhash

`func (o *User) SetSmbhash(v string)`

SetSmbhash sets Smbhash field to given value.

### HasSmbhash

`func (o *User) HasSmbhash() bool`

HasSmbhash returns a boolean if a field has been set.

### GetHome

`func (o *User) GetHome() string`

GetHome returns the Home field if non-nil, zero value otherwise.

### GetHomeOk

`func (o *User) GetHomeOk() (*string, bool)`

GetHomeOk returns a tuple with the Home field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHome

`func (o *User) SetHome(v string)`

SetHome sets Home field to given value.

### HasHome

`func (o *User) HasHome() bool`

HasHome returns a boolean if a field has been set.

### GetShell

`func (o *User) GetShell() string`

GetShell returns the Shell field if non-nil, zero value otherwise.

### GetShellOk

`func (o *User) GetShellOk() (*string, bool)`

GetShellOk returns a tuple with the Shell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShell

`func (o *User) SetShell(v string)`

SetShell sets Shell field to given value.

### HasShell

`func (o *User) HasShell() bool`

HasShell returns a boolean if a field has been set.

### GetFullName

`func (o *User) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *User) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *User) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetBuiltin

`func (o *User) GetBuiltin() bool`

GetBuiltin returns the Builtin field if non-nil, zero value otherwise.

### GetBuiltinOk

`func (o *User) GetBuiltinOk() (*bool, bool)`

GetBuiltinOk returns a tuple with the Builtin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltin

`func (o *User) SetBuiltin(v bool)`

SetBuiltin sets Builtin field to given value.

### HasBuiltin

`func (o *User) HasBuiltin() bool`

HasBuiltin returns a boolean if a field has been set.

### GetSmb

`func (o *User) GetSmb() bool`

GetSmb returns the Smb field if non-nil, zero value otherwise.

### GetSmbOk

`func (o *User) GetSmbOk() (*bool, bool)`

GetSmbOk returns a tuple with the Smb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmb

`func (o *User) SetSmb(v bool)`

SetSmb sets Smb field to given value.

### HasSmb

`func (o *User) HasSmb() bool`

HasSmb returns a boolean if a field has been set.

### GetPasswordDisabled

`func (o *User) GetPasswordDisabled() bool`

GetPasswordDisabled returns the PasswordDisabled field if non-nil, zero value otherwise.

### GetPasswordDisabledOk

`func (o *User) GetPasswordDisabledOk() (*bool, bool)`

GetPasswordDisabledOk returns a tuple with the PasswordDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordDisabled

`func (o *User) SetPasswordDisabled(v bool)`

SetPasswordDisabled sets PasswordDisabled field to given value.

### HasPasswordDisabled

`func (o *User) HasPasswordDisabled() bool`

HasPasswordDisabled returns a boolean if a field has been set.

### GetLocked

`func (o *User) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *User) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *User) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *User) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetSudo

`func (o *User) GetSudo() bool`

GetSudo returns the Sudo field if non-nil, zero value otherwise.

### GetSudoOk

`func (o *User) GetSudoOk() (*bool, bool)`

GetSudoOk returns a tuple with the Sudo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudo

`func (o *User) SetSudo(v bool)`

SetSudo sets Sudo field to given value.

### HasSudo

`func (o *User) HasSudo() bool`

HasSudo returns a boolean if a field has been set.

### GetSudoNopasswd

`func (o *User) GetSudoNopasswd() bool`

GetSudoNopasswd returns the SudoNopasswd field if non-nil, zero value otherwise.

### GetSudoNopasswdOk

`func (o *User) GetSudoNopasswdOk() (*bool, bool)`

GetSudoNopasswdOk returns a tuple with the SudoNopasswd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudoNopasswd

`func (o *User) SetSudoNopasswd(v bool)`

SetSudoNopasswd sets SudoNopasswd field to given value.

### HasSudoNopasswd

`func (o *User) HasSudoNopasswd() bool`

HasSudoNopasswd returns a boolean if a field has been set.

### GetSudoCommands

`func (o *User) GetSudoCommands() []string`

GetSudoCommands returns the SudoCommands field if non-nil, zero value otherwise.

### GetSudoCommandsOk

`func (o *User) GetSudoCommandsOk() (*[]string, bool)`

GetSudoCommandsOk returns a tuple with the SudoCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudoCommands

`func (o *User) SetSudoCommands(v []string)`

SetSudoCommands sets SudoCommands field to given value.

### HasSudoCommands

`func (o *User) HasSudoCommands() bool`

HasSudoCommands returns a boolean if a field has been set.

### GetMicrosoftAccount

`func (o *User) GetMicrosoftAccount() bool`

GetMicrosoftAccount returns the MicrosoftAccount field if non-nil, zero value otherwise.

### GetMicrosoftAccountOk

`func (o *User) GetMicrosoftAccountOk() (*bool, bool)`

GetMicrosoftAccountOk returns a tuple with the MicrosoftAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoftAccount

`func (o *User) SetMicrosoftAccount(v bool)`

SetMicrosoftAccount sets MicrosoftAccount field to given value.

### HasMicrosoftAccount

`func (o *User) HasMicrosoftAccount() bool`

HasMicrosoftAccount returns a boolean if a field has been set.

### GetAttributes

`func (o *User) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *User) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *User) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *User) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *User) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *User) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *User) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetGroup

`func (o *User) GetGroup() UserGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *User) GetGroupOk() (*UserGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *User) SetGroup(v UserGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *User) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetGroups

`func (o *User) GetGroups() []int32`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *User) GetGroupsOk() (*[]int32, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *User) SetGroups(v []int32)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *User) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetSshpubkey

`func (o *User) GetSshpubkey() string`

GetSshpubkey returns the Sshpubkey field if non-nil, zero value otherwise.

### GetSshpubkeyOk

`func (o *User) GetSshpubkeyOk() (*string, bool)`

GetSshpubkeyOk returns a tuple with the Sshpubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshpubkey

`func (o *User) SetSshpubkey(v string)`

SetSshpubkey sets Sshpubkey field to given value.

### HasSshpubkey

`func (o *User) HasSshpubkey() bool`

HasSshpubkey returns a boolean if a field has been set.

### SetSshpubkeyNil

`func (o *User) SetSshpubkeyNil(b bool)`

 SetSshpubkeyNil sets the value for Sshpubkey to be an explicit nil

### UnsetSshpubkey
`func (o *User) UnsetSshpubkey()`

UnsetSshpubkey ensures that no value is present for Sshpubkey, not even an explicit nil
### GetLocal

`func (o *User) GetLocal() bool`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *User) GetLocalOk() (*bool, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *User) SetLocal(v bool)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *User) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### GetIdTypeBoth

`func (o *User) GetIdTypeBoth() bool`

GetIdTypeBoth returns the IdTypeBoth field if non-nil, zero value otherwise.

### GetIdTypeBothOk

`func (o *User) GetIdTypeBothOk() (*bool, bool)`

GetIdTypeBothOk returns a tuple with the IdTypeBoth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTypeBoth

`func (o *User) SetIdTypeBoth(v bool)`

SetIdTypeBoth sets IdTypeBoth field to given value.

### HasIdTypeBoth

`func (o *User) HasIdTypeBoth() bool`

HasIdTypeBoth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


