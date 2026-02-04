# \AuthServiceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthServiceAddTeamMember**](AuthServiceAPI.md#AuthServiceAddTeamMember) | **Post** /v1/auth/teams/members | 
[**AuthServiceCreateSSOProvider**](AuthServiceAPI.md#AuthServiceCreateSSOProvider) | **Post** /v1/auth/sso/providers | 
[**AuthServiceGetAuditLogs**](AuthServiceAPI.md#AuthServiceGetAuditLogs) | **Get** /v1/auth/audit_logs | 
[**AuthServiceGetSSOProvider**](AuthServiceAPI.md#AuthServiceGetSSOProvider) | **Get** /v1/auth/sso/providers/{id} | 
[**AuthServiceInitiateSSO**](AuthServiceAPI.md#AuthServiceInitiateSSO) | **Post** /v1/auth/sso/initiate | 
[**AuthServiceListTeamMembers**](AuthServiceAPI.md#AuthServiceListTeamMembers) | **Get** /v1/auth/teams/{orgId}/members | 
[**AuthServiceRemoveTeamMember**](AuthServiceAPI.md#AuthServiceRemoveTeamMember) | **Delete** /v1/auth/teams/members | 
[**AuthServiceValidateKey**](AuthServiceAPI.md#AuthServiceValidateKey) | **Post** /v1/auth/validate | 
[**AuthServiceValidateToken**](AuthServiceAPI.md#AuthServiceValidateToken) | **Post** /v1/auth/validate_token | 



## AuthServiceAddTeamMember

> AuthMembership AuthServiceAddTeamMember(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sapliy/fintech-sdk-go"
)

func main() {
	body := *openapiclient.NewAuthAddTeamMemberRequest() // AuthAddTeamMemberRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthServiceAPI.AuthServiceAddTeamMember(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthServiceAPI.AuthServiceAddTeamMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthServiceAddTeamMember`: AuthMembership
	fmt.Fprintf(os.Stdout, "Response from `AuthServiceAPI.AuthServiceAddTeamMember`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthServiceAddTeamMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthAddTeamMemberRequest**](AuthAddTeamMemberRequest.md) |  | 

### Return type

[**AuthMembership**](AuthMembership.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthServiceCreateSSOProvider

> AuthSSOProvider AuthServiceCreateSSOProvider(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sapliy/fintech-sdk-go"
)

func main() {
	body := *openapiclient.NewAuthCreateSSOProviderRequest() // AuthCreateSSOProviderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthServiceAPI.AuthServiceCreateSSOProvider(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthServiceAPI.AuthServiceCreateSSOProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthServiceCreateSSOProvider`: AuthSSOProvider
	fmt.Fprintf(os.Stdout, "Response from `AuthServiceAPI.AuthServiceCreateSSOProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthServiceCreateSSOProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthCreateSSOProviderRequest**](AuthCreateSSOProviderRequest.md) |  | 

### Return type

[**AuthSSOProvider**](AuthSSOProvider.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthServiceGetAuditLogs

> AuthGetAuditLogsResponse AuthServiceGetAuditLogs(ctx).OrgId(orgId).Limit(limit).Offset(offset).Action(action).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sapliy/fintech-sdk-go"
)

func main() {
	orgId := "orgId_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)
	action := "action_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthServiceAPI.AuthServiceGetAuditLogs(context.Background()).OrgId(orgId).Limit(limit).Offset(offset).Action(action).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthServiceAPI.AuthServiceGetAuditLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthServiceGetAuditLogs`: AuthGetAuditLogsResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthServiceAPI.AuthServiceGetAuditLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthServiceGetAuditLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | **string** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **action** | **string** |  | 

### Return type

[**AuthGetAuditLogsResponse**](AuthGetAuditLogsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthServiceGetSSOProvider

> AuthSSOProvider AuthServiceGetSSOProvider(ctx, id).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sapliy/fintech-sdk-go"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthServiceAPI.AuthServiceGetSSOProvider(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthServiceAPI.AuthServiceGetSSOProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthServiceGetSSOProvider`: AuthSSOProvider
	fmt.Fprintf(os.Stdout, "Response from `AuthServiceAPI.AuthServiceGetSSOProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthServiceGetSSOProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthSSOProvider**](AuthSSOProvider.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthServiceInitiateSSO

> AuthInitiateSSOResponse AuthServiceInitiateSSO(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sapliy/fintech-sdk-go"
)

func main() {
	body := *openapiclient.NewAuthInitiateSSORequest() // AuthInitiateSSORequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthServiceAPI.AuthServiceInitiateSSO(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthServiceAPI.AuthServiceInitiateSSO``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthServiceInitiateSSO`: AuthInitiateSSOResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthServiceAPI.AuthServiceInitiateSSO`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthServiceInitiateSSORequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthInitiateSSORequest**](AuthInitiateSSORequest.md) |  | 

### Return type

[**AuthInitiateSSOResponse**](AuthInitiateSSOResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthServiceListTeamMembers

> AuthListTeamMembersResponse AuthServiceListTeamMembers(ctx, orgId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sapliy/fintech-sdk-go"
)

func main() {
	orgId := "orgId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthServiceAPI.AuthServiceListTeamMembers(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthServiceAPI.AuthServiceListTeamMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthServiceListTeamMembers`: AuthListTeamMembersResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthServiceAPI.AuthServiceListTeamMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthServiceListTeamMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthListTeamMembersResponse**](AuthListTeamMembersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthServiceRemoveTeamMember

> AuthRemoveTeamMemberResponse AuthServiceRemoveTeamMember(ctx).OrgId(orgId).UserId(userId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sapliy/fintech-sdk-go"
)

func main() {
	orgId := "orgId_example" // string |  (optional)
	userId := "userId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthServiceAPI.AuthServiceRemoveTeamMember(context.Background()).OrgId(orgId).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthServiceAPI.AuthServiceRemoveTeamMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthServiceRemoveTeamMember`: AuthRemoveTeamMemberResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthServiceAPI.AuthServiceRemoveTeamMember`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthServiceRemoveTeamMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | **string** |  | 
 **userId** | **string** |  | 

### Return type

[**AuthRemoveTeamMemberResponse**](AuthRemoveTeamMemberResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthServiceValidateKey

> AuthValidateKeyResponse AuthServiceValidateKey(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sapliy/fintech-sdk-go"
)

func main() {
	body := *openapiclient.NewAuthValidateKeyRequest() // AuthValidateKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthServiceAPI.AuthServiceValidateKey(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthServiceAPI.AuthServiceValidateKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthServiceValidateKey`: AuthValidateKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthServiceAPI.AuthServiceValidateKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthServiceValidateKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthValidateKeyRequest**](AuthValidateKeyRequest.md) |  | 

### Return type

[**AuthValidateKeyResponse**](AuthValidateKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthServiceValidateToken

> AuthValidateTokenResponse AuthServiceValidateToken(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sapliy/fintech-sdk-go"
)

func main() {
	body := *openapiclient.NewAuthValidateTokenRequest() // AuthValidateTokenRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthServiceAPI.AuthServiceValidateToken(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthServiceAPI.AuthServiceValidateToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthServiceValidateToken`: AuthValidateTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthServiceAPI.AuthServiceValidateToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthServiceValidateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthValidateTokenRequest**](AuthValidateTokenRequest.md) |  | 

### Return type

[**AuthValidateTokenResponse**](AuthValidateTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

