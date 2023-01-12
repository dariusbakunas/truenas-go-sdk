# NetworkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**Ipv4gateway** | Pointer to **string** |  | [optional] 
**Ipv6gateway** | Pointer to **string** |  | [optional] 
**Nameserver1** | Pointer to **string** |  | [optional] 
**Nameserver2** | Pointer to **string** |  | [optional] 
**Nameserver3** | Pointer to **string** |  | [optional] 
**Httpproxy** | Pointer to **string** |  | [optional] 
**NetwaitEnabled** | Pointer to **bool** |  | [optional] 
**NetwaitIp** | Pointer to **[]string** |  | [optional] 
**Hosts** | Pointer to **string** |  | [optional] 
**Domains** | Pointer to **[]string** |  | [optional] 
**ServiceAnnouncement** | Pointer to [**NetworkConfigServiceAnnouncement**](NetworkConfigServiceAnnouncement.md) |  | [optional] 
**HostnameLocal** | Pointer to **string** |  | [optional] 

## Methods

### NewNetworkConfig

`func NewNetworkConfig() *NetworkConfig`

NewNetworkConfig instantiates a new NetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkConfigWithDefaults

`func NewNetworkConfigWithDefaults() *NetworkConfig`

NewNetworkConfigWithDefaults instantiates a new NetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkConfig) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkConfig) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkConfig) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHostname

`func (o *NetworkConfig) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *NetworkConfig) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *NetworkConfig) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *NetworkConfig) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetDomain

`func (o *NetworkConfig) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *NetworkConfig) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *NetworkConfig) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *NetworkConfig) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetIpv4gateway

`func (o *NetworkConfig) GetIpv4gateway() string`

GetIpv4gateway returns the Ipv4gateway field if non-nil, zero value otherwise.

### GetIpv4gatewayOk

`func (o *NetworkConfig) GetIpv4gatewayOk() (*string, bool)`

GetIpv4gatewayOk returns a tuple with the Ipv4gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4gateway

`func (o *NetworkConfig) SetIpv4gateway(v string)`

SetIpv4gateway sets Ipv4gateway field to given value.

### HasIpv4gateway

`func (o *NetworkConfig) HasIpv4gateway() bool`

HasIpv4gateway returns a boolean if a field has been set.

### GetIpv6gateway

`func (o *NetworkConfig) GetIpv6gateway() string`

GetIpv6gateway returns the Ipv6gateway field if non-nil, zero value otherwise.

### GetIpv6gatewayOk

`func (o *NetworkConfig) GetIpv6gatewayOk() (*string, bool)`

GetIpv6gatewayOk returns a tuple with the Ipv6gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6gateway

`func (o *NetworkConfig) SetIpv6gateway(v string)`

SetIpv6gateway sets Ipv6gateway field to given value.

### HasIpv6gateway

`func (o *NetworkConfig) HasIpv6gateway() bool`

HasIpv6gateway returns a boolean if a field has been set.

### GetNameserver1

`func (o *NetworkConfig) GetNameserver1() string`

GetNameserver1 returns the Nameserver1 field if non-nil, zero value otherwise.

### GetNameserver1Ok

`func (o *NetworkConfig) GetNameserver1Ok() (*string, bool)`

GetNameserver1Ok returns a tuple with the Nameserver1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameserver1

`func (o *NetworkConfig) SetNameserver1(v string)`

SetNameserver1 sets Nameserver1 field to given value.

### HasNameserver1

`func (o *NetworkConfig) HasNameserver1() bool`

HasNameserver1 returns a boolean if a field has been set.

### GetNameserver2

`func (o *NetworkConfig) GetNameserver2() string`

GetNameserver2 returns the Nameserver2 field if non-nil, zero value otherwise.

### GetNameserver2Ok

`func (o *NetworkConfig) GetNameserver2Ok() (*string, bool)`

GetNameserver2Ok returns a tuple with the Nameserver2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameserver2

`func (o *NetworkConfig) SetNameserver2(v string)`

SetNameserver2 sets Nameserver2 field to given value.

### HasNameserver2

`func (o *NetworkConfig) HasNameserver2() bool`

HasNameserver2 returns a boolean if a field has been set.

### GetNameserver3

`func (o *NetworkConfig) GetNameserver3() string`

GetNameserver3 returns the Nameserver3 field if non-nil, zero value otherwise.

### GetNameserver3Ok

`func (o *NetworkConfig) GetNameserver3Ok() (*string, bool)`

GetNameserver3Ok returns a tuple with the Nameserver3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameserver3

`func (o *NetworkConfig) SetNameserver3(v string)`

SetNameserver3 sets Nameserver3 field to given value.

### HasNameserver3

`func (o *NetworkConfig) HasNameserver3() bool`

HasNameserver3 returns a boolean if a field has been set.

### GetHttpproxy

`func (o *NetworkConfig) GetHttpproxy() string`

GetHttpproxy returns the Httpproxy field if non-nil, zero value otherwise.

### GetHttpproxyOk

`func (o *NetworkConfig) GetHttpproxyOk() (*string, bool)`

GetHttpproxyOk returns a tuple with the Httpproxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpproxy

`func (o *NetworkConfig) SetHttpproxy(v string)`

SetHttpproxy sets Httpproxy field to given value.

### HasHttpproxy

`func (o *NetworkConfig) HasHttpproxy() bool`

HasHttpproxy returns a boolean if a field has been set.

### GetNetwaitEnabled

`func (o *NetworkConfig) GetNetwaitEnabled() bool`

GetNetwaitEnabled returns the NetwaitEnabled field if non-nil, zero value otherwise.

### GetNetwaitEnabledOk

`func (o *NetworkConfig) GetNetwaitEnabledOk() (*bool, bool)`

GetNetwaitEnabledOk returns a tuple with the NetwaitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwaitEnabled

`func (o *NetworkConfig) SetNetwaitEnabled(v bool)`

SetNetwaitEnabled sets NetwaitEnabled field to given value.

### HasNetwaitEnabled

`func (o *NetworkConfig) HasNetwaitEnabled() bool`

HasNetwaitEnabled returns a boolean if a field has been set.

### GetNetwaitIp

`func (o *NetworkConfig) GetNetwaitIp() []string`

GetNetwaitIp returns the NetwaitIp field if non-nil, zero value otherwise.

### GetNetwaitIpOk

`func (o *NetworkConfig) GetNetwaitIpOk() (*[]string, bool)`

GetNetwaitIpOk returns a tuple with the NetwaitIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwaitIp

`func (o *NetworkConfig) SetNetwaitIp(v []string)`

SetNetwaitIp sets NetwaitIp field to given value.

### HasNetwaitIp

`func (o *NetworkConfig) HasNetwaitIp() bool`

HasNetwaitIp returns a boolean if a field has been set.

### GetHosts

`func (o *NetworkConfig) GetHosts() string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *NetworkConfig) GetHostsOk() (*string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *NetworkConfig) SetHosts(v string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *NetworkConfig) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetDomains

`func (o *NetworkConfig) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *NetworkConfig) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *NetworkConfig) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *NetworkConfig) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetServiceAnnouncement

`func (o *NetworkConfig) GetServiceAnnouncement() NetworkConfigServiceAnnouncement`

GetServiceAnnouncement returns the ServiceAnnouncement field if non-nil, zero value otherwise.

### GetServiceAnnouncementOk

`func (o *NetworkConfig) GetServiceAnnouncementOk() (*NetworkConfigServiceAnnouncement, bool)`

GetServiceAnnouncementOk returns a tuple with the ServiceAnnouncement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAnnouncement

`func (o *NetworkConfig) SetServiceAnnouncement(v NetworkConfigServiceAnnouncement)`

SetServiceAnnouncement sets ServiceAnnouncement field to given value.

### HasServiceAnnouncement

`func (o *NetworkConfig) HasServiceAnnouncement() bool`

HasServiceAnnouncement returns a boolean if a field has been set.

### GetHostnameLocal

`func (o *NetworkConfig) GetHostnameLocal() string`

GetHostnameLocal returns the HostnameLocal field if non-nil, zero value otherwise.

### GetHostnameLocalOk

`func (o *NetworkConfig) GetHostnameLocalOk() (*string, bool)`

GetHostnameLocalOk returns a tuple with the HostnameLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameLocal

`func (o *NetworkConfig) SetHostnameLocal(v string)`

SetHostnameLocal sets HostnameLocal field to given value.

### HasHostnameLocal

`func (o *NetworkConfig) HasHostnameLocal() bool`

HasHostnameLocal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


