# TypesOrgQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Association** | [**map[string]TypesQuota**](TypesQuota.md) |  | 
**Conn** | [**TypesQuota**](TypesQuota.md) |  | 
**Workspace** | [**TypesQuota**](TypesQuota.md) |  | 

## Methods

### NewTypesOrgQuota

`func NewTypesOrgQuota(association map[string]TypesQuota, conn TypesQuota, workspace TypesQuota, ) *TypesOrgQuota`

NewTypesOrgQuota instantiates a new TypesOrgQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesOrgQuotaWithDefaults

`func NewTypesOrgQuotaWithDefaults() *TypesOrgQuota`

NewTypesOrgQuotaWithDefaults instantiates a new TypesOrgQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociation

`func (o *TypesOrgQuota) GetAssociation() map[string]TypesQuota`

GetAssociation returns the Association field if non-nil, zero value otherwise.

### GetAssociationOk

`func (o *TypesOrgQuota) GetAssociationOk() (*map[string]TypesQuota, bool)`

GetAssociationOk returns a tuple with the Association field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociation

`func (o *TypesOrgQuota) SetAssociation(v map[string]TypesQuota)`

SetAssociation sets Association field to given value.


### GetConn

`func (o *TypesOrgQuota) GetConn() TypesQuota`

GetConn returns the Conn field if non-nil, zero value otherwise.

### GetConnOk

`func (o *TypesOrgQuota) GetConnOk() (*TypesQuota, bool)`

GetConnOk returns a tuple with the Conn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConn

`func (o *TypesOrgQuota) SetConn(v TypesQuota)`

SetConn sets Conn field to given value.


### GetWorkspace

`func (o *TypesOrgQuota) GetWorkspace() TypesQuota`

GetWorkspace returns the Workspace field if non-nil, zero value otherwise.

### GetWorkspaceOk

`func (o *TypesOrgQuota) GetWorkspaceOk() (*TypesQuota, bool)`

GetWorkspaceOk returns a tuple with the Workspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspace

`func (o *TypesOrgQuota) SetWorkspace(v TypesQuota)`

SetWorkspace sets Workspace field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


