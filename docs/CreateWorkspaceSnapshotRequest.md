# CreateWorkspaceSnapshotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**WorkspaceSnapshotData**](WorkspaceSnapshotData.md) |  | 
**Tags** | Pointer to **interface{}** |  | [optional] 
**Title** | Pointer to **string** | The title of the snapshot. | [optional] 
**Visibility** | Pointer to **string** | The visibility of the snapshot to create. | [optional] 

## Methods

### NewCreateWorkspaceSnapshotRequest

`func NewCreateWorkspaceSnapshotRequest(data WorkspaceSnapshotData, ) *CreateWorkspaceSnapshotRequest`

NewCreateWorkspaceSnapshotRequest instantiates a new CreateWorkspaceSnapshotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWorkspaceSnapshotRequestWithDefaults

`func NewCreateWorkspaceSnapshotRequestWithDefaults() *CreateWorkspaceSnapshotRequest`

NewCreateWorkspaceSnapshotRequestWithDefaults instantiates a new CreateWorkspaceSnapshotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateWorkspaceSnapshotRequest) GetData() WorkspaceSnapshotData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateWorkspaceSnapshotRequest) GetDataOk() (*WorkspaceSnapshotData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateWorkspaceSnapshotRequest) SetData(v WorkspaceSnapshotData)`

SetData sets Data field to given value.


### GetTags

`func (o *CreateWorkspaceSnapshotRequest) GetTags() interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateWorkspaceSnapshotRequest) GetTagsOk() (*interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateWorkspaceSnapshotRequest) SetTags(v interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateWorkspaceSnapshotRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *CreateWorkspaceSnapshotRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *CreateWorkspaceSnapshotRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetTitle

`func (o *CreateWorkspaceSnapshotRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateWorkspaceSnapshotRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateWorkspaceSnapshotRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateWorkspaceSnapshotRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetVisibility

`func (o *CreateWorkspaceSnapshotRequest) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *CreateWorkspaceSnapshotRequest) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *CreateWorkspaceSnapshotRequest) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *CreateWorkspaceSnapshotRequest) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


