# AuthListTeamMembersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Memberships** | Pointer to [**[]AuthMembership**](AuthMembership.md) |  | [optional] 

## Methods

### NewAuthListTeamMembersResponse

`func NewAuthListTeamMembersResponse() *AuthListTeamMembersResponse`

NewAuthListTeamMembersResponse instantiates a new AuthListTeamMembersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthListTeamMembersResponseWithDefaults

`func NewAuthListTeamMembersResponseWithDefaults() *AuthListTeamMembersResponse`

NewAuthListTeamMembersResponseWithDefaults instantiates a new AuthListTeamMembersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberships

`func (o *AuthListTeamMembersResponse) GetMemberships() []AuthMembership`

GetMemberships returns the Memberships field if non-nil, zero value otherwise.

### GetMembershipsOk

`func (o *AuthListTeamMembersResponse) GetMembershipsOk() (*[]AuthMembership, bool)`

GetMembershipsOk returns a tuple with the Memberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberships

`func (o *AuthListTeamMembersResponse) SetMemberships(v []AuthMembership)`

SetMemberships sets Memberships field to given value.

### HasMemberships

`func (o *AuthListTeamMembersResponse) HasMemberships() bool`

HasMemberships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


