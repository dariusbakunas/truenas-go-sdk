# Certificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Certificate** | Pointer to **string** |  | [optional] 
**Privatekey** | Pointer to **string** |  | [optional] 
**RootPath** | Pointer to **string** |  | [optional] 
**CertificatePath** | Pointer to **string** |  | [optional] 
**PrivatekeyPath** | Pointer to **string** |  | [optional] 
**CsrPath** | Pointer to **string** |  | [optional] 
**CertType** | Pointer to **string** |  | [optional] 
**Revoked** | Pointer to **bool** |  | [optional] 
**CanBeRevoked** | Pointer to **bool** |  | [optional] 
**Internal** | Pointer to **string** |  | [optional] 
**CATypeExisting** | Pointer to **bool** |  | [optional] 
**CATypeInternal** | Pointer to **bool** |  | [optional] 
**CATypeIntermediate** | Pointer to **bool** |  | [optional] 
**CertTypeExisting** | Pointer to **bool** |  | [optional] 
**CertTypeInternal** | Pointer to **bool** |  | [optional] 
**CertTypeCSR** | Pointer to **bool** |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 
**ChainList** | Pointer to **[]string** |  | [optional] 
**KeyLength** | Pointer to **int32** |  | [optional] 
**KeyType** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**OrganizationalUnit** | Pointer to **string** |  | [optional] 
**Common** | Pointer to **string** |  | [optional] 
**San** | Pointer to **[]string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**DN** | Pointer to **string** |  | [optional] 
**SubjectNameHash** | Pointer to **int64** |  | [optional] 
**Extensions** | Pointer to **map[string]interface{}** |  | [optional] 
**DigestAlgorithm** | Pointer to **string** |  | [optional] 
**Lifetime** | Pointer to **int32** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**Until** | Pointer to **string** |  | [optional] 
**Serial** | Pointer to **int32** |  | [optional] 
**Chain** | Pointer to **bool** |  | [optional] 
**Fingerprint** | Pointer to **string** |  | [optional] 
**Parsed** | Pointer to **bool** |  | [optional] 

## Methods

### NewCertificate

`func NewCertificate() *Certificate`

NewCertificate instantiates a new Certificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateWithDefaults

`func NewCertificateWithDefaults() *Certificate`

NewCertificateWithDefaults instantiates a new Certificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Certificate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Certificate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Certificate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Certificate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Certificate) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Certificate) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Certificate) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *Certificate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *Certificate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Certificate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Certificate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Certificate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCertificate

`func (o *Certificate) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *Certificate) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *Certificate) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *Certificate) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetPrivatekey

`func (o *Certificate) GetPrivatekey() string`

GetPrivatekey returns the Privatekey field if non-nil, zero value otherwise.

### GetPrivatekeyOk

`func (o *Certificate) GetPrivatekeyOk() (*string, bool)`

GetPrivatekeyOk returns a tuple with the Privatekey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivatekey

`func (o *Certificate) SetPrivatekey(v string)`

SetPrivatekey sets Privatekey field to given value.

### HasPrivatekey

`func (o *Certificate) HasPrivatekey() bool`

HasPrivatekey returns a boolean if a field has been set.

### GetRootPath

`func (o *Certificate) GetRootPath() string`

GetRootPath returns the RootPath field if non-nil, zero value otherwise.

### GetRootPathOk

`func (o *Certificate) GetRootPathOk() (*string, bool)`

GetRootPathOk returns a tuple with the RootPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPath

`func (o *Certificate) SetRootPath(v string)`

SetRootPath sets RootPath field to given value.

### HasRootPath

`func (o *Certificate) HasRootPath() bool`

HasRootPath returns a boolean if a field has been set.

### GetCertificatePath

`func (o *Certificate) GetCertificatePath() string`

GetCertificatePath returns the CertificatePath field if non-nil, zero value otherwise.

### GetCertificatePathOk

`func (o *Certificate) GetCertificatePathOk() (*string, bool)`

GetCertificatePathOk returns a tuple with the CertificatePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePath

`func (o *Certificate) SetCertificatePath(v string)`

SetCertificatePath sets CertificatePath field to given value.

### HasCertificatePath

`func (o *Certificate) HasCertificatePath() bool`

HasCertificatePath returns a boolean if a field has been set.

### GetPrivatekeyPath

`func (o *Certificate) GetPrivatekeyPath() string`

GetPrivatekeyPath returns the PrivatekeyPath field if non-nil, zero value otherwise.

### GetPrivatekeyPathOk

`func (o *Certificate) GetPrivatekeyPathOk() (*string, bool)`

GetPrivatekeyPathOk returns a tuple with the PrivatekeyPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivatekeyPath

`func (o *Certificate) SetPrivatekeyPath(v string)`

SetPrivatekeyPath sets PrivatekeyPath field to given value.

### HasPrivatekeyPath

`func (o *Certificate) HasPrivatekeyPath() bool`

HasPrivatekeyPath returns a boolean if a field has been set.

### GetCsrPath

`func (o *Certificate) GetCsrPath() string`

GetCsrPath returns the CsrPath field if non-nil, zero value otherwise.

### GetCsrPathOk

`func (o *Certificate) GetCsrPathOk() (*string, bool)`

GetCsrPathOk returns a tuple with the CsrPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrPath

`func (o *Certificate) SetCsrPath(v string)`

SetCsrPath sets CsrPath field to given value.

### HasCsrPath

`func (o *Certificate) HasCsrPath() bool`

HasCsrPath returns a boolean if a field has been set.

### GetCertType

`func (o *Certificate) GetCertType() string`

GetCertType returns the CertType field if non-nil, zero value otherwise.

### GetCertTypeOk

`func (o *Certificate) GetCertTypeOk() (*string, bool)`

GetCertTypeOk returns a tuple with the CertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertType

`func (o *Certificate) SetCertType(v string)`

SetCertType sets CertType field to given value.

### HasCertType

`func (o *Certificate) HasCertType() bool`

HasCertType returns a boolean if a field has been set.

### GetRevoked

`func (o *Certificate) GetRevoked() bool`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *Certificate) GetRevokedOk() (*bool, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *Certificate) SetRevoked(v bool)`

SetRevoked sets Revoked field to given value.

### HasRevoked

`func (o *Certificate) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.

### GetCanBeRevoked

`func (o *Certificate) GetCanBeRevoked() bool`

GetCanBeRevoked returns the CanBeRevoked field if non-nil, zero value otherwise.

### GetCanBeRevokedOk

`func (o *Certificate) GetCanBeRevokedOk() (*bool, bool)`

GetCanBeRevokedOk returns a tuple with the CanBeRevoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanBeRevoked

`func (o *Certificate) SetCanBeRevoked(v bool)`

SetCanBeRevoked sets CanBeRevoked field to given value.

### HasCanBeRevoked

`func (o *Certificate) HasCanBeRevoked() bool`

HasCanBeRevoked returns a boolean if a field has been set.

### GetInternal

`func (o *Certificate) GetInternal() string`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *Certificate) GetInternalOk() (*string, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *Certificate) SetInternal(v string)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *Certificate) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetCATypeExisting

`func (o *Certificate) GetCATypeExisting() bool`

GetCATypeExisting returns the CATypeExisting field if non-nil, zero value otherwise.

### GetCATypeExistingOk

`func (o *Certificate) GetCATypeExistingOk() (*bool, bool)`

GetCATypeExistingOk returns a tuple with the CATypeExisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCATypeExisting

`func (o *Certificate) SetCATypeExisting(v bool)`

SetCATypeExisting sets CATypeExisting field to given value.

### HasCATypeExisting

`func (o *Certificate) HasCATypeExisting() bool`

HasCATypeExisting returns a boolean if a field has been set.

### GetCATypeInternal

`func (o *Certificate) GetCATypeInternal() bool`

GetCATypeInternal returns the CATypeInternal field if non-nil, zero value otherwise.

### GetCATypeInternalOk

`func (o *Certificate) GetCATypeInternalOk() (*bool, bool)`

GetCATypeInternalOk returns a tuple with the CATypeInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCATypeInternal

`func (o *Certificate) SetCATypeInternal(v bool)`

SetCATypeInternal sets CATypeInternal field to given value.

### HasCATypeInternal

`func (o *Certificate) HasCATypeInternal() bool`

HasCATypeInternal returns a boolean if a field has been set.

### GetCATypeIntermediate

`func (o *Certificate) GetCATypeIntermediate() bool`

GetCATypeIntermediate returns the CATypeIntermediate field if non-nil, zero value otherwise.

### GetCATypeIntermediateOk

`func (o *Certificate) GetCATypeIntermediateOk() (*bool, bool)`

GetCATypeIntermediateOk returns a tuple with the CATypeIntermediate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCATypeIntermediate

`func (o *Certificate) SetCATypeIntermediate(v bool)`

SetCATypeIntermediate sets CATypeIntermediate field to given value.

### HasCATypeIntermediate

`func (o *Certificate) HasCATypeIntermediate() bool`

HasCATypeIntermediate returns a boolean if a field has been set.

### GetCertTypeExisting

`func (o *Certificate) GetCertTypeExisting() bool`

GetCertTypeExisting returns the CertTypeExisting field if non-nil, zero value otherwise.

### GetCertTypeExistingOk

`func (o *Certificate) GetCertTypeExistingOk() (*bool, bool)`

GetCertTypeExistingOk returns a tuple with the CertTypeExisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertTypeExisting

`func (o *Certificate) SetCertTypeExisting(v bool)`

SetCertTypeExisting sets CertTypeExisting field to given value.

### HasCertTypeExisting

`func (o *Certificate) HasCertTypeExisting() bool`

HasCertTypeExisting returns a boolean if a field has been set.

### GetCertTypeInternal

`func (o *Certificate) GetCertTypeInternal() bool`

GetCertTypeInternal returns the CertTypeInternal field if non-nil, zero value otherwise.

### GetCertTypeInternalOk

`func (o *Certificate) GetCertTypeInternalOk() (*bool, bool)`

GetCertTypeInternalOk returns a tuple with the CertTypeInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertTypeInternal

`func (o *Certificate) SetCertTypeInternal(v bool)`

SetCertTypeInternal sets CertTypeInternal field to given value.

### HasCertTypeInternal

`func (o *Certificate) HasCertTypeInternal() bool`

HasCertTypeInternal returns a boolean if a field has been set.

### GetCertTypeCSR

`func (o *Certificate) GetCertTypeCSR() bool`

GetCertTypeCSR returns the CertTypeCSR field if non-nil, zero value otherwise.

### GetCertTypeCSROk

`func (o *Certificate) GetCertTypeCSROk() (*bool, bool)`

GetCertTypeCSROk returns a tuple with the CertTypeCSR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertTypeCSR

`func (o *Certificate) SetCertTypeCSR(v bool)`

SetCertTypeCSR sets CertTypeCSR field to given value.

### HasCertTypeCSR

`func (o *Certificate) HasCertTypeCSR() bool`

HasCertTypeCSR returns a boolean if a field has been set.

### GetIssuer

`func (o *Certificate) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *Certificate) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *Certificate) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *Certificate) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetChainList

`func (o *Certificate) GetChainList() []string`

GetChainList returns the ChainList field if non-nil, zero value otherwise.

### GetChainListOk

`func (o *Certificate) GetChainListOk() (*[]string, bool)`

GetChainListOk returns a tuple with the ChainList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainList

`func (o *Certificate) SetChainList(v []string)`

SetChainList sets ChainList field to given value.

### HasChainList

`func (o *Certificate) HasChainList() bool`

HasChainList returns a boolean if a field has been set.

### GetKeyLength

`func (o *Certificate) GetKeyLength() int32`

GetKeyLength returns the KeyLength field if non-nil, zero value otherwise.

### GetKeyLengthOk

`func (o *Certificate) GetKeyLengthOk() (*int32, bool)`

GetKeyLengthOk returns a tuple with the KeyLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyLength

`func (o *Certificate) SetKeyLength(v int32)`

SetKeyLength sets KeyLength field to given value.

### HasKeyLength

`func (o *Certificate) HasKeyLength() bool`

HasKeyLength returns a boolean if a field has been set.

### GetKeyType

`func (o *Certificate) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *Certificate) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *Certificate) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *Certificate) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetCountry

`func (o *Certificate) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Certificate) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Certificate) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Certificate) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetState

`func (o *Certificate) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Certificate) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Certificate) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Certificate) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCity

`func (o *Certificate) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Certificate) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Certificate) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Certificate) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetOrganization

`func (o *Certificate) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *Certificate) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *Certificate) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *Certificate) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOrganizationalUnit

`func (o *Certificate) GetOrganizationalUnit() string`

GetOrganizationalUnit returns the OrganizationalUnit field if non-nil, zero value otherwise.

### GetOrganizationalUnitOk

`func (o *Certificate) GetOrganizationalUnitOk() (*string, bool)`

GetOrganizationalUnitOk returns a tuple with the OrganizationalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalUnit

`func (o *Certificate) SetOrganizationalUnit(v string)`

SetOrganizationalUnit sets OrganizationalUnit field to given value.

### HasOrganizationalUnit

`func (o *Certificate) HasOrganizationalUnit() bool`

HasOrganizationalUnit returns a boolean if a field has been set.

### GetCommon

`func (o *Certificate) GetCommon() string`

GetCommon returns the Common field if non-nil, zero value otherwise.

### GetCommonOk

`func (o *Certificate) GetCommonOk() (*string, bool)`

GetCommonOk returns a tuple with the Common field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommon

`func (o *Certificate) SetCommon(v string)`

SetCommon sets Common field to given value.

### HasCommon

`func (o *Certificate) HasCommon() bool`

HasCommon returns a boolean if a field has been set.

### GetSan

`func (o *Certificate) GetSan() []string`

GetSan returns the San field if non-nil, zero value otherwise.

### GetSanOk

`func (o *Certificate) GetSanOk() (*[]string, bool)`

GetSanOk returns a tuple with the San field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSan

`func (o *Certificate) SetSan(v []string)`

SetSan sets San field to given value.

### HasSan

`func (o *Certificate) HasSan() bool`

HasSan returns a boolean if a field has been set.

### GetEmail

`func (o *Certificate) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Certificate) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Certificate) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Certificate) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDN

`func (o *Certificate) GetDN() string`

GetDN returns the DN field if non-nil, zero value otherwise.

### GetDNOk

`func (o *Certificate) GetDNOk() (*string, bool)`

GetDNOk returns a tuple with the DN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDN

`func (o *Certificate) SetDN(v string)`

SetDN sets DN field to given value.

### HasDN

`func (o *Certificate) HasDN() bool`

HasDN returns a boolean if a field has been set.

### GetSubjectNameHash

`func (o *Certificate) GetSubjectNameHash() int64`

GetSubjectNameHash returns the SubjectNameHash field if non-nil, zero value otherwise.

### GetSubjectNameHashOk

`func (o *Certificate) GetSubjectNameHashOk() (*int64, bool)`

GetSubjectNameHashOk returns a tuple with the SubjectNameHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectNameHash

`func (o *Certificate) SetSubjectNameHash(v int64)`

SetSubjectNameHash sets SubjectNameHash field to given value.

### HasSubjectNameHash

`func (o *Certificate) HasSubjectNameHash() bool`

HasSubjectNameHash returns a boolean if a field has been set.

### GetExtensions

`func (o *Certificate) GetExtensions() map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Certificate) GetExtensionsOk() (*map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Certificate) SetExtensions(v map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *Certificate) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetDigestAlgorithm

`func (o *Certificate) GetDigestAlgorithm() string`

GetDigestAlgorithm returns the DigestAlgorithm field if non-nil, zero value otherwise.

### GetDigestAlgorithmOk

`func (o *Certificate) GetDigestAlgorithmOk() (*string, bool)`

GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestAlgorithm

`func (o *Certificate) SetDigestAlgorithm(v string)`

SetDigestAlgorithm sets DigestAlgorithm field to given value.

### HasDigestAlgorithm

`func (o *Certificate) HasDigestAlgorithm() bool`

HasDigestAlgorithm returns a boolean if a field has been set.

### GetLifetime

`func (o *Certificate) GetLifetime() int32`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *Certificate) GetLifetimeOk() (*int32, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *Certificate) SetLifetime(v int32)`

SetLifetime sets Lifetime field to given value.

### HasLifetime

`func (o *Certificate) HasLifetime() bool`

HasLifetime returns a boolean if a field has been set.

### GetFrom

`func (o *Certificate) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Certificate) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *Certificate) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *Certificate) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetUntil

`func (o *Certificate) GetUntil() string`

GetUntil returns the Until field if non-nil, zero value otherwise.

### GetUntilOk

`func (o *Certificate) GetUntilOk() (*string, bool)`

GetUntilOk returns a tuple with the Until field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntil

`func (o *Certificate) SetUntil(v string)`

SetUntil sets Until field to given value.

### HasUntil

`func (o *Certificate) HasUntil() bool`

HasUntil returns a boolean if a field has been set.

### GetSerial

`func (o *Certificate) GetSerial() int32`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *Certificate) GetSerialOk() (*int32, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *Certificate) SetSerial(v int32)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *Certificate) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetChain

`func (o *Certificate) GetChain() bool`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *Certificate) GetChainOk() (*bool, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *Certificate) SetChain(v bool)`

SetChain sets Chain field to given value.

### HasChain

`func (o *Certificate) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetFingerprint

`func (o *Certificate) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *Certificate) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *Certificate) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *Certificate) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetParsed

`func (o *Certificate) GetParsed() bool`

GetParsed returns the Parsed field if non-nil, zero value otherwise.

### GetParsedOk

`func (o *Certificate) GetParsedOk() (*bool, bool)`

GetParsedOk returns a tuple with the Parsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsed

`func (o *Certificate) SetParsed(v bool)`

SetParsed sets Parsed field to given value.

### HasParsed

`func (o *Certificate) HasParsed() bool`

HasParsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


