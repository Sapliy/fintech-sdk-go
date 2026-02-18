# LedgerAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Type** | **string** |  | 
**Currency** | **string** |  | 
**Balance** | Pointer to **int64** |  | [optional] 

## Methods

### NewLedgerAccount

`func NewLedgerAccount(id string, name string, type_ string, currency string, ) *LedgerAccount`

NewLedgerAccount instantiates a new LedgerAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLedgerAccountWithDefaults

`func NewLedgerAccountWithDefaults() *LedgerAccount`

NewLedgerAccountWithDefaults instantiates a new LedgerAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LedgerAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LedgerAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LedgerAccount) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *LedgerAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LedgerAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LedgerAccount) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *LedgerAccount) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LedgerAccount) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LedgerAccount) SetType(v string)`

SetType sets Type field to given value.


### GetCurrency

`func (o *LedgerAccount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *LedgerAccount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *LedgerAccount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetBalance

`func (o *LedgerAccount) GetBalance() int64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *LedgerAccount) GetBalanceOk() (*int64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *LedgerAccount) SetBalance(v int64)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *LedgerAccount) HasBalance() bool`

HasBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


