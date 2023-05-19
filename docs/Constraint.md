# Constraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Approval** | Pointer to **string** | The status of the setting requested for the constraint at this level. | [optional] 
**DatatankId** | Pointer to **string** | The datatank id for which the constraint applies to. | [optional] 
**Description** | Pointer to **string** | A short description of the constraint. | [optional] 
**Limit** | Pointer to **interface{}** |  | [optional] 
**Name** | **string** | The name of the constraint. | 
**OrgId** | Pointer to **string** | The org id for which the constraint applies to. | [optional] 
**PipelineId** | Pointer to **string** | The pipeline id for which the constraint applies to. | [optional] 
**PlanId** | Pointer to **string** | The plan id for which the constraint applies to. | [optional] 
**Setting** | Pointer to **interface{}** |  | [optional] 
**Statistics** | Pointer to **interface{}** |  | [optional] 
**Target** | **string** | The resource which the constraint targets. | 
**Type** | **string** | The type of constraint. Can be one of rate, quota, range, feature or constant. | 
**UserId** | Pointer to **string** | The user id for which the constraint applies to. | [optional] 
**Value** | **interface{}** |  | 
**ValueFrom** | **string** | The id of the item from which the value is inherited. | 
**WorkspaceId** | Pointer to **string** | The workspace id for which the constraint applies to. | [optional] 

## Methods

### NewConstraint

`func NewConstraint(name string, target string, type_ string, value interface{}, valueFrom string, ) *Constraint`

NewConstraint instantiates a new Constraint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConstraintWithDefaults

`func NewConstraintWithDefaults() *Constraint`

NewConstraintWithDefaults instantiates a new Constraint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApproval

`func (o *Constraint) GetApproval() string`

GetApproval returns the Approval field if non-nil, zero value otherwise.

### GetApprovalOk

`func (o *Constraint) GetApprovalOk() (*string, bool)`

GetApprovalOk returns a tuple with the Approval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproval

`func (o *Constraint) SetApproval(v string)`

SetApproval sets Approval field to given value.

### HasApproval

`func (o *Constraint) HasApproval() bool`

HasApproval returns a boolean if a field has been set.

### GetDatatankId

`func (o *Constraint) GetDatatankId() string`

GetDatatankId returns the DatatankId field if non-nil, zero value otherwise.

### GetDatatankIdOk

`func (o *Constraint) GetDatatankIdOk() (*string, bool)`

GetDatatankIdOk returns a tuple with the DatatankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatatankId

`func (o *Constraint) SetDatatankId(v string)`

SetDatatankId sets DatatankId field to given value.

### HasDatatankId

`func (o *Constraint) HasDatatankId() bool`

HasDatatankId returns a boolean if a field has been set.

### GetDescription

`func (o *Constraint) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Constraint) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Constraint) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Constraint) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLimit

`func (o *Constraint) GetLimit() interface{}`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Constraint) GetLimitOk() (*interface{}, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Constraint) SetLimit(v interface{})`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Constraint) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *Constraint) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *Constraint) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetName

`func (o *Constraint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Constraint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Constraint) SetName(v string)`

SetName sets Name field to given value.


### GetOrgId

`func (o *Constraint) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Constraint) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Constraint) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *Constraint) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPipelineId

`func (o *Constraint) GetPipelineId() string`

GetPipelineId returns the PipelineId field if non-nil, zero value otherwise.

### GetPipelineIdOk

`func (o *Constraint) GetPipelineIdOk() (*string, bool)`

GetPipelineIdOk returns a tuple with the PipelineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineId

`func (o *Constraint) SetPipelineId(v string)`

SetPipelineId sets PipelineId field to given value.

### HasPipelineId

`func (o *Constraint) HasPipelineId() bool`

HasPipelineId returns a boolean if a field has been set.

### GetPlanId

`func (o *Constraint) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *Constraint) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *Constraint) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *Constraint) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetSetting

`func (o *Constraint) GetSetting() interface{}`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *Constraint) GetSettingOk() (*interface{}, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetting

`func (o *Constraint) SetSetting(v interface{})`

SetSetting sets Setting field to given value.

### HasSetting

`func (o *Constraint) HasSetting() bool`

HasSetting returns a boolean if a field has been set.

### SetSettingNil

`func (o *Constraint) SetSettingNil(b bool)`

 SetSettingNil sets the value for Setting to be an explicit nil

### UnsetSetting
`func (o *Constraint) UnsetSetting()`

UnsetSetting ensures that no value is present for Setting, not even an explicit nil
### GetStatistics

`func (o *Constraint) GetStatistics() interface{}`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *Constraint) GetStatisticsOk() (*interface{}, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *Constraint) SetStatistics(v interface{})`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *Constraint) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### SetStatisticsNil

`func (o *Constraint) SetStatisticsNil(b bool)`

 SetStatisticsNil sets the value for Statistics to be an explicit nil

### UnsetStatistics
`func (o *Constraint) UnsetStatistics()`

UnsetStatistics ensures that no value is present for Statistics, not even an explicit nil
### GetTarget

`func (o *Constraint) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *Constraint) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *Constraint) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetType

`func (o *Constraint) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Constraint) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Constraint) SetType(v string)`

SetType sets Type field to given value.


### GetUserId

`func (o *Constraint) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Constraint) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Constraint) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Constraint) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetValue

`func (o *Constraint) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Constraint) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Constraint) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *Constraint) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *Constraint) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetValueFrom

`func (o *Constraint) GetValueFrom() string`

GetValueFrom returns the ValueFrom field if non-nil, zero value otherwise.

### GetValueFromOk

`func (o *Constraint) GetValueFromOk() (*string, bool)`

GetValueFromOk returns a tuple with the ValueFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueFrom

`func (o *Constraint) SetValueFrom(v string)`

SetValueFrom sets ValueFrom field to given value.


### GetWorkspaceId

`func (o *Constraint) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *Constraint) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *Constraint) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *Constraint) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


