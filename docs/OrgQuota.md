# OrgQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregator** | [**map[string]Quota**](Quota.md) |  | 
**Association** | [**map[string]Quota**](Quota.md) |  | 
**Conn** | [**Quota**](Quota.md) |  | 
**Member** | [**Quota**](Quota.md) |  | 
**Mod** | [**map[string]Quota**](Quota.md) |  | 
**Pipeline** | [**map[string]Quota**](Quota.md) |  | 
**Snapshot** | [**map[string]Quota**](Quota.md) |  | 
**Workspace** | [**Quota**](Quota.md) |  | 

## Methods

### NewOrgQuota

`func NewOrgQuota(aggregator map[string]Quota, association map[string]Quota, conn Quota, member Quota, mod map[string]Quota, pipeline map[string]Quota, snapshot map[string]Quota, workspace Quota, ) *OrgQuota`

NewOrgQuota instantiates a new OrgQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgQuotaWithDefaults

`func NewOrgQuotaWithDefaults() *OrgQuota`

NewOrgQuotaWithDefaults instantiates a new OrgQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregator

`func (o *OrgQuota) GetAggregator() map[string]Quota`

GetAggregator returns the Aggregator field if non-nil, zero value otherwise.

### GetAggregatorOk

`func (o *OrgQuota) GetAggregatorOk() (*map[string]Quota, bool)`

GetAggregatorOk returns a tuple with the Aggregator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregator

`func (o *OrgQuota) SetAggregator(v map[string]Quota)`

SetAggregator sets Aggregator field to given value.


### GetAssociation

`func (o *OrgQuota) GetAssociation() map[string]Quota`

GetAssociation returns the Association field if non-nil, zero value otherwise.

### GetAssociationOk

`func (o *OrgQuota) GetAssociationOk() (*map[string]Quota, bool)`

GetAssociationOk returns a tuple with the Association field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociation

`func (o *OrgQuota) SetAssociation(v map[string]Quota)`

SetAssociation sets Association field to given value.


### GetConn

`func (o *OrgQuota) GetConn() Quota`

GetConn returns the Conn field if non-nil, zero value otherwise.

### GetConnOk

`func (o *OrgQuota) GetConnOk() (*Quota, bool)`

GetConnOk returns a tuple with the Conn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConn

`func (o *OrgQuota) SetConn(v Quota)`

SetConn sets Conn field to given value.


### GetMember

`func (o *OrgQuota) GetMember() Quota`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *OrgQuota) GetMemberOk() (*Quota, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *OrgQuota) SetMember(v Quota)`

SetMember sets Member field to given value.


### GetMod

`func (o *OrgQuota) GetMod() map[string]Quota`

GetMod returns the Mod field if non-nil, zero value otherwise.

### GetModOk

`func (o *OrgQuota) GetModOk() (*map[string]Quota, bool)`

GetModOk returns a tuple with the Mod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMod

`func (o *OrgQuota) SetMod(v map[string]Quota)`

SetMod sets Mod field to given value.


### GetPipeline

`func (o *OrgQuota) GetPipeline() map[string]Quota`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *OrgQuota) GetPipelineOk() (*map[string]Quota, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *OrgQuota) SetPipeline(v map[string]Quota)`

SetPipeline sets Pipeline field to given value.


### GetSnapshot

`func (o *OrgQuota) GetSnapshot() map[string]Quota`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *OrgQuota) GetSnapshotOk() (*map[string]Quota, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *OrgQuota) SetSnapshot(v map[string]Quota)`

SetSnapshot sets Snapshot field to given value.


### GetWorkspace

`func (o *OrgQuota) GetWorkspace() Quota`

GetWorkspace returns the Workspace field if non-nil, zero value otherwise.

### GetWorkspaceOk

`func (o *OrgQuota) GetWorkspaceOk() (*Quota, bool)`

GetWorkspaceOk returns a tuple with the Workspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspace

`func (o *OrgQuota) SetWorkspace(v Quota)`

SetWorkspace sets Workspace field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


