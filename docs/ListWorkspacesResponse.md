# ListWorkspacesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Workspace**](Workspace.md) |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewListWorkspacesResponse

`func NewListWorkspacesResponse() *ListWorkspacesResponse`

NewListWorkspacesResponse instantiates a new ListWorkspacesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWorkspacesResponseWithDefaults

`func NewListWorkspacesResponseWithDefaults() *ListWorkspacesResponse`

NewListWorkspacesResponseWithDefaults instantiates a new ListWorkspacesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListWorkspacesResponse) GetItems() []Workspace`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListWorkspacesResponse) GetItemsOk() (*[]Workspace, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListWorkspacesResponse) SetItems(v []Workspace)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListWorkspacesResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNextToken

`func (o *ListWorkspacesResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *ListWorkspacesResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *ListWorkspacesResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *ListWorkspacesResponse) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


