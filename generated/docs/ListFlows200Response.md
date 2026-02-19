# ListFlows200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flows** | Pointer to [**[]AutomationFlow**](AutomationFlow.md) |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 

## Methods

### NewListFlows200Response

`func NewListFlows200Response() *ListFlows200Response`

NewListFlows200Response instantiates a new ListFlows200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFlows200ResponseWithDefaults

`func NewListFlows200ResponseWithDefaults() *ListFlows200Response`

NewListFlows200ResponseWithDefaults instantiates a new ListFlows200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlows

`func (o *ListFlows200Response) GetFlows() []AutomationFlow`

GetFlows returns the Flows field if non-nil, zero value otherwise.

### GetFlowsOk

`func (o *ListFlows200Response) GetFlowsOk() (*[]AutomationFlow, bool)`

GetFlowsOk returns a tuple with the Flows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlows

`func (o *ListFlows200Response) SetFlows(v []AutomationFlow)`

SetFlows sets Flows field to given value.

### HasFlows

`func (o *ListFlows200Response) HasFlows() bool`

HasFlows returns a boolean if a field has been set.

### GetCount

`func (o *ListFlows200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListFlows200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListFlows200Response) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListFlows200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


