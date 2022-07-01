# InviteOrgWorkspaceUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Handle** | Pointer to **string** |  | [optional] 
**Role** | **string** |  | 

## Methods

### NewInviteOrgWorkspaceUserRequest

`func NewInviteOrgWorkspaceUserRequest(role string, ) *InviteOrgWorkspaceUserRequest`

NewInviteOrgWorkspaceUserRequest instantiates a new InviteOrgWorkspaceUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteOrgWorkspaceUserRequestWithDefaults

`func NewInviteOrgWorkspaceUserRequestWithDefaults() *InviteOrgWorkspaceUserRequest`

NewInviteOrgWorkspaceUserRequestWithDefaults instantiates a new InviteOrgWorkspaceUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *InviteOrgWorkspaceUserRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InviteOrgWorkspaceUserRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InviteOrgWorkspaceUserRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *InviteOrgWorkspaceUserRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetHandle

`func (o *InviteOrgWorkspaceUserRequest) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *InviteOrgWorkspaceUserRequest) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *InviteOrgWorkspaceUserRequest) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *InviteOrgWorkspaceUserRequest) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetRole

`func (o *InviteOrgWorkspaceUserRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *InviteOrgWorkspaceUserRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *InviteOrgWorkspaceUserRequest) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


