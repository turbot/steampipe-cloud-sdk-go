# WorkspaceSnapshotData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | The action type of the snapshot. | 
**DashboardNode** | **map[string]interface{}** |  | 
**EndTime** | **string** | The time the dashboard execution ended. | 
**Inputs** | Pointer to **map[string]interface{}** |  | [optional] 
**Layout** | **map[string]interface{}** |  | 
**Panels** | **map[string]interface{}** |  | 
**SchemaVersion** | **string** | The schema version of this snapshot. | 
**SearchPath** | **[]string** | The search path used for this snapshot. | 
**StartTime** | **string** | The time the dashboard execution started. | 
**Variables** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewWorkspaceSnapshotData

`func NewWorkspaceSnapshotData(action string, dashboardNode map[string]interface{}, endTime string, layout map[string]interface{}, panels map[string]interface{}, schemaVersion string, searchPath []string, startTime string, ) *WorkspaceSnapshotData`

NewWorkspaceSnapshotData instantiates a new WorkspaceSnapshotData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceSnapshotDataWithDefaults

`func NewWorkspaceSnapshotDataWithDefaults() *WorkspaceSnapshotData`

NewWorkspaceSnapshotDataWithDefaults instantiates a new WorkspaceSnapshotData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *WorkspaceSnapshotData) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WorkspaceSnapshotData) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WorkspaceSnapshotData) SetAction(v string)`

SetAction sets Action field to given value.


### GetDashboardNode

`func (o *WorkspaceSnapshotData) GetDashboardNode() map[string]interface{}`

GetDashboardNode returns the DashboardNode field if non-nil, zero value otherwise.

### GetDashboardNodeOk

`func (o *WorkspaceSnapshotData) GetDashboardNodeOk() (*map[string]interface{}, bool)`

GetDashboardNodeOk returns a tuple with the DashboardNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardNode

`func (o *WorkspaceSnapshotData) SetDashboardNode(v map[string]interface{})`

SetDashboardNode sets DashboardNode field to given value.


### GetEndTime

`func (o *WorkspaceSnapshotData) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *WorkspaceSnapshotData) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *WorkspaceSnapshotData) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.


### GetInputs

`func (o *WorkspaceSnapshotData) GetInputs() map[string]interface{}`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *WorkspaceSnapshotData) GetInputsOk() (*map[string]interface{}, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *WorkspaceSnapshotData) SetInputs(v map[string]interface{})`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *WorkspaceSnapshotData) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetLayout

`func (o *WorkspaceSnapshotData) GetLayout() map[string]interface{}`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *WorkspaceSnapshotData) GetLayoutOk() (*map[string]interface{}, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *WorkspaceSnapshotData) SetLayout(v map[string]interface{})`

SetLayout sets Layout field to given value.


### GetPanels

`func (o *WorkspaceSnapshotData) GetPanels() map[string]interface{}`

GetPanels returns the Panels field if non-nil, zero value otherwise.

### GetPanelsOk

`func (o *WorkspaceSnapshotData) GetPanelsOk() (*map[string]interface{}, bool)`

GetPanelsOk returns a tuple with the Panels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanels

`func (o *WorkspaceSnapshotData) SetPanels(v map[string]interface{})`

SetPanels sets Panels field to given value.


### GetSchemaVersion

`func (o *WorkspaceSnapshotData) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *WorkspaceSnapshotData) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *WorkspaceSnapshotData) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.


### GetSearchPath

`func (o *WorkspaceSnapshotData) GetSearchPath() []string`

GetSearchPath returns the SearchPath field if non-nil, zero value otherwise.

### GetSearchPathOk

`func (o *WorkspaceSnapshotData) GetSearchPathOk() (*[]string, bool)`

GetSearchPathOk returns a tuple with the SearchPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchPath

`func (o *WorkspaceSnapshotData) SetSearchPath(v []string)`

SetSearchPath sets SearchPath field to given value.


### GetStartTime

`func (o *WorkspaceSnapshotData) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *WorkspaceSnapshotData) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *WorkspaceSnapshotData) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetVariables

`func (o *WorkspaceSnapshotData) GetVariables() map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *WorkspaceSnapshotData) GetVariablesOk() (*map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *WorkspaceSnapshotData) SetVariables(v map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *WorkspaceSnapshotData) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


