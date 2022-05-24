# CreateOrgNotificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | **[]string** |  | 
**Sender** | Pointer to [**NotificationRecipientRequest**](NotificationRecipientRequest.md) |  | [optional] 
**Title** | **string** |  | 

## Methods

### NewCreateOrgNotificationRequest

`func NewCreateOrgNotificationRequest(events []string, title string, ) *CreateOrgNotificationRequest`

NewCreateOrgNotificationRequest instantiates a new CreateOrgNotificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrgNotificationRequestWithDefaults

`func NewCreateOrgNotificationRequestWithDefaults() *CreateOrgNotificationRequest`

NewCreateOrgNotificationRequestWithDefaults instantiates a new CreateOrgNotificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *CreateOrgNotificationRequest) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CreateOrgNotificationRequest) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CreateOrgNotificationRequest) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetSender

`func (o *CreateOrgNotificationRequest) GetSender() NotificationRecipientRequest`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *CreateOrgNotificationRequest) GetSenderOk() (*NotificationRecipientRequest, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *CreateOrgNotificationRequest) SetSender(v NotificationRecipientRequest)`

SetSender sets Sender field to given value.

### HasSender

`func (o *CreateOrgNotificationRequest) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetTitle

`func (o *CreateOrgNotificationRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateOrgNotificationRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateOrgNotificationRequest) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


