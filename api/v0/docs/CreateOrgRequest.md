# CreateOrgRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvatarUrl** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Handle** | **string** |  | 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateOrgRequest

`func NewCreateOrgRequest(handle string, ) *CreateOrgRequest`

NewCreateOrgRequest instantiates a new CreateOrgRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrgRequestWithDefaults

`func NewCreateOrgRequestWithDefaults() *CreateOrgRequest`

NewCreateOrgRequestWithDefaults instantiates a new CreateOrgRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatarUrl

`func (o *CreateOrgRequest) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *CreateOrgRequest) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *CreateOrgRequest) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *CreateOrgRequest) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetDisplayName

`func (o *CreateOrgRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateOrgRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateOrgRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateOrgRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHandle

`func (o *CreateOrgRequest) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *CreateOrgRequest) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *CreateOrgRequest) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetUrl

`func (o *CreateOrgRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateOrgRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateOrgRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CreateOrgRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


