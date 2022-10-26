# UserPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommunicationCommunityUpdates** | **string** | Whether the user is subscribed to community update emails. | 
**CommunicationProductUpdates** | **string** | Whether the user is subscribed to product update emails. | 
**CommunicationTipsAndTricks** | **string** | Whether the user is subscribed to tips and tricks emails. | 
**CreatedAt** | **string** | The time when the user setting was created. | 
**Id** | Pointer to **string** | The unique identifier of the user preferences record. | [optional] 
**UpdatedAt** | Pointer to **string** | The time when the user setting was last updated. | [optional] 
**VersionId** | **int32** | The current version of the user setting. | 

## Methods

### NewUserPreferences

`func NewUserPreferences(communicationCommunityUpdates string, communicationProductUpdates string, communicationTipsAndTricks string, createdAt string, versionId int32, ) *UserPreferences`

NewUserPreferences instantiates a new UserPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPreferencesWithDefaults

`func NewUserPreferencesWithDefaults() *UserPreferences`

NewUserPreferencesWithDefaults instantiates a new UserPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommunicationCommunityUpdates

`func (o *UserPreferences) GetCommunicationCommunityUpdates() string`

GetCommunicationCommunityUpdates returns the CommunicationCommunityUpdates field if non-nil, zero value otherwise.

### GetCommunicationCommunityUpdatesOk

`func (o *UserPreferences) GetCommunicationCommunityUpdatesOk() (*string, bool)`

GetCommunicationCommunityUpdatesOk returns a tuple with the CommunicationCommunityUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationCommunityUpdates

`func (o *UserPreferences) SetCommunicationCommunityUpdates(v string)`

SetCommunicationCommunityUpdates sets CommunicationCommunityUpdates field to given value.


### GetCommunicationProductUpdates

`func (o *UserPreferences) GetCommunicationProductUpdates() string`

GetCommunicationProductUpdates returns the CommunicationProductUpdates field if non-nil, zero value otherwise.

### GetCommunicationProductUpdatesOk

`func (o *UserPreferences) GetCommunicationProductUpdatesOk() (*string, bool)`

GetCommunicationProductUpdatesOk returns a tuple with the CommunicationProductUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationProductUpdates

`func (o *UserPreferences) SetCommunicationProductUpdates(v string)`

SetCommunicationProductUpdates sets CommunicationProductUpdates field to given value.


### GetCommunicationTipsAndTricks

`func (o *UserPreferences) GetCommunicationTipsAndTricks() string`

GetCommunicationTipsAndTricks returns the CommunicationTipsAndTricks field if non-nil, zero value otherwise.

### GetCommunicationTipsAndTricksOk

`func (o *UserPreferences) GetCommunicationTipsAndTricksOk() (*string, bool)`

GetCommunicationTipsAndTricksOk returns a tuple with the CommunicationTipsAndTricks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationTipsAndTricks

`func (o *UserPreferences) SetCommunicationTipsAndTricks(v string)`

SetCommunicationTipsAndTricks sets CommunicationTipsAndTricks field to given value.


### GetCreatedAt

`func (o *UserPreferences) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserPreferences) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserPreferences) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *UserPreferences) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserPreferences) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserPreferences) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserPreferences) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *UserPreferences) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserPreferences) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserPreferences) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UserPreferences) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersionId

`func (o *UserPreferences) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *UserPreferences) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *UserPreferences) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


