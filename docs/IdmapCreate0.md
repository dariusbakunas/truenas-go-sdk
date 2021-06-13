# IdmapCreate0

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

### NewIdmapCreate0

`func NewIdmapCreate0() *IdmapCreate0`

NewIdmapCreate0 instantiates a new IdmapCreate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdmapCreate0WithDefaults

`func NewIdmapCreate0WithDefaults() *IdmapCreate0`

NewIdmapCreate0WithDefaults instantiates a new IdmapCreate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IdmapCreate0) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdmapCreate0) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdmapCreate0) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdmapCreate0) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDnsDomainName

`func (o *IdmapCreate0) GetDnsDomainName() string`

GetDnsDomainName returns the DnsDomainName field if non-nil, zero value otherwise.

### GetDnsDomainNameOk

`func (o *IdmapCreate0) GetDnsDomainNameOk() (*string, bool)`

GetDnsDomainNameOk returns a tuple with the DnsDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsDomainName

`func (o *IdmapCreate0) SetDnsDomainName(v string)`

SetDnsDomainName sets DnsDomainName field to given value.

### HasDnsDomainName

`func (o *IdmapCreate0) HasDnsDomainName() bool`

HasDnsDomainName returns a boolean if a field has been set.

### GetRangeLow

`func (o *IdmapCreate0) GetRangeLow() int32`

GetRangeLow returns the RangeLow field if non-nil, zero value otherwise.

### GetRangeLowOk

`func (o *IdmapCreate0) GetRangeLowOk() (*int32, bool)`

GetRangeLowOk returns a tuple with the RangeLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeLow

`func (o *IdmapCreate0) SetRangeLow(v int32)`

SetRangeLow sets RangeLow field to given value.

### HasRangeLow

`func (o *IdmapCreate0) HasRangeLow() bool`

HasRangeLow returns a boolean if a field has been set.

### GetRangeHigh

`func (o *IdmapCreate0) GetRangeHigh() int32`

GetRangeHigh returns the RangeHigh field if non-nil, zero value otherwise.

### GetRangeHighOk

`func (o *IdmapCreate0) GetRangeHighOk() (*int32, bool)`

GetRangeHighOk returns a tuple with the RangeHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeHigh

`func (o *IdmapCreate0) SetRangeHigh(v int32)`

SetRangeHigh sets RangeHigh field to given value.

### HasRangeHigh

`func (o *IdmapCreate0) HasRangeHigh() bool`

HasRangeHigh returns a boolean if a field has been set.

### GetIdmapBackend

`func (o *IdmapCreate0) GetIdmapBackend() string`

GetIdmapBackend returns the IdmapBackend field if non-nil, zero value otherwise.

### GetIdmapBackendOk

`func (o *IdmapCreate0) GetIdmapBackendOk() (*string, bool)`

GetIdmapBackendOk returns a tuple with the IdmapBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdmapBackend

`func (o *IdmapCreate0) SetIdmapBackend(v string)`

SetIdmapBackend sets IdmapBackend field to given value.

### HasIdmapBackend

`func (o *IdmapCreate0) HasIdmapBackend() bool`

HasIdmapBackend returns a boolean if a field has been set.

### GetCertificate

`func (o *IdmapCreate0) GetCertificate() int32`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *IdmapCreate0) GetCertificateOk() (*int32, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *IdmapCreate0) SetCertificate(v int32)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *IdmapCreate0) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *IdmapCreate0) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *IdmapCreate0) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetOptions

`func (o *IdmapCreate0) GetOptions() IdmapCreate0Options`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *IdmapCreate0) GetOptionsOk() (*IdmapCreate0Options, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *IdmapCreate0) SetOptions(v IdmapCreate0Options)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *IdmapCreate0) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


