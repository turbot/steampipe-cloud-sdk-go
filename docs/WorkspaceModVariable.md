# WorkspaceModVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** |  | 
**CreatedBy** | Pointer to [**User**](User.md) |  | [optional] 
**CreatedById** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**ModAlias** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to [**User**](User.md) |  | [optional] 
**UpdatedById** | **string** |  | 
**Value** | Pointer to **interface{}** |  | [optional] 
**ValueDefault** | Pointer to **interface{}** |  | [optional] 
**ValueSetting** | Pointer to **interface{}** |  | [optional] 
**VersionId** | **int32** |  | 

## Methods

### NewWorkspaceModVariable

`func NewWorkspaceModVariable(createdAt string, createdById string, id string, updatedById string, versionId int32, ) *WorkspaceModVariable`

NewWorkspaceModVariable instantiates a new WorkspaceModVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceModVariableWithDefaults

`func NewWorkspaceModVariableWithDefaults() *WorkspaceModVariable`

NewWorkspaceModVariableWithDefaults instantiates a new WorkspaceModVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *WorkspaceModVariable) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkspaceModVariable) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkspaceModVariable) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *WorkspaceModVariable) GetCreatedBy() User`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *WorkspaceModVariable) GetCreatedByOk() (*User, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *WorkspaceModVariable) SetCreatedBy(v User)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *WorkspaceModVariable) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedById

`func (o *WorkspaceModVariable) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *WorkspaceModVariable) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *WorkspaceModVariable) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetDescription

`func (o *WorkspaceModVariable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkspaceModVariable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkspaceModVariable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkspaceModVariable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *WorkspaceModVariable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkspaceModVariable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkspaceModVariable) SetId(v string)`

SetId sets Id field to given value.


### GetModAlias

`func (o *WorkspaceModVariable) GetModAlias() string`

GetModAlias returns the ModAlias field if non-nil, zero value otherwise.

### GetModAliasOk

`func (o *WorkspaceModVariable) GetModAliasOk() (*string, bool)`

GetModAliasOk returns a tuple with the ModAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModAlias

`func (o *WorkspaceModVariable) SetModAlias(v string)`

SetModAlias sets ModAlias field to given value.

### HasModAlias

`func (o *WorkspaceModVariable) HasModAlias() bool`

HasModAlias returns a boolean if a field has been set.

### GetName

`func (o *WorkspaceModVariable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkspaceModVariable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkspaceModVariable) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkspaceModVariable) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *WorkspaceModVariable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkspaceModVariable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkspaceModVariable) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkspaceModVariable) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *WorkspaceModVariable) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WorkspaceModVariable) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WorkspaceModVariable) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WorkspaceModVariable) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *WorkspaceModVariable) GetUpdatedBy() User`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *WorkspaceModVariable) GetUpdatedByOk() (*User, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *WorkspaceModVariable) SetUpdatedBy(v User)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *WorkspaceModVariable) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedById

`func (o *WorkspaceModVariable) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *WorkspaceModVariable) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *WorkspaceModVariable) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.


### GetValue

`func (o *WorkspaceModVariable) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WorkspaceModVariable) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WorkspaceModVariable) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *WorkspaceModVariable) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *WorkspaceModVariable) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *WorkspaceModVariable) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetValueDefault

`func (o *WorkspaceModVariable) GetValueDefault() interface{}`

GetValueDefault returns the ValueDefault field if non-nil, zero value otherwise.

### GetValueDefaultOk

`func (o *WorkspaceModVariable) GetValueDefaultOk() (*interface{}, bool)`

GetValueDefaultOk returns a tuple with the ValueDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueDefault

`func (o *WorkspaceModVariable) SetValueDefault(v interface{})`

SetValueDefault sets ValueDefault field to given value.

### HasValueDefault

`func (o *WorkspaceModVariable) HasValueDefault() bool`

HasValueDefault returns a boolean if a field has been set.

### SetValueDefaultNil

`func (o *WorkspaceModVariable) SetValueDefaultNil(b bool)`

 SetValueDefaultNil sets the value for ValueDefault to be an explicit nil

### UnsetValueDefault
`func (o *WorkspaceModVariable) UnsetValueDefault()`

UnsetValueDefault ensures that no value is present for ValueDefault, not even an explicit nil
### GetValueSetting

`func (o *WorkspaceModVariable) GetValueSetting() interface{}`

GetValueSetting returns the ValueSetting field if non-nil, zero value otherwise.

### GetValueSettingOk

`func (o *WorkspaceModVariable) GetValueSettingOk() (*interface{}, bool)`

GetValueSettingOk returns a tuple with the ValueSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueSetting

`func (o *WorkspaceModVariable) SetValueSetting(v interface{})`

SetValueSetting sets ValueSetting field to given value.

### HasValueSetting

`func (o *WorkspaceModVariable) HasValueSetting() bool`

HasValueSetting returns a boolean if a field has been set.

### SetValueSettingNil

`func (o *WorkspaceModVariable) SetValueSettingNil(b bool)`

 SetValueSettingNil sets the value for ValueSetting to be an explicit nil

### UnsetValueSetting
`func (o *WorkspaceModVariable) UnsetValueSetting()`

UnsetValueSetting ensures that no value is present for ValueSetting, not even an explicit nil
### GetVersionId

`func (o *WorkspaceModVariable) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *WorkspaceModVariable) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *WorkspaceModVariable) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


