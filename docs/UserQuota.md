# UserQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregator** | [**map[string]Quota**](Quota.md) |  | 
**Association** | [**map[string]Quota**](Quota.md) |  | 
**Conn** | [**Quota**](Quota.md) |  | 
**Mod** | [**map[string]Quota**](Quota.md) |  | 
**Organization** | [**Quota**](Quota.md) |  | 
**Pipeline** | [**map[string]Quota**](Quota.md) |  | 
**Snapshot** | [**map[string]Quota**](Quota.md) |  | 
**Token** | [**Quota**](Quota.md) |  | 
**Workspace** | [**Quota**](Quota.md) |  | 

## Methods

### NewUserQuota

`func NewUserQuota(aggregator map[string]Quota, association map[string]Quota, conn Quota, mod map[string]Quota, organization Quota, pipeline map[string]Quota, snapshot map[string]Quota, token Quota, workspace Quota, ) *UserQuota`

NewUserQuota instantiates a new UserQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserQuotaWithDefaults

`func NewUserQuotaWithDefaults() *UserQuota`

NewUserQuotaWithDefaults instantiates a new UserQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregator

`func (o *UserQuota) GetAggregator() map[string]Quota`

GetAggregator returns the Aggregator field if non-nil, zero value otherwise.

### GetAggregatorOk

`func (o *UserQuota) GetAggregatorOk() (*map[string]Quota, bool)`

GetAggregatorOk returns a tuple with the Aggregator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregator

`func (o *UserQuota) SetAggregator(v map[string]Quota)`

SetAggregator sets Aggregator field to given value.


### GetAssociation

`func (o *UserQuota) GetAssociation() map[string]Quota`

GetAssociation returns the Association field if non-nil, zero value otherwise.

### GetAssociationOk

`func (o *UserQuota) GetAssociationOk() (*map[string]Quota, bool)`

GetAssociationOk returns a tuple with the Association field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociation

`func (o *UserQuota) SetAssociation(v map[string]Quota)`

SetAssociation sets Association field to given value.


### GetConn

`func (o *UserQuota) GetConn() Quota`

GetConn returns the Conn field if non-nil, zero value otherwise.

### GetConnOk

`func (o *UserQuota) GetConnOk() (*Quota, bool)`

GetConnOk returns a tuple with the Conn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConn

`func (o *UserQuota) SetConn(v Quota)`

SetConn sets Conn field to given value.


### GetMod

`func (o *UserQuota) GetMod() map[string]Quota`

GetMod returns the Mod field if non-nil, zero value otherwise.

### GetModOk

`func (o *UserQuota) GetModOk() (*map[string]Quota, bool)`

GetModOk returns a tuple with the Mod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMod

`func (o *UserQuota) SetMod(v map[string]Quota)`

SetMod sets Mod field to given value.


### GetOrganization

`func (o *UserQuota) GetOrganization() Quota`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *UserQuota) GetOrganizationOk() (*Quota, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *UserQuota) SetOrganization(v Quota)`

SetOrganization sets Organization field to given value.


### GetPipeline

`func (o *UserQuota) GetPipeline() map[string]Quota`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *UserQuota) GetPipelineOk() (*map[string]Quota, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *UserQuota) SetPipeline(v map[string]Quota)`

SetPipeline sets Pipeline field to given value.


### GetSnapshot

`func (o *UserQuota) GetSnapshot() map[string]Quota`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *UserQuota) GetSnapshotOk() (*map[string]Quota, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *UserQuota) SetSnapshot(v map[string]Quota)`

SetSnapshot sets Snapshot field to given value.


### GetToken

`func (o *UserQuota) GetToken() Quota`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UserQuota) GetTokenOk() (*Quota, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UserQuota) SetToken(v Quota)`

SetToken sets Token field to given value.


### GetWorkspace

`func (o *UserQuota) GetWorkspace() Quota`

GetWorkspace returns the Workspace field if non-nil, zero value otherwise.

### GetWorkspaceOk

`func (o *UserQuota) GetWorkspaceOk() (*Quota, bool)`

GetWorkspaceOk returns a tuple with the Workspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspace

`func (o *UserQuota) SetWorkspace(v Quota)`

SetWorkspace sets Workspace field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


