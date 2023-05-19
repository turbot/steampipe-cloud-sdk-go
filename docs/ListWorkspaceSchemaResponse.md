# ListWorkspaceSchemaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]WorkspaceSchema**](WorkspaceSchema.md) |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewListWorkspaceSchemaResponse

`func NewListWorkspaceSchemaResponse() *ListWorkspaceSchemaResponse`

NewListWorkspaceSchemaResponse instantiates a new ListWorkspaceSchemaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWorkspaceSchemaResponseWithDefaults

`func NewListWorkspaceSchemaResponseWithDefaults() *ListWorkspaceSchemaResponse`

NewListWorkspaceSchemaResponseWithDefaults instantiates a new ListWorkspaceSchemaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListWorkspaceSchemaResponse) GetItems() []WorkspaceSchema`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListWorkspaceSchemaResponse) GetItemsOk() (*[]WorkspaceSchema, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListWorkspaceSchemaResponse) SetItems(v []WorkspaceSchema)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListWorkspaceSchemaResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNextToken

`func (o *ListWorkspaceSchemaResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *ListWorkspaceSchemaResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *ListWorkspaceSchemaResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *ListWorkspaceSchemaResponse) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


