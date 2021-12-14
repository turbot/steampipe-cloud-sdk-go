# OrgQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Association** | [**map[string]Quota**](Quota.md) |  | 
**Conn** | [**Quota**](Quota.md) |  | 
**Workspace** | [**Quota**](Quota.md) |  | 

## Methods

### NewOrgQuota

`func NewOrgQuota(association map[string]Quota, conn Quota, workspace Quota, ) *OrgQuota`

NewOrgQuota instantiates a new OrgQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgQuotaWithDefaults

`func NewOrgQuotaWithDefaults() *OrgQuota`

NewOrgQuotaWithDefaults instantiates a new OrgQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


