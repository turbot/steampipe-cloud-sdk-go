# UserLoginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**State** | Pointer to **string** |  | [optional] 

## Methods

### NewUserLoginRequest

`func NewUserLoginRequest(email string, ) *UserLoginRequest`

NewUserLoginRequest instantiates a new UserLoginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserLoginRequestWithDefaults

`func NewUserLoginRequestWithDefaults() *UserLoginRequest`

NewUserLoginRequestWithDefaults instantiates a new UserLoginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserLoginRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserLoginRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserLoginRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetState

`func (o *UserLoginRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UserLoginRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UserLoginRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UserLoginRequest) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


