# FailoverControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**FailoverControl0**](FailoverControl0.md) |  | [optional] 
**Options** | Pointer to [**FailoverControl1**](FailoverControl1.md) |  | [optional] [default to {}]

## Methods

### NewFailoverControl

`func NewFailoverControl() *FailoverControl`

NewFailoverControl instantiates a new FailoverControl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailoverControlWithDefaults

`func NewFailoverControlWithDefaults() *FailoverControl`

NewFailoverControlWithDefaults instantiates a new FailoverControl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *FailoverControl) GetAction() FailoverControl0`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FailoverControl) GetActionOk() (*FailoverControl0, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FailoverControl) SetAction(v FailoverControl0)`

SetAction sets Action field to given value.

### HasAction

`func (o *FailoverControl) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetOptions

`func (o *FailoverControl) GetOptions() FailoverControl1`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *FailoverControl) GetOptionsOk() (*FailoverControl1, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *FailoverControl) SetOptions(v FailoverControl1)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *FailoverControl) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


