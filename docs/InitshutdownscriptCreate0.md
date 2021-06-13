# InitshutdownscriptCreate0

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

### NewInitshutdownscriptCreate0

`func NewInitshutdownscriptCreate0() *InitshutdownscriptCreate0`

NewInitshutdownscriptCreate0 instantiates a new InitshutdownscriptCreate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitshutdownscriptCreate0WithDefaults

`func NewInitshutdownscriptCreate0WithDefaults() *InitshutdownscriptCreate0`

NewInitshutdownscriptCreate0WithDefaults instantiates a new InitshutdownscriptCreate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InitshutdownscriptCreate0) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InitshutdownscriptCreate0) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InitshutdownscriptCreate0) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InitshutdownscriptCreate0) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCommand

`func (o *InitshutdownscriptCreate0) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *InitshutdownscriptCreate0) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *InitshutdownscriptCreate0) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *InitshutdownscriptCreate0) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### SetCommandNil

`func (o *InitshutdownscriptCreate0) SetCommandNil(b bool)`

 SetCommandNil sets the value for Command to be an explicit nil

### UnsetCommand
`func (o *InitshutdownscriptCreate0) UnsetCommand()`

UnsetCommand ensures that no value is present for Command, not even an explicit nil
### GetScriptText

`func (o *InitshutdownscriptCreate0) GetScriptText() string`

GetScriptText returns the ScriptText field if non-nil, zero value otherwise.

### GetScriptTextOk

`func (o *InitshutdownscriptCreate0) GetScriptTextOk() (*string, bool)`

GetScriptTextOk returns a tuple with the ScriptText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptText

`func (o *InitshutdownscriptCreate0) SetScriptText(v string)`

SetScriptText sets ScriptText field to given value.

### HasScriptText

`func (o *InitshutdownscriptCreate0) HasScriptText() bool`

HasScriptText returns a boolean if a field has been set.

### SetScriptTextNil

`func (o *InitshutdownscriptCreate0) SetScriptTextNil(b bool)`

 SetScriptTextNil sets the value for ScriptText to be an explicit nil

### UnsetScriptText
`func (o *InitshutdownscriptCreate0) UnsetScriptText()`

UnsetScriptText ensures that no value is present for ScriptText, not even an explicit nil
### GetScript

`func (o *InitshutdownscriptCreate0) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *InitshutdownscriptCreate0) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *InitshutdownscriptCreate0) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *InitshutdownscriptCreate0) HasScript() bool`

HasScript returns a boolean if a field has been set.

### SetScriptNil

`func (o *InitshutdownscriptCreate0) SetScriptNil(b bool)`

 SetScriptNil sets the value for Script to be an explicit nil

### UnsetScript
`func (o *InitshutdownscriptCreate0) UnsetScript()`

UnsetScript ensures that no value is present for Script, not even an explicit nil
### GetWhen

`func (o *InitshutdownscriptCreate0) GetWhen() string`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *InitshutdownscriptCreate0) GetWhenOk() (*string, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *InitshutdownscriptCreate0) SetWhen(v string)`

SetWhen sets When field to given value.

### HasWhen

`func (o *InitshutdownscriptCreate0) HasWhen() bool`

HasWhen returns a boolean if a field has been set.

### GetEnabled

`func (o *InitshutdownscriptCreate0) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InitshutdownscriptCreate0) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InitshutdownscriptCreate0) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InitshutdownscriptCreate0) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTimeout

`func (o *InitshutdownscriptCreate0) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *InitshutdownscriptCreate0) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *InitshutdownscriptCreate0) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *InitshutdownscriptCreate0) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetComment

`func (o *InitshutdownscriptCreate0) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *InitshutdownscriptCreate0) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *InitshutdownscriptCreate0) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *InitshutdownscriptCreate0) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


