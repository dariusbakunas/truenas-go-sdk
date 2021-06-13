# \FilesystemApi

All URIs are relative to *http://nas.local.geekspace.us/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilesystemAclIsTrivialPost**](FilesystemApi.md#FilesystemAclIsTrivialPost) | **Post** /filesystem/acl_is_trivial | 
[**FilesystemChownPost**](FilesystemApi.md#FilesystemChownPost) | **Post** /filesystem/chown | 
[**FilesystemDefaultAclChoicesGet**](FilesystemApi.md#FilesystemDefaultAclChoicesGet) | **Get** /filesystem/default_acl_choices | 
[**FilesystemGetDefaultAclPost**](FilesystemApi.md#FilesystemGetDefaultAclPost) | **Post** /filesystem/get_default_acl | 
[**FilesystemGetaclPost**](FilesystemApi.md#FilesystemGetaclPost) | **Post** /filesystem/getacl | 
[**FilesystemListdirPost**](FilesystemApi.md#FilesystemListdirPost) | **Post** /filesystem/listdir | 
[**FilesystemSetaclPost**](FilesystemApi.md#FilesystemSetaclPost) | **Post** /filesystem/setacl | 
[**FilesystemSetpermPost**](FilesystemApi.md#FilesystemSetpermPost) | **Post** /filesystem/setperm | 
[**FilesystemStatPost**](FilesystemApi.md#FilesystemStatPost) | **Post** /filesystem/stat | 
[**FilesystemStatfsPost**](FilesystemApi.md#FilesystemStatfsPost) | **Post** /filesystem/statfs | 



## FilesystemAclIsTrivialPost

> FilesystemAclIsTrivialPost(ctx).Body(body).Execute()





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
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FilesystemApi.FilesystemAclIsTrivialPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesystemApi.FilesystemAclIsTrivialPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesystemAclIsTrivialPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

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


## FilesystemChownPost

> FilesystemChownPost(ctx).FilesystemChown0(filesystemChown0).Execute()





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
    filesystemChown0 := *openapiclient.NewFilesystemChown0() // FilesystemChown0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FilesystemApi.FilesystemChownPost(context.Background()).FilesystemChown0(filesystemChown0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesystemApi.FilesystemChownPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesystemChownPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filesystemChown0** | [**FilesystemChown0**](FilesystemChown0.md) |  | 

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


## FilesystemDefaultAclChoicesGet

> FilesystemDefaultAclChoicesGet(ctx).Execute()





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
    resp, r, err := api_client.FilesystemApi.FilesystemDefaultAclChoicesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesystemApi.FilesystemDefaultAclChoicesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFilesystemDefaultAclChoicesGetRequest struct via the builder pattern


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


## FilesystemGetDefaultAclPost

> FilesystemGetDefaultAclPost(ctx).FilesystemGetDefaultAcl(filesystemGetDefaultAcl).Execute()





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
    filesystemGetDefaultAcl := *openapiclient.NewFilesystemGetDefaultAcl() // FilesystemGetDefaultAcl |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FilesystemApi.FilesystemGetDefaultAclPost(context.Background()).FilesystemGetDefaultAcl(filesystemGetDefaultAcl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesystemApi.FilesystemGetDefaultAclPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesystemGetDefaultAclPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filesystemGetDefaultAcl** | [**FilesystemGetDefaultAcl**](FilesystemGetDefaultAcl.md) |  | 

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


## FilesystemGetaclPost

> FilesystemGetaclPost(ctx).FilesystemGetacl(filesystemGetacl).Execute()





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
    filesystemGetacl := *openapiclient.NewFilesystemGetacl() // FilesystemGetacl |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FilesystemApi.FilesystemGetaclPost(context.Background()).FilesystemGetacl(filesystemGetacl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesystemApi.FilesystemGetaclPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesystemGetaclPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filesystemGetacl** | [**FilesystemGetacl**](FilesystemGetacl.md) |  | 

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


## FilesystemListdirPost

> FilesystemListdirPost(ctx).FilesystemListdir(filesystemListdir).Execute()





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
    filesystemListdir := *openapiclient.NewFilesystemListdir() // FilesystemListdir |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FilesystemApi.FilesystemListdirPost(context.Background()).FilesystemListdir(filesystemListdir).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesystemApi.FilesystemListdirPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesystemListdirPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filesystemListdir** | [**FilesystemListdir**](FilesystemListdir.md) |  | 

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


## FilesystemSetaclPost

> FilesystemSetaclPost(ctx).FilesystemSetacl0(filesystemSetacl0).Execute()





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
    filesystemSetacl0 := *openapiclient.NewFilesystemSetacl0() // FilesystemSetacl0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FilesystemApi.FilesystemSetaclPost(context.Background()).FilesystemSetacl0(filesystemSetacl0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesystemApi.FilesystemSetaclPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesystemSetaclPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filesystemSetacl0** | [**FilesystemSetacl0**](FilesystemSetacl0.md) |  | 

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


## FilesystemSetpermPost

> FilesystemSetpermPost(ctx).FilesystemSetperm0(filesystemSetperm0).Execute()





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
    filesystemSetperm0 := *openapiclient.NewFilesystemSetperm0() // FilesystemSetperm0 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FilesystemApi.FilesystemSetpermPost(context.Background()).FilesystemSetperm0(filesystemSetperm0).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesystemApi.FilesystemSetpermPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesystemSetpermPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filesystemSetperm0** | [**FilesystemSetperm0**](FilesystemSetperm0.md) |  | 

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


## FilesystemStatPost

> FilesystemStatPost(ctx).Body(body).Execute()





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
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FilesystemApi.FilesystemStatPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesystemApi.FilesystemStatPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesystemStatPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

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


## FilesystemStatfsPost

> FilesystemStatfsPost(ctx).Body(body).Execute()





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
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FilesystemApi.FilesystemStatfsPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesystemApi.FilesystemStatfsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesystemStatfsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

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

