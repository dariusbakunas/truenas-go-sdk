# IdmapUpdate1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**DnsDomainName** | Pointer to **string** |  | [optional] 
**RangeLow** | Pointer to **int32** |  | [optional] 
**RangeHigh** | Pointer to **int32** |  | [optional] 
**IdmapBackend** | Pointer to **string** |  | [optional] 
**Certificate** | Pointer to **NullableInt32** |  | [optional] 
**Options** | Pointer to [**IdmapCreate0Options**](IdmapCreate0Options.md) |  | [optional] 

## Methods

### NewIdmapUpdate1

`func NewIdmapUpdate1() *IdmapUpdate1`

NewIdmapUpdate1 instantiates a new IdmapUpdate1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdmapUpdate1WithDefaults

`func NewIdmapUpdate1WithDefaults() *IdmapUpdate1`

NewIdmapUpdate1WithDefaults instantiates a new IdmapUpdate1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IdmapUpdate1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdmapUpdate1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdmapUpdate1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdmapUpdate1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDnsDomainName

`func (o *IdmapUpdate1) GetDnsDomainName() string`

GetDnsDomainName returns the DnsDomainName field if non-nil, zero value otherwise.

### GetDnsDomainNameOk

`func (o *IdmapUpdate1) GetDnsDomainNameOk() (*string, bool)`

GetDnsDomainNameOk returns a tuple with the DnsDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsDomainName

`func (o *IdmapUpdate1) SetDnsDomainName(v string)`

SetDnsDomainName sets DnsDomainName field to given value.

### HasDnsDomainName

`func (o *IdmapUpdate1) HasDnsDomainName() bool`

HasDnsDomainName returns a boolean if a field has been set.

### GetRangeLow

`func (o *IdmapUpdate1) GetRangeLow() int32`

GetRangeLow returns the RangeLow field if non-nil, zero value otherwise.

### GetRangeLowOk

`func (o *IdmapUpdate1) GetRangeLowOk() (*int32, bool)`

GetRangeLowOk returns a tuple with the RangeLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeLow

`func (o *IdmapUpdate1) SetRangeLow(v int32)`

SetRangeLow sets RangeLow field to given value.

### HasRangeLow

`func (o *IdmapUpdate1) HasRangeLow() bool`

HasRangeLow returns a boolean if a field has been set.

### GetRangeHigh

`func (o *IdmapUpdate1) GetRangeHigh() int32`

GetRangeHigh returns the RangeHigh field if non-nil, zero value otherwise.

### GetRangeHighOk

`func (o *IdmapUpdate1) GetRangeHighOk() (*int32, bool)`

GetRangeHighOk returns a tuple with the RangeHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeHigh

`func (o *IdmapUpdate1) SetRangeHigh(v int32)`

SetRangeHigh sets RangeHigh field to given value.

### HasRangeHigh

`func (o *IdmapUpdate1) HasRangeHigh() bool`

HasRangeHigh returns a boolean if a field has been set.

### GetIdmapBackend

`func (o *IdmapUpdate1) GetIdmapBackend() string`

GetIdmapBackend returns the IdmapBackend field if non-nil, zero value otherwise.

### GetIdmapBackendOk

`func (o *IdmapUpdate1) GetIdmapBackendOk() (*string, bool)`

GetIdmapBackendOk returns a tuple with the IdmapBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdmapBackend

`func (o *IdmapUpdate1) SetIdmapBackend(v string)`

SetIdmapBackend sets IdmapBackend field to given value.

### HasIdmapBackend

`func (o *IdmapUpdate1) HasIdmapBackend() bool`

HasIdmapBackend returns a boolean if a field has been set.

### GetCertificate

`func (o *IdmapUpdate1) GetCertificate() int32`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *IdmapUpdate1) GetCertificateOk() (*int32, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *IdmapUpdate1) SetCertificate(v int32)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *IdmapUpdate1) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *IdmapUpdate1) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *IdmapUpdate1) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetOptions

`func (o *IdmapUpdate1) GetOptions() IdmapCreate0Options`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *IdmapUpdate1) GetOptionsOk() (*IdmapCreate0Options, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *IdmapUpdate1) SetOptions(v IdmapCreate0Options)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *IdmapUpdate1) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


