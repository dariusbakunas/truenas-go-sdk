# KmipUpdate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**ForceClear** | Pointer to **bool** |  | [optional] 
**ManageSedDisks** | Pointer to **bool** |  | [optional] 
**ManageZfsKeys** | Pointer to **bool** |  | [optional] 
**ChangeServer** | Pointer to **bool** |  | [optional] 
**Validate** | Pointer to **bool** |  | [optional] 
**Certificate** | Pointer to **NullableInt32** |  | [optional] 
**CertificateAuthority** | Pointer to **NullableInt32** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Server** | Pointer to **string** |  | [optional] 

## Methods

### NewKmipUpdate0

`func NewKmipUpdate0() *KmipUpdate0`

NewKmipUpdate0 instantiates a new KmipUpdate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmipUpdate0WithDefaults

`func NewKmipUpdate0WithDefaults() *KmipUpdate0`

NewKmipUpdate0WithDefaults instantiates a new KmipUpdate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *KmipUpdate0) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *KmipUpdate0) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *KmipUpdate0) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *KmipUpdate0) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetForceClear

`func (o *KmipUpdate0) GetForceClear() bool`

GetForceClear returns the ForceClear field if non-nil, zero value otherwise.

### GetForceClearOk

`func (o *KmipUpdate0) GetForceClearOk() (*bool, bool)`

GetForceClearOk returns a tuple with the ForceClear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceClear

`func (o *KmipUpdate0) SetForceClear(v bool)`

SetForceClear sets ForceClear field to given value.

### HasForceClear

`func (o *KmipUpdate0) HasForceClear() bool`

HasForceClear returns a boolean if a field has been set.

### GetManageSedDisks

`func (o *KmipUpdate0) GetManageSedDisks() bool`

GetManageSedDisks returns the ManageSedDisks field if non-nil, zero value otherwise.

### GetManageSedDisksOk

`func (o *KmipUpdate0) GetManageSedDisksOk() (*bool, bool)`

GetManageSedDisksOk returns a tuple with the ManageSedDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageSedDisks

`func (o *KmipUpdate0) SetManageSedDisks(v bool)`

SetManageSedDisks sets ManageSedDisks field to given value.

### HasManageSedDisks

`func (o *KmipUpdate0) HasManageSedDisks() bool`

HasManageSedDisks returns a boolean if a field has been set.

### GetManageZfsKeys

`func (o *KmipUpdate0) GetManageZfsKeys() bool`

GetManageZfsKeys returns the ManageZfsKeys field if non-nil, zero value otherwise.

### GetManageZfsKeysOk

`func (o *KmipUpdate0) GetManageZfsKeysOk() (*bool, bool)`

GetManageZfsKeysOk returns a tuple with the ManageZfsKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageZfsKeys

`func (o *KmipUpdate0) SetManageZfsKeys(v bool)`

SetManageZfsKeys sets ManageZfsKeys field to given value.

### HasManageZfsKeys

`func (o *KmipUpdate0) HasManageZfsKeys() bool`

HasManageZfsKeys returns a boolean if a field has been set.

### GetChangeServer

`func (o *KmipUpdate0) GetChangeServer() bool`

GetChangeServer returns the ChangeServer field if non-nil, zero value otherwise.

### GetChangeServerOk

`func (o *KmipUpdate0) GetChangeServerOk() (*bool, bool)`

GetChangeServerOk returns a tuple with the ChangeServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeServer

`func (o *KmipUpdate0) SetChangeServer(v bool)`

SetChangeServer sets ChangeServer field to given value.

### HasChangeServer

`func (o *KmipUpdate0) HasChangeServer() bool`

HasChangeServer returns a boolean if a field has been set.

### GetValidate

`func (o *KmipUpdate0) GetValidate() bool`

GetValidate returns the Validate field if non-nil, zero value otherwise.

### GetValidateOk

`func (o *KmipUpdate0) GetValidateOk() (*bool, bool)`

GetValidateOk returns a tuple with the Validate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidate

`func (o *KmipUpdate0) SetValidate(v bool)`

SetValidate sets Validate field to given value.

### HasValidate

`func (o *KmipUpdate0) HasValidate() bool`

HasValidate returns a boolean if a field has been set.

### GetCertificate

`func (o *KmipUpdate0) GetCertificate() int32`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *KmipUpdate0) GetCertificateOk() (*int32, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *KmipUpdate0) SetCertificate(v int32)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *KmipUpdate0) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *KmipUpdate0) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *KmipUpdate0) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetCertificateAuthority

`func (o *KmipUpdate0) GetCertificateAuthority() int32`

GetCertificateAuthority returns the CertificateAuthority field if non-nil, zero value otherwise.

### GetCertificateAuthorityOk

`func (o *KmipUpdate0) GetCertificateAuthorityOk() (*int32, bool)`

GetCertificateAuthorityOk returns a tuple with the CertificateAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthority

`func (o *KmipUpdate0) SetCertificateAuthority(v int32)`

SetCertificateAuthority sets CertificateAuthority field to given value.

### HasCertificateAuthority

`func (o *KmipUpdate0) HasCertificateAuthority() bool`

HasCertificateAuthority returns a boolean if a field has been set.

### SetCertificateAuthorityNil

`func (o *KmipUpdate0) SetCertificateAuthorityNil(b bool)`

 SetCertificateAuthorityNil sets the value for CertificateAuthority to be an explicit nil

### UnsetCertificateAuthority
`func (o *KmipUpdate0) UnsetCertificateAuthority()`

UnsetCertificateAuthority ensures that no value is present for CertificateAuthority, not even an explicit nil
### GetPort

`func (o *KmipUpdate0) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *KmipUpdate0) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *KmipUpdate0) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *KmipUpdate0) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetServer

`func (o *KmipUpdate0) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *KmipUpdate0) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *KmipUpdate0) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *KmipUpdate0) HasServer() bool`

HasServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


