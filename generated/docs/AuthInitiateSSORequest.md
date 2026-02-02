# AuthInitiateSSORequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**RedirectUri** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthInitiateSSORequest

`func NewAuthInitiateSSORequest() *AuthInitiateSSORequest`

NewAuthInitiateSSORequest instantiates a new AuthInitiateSSORequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthInitiateSSORequestWithDefaults

`func NewAuthInitiateSSORequestWithDefaults() *AuthInitiateSSORequest`

NewAuthInitiateSSORequestWithDefaults instantiates a new AuthInitiateSSORequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *AuthInitiateSSORequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AuthInitiateSSORequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AuthInitiateSSORequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AuthInitiateSSORequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetRedirectUri

`func (o *AuthInitiateSSORequest) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *AuthInitiateSSORequest) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *AuthInitiateSSORequest) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *AuthInitiateSSORequest) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


