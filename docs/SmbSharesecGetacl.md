# SmbSharesecGetacl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShareName** | Pointer to **string** |  | [optional] 
**Options** | Pointer to [**SmbSharesecGetacl1**](SmbSharesecGetacl1.md) |  | [optional] [default to {}]

## Methods

### NewSmbSharesecGetacl

`func NewSmbSharesecGetacl() *SmbSharesecGetacl`

NewSmbSharesecGetacl instantiates a new SmbSharesecGetacl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmbSharesecGetaclWithDefaults

`func NewSmbSharesecGetaclWithDefaults() *SmbSharesecGetacl`

NewSmbSharesecGetaclWithDefaults instantiates a new SmbSharesecGetacl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShareName

`func (o *SmbSharesecGetacl) GetShareName() string`

GetShareName returns the ShareName field if non-nil, zero value otherwise.

### GetShareNameOk

`func (o *SmbSharesecGetacl) GetShareNameOk() (*string, bool)`

GetShareNameOk returns a tuple with the ShareName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareName

`func (o *SmbSharesecGetacl) SetShareName(v string)`

SetShareName sets ShareName field to given value.

### HasShareName

`func (o *SmbSharesecGetacl) HasShareName() bool`

HasShareName returns a boolean if a field has been set.

### GetOptions

`func (o *SmbSharesecGetacl) GetOptions() SmbSharesecGetacl1`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SmbSharesecGetacl) GetOptionsOk() (*SmbSharesecGetacl1, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SmbSharesecGetacl) SetOptions(v SmbSharesecGetacl1)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SmbSharesecGetacl) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


