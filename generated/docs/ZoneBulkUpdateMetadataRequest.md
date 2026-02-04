# ZoneBulkUpdateMetadataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZoneIds** | Pointer to **[]string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewZoneBulkUpdateMetadataRequest

`func NewZoneBulkUpdateMetadataRequest() *ZoneBulkUpdateMetadataRequest`

NewZoneBulkUpdateMetadataRequest instantiates a new ZoneBulkUpdateMetadataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneBulkUpdateMetadataRequestWithDefaults

`func NewZoneBulkUpdateMetadataRequestWithDefaults() *ZoneBulkUpdateMetadataRequest`

NewZoneBulkUpdateMetadataRequestWithDefaults instantiates a new ZoneBulkUpdateMetadataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoneIds

`func (o *ZoneBulkUpdateMetadataRequest) GetZoneIds() []string`

GetZoneIds returns the ZoneIds field if non-nil, zero value otherwise.

### GetZoneIdsOk

`func (o *ZoneBulkUpdateMetadataRequest) GetZoneIdsOk() (*[]string, bool)`

GetZoneIdsOk returns a tuple with the ZoneIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneIds

`func (o *ZoneBulkUpdateMetadataRequest) SetZoneIds(v []string)`

SetZoneIds sets ZoneIds field to given value.

### HasZoneIds

`func (o *ZoneBulkUpdateMetadataRequest) HasZoneIds() bool`

HasZoneIds returns a boolean if a field has been set.

### GetMetadata

`func (o *ZoneBulkUpdateMetadataRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ZoneBulkUpdateMetadataRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ZoneBulkUpdateMetadataRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ZoneBulkUpdateMetadataRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


