# UserSignupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**State** | Pointer to **string** |  | [optional] 

## Methods

### NewUserSignupRequest

`func NewUserSignupRequest(email string, ) *UserSignupRequest`

NewUserSignupRequest instantiates a new UserSignupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSignupRequestWithDefaults

`func NewUserSignupRequestWithDefaults() *UserSignupRequest`

NewUserSignupRequestWithDefaults instantiates a new UserSignupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserSignupRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserSignupRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserSignupRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetState

`func (o *UserSignupRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UserSignupRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UserSignupRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UserSignupRequest) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


