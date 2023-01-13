# CreateUserParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **int32** |  | [optional] 
**Username** | **string** |  | 
**Group** | Pointer to **int32** |  | [optional] 
**GroupCreate** | Pointer to **bool** |  | [optional] 
**Home** | Pointer to **string** |  | [optional] 
**HomeMode** | Pointer to **string** |  | [optional] 
**Shell** | Pointer to **string** |  | [optional] 
**FullName** | **string** |  | 
**Email** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**PasswordDisabled** | Pointer to **bool** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**MicrosoftAccount** | Pointer to **bool** |  | [optional] 
**Smb** | Pointer to **bool** |  | [optional] 
**Sudo** | Pointer to **bool** |  | [optional] 
**SudoNopasswd** | Pointer to **bool** |  | [optional] 
**SudoCommands** | Pointer to **[]string** |  | [optional] 
**Sshpubkey** | Pointer to **NullableString** |  | [optional] 
**Groups** | Pointer to **[]int32** |  | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCreateUserParams

`func NewCreateUserParams(username string, fullName string, ) *CreateUserParams`

NewCreateUserParams instantiates a new CreateUserParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserParamsWithDefaults

`func NewCreateUserParamsWithDefaults() *CreateUserParams`

NewCreateUserParamsWithDefaults instantiates a new CreateUserParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *CreateUserParams) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *CreateUserParams) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *CreateUserParams) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *CreateUserParams) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUsername

`func (o *CreateUserParams) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateUserParams) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateUserParams) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetGroup

`func (o *CreateUserParams) GetGroup() int32`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *CreateUserParams) GetGroupOk() (*int32, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *CreateUserParams) SetGroup(v int32)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *CreateUserParams) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetGroupCreate

`func (o *CreateUserParams) GetGroupCreate() bool`

GetGroupCreate returns the GroupCreate field if non-nil, zero value otherwise.

### GetGroupCreateOk

`func (o *CreateUserParams) GetGroupCreateOk() (*bool, bool)`

GetGroupCreateOk returns a tuple with the GroupCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupCreate

`func (o *CreateUserParams) SetGroupCreate(v bool)`

SetGroupCreate sets GroupCreate field to given value.

### HasGroupCreate

`func (o *CreateUserParams) HasGroupCreate() bool`

HasGroupCreate returns a boolean if a field has been set.

### GetHome

`func (o *CreateUserParams) GetHome() string`

GetHome returns the Home field if non-nil, zero value otherwise.

### GetHomeOk

`func (o *CreateUserParams) GetHomeOk() (*string, bool)`

GetHomeOk returns a tuple with the Home field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHome

`func (o *CreateUserParams) SetHome(v string)`

SetHome sets Home field to given value.

### HasHome

`func (o *CreateUserParams) HasHome() bool`

HasHome returns a boolean if a field has been set.

### GetHomeMode

`func (o *CreateUserParams) GetHomeMode() string`

GetHomeMode returns the HomeMode field if non-nil, zero value otherwise.

### GetHomeModeOk

`func (o *CreateUserParams) GetHomeModeOk() (*string, bool)`

GetHomeModeOk returns a tuple with the HomeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeMode

`func (o *CreateUserParams) SetHomeMode(v string)`

SetHomeMode sets HomeMode field to given value.

### HasHomeMode

`func (o *CreateUserParams) HasHomeMode() bool`

HasHomeMode returns a boolean if a field has been set.

### GetShell

`func (o *CreateUserParams) GetShell() string`

GetShell returns the Shell field if non-nil, zero value otherwise.

### GetShellOk

`func (o *CreateUserParams) GetShellOk() (*string, bool)`

GetShellOk returns a tuple with the Shell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShell

`func (o *CreateUserParams) SetShell(v string)`

SetShell sets Shell field to given value.

### HasShell

`func (o *CreateUserParams) HasShell() bool`

HasShell returns a boolean if a field has been set.

### GetFullName

`func (o *CreateUserParams) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *CreateUserParams) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *CreateUserParams) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetEmail

`func (o *CreateUserParams) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateUserParams) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateUserParams) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateUserParams) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *CreateUserParams) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CreateUserParams) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPassword

`func (o *CreateUserParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateUserParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateUserParams) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateUserParams) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordDisabled

`func (o *CreateUserParams) GetPasswordDisabled() bool`

GetPasswordDisabled returns the PasswordDisabled field if non-nil, zero value otherwise.

### GetPasswordDisabledOk

`func (o *CreateUserParams) GetPasswordDisabledOk() (*bool, bool)`

GetPasswordDisabledOk returns a tuple with the PasswordDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordDisabled

`func (o *CreateUserParams) SetPasswordDisabled(v bool)`

SetPasswordDisabled sets PasswordDisabled field to given value.

### HasPasswordDisabled

`func (o *CreateUserParams) HasPasswordDisabled() bool`

HasPasswordDisabled returns a boolean if a field has been set.

### GetLocked

`func (o *CreateUserParams) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *CreateUserParams) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *CreateUserParams) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *CreateUserParams) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetMicrosoftAccount

`func (o *CreateUserParams) GetMicrosoftAccount() bool`

GetMicrosoftAccount returns the MicrosoftAccount field if non-nil, zero value otherwise.

### GetMicrosoftAccountOk

`func (o *CreateUserParams) GetMicrosoftAccountOk() (*bool, bool)`

GetMicrosoftAccountOk returns a tuple with the MicrosoftAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoftAccount

`func (o *CreateUserParams) SetMicrosoftAccount(v bool)`

SetMicrosoftAccount sets MicrosoftAccount field to given value.

### HasMicrosoftAccount

`func (o *CreateUserParams) HasMicrosoftAccount() bool`

HasMicrosoftAccount returns a boolean if a field has been set.

### GetSmb

`func (o *CreateUserParams) GetSmb() bool`

GetSmb returns the Smb field if non-nil, zero value otherwise.

### GetSmbOk

`func (o *CreateUserParams) GetSmbOk() (*bool, bool)`

GetSmbOk returns a tuple with the Smb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmb

`func (o *CreateUserParams) SetSmb(v bool)`

SetSmb sets Smb field to given value.

### HasSmb

`func (o *CreateUserParams) HasSmb() bool`

HasSmb returns a boolean if a field has been set.

### GetSudo

`func (o *CreateUserParams) GetSudo() bool`

GetSudo returns the Sudo field if non-nil, zero value otherwise.

### GetSudoOk

`func (o *CreateUserParams) GetSudoOk() (*bool, bool)`

GetSudoOk returns a tuple with the Sudo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudo

`func (o *CreateUserParams) SetSudo(v bool)`

SetSudo sets Sudo field to given value.

### HasSudo

`func (o *CreateUserParams) HasSudo() bool`

HasSudo returns a boolean if a field has been set.

### GetSudoNopasswd

`func (o *CreateUserParams) GetSudoNopasswd() bool`

GetSudoNopasswd returns the SudoNopasswd field if non-nil, zero value otherwise.

### GetSudoNopasswdOk

`func (o *CreateUserParams) GetSudoNopasswdOk() (*bool, bool)`

GetSudoNopasswdOk returns a tuple with the SudoNopasswd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudoNopasswd

`func (o *CreateUserParams) SetSudoNopasswd(v bool)`

SetSudoNopasswd sets SudoNopasswd field to given value.

### HasSudoNopasswd

`func (o *CreateUserParams) HasSudoNopasswd() bool`

HasSudoNopasswd returns a boolean if a field has been set.

### GetSudoCommands

`func (o *CreateUserParams) GetSudoCommands() []string`

GetSudoCommands returns the SudoCommands field if non-nil, zero value otherwise.

### GetSudoCommandsOk

`func (o *CreateUserParams) GetSudoCommandsOk() (*[]string, bool)`

GetSudoCommandsOk returns a tuple with the SudoCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudoCommands

`func (o *CreateUserParams) SetSudoCommands(v []string)`

SetSudoCommands sets SudoCommands field to given value.

### HasSudoCommands

`func (o *CreateUserParams) HasSudoCommands() bool`

HasSudoCommands returns a boolean if a field has been set.

### GetSshpubkey

`func (o *CreateUserParams) GetSshpubkey() string`

GetSshpubkey returns the Sshpubkey field if non-nil, zero value otherwise.

### GetSshpubkeyOk

`func (o *CreateUserParams) GetSshpubkeyOk() (*string, bool)`

GetSshpubkeyOk returns a tuple with the Sshpubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshpubkey

`func (o *CreateUserParams) SetSshpubkey(v string)`

SetSshpubkey sets Sshpubkey field to given value.

### HasSshpubkey

`func (o *CreateUserParams) HasSshpubkey() bool`

HasSshpubkey returns a boolean if a field has been set.

### SetSshpubkeyNil

`func (o *CreateUserParams) SetSshpubkeyNil(b bool)`

 SetSshpubkeyNil sets the value for Sshpubkey to be an explicit nil

### UnsetSshpubkey
`func (o *CreateUserParams) UnsetSshpubkey()`

UnsetSshpubkey ensures that no value is present for Sshpubkey, not even an explicit nil
### GetGroups

`func (o *CreateUserParams) GetGroups() []int32`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *CreateUserParams) GetGroupsOk() (*[]int32, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *CreateUserParams) SetGroups(v []int32)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *CreateUserParams) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetAttributes

`func (o *CreateUserParams) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CreateUserParams) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CreateUserParams) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CreateUserParams) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


