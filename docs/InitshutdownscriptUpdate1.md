# InitshutdownscriptUpdate1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Command** | Pointer to **NullableString** |  | [optional] 
**ScriptText** | Pointer to **NullableString** |  | [optional] 
**Script** | Pointer to **NullableString** |  | [optional] 
**When** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Timeout** | Pointer to **int32** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 

## Methods

### NewInitshutdownscriptUpdate1

`func NewInitshutdownscriptUpdate1() *InitshutdownscriptUpdate1`

NewInitshutdownscriptUpdate1 instantiates a new InitshutdownscriptUpdate1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitshutdownscriptUpdate1WithDefaults

`func NewInitshutdownscriptUpdate1WithDefaults() *InitshutdownscriptUpdate1`

NewInitshutdownscriptUpdate1WithDefaults instantiates a new InitshutdownscriptUpdate1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InitshutdownscriptUpdate1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InitshutdownscriptUpdate1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InitshutdownscriptUpdate1) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InitshutdownscriptUpdate1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCommand

`func (o *InitshutdownscriptUpdate1) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *InitshutdownscriptUpdate1) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *InitshutdownscriptUpdate1) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *InitshutdownscriptUpdate1) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### SetCommandNil

`func (o *InitshutdownscriptUpdate1) SetCommandNil(b bool)`

 SetCommandNil sets the value for Command to be an explicit nil

### UnsetCommand
`func (o *InitshutdownscriptUpdate1) UnsetCommand()`

UnsetCommand ensures that no value is present for Command, not even an explicit nil
### GetScriptText

`func (o *InitshutdownscriptUpdate1) GetScriptText() string`

GetScriptText returns the ScriptText field if non-nil, zero value otherwise.

### GetScriptTextOk

`func (o *InitshutdownscriptUpdate1) GetScriptTextOk() (*string, bool)`

GetScriptTextOk returns a tuple with the ScriptText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptText

`func (o *InitshutdownscriptUpdate1) SetScriptText(v string)`

SetScriptText sets ScriptText field to given value.

### HasScriptText

`func (o *InitshutdownscriptUpdate1) HasScriptText() bool`

HasScriptText returns a boolean if a field has been set.

### SetScriptTextNil

`func (o *InitshutdownscriptUpdate1) SetScriptTextNil(b bool)`

 SetScriptTextNil sets the value for ScriptText to be an explicit nil

### UnsetScriptText
`func (o *InitshutdownscriptUpdate1) UnsetScriptText()`

UnsetScriptText ensures that no value is present for ScriptText, not even an explicit nil
### GetScript

`func (o *InitshutdownscriptUpdate1) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *InitshutdownscriptUpdate1) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *InitshutdownscriptUpdate1) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *InitshutdownscriptUpdate1) HasScript() bool`

HasScript returns a boolean if a field has been set.

### SetScriptNil

`func (o *InitshutdownscriptUpdate1) SetScriptNil(b bool)`

 SetScriptNil sets the value for Script to be an explicit nil

### UnsetScript
`func (o *InitshutdownscriptUpdate1) UnsetScript()`

UnsetScript ensures that no value is present for Script, not even an explicit nil
### GetWhen

`func (o *InitshutdownscriptUpdate1) GetWhen() string`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *InitshutdownscriptUpdate1) GetWhenOk() (*string, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *InitshutdownscriptUpdate1) SetWhen(v string)`

SetWhen sets When field to given value.

### HasWhen

`func (o *InitshutdownscriptUpdate1) HasWhen() bool`

HasWhen returns a boolean if a field has been set.

### GetEnabled

`func (o *InitshutdownscriptUpdate1) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InitshutdownscriptUpdate1) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InitshutdownscriptUpdate1) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InitshutdownscriptUpdate1) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTimeout

`func (o *InitshutdownscriptUpdate1) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *InitshutdownscriptUpdate1) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *InitshutdownscriptUpdate1) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *InitshutdownscriptUpdate1) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetComment

`func (o *InitshutdownscriptUpdate1) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *InitshutdownscriptUpdate1) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *InitshutdownscriptUpdate1) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *InitshutdownscriptUpdate1) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


