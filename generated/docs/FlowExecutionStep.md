# FlowExecutionStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**InputJson** | Pointer to **string** |  | [optional] 
**OutputJson** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewFlowExecutionStep

`func NewFlowExecutionStep() *FlowExecutionStep`

NewFlowExecutionStep instantiates a new FlowExecutionStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowExecutionStepWithDefaults

`func NewFlowExecutionStepWithDefaults() *FlowExecutionStep`

NewFlowExecutionStepWithDefaults instantiates a new FlowExecutionStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeId

`func (o *FlowExecutionStep) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *FlowExecutionStep) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *FlowExecutionStep) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *FlowExecutionStep) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetStatus

`func (o *FlowExecutionStep) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FlowExecutionStep) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FlowExecutionStep) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FlowExecutionStep) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetInputJson

`func (o *FlowExecutionStep) GetInputJson() string`

GetInputJson returns the InputJson field if non-nil, zero value otherwise.

### GetInputJsonOk

`func (o *FlowExecutionStep) GetInputJsonOk() (*string, bool)`

GetInputJsonOk returns a tuple with the InputJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputJson

`func (o *FlowExecutionStep) SetInputJson(v string)`

SetInputJson sets InputJson field to given value.

### HasInputJson

`func (o *FlowExecutionStep) HasInputJson() bool`

HasInputJson returns a boolean if a field has been set.

### GetOutputJson

`func (o *FlowExecutionStep) GetOutputJson() string`

GetOutputJson returns the OutputJson field if non-nil, zero value otherwise.

### GetOutputJsonOk

`func (o *FlowExecutionStep) GetOutputJsonOk() (*string, bool)`

GetOutputJsonOk returns a tuple with the OutputJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputJson

`func (o *FlowExecutionStep) SetOutputJson(v string)`

SetOutputJson sets OutputJson field to given value.

### HasOutputJson

`func (o *FlowExecutionStep) HasOutputJson() bool`

HasOutputJson returns a boolean if a field has been set.

### GetError

`func (o *FlowExecutionStep) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *FlowExecutionStep) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *FlowExecutionStep) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *FlowExecutionStep) HasError() bool`

HasError returns a boolean if a field has been set.

### GetStartedAt

`func (o *FlowExecutionStep) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *FlowExecutionStep) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *FlowExecutionStep) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *FlowExecutionStep) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


