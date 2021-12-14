# InviteOrgUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Handle** | Pointer to **string** |  | [optional] 
**Role** | **string** |  | 

## Methods

### NewInviteOrgUserRequest

`func NewInviteOrgUserRequest(role string, ) *InviteOrgUserRequest`

NewInviteOrgUserRequest instantiates a new InviteOrgUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteOrgUserRequestWithDefaults

`func NewInviteOrgUserRequestWithDefaults() *InviteOrgUserRequest`

NewInviteOrgUserRequestWithDefaults instantiates a new InviteOrgUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *InviteOrgUserRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InviteOrgUserRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InviteOrgUserRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *InviteOrgUserRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetHandle

`func (o *InviteOrgUserRequest) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *InviteOrgUserRequest) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *InviteOrgUserRequest) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *InviteOrgUserRequest) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetRole

`func (o *InviteOrgUserRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *InviteOrgUserRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *InviteOrgUserRequest) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


