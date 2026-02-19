# AutomationFlowExecutionStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | **string** |  | 
**Status** | **string** |  | 
**Input** | Pointer to **map[string]interface{}** |  | [optional] 
**Output** | Pointer to **map[string]interface{}** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewAutomationFlowExecutionStep

`func NewAutomationFlowExecutionStep(nodeId string, status string, ) *AutomationFlowExecutionStep`

NewAutomationFlowExecutionStep instantiates a new AutomationFlowExecutionStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationFlowExecutionStepWithDefaults

`func NewAutomationFlowExecutionStepWithDefaults() *AutomationFlowExecutionStep`

NewAutomationFlowExecutionStepWithDefaults instantiates a new AutomationFlowExecutionStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeId

`func (o *AutomationFlowExecutionStep) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *AutomationFlowExecutionStep) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *AutomationFlowExecutionStep) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetStatus

`func (o *AutomationFlowExecutionStep) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AutomationFlowExecutionStep) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AutomationFlowExecutionStep) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetInput

`func (o *AutomationFlowExecutionStep) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *AutomationFlowExecutionStep) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *AutomationFlowExecutionStep) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *AutomationFlowExecutionStep) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetOutput

`func (o *AutomationFlowExecutionStep) GetOutput() map[string]interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *AutomationFlowExecutionStep) GetOutputOk() (*map[string]interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *AutomationFlowExecutionStep) SetOutput(v map[string]interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *AutomationFlowExecutionStep) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetError

`func (o *AutomationFlowExecutionStep) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AutomationFlowExecutionStep) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AutomationFlowExecutionStep) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AutomationFlowExecutionStep) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


