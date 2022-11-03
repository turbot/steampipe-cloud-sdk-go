/*
Steampipe Cloud

Steampipe Cloud is a hosted version of Steampipe (https://steampipe.io), an open source tool to instantly query your cloud services (e.g. AWS, Azure, GCP and more) with SQL. No DB required.

API version: {{OPEN_API_VERSION}}
Contact: help@steampipe.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package steampipecloud

import (
	"encoding/json"
)

// LogRecord struct for LogRecord
type LogRecord struct {
	// The avatar URL of the actor.
	ActorAvatarUrl string `json:"actor_avatar_url"`
	// The display name of the actor.
	ActorDisplayName string `json:"actor_display_name"`
	// The actor handle who executed the query.
	ActorHandle string `json:"actor_handle"`
	// The actor ID who executed the query.
	ActorId string `json:"actor_id"`
	// The created time of the log.
	CreatedAt string `json:"created_at"`
	// The duration of the query.
	Duration *int32 `json:"duration,omitempty"`
	// The unique identifier of the DB log.
	Id string `json:"id"`
	// The time when the log got captured in the postgres.
	LogTimestamp *string `json:"log_timestamp,omitempty"`
	// The query being executed in the workspace.
	Query *string `json:"query,omitempty"`
	// The workspace handle where the query was executed.
	WorkspaceHandle string `json:"workspace_handle"`
	// The workspace ID where the query was executed.
	WorkspaceId string `json:"workspace_id"`
}

// NewLogRecord instantiates a new LogRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogRecord(actorAvatarUrl string, actorDisplayName string, actorHandle string, actorId string, createdAt string, id string, workspaceHandle string, workspaceId string) *LogRecord {
	this := LogRecord{}
	this.ActorAvatarUrl = actorAvatarUrl
	this.ActorDisplayName = actorDisplayName
	this.ActorHandle = actorHandle
	this.ActorId = actorId
	this.CreatedAt = createdAt
	this.Id = id
	this.WorkspaceHandle = workspaceHandle
	this.WorkspaceId = workspaceId
	return &this
}

// NewLogRecordWithDefaults instantiates a new LogRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogRecordWithDefaults() *LogRecord {
	this := LogRecord{}
	return &this
}

// GetActorAvatarUrl returns the ActorAvatarUrl field value
func (o *LogRecord) GetActorAvatarUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActorAvatarUrl
}

// GetActorAvatarUrlOk returns a tuple with the ActorAvatarUrl field value
// and a boolean to check if the value has been set.
func (o *LogRecord) GetActorAvatarUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActorAvatarUrl, true
}

// SetActorAvatarUrl sets field value
func (o *LogRecord) SetActorAvatarUrl(v string) {
	o.ActorAvatarUrl = v
}

// GetActorDisplayName returns the ActorDisplayName field value
func (o *LogRecord) GetActorDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActorDisplayName
}

// GetActorDisplayNameOk returns a tuple with the ActorDisplayName field value
// and a boolean to check if the value has been set.
func (o *LogRecord) GetActorDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActorDisplayName, true
}

// SetActorDisplayName sets field value
func (o *LogRecord) SetActorDisplayName(v string) {
	o.ActorDisplayName = v
}

// GetActorHandle returns the ActorHandle field value
func (o *LogRecord) GetActorHandle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActorHandle
}

// GetActorHandleOk returns a tuple with the ActorHandle field value
// and a boolean to check if the value has been set.
func (o *LogRecord) GetActorHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActorHandle, true
}

// SetActorHandle sets field value
func (o *LogRecord) SetActorHandle(v string) {
	o.ActorHandle = v
}

// GetActorId returns the ActorId field value
func (o *LogRecord) GetActorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActorId
}

// GetActorIdOk returns a tuple with the ActorId field value
// and a boolean to check if the value has been set.
func (o *LogRecord) GetActorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActorId, true
}

// SetActorId sets field value
func (o *LogRecord) SetActorId(v string) {
	o.ActorId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *LogRecord) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *LogRecord) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *LogRecord) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *LogRecord) GetDuration() int32 {
	if o == nil || o.Duration == nil {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogRecord) GetDurationOk() (*int32, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *LogRecord) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *LogRecord) SetDuration(v int32) {
	o.Duration = &v
}

// GetId returns the Id field value
func (o *LogRecord) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LogRecord) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LogRecord) SetId(v string) {
	o.Id = v
}

// GetLogTimestamp returns the LogTimestamp field value if set, zero value otherwise.
func (o *LogRecord) GetLogTimestamp() string {
	if o == nil || o.LogTimestamp == nil {
		var ret string
		return ret
	}
	return *o.LogTimestamp
}

// GetLogTimestampOk returns a tuple with the LogTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogRecord) GetLogTimestampOk() (*string, bool) {
	if o == nil || o.LogTimestamp == nil {
		return nil, false
	}
	return o.LogTimestamp, true
}

// HasLogTimestamp returns a boolean if a field has been set.
func (o *LogRecord) HasLogTimestamp() bool {
	if o != nil && o.LogTimestamp != nil {
		return true
	}

	return false
}

// SetLogTimestamp gets a reference to the given string and assigns it to the LogTimestamp field.
func (o *LogRecord) SetLogTimestamp(v string) {
	o.LogTimestamp = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *LogRecord) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogRecord) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *LogRecord) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *LogRecord) SetQuery(v string) {
	o.Query = &v
}

// GetWorkspaceHandle returns the WorkspaceHandle field value
func (o *LogRecord) GetWorkspaceHandle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkspaceHandle
}

// GetWorkspaceHandleOk returns a tuple with the WorkspaceHandle field value
// and a boolean to check if the value has been set.
func (o *LogRecord) GetWorkspaceHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkspaceHandle, true
}

// SetWorkspaceHandle sets field value
func (o *LogRecord) SetWorkspaceHandle(v string) {
	o.WorkspaceHandle = v
}

// GetWorkspaceId returns the WorkspaceId field value
func (o *LogRecord) GetWorkspaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value
// and a boolean to check if the value has been set.
func (o *LogRecord) GetWorkspaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkspaceId, true
}

// SetWorkspaceId sets field value
func (o *LogRecord) SetWorkspaceId(v string) {
	o.WorkspaceId = v
}

func (o LogRecord) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["actor_avatar_url"] = o.ActorAvatarUrl
	}
	if true {
		toSerialize["actor_display_name"] = o.ActorDisplayName
	}
	if true {
		toSerialize["actor_handle"] = o.ActorHandle
	}
	if true {
		toSerialize["actor_id"] = o.ActorId
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.LogTimestamp != nil {
		toSerialize["log_timestamp"] = o.LogTimestamp
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if true {
		toSerialize["workspace_handle"] = o.WorkspaceHandle
	}
	if true {
		toSerialize["workspace_id"] = o.WorkspaceId
	}
	return json.Marshal(toSerialize)
}

type NullableLogRecord struct {
	value *LogRecord
	isSet bool
}

func (v NullableLogRecord) Get() *LogRecord {
	return v.value
}

func (v *NullableLogRecord) Set(val *LogRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableLogRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableLogRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogRecord(val *LogRecord) *NullableLogRecord {
	return &NullableLogRecord{value: val, isSet: true}
}

func (v NullableLogRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
