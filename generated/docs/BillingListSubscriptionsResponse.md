# BillingListSubscriptionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscriptions** | Pointer to [**[]BillingSubscription**](BillingSubscription.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewBillingListSubscriptionsResponse

`func NewBillingListSubscriptionsResponse() *BillingListSubscriptionsResponse`

NewBillingListSubscriptionsResponse instantiates a new BillingListSubscriptionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingListSubscriptionsResponseWithDefaults

`func NewBillingListSubscriptionsResponseWithDefaults() *BillingListSubscriptionsResponse`

NewBillingListSubscriptionsResponseWithDefaults instantiates a new BillingListSubscriptionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptions

`func (o *BillingListSubscriptionsResponse) GetSubscriptions() []BillingSubscription`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *BillingListSubscriptionsResponse) GetSubscriptionsOk() (*[]BillingSubscription, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *BillingListSubscriptionsResponse) SetSubscriptions(v []BillingSubscription)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *BillingListSubscriptionsResponse) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetTotalCount

`func (o *BillingListSubscriptionsResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *BillingListSubscriptionsResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *BillingListSubscriptionsResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *BillingListSubscriptionsResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


