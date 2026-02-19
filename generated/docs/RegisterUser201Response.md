# RegisterUser201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 

## Methods

### NewRegisterUser201Response

`func NewRegisterUser201Response() *RegisterUser201Response`

NewRegisterUser201Response instantiates a new RegisterUser201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterUser201ResponseWithDefaults

`func NewRegisterUser201ResponseWithDefaults() *RegisterUser201Response`

NewRegisterUser201ResponseWithDefaults instantiates a new RegisterUser201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *RegisterUser201Response) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RegisterUser201Response) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RegisterUser201Response) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RegisterUser201Response) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUser

`func (o *RegisterUser201Response) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RegisterUser201Response) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RegisterUser201Response) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *RegisterUser201Response) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


