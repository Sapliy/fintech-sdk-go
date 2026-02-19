# AutomationFlowTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**EventType** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewAutomationFlowTrigger

`func NewAutomationFlowTrigger(type_ string, ) *AutomationFlowTrigger`

NewAutomationFlowTrigger instantiates a new AutomationFlowTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationFlowTriggerWithDefaults

`func NewAutomationFlowTriggerWithDefaults() *AutomationFlowTrigger`

NewAutomationFlowTriggerWithDefaults instantiates a new AutomationFlowTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AutomationFlowTrigger) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AutomationFlowTrigger) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AutomationFlowTrigger) SetType(v string)`

SetType sets Type field to given value.


### GetEventType

`func (o *AutomationFlowTrigger) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *AutomationFlowTrigger) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *AutomationFlowTrigger) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *AutomationFlowTrigger) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetConfig

`func (o *AutomationFlowTrigger) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AutomationFlowTrigger) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AutomationFlowTrigger) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AutomationFlowTrigger) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


