# JailExec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jail** | Pointer to **string** |  | [optional] 
**Command** | Pointer to **[]interface{}** |  | [optional] 
**Options** | Pointer to [**JailExec2**](JailExec2.md) |  | [optional] [default to {}]

## Methods

### NewJailExec

`func NewJailExec() *JailExec`

NewJailExec instantiates a new JailExec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJailExecWithDefaults

`func NewJailExecWithDefaults() *JailExec`

NewJailExecWithDefaults instantiates a new JailExec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJail

`func (o *JailExec) GetJail() string`

GetJail returns the Jail field if non-nil, zero value otherwise.

### GetJailOk

`func (o *JailExec) GetJailOk() (*string, bool)`

GetJailOk returns a tuple with the Jail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJail

`func (o *JailExec) SetJail(v string)`

SetJail sets Jail field to given value.

### HasJail

`func (o *JailExec) HasJail() bool`

HasJail returns a boolean if a field has been set.

### GetCommand

`func (o *JailExec) GetCommand() []interface{}`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *JailExec) GetCommandOk() (*[]interface{}, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *JailExec) SetCommand(v []interface{})`

SetCommand sets Command field to given value.

### HasCommand

`func (o *JailExec) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetOptions

`func (o *JailExec) GetOptions() JailExec2`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *JailExec) GetOptionsOk() (*JailExec2, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *JailExec) SetOptions(v JailExec2)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *JailExec) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


