# AuthGetAuditLogsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | Pointer to [**[]AuthAuditLog**](AuthAuditLog.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewAuthGetAuditLogsResponse

`func NewAuthGetAuditLogsResponse() *AuthGetAuditLogsResponse`

NewAuthGetAuditLogsResponse instantiates a new AuthGetAuditLogsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthGetAuditLogsResponseWithDefaults

`func NewAuthGetAuditLogsResponseWithDefaults() *AuthGetAuditLogsResponse`

NewAuthGetAuditLogsResponseWithDefaults instantiates a new AuthGetAuditLogsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *AuthGetAuditLogsResponse) GetLogs() []AuthAuditLog`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *AuthGetAuditLogsResponse) GetLogsOk() (*[]AuthAuditLog, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *AuthGetAuditLogsResponse) SetLogs(v []AuthAuditLog)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *AuthGetAuditLogsResponse) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetTotalCount

`func (o *AuthGetAuditLogsResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *AuthGetAuditLogsResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *AuthGetAuditLogsResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *AuthGetAuditLogsResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


