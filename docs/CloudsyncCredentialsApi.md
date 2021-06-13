# \CloudsyncCredentialsApi

All URIs are relative to *http://nas.local.geekspace.us/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudsyncCredentialsGet**](CloudsyncCredentialsApi.md#CloudsyncCredentialsGet) | **Get** /cloudsync/credentials | 
[**CloudsyncCredentialsIdIdDelete**](CloudsyncCredentialsApi.md#CloudsyncCredentialsIdIdDelete) | **Delete** /cloudsync/credentials/id/{id} | 
[**CloudsyncCredentialsIdIdGet**](CloudsyncCredentialsApi.md#CloudsyncCredentialsIdIdGet) | **Get** /cloudsync/credentials/id/{id} | 
[**CloudsyncCredentialsIdIdPut**](CloudsyncCredentialsApi.md#CloudsyncCredentialsIdIdPut) | **Put** /cloudsync/credentials/id/{id} | 
[**CloudsyncCredentialsPost**](CloudsyncCredentialsApi.md#CloudsyncCredentialsPost) | **Post** /cloudsync/credentials | 
[**CloudsyncCredentialsVerifyPost**](CloudsyncCredentialsApi.md#CloudsyncCredentialsVerifyPost) | **Post** /cloudsync/credentials/verify | 



## CloudsyncCredentialsGet

> CloudsyncCredentialsGet(ctx).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()



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
    resp, r, err := api_client.CloudsyncCredentialsApi.CloudsyncCredentialsGet(context.Background()).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsyncCredentialsApi.CloudsyncCredentialsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudsyncCredentialsGetRequest struct via the builder pattern


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


## CloudsyncCredentialsIdIdDelete

> CloudsyncCredentialsIdIdDelete(ctx, id).Execute()





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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsyncCredentialsApi.CloudsyncCredentialsIdIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsyncCredentialsApi.CloudsyncCredentialsIdIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudsyncCredentialsIdIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## CloudsyncCredentialsIdIdGet

> CloudsyncCredentialsIdIdGet(ctx, id).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()



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
    resp, r, err := api_client.CloudsyncCredentialsApi.CloudsyncCredentialsIdIdGet(context.Background(), id).Limit(limit).Offset(offset).Count(count).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsyncCredentialsApi.CloudsyncCredentialsIdIdGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudsyncCredentialsIdIdGetRequest struct via the builder pattern


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


## CloudsyncCredentialsIdIdPut

> CloudsyncCredentialsIdIdPut(ctx, id).CloudsyncCredentialsUpdate1(cloudsyncCredentialsUpdate1).Execute()





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
    cloudsyncCredentialsUpdate1 := *openapiclient.NewCloudsyncCredentialsUpdate1() // CloudsyncCredentialsUpdate1 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsyncCredentialsApi.CloudsyncCredentialsIdIdPut(context.Background(), id).CloudsyncCredentialsUpdate1(cloudsyncCredentialsUpdate1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsyncCredentialsApi.CloudsyncCredentialsIdIdPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudsyncCredentialsIdIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudsyncCredentialsUpdate1** | [**CloudsyncCredentialsUpdate1**](CloudsyncCredentialsUpdate1.md) |  | 

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


## CloudsyncCredentialsPost

> CloudsyncCredentialsPost(ctx).CloudsyncCredentialsCreate0(cloudsyncCredentialsCreate0).Execute()





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
    cloudsyncCredentialsCreate0 := *openapiclient.NewCloudsyncCredentialsCreate0() // CloudsyncCredentialsCreate0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsyncCredentialsApi.CloudsyncCredentialsPost(context.Background()).CloudsyncCredentialsCreate0(cloudsyncCredentialsCreate0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsyncCredentialsApi.CloudsyncCredentialsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudsyncCredentialsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudsyncCredentialsCreate0** | [**CloudsyncCredentialsCreate0**](CloudsyncCredentialsCreate0.md) |  | 

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


## CloudsyncCredentialsVerifyPost

> CloudsyncCredentialsVerifyPost(ctx).CloudsyncCredentialsVerify0(cloudsyncCredentialsVerify0).Execute()





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
    cloudsyncCredentialsVerify0 := *openapiclient.NewCloudsyncCredentialsVerify0() // CloudsyncCredentialsVerify0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudsyncCredentialsApi.CloudsyncCredentialsVerifyPost(context.Background()).CloudsyncCredentialsVerify0(cloudsyncCredentialsVerify0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudsyncCredentialsApi.CloudsyncCredentialsVerifyPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudsyncCredentialsVerifyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudsyncCredentialsVerify0** | [**CloudsyncCredentialsVerify0**](CloudsyncCredentialsVerify0.md) |  | 

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

