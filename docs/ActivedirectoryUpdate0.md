# ActivedirectoryUpdate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domainname** | Pointer to **string** |  | [optional] 
**Bindname** | Pointer to **string** |  | [optional] 
**Bindpw** | Pointer to **string** |  | [optional] 
**VerboseLogging** | Pointer to **bool** |  | [optional] 
**UseDefaultDomain** | Pointer to **bool** |  | [optional] 
**AllowTrustedDoms** | Pointer to **bool** |  | [optional] 
**AllowDnsUpdates** | Pointer to **bool** |  | [optional] 
**DisableFreenasCache** | Pointer to **bool** |  | [optional] 
**RestrictPam** | Pointer to **bool** |  | [optional] 
**Site** | Pointer to **NullableString** |  | [optional] 
**KerberosRealm** | Pointer to **NullableInt32** |  | [optional] 
**KerberosPrincipal** | Pointer to **NullableString** |  | [optional] 
**Timeout** | Pointer to **int32** |  | [optional] 
**DnsTimeout** | Pointer to **int32** |  | [optional] 
**NssInfo** | Pointer to **NullableString** |  | [optional] 
**Createcomputer** | Pointer to **string** |  | [optional] 
**Netbiosname** | Pointer to **string** |  | [optional] 
**NetbiosnameB** | Pointer to **string** |  | [optional] 
**Netbiosalias** | Pointer to **[]interface{}** |  | [optional] 
**Enable** | Pointer to **bool** |  | [optional] 

## Methods

### NewActivedirectoryUpdate0

`func NewActivedirectoryUpdate0() *ActivedirectoryUpdate0`

NewActivedirectoryUpdate0 instantiates a new ActivedirectoryUpdate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivedirectoryUpdate0WithDefaults

`func NewActivedirectoryUpdate0WithDefaults() *ActivedirectoryUpdate0`

NewActivedirectoryUpdate0WithDefaults instantiates a new ActivedirectoryUpdate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainname

`func (o *ActivedirectoryUpdate0) GetDomainname() string`

GetDomainname returns the Domainname field if non-nil, zero value otherwise.

### GetDomainnameOk

`func (o *ActivedirectoryUpdate0) GetDomainnameOk() (*string, bool)`

GetDomainnameOk returns a tuple with the Domainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainname

`func (o *ActivedirectoryUpdate0) SetDomainname(v string)`

SetDomainname sets Domainname field to given value.

### HasDomainname

`func (o *ActivedirectoryUpdate0) HasDomainname() bool`

HasDomainname returns a boolean if a field has been set.

### GetBindname

`func (o *ActivedirectoryUpdate0) GetBindname() string`

GetBindname returns the Bindname field if non-nil, zero value otherwise.

### GetBindnameOk

`func (o *ActivedirectoryUpdate0) GetBindnameOk() (*string, bool)`

GetBindnameOk returns a tuple with the Bindname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindname

`func (o *ActivedirectoryUpdate0) SetBindname(v string)`

SetBindname sets Bindname field to given value.

### HasBindname

`func (o *ActivedirectoryUpdate0) HasBindname() bool`

HasBindname returns a boolean if a field has been set.

### GetBindpw

`func (o *ActivedirectoryUpdate0) GetBindpw() string`

GetBindpw returns the Bindpw field if non-nil, zero value otherwise.

### GetBindpwOk

`func (o *ActivedirectoryUpdate0) GetBindpwOk() (*string, bool)`

GetBindpwOk returns a tuple with the Bindpw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindpw

`func (o *ActivedirectoryUpdate0) SetBindpw(v string)`

SetBindpw sets Bindpw field to given value.

### HasBindpw

`func (o *ActivedirectoryUpdate0) HasBindpw() bool`

HasBindpw returns a boolean if a field has been set.

### GetVerboseLogging

`func (o *ActivedirectoryUpdate0) GetVerboseLogging() bool`

GetVerboseLogging returns the VerboseLogging field if non-nil, zero value otherwise.

### GetVerboseLoggingOk

`func (o *ActivedirectoryUpdate0) GetVerboseLoggingOk() (*bool, bool)`

GetVerboseLoggingOk returns a tuple with the VerboseLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseLogging

`func (o *ActivedirectoryUpdate0) SetVerboseLogging(v bool)`

SetVerboseLogging sets VerboseLogging field to given value.

### HasVerboseLogging

`func (o *ActivedirectoryUpdate0) HasVerboseLogging() bool`

HasVerboseLogging returns a boolean if a field has been set.

### GetUseDefaultDomain

`func (o *ActivedirectoryUpdate0) GetUseDefaultDomain() bool`

GetUseDefaultDomain returns the UseDefaultDomain field if non-nil, zero value otherwise.

### GetUseDefaultDomainOk

`func (o *ActivedirectoryUpdate0) GetUseDefaultDomainOk() (*bool, bool)`

GetUseDefaultDomainOk returns a tuple with the UseDefaultDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDefaultDomain

`func (o *ActivedirectoryUpdate0) SetUseDefaultDomain(v bool)`

SetUseDefaultDomain sets UseDefaultDomain field to given value.

### HasUseDefaultDomain

`func (o *ActivedirectoryUpdate0) HasUseDefaultDomain() bool`

HasUseDefaultDomain returns a boolean if a field has been set.

### GetAllowTrustedDoms

`func (o *ActivedirectoryUpdate0) GetAllowTrustedDoms() bool`

GetAllowTrustedDoms returns the AllowTrustedDoms field if non-nil, zero value otherwise.

### GetAllowTrustedDomsOk

`func (o *ActivedirectoryUpdate0) GetAllowTrustedDomsOk() (*bool, bool)`

GetAllowTrustedDomsOk returns a tuple with the AllowTrustedDoms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTrustedDoms

`func (o *ActivedirectoryUpdate0) SetAllowTrustedDoms(v bool)`

SetAllowTrustedDoms sets AllowTrustedDoms field to given value.

### HasAllowTrustedDoms

`func (o *ActivedirectoryUpdate0) HasAllowTrustedDoms() bool`

HasAllowTrustedDoms returns a boolean if a field has been set.

### GetAllowDnsUpdates

`func (o *ActivedirectoryUpdate0) GetAllowDnsUpdates() bool`

GetAllowDnsUpdates returns the AllowDnsUpdates field if non-nil, zero value otherwise.

### GetAllowDnsUpdatesOk

`func (o *ActivedirectoryUpdate0) GetAllowDnsUpdatesOk() (*bool, bool)`

GetAllowDnsUpdatesOk returns a tuple with the AllowDnsUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDnsUpdates

`func (o *ActivedirectoryUpdate0) SetAllowDnsUpdates(v bool)`

SetAllowDnsUpdates sets AllowDnsUpdates field to given value.

### HasAllowDnsUpdates

`func (o *ActivedirectoryUpdate0) HasAllowDnsUpdates() bool`

HasAllowDnsUpdates returns a boolean if a field has been set.

### GetDisableFreenasCache

`func (o *ActivedirectoryUpdate0) GetDisableFreenasCache() bool`

GetDisableFreenasCache returns the DisableFreenasCache field if non-nil, zero value otherwise.

### GetDisableFreenasCacheOk

`func (o *ActivedirectoryUpdate0) GetDisableFreenasCacheOk() (*bool, bool)`

GetDisableFreenasCacheOk returns a tuple with the DisableFreenasCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableFreenasCache

`func (o *ActivedirectoryUpdate0) SetDisableFreenasCache(v bool)`

SetDisableFreenasCache sets DisableFreenasCache field to given value.

### HasDisableFreenasCache

`func (o *ActivedirectoryUpdate0) HasDisableFreenasCache() bool`

HasDisableFreenasCache returns a boolean if a field has been set.

### GetRestrictPam

`func (o *ActivedirectoryUpdate0) GetRestrictPam() bool`

GetRestrictPam returns the RestrictPam field if non-nil, zero value otherwise.

### GetRestrictPamOk

`func (o *ActivedirectoryUpdate0) GetRestrictPamOk() (*bool, bool)`

GetRestrictPamOk returns a tuple with the RestrictPam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictPam

`func (o *ActivedirectoryUpdate0) SetRestrictPam(v bool)`

SetRestrictPam sets RestrictPam field to given value.

### HasRestrictPam

`func (o *ActivedirectoryUpdate0) HasRestrictPam() bool`

HasRestrictPam returns a boolean if a field has been set.

### GetSite

`func (o *ActivedirectoryUpdate0) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *ActivedirectoryUpdate0) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *ActivedirectoryUpdate0) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *ActivedirectoryUpdate0) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *ActivedirectoryUpdate0) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *ActivedirectoryUpdate0) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetKerberosRealm

`func (o *ActivedirectoryUpdate0) GetKerberosRealm() int32`

GetKerberosRealm returns the KerberosRealm field if non-nil, zero value otherwise.

### GetKerberosRealmOk

`func (o *ActivedirectoryUpdate0) GetKerberosRealmOk() (*int32, bool)`

GetKerberosRealmOk returns a tuple with the KerberosRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosRealm

`func (o *ActivedirectoryUpdate0) SetKerberosRealm(v int32)`

SetKerberosRealm sets KerberosRealm field to given value.

### HasKerberosRealm

`func (o *ActivedirectoryUpdate0) HasKerberosRealm() bool`

HasKerberosRealm returns a boolean if a field has been set.

### SetKerberosRealmNil

`func (o *ActivedirectoryUpdate0) SetKerberosRealmNil(b bool)`

 SetKerberosRealmNil sets the value for KerberosRealm to be an explicit nil

### UnsetKerberosRealm
`func (o *ActivedirectoryUpdate0) UnsetKerberosRealm()`

UnsetKerberosRealm ensures that no value is present for KerberosRealm, not even an explicit nil
### GetKerberosPrincipal

`func (o *ActivedirectoryUpdate0) GetKerberosPrincipal() string`

GetKerberosPrincipal returns the KerberosPrincipal field if non-nil, zero value otherwise.

### GetKerberosPrincipalOk

`func (o *ActivedirectoryUpdate0) GetKerberosPrincipalOk() (*string, bool)`

GetKerberosPrincipalOk returns a tuple with the KerberosPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosPrincipal

`func (o *ActivedirectoryUpdate0) SetKerberosPrincipal(v string)`

SetKerberosPrincipal sets KerberosPrincipal field to given value.

### HasKerberosPrincipal

`func (o *ActivedirectoryUpdate0) HasKerberosPrincipal() bool`

HasKerberosPrincipal returns a boolean if a field has been set.

### SetKerberosPrincipalNil

`func (o *ActivedirectoryUpdate0) SetKerberosPrincipalNil(b bool)`

 SetKerberosPrincipalNil sets the value for KerberosPrincipal to be an explicit nil

### UnsetKerberosPrincipal
`func (o *ActivedirectoryUpdate0) UnsetKerberosPrincipal()`

UnsetKerberosPrincipal ensures that no value is present for KerberosPrincipal, not even an explicit nil
### GetTimeout

`func (o *ActivedirectoryUpdate0) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ActivedirectoryUpdate0) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ActivedirectoryUpdate0) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ActivedirectoryUpdate0) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetDnsTimeout

`func (o *ActivedirectoryUpdate0) GetDnsTimeout() int32`

GetDnsTimeout returns the DnsTimeout field if non-nil, zero value otherwise.

### GetDnsTimeoutOk

`func (o *ActivedirectoryUpdate0) GetDnsTimeoutOk() (*int32, bool)`

GetDnsTimeoutOk returns a tuple with the DnsTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsTimeout

`func (o *ActivedirectoryUpdate0) SetDnsTimeout(v int32)`

SetDnsTimeout sets DnsTimeout field to given value.

### HasDnsTimeout

`func (o *ActivedirectoryUpdate0) HasDnsTimeout() bool`

HasDnsTimeout returns a boolean if a field has been set.

### GetNssInfo

`func (o *ActivedirectoryUpdate0) GetNssInfo() string`

GetNssInfo returns the NssInfo field if non-nil, zero value otherwise.

### GetNssInfoOk

`func (o *ActivedirectoryUpdate0) GetNssInfoOk() (*string, bool)`

GetNssInfoOk returns a tuple with the NssInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNssInfo

`func (o *ActivedirectoryUpdate0) SetNssInfo(v string)`

SetNssInfo sets NssInfo field to given value.

### HasNssInfo

`func (o *ActivedirectoryUpdate0) HasNssInfo() bool`

HasNssInfo returns a boolean if a field has been set.

### SetNssInfoNil

`func (o *ActivedirectoryUpdate0) SetNssInfoNil(b bool)`

 SetNssInfoNil sets the value for NssInfo to be an explicit nil

### UnsetNssInfo
`func (o *ActivedirectoryUpdate0) UnsetNssInfo()`

UnsetNssInfo ensures that no value is present for NssInfo, not even an explicit nil
### GetCreatecomputer

`func (o *ActivedirectoryUpdate0) GetCreatecomputer() string`

GetCreatecomputer returns the Createcomputer field if non-nil, zero value otherwise.

### GetCreatecomputerOk

`func (o *ActivedirectoryUpdate0) GetCreatecomputerOk() (*string, bool)`

GetCreatecomputerOk returns a tuple with the Createcomputer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatecomputer

`func (o *ActivedirectoryUpdate0) SetCreatecomputer(v string)`

SetCreatecomputer sets Createcomputer field to given value.

### HasCreatecomputer

`func (o *ActivedirectoryUpdate0) HasCreatecomputer() bool`

HasCreatecomputer returns a boolean if a field has been set.

### GetNetbiosname

`func (o *ActivedirectoryUpdate0) GetNetbiosname() string`

GetNetbiosname returns the Netbiosname field if non-nil, zero value otherwise.

### GetNetbiosnameOk

`func (o *ActivedirectoryUpdate0) GetNetbiosnameOk() (*string, bool)`

GetNetbiosnameOk returns a tuple with the Netbiosname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbiosname

`func (o *ActivedirectoryUpdate0) SetNetbiosname(v string)`

SetNetbiosname sets Netbiosname field to given value.

### HasNetbiosname

`func (o *ActivedirectoryUpdate0) HasNetbiosname() bool`

HasNetbiosname returns a boolean if a field has been set.

### GetNetbiosnameB

`func (o *ActivedirectoryUpdate0) GetNetbiosnameB() string`

GetNetbiosnameB returns the NetbiosnameB field if non-nil, zero value otherwise.

### GetNetbiosnameBOk

`func (o *ActivedirectoryUpdate0) GetNetbiosnameBOk() (*string, bool)`

GetNetbiosnameBOk returns a tuple with the NetbiosnameB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbiosnameB

`func (o *ActivedirectoryUpdate0) SetNetbiosnameB(v string)`

SetNetbiosnameB sets NetbiosnameB field to given value.

### HasNetbiosnameB

`func (o *ActivedirectoryUpdate0) HasNetbiosnameB() bool`

HasNetbiosnameB returns a boolean if a field has been set.

### GetNetbiosalias

`func (o *ActivedirectoryUpdate0) GetNetbiosalias() []interface{}`

GetNetbiosalias returns the Netbiosalias field if non-nil, zero value otherwise.

### GetNetbiosaliasOk

`func (o *ActivedirectoryUpdate0) GetNetbiosaliasOk() (*[]interface{}, bool)`

GetNetbiosaliasOk returns a tuple with the Netbiosalias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbiosalias

`func (o *ActivedirectoryUpdate0) SetNetbiosalias(v []interface{})`

SetNetbiosalias sets Netbiosalias field to given value.

### HasNetbiosalias

`func (o *ActivedirectoryUpdate0) HasNetbiosalias() bool`

HasNetbiosalias returns a boolean if a field has been set.

### GetEnable

`func (o *ActivedirectoryUpdate0) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *ActivedirectoryUpdate0) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *ActivedirectoryUpdate0) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *ActivedirectoryUpdate0) HasEnable() bool`

HasEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


