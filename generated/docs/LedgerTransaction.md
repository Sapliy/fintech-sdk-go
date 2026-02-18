# LedgerTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ReferenceId** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**Entries** | Pointer to [**[]LedgerEntry**](LedgerEntry.md) |  | [optional] 

## Methods

### NewLedgerTransaction

`func NewLedgerTransaction(id string, referenceId string, status string, ) *LedgerTransaction`

NewLedgerTransaction instantiates a new LedgerTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLedgerTransactionWithDefaults

`func NewLedgerTransactionWithDefaults() *LedgerTransaction`

NewLedgerTransactionWithDefaults instantiates a new LedgerTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LedgerTransaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LedgerTransaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LedgerTransaction) SetId(v string)`

SetId sets Id field to given value.


### GetReferenceId

`func (o *LedgerTransaction) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *LedgerTransaction) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *LedgerTransaction) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetDescription

`func (o *LedgerTransaction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LedgerTransaction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LedgerTransaction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LedgerTransaction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *LedgerTransaction) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LedgerTransaction) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LedgerTransaction) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetEntries

`func (o *LedgerTransaction) GetEntries() []LedgerEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *LedgerTransaction) GetEntriesOk() (*[]LedgerEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *LedgerTransaction) SetEntries(v []LedgerEntry)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *LedgerTransaction) HasEntries() bool`

HasEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


