# TypesListUsersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]TypesUser**](TypesUser.md) |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewTypesListUsersResponse

`func NewTypesListUsersResponse() *TypesListUsersResponse`

NewTypesListUsersResponse instantiates a new TypesListUsersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesListUsersResponseWithDefaults

`func NewTypesListUsersResponseWithDefaults() *TypesListUsersResponse`

NewTypesListUsersResponseWithDefaults instantiates a new TypesListUsersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *TypesListUsersResponse) GetItems() []TypesUser`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TypesListUsersResponse) GetItemsOk() (*[]TypesUser, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TypesListUsersResponse) SetItems(v []TypesUser)`

SetItems sets Items field to given value.

### HasItems

`func (o *TypesListUsersResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNextToken

`func (o *TypesListUsersResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *TypesListUsersResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *TypesListUsersResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *TypesListUsersResponse) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


