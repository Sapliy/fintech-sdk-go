# AutomationFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**OrgId** | Pointer to **string** |  | [optional] 
**ZoneId** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**Trigger** | Pointer to [**AutomationFlowTrigger**](AutomationFlowTrigger.md) |  | [optional] 
**Nodes** | Pointer to [**[]AutomationFlowNode**](AutomationFlowNode.md) |  | [optional] 
**Edges** | Pointer to [**[]AutomationFlowEdge**](AutomationFlowEdge.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAutomationFlow

`func NewAutomationFlow(id string, zoneId string, name string, ) *AutomationFlow`

NewAutomationFlow instantiates a new AutomationFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationFlowWithDefaults

`func NewAutomationFlowWithDefaults() *AutomationFlow`

NewAutomationFlowWithDefaults instantiates a new AutomationFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutomationFlow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutomationFlow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutomationFlow) SetId(v string)`

SetId sets Id field to given value.


### GetOrgId

`func (o *AutomationFlow) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *AutomationFlow) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *AutomationFlow) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *AutomationFlow) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetZoneId

`func (o *AutomationFlow) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *AutomationFlow) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *AutomationFlow) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.


### GetName

`func (o *AutomationFlow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutomationFlow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutomationFlow) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AutomationFlow) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AutomationFlow) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AutomationFlow) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AutomationFlow) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AutomationFlow) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AutomationFlow) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AutomationFlow) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AutomationFlow) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetVersion

`func (o *AutomationFlow) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AutomationFlow) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AutomationFlow) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AutomationFlow) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetTrigger

`func (o *AutomationFlow) GetTrigger() AutomationFlowTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *AutomationFlow) GetTriggerOk() (*AutomationFlowTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *AutomationFlow) SetTrigger(v AutomationFlowTrigger)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *AutomationFlow) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetNodes

`func (o *AutomationFlow) GetNodes() []AutomationFlowNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *AutomationFlow) GetNodesOk() (*[]AutomationFlowNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *AutomationFlow) SetNodes(v []AutomationFlowNode)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *AutomationFlow) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetEdges

`func (o *AutomationFlow) GetEdges() []AutomationFlowEdge`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *AutomationFlow) GetEdgesOk() (*[]AutomationFlowEdge, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *AutomationFlow) SetEdges(v []AutomationFlowEdge)`

SetEdges sets Edges field to given value.

### HasEdges

`func (o *AutomationFlow) HasEdges() bool`

HasEdges returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AutomationFlow) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AutomationFlow) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AutomationFlow) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AutomationFlow) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AutomationFlow) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AutomationFlow) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AutomationFlow) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AutomationFlow) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


