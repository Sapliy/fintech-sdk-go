# FlowFlowExecution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**FlowId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**CurrentNodeId** | Pointer to **string** |  | [optional] 
**InputJson** | Pointer to **string** |  | [optional] 
**MetadataJson** | Pointer to **string** |  | [optional] 
**Steps** | Pointer to [**[]FlowExecutionStep**](FlowExecutionStep.md) |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**FinishedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewFlowFlowExecution

`func NewFlowFlowExecution() *FlowFlowExecution`

NewFlowFlowExecution instantiates a new FlowFlowExecution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowFlowExecutionWithDefaults

`func NewFlowFlowExecutionWithDefaults() *FlowFlowExecution`

NewFlowFlowExecutionWithDefaults instantiates a new FlowFlowExecution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FlowFlowExecution) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlowFlowExecution) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlowFlowExecution) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FlowFlowExecution) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFlowId

`func (o *FlowFlowExecution) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *FlowFlowExecution) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *FlowFlowExecution) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.

### HasFlowId

`func (o *FlowFlowExecution) HasFlowId() bool`

HasFlowId returns a boolean if a field has been set.

### GetStatus

`func (o *FlowFlowExecution) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FlowFlowExecution) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FlowFlowExecution) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FlowFlowExecution) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCurrentNodeId

`func (o *FlowFlowExecution) GetCurrentNodeId() string`

GetCurrentNodeId returns the CurrentNodeId field if non-nil, zero value otherwise.

### GetCurrentNodeIdOk

`func (o *FlowFlowExecution) GetCurrentNodeIdOk() (*string, bool)`

GetCurrentNodeIdOk returns a tuple with the CurrentNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentNodeId

`func (o *FlowFlowExecution) SetCurrentNodeId(v string)`

SetCurrentNodeId sets CurrentNodeId field to given value.

### HasCurrentNodeId

`func (o *FlowFlowExecution) HasCurrentNodeId() bool`

HasCurrentNodeId returns a boolean if a field has been set.

### GetInputJson

`func (o *FlowFlowExecution) GetInputJson() string`

GetInputJson returns the InputJson field if non-nil, zero value otherwise.

### GetInputJsonOk

`func (o *FlowFlowExecution) GetInputJsonOk() (*string, bool)`

GetInputJsonOk returns a tuple with the InputJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputJson

`func (o *FlowFlowExecution) SetInputJson(v string)`

SetInputJson sets InputJson field to given value.

### HasInputJson

`func (o *FlowFlowExecution) HasInputJson() bool`

HasInputJson returns a boolean if a field has been set.

### GetMetadataJson

`func (o *FlowFlowExecution) GetMetadataJson() string`

GetMetadataJson returns the MetadataJson field if non-nil, zero value otherwise.

### GetMetadataJsonOk

`func (o *FlowFlowExecution) GetMetadataJsonOk() (*string, bool)`

GetMetadataJsonOk returns a tuple with the MetadataJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataJson

`func (o *FlowFlowExecution) SetMetadataJson(v string)`

SetMetadataJson sets MetadataJson field to given value.

### HasMetadataJson

`func (o *FlowFlowExecution) HasMetadataJson() bool`

HasMetadataJson returns a boolean if a field has been set.

### GetSteps

`func (o *FlowFlowExecution) GetSteps() []FlowExecutionStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *FlowFlowExecution) GetStepsOk() (*[]FlowExecutionStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *FlowFlowExecution) SetSteps(v []FlowExecutionStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *FlowFlowExecution) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetStartedAt

`func (o *FlowFlowExecution) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *FlowFlowExecution) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *FlowFlowExecution) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *FlowFlowExecution) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetFinishedAt

`func (o *FlowFlowExecution) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *FlowFlowExecution) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *FlowFlowExecution) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *FlowFlowExecution) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


