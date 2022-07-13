# ActorOrgWorkspace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | **string** | The role of the actor in the workspace. | 
**Workspace** | Pointer to [**Workspace**](Workspace.md) |  | [optional] 

## Methods

### NewActorOrgWorkspace

`func NewActorOrgWorkspace(role string, ) *ActorOrgWorkspace`

NewActorOrgWorkspace instantiates a new ActorOrgWorkspace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActorOrgWorkspaceWithDefaults

`func NewActorOrgWorkspaceWithDefaults() *ActorOrgWorkspace`

NewActorOrgWorkspaceWithDefaults instantiates a new ActorOrgWorkspace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *ActorOrgWorkspace) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ActorOrgWorkspace) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ActorOrgWorkspace) SetRole(v string)`

SetRole sets Role field to given value.


### GetWorkspace

`func (o *ActorOrgWorkspace) GetWorkspace() Workspace`

GetWorkspace returns the Workspace field if non-nil, zero value otherwise.

### GetWorkspaceOk

`func (o *ActorOrgWorkspace) GetWorkspaceOk() (*Workspace, bool)`

GetWorkspaceOk returns a tuple with the Workspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspace

`func (o *ActorOrgWorkspace) SetWorkspace(v Workspace)`

SetWorkspace sets Workspace field to given value.

### HasWorkspace

`func (o *ActorOrgWorkspace) HasWorkspace() bool`

HasWorkspace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


