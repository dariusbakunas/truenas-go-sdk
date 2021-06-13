# FailoverCallRemote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | Pointer to **string** |  | [optional] 
**Args** | Pointer to **[]interface{}** |  | [optional] [default to []]
**Options** | Pointer to [**FailoverCallRemote2**](FailoverCallRemote2.md) |  | [optional] [default to {}]

## Methods

### NewFailoverCallRemote

`func NewFailoverCallRemote() *FailoverCallRemote`

NewFailoverCallRemote instantiates a new FailoverCallRemote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailoverCallRemoteWithDefaults

`func NewFailoverCallRemoteWithDefaults() *FailoverCallRemote`

NewFailoverCallRemoteWithDefaults instantiates a new FailoverCallRemote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *FailoverCallRemote) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *FailoverCallRemote) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *FailoverCallRemote) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *FailoverCallRemote) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetArgs

`func (o *FailoverCallRemote) GetArgs() []interface{}`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *FailoverCallRemote) GetArgsOk() (*[]interface{}, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *FailoverCallRemote) SetArgs(v []interface{})`

SetArgs sets Args field to given value.

### HasArgs

`func (o *FailoverCallRemote) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetOptions

`func (o *FailoverCallRemote) GetOptions() FailoverCallRemote2`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *FailoverCallRemote) GetOptionsOk() (*FailoverCallRemote2, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *FailoverCallRemote) SetOptions(v FailoverCallRemote2)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *FailoverCallRemote) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


