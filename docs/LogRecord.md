# LogRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActorAvatarUrl** | **string** | The avatar URL of the actor. | 
**ActorDisplayName** | **string** | The display name of the actor. | 
**ActorHandle** | **string** | The actor handle who executed the query. | 
**ActorId** | **string** | The actor ID who executed the query. | 
**CreatedAt** | **string** | The created time of the log. | 
**Duration** | Pointer to **int32** | The duration of the query. | [optional] 
**Id** | **string** | The unique identifier of the DB log. | 
**LogTimestamp** | Pointer to **string** | The time when the log got captured in the postgres. | [optional] 
**Query** | Pointer to **string** | The query being executed in the workspace. | [optional] 
**WorkspaceHandle** | **string** | The workspace handle where the query was executed. | 
**WorkspaceId** | **string** | The workspace ID where the query was executed. | 

## Methods

### NewLogRecord

`func NewLogRecord(actorAvatarUrl string, actorDisplayName string, actorHandle string, actorId string, createdAt string, id string, workspaceHandle string, workspaceId string, ) *LogRecord`

NewLogRecord instantiates a new LogRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogRecordWithDefaults

`func NewLogRecordWithDefaults() *LogRecord`

NewLogRecordWithDefaults instantiates a new LogRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActorAvatarUrl

`func (o *LogRecord) GetActorAvatarUrl() string`

GetActorAvatarUrl returns the ActorAvatarUrl field if non-nil, zero value otherwise.

### GetActorAvatarUrlOk

`func (o *LogRecord) GetActorAvatarUrlOk() (*string, bool)`

GetActorAvatarUrlOk returns a tuple with the ActorAvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorAvatarUrl

`func (o *LogRecord) SetActorAvatarUrl(v string)`

SetActorAvatarUrl sets ActorAvatarUrl field to given value.


### GetActorDisplayName

`func (o *LogRecord) GetActorDisplayName() string`

GetActorDisplayName returns the ActorDisplayName field if non-nil, zero value otherwise.

### GetActorDisplayNameOk

`func (o *LogRecord) GetActorDisplayNameOk() (*string, bool)`

GetActorDisplayNameOk returns a tuple with the ActorDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorDisplayName

`func (o *LogRecord) SetActorDisplayName(v string)`

SetActorDisplayName sets ActorDisplayName field to given value.


### GetActorHandle

`func (o *LogRecord) GetActorHandle() string`

GetActorHandle returns the ActorHandle field if non-nil, zero value otherwise.

### GetActorHandleOk

`func (o *LogRecord) GetActorHandleOk() (*string, bool)`

GetActorHandleOk returns a tuple with the ActorHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorHandle

`func (o *LogRecord) SetActorHandle(v string)`

SetActorHandle sets ActorHandle field to given value.


### GetActorId

`func (o *LogRecord) GetActorId() string`

GetActorId returns the ActorId field if non-nil, zero value otherwise.

### GetActorIdOk

`func (o *LogRecord) GetActorIdOk() (*string, bool)`

GetActorIdOk returns a tuple with the ActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorId

`func (o *LogRecord) SetActorId(v string)`

SetActorId sets ActorId field to given value.


### GetCreatedAt

`func (o *LogRecord) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LogRecord) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LogRecord) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDuration

`func (o *LogRecord) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *LogRecord) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *LogRecord) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *LogRecord) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetId

`func (o *LogRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogRecord) SetId(v string)`

SetId sets Id field to given value.


### GetLogTimestamp

`func (o *LogRecord) GetLogTimestamp() string`

GetLogTimestamp returns the LogTimestamp field if non-nil, zero value otherwise.

### GetLogTimestampOk

`func (o *LogRecord) GetLogTimestampOk() (*string, bool)`

GetLogTimestampOk returns a tuple with the LogTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogTimestamp

`func (o *LogRecord) SetLogTimestamp(v string)`

SetLogTimestamp sets LogTimestamp field to given value.

### HasLogTimestamp

`func (o *LogRecord) HasLogTimestamp() bool`

HasLogTimestamp returns a boolean if a field has been set.

### GetQuery

`func (o *LogRecord) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *LogRecord) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *LogRecord) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *LogRecord) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetWorkspaceHandle

`func (o *LogRecord) GetWorkspaceHandle() string`

GetWorkspaceHandle returns the WorkspaceHandle field if non-nil, zero value otherwise.

### GetWorkspaceHandleOk

`func (o *LogRecord) GetWorkspaceHandleOk() (*string, bool)`

GetWorkspaceHandleOk returns a tuple with the WorkspaceHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceHandle

`func (o *LogRecord) SetWorkspaceHandle(v string)`

SetWorkspaceHandle sets WorkspaceHandle field to given value.


### GetWorkspaceId

`func (o *LogRecord) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *LogRecord) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *LogRecord) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


