# Node

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | **string** |  | 
**Position** | Pointer to **map[string]interface{}** |  | [optional] 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewNode

`func NewNode(id string, type_ string, ) *Node`

NewNode instantiates a new Node object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeWithDefaults

`func NewNodeWithDefaults() *Node`

NewNodeWithDefaults instantiates a new Node object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Node) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Node) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Node) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *Node) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Node) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Node) SetType(v string)`

SetType sets Type field to given value.


### GetPosition

`func (o *Node) GetPosition() map[string]interface{}`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *Node) GetPositionOk() (*map[string]interface{}, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *Node) SetPosition(v map[string]interface{})`

SetPosition sets Position field to given value.

### HasPosition

`func (o *Node) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetData

`func (o *Node) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Node) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Node) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *Node) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


