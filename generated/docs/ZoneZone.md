# ZoneZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**SecretKey** | Pointer to **string** |  | [optional] 
**PublishableKey** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TemplateName** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewZoneZone

`func NewZoneZone() *ZoneZone`

NewZoneZone instantiates a new ZoneZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneZoneWithDefaults

`func NewZoneZoneWithDefaults() *ZoneZone`

NewZoneZoneWithDefaults instantiates a new ZoneZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ZoneZone) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ZoneZone) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ZoneZone) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ZoneZone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrgId

`func (o *ZoneZone) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ZoneZone) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ZoneZone) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ZoneZone) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetName

`func (o *ZoneZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ZoneZone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ZoneZone) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ZoneZone) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMode

`func (o *ZoneZone) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ZoneZone) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ZoneZone) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ZoneZone) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetSecretKey

`func (o *ZoneZone) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *ZoneZone) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *ZoneZone) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *ZoneZone) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetPublishableKey

`func (o *ZoneZone) GetPublishableKey() string`

GetPublishableKey returns the PublishableKey field if non-nil, zero value otherwise.

### GetPublishableKeyOk

`func (o *ZoneZone) GetPublishableKeyOk() (*string, bool)`

GetPublishableKeyOk returns a tuple with the PublishableKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishableKey

`func (o *ZoneZone) SetPublishableKey(v string)`

SetPublishableKey sets PublishableKey field to given value.

### HasPublishableKey

`func (o *ZoneZone) HasPublishableKey() bool`

HasPublishableKey returns a boolean if a field has been set.

### GetStatus

`func (o *ZoneZone) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ZoneZone) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ZoneZone) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ZoneZone) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTemplateName

`func (o *ZoneZone) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *ZoneZone) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *ZoneZone) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *ZoneZone) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### GetMetadata

`func (o *ZoneZone) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ZoneZone) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ZoneZone) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ZoneZone) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ZoneZone) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ZoneZone) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ZoneZone) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ZoneZone) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


