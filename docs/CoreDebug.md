# CoreDebug

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Engine** | Pointer to [**CoreDebug0**](CoreDebug0.md) |  | [optional] 
**Options** | Pointer to [**CoreDebug1**](CoreDebug1.md) |  | [optional] [default to {}]

## Methods

### NewCoreDebug

`func NewCoreDebug() *CoreDebug`

NewCoreDebug instantiates a new CoreDebug object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreDebugWithDefaults

`func NewCoreDebugWithDefaults() *CoreDebug`

NewCoreDebugWithDefaults instantiates a new CoreDebug object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngine

`func (o *CoreDebug) GetEngine() CoreDebug0`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *CoreDebug) GetEngineOk() (*CoreDebug0, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *CoreDebug) SetEngine(v CoreDebug0)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *CoreDebug) HasEngine() bool`

HasEngine returns a boolean if a field has been set.

### GetOptions

`func (o *CoreDebug) GetOptions() CoreDebug1`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CoreDebug) GetOptionsOk() (*CoreDebug1, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CoreDebug) SetOptions(v CoreDebug1)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CoreDebug) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


