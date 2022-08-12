# AuditRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionType** | **string** | The action performed on the resource. | 
**ActorAvatarUrl** | **string** |  | 
**ActorDisplayName** | **string** | The display name of an actor. | 
**ActorHandle** | **string** | The handle name of an actor. | 
**ActorId** | **string** | The unique identifier of an actor. | 
**ActorIp** | **string** |  | 
**CreatedAt** | **string** | The time when the audit log was recorded. | 
**Data** | **map[string]interface{}** |  | 
**Id** | **string** | The unique identifier for an audit log. | 
**IdentityHandle** | **string** | The handle name for an identity where the action has been performed. | 
**IdentityId** | **string** | The unique identifier for an identity where the action has been performed. | 
**ProcessId** | Pointer to **string** |  | [optional] 
**TargetHandle** | Pointer to **string** | The handle name of the entity on which the action has been performed. | [optional] 
**TargetId** | Pointer to **string** | The unique identifier of the entity on which the action has been performed. | [optional] 

## Methods

### NewAuditRecord

`func NewAuditRecord(actionType string, actorAvatarUrl string, actorDisplayName string, actorHandle string, actorId string, actorIp string, createdAt string, data map[string]interface{}, id string, identityHandle string, identityId string, ) *AuditRecord`

NewAuditRecord instantiates a new AuditRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditRecordWithDefaults

`func NewAuditRecordWithDefaults() *AuditRecord`

NewAuditRecordWithDefaults instantiates a new AuditRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionType

`func (o *AuditRecord) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *AuditRecord) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *AuditRecord) SetActionType(v string)`

SetActionType sets ActionType field to given value.


### GetActorAvatarUrl

`func (o *AuditRecord) GetActorAvatarUrl() string`

GetActorAvatarUrl returns the ActorAvatarUrl field if non-nil, zero value otherwise.

### GetActorAvatarUrlOk

`func (o *AuditRecord) GetActorAvatarUrlOk() (*string, bool)`

GetActorAvatarUrlOk returns a tuple with the ActorAvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorAvatarUrl

`func (o *AuditRecord) SetActorAvatarUrl(v string)`

SetActorAvatarUrl sets ActorAvatarUrl field to given value.


### GetActorDisplayName

`func (o *AuditRecord) GetActorDisplayName() string`

GetActorDisplayName returns the ActorDisplayName field if non-nil, zero value otherwise.

### GetActorDisplayNameOk

`func (o *AuditRecord) GetActorDisplayNameOk() (*string, bool)`

GetActorDisplayNameOk returns a tuple with the ActorDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorDisplayName

`func (o *AuditRecord) SetActorDisplayName(v string)`

SetActorDisplayName sets ActorDisplayName field to given value.


### GetActorHandle

`func (o *AuditRecord) GetActorHandle() string`

GetActorHandle returns the ActorHandle field if non-nil, zero value otherwise.

### GetActorHandleOk

`func (o *AuditRecord) GetActorHandleOk() (*string, bool)`

GetActorHandleOk returns a tuple with the ActorHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorHandle

`func (o *AuditRecord) SetActorHandle(v string)`

SetActorHandle sets ActorHandle field to given value.


### GetActorId

`func (o *AuditRecord) GetActorId() string`

GetActorId returns the ActorId field if non-nil, zero value otherwise.

### GetActorIdOk

`func (o *AuditRecord) GetActorIdOk() (*string, bool)`

GetActorIdOk returns a tuple with the ActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorId

`func (o *AuditRecord) SetActorId(v string)`

SetActorId sets ActorId field to given value.


### GetActorIp

`func (o *AuditRecord) GetActorIp() string`

GetActorIp returns the ActorIp field if non-nil, zero value otherwise.

### GetActorIpOk

`func (o *AuditRecord) GetActorIpOk() (*string, bool)`

GetActorIpOk returns a tuple with the ActorIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorIp

`func (o *AuditRecord) SetActorIp(v string)`

SetActorIp sets ActorIp field to given value.


### GetCreatedAt

`func (o *AuditRecord) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuditRecord) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuditRecord) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetData

`func (o *AuditRecord) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AuditRecord) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AuditRecord) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetId

`func (o *AuditRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditRecord) SetId(v string)`

SetId sets Id field to given value.


### GetIdentityHandle

`func (o *AuditRecord) GetIdentityHandle() string`

GetIdentityHandle returns the IdentityHandle field if non-nil, zero value otherwise.

### GetIdentityHandleOk

`func (o *AuditRecord) GetIdentityHandleOk() (*string, bool)`

GetIdentityHandleOk returns a tuple with the IdentityHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityHandle

`func (o *AuditRecord) SetIdentityHandle(v string)`

SetIdentityHandle sets IdentityHandle field to given value.


### GetIdentityId

`func (o *AuditRecord) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *AuditRecord) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *AuditRecord) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.


### GetProcessId

`func (o *AuditRecord) GetProcessId() string`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *AuditRecord) GetProcessIdOk() (*string, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *AuditRecord) SetProcessId(v string)`

SetProcessId sets ProcessId field to given value.

### HasProcessId

`func (o *AuditRecord) HasProcessId() bool`

HasProcessId returns a boolean if a field has been set.

### GetTargetHandle

`func (o *AuditRecord) GetTargetHandle() string`

GetTargetHandle returns the TargetHandle field if non-nil, zero value otherwise.

### GetTargetHandleOk

`func (o *AuditRecord) GetTargetHandleOk() (*string, bool)`

GetTargetHandleOk returns a tuple with the TargetHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetHandle

`func (o *AuditRecord) SetTargetHandle(v string)`

SetTargetHandle sets TargetHandle field to given value.

### HasTargetHandle

`func (o *AuditRecord) HasTargetHandle() bool`

HasTargetHandle returns a boolean if a field has been set.

### GetTargetId

`func (o *AuditRecord) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *AuditRecord) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *AuditRecord) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *AuditRecord) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


