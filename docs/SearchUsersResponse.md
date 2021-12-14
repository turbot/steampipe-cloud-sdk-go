# SearchUsersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]UserSearch**](UserSearch.md) |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewSearchUsersResponse

`func NewSearchUsersResponse() *SearchUsersResponse`

NewSearchUsersResponse instantiates a new SearchUsersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchUsersResponseWithDefaults

`func NewSearchUsersResponseWithDefaults() *SearchUsersResponse`

NewSearchUsersResponseWithDefaults instantiates a new SearchUsersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *SearchUsersResponse) GetItems() []UserSearch`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SearchUsersResponse) GetItemsOk() (*[]UserSearch, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SearchUsersResponse) SetItems(v []UserSearch)`

SetItems sets Items field to given value.

### HasItems

`func (o *SearchUsersResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNextToken

`func (o *SearchUsersResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *SearchUsersResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *SearchUsersResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *SearchUsersResponse) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


