# CompositeValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** |  | [optional] 
**Rawvalue** | **string** |  | 
**Source** | Pointer to **string** |  | [optional] 

## Methods

### NewCompositeValue

`func NewCompositeValue(rawvalue string, ) *CompositeValue`

NewCompositeValue instantiates a new CompositeValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompositeValueWithDefaults

`func NewCompositeValueWithDefaults() *CompositeValue`

NewCompositeValueWithDefaults instantiates a new CompositeValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *CompositeValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CompositeValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CompositeValue) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CompositeValue) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetRawvalue

`func (o *CompositeValue) GetRawvalue() string`

GetRawvalue returns the Rawvalue field if non-nil, zero value otherwise.

### GetRawvalueOk

`func (o *CompositeValue) GetRawvalueOk() (*string, bool)`

GetRawvalueOk returns a tuple with the Rawvalue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawvalue

`func (o *CompositeValue) SetRawvalue(v string)`

SetRawvalue sets Rawvalue field to given value.


### GetSource

`func (o *CompositeValue) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CompositeValue) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CompositeValue) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CompositeValue) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


