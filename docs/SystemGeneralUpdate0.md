# SystemGeneralUpdate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UiCertificate** | Pointer to **NullableInt32** |  | [optional] 
**UiHttpsport** | Pointer to **int32** |  | [optional] 
**UiHttpsredirect** | Pointer to **bool** |  | [optional] 
**UiHttpsprotocols** | Pointer to **[]string** |  | [optional] 
**UiPort** | Pointer to **int32** |  | [optional] 
**UiAddress** | Pointer to **[]string** |  | [optional] 
**UiV6address** | Pointer to **[]string** |  | [optional] 
**Kbdmap** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**Sysloglevel** | Pointer to **string** |  | [optional] 
**Syslogserver** | Pointer to **string** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**CrashReporting** | Pointer to **NullableBool** |  | [optional] 
**UsageCollection** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewSystemGeneralUpdate0

`func NewSystemGeneralUpdate0() *SystemGeneralUpdate0`

NewSystemGeneralUpdate0 instantiates a new SystemGeneralUpdate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemGeneralUpdate0WithDefaults

`func NewSystemGeneralUpdate0WithDefaults() *SystemGeneralUpdate0`

NewSystemGeneralUpdate0WithDefaults instantiates a new SystemGeneralUpdate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUiCertificate

`func (o *SystemGeneralUpdate0) GetUiCertificate() int32`

GetUiCertificate returns the UiCertificate field if non-nil, zero value otherwise.

### GetUiCertificateOk

`func (o *SystemGeneralUpdate0) GetUiCertificateOk() (*int32, bool)`

GetUiCertificateOk returns a tuple with the UiCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiCertificate

`func (o *SystemGeneralUpdate0) SetUiCertificate(v int32)`

SetUiCertificate sets UiCertificate field to given value.

### HasUiCertificate

`func (o *SystemGeneralUpdate0) HasUiCertificate() bool`

HasUiCertificate returns a boolean if a field has been set.

### SetUiCertificateNil

`func (o *SystemGeneralUpdate0) SetUiCertificateNil(b bool)`

 SetUiCertificateNil sets the value for UiCertificate to be an explicit nil

### UnsetUiCertificate
`func (o *SystemGeneralUpdate0) UnsetUiCertificate()`

UnsetUiCertificate ensures that no value is present for UiCertificate, not even an explicit nil
### GetUiHttpsport

`func (o *SystemGeneralUpdate0) GetUiHttpsport() int32`

GetUiHttpsport returns the UiHttpsport field if non-nil, zero value otherwise.

### GetUiHttpsportOk

`func (o *SystemGeneralUpdate0) GetUiHttpsportOk() (*int32, bool)`

GetUiHttpsportOk returns a tuple with the UiHttpsport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiHttpsport

`func (o *SystemGeneralUpdate0) SetUiHttpsport(v int32)`

SetUiHttpsport sets UiHttpsport field to given value.

### HasUiHttpsport

`func (o *SystemGeneralUpdate0) HasUiHttpsport() bool`

HasUiHttpsport returns a boolean if a field has been set.

### GetUiHttpsredirect

`func (o *SystemGeneralUpdate0) GetUiHttpsredirect() bool`

GetUiHttpsredirect returns the UiHttpsredirect field if non-nil, zero value otherwise.

### GetUiHttpsredirectOk

`func (o *SystemGeneralUpdate0) GetUiHttpsredirectOk() (*bool, bool)`

GetUiHttpsredirectOk returns a tuple with the UiHttpsredirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiHttpsredirect

`func (o *SystemGeneralUpdate0) SetUiHttpsredirect(v bool)`

SetUiHttpsredirect sets UiHttpsredirect field to given value.

### HasUiHttpsredirect

`func (o *SystemGeneralUpdate0) HasUiHttpsredirect() bool`

HasUiHttpsredirect returns a boolean if a field has been set.

### GetUiHttpsprotocols

`func (o *SystemGeneralUpdate0) GetUiHttpsprotocols() []string`

GetUiHttpsprotocols returns the UiHttpsprotocols field if non-nil, zero value otherwise.

### GetUiHttpsprotocolsOk

`func (o *SystemGeneralUpdate0) GetUiHttpsprotocolsOk() (*[]string, bool)`

GetUiHttpsprotocolsOk returns a tuple with the UiHttpsprotocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiHttpsprotocols

`func (o *SystemGeneralUpdate0) SetUiHttpsprotocols(v []string)`

SetUiHttpsprotocols sets UiHttpsprotocols field to given value.

### HasUiHttpsprotocols

`func (o *SystemGeneralUpdate0) HasUiHttpsprotocols() bool`

HasUiHttpsprotocols returns a boolean if a field has been set.

### GetUiPort

`func (o *SystemGeneralUpdate0) GetUiPort() int32`

GetUiPort returns the UiPort field if non-nil, zero value otherwise.

### GetUiPortOk

`func (o *SystemGeneralUpdate0) GetUiPortOk() (*int32, bool)`

GetUiPortOk returns a tuple with the UiPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiPort

`func (o *SystemGeneralUpdate0) SetUiPort(v int32)`

SetUiPort sets UiPort field to given value.

### HasUiPort

`func (o *SystemGeneralUpdate0) HasUiPort() bool`

HasUiPort returns a boolean if a field has been set.

### GetUiAddress

`func (o *SystemGeneralUpdate0) GetUiAddress() []string`

GetUiAddress returns the UiAddress field if non-nil, zero value otherwise.

### GetUiAddressOk

`func (o *SystemGeneralUpdate0) GetUiAddressOk() (*[]string, bool)`

GetUiAddressOk returns a tuple with the UiAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiAddress

`func (o *SystemGeneralUpdate0) SetUiAddress(v []string)`

SetUiAddress sets UiAddress field to given value.

### HasUiAddress

`func (o *SystemGeneralUpdate0) HasUiAddress() bool`

HasUiAddress returns a boolean if a field has been set.

### GetUiV6address

`func (o *SystemGeneralUpdate0) GetUiV6address() []string`

GetUiV6address returns the UiV6address field if non-nil, zero value otherwise.

### GetUiV6addressOk

`func (o *SystemGeneralUpdate0) GetUiV6addressOk() (*[]string, bool)`

GetUiV6addressOk returns a tuple with the UiV6address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiV6address

`func (o *SystemGeneralUpdate0) SetUiV6address(v []string)`

SetUiV6address sets UiV6address field to given value.

### HasUiV6address

`func (o *SystemGeneralUpdate0) HasUiV6address() bool`

HasUiV6address returns a boolean if a field has been set.

### GetKbdmap

`func (o *SystemGeneralUpdate0) GetKbdmap() string`

GetKbdmap returns the Kbdmap field if non-nil, zero value otherwise.

### GetKbdmapOk

`func (o *SystemGeneralUpdate0) GetKbdmapOk() (*string, bool)`

GetKbdmapOk returns a tuple with the Kbdmap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKbdmap

`func (o *SystemGeneralUpdate0) SetKbdmap(v string)`

SetKbdmap sets Kbdmap field to given value.

### HasKbdmap

`func (o *SystemGeneralUpdate0) HasKbdmap() bool`

HasKbdmap returns a boolean if a field has been set.

### GetLanguage

`func (o *SystemGeneralUpdate0) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *SystemGeneralUpdate0) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *SystemGeneralUpdate0) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *SystemGeneralUpdate0) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetSysloglevel

`func (o *SystemGeneralUpdate0) GetSysloglevel() string`

GetSysloglevel returns the Sysloglevel field if non-nil, zero value otherwise.

### GetSysloglevelOk

`func (o *SystemGeneralUpdate0) GetSysloglevelOk() (*string, bool)`

GetSysloglevelOk returns a tuple with the Sysloglevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysloglevel

`func (o *SystemGeneralUpdate0) SetSysloglevel(v string)`

SetSysloglevel sets Sysloglevel field to given value.

### HasSysloglevel

`func (o *SystemGeneralUpdate0) HasSysloglevel() bool`

HasSysloglevel returns a boolean if a field has been set.

### GetSyslogserver

`func (o *SystemGeneralUpdate0) GetSyslogserver() string`

GetSyslogserver returns the Syslogserver field if non-nil, zero value otherwise.

### GetSyslogserverOk

`func (o *SystemGeneralUpdate0) GetSyslogserverOk() (*string, bool)`

GetSyslogserverOk returns a tuple with the Syslogserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogserver

`func (o *SystemGeneralUpdate0) SetSyslogserver(v string)`

SetSyslogserver sets Syslogserver field to given value.

### HasSyslogserver

`func (o *SystemGeneralUpdate0) HasSyslogserver() bool`

HasSyslogserver returns a boolean if a field has been set.

### GetTimezone

`func (o *SystemGeneralUpdate0) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *SystemGeneralUpdate0) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *SystemGeneralUpdate0) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *SystemGeneralUpdate0) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetCrashReporting

`func (o *SystemGeneralUpdate0) GetCrashReporting() bool`

GetCrashReporting returns the CrashReporting field if non-nil, zero value otherwise.

### GetCrashReportingOk

`func (o *SystemGeneralUpdate0) GetCrashReportingOk() (*bool, bool)`

GetCrashReportingOk returns a tuple with the CrashReporting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrashReporting

`func (o *SystemGeneralUpdate0) SetCrashReporting(v bool)`

SetCrashReporting sets CrashReporting field to given value.

### HasCrashReporting

`func (o *SystemGeneralUpdate0) HasCrashReporting() bool`

HasCrashReporting returns a boolean if a field has been set.

### SetCrashReportingNil

`func (o *SystemGeneralUpdate0) SetCrashReportingNil(b bool)`

 SetCrashReportingNil sets the value for CrashReporting to be an explicit nil

### UnsetCrashReporting
`func (o *SystemGeneralUpdate0) UnsetCrashReporting()`

UnsetCrashReporting ensures that no value is present for CrashReporting, not even an explicit nil
### GetUsageCollection

`func (o *SystemGeneralUpdate0) GetUsageCollection() bool`

GetUsageCollection returns the UsageCollection field if non-nil, zero value otherwise.

### GetUsageCollectionOk

`func (o *SystemGeneralUpdate0) GetUsageCollectionOk() (*bool, bool)`

GetUsageCollectionOk returns a tuple with the UsageCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCollection

`func (o *SystemGeneralUpdate0) SetUsageCollection(v bool)`

SetUsageCollection sets UsageCollection field to given value.

### HasUsageCollection

`func (o *SystemGeneralUpdate0) HasUsageCollection() bool`

HasUsageCollection returns a boolean if a field has been set.

### SetUsageCollectionNil

`func (o *SystemGeneralUpdate0) SetUsageCollectionNil(b bool)`

 SetUsageCollectionNil sets the value for UsageCollection to be an explicit nil

### UnsetUsageCollection
`func (o *SystemGeneralUpdate0) UnsetUsageCollection()`

UnsetUsageCollection ensures that no value is present for UsageCollection, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


