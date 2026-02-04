# LedgerBulkRecordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transactions** | Pointer to [**[]LedgerRecordTransactionRequest**](LedgerRecordTransactionRequest.md) |  | [optional] 

## Methods

### NewLedgerBulkRecordRequest

`func NewLedgerBulkRecordRequest() *LedgerBulkRecordRequest`

NewLedgerBulkRecordRequest instantiates a new LedgerBulkRecordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLedgerBulkRecordRequestWithDefaults

`func NewLedgerBulkRecordRequestWithDefaults() *LedgerBulkRecordRequest`

NewLedgerBulkRecordRequestWithDefaults instantiates a new LedgerBulkRecordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactions

`func (o *LedgerBulkRecordRequest) GetTransactions() []LedgerRecordTransactionRequest`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *LedgerBulkRecordRequest) GetTransactionsOk() (*[]LedgerRecordTransactionRequest, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *LedgerBulkRecordRequest) SetTransactions(v []LedgerRecordTransactionRequest)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *LedgerBulkRecordRequest) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


