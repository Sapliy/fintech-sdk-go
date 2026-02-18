# V1LedgerAccountsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | **string** |  | 
**Currency** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewV1LedgerAccountsPostRequest

`func NewV1LedgerAccountsPostRequest(name string, type_ string, ) *V1LedgerAccountsPostRequest`

NewV1LedgerAccountsPostRequest instantiates a new V1LedgerAccountsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1LedgerAccountsPostRequestWithDefaults

`func NewV1LedgerAccountsPostRequestWithDefaults() *V1LedgerAccountsPostRequest`

NewV1LedgerAccountsPostRequestWithDefaults instantiates a new V1LedgerAccountsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V1LedgerAccountsPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1LedgerAccountsPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1LedgerAccountsPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *V1LedgerAccountsPostRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1LedgerAccountsPostRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1LedgerAccountsPostRequest) SetType(v string)`

SetType sets Type field to given value.


### GetCurrency

`func (o *V1LedgerAccountsPostRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *V1LedgerAccountsPostRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *V1LedgerAccountsPostRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *V1LedgerAccountsPostRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetUserId

`func (o *V1LedgerAccountsPostRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *V1LedgerAccountsPostRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *V1LedgerAccountsPostRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *V1LedgerAccountsPostRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


