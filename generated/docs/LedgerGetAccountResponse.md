# LedgerGetAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**Balance** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewLedgerGetAccountResponse

`func NewLedgerGetAccountResponse() *LedgerGetAccountResponse`

NewLedgerGetAccountResponse instantiates a new LedgerGetAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLedgerGetAccountResponseWithDefaults

`func NewLedgerGetAccountResponseWithDefaults() *LedgerGetAccountResponse`

NewLedgerGetAccountResponseWithDefaults instantiates a new LedgerGetAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *LedgerGetAccountResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *LedgerGetAccountResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *LedgerGetAccountResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *LedgerGetAccountResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetBalance

`func (o *LedgerGetAccountResponse) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *LedgerGetAccountResponse) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *LedgerGetAccountResponse) SetBalance(v string)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *LedgerGetAccountResponse) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetCurrency

`func (o *LedgerGetAccountResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *LedgerGetAccountResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *LedgerGetAccountResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *LedgerGetAccountResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LedgerGetAccountResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LedgerGetAccountResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LedgerGetAccountResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LedgerGetAccountResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


