# NotificationRecipientRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recipients** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewNotificationRecipientRequest

`func NewNotificationRecipientRequest() *NotificationRecipientRequest`

NewNotificationRecipientRequest instantiates a new NotificationRecipientRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRecipientRequestWithDefaults

`func NewNotificationRecipientRequestWithDefaults() *NotificationRecipientRequest`

NewNotificationRecipientRequestWithDefaults instantiates a new NotificationRecipientRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipients

`func (o *NotificationRecipientRequest) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *NotificationRecipientRequest) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *NotificationRecipientRequest) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *NotificationRecipientRequest) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetType

`func (o *NotificationRecipientRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationRecipientRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationRecipientRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NotificationRecipientRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


