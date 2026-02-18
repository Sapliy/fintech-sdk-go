# Flow

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
**Trigger** | Pointer to [**Trigger**](Trigger.md) |  | [optional] 
**Nodes** | Pointer to [**[]Node**](Node.md) |  | [optional] 
**Edges** | Pointer to [**[]Edge**](Edge.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewFlow

`func NewFlow(id string, zoneId string, name string, ) *Flow`

NewFlow instantiates a new Flow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowWithDefaults

`func NewFlowWithDefaults() *Flow`

NewFlowWithDefaults instantiates a new Flow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Flow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Flow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Flow) SetId(v string)`

SetId sets Id field to given value.


### GetOrgId

`func (o *Flow) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Flow) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Flow) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *Flow) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetZoneId

`func (o *Flow) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *Flow) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *Flow) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.


### GetName

`func (o *Flow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Flow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Flow) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Flow) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Flow) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Flow) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Flow) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *Flow) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Flow) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Flow) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Flow) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetVersion

`func (o *Flow) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Flow) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Flow) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Flow) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetTrigger

`func (o *Flow) GetTrigger() Trigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *Flow) GetTriggerOk() (*Trigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *Flow) SetTrigger(v Trigger)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *Flow) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetNodes

`func (o *Flow) GetNodes() []Node`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *Flow) GetNodesOk() (*[]Node, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *Flow) SetNodes(v []Node)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *Flow) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetEdges

`func (o *Flow) GetEdges() []Edge`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *Flow) GetEdgesOk() (*[]Edge, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *Flow) SetEdges(v []Edge)`

SetEdges sets Edges field to given value.

### HasEdges

`func (o *Flow) HasEdges() bool`

HasEdges returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Flow) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Flow) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Flow) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Flow) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Flow) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Flow) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Flow) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Flow) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


