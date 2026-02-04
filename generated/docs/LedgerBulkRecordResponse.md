# LedgerBulkRecordResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Responses** | Pointer to [**[]LedgerRecordTransactionResponse**](LedgerRecordTransactionResponse.md) |  | [optional] 

## Methods

### NewLedgerBulkRecordResponse

`func NewLedgerBulkRecordResponse() *LedgerBulkRecordResponse`

NewLedgerBulkRecordResponse instantiates a new LedgerBulkRecordResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLedgerBulkRecordResponseWithDefaults

`func NewLedgerBulkRecordResponseWithDefaults() *LedgerBulkRecordResponse`

NewLedgerBulkRecordResponseWithDefaults instantiates a new LedgerBulkRecordResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponses

`func (o *LedgerBulkRecordResponse) GetResponses() []LedgerRecordTransactionResponse`

GetResponses returns the Responses field if non-nil, zero value otherwise.

### GetResponsesOk

`func (o *LedgerBulkRecordResponse) GetResponsesOk() (*[]LedgerRecordTransactionResponse, bool)`

GetResponsesOk returns a tuple with the Responses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponses

`func (o *LedgerBulkRecordResponse) SetResponses(v []LedgerRecordTransactionResponse)`

SetResponses sets Responses field to given value.

### HasResponses

`func (o *LedgerBulkRecordResponse) HasResponses() bool`

HasResponses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


