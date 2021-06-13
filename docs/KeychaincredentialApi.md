# \KeychaincredentialApi

All URIs are relative to *http://nas.local.geekspace.us/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KeychaincredentialGenerateSshKeyPairGet**](KeychaincredentialApi.md#KeychaincredentialGenerateSshKeyPairGet) | **Get** /keychaincredential/generate_ssh_key_pair | 
[**KeychaincredentialGet**](KeychaincredentialApi.md#KeychaincredentialGet) | **Get** /keychaincredential | 
[**KeychaincredentialIdIdDelete**](KeychaincredentialApi.md#KeychaincredentialIdIdDelete) | **Delete** /keychaincredential/id/{id} | 
[**KeychaincredentialIdIdGet**](KeychaincredentialApi.md#KeychaincredentialIdIdGet) | **Get** /keychaincredential/id/{id} | 
[**KeychaincredentialIdIdPut**](KeychaincredentialApi.md#KeychaincredentialIdIdPut) | **Put** /keychaincredential/id/{id} | 
[**KeychaincredentialPost**](KeychaincredentialApi.md#KeychaincredentialPost) | **Post** /keychaincredential | 
[**KeychaincredentialRemoteSshHostKeyScanPost**](KeychaincredentialApi.md#KeychaincredentialRemoteSshHostKeyScanPost) | **Post** /keychaincredential/remote_ssh_host_key_scan | 
[**KeychaincredentialRemoteSshSemiautomaticSetupPost**](KeychaincredentialApi.md#KeychaincredentialRemoteSshSemiautomaticSetupPost) | **Post** /keychaincredential/remote_ssh_semiautomatic_setup | 
[**KeychaincredentialUsedByPost**](KeychaincredentialApi.md#KeychaincredentialUsedByPost) | **Post** /keychaincredential/used_by | 



## KeychaincredentialGenerateSshKeyPairGet

> KeychaincredentialGenerateSshKeyPairGet(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeychaincredentialApi.KeychaincredentialGenerateSshKeyPairGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeychaincredentialApi.KeychaincredentialGenerateSshKeyPairGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiKeychaincredentialGenerateSshKeyPairGetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeychaincredentialGet

> KeychaincredentialGet(ctx).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    count := true // bool |  (optional)
    sort := "sort_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeychaincredentialApi.KeychaincredentialGet(context.Background()).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeychaincredentialApi.KeychaincredentialGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeychaincredentialGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **count** | **bool** |  | 
 **sort** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeychaincredentialIdIdDelete

> KeychaincredentialIdIdDelete(ctx, id).KeychaincredentialDelete1(keychaincredentialDelete1).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | 
    keychaincredentialDelete1 := *openapiclient.NewKeychaincredentialDelete1() // KeychaincredentialDelete1 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeychaincredentialApi.KeychaincredentialIdIdDelete(context.Background(), id).KeychaincredentialDelete1(keychaincredentialDelete1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeychaincredentialApi.KeychaincredentialIdIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKeychaincredentialIdIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **keychaincredentialDelete1** | [**KeychaincredentialDelete1**](KeychaincredentialDelete1.md) |  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeychaincredentialIdIdGet

> KeychaincredentialIdIdGet(ctx, id).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := []interface{}{interface{}(123)} // []interface{} |  (default to [])
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    count := true // bool |  (optional)
    sort := "sort_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeychaincredentialApi.KeychaincredentialIdIdGet(context.Background(), id).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeychaincredentialApi.KeychaincredentialIdIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**[]interface{}**](interface{}.md) |  | [default to []]

### Other Parameters

Other parameters are passed through a pointer to a apiKeychaincredentialIdIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **count** | **bool** |  | 
 **sort** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeychaincredentialIdIdPut

> KeychaincredentialIdIdPut(ctx, id).KeychaincredentialUpdate1(keychaincredentialUpdate1).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | 
    keychaincredentialUpdate1 := *openapiclient.NewKeychaincredentialUpdate1() // KeychaincredentialUpdate1 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeychaincredentialApi.KeychaincredentialIdIdPut(context.Background(), id).KeychaincredentialUpdate1(keychaincredentialUpdate1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeychaincredentialApi.KeychaincredentialIdIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKeychaincredentialIdIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **keychaincredentialUpdate1** | [**KeychaincredentialUpdate1**](KeychaincredentialUpdate1.md) |  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeychaincredentialPost

> KeychaincredentialPost(ctx).KeychaincredentialCreate0(keychaincredentialCreate0).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    keychaincredentialCreate0 := *openapiclient.NewKeychaincredentialCreate0() // KeychaincredentialCreate0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeychaincredentialApi.KeychaincredentialPost(context.Background()).KeychaincredentialCreate0(keychaincredentialCreate0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeychaincredentialApi.KeychaincredentialPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeychaincredentialPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keychaincredentialCreate0** | [**KeychaincredentialCreate0**](KeychaincredentialCreate0.md) |  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeychaincredentialRemoteSshHostKeyScanPost

> KeychaincredentialRemoteSshHostKeyScanPost(ctx).KeychaincredentialRemoteSshHostKeyScan0(keychaincredentialRemoteSshHostKeyScan0).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    keychaincredentialRemoteSshHostKeyScan0 := *openapiclient.NewKeychaincredentialRemoteSshHostKeyScan0() // KeychaincredentialRemoteSshHostKeyScan0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeychaincredentialApi.KeychaincredentialRemoteSshHostKeyScanPost(context.Background()).KeychaincredentialRemoteSshHostKeyScan0(keychaincredentialRemoteSshHostKeyScan0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeychaincredentialApi.KeychaincredentialRemoteSshHostKeyScanPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeychaincredentialRemoteSshHostKeyScanPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keychaincredentialRemoteSshHostKeyScan0** | [**KeychaincredentialRemoteSshHostKeyScan0**](KeychaincredentialRemoteSshHostKeyScan0.md) |  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeychaincredentialRemoteSshSemiautomaticSetupPost

> KeychaincredentialRemoteSshSemiautomaticSetupPost(ctx).KeychaincredentialRemoteSshSemiautomaticSetup0(keychaincredentialRemoteSshSemiautomaticSetup0).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    keychaincredentialRemoteSshSemiautomaticSetup0 := *openapiclient.NewKeychaincredentialRemoteSshSemiautomaticSetup0() // KeychaincredentialRemoteSshSemiautomaticSetup0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeychaincredentialApi.KeychaincredentialRemoteSshSemiautomaticSetupPost(context.Background()).KeychaincredentialRemoteSshSemiautomaticSetup0(keychaincredentialRemoteSshSemiautomaticSetup0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeychaincredentialApi.KeychaincredentialRemoteSshSemiautomaticSetupPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeychaincredentialRemoteSshSemiautomaticSetupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keychaincredentialRemoteSshSemiautomaticSetup0** | [**KeychaincredentialRemoteSshSemiautomaticSetup0**](KeychaincredentialRemoteSshSemiautomaticSetup0.md) |  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeychaincredentialUsedByPost

> KeychaincredentialUsedByPost(ctx).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KeychaincredentialApi.KeychaincredentialUsedByPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeychaincredentialApi.KeychaincredentialUsedByPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeychaincredentialUsedByPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **int32** |  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

