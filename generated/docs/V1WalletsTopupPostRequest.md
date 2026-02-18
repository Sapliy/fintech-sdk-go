# V1WalletsTopupPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** |  | 
**Currency** | **string** |  | 
**ReferenceId** | **string** |  | 

## Methods

### NewV1WalletsTopupPostRequest

`func NewV1WalletsTopupPostRequest(amount int64, currency string, referenceId string, ) *V1WalletsTopupPostRequest`

NewV1WalletsTopupPostRequest instantiates a new V1WalletsTopupPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1WalletsTopupPostRequestWithDefaults

`func NewV1WalletsTopupPostRequestWithDefaults() *V1WalletsTopupPostRequest`

NewV1WalletsTopupPostRequestWithDefaults instantiates a new V1WalletsTopupPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *V1WalletsTopupPostRequest) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *V1WalletsTopupPostRequest) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *V1WalletsTopupPostRequest) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *V1WalletsTopupPostRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *V1WalletsTopupPostRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *V1WalletsTopupPostRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetReferenceId

`func (o *V1WalletsTopupPostRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *V1WalletsTopupPostRequest) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *V1WalletsTopupPostRequest) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


