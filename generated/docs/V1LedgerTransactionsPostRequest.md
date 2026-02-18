# V1LedgerTransactionsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceId** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Entries** | [**[]LedgerEntry**](LedgerEntry.md) |  | 

## Methods

### NewV1LedgerTransactionsPostRequest

`func NewV1LedgerTransactionsPostRequest(referenceId string, entries []LedgerEntry, ) *V1LedgerTransactionsPostRequest`

NewV1LedgerTransactionsPostRequest instantiates a new V1LedgerTransactionsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1LedgerTransactionsPostRequestWithDefaults

`func NewV1LedgerTransactionsPostRequestWithDefaults() *V1LedgerTransactionsPostRequest`

NewV1LedgerTransactionsPostRequestWithDefaults instantiates a new V1LedgerTransactionsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceId

`func (o *V1LedgerTransactionsPostRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *V1LedgerTransactionsPostRequest) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *V1LedgerTransactionsPostRequest) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetDescription

`func (o *V1LedgerTransactionsPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1LedgerTransactionsPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1LedgerTransactionsPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1LedgerTransactionsPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntries

`func (o *V1LedgerTransactionsPostRequest) GetEntries() []LedgerEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *V1LedgerTransactionsPostRequest) GetEntriesOk() (*[]LedgerEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *V1LedgerTransactionsPostRequest) SetEntries(v []LedgerEntry)`

SetEntries sets Entries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


