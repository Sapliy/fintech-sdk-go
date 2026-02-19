# AutomationFlowExecution

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
**Steps** | Pointer to [**[]AutomationFlowExecutionStep**](AutomationFlowExecutionStep.md) |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**EndedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAutomationFlowExecution

`func NewAutomationFlowExecution(id string, flowId string, status string, ) *AutomationFlowExecution`

NewAutomationFlowExecution instantiates a new AutomationFlowExecution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationFlowExecutionWithDefaults

`func NewAutomationFlowExecutionWithDefaults() *AutomationFlowExecution`

NewAutomationFlowExecutionWithDefaults instantiates a new AutomationFlowExecution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutomationFlowExecution) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutomationFlowExecution) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutomationFlowExecution) SetId(v string)`

SetId sets Id field to given value.


### GetFlowId

`func (o *AutomationFlowExecution) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *AutomationFlowExecution) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *AutomationFlowExecution) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.


### GetFlowVersion

`func (o *AutomationFlowExecution) GetFlowVersion() int32`

GetFlowVersion returns the FlowVersion field if non-nil, zero value otherwise.

### GetFlowVersionOk

`func (o *AutomationFlowExecution) GetFlowVersionOk() (*int32, bool)`

GetFlowVersionOk returns a tuple with the FlowVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowVersion

`func (o *AutomationFlowExecution) SetFlowVersion(v int32)`

SetFlowVersion sets FlowVersion field to given value.

### HasFlowVersion

`func (o *AutomationFlowExecution) HasFlowVersion() bool`

HasFlowVersion returns a boolean if a field has been set.

### GetTriggerId

`func (o *AutomationFlowExecution) GetTriggerId() string`

GetTriggerId returns the TriggerId field if non-nil, zero value otherwise.

### GetTriggerIdOk

`func (o *AutomationFlowExecution) GetTriggerIdOk() (*string, bool)`

GetTriggerIdOk returns a tuple with the TriggerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerId

`func (o *AutomationFlowExecution) SetTriggerId(v string)`

SetTriggerId sets TriggerId field to given value.

### HasTriggerId

`func (o *AutomationFlowExecution) HasTriggerId() bool`

HasTriggerId returns a boolean if a field has been set.

### GetStatus

`func (o *AutomationFlowExecution) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AutomationFlowExecution) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AutomationFlowExecution) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCurrentNodeId

`func (o *AutomationFlowExecution) GetCurrentNodeId() string`

GetCurrentNodeId returns the CurrentNodeId field if non-nil, zero value otherwise.

### GetCurrentNodeIdOk

`func (o *AutomationFlowExecution) GetCurrentNodeIdOk() (*string, bool)`

GetCurrentNodeIdOk returns a tuple with the CurrentNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentNodeId

`func (o *AutomationFlowExecution) SetCurrentNodeId(v string)`

SetCurrentNodeId sets CurrentNodeId field to given value.

### HasCurrentNodeId

`func (o *AutomationFlowExecution) HasCurrentNodeId() bool`

HasCurrentNodeId returns a boolean if a field has been set.

### GetInput

`func (o *AutomationFlowExecution) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *AutomationFlowExecution) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *AutomationFlowExecution) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *AutomationFlowExecution) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetOutput

`func (o *AutomationFlowExecution) GetOutput() map[string]interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *AutomationFlowExecution) GetOutputOk() (*map[string]interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *AutomationFlowExecution) SetOutput(v map[string]interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *AutomationFlowExecution) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetSteps

`func (o *AutomationFlowExecution) GetSteps() []AutomationFlowExecutionStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *AutomationFlowExecution) GetStepsOk() (*[]AutomationFlowExecutionStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *AutomationFlowExecution) SetSteps(v []AutomationFlowExecutionStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *AutomationFlowExecution) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetStartedAt

`func (o *AutomationFlowExecution) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *AutomationFlowExecution) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *AutomationFlowExecution) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *AutomationFlowExecution) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetEndedAt

`func (o *AutomationFlowExecution) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *AutomationFlowExecution) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *AutomationFlowExecution) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *AutomationFlowExecution) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


