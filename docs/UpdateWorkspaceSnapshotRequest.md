# UpdateWorkspaceSnapshotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **interface{}** |  | [optional] 
**Title** | Pointer to **string** | The updated title for the snapshot. | [optional] 
**Visibility** | Pointer to **string** | The updated visibility for the snapshot. | [optional] 

## Methods

### NewUpdateWorkspaceSnapshotRequest

`func NewUpdateWorkspaceSnapshotRequest() *UpdateWorkspaceSnapshotRequest`

NewUpdateWorkspaceSnapshotRequest instantiates a new UpdateWorkspaceSnapshotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWorkspaceSnapshotRequestWithDefaults

`func NewUpdateWorkspaceSnapshotRequestWithDefaults() *UpdateWorkspaceSnapshotRequest`

NewUpdateWorkspaceSnapshotRequestWithDefaults instantiates a new UpdateWorkspaceSnapshotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *UpdateWorkspaceSnapshotRequest) GetTags() interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateWorkspaceSnapshotRequest) GetTagsOk() (*interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateWorkspaceSnapshotRequest) SetTags(v interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateWorkspaceSnapshotRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *UpdateWorkspaceSnapshotRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *UpdateWorkspaceSnapshotRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetTitle

`func (o *UpdateWorkspaceSnapshotRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateWorkspaceSnapshotRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateWorkspaceSnapshotRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateWorkspaceSnapshotRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateWorkspaceSnapshotRequest) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateWorkspaceSnapshotRequest) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateWorkspaceSnapshotRequest) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateWorkspaceSnapshotRequest) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


