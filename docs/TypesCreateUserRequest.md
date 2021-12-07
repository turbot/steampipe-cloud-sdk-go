# TypesCreateUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**Email** | **string** |  | 
**Handle** | **string** |  | 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewTypesCreateUserRequest

`func NewTypesCreateUserRequest(email string, handle string, ) *TypesCreateUserRequest`

NewTypesCreateUserRequest instantiates a new TypesCreateUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesCreateUserRequestWithDefaults

`func NewTypesCreateUserRequestWithDefaults() *TypesCreateUserRequest`

NewTypesCreateUserRequestWithDefaults instantiates a new TypesCreateUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *TypesCreateUserRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TypesCreateUserRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TypesCreateUserRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *TypesCreateUserRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmail

`func (o *TypesCreateUserRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TypesCreateUserRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TypesCreateUserRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetHandle

`func (o *TypesCreateUserRequest) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *TypesCreateUserRequest) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *TypesCreateUserRequest) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetUrl

`func (o *TypesCreateUserRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TypesCreateUserRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TypesCreateUserRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TypesCreateUserRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


