# NotificationRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActorId** | **string** |  | 
**CreatedAt** | **string** |  | 
**Events** | Pointer to [**[]NotificationRuleEvent**](NotificationRuleEvent.md) |  | [optional] 
**Id** | **string** |  | 
**IdentityId** | Pointer to **string** |  | [optional] 
**RecipientType** | **string** |  | 
**Recipients** | Pointer to [**[]NotificationRuleRecipient**](NotificationRuleRecipient.md) |  | [optional] 
**TargetId** | Pointer to **string** |  | [optional] 
**Title** | **string** |  | 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewNotificationRule

`func NewNotificationRule(actorId string, createdAt string, id string, recipientType string, title string, ) *NotificationRule`

NewNotificationRule instantiates a new NotificationRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRuleWithDefaults

`func NewNotificationRuleWithDefaults() *NotificationRule`

NewNotificationRuleWithDefaults instantiates a new NotificationRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActorId

`func (o *NotificationRule) GetActorId() string`

GetActorId returns the ActorId field if non-nil, zero value otherwise.

### GetActorIdOk

`func (o *NotificationRule) GetActorIdOk() (*string, bool)`

GetActorIdOk returns a tuple with the ActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorId

`func (o *NotificationRule) SetActorId(v string)`

SetActorId sets ActorId field to given value.


### GetCreatedAt

`func (o *NotificationRule) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationRule) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationRule) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetEvents

`func (o *NotificationRule) GetEvents() []NotificationRuleEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *NotificationRule) GetEventsOk() (*[]NotificationRuleEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *NotificationRule) SetEvents(v []NotificationRuleEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *NotificationRule) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetId

`func (o *NotificationRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationRule) SetId(v string)`

SetId sets Id field to given value.


### GetIdentityId

`func (o *NotificationRule) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *NotificationRule) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *NotificationRule) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *NotificationRule) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetRecipientType

`func (o *NotificationRule) GetRecipientType() string`

GetRecipientType returns the RecipientType field if non-nil, zero value otherwise.

### GetRecipientTypeOk

`func (o *NotificationRule) GetRecipientTypeOk() (*string, bool)`

GetRecipientTypeOk returns a tuple with the RecipientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientType

`func (o *NotificationRule) SetRecipientType(v string)`

SetRecipientType sets RecipientType field to given value.


### GetRecipients

`func (o *NotificationRule) GetRecipients() []NotificationRuleRecipient`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *NotificationRule) GetRecipientsOk() (*[]NotificationRuleRecipient, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *NotificationRule) SetRecipients(v []NotificationRuleRecipient)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *NotificationRule) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetTargetId

`func (o *NotificationRule) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *NotificationRule) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *NotificationRule) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *NotificationRule) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetTitle

`func (o *NotificationRule) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NotificationRule) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NotificationRule) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUpdatedAt

`func (o *NotificationRule) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NotificationRule) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NotificationRule) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NotificationRule) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


