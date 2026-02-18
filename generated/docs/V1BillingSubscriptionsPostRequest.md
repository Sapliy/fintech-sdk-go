# V1BillingSubscriptionsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanId** | **string** |  | 
**CustomerId** | Pointer to **string** |  | [optional] 

## Methods

### NewV1BillingSubscriptionsPostRequest

`func NewV1BillingSubscriptionsPostRequest(planId string, ) *V1BillingSubscriptionsPostRequest`

NewV1BillingSubscriptionsPostRequest instantiates a new V1BillingSubscriptionsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1BillingSubscriptionsPostRequestWithDefaults

`func NewV1BillingSubscriptionsPostRequestWithDefaults() *V1BillingSubscriptionsPostRequest`

NewV1BillingSubscriptionsPostRequestWithDefaults instantiates a new V1BillingSubscriptionsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanId

`func (o *V1BillingSubscriptionsPostRequest) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *V1BillingSubscriptionsPostRequest) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *V1BillingSubscriptionsPostRequest) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetCustomerId

`func (o *V1BillingSubscriptionsPostRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *V1BillingSubscriptionsPostRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *V1BillingSubscriptionsPostRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *V1BillingSubscriptionsPostRequest) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


