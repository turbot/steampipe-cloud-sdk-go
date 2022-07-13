# ActorOrg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Org** | [**Org**](Org.md) |  | 
**Role** | **string** | The role of the actor in the org. | 
**Workspaces** | Pointer to [**[]ActorOrgWorkspace**](ActorOrgWorkspace.md) | Any workspaces the actor has explicit permission to within the org. Any org owner will have implicit owner roles over all workspaces, so they will not be listed here. | [optional] 

## Methods

### NewActorOrg

`func NewActorOrg(org Org, role string, ) *ActorOrg`

NewActorOrg instantiates a new ActorOrg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActorOrgWithDefaults

`func NewActorOrgWithDefaults() *ActorOrg`

NewActorOrgWithDefaults instantiates a new ActorOrg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrg

`func (o *ActorOrg) GetOrg() Org`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *ActorOrg) GetOrgOk() (*Org, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *ActorOrg) SetOrg(v Org)`

SetOrg sets Org field to given value.


### GetRole

`func (o *ActorOrg) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ActorOrg) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ActorOrg) SetRole(v string)`

SetRole sets Role field to given value.


### GetWorkspaces

`func (o *ActorOrg) GetWorkspaces() []ActorOrgWorkspace`

GetWorkspaces returns the Workspaces field if non-nil, zero value otherwise.

### GetWorkspacesOk

`func (o *ActorOrg) GetWorkspacesOk() (*[]ActorOrgWorkspace, bool)`

GetWorkspacesOk returns a tuple with the Workspaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaces

`func (o *ActorOrg) SetWorkspaces(v []ActorOrgWorkspace)`

SetWorkspaces sets Workspaces field to given value.

### HasWorkspaces

`func (o *ActorOrg) HasWorkspaces() bool`

HasWorkspaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


