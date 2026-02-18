# FlowExecution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**FlowId** | **string** |  | 
**FlowVersion** | Pointer to **int32** |  | [optional] 
**TriggerId** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**CurrentNodeId** | Pointer to **string** |  | [optional] 
**Input** | Pointer to **map[string]interface{}** |  | [optional] 
**Output** | Pointer to **map[string]interface{}** |  | [optional] 
**Steps** | Pointer to [**[]ExecutionStep**](ExecutionStep.md) |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**EndedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewFlowExecution

`func NewFlowExecution(id string, flowId string, status string, ) *FlowExecution`

NewFlowExecution instantiates a new FlowExecution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowExecutionWithDefaults

`func NewFlowExecutionWithDefaults() *FlowExecution`

NewFlowExecutionWithDefaults instantiates a new FlowExecution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FlowExecution) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlowExecution) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlowExecution) SetId(v string)`

SetId sets Id field to given value.


### GetFlowId

`func (o *FlowExecution) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *FlowExecution) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *FlowExecution) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.


### GetFlowVersion

`func (o *FlowExecution) GetFlowVersion() int32`

GetFlowVersion returns the FlowVersion field if non-nil, zero value otherwise.

### GetFlowVersionOk

`func (o *FlowExecution) GetFlowVersionOk() (*int32, bool)`

GetFlowVersionOk returns a tuple with the FlowVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowVersion

`func (o *FlowExecution) SetFlowVersion(v int32)`

SetFlowVersion sets FlowVersion field to given value.

### HasFlowVersion

`func (o *FlowExecution) HasFlowVersion() bool`

HasFlowVersion returns a boolean if a field has been set.

### GetTriggerId

`func (o *FlowExecution) GetTriggerId() string`

GetTriggerId returns the TriggerId field if non-nil, zero value otherwise.

### GetTriggerIdOk

`func (o *FlowExecution) GetTriggerIdOk() (*string, bool)`

GetTriggerIdOk returns a tuple with the TriggerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerId

`func (o *FlowExecution) SetTriggerId(v string)`

SetTriggerId sets TriggerId field to given value.

### HasTriggerId

`func (o *FlowExecution) HasTriggerId() bool`

HasTriggerId returns a boolean if a field has been set.

### GetStatus

`func (o *FlowExecution) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FlowExecution) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FlowExecution) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCurrentNodeId

`func (o *FlowExecution) GetCurrentNodeId() string`

GetCurrentNodeId returns the CurrentNodeId field if non-nil, zero value otherwise.

### GetCurrentNodeIdOk

`func (o *FlowExecution) GetCurrentNodeIdOk() (*string, bool)`

GetCurrentNodeIdOk returns a tuple with the CurrentNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentNodeId

`func (o *FlowExecution) SetCurrentNodeId(v string)`

SetCurrentNodeId sets CurrentNodeId field to given value.

### HasCurrentNodeId

`func (o *FlowExecution) HasCurrentNodeId() bool`

HasCurrentNodeId returns a boolean if a field has been set.

### GetInput

`func (o *FlowExecution) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *FlowExecution) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *FlowExecution) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *FlowExecution) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetOutput

`func (o *FlowExecution) GetOutput() map[string]interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *FlowExecution) GetOutputOk() (*map[string]interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *FlowExecution) SetOutput(v map[string]interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *FlowExecution) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetSteps

`func (o *FlowExecution) GetSteps() []ExecutionStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *FlowExecution) GetStepsOk() (*[]ExecutionStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *FlowExecution) SetSteps(v []ExecutionStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *FlowExecution) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetStartedAt

`func (o *FlowExecution) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *FlowExecution) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *FlowExecution) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *FlowExecution) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetEndedAt

`func (o *FlowExecution) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *FlowExecution) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *FlowExecution) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *FlowExecution) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


