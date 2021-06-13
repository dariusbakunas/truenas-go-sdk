# CertificateCreate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tos** | Pointer to **bool** |  | [optional] 
**DnsMapping** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**CsrId** | Pointer to **int32** |  | [optional] 
**Signedby** | Pointer to **int32** |  | [optional] 
**KeyLength** | Pointer to **int32** |  | [optional] 
**RenewDays** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **int32** |  | [optional] 
**Lifetime** | Pointer to **int32** |  | [optional] 
**Serial** | Pointer to **int32** |  | [optional] 
**AcmeDirectoryUri** | Pointer to **string** |  | [optional] 
**Certificate** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**Common** | Pointer to **NullableString** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**CSR** | Pointer to **string** |  | [optional] 
**EcCurve** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**KeyType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**OrganizationalUnit** | Pointer to **string** |  | [optional] 
**Passphrase** | Pointer to **string** |  | [optional] 
**Privatekey** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**CreateType** | Pointer to **string** |  | [optional] 
**DigestAlgorithm** | Pointer to **string** |  | [optional] 
**San** | Pointer to **[]string** |  | [optional] 
**CertExtensions** | Pointer to [**CertificateCreate0CertExtensions**](CertificateCreate0CertExtensions.md) |  | [optional] 

## Methods

### NewCertificateCreate0

`func NewCertificateCreate0() *CertificateCreate0`

NewCertificateCreate0 instantiates a new CertificateCreate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateCreate0WithDefaults

`func NewCertificateCreate0WithDefaults() *CertificateCreate0`

NewCertificateCreate0WithDefaults instantiates a new CertificateCreate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTos

`func (o *CertificateCreate0) GetTos() bool`

GetTos returns the Tos field if non-nil, zero value otherwise.

### GetTosOk

`func (o *CertificateCreate0) GetTosOk() (*bool, bool)`

GetTosOk returns a tuple with the Tos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTos

`func (o *CertificateCreate0) SetTos(v bool)`

SetTos sets Tos field to given value.

### HasTos

`func (o *CertificateCreate0) HasTos() bool`

HasTos returns a boolean if a field has been set.

### GetDnsMapping

`func (o *CertificateCreate0) GetDnsMapping() map[string]map[string]interface{}`

GetDnsMapping returns the DnsMapping field if non-nil, zero value otherwise.

### GetDnsMappingOk

`func (o *CertificateCreate0) GetDnsMappingOk() (*map[string]map[string]interface{}, bool)`

GetDnsMappingOk returns a tuple with the DnsMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsMapping

`func (o *CertificateCreate0) SetDnsMapping(v map[string]map[string]interface{})`

SetDnsMapping sets DnsMapping field to given value.

### HasDnsMapping

`func (o *CertificateCreate0) HasDnsMapping() bool`

HasDnsMapping returns a boolean if a field has been set.

### GetCsrId

`func (o *CertificateCreate0) GetCsrId() int32`

GetCsrId returns the CsrId field if non-nil, zero value otherwise.

### GetCsrIdOk

`func (o *CertificateCreate0) GetCsrIdOk() (*int32, bool)`

GetCsrIdOk returns a tuple with the CsrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrId

`func (o *CertificateCreate0) SetCsrId(v int32)`

SetCsrId sets CsrId field to given value.

### HasCsrId

`func (o *CertificateCreate0) HasCsrId() bool`

HasCsrId returns a boolean if a field has been set.

### GetSignedby

`func (o *CertificateCreate0) GetSignedby() int32`

GetSignedby returns the Signedby field if non-nil, zero value otherwise.

### GetSignedbyOk

`func (o *CertificateCreate0) GetSignedbyOk() (*int32, bool)`

GetSignedbyOk returns a tuple with the Signedby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedby

`func (o *CertificateCreate0) SetSignedby(v int32)`

SetSignedby sets Signedby field to given value.

### HasSignedby

`func (o *CertificateCreate0) HasSignedby() bool`

HasSignedby returns a boolean if a field has been set.

### GetKeyLength

`func (o *CertificateCreate0) GetKeyLength() int32`

GetKeyLength returns the KeyLength field if non-nil, zero value otherwise.

### GetKeyLengthOk

`func (o *CertificateCreate0) GetKeyLengthOk() (*int32, bool)`

GetKeyLengthOk returns a tuple with the KeyLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyLength

`func (o *CertificateCreate0) SetKeyLength(v int32)`

SetKeyLength sets KeyLength field to given value.

### HasKeyLength

`func (o *CertificateCreate0) HasKeyLength() bool`

HasKeyLength returns a boolean if a field has been set.

### GetRenewDays

`func (o *CertificateCreate0) GetRenewDays() int32`

GetRenewDays returns the RenewDays field if non-nil, zero value otherwise.

### GetRenewDaysOk

`func (o *CertificateCreate0) GetRenewDaysOk() (*int32, bool)`

GetRenewDaysOk returns a tuple with the RenewDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewDays

`func (o *CertificateCreate0) SetRenewDays(v int32)`

SetRenewDays sets RenewDays field to given value.

### HasRenewDays

`func (o *CertificateCreate0) HasRenewDays() bool`

HasRenewDays returns a boolean if a field has been set.

### GetType

`func (o *CertificateCreate0) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CertificateCreate0) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CertificateCreate0) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *CertificateCreate0) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLifetime

`func (o *CertificateCreate0) GetLifetime() int32`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *CertificateCreate0) GetLifetimeOk() (*int32, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *CertificateCreate0) SetLifetime(v int32)`

SetLifetime sets Lifetime field to given value.

### HasLifetime

`func (o *CertificateCreate0) HasLifetime() bool`

HasLifetime returns a boolean if a field has been set.

### GetSerial

`func (o *CertificateCreate0) GetSerial() int32`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *CertificateCreate0) GetSerialOk() (*int32, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *CertificateCreate0) SetSerial(v int32)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *CertificateCreate0) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetAcmeDirectoryUri

`func (o *CertificateCreate0) GetAcmeDirectoryUri() string`

GetAcmeDirectoryUri returns the AcmeDirectoryUri field if non-nil, zero value otherwise.

### GetAcmeDirectoryUriOk

`func (o *CertificateCreate0) GetAcmeDirectoryUriOk() (*string, bool)`

GetAcmeDirectoryUriOk returns a tuple with the AcmeDirectoryUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeDirectoryUri

`func (o *CertificateCreate0) SetAcmeDirectoryUri(v string)`

SetAcmeDirectoryUri sets AcmeDirectoryUri field to given value.

### HasAcmeDirectoryUri

`func (o *CertificateCreate0) HasAcmeDirectoryUri() bool`

HasAcmeDirectoryUri returns a boolean if a field has been set.

### GetCertificate

`func (o *CertificateCreate0) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CertificateCreate0) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CertificateCreate0) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *CertificateCreate0) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetCity

`func (o *CertificateCreate0) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CertificateCreate0) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CertificateCreate0) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CertificateCreate0) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCommon

`func (o *CertificateCreate0) GetCommon() string`

GetCommon returns the Common field if non-nil, zero value otherwise.

### GetCommonOk

`func (o *CertificateCreate0) GetCommonOk() (*string, bool)`

GetCommonOk returns a tuple with the Common field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommon

`func (o *CertificateCreate0) SetCommon(v string)`

SetCommon sets Common field to given value.

### HasCommon

`func (o *CertificateCreate0) HasCommon() bool`

HasCommon returns a boolean if a field has been set.

### SetCommonNil

`func (o *CertificateCreate0) SetCommonNil(b bool)`

 SetCommonNil sets the value for Common to be an explicit nil

### UnsetCommon
`func (o *CertificateCreate0) UnsetCommon()`

UnsetCommon ensures that no value is present for Common, not even an explicit nil
### GetCountry

`func (o *CertificateCreate0) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CertificateCreate0) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CertificateCreate0) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CertificateCreate0) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCSR

`func (o *CertificateCreate0) GetCSR() string`

GetCSR returns the CSR field if non-nil, zero value otherwise.

### GetCSROk

`func (o *CertificateCreate0) GetCSROk() (*string, bool)`

GetCSROk returns a tuple with the CSR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCSR

`func (o *CertificateCreate0) SetCSR(v string)`

SetCSR sets CSR field to given value.

### HasCSR

`func (o *CertificateCreate0) HasCSR() bool`

HasCSR returns a boolean if a field has been set.

### GetEcCurve

`func (o *CertificateCreate0) GetEcCurve() string`

GetEcCurve returns the EcCurve field if non-nil, zero value otherwise.

### GetEcCurveOk

`func (o *CertificateCreate0) GetEcCurveOk() (*string, bool)`

GetEcCurveOk returns a tuple with the EcCurve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcCurve

`func (o *CertificateCreate0) SetEcCurve(v string)`

SetEcCurve sets EcCurve field to given value.

### HasEcCurve

`func (o *CertificateCreate0) HasEcCurve() bool`

HasEcCurve returns a boolean if a field has been set.

### GetEmail

`func (o *CertificateCreate0) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CertificateCreate0) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CertificateCreate0) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CertificateCreate0) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetKeyType

`func (o *CertificateCreate0) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *CertificateCreate0) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *CertificateCreate0) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *CertificateCreate0) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetName

`func (o *CertificateCreate0) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateCreate0) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateCreate0) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CertificateCreate0) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganization

`func (o *CertificateCreate0) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *CertificateCreate0) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *CertificateCreate0) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *CertificateCreate0) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOrganizationalUnit

`func (o *CertificateCreate0) GetOrganizationalUnit() string`

GetOrganizationalUnit returns the OrganizationalUnit field if non-nil, zero value otherwise.

### GetOrganizationalUnitOk

`func (o *CertificateCreate0) GetOrganizationalUnitOk() (*string, bool)`

GetOrganizationalUnitOk returns a tuple with the OrganizationalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalUnit

`func (o *CertificateCreate0) SetOrganizationalUnit(v string)`

SetOrganizationalUnit sets OrganizationalUnit field to given value.

### HasOrganizationalUnit

`func (o *CertificateCreate0) HasOrganizationalUnit() bool`

HasOrganizationalUnit returns a boolean if a field has been set.

### GetPassphrase

`func (o *CertificateCreate0) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *CertificateCreate0) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *CertificateCreate0) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *CertificateCreate0) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetPrivatekey

`func (o *CertificateCreate0) GetPrivatekey() string`

GetPrivatekey returns the Privatekey field if non-nil, zero value otherwise.

### GetPrivatekeyOk

`func (o *CertificateCreate0) GetPrivatekeyOk() (*string, bool)`

GetPrivatekeyOk returns a tuple with the Privatekey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivatekey

`func (o *CertificateCreate0) SetPrivatekey(v string)`

SetPrivatekey sets Privatekey field to given value.

### HasPrivatekey

`func (o *CertificateCreate0) HasPrivatekey() bool`

HasPrivatekey returns a boolean if a field has been set.

### GetState

`func (o *CertificateCreate0) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CertificateCreate0) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CertificateCreate0) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CertificateCreate0) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCreateType

`func (o *CertificateCreate0) GetCreateType() string`

GetCreateType returns the CreateType field if non-nil, zero value otherwise.

### GetCreateTypeOk

`func (o *CertificateCreate0) GetCreateTypeOk() (*string, bool)`

GetCreateTypeOk returns a tuple with the CreateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateType

`func (o *CertificateCreate0) SetCreateType(v string)`

SetCreateType sets CreateType field to given value.

### HasCreateType

`func (o *CertificateCreate0) HasCreateType() bool`

HasCreateType returns a boolean if a field has been set.

### GetDigestAlgorithm

`func (o *CertificateCreate0) GetDigestAlgorithm() string`

GetDigestAlgorithm returns the DigestAlgorithm field if non-nil, zero value otherwise.

### GetDigestAlgorithmOk

`func (o *CertificateCreate0) GetDigestAlgorithmOk() (*string, bool)`

GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestAlgorithm

`func (o *CertificateCreate0) SetDigestAlgorithm(v string)`

SetDigestAlgorithm sets DigestAlgorithm field to given value.

### HasDigestAlgorithm

`func (o *CertificateCreate0) HasDigestAlgorithm() bool`

HasDigestAlgorithm returns a boolean if a field has been set.

### GetSan

`func (o *CertificateCreate0) GetSan() []string`

GetSan returns the San field if non-nil, zero value otherwise.

### GetSanOk

`func (o *CertificateCreate0) GetSanOk() (*[]string, bool)`

GetSanOk returns a tuple with the San field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSan

`func (o *CertificateCreate0) SetSan(v []string)`

SetSan sets San field to given value.

### HasSan

`func (o *CertificateCreate0) HasSan() bool`

HasSan returns a boolean if a field has been set.

### GetCertExtensions

`func (o *CertificateCreate0) GetCertExtensions() CertificateCreate0CertExtensions`

GetCertExtensions returns the CertExtensions field if non-nil, zero value otherwise.

### GetCertExtensionsOk

`func (o *CertificateCreate0) GetCertExtensionsOk() (*CertificateCreate0CertExtensions, bool)`

GetCertExtensionsOk returns a tuple with the CertExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertExtensions

`func (o *CertificateCreate0) SetCertExtensions(v CertificateCreate0CertExtensions)`

SetCertExtensions sets CertExtensions field to given value.

### HasCertExtensions

`func (o *CertificateCreate0) HasCertExtensions() bool`

HasCertExtensions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


