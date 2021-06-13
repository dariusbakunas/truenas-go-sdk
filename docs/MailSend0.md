# MailSend0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Html** | Pointer to **NullableString** |  | [optional] 
**To** | Pointer to **[]string** |  | [optional] 
**Cc** | Pointer to **[]string** |  | [optional] 
**Interval** | Pointer to **NullableInt32** |  | [optional] 
**Channel** | Pointer to **NullableString** |  | [optional] 
**Timeout** | Pointer to **int32** |  | [optional] 
**Attachments** | Pointer to **bool** |  | [optional] 
**Queue** | Pointer to **bool** |  | [optional] 
**ExtraHeaders** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewMailSend0

`func NewMailSend0() *MailSend0`

NewMailSend0 instantiates a new MailSend0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailSend0WithDefaults

`func NewMailSend0WithDefaults() *MailSend0`

NewMailSend0WithDefaults instantiates a new MailSend0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *MailSend0) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *MailSend0) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *MailSend0) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *MailSend0) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetText

`func (o *MailSend0) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *MailSend0) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *MailSend0) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *MailSend0) HasText() bool`

HasText returns a boolean if a field has been set.

### GetHtml

`func (o *MailSend0) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *MailSend0) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *MailSend0) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *MailSend0) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### SetHtmlNil

`func (o *MailSend0) SetHtmlNil(b bool)`

 SetHtmlNil sets the value for Html to be an explicit nil

### UnsetHtml
`func (o *MailSend0) UnsetHtml()`

UnsetHtml ensures that no value is present for Html, not even an explicit nil
### GetTo

`func (o *MailSend0) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *MailSend0) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *MailSend0) SetTo(v []string)`

SetTo sets To field to given value.

### HasTo

`func (o *MailSend0) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetCc

`func (o *MailSend0) GetCc() []string`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *MailSend0) GetCcOk() (*[]string, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *MailSend0) SetCc(v []string)`

SetCc sets Cc field to given value.

### HasCc

`func (o *MailSend0) HasCc() bool`

HasCc returns a boolean if a field has been set.

### GetInterval

`func (o *MailSend0) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *MailSend0) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *MailSend0) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *MailSend0) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### SetIntervalNil

`func (o *MailSend0) SetIntervalNil(b bool)`

 SetIntervalNil sets the value for Interval to be an explicit nil

### UnsetInterval
`func (o *MailSend0) UnsetInterval()`

UnsetInterval ensures that no value is present for Interval, not even an explicit nil
### GetChannel

`func (o *MailSend0) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *MailSend0) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *MailSend0) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *MailSend0) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### SetChannelNil

`func (o *MailSend0) SetChannelNil(b bool)`

 SetChannelNil sets the value for Channel to be an explicit nil

### UnsetChannel
`func (o *MailSend0) UnsetChannel()`

UnsetChannel ensures that no value is present for Channel, not even an explicit nil
### GetTimeout

`func (o *MailSend0) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *MailSend0) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *MailSend0) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *MailSend0) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetAttachments

`func (o *MailSend0) GetAttachments() bool`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *MailSend0) GetAttachmentsOk() (*bool, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *MailSend0) SetAttachments(v bool)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *MailSend0) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetQueue

`func (o *MailSend0) GetQueue() bool`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *MailSend0) GetQueueOk() (*bool, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *MailSend0) SetQueue(v bool)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *MailSend0) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### GetExtraHeaders

`func (o *MailSend0) GetExtraHeaders() map[string]map[string]interface{}`

GetExtraHeaders returns the ExtraHeaders field if non-nil, zero value otherwise.

### GetExtraHeadersOk

`func (o *MailSend0) GetExtraHeadersOk() (*map[string]map[string]interface{}, bool)`

GetExtraHeadersOk returns a tuple with the ExtraHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraHeaders

`func (o *MailSend0) SetExtraHeaders(v map[string]map[string]interface{})`

SetExtraHeaders sets ExtraHeaders field to given value.

### HasExtraHeaders

`func (o *MailSend0) HasExtraHeaders() bool`

HasExtraHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


