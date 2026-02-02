# AuthAddTeamMemberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthAddTeamMemberRequest

`func NewAuthAddTeamMemberRequest() *AuthAddTeamMemberRequest`

NewAuthAddTeamMemberRequest instantiates a new AuthAddTeamMemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthAddTeamMemberRequestWithDefaults

`func NewAuthAddTeamMemberRequestWithDefaults() *AuthAddTeamMemberRequest`

NewAuthAddTeamMemberRequestWithDefaults instantiates a new AuthAddTeamMemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *AuthAddTeamMemberRequest) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *AuthAddTeamMemberRequest) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *AuthAddTeamMemberRequest) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *AuthAddTeamMemberRequest) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetUserId

`func (o *AuthAddTeamMemberRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AuthAddTeamMemberRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AuthAddTeamMemberRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AuthAddTeamMemberRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetRole

`func (o *AuthAddTeamMemberRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AuthAddTeamMemberRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AuthAddTeamMemberRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AuthAddTeamMemberRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


