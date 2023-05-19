# ConstraintOverrideRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatatankId** | Pointer to **string** | The id of the datatank where the override will be set. | [optional] 
**Name** | **string** | The name of the constraint to request override for. | 
**OrgId** | Pointer to **string** | The id of the org where the override will be set. | [optional] 
**PipelineId** | Pointer to **string** | The id of the pipeline where the override will be set. | [optional] 
**Setting** | **interface{}** |  | 
**UserId** | Pointer to **string** | The id of the user where the override will be set. | [optional] 
**WorkspaceId** | Pointer to **string** | The id of the workspace where the override will be set. | [optional] 

## Methods

### NewConstraintOverrideRequest

`func NewConstraintOverrideRequest(name string, setting interface{}, ) *ConstraintOverrideRequest`

NewConstraintOverrideRequest instantiates a new ConstraintOverrideRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConstraintOverrideRequestWithDefaults

`func NewConstraintOverrideRequestWithDefaults() *ConstraintOverrideRequest`

NewConstraintOverrideRequestWithDefaults instantiates a new ConstraintOverrideRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatatankId

`func (o *ConstraintOverrideRequest) GetDatatankId() string`

GetDatatankId returns the DatatankId field if non-nil, zero value otherwise.

### GetDatatankIdOk

`func (o *ConstraintOverrideRequest) GetDatatankIdOk() (*string, bool)`

GetDatatankIdOk returns a tuple with the DatatankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatatankId

`func (o *ConstraintOverrideRequest) SetDatatankId(v string)`

SetDatatankId sets DatatankId field to given value.

### HasDatatankId

`func (o *ConstraintOverrideRequest) HasDatatankId() bool`

HasDatatankId returns a boolean if a field has been set.

### GetName

`func (o *ConstraintOverrideRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConstraintOverrideRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConstraintOverrideRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOrgId

`func (o *ConstraintOverrideRequest) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ConstraintOverrideRequest) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ConstraintOverrideRequest) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ConstraintOverrideRequest) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPipelineId

`func (o *ConstraintOverrideRequest) GetPipelineId() string`

GetPipelineId returns the PipelineId field if non-nil, zero value otherwise.

### GetPipelineIdOk

`func (o *ConstraintOverrideRequest) GetPipelineIdOk() (*string, bool)`

GetPipelineIdOk returns a tuple with the PipelineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineId

`func (o *ConstraintOverrideRequest) SetPipelineId(v string)`

SetPipelineId sets PipelineId field to given value.

### HasPipelineId

`func (o *ConstraintOverrideRequest) HasPipelineId() bool`

HasPipelineId returns a boolean if a field has been set.

### GetSetting

`func (o *ConstraintOverrideRequest) GetSetting() interface{}`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *ConstraintOverrideRequest) GetSettingOk() (*interface{}, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetting

`func (o *ConstraintOverrideRequest) SetSetting(v interface{})`

SetSetting sets Setting field to given value.


### SetSettingNil

`func (o *ConstraintOverrideRequest) SetSettingNil(b bool)`

 SetSettingNil sets the value for Setting to be an explicit nil

### UnsetSetting
`func (o *ConstraintOverrideRequest) UnsetSetting()`

UnsetSetting ensures that no value is present for Setting, not even an explicit nil
### GetUserId

`func (o *ConstraintOverrideRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ConstraintOverrideRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ConstraintOverrideRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ConstraintOverrideRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *ConstraintOverrideRequest) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *ConstraintOverrideRequest) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *ConstraintOverrideRequest) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *ConstraintOverrideRequest) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


