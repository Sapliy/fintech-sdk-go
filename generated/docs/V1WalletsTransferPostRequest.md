# V1WalletsTransferPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToUserId** | **string** |  | 
**Amount** | **int64** |  | 
**Currency** | **string** |  | 
**ReferenceId** | **string** |  | 

## Methods

### NewV1WalletsTransferPostRequest

`func NewV1WalletsTransferPostRequest(toUserId string, amount int64, currency string, referenceId string, ) *V1WalletsTransferPostRequest`

NewV1WalletsTransferPostRequest instantiates a new V1WalletsTransferPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1WalletsTransferPostRequestWithDefaults

`func NewV1WalletsTransferPostRequestWithDefaults() *V1WalletsTransferPostRequest`

NewV1WalletsTransferPostRequestWithDefaults instantiates a new V1WalletsTransferPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToUserId

`func (o *V1WalletsTransferPostRequest) GetToUserId() string`

GetToUserId returns the ToUserId field if non-nil, zero value otherwise.

### GetToUserIdOk

`func (o *V1WalletsTransferPostRequest) GetToUserIdOk() (*string, bool)`

GetToUserIdOk returns a tuple with the ToUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToUserId

`func (o *V1WalletsTransferPostRequest) SetToUserId(v string)`

SetToUserId sets ToUserId field to given value.


### GetAmount

`func (o *V1WalletsTransferPostRequest) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *V1WalletsTransferPostRequest) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *V1WalletsTransferPostRequest) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *V1WalletsTransferPostRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *V1WalletsTransferPostRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *V1WalletsTransferPostRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetReferenceId

`func (o *V1WalletsTransferPostRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *V1WalletsTransferPostRequest) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *V1WalletsTransferPostRequest) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


