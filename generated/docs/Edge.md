# Edge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Source** | **string** |  | 
**Target** | **string** |  | 
**SourceHandle** | Pointer to **string** |  | [optional] 

## Methods

### NewEdge

`func NewEdge(id string, source string, target string, ) *Edge`

NewEdge instantiates a new Edge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeWithDefaults

`func NewEdgeWithDefaults() *Edge`

NewEdgeWithDefaults instantiates a new Edge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Edge) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Edge) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Edge) SetId(v string)`

SetId sets Id field to given value.


### GetSource

`func (o *Edge) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Edge) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Edge) SetSource(v string)`

SetSource sets Source field to given value.


### GetTarget

`func (o *Edge) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *Edge) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *Edge) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetSourceHandle

`func (o *Edge) GetSourceHandle() string`

GetSourceHandle returns the SourceHandle field if non-nil, zero value otherwise.

### GetSourceHandleOk

`func (o *Edge) GetSourceHandleOk() (*string, bool)`

GetSourceHandleOk returns a tuple with the SourceHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceHandle

`func (o *Edge) SetSourceHandle(v string)`

SetSourceHandle sets SourceHandle field to given value.

### HasSourceHandle

`func (o *Edge) HasSourceHandle() bool`

HasSourceHandle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


