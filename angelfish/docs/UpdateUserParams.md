# UpdateUserParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **int32** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Group** | Pointer to **int32** |  | [optional] 
**Home** | Pointer to **string** |  | [optional] 
**HomeMode** | Pointer to **string** |  | [optional] 
**Shell** | Pointer to **string** |  | [optional] 
**FullName** | Pointer to **string** |  | [optional] 
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

### NewUpdateUserParams

`func NewUpdateUserParams() *UpdateUserParams`

NewUpdateUserParams instantiates a new UpdateUserParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserParamsWithDefaults

`func NewUpdateUserParamsWithDefaults() *UpdateUserParams`

NewUpdateUserParamsWithDefaults instantiates a new UpdateUserParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *UpdateUserParams) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *UpdateUserParams) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *UpdateUserParams) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *UpdateUserParams) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUsername

`func (o *UpdateUserParams) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateUserParams) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateUserParams) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateUserParams) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetGroup

`func (o *UpdateUserParams) GetGroup() int32`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *UpdateUserParams) GetGroupOk() (*int32, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *UpdateUserParams) SetGroup(v int32)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *UpdateUserParams) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetHome

`func (o *UpdateUserParams) GetHome() string`

GetHome returns the Home field if non-nil, zero value otherwise.

### GetHomeOk

`func (o *UpdateUserParams) GetHomeOk() (*string, bool)`

GetHomeOk returns a tuple with the Home field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHome

`func (o *UpdateUserParams) SetHome(v string)`

SetHome sets Home field to given value.

### HasHome

`func (o *UpdateUserParams) HasHome() bool`

HasHome returns a boolean if a field has been set.

### GetHomeMode

`func (o *UpdateUserParams) GetHomeMode() string`

GetHomeMode returns the HomeMode field if non-nil, zero value otherwise.

### GetHomeModeOk

`func (o *UpdateUserParams) GetHomeModeOk() (*string, bool)`

GetHomeModeOk returns a tuple with the HomeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeMode

`func (o *UpdateUserParams) SetHomeMode(v string)`

SetHomeMode sets HomeMode field to given value.

### HasHomeMode

`func (o *UpdateUserParams) HasHomeMode() bool`

HasHomeMode returns a boolean if a field has been set.

### GetShell

`func (o *UpdateUserParams) GetShell() string`

GetShell returns the Shell field if non-nil, zero value otherwise.

### GetShellOk

`func (o *UpdateUserParams) GetShellOk() (*string, bool)`

GetShellOk returns a tuple with the Shell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShell

`func (o *UpdateUserParams) SetShell(v string)`

SetShell sets Shell field to given value.

### HasShell

`func (o *UpdateUserParams) HasShell() bool`

HasShell returns a boolean if a field has been set.

### GetFullName

`func (o *UpdateUserParams) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *UpdateUserParams) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *UpdateUserParams) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *UpdateUserParams) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateUserParams) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateUserParams) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateUserParams) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateUserParams) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *UpdateUserParams) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UpdateUserParams) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPassword

`func (o *UpdateUserParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateUserParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateUserParams) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateUserParams) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordDisabled

`func (o *UpdateUserParams) GetPasswordDisabled() bool`

GetPasswordDisabled returns the PasswordDisabled field if non-nil, zero value otherwise.

### GetPasswordDisabledOk

`func (o *UpdateUserParams) GetPasswordDisabledOk() (*bool, bool)`

GetPasswordDisabledOk returns a tuple with the PasswordDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordDisabled

`func (o *UpdateUserParams) SetPasswordDisabled(v bool)`

SetPasswordDisabled sets PasswordDisabled field to given value.

### HasPasswordDisabled

`func (o *UpdateUserParams) HasPasswordDisabled() bool`

HasPasswordDisabled returns a boolean if a field has been set.

### GetLocked

`func (o *UpdateUserParams) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *UpdateUserParams) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *UpdateUserParams) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *UpdateUserParams) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetMicrosoftAccount

`func (o *UpdateUserParams) GetMicrosoftAccount() bool`

GetMicrosoftAccount returns the MicrosoftAccount field if non-nil, zero value otherwise.

### GetMicrosoftAccountOk

`func (o *UpdateUserParams) GetMicrosoftAccountOk() (*bool, bool)`

GetMicrosoftAccountOk returns a tuple with the MicrosoftAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoftAccount

`func (o *UpdateUserParams) SetMicrosoftAccount(v bool)`

SetMicrosoftAccount sets MicrosoftAccount field to given value.

### HasMicrosoftAccount

`func (o *UpdateUserParams) HasMicrosoftAccount() bool`

HasMicrosoftAccount returns a boolean if a field has been set.

### GetSmb

`func (o *UpdateUserParams) GetSmb() bool`

GetSmb returns the Smb field if non-nil, zero value otherwise.

### GetSmbOk

`func (o *UpdateUserParams) GetSmbOk() (*bool, bool)`

GetSmbOk returns a tuple with the Smb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmb

`func (o *UpdateUserParams) SetSmb(v bool)`

SetSmb sets Smb field to given value.

### HasSmb

`func (o *UpdateUserParams) HasSmb() bool`

HasSmb returns a boolean if a field has been set.

### GetSudo

`func (o *UpdateUserParams) GetSudo() bool`

GetSudo returns the Sudo field if non-nil, zero value otherwise.

### GetSudoOk

`func (o *UpdateUserParams) GetSudoOk() (*bool, bool)`

GetSudoOk returns a tuple with the Sudo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudo

`func (o *UpdateUserParams) SetSudo(v bool)`

SetSudo sets Sudo field to given value.

### HasSudo

`func (o *UpdateUserParams) HasSudo() bool`

HasSudo returns a boolean if a field has been set.

### GetSudoNopasswd

`func (o *UpdateUserParams) GetSudoNopasswd() bool`

GetSudoNopasswd returns the SudoNopasswd field if non-nil, zero value otherwise.

### GetSudoNopasswdOk

`func (o *UpdateUserParams) GetSudoNopasswdOk() (*bool, bool)`

GetSudoNopasswdOk returns a tuple with the SudoNopasswd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudoNopasswd

`func (o *UpdateUserParams) SetSudoNopasswd(v bool)`

SetSudoNopasswd sets SudoNopasswd field to given value.

### HasSudoNopasswd

`func (o *UpdateUserParams) HasSudoNopasswd() bool`

HasSudoNopasswd returns a boolean if a field has been set.

### GetSudoCommands

`func (o *UpdateUserParams) GetSudoCommands() []string`

GetSudoCommands returns the SudoCommands field if non-nil, zero value otherwise.

### GetSudoCommandsOk

`func (o *UpdateUserParams) GetSudoCommandsOk() (*[]string, bool)`

GetSudoCommandsOk returns a tuple with the SudoCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudoCommands

`func (o *UpdateUserParams) SetSudoCommands(v []string)`

SetSudoCommands sets SudoCommands field to given value.

### HasSudoCommands

`func (o *UpdateUserParams) HasSudoCommands() bool`

HasSudoCommands returns a boolean if a field has been set.

### GetSshpubkey

`func (o *UpdateUserParams) GetSshpubkey() string`

GetSshpubkey returns the Sshpubkey field if non-nil, zero value otherwise.

### GetSshpubkeyOk

`func (o *UpdateUserParams) GetSshpubkeyOk() (*string, bool)`

GetSshpubkeyOk returns a tuple with the Sshpubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshpubkey

`func (o *UpdateUserParams) SetSshpubkey(v string)`

SetSshpubkey sets Sshpubkey field to given value.

### HasSshpubkey

`func (o *UpdateUserParams) HasSshpubkey() bool`

HasSshpubkey returns a boolean if a field has been set.

### SetSshpubkeyNil

`func (o *UpdateUserParams) SetSshpubkeyNil(b bool)`

 SetSshpubkeyNil sets the value for Sshpubkey to be an explicit nil

### UnsetSshpubkey
`func (o *UpdateUserParams) UnsetSshpubkey()`

UnsetSshpubkey ensures that no value is present for Sshpubkey, not even an explicit nil
### GetGroups

`func (o *UpdateUserParams) GetGroups() []int32`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UpdateUserParams) GetGroupsOk() (*[]int32, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UpdateUserParams) SetGroups(v []int32)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *UpdateUserParams) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetAttributes

`func (o *UpdateUserParams) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UpdateUserParams) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UpdateUserParams) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *UpdateUserParams) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


