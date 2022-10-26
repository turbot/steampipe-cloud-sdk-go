# TemporaryTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientIp** | **string** |  | 
**Code** | **string** |  | 
**CreatedAt** | **string** |  | 
**Id** | **string** |  | 
**State** | **string** |  | 
**Token** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** | The time of the last update in ISO 8601 UTC. | [optional] 

## Methods

### NewTemporaryTokenRequest

`func NewTemporaryTokenRequest(clientIp string, code string, createdAt string, id string, state string, ) *TemporaryTokenRequest`

NewTemporaryTokenRequest instantiates a new TemporaryTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemporaryTokenRequestWithDefaults

`func NewTemporaryTokenRequestWithDefaults() *TemporaryTokenRequest`

NewTemporaryTokenRequestWithDefaults instantiates a new TemporaryTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientIp

`func (o *TemporaryTokenRequest) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *TemporaryTokenRequest) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *TemporaryTokenRequest) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.


### GetCode

`func (o *TemporaryTokenRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TemporaryTokenRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TemporaryTokenRequest) SetCode(v string)`

SetCode sets Code field to given value.


### GetCreatedAt

`func (o *TemporaryTokenRequest) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TemporaryTokenRequest) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TemporaryTokenRequest) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *TemporaryTokenRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemporaryTokenRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemporaryTokenRequest) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *TemporaryTokenRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TemporaryTokenRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TemporaryTokenRequest) SetState(v string)`

SetState sets State field to given value.


### GetToken

`func (o *TemporaryTokenRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TemporaryTokenRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TemporaryTokenRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TemporaryTokenRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TemporaryTokenRequest) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TemporaryTokenRequest) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TemporaryTokenRequest) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TemporaryTokenRequest) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


