# ListWorkspaceConnectionAssociationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]WorkspaceConnectionAssociation**](WorkspaceConnectionAssociation.md) |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewListWorkspaceConnectionAssociationsResponse

`func NewListWorkspaceConnectionAssociationsResponse() *ListWorkspaceConnectionAssociationsResponse`

NewListWorkspaceConnectionAssociationsResponse instantiates a new ListWorkspaceConnectionAssociationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWorkspaceConnectionAssociationsResponseWithDefaults

`func NewListWorkspaceConnectionAssociationsResponseWithDefaults() *ListWorkspaceConnectionAssociationsResponse`

NewListWorkspaceConnectionAssociationsResponseWithDefaults instantiates a new ListWorkspaceConnectionAssociationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListWorkspaceConnectionAssociationsResponse) GetItems() []WorkspaceConnectionAssociation`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListWorkspaceConnectionAssociationsResponse) GetItemsOk() (*[]WorkspaceConnectionAssociation, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListWorkspaceConnectionAssociationsResponse) SetItems(v []WorkspaceConnectionAssociation)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListWorkspaceConnectionAssociationsResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNextToken

`func (o *ListWorkspaceConnectionAssociationsResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *ListWorkspaceConnectionAssociationsResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *ListWorkspaceConnectionAssociationsResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *ListWorkspaceConnectionAssociationsResponse) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


