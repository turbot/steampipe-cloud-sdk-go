# TypesAuditRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionType** | **string** |  | 
**ActorAvatarUrl** | **string** |  | 
**ActorDisplayName** | **string** |  | 
**ActorHandle** | **string** |  | 
**ActorId** | **string** |  | 
**ActorIp** | **string** |  | 
**CreatedAt** | **string** |  | 
**Data** | **map[string]interface{}** |  | 
**Id** | **string** |  | 
**IdentityHandle** | **string** |  | 
**IdentityId** | **string** |  | 
**TargetHandle** | Pointer to **string** |  | [optional] 
**TargetId** | Pointer to **string** |  | [optional] 

## Methods

### NewTypesAuditRecord

`func NewTypesAuditRecord(actionType string, actorAvatarUrl string, actorDisplayName string, actorHandle string, actorId string, actorIp string, createdAt string, data map[string]interface{}, id string, identityHandle string, identityId string, ) *TypesAuditRecord`

NewTypesAuditRecord instantiates a new TypesAuditRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesAuditRecordWithDefaults

`func NewTypesAuditRecordWithDefaults() *TypesAuditRecord`

NewTypesAuditRecordWithDefaults instantiates a new TypesAuditRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionType

`func (o *TypesAuditRecord) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *TypesAuditRecord) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *TypesAuditRecord) SetActionType(v string)`

SetActionType sets ActionType field to given value.


### GetActorAvatarUrl

`func (o *TypesAuditRecord) GetActorAvatarUrl() string`

GetActorAvatarUrl returns the ActorAvatarUrl field if non-nil, zero value otherwise.

### GetActorAvatarUrlOk

`func (o *TypesAuditRecord) GetActorAvatarUrlOk() (*string, bool)`

GetActorAvatarUrlOk returns a tuple with the ActorAvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorAvatarUrl

`func (o *TypesAuditRecord) SetActorAvatarUrl(v string)`

SetActorAvatarUrl sets ActorAvatarUrl field to given value.


### GetActorDisplayName

`func (o *TypesAuditRecord) GetActorDisplayName() string`

GetActorDisplayName returns the ActorDisplayName field if non-nil, zero value otherwise.

### GetActorDisplayNameOk

`func (o *TypesAuditRecord) GetActorDisplayNameOk() (*string, bool)`

GetActorDisplayNameOk returns a tuple with the ActorDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorDisplayName

`func (o *TypesAuditRecord) SetActorDisplayName(v string)`

SetActorDisplayName sets ActorDisplayName field to given value.


### GetActorHandle

`func (o *TypesAuditRecord) GetActorHandle() string`

GetActorHandle returns the ActorHandle field if non-nil, zero value otherwise.

### GetActorHandleOk

`func (o *TypesAuditRecord) GetActorHandleOk() (*string, bool)`

GetActorHandleOk returns a tuple with the ActorHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorHandle

`func (o *TypesAuditRecord) SetActorHandle(v string)`

SetActorHandle sets ActorHandle field to given value.


### GetActorId

`func (o *TypesAuditRecord) GetActorId() string`

GetActorId returns the ActorId field if non-nil, zero value otherwise.

### GetActorIdOk

`func (o *TypesAuditRecord) GetActorIdOk() (*string, bool)`

GetActorIdOk returns a tuple with the ActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorId

`func (o *TypesAuditRecord) SetActorId(v string)`

SetActorId sets ActorId field to given value.


### GetActorIp

`func (o *TypesAuditRecord) GetActorIp() string`

GetActorIp returns the ActorIp field if non-nil, zero value otherwise.

### GetActorIpOk

`func (o *TypesAuditRecord) GetActorIpOk() (*string, bool)`

GetActorIpOk returns a tuple with the ActorIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorIp

`func (o *TypesAuditRecord) SetActorIp(v string)`

SetActorIp sets ActorIp field to given value.


### GetCreatedAt

`func (o *TypesAuditRecord) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TypesAuditRecord) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TypesAuditRecord) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetData

`func (o *TypesAuditRecord) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TypesAuditRecord) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TypesAuditRecord) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetId

`func (o *TypesAuditRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TypesAuditRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TypesAuditRecord) SetId(v string)`

SetId sets Id field to given value.


### GetIdentityHandle

`func (o *TypesAuditRecord) GetIdentityHandle() string`

GetIdentityHandle returns the IdentityHandle field if non-nil, zero value otherwise.

### GetIdentityHandleOk

`func (o *TypesAuditRecord) GetIdentityHandleOk() (*string, bool)`

GetIdentityHandleOk returns a tuple with the IdentityHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityHandle

`func (o *TypesAuditRecord) SetIdentityHandle(v string)`

SetIdentityHandle sets IdentityHandle field to given value.


### GetIdentityId

`func (o *TypesAuditRecord) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *TypesAuditRecord) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *TypesAuditRecord) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.


### GetTargetHandle

`func (o *TypesAuditRecord) GetTargetHandle() string`

GetTargetHandle returns the TargetHandle field if non-nil, zero value otherwise.

### GetTargetHandleOk

`func (o *TypesAuditRecord) GetTargetHandleOk() (*string, bool)`

GetTargetHandleOk returns a tuple with the TargetHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetHandle

`func (o *TypesAuditRecord) SetTargetHandle(v string)`

SetTargetHandle sets TargetHandle field to given value.

### HasTargetHandle

`func (o *TypesAuditRecord) HasTargetHandle() bool`

HasTargetHandle returns a boolean if a field has been set.

### GetTargetId

`func (o *TypesAuditRecord) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *TypesAuditRecord) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *TypesAuditRecord) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *TypesAuditRecord) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


