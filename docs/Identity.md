# Identity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvatarUrl** | Pointer to **string** | The avatar URL of the identity. | [optional] 
**DisplayName** | Pointer to **string** | The display name of the identity. | [optional] 
**Handle** | **string** | The handle name of the identity. | 
**Id** | **string** | The unique identifier of the identity. | 
**Type** | **string** | The type of identity e.g. user or org. | 

## Methods

### NewIdentity

`func NewIdentity(handle string, id string, type_ string, ) *Identity`

NewIdentity instantiates a new Identity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityWithDefaults

`func NewIdentityWithDefaults() *Identity`

NewIdentityWithDefaults instantiates a new Identity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatarUrl

`func (o *Identity) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *Identity) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *Identity) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *Identity) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetDisplayName

`func (o *Identity) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Identity) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Identity) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Identity) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHandle

`func (o *Identity) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *Identity) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *Identity) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetId

`func (o *Identity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Identity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Identity) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *Identity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Identity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Identity) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


