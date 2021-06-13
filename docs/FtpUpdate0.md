# FtpUpdate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **int32** |  | [optional] 
**Clients** | Pointer to **int32** |  | [optional] 
**Ipconnections** | Pointer to **int32** |  | [optional] 
**Loginattempt** | Pointer to **int32** |  | [optional] 
**Timeout** | Pointer to **int32** |  | [optional] 
**Rootlogin** | Pointer to **bool** |  | [optional] 
**Onlyanonymous** | Pointer to **bool** |  | [optional] 
**Anonpath** | Pointer to **NullableString** |  | [optional] 
**Onlylocal** | Pointer to **bool** |  | [optional] 
**Banner** | Pointer to **string** |  | [optional] 
**Filemask** | Pointer to **string** |  | [optional] 
**Dirmask** | Pointer to **string** |  | [optional] 
**Fxp** | Pointer to **bool** |  | [optional] 
**Resume** | Pointer to **bool** |  | [optional] 
**Defaultroot** | Pointer to **bool** |  | [optional] 
**Ident** | Pointer to **bool** |  | [optional] 
**Reversedns** | Pointer to **bool** |  | [optional] 
**Masqaddress** | Pointer to **string** |  | [optional] 
**Passiveportsmin** | Pointer to **int32** |  | [optional] 
**Passiveportsmax** | Pointer to **int32** |  | [optional] 
**Localuserbw** | Pointer to **int32** |  | [optional] 
**Localuserdlbw** | Pointer to **int32** |  | [optional] 
**Anonuserbw** | Pointer to **int32** |  | [optional] 
**Anonuserdlbw** | Pointer to **int32** |  | [optional] 
**Tls** | Pointer to **bool** |  | [optional] 
**TlsPolicy** | Pointer to **string** |  | [optional] 
**TlsOptAllowClientRenegotiations** | Pointer to **bool** |  | [optional] 
**TlsOptAllowDotLogin** | Pointer to **bool** |  | [optional] 
**TlsOptAllowPerUser** | Pointer to **bool** |  | [optional] 
**TlsOptCommonNameRequired** | Pointer to **bool** |  | [optional] 
**TlsOptEnableDiags** | Pointer to **bool** |  | [optional] 
**TlsOptExportCertData** | Pointer to **bool** |  | [optional] 
**TlsOptNoCertRequest** | Pointer to **bool** |  | [optional] 
**TlsOptNoEmptyFragments** | Pointer to **bool** |  | [optional] 
**TlsOptNoSessionReuseRequired** | Pointer to **bool** |  | [optional] 
**TlsOptStdenvvars** | Pointer to **bool** |  | [optional] 
**TlsOptDnsNameRequired** | Pointer to **bool** |  | [optional] 
**TlsOptIpAddressRequired** | Pointer to **bool** |  | [optional] 
**SsltlsCertificate** | Pointer to **NullableInt32** |  | [optional] 
**Options** | Pointer to **string** |  | [optional] 

## Methods

### NewFtpUpdate0

`func NewFtpUpdate0() *FtpUpdate0`

NewFtpUpdate0 instantiates a new FtpUpdate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFtpUpdate0WithDefaults

`func NewFtpUpdate0WithDefaults() *FtpUpdate0`

NewFtpUpdate0WithDefaults instantiates a new FtpUpdate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *FtpUpdate0) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *FtpUpdate0) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *FtpUpdate0) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *FtpUpdate0) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetClients

`func (o *FtpUpdate0) GetClients() int32`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *FtpUpdate0) GetClientsOk() (*int32, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *FtpUpdate0) SetClients(v int32)`

SetClients sets Clients field to given value.

### HasClients

`func (o *FtpUpdate0) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetIpconnections

`func (o *FtpUpdate0) GetIpconnections() int32`

GetIpconnections returns the Ipconnections field if non-nil, zero value otherwise.

### GetIpconnectionsOk

`func (o *FtpUpdate0) GetIpconnectionsOk() (*int32, bool)`

GetIpconnectionsOk returns a tuple with the Ipconnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconnections

`func (o *FtpUpdate0) SetIpconnections(v int32)`

SetIpconnections sets Ipconnections field to given value.

### HasIpconnections

`func (o *FtpUpdate0) HasIpconnections() bool`

HasIpconnections returns a boolean if a field has been set.

### GetLoginattempt

`func (o *FtpUpdate0) GetLoginattempt() int32`

GetLoginattempt returns the Loginattempt field if non-nil, zero value otherwise.

### GetLoginattemptOk

`func (o *FtpUpdate0) GetLoginattemptOk() (*int32, bool)`

GetLoginattemptOk returns a tuple with the Loginattempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginattempt

`func (o *FtpUpdate0) SetLoginattempt(v int32)`

SetLoginattempt sets Loginattempt field to given value.

### HasLoginattempt

`func (o *FtpUpdate0) HasLoginattempt() bool`

HasLoginattempt returns a boolean if a field has been set.

### GetTimeout

`func (o *FtpUpdate0) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *FtpUpdate0) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *FtpUpdate0) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *FtpUpdate0) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetRootlogin

`func (o *FtpUpdate0) GetRootlogin() bool`

GetRootlogin returns the Rootlogin field if non-nil, zero value otherwise.

### GetRootloginOk

`func (o *FtpUpdate0) GetRootloginOk() (*bool, bool)`

GetRootloginOk returns a tuple with the Rootlogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootlogin

`func (o *FtpUpdate0) SetRootlogin(v bool)`

SetRootlogin sets Rootlogin field to given value.

### HasRootlogin

`func (o *FtpUpdate0) HasRootlogin() bool`

HasRootlogin returns a boolean if a field has been set.

### GetOnlyanonymous

`func (o *FtpUpdate0) GetOnlyanonymous() bool`

GetOnlyanonymous returns the Onlyanonymous field if non-nil, zero value otherwise.

### GetOnlyanonymousOk

`func (o *FtpUpdate0) GetOnlyanonymousOk() (*bool, bool)`

GetOnlyanonymousOk returns a tuple with the Onlyanonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyanonymous

`func (o *FtpUpdate0) SetOnlyanonymous(v bool)`

SetOnlyanonymous sets Onlyanonymous field to given value.

### HasOnlyanonymous

`func (o *FtpUpdate0) HasOnlyanonymous() bool`

HasOnlyanonymous returns a boolean if a field has been set.

### GetAnonpath

`func (o *FtpUpdate0) GetAnonpath() string`

GetAnonpath returns the Anonpath field if non-nil, zero value otherwise.

### GetAnonpathOk

`func (o *FtpUpdate0) GetAnonpathOk() (*string, bool)`

GetAnonpathOk returns a tuple with the Anonpath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonpath

`func (o *FtpUpdate0) SetAnonpath(v string)`

SetAnonpath sets Anonpath field to given value.

### HasAnonpath

`func (o *FtpUpdate0) HasAnonpath() bool`

HasAnonpath returns a boolean if a field has been set.

### SetAnonpathNil

`func (o *FtpUpdate0) SetAnonpathNil(b bool)`

 SetAnonpathNil sets the value for Anonpath to be an explicit nil

### UnsetAnonpath
`func (o *FtpUpdate0) UnsetAnonpath()`

UnsetAnonpath ensures that no value is present for Anonpath, not even an explicit nil
### GetOnlylocal

`func (o *FtpUpdate0) GetOnlylocal() bool`

GetOnlylocal returns the Onlylocal field if non-nil, zero value otherwise.

### GetOnlylocalOk

`func (o *FtpUpdate0) GetOnlylocalOk() (*bool, bool)`

GetOnlylocalOk returns a tuple with the Onlylocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlylocal

`func (o *FtpUpdate0) SetOnlylocal(v bool)`

SetOnlylocal sets Onlylocal field to given value.

### HasOnlylocal

`func (o *FtpUpdate0) HasOnlylocal() bool`

HasOnlylocal returns a boolean if a field has been set.

### GetBanner

`func (o *FtpUpdate0) GetBanner() string`

GetBanner returns the Banner field if non-nil, zero value otherwise.

### GetBannerOk

`func (o *FtpUpdate0) GetBannerOk() (*string, bool)`

GetBannerOk returns a tuple with the Banner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanner

`func (o *FtpUpdate0) SetBanner(v string)`

SetBanner sets Banner field to given value.

### HasBanner

`func (o *FtpUpdate0) HasBanner() bool`

HasBanner returns a boolean if a field has been set.

### GetFilemask

`func (o *FtpUpdate0) GetFilemask() string`

GetFilemask returns the Filemask field if non-nil, zero value otherwise.

### GetFilemaskOk

`func (o *FtpUpdate0) GetFilemaskOk() (*string, bool)`

GetFilemaskOk returns a tuple with the Filemask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilemask

`func (o *FtpUpdate0) SetFilemask(v string)`

SetFilemask sets Filemask field to given value.

### HasFilemask

`func (o *FtpUpdate0) HasFilemask() bool`

HasFilemask returns a boolean if a field has been set.

### GetDirmask

`func (o *FtpUpdate0) GetDirmask() string`

GetDirmask returns the Dirmask field if non-nil, zero value otherwise.

### GetDirmaskOk

`func (o *FtpUpdate0) GetDirmaskOk() (*string, bool)`

GetDirmaskOk returns a tuple with the Dirmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirmask

`func (o *FtpUpdate0) SetDirmask(v string)`

SetDirmask sets Dirmask field to given value.

### HasDirmask

`func (o *FtpUpdate0) HasDirmask() bool`

HasDirmask returns a boolean if a field has been set.

### GetFxp

`func (o *FtpUpdate0) GetFxp() bool`

GetFxp returns the Fxp field if non-nil, zero value otherwise.

### GetFxpOk

`func (o *FtpUpdate0) GetFxpOk() (*bool, bool)`

GetFxpOk returns a tuple with the Fxp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFxp

`func (o *FtpUpdate0) SetFxp(v bool)`

SetFxp sets Fxp field to given value.

### HasFxp

`func (o *FtpUpdate0) HasFxp() bool`

HasFxp returns a boolean if a field has been set.

### GetResume

`func (o *FtpUpdate0) GetResume() bool`

GetResume returns the Resume field if non-nil, zero value otherwise.

### GetResumeOk

`func (o *FtpUpdate0) GetResumeOk() (*bool, bool)`

GetResumeOk returns a tuple with the Resume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResume

`func (o *FtpUpdate0) SetResume(v bool)`

SetResume sets Resume field to given value.

### HasResume

`func (o *FtpUpdate0) HasResume() bool`

HasResume returns a boolean if a field has been set.

### GetDefaultroot

`func (o *FtpUpdate0) GetDefaultroot() bool`

GetDefaultroot returns the Defaultroot field if non-nil, zero value otherwise.

### GetDefaultrootOk

`func (o *FtpUpdate0) GetDefaultrootOk() (*bool, bool)`

GetDefaultrootOk returns a tuple with the Defaultroot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultroot

`func (o *FtpUpdate0) SetDefaultroot(v bool)`

SetDefaultroot sets Defaultroot field to given value.

### HasDefaultroot

`func (o *FtpUpdate0) HasDefaultroot() bool`

HasDefaultroot returns a boolean if a field has been set.

### GetIdent

`func (o *FtpUpdate0) GetIdent() bool`

GetIdent returns the Ident field if non-nil, zero value otherwise.

### GetIdentOk

`func (o *FtpUpdate0) GetIdentOk() (*bool, bool)`

GetIdentOk returns a tuple with the Ident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdent

`func (o *FtpUpdate0) SetIdent(v bool)`

SetIdent sets Ident field to given value.

### HasIdent

`func (o *FtpUpdate0) HasIdent() bool`

HasIdent returns a boolean if a field has been set.

### GetReversedns

`func (o *FtpUpdate0) GetReversedns() bool`

GetReversedns returns the Reversedns field if non-nil, zero value otherwise.

### GetReversednsOk

`func (o *FtpUpdate0) GetReversednsOk() (*bool, bool)`

GetReversednsOk returns a tuple with the Reversedns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReversedns

`func (o *FtpUpdate0) SetReversedns(v bool)`

SetReversedns sets Reversedns field to given value.

### HasReversedns

`func (o *FtpUpdate0) HasReversedns() bool`

HasReversedns returns a boolean if a field has been set.

### GetMasqaddress

`func (o *FtpUpdate0) GetMasqaddress() string`

GetMasqaddress returns the Masqaddress field if non-nil, zero value otherwise.

### GetMasqaddressOk

`func (o *FtpUpdate0) GetMasqaddressOk() (*string, bool)`

GetMasqaddressOk returns a tuple with the Masqaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasqaddress

`func (o *FtpUpdate0) SetMasqaddress(v string)`

SetMasqaddress sets Masqaddress field to given value.

### HasMasqaddress

`func (o *FtpUpdate0) HasMasqaddress() bool`

HasMasqaddress returns a boolean if a field has been set.

### GetPassiveportsmin

`func (o *FtpUpdate0) GetPassiveportsmin() int32`

GetPassiveportsmin returns the Passiveportsmin field if non-nil, zero value otherwise.

### GetPassiveportsminOk

`func (o *FtpUpdate0) GetPassiveportsminOk() (*int32, bool)`

GetPassiveportsminOk returns a tuple with the Passiveportsmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassiveportsmin

`func (o *FtpUpdate0) SetPassiveportsmin(v int32)`

SetPassiveportsmin sets Passiveportsmin field to given value.

### HasPassiveportsmin

`func (o *FtpUpdate0) HasPassiveportsmin() bool`

HasPassiveportsmin returns a boolean if a field has been set.

### GetPassiveportsmax

`func (o *FtpUpdate0) GetPassiveportsmax() int32`

GetPassiveportsmax returns the Passiveportsmax field if non-nil, zero value otherwise.

### GetPassiveportsmaxOk

`func (o *FtpUpdate0) GetPassiveportsmaxOk() (*int32, bool)`

GetPassiveportsmaxOk returns a tuple with the Passiveportsmax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassiveportsmax

`func (o *FtpUpdate0) SetPassiveportsmax(v int32)`

SetPassiveportsmax sets Passiveportsmax field to given value.

### HasPassiveportsmax

`func (o *FtpUpdate0) HasPassiveportsmax() bool`

HasPassiveportsmax returns a boolean if a field has been set.

### GetLocaluserbw

`func (o *FtpUpdate0) GetLocaluserbw() int32`

GetLocaluserbw returns the Localuserbw field if non-nil, zero value otherwise.

### GetLocaluserbwOk

`func (o *FtpUpdate0) GetLocaluserbwOk() (*int32, bool)`

GetLocaluserbwOk returns a tuple with the Localuserbw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocaluserbw

`func (o *FtpUpdate0) SetLocaluserbw(v int32)`

SetLocaluserbw sets Localuserbw field to given value.

### HasLocaluserbw

`func (o *FtpUpdate0) HasLocaluserbw() bool`

HasLocaluserbw returns a boolean if a field has been set.

### GetLocaluserdlbw

`func (o *FtpUpdate0) GetLocaluserdlbw() int32`

GetLocaluserdlbw returns the Localuserdlbw field if non-nil, zero value otherwise.

### GetLocaluserdlbwOk

`func (o *FtpUpdate0) GetLocaluserdlbwOk() (*int32, bool)`

GetLocaluserdlbwOk returns a tuple with the Localuserdlbw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocaluserdlbw

`func (o *FtpUpdate0) SetLocaluserdlbw(v int32)`

SetLocaluserdlbw sets Localuserdlbw field to given value.

### HasLocaluserdlbw

`func (o *FtpUpdate0) HasLocaluserdlbw() bool`

HasLocaluserdlbw returns a boolean if a field has been set.

### GetAnonuserbw

`func (o *FtpUpdate0) GetAnonuserbw() int32`

GetAnonuserbw returns the Anonuserbw field if non-nil, zero value otherwise.

### GetAnonuserbwOk

`func (o *FtpUpdate0) GetAnonuserbwOk() (*int32, bool)`

GetAnonuserbwOk returns a tuple with the Anonuserbw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonuserbw

`func (o *FtpUpdate0) SetAnonuserbw(v int32)`

SetAnonuserbw sets Anonuserbw field to given value.

### HasAnonuserbw

`func (o *FtpUpdate0) HasAnonuserbw() bool`

HasAnonuserbw returns a boolean if a field has been set.

### GetAnonuserdlbw

`func (o *FtpUpdate0) GetAnonuserdlbw() int32`

GetAnonuserdlbw returns the Anonuserdlbw field if non-nil, zero value otherwise.

### GetAnonuserdlbwOk

`func (o *FtpUpdate0) GetAnonuserdlbwOk() (*int32, bool)`

GetAnonuserdlbwOk returns a tuple with the Anonuserdlbw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonuserdlbw

`func (o *FtpUpdate0) SetAnonuserdlbw(v int32)`

SetAnonuserdlbw sets Anonuserdlbw field to given value.

### HasAnonuserdlbw

`func (o *FtpUpdate0) HasAnonuserdlbw() bool`

HasAnonuserdlbw returns a boolean if a field has been set.

### GetTls

`func (o *FtpUpdate0) GetTls() bool`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *FtpUpdate0) GetTlsOk() (*bool, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *FtpUpdate0) SetTls(v bool)`

SetTls sets Tls field to given value.

### HasTls

`func (o *FtpUpdate0) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetTlsPolicy

`func (o *FtpUpdate0) GetTlsPolicy() string`

GetTlsPolicy returns the TlsPolicy field if non-nil, zero value otherwise.

### GetTlsPolicyOk

`func (o *FtpUpdate0) GetTlsPolicyOk() (*string, bool)`

GetTlsPolicyOk returns a tuple with the TlsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsPolicy

`func (o *FtpUpdate0) SetTlsPolicy(v string)`

SetTlsPolicy sets TlsPolicy field to given value.

### HasTlsPolicy

`func (o *FtpUpdate0) HasTlsPolicy() bool`

HasTlsPolicy returns a boolean if a field has been set.

### GetTlsOptAllowClientRenegotiations

`func (o *FtpUpdate0) GetTlsOptAllowClientRenegotiations() bool`

GetTlsOptAllowClientRenegotiations returns the TlsOptAllowClientRenegotiations field if non-nil, zero value otherwise.

### GetTlsOptAllowClientRenegotiationsOk

`func (o *FtpUpdate0) GetTlsOptAllowClientRenegotiationsOk() (*bool, bool)`

GetTlsOptAllowClientRenegotiationsOk returns a tuple with the TlsOptAllowClientRenegotiations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsOptAllowClientRenegotiations

`func (o *FtpUpdate0) SetTlsOptAllowClientRenegotiations(v bool)`

SetTlsOptAllowClientRenegotiations sets TlsOptAllowClientRenegotiations field to given value.

### HasTlsOptAllowClientRenegotiations

`func (o *FtpUpdate0) HasTlsOptAllowClientRenegotiations() bool`

HasTlsOptAllowClientRenegotiations returns a boolean if a field has been set.

### GetTlsOptAllowDotLogin

`func (o *FtpUpdate0) GetTlsOptAllowDotLogin() bool`

GetTlsOptAllowDotLogin returns the TlsOptAllowDotLogin field if non-nil, zero value otherwise.

### GetTlsOptAllowDotLoginOk

`func (o *FtpUpdate0) GetTlsOptAllowDotLoginOk() (*bool, bool)`

GetTlsOptAllowDotLoginOk returns a tuple with the TlsOptAllowDotLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsOptAllowDotLogin

`func (o *FtpUpdate0) SetTlsOptAllowDotLogin(v bool)`

SetTlsOptAllowDotLogin sets TlsOptAllowDotLogin field to given value.

### HasTlsOptAllowDotLogin

`func (o *FtpUpdate0) HasTlsOptAllowDotLogin() bool`

HasTlsOptAllowDotLogin returns a boolean if a field has been set.

### GetTlsOptAllowPerUser

`func (o *FtpUpdate0) GetTlsOptAllowPerUser() bool`

GetTlsOptAllowPerUser returns the TlsOptAllowPerUser field if non-nil, zero value otherwise.

### GetTlsOptAllowPerUserOk

`func (o *FtpUpdate0) GetTlsOptAllowPerUserOk() (*bool, bool)`

GetTlsOptAllowPerUserOk returns a tuple with the TlsOptAllowPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsOptAllowPerUser

`func (o *FtpUpdate0) SetTlsOptAllowPerUser(v bool)`

SetTlsOptAllowPerUser sets TlsOptAllowPerUser field to given value.

### HasTlsOptAllowPerUser

`func (o *FtpUpdate0) HasTlsOptAllowPerUser() bool`

HasTlsOptAllowPerUser returns a boolean if a field has been set.

### GetTlsOptCommonNameRequired

`func (o *FtpUpdate0) GetTlsOptCommonNameRequired() bool`

GetTlsOptCommonNameRequired returns the TlsOptCommonNameRequired field if non-nil, zero value otherwise.

### GetTlsOptCommonNameRequiredOk

`func (o *FtpUpdate0) GetTlsOptCommonNameRequiredOk() (*bool, bool)`

GetTlsOptCommonNameRequiredOk returns a tuple with the TlsOptCommonNameRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsOptCommonNameRequired

`func (o *FtpUpdate0) SetTlsOptCommonNameRequired(v bool)`

SetTlsOptCommonNameRequired sets TlsOptCommonNameRequired field to given value.

### HasTlsOptCommonNameRequired

`func (o *FtpUpdate0) HasTlsOptCommonNameRequired() bool`

HasTlsOptCommonNameRequired returns a boolean if a field has been set.

### GetTlsOptEnableDiags

`func (o *FtpUpdate0) GetTlsOptEnableDiags() bool`

GetTlsOptEnableDiags returns the TlsOptEnableDiags field if non-nil, zero value otherwise.

### GetTlsOptEnableDiagsOk

`func (o *FtpUpdate0) GetTlsOptEnableDiagsOk() (*bool, bool)`

GetTlsOptEnableDiagsOk returns a tuple with the TlsOptEnableDiags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsOptEnableDiags

`func (o *FtpUpdate0) SetTlsOptEnableDiags(v bool)`

SetTlsOptEnableDiags sets TlsOptEnableDiags field to given value.

### HasTlsOptEnableDiags

`func (o *FtpUpdate0) HasTlsOptEnableDiags() bool`

HasTlsOptEnableDiags returns a boolean if a field has been set.

### GetTlsOptExportCertData

`func (o *FtpUpdate0) GetTlsOptExportCertData() bool`

GetTlsOptExportCertData returns the TlsOptExportCertData field if non-nil, zero value otherwise.

### GetTlsOptExportCertDataOk

`func (o *FtpUpdate0) GetTlsOptExportCertDataOk() (*bool, bool)`

GetTlsOptExportCertDataOk returns a tuple with the TlsOptExportCertData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsOptExportCertData

`func (o *FtpUpdate0) SetTlsOptExportCertData(v bool)`

SetTlsOptExportCertData sets TlsOptExportCertData field to given value.

### HasTlsOptExportCertData

`func (o *FtpUpdate0) HasTlsOptExportCertData() bool`

HasTlsOptExportCertData returns a boolean if a field has been set.

### GetTlsOptNoCertRequest

`func (o *FtpUpdate0) GetTlsOptNoCertRequest() bool`

GetTlsOptNoCertRequest returns the TlsOptNoCertRequest field if non-nil, zero value otherwise.

### GetTlsOptNoCertRequestOk

`func (o *FtpUpdate0) GetTlsOptNoCertRequestOk() (*bool, bool)`

GetTlsOptNoCertRequestOk returns a tuple with the TlsOptNoCertRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsOptNoCertRequest

`func (o *FtpUpdate0) SetTlsOptNoCertRequest(v bool)`

SetTlsOptNoCertRequest sets TlsOptNoCertRequest field to given value.

### HasTlsOptNoCertRequest

`func (o *FtpUpdate0) HasTlsOptNoCertRequest() bool`

HasTlsOptNoCertRequest returns a boolean if a field has been set.

### GetTlsOptNoEmptyFragments

`func (o *FtpUpdate0) GetTlsOptNoEmptyFragments() bool`

GetTlsOptNoEmptyFragments returns the TlsOptNoEmptyFragments field if non-nil, zero value otherwise.

### GetTlsOptNoEmptyFragmentsOk

`func (o *FtpUpdate0) GetTlsOptNoEmptyFragmentsOk() (*bool, bool)`

GetTlsOptNoEmptyFragmentsOk returns a tuple with the TlsOptNoEmptyFragments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsOptNoEmptyFragments

`func (o *FtpUpdate0) SetTlsOptNoEmptyFragments(v bool)`

SetTlsOptNoEmptyFragments sets TlsOptNoEmptyFragments field to given value.

### HasTlsOptNoEmptyFragments

`func (o *FtpUpdate0) HasTlsOptNoEmptyFragments() bool`

HasTlsOptNoEmptyFragments returns a boolean if a field has been set.

### GetTlsOptNoSessionReuseRequired

`func (o *FtpUpdate0) GetTlsOptNoSessionReuseRequired() bool`

GetTlsOptNoSessionReuseRequired returns the TlsOptNoSessionReuseRequired field if non-nil, zero value otherwise.

### GetTlsOptNoSessionReuseRequiredOk

`func (o *FtpUpdate0) GetTlsOptNoSessionReuseRequiredOk() (*bool, bool)`

GetTlsOptNoSessionReuseRequiredOk returns a tuple with the TlsOptNoSessionReuseRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsOptNoSessionReuseRequired

`func (o *FtpUpdate0) SetTlsOptNoSessionReuseRequired(v bool)`

SetTlsOptNoSessionReuseRequired sets TlsOptNoSessionReuseRequired field to given value.

### HasTlsOptNoSessionReuseRequired

`func (o *FtpUpdate0) HasTlsOptNoSessionReuseRequired() bool`

HasTlsOptNoSessionReuseRequired returns a boolean if a field has been set.

### GetTlsOptStdenvvars

`func (o *FtpUpdate0) GetTlsOptStdenvvars() bool`

GetTlsOptStdenvvars returns the TlsOptStdenvvars field if non-nil, zero value otherwise.

### GetTlsOptStdenvvarsOk

`func (o *FtpUpdate0) GetTlsOptStdenvvarsOk() (*bool, bool)`

GetTlsOptStdenvvarsOk returns a tuple with the TlsOptStdenvvars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsOptStdenvvars

`func (o *FtpUpdate0) SetTlsOptStdenvvars(v bool)`

SetTlsOptStdenvvars sets TlsOptStdenvvars field to given value.

### HasTlsOptStdenvvars

`func (o *FtpUpdate0) HasTlsOptStdenvvars() bool`

HasTlsOptStdenvvars returns a boolean if a field has been set.

### GetTlsOptDnsNameRequired

`func (o *FtpUpdate0) GetTlsOptDnsNameRequired() bool`

GetTlsOptDnsNameRequired returns the TlsOptDnsNameRequired field if non-nil, zero value otherwise.

### GetTlsOptDnsNameRequiredOk

`func (o *FtpUpdate0) GetTlsOptDnsNameRequiredOk() (*bool, bool)`

GetTlsOptDnsNameRequiredOk returns a tuple with the TlsOptDnsNameRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsOptDnsNameRequired

`func (o *FtpUpdate0) SetTlsOptDnsNameRequired(v bool)`

SetTlsOptDnsNameRequired sets TlsOptDnsNameRequired field to given value.

### HasTlsOptDnsNameRequired

`func (o *FtpUpdate0) HasTlsOptDnsNameRequired() bool`

HasTlsOptDnsNameRequired returns a boolean if a field has been set.

### GetTlsOptIpAddressRequired

`func (o *FtpUpdate0) GetTlsOptIpAddressRequired() bool`

GetTlsOptIpAddressRequired returns the TlsOptIpAddressRequired field if non-nil, zero value otherwise.

### GetTlsOptIpAddressRequiredOk

`func (o *FtpUpdate0) GetTlsOptIpAddressRequiredOk() (*bool, bool)`

GetTlsOptIpAddressRequiredOk returns a tuple with the TlsOptIpAddressRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsOptIpAddressRequired

`func (o *FtpUpdate0) SetTlsOptIpAddressRequired(v bool)`

SetTlsOptIpAddressRequired sets TlsOptIpAddressRequired field to given value.

### HasTlsOptIpAddressRequired

`func (o *FtpUpdate0) HasTlsOptIpAddressRequired() bool`

HasTlsOptIpAddressRequired returns a boolean if a field has been set.

### GetSsltlsCertificate

`func (o *FtpUpdate0) GetSsltlsCertificate() int32`

GetSsltlsCertificate returns the SsltlsCertificate field if non-nil, zero value otherwise.

### GetSsltlsCertificateOk

`func (o *FtpUpdate0) GetSsltlsCertificateOk() (*int32, bool)`

GetSsltlsCertificateOk returns a tuple with the SsltlsCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsltlsCertificate

`func (o *FtpUpdate0) SetSsltlsCertificate(v int32)`

SetSsltlsCertificate sets SsltlsCertificate field to given value.

### HasSsltlsCertificate

`func (o *FtpUpdate0) HasSsltlsCertificate() bool`

HasSsltlsCertificate returns a boolean if a field has been set.

### SetSsltlsCertificateNil

`func (o *FtpUpdate0) SetSsltlsCertificateNil(b bool)`

 SetSsltlsCertificateNil sets the value for SsltlsCertificate to be an explicit nil

### UnsetSsltlsCertificate
`func (o *FtpUpdate0) UnsetSsltlsCertificate()`

UnsetSsltlsCertificate ensures that no value is present for SsltlsCertificate, not even an explicit nil
### GetOptions

`func (o *FtpUpdate0) GetOptions() string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *FtpUpdate0) GetOptionsOk() (*string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *FtpUpdate0) SetOptions(v string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *FtpUpdate0) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


