# TransferWalletRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToUserId** | **string** |  | 
**Amount** | **int64** |  | 
**Currency** | **string** |  | 
**ReferenceId** | **string** |  | 

## Methods

### NewTransferWalletRequest

`func NewTransferWalletRequest(toUserId string, amount int64, currency string, referenceId string, ) *TransferWalletRequest`

NewTransferWalletRequest instantiates a new TransferWalletRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferWalletRequestWithDefaults

`func NewTransferWalletRequestWithDefaults() *TransferWalletRequest`

NewTransferWalletRequestWithDefaults instantiates a new TransferWalletRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToUserId

`func (o *TransferWalletRequest) GetToUserId() string`

GetToUserId returns the ToUserId field if non-nil, zero value otherwise.

### GetToUserIdOk

`func (o *TransferWalletRequest) GetToUserIdOk() (*string, bool)`

GetToUserIdOk returns a tuple with the ToUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToUserId

`func (o *TransferWalletRequest) SetToUserId(v string)`

SetToUserId sets ToUserId field to given value.


### GetAmount

`func (o *TransferWalletRequest) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferWalletRequest) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferWalletRequest) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *TransferWalletRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TransferWalletRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TransferWalletRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetReferenceId

`func (o *TransferWalletRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *TransferWalletRequest) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *TransferWalletRequest) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


