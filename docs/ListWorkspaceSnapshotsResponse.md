# ListWorkspaceSnapshotsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]WorkspaceSnapshot**](WorkspaceSnapshot.md) | Metadata  ListResponseMetadata &#x60;json:\&quot;metadata\&quot;&#x60; | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewListWorkspaceSnapshotsResponse

`func NewListWorkspaceSnapshotsResponse() *ListWorkspaceSnapshotsResponse`

NewListWorkspaceSnapshotsResponse instantiates a new ListWorkspaceSnapshotsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWorkspaceSnapshotsResponseWithDefaults

`func NewListWorkspaceSnapshotsResponseWithDefaults() *ListWorkspaceSnapshotsResponse`

NewListWorkspaceSnapshotsResponseWithDefaults instantiates a new ListWorkspaceSnapshotsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListWorkspaceSnapshotsResponse) GetItems() []WorkspaceSnapshot`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListWorkspaceSnapshotsResponse) GetItemsOk() (*[]WorkspaceSnapshot, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListWorkspaceSnapshotsResponse) SetItems(v []WorkspaceSnapshot)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListWorkspaceSnapshotsResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNextToken

`func (o *ListWorkspaceSnapshotsResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *ListWorkspaceSnapshotsResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *ListWorkspaceSnapshotsResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *ListWorkspaceSnapshotsResponse) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


