# \Actors

All URIs are relative to *https://cloud.steampipe.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](Actors.md#Get) | **Get** /actor | Actor information
[**ListActivity**](Actors.md#ListActivity) | **Get** /actor/activity | List actor activity
[**ListConnections**](Actors.md#ListConnections) | **Get** /actor/conn | List actor connections
[**ListWorkspaces**](Actors.md#ListWorkspaces) | **Get** /actor/workspace | List actor workspaces



## Get

> User Get(ctx).Execute()

Actor information



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
    resp, r, err := api_client.Actors.Get(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Actors.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: User
    fmt.Fprintf(os.Stdout, "Response from `Actors.Get`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern
[]
**** | [****](.md) |  | []

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListActivity

> ListAuditLogsResponse ListActivity(ctx).Limit(limit).NextToken(nextToken).Execute()

List actor activity



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
    limit := int32(56) // int32 | Pagination limit (optional) (default to 20)
    nextToken := "nextToken_example" // string | An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Actors.ListActivity(context.Background()).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Actors.ListActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListActivity`: ListAuditLogsResponse
    fmt.Fprintf(os.Stdout, "Response from `Actors.ListActivity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------[CodegenParameter{isFormParam&#x3D;false, isQueryParam&#x3D;true, isPathParam&#x3D;false, isHeaderParam&#x3D;false, isCookieParam&#x3D;false, isBodyParam&#x3D;false, isContainer&#x3D;false, isCollectionFormatMulti&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isExplode&#x3D;true, baseName&#x3D;&#39;limit&#39;, paramName&#x3D;&#39;limit&#39;, dataType&#x3D;&#39;int32&#39;, datatypeWithEnum&#x3D;&#39;null&#39;, dataFormat&#x3D;&#39;null&#39;, collectionFormat&#x3D;&#39;null&#39;, description&#x3D;&#39;Pagination limit&#39;, unescapedDescription&#x3D;&#39;Pagination limit&#39;, baseType&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;20&#39;, enumName&#x3D;&#39;null&#39;, style&#x3D;&#39;form&#39;, deepObject&#x3D;&#39;false&#39;, allowEmptyValue&#x3D;&#39;false&#39;, example&#x3D;&#39;56&#39;, jsonSchema&#x3D;&#39;{
  &quot;name&quot; : &quot;limit&quot;,
  &quot;in&quot; : &quot;query&quot;,
  &quot;description&quot; : &quot;Pagination limit&quot;,
  &quot;required&quot; : false,
  &quot;style&quot; : &quot;form&quot;,
  &quot;explode&quot; : true,
  &quot;schema&quot; : {
    &quot;maximum&quot; : 100,
    &quot;minimum&quot; : 1,
    &quot;type&quot; : &quot;integer&quot;,
    &quot;default&quot; : 20
  }
}&#39;, isString&#x3D;false, isNumeric&#x3D;false, isInteger&#x3D;true, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;true, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isAnyType&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isFile&#x3D;false, isEnum&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, mostInnerItems&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], vendorExtensions&#x3D;{x-go-example&#x3D;int32(56), x-export-param-name&#x3D;Limit}, hasValidation&#x3D;true, maxProperties&#x3D;null, minProperties&#x3D;null, isNullable&#x3D;false, isDeprecated&#x3D;false, required&#x3D;false, maximum&#x3D;&#39;100&#39;, exclusiveMaximum&#x3D;false, minimum&#x3D;&#39;1&#39;, exclusiveMinimum&#x3D;false, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, uniqueItems&#x3D;false, contentType&#x3D;null, multipleOf&#x3D;null, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false, schema&#x3D;CodegenProperty{openApiType&#x3D;&#39;integer&#39;, baseName&#x3D;&#39;limit&#39;, complexType&#x3D;&#39;null&#39;, getter&#x3D;&#39;GetLimit&#39;, setter&#x3D;&#39;setLimit&#39;, description&#x3D;&#39;null&#39;, dataType&#x3D;&#39;int32&#39;, datatypeWithEnum&#x3D;&#39;int32&#39;, dataFormat&#x3D;&#39;null&#39;, name&#x3D;&#39;Limit&#39;, min&#x3D;&#39;null&#39;, max&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;20&#39;, defaultValueWithParam&#x3D;&#39; &#x3D; data.limit;&#39;, baseType&#x3D;&#39;int32&#39;, containerType&#x3D;&#39;null&#39;, title&#x3D;&#39;null&#39;, unescapedDescription&#x3D;&#39;null&#39;, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;maximum&quot; : 100,
  &quot;minimum&quot; : 1,
  &quot;type&quot; : &quot;integer&quot;,
  &quot;default&quot; : 20
}&#39;, minimum&#x3D;&#39;1&#39;, maximum&#x3D;&#39;100&#39;, exclusiveMinimum&#x3D;false, exclusiveMaximum&#x3D;false, required&#x3D;false, deprecated&#x3D;false, hasMoreNonReadOnly&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isContainer&#x3D;false, isString&#x3D;false, isNumeric&#x3D;true, isInteger&#x3D;true, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;true, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isFile&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isEnum&#x3D;false, isReadOnly&#x3D;false, isWriteOnly&#x3D;false, isNullable&#x3D;false, isSelfReference&#x3D;false, isCircularReference&#x3D;false, isDiscriminator&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], mostInnerItems&#x3D;null, vendorExtensions&#x3D;{}, hasValidation&#x3D;true, isInherited&#x3D;false, discriminatorValue&#x3D;&#39;null&#39;, nameInCamelCase&#x3D;&#39;limit&#39;, nameInSnakeCase&#x3D;&#39;LIMIT&#39;, enumName&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, maxProperties&#x3D;null, minProperties&#x3D;null, uniqueItems&#x3D;false, multipleOf&#x3D;null, isXmlAttribute&#x3D;false, xmlPrefix&#x3D;&#39;null&#39;, xmlName&#x3D;&#39;null&#39;, xmlNamespace&#x3D;&#39;null&#39;, isXmlWrapped&#x3D;false, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false}, content&#x3D;null}, CodegenParameter{isFormParam&#x3D;false, isQueryParam&#x3D;true, isPathParam&#x3D;false, isHeaderParam&#x3D;false, isCookieParam&#x3D;false, isBodyParam&#x3D;false, isContainer&#x3D;false, isCollectionFormatMulti&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isExplode&#x3D;true, baseName&#x3D;&#39;next_token&#39;, paramName&#x3D;&#39;nextToken&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;null&#39;, dataFormat&#x3D;&#39;null&#39;, collectionFormat&#x3D;&#39;null&#39;, description&#x3D;&#39;An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request.&#39;, unescapedDescription&#x3D;&#39;An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request.&#39;, baseType&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, enumName&#x3D;&#39;null&#39;, style&#x3D;&#39;form&#39;, deepObject&#x3D;&#39;false&#39;, allowEmptyValue&#x3D;&#39;false&#39;, example&#x3D;&#39;nextToken_example&#39;, jsonSchema&#x3D;&#39;{
  &quot;name&quot; : &quot;next_token&quot;,
  &quot;in&quot; : &quot;query&quot;,
  &quot;description&quot; : &quot;An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request.&quot;,
  &quot;required&quot; : false,
  &quot;style&quot; : &quot;form&quot;,
  &quot;explode&quot; : true,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;string&quot;
  }
}&#39;, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isAnyType&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isFile&#x3D;false, isEnum&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, mostInnerItems&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], vendorExtensions&#x3D;{x-go-example&#x3D;&quot;nextToken_example&quot;, x-export-param-name&#x3D;NextToken}, hasValidation&#x3D;false, maxProperties&#x3D;null, minProperties&#x3D;null, isNullable&#x3D;false, isDeprecated&#x3D;false, required&#x3D;false, maximum&#x3D;&#39;null&#39;, exclusiveMaximum&#x3D;false, minimum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, uniqueItems&#x3D;false, contentType&#x3D;null, multipleOf&#x3D;null, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false, schema&#x3D;CodegenProperty{openApiType&#x3D;&#39;string&#39;, baseName&#x3D;&#39;next_token&#39;, complexType&#x3D;&#39;null&#39;, getter&#x3D;&#39;GetNextToken&#39;, setter&#x3D;&#39;setNextToken&#39;, description&#x3D;&#39;null&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;string&#39;, dataFormat&#x3D;&#39;null&#39;, name&#x3D;&#39;NextToken&#39;, min&#x3D;&#39;null&#39;, max&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, defaultValueWithParam&#x3D;&#39; &#x3D; data.next_token;&#39;, baseType&#x3D;&#39;string&#39;, containerType&#x3D;&#39;null&#39;, title&#x3D;&#39;null&#39;, unescapedDescription&#x3D;&#39;null&#39;, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;type&quot; : &quot;string&quot;
}&#39;, minimum&#x3D;&#39;null&#39;, maximum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, exclusiveMaximum&#x3D;false, required&#x3D;false, deprecated&#x3D;false, hasMoreNonReadOnly&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isContainer&#x3D;false, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isFile&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isEnum&#x3D;false, isReadOnly&#x3D;false, isWriteOnly&#x3D;false, isNullable&#x3D;false, isSelfReference&#x3D;false, isCircularReference&#x3D;false, isDiscriminator&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], mostInnerItems&#x3D;null, vendorExtensions&#x3D;{}, hasValidation&#x3D;false, isInherited&#x3D;false, discriminatorValue&#x3D;&#39;null&#39;, nameInCamelCase&#x3D;&#39;nextToken&#39;, nameInSnakeCase&#x3D;&#39;NEXT_TOKEN&#39;, enumName&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, maxProperties&#x3D;null, minProperties&#x3D;null, uniqueItems&#x3D;false, multipleOf&#x3D;null, isXmlAttribute&#x3D;false, xmlPrefix&#x3D;&#39;null&#39;, xmlName&#x3D;&#39;null&#39;, xmlNamespace&#x3D;&#39;null&#39;, isXmlWrapped&#x3D;false, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false}, content&#x3D;null}]
**** | [****](.md) |  | [CodegenParameter{isFormParam&#x3D;false, isQueryParam&#x3D;true, isPathParam&#x3D;false, isHeaderParam&#x3D;false, isCookieParam&#x3D;false, isBodyParam&#x3D;false, isContainer&#x3D;false, isCollectionFormatMulti&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isExplode&#x3D;true, baseName&#x3D;&#39;limit&#39;, paramName&#x3D;&#39;limit&#39;, dataType&#x3D;&#39;int32&#39;, datatypeWithEnum&#x3D;&#39;null&#39;, dataFormat&#x3D;&#39;null&#39;, collectionFormat&#x3D;&#39;null&#39;, description&#x3D;&#39;Pagination limit&#39;, unescapedDescription&#x3D;&#39;Pagination limit&#39;, baseType&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;20&#39;, enumName&#x3D;&#39;null&#39;, style&#x3D;&#39;form&#39;, deepObject&#x3D;&#39;false&#39;, allowEmptyValue&#x3D;&#39;false&#39;, example&#x3D;&#39;56&#39;, jsonSchema&#x3D;&#39;{
  &quot;name&quot; : &quot;limit&quot;,
  &quot;in&quot; : &quot;query&quot;,
  &quot;description&quot; : &quot;Pagination limit&quot;,
  &quot;required&quot; : false,
  &quot;style&quot; : &quot;form&quot;,
  &quot;explode&quot; : true,
  &quot;schema&quot; : {
    &quot;maximum&quot; : 100,
    &quot;minimum&quot; : 1,
    &quot;type&quot; : &quot;integer&quot;,
    &quot;default&quot; : 20
  }
}&#39;, isString&#x3D;false, isNumeric&#x3D;false, isInteger&#x3D;true, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;true, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isAnyType&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isFile&#x3D;false, isEnum&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, mostInnerItems&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], vendorExtensions&#x3D;{x-go-example&#x3D;int32(56), x-export-param-name&#x3D;Limit}, hasValidation&#x3D;true, maxProperties&#x3D;null, minProperties&#x3D;null, isNullable&#x3D;false, isDeprecated&#x3D;false, required&#x3D;false, maximum&#x3D;&#39;100&#39;, exclusiveMaximum&#x3D;false, minimum&#x3D;&#39;1&#39;, exclusiveMinimum&#x3D;false, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, uniqueItems&#x3D;false, contentType&#x3D;null, multipleOf&#x3D;null, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false, schema&#x3D;CodegenProperty{openApiType&#x3D;&#39;integer&#39;, baseName&#x3D;&#39;limit&#39;, complexType&#x3D;&#39;null&#39;, getter&#x3D;&#39;GetLimit&#39;, setter&#x3D;&#39;setLimit&#39;, description&#x3D;&#39;null&#39;, dataType&#x3D;&#39;int32&#39;, datatypeWithEnum&#x3D;&#39;int32&#39;, dataFormat&#x3D;&#39;null&#39;, name&#x3D;&#39;Limit&#39;, min&#x3D;&#39;null&#39;, max&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;20&#39;, defaultValueWithParam&#x3D;&#39; &#x3D; data.limit;&#39;, baseType&#x3D;&#39;int32&#39;, containerType&#x3D;&#39;null&#39;, title&#x3D;&#39;null&#39;, unescapedDescription&#x3D;&#39;null&#39;, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;maximum&quot; : 100,
  &quot;minimum&quot; : 1,
  &quot;type&quot; : &quot;integer&quot;,
  &quot;default&quot; : 20
}&#39;, minimum&#x3D;&#39;1&#39;, maximum&#x3D;&#39;100&#39;, exclusiveMinimum&#x3D;false, exclusiveMaximum&#x3D;false, required&#x3D;false, deprecated&#x3D;false, hasMoreNonReadOnly&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isContainer&#x3D;false, isString&#x3D;false, isNumeric&#x3D;true, isInteger&#x3D;true, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;true, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isFile&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isEnum&#x3D;false, isReadOnly&#x3D;false, isWriteOnly&#x3D;false, isNullable&#x3D;false, isSelfReference&#x3D;false, isCircularReference&#x3D;false, isDiscriminator&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], mostInnerItems&#x3D;null, vendorExtensions&#x3D;{}, hasValidation&#x3D;true, isInherited&#x3D;false, discriminatorValue&#x3D;&#39;null&#39;, nameInCamelCase&#x3D;&#39;limit&#39;, nameInSnakeCase&#x3D;&#39;LIMIT&#39;, enumName&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, maxProperties&#x3D;null, minProperties&#x3D;null, uniqueItems&#x3D;false, multipleOf&#x3D;null, isXmlAttribute&#x3D;false, xmlPrefix&#x3D;&#39;null&#39;, xmlName&#x3D;&#39;null&#39;, xmlNamespace&#x3D;&#39;null&#39;, isXmlWrapped&#x3D;false, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false}, content&#x3D;null}, CodegenParameter{isFormParam&#x3D;false, isQueryParam&#x3D;true, isPathParam&#x3D;false, isHeaderParam&#x3D;false, isCookieParam&#x3D;false, isBodyParam&#x3D;false, isContainer&#x3D;false, isCollectionFormatMulti&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isExplode&#x3D;true, baseName&#x3D;&#39;next_token&#39;, paramName&#x3D;&#39;nextToken&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;null&#39;, dataFormat&#x3D;&#39;null&#39;, collectionFormat&#x3D;&#39;null&#39;, description&#x3D;&#39;An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request.&#39;, unescapedDescription&#x3D;&#39;An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request.&#39;, baseType&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, enumName&#x3D;&#39;null&#39;, style&#x3D;&#39;form&#39;, deepObject&#x3D;&#39;false&#39;, allowEmptyValue&#x3D;&#39;false&#39;, example&#x3D;&#39;nextToken_example&#39;, jsonSchema&#x3D;&#39;{
  &quot;name&quot; : &quot;next_token&quot;,
  &quot;in&quot; : &quot;query&quot;,
  &quot;description&quot; : &quot;An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request.&quot;,
  &quot;required&quot; : false,
  &quot;style&quot; : &quot;form&quot;,
  &quot;explode&quot; : true,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;string&quot;
  }
}&#39;, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isAnyType&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isFile&#x3D;false, isEnum&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, mostInnerItems&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], vendorExtensions&#x3D;{x-go-example&#x3D;&quot;nextToken_example&quot;, x-export-param-name&#x3D;NextToken}, hasValidation&#x3D;false, maxProperties&#x3D;null, minProperties&#x3D;null, isNullable&#x3D;false, isDeprecated&#x3D;false, required&#x3D;false, maximum&#x3D;&#39;null&#39;, exclusiveMaximum&#x3D;false, minimum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, uniqueItems&#x3D;false, contentType&#x3D;null, multipleOf&#x3D;null, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false, schema&#x3D;CodegenProperty{openApiType&#x3D;&#39;string&#39;, baseName&#x3D;&#39;next_token&#39;, complexType&#x3D;&#39;null&#39;, getter&#x3D;&#39;GetNextToken&#39;, setter&#x3D;&#39;setNextToken&#39;, description&#x3D;&#39;null&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;string&#39;, dataFormat&#x3D;&#39;null&#39;, name&#x3D;&#39;NextToken&#39;, min&#x3D;&#39;null&#39;, max&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, defaultValueWithParam&#x3D;&#39; &#x3D; data.next_token;&#39;, baseType&#x3D;&#39;string&#39;, containerType&#x3D;&#39;null&#39;, title&#x3D;&#39;null&#39;, unescapedDescription&#x3D;&#39;null&#39;, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;type&quot; : &quot;string&quot;
}&#39;, minimum&#x3D;&#39;null&#39;, maximum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, exclusiveMaximum&#x3D;false, required&#x3D;false, deprecated&#x3D;false, hasMoreNonReadOnly&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isContainer&#x3D;false, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isFile&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isEnum&#x3D;false, isReadOnly&#x3D;false, isWriteOnly&#x3D;false, isNullable&#x3D;false, isSelfReference&#x3D;false, isCircularReference&#x3D;false, isDiscriminator&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], mostInnerItems&#x3D;null, vendorExtensions&#x3D;{}, hasValidation&#x3D;false, isInherited&#x3D;false, discriminatorValue&#x3D;&#39;null&#39;, nameInCamelCase&#x3D;&#39;nextToken&#39;, nameInSnakeCase&#x3D;&#39;NEXT_TOKEN&#39;, enumName&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, maxProperties&#x3D;null, minProperties&#x3D;null, uniqueItems&#x3D;false, multipleOf&#x3D;null, isXmlAttribute&#x3D;false, xmlPrefix&#x3D;&#39;null&#39;, xmlName&#x3D;&#39;null&#39;, xmlNamespace&#x3D;&#39;null&#39;, isXmlWrapped&#x3D;false, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false}, content&#x3D;null}]

### Return type

[**ListAuditLogsResponse**](ListAuditLogsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnections

> ListConnectionsResponse ListConnections(ctx).Limit(limit).NextToken(nextToken).Execute()

List actor connections



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
    limit := int32(56) // int32 | Pagination limit (optional) (default to 20)
    nextToken := "nextToken_example" // string | An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Actors.ListConnections(context.Background()).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Actors.ListConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnections`: ListConnectionsResponse
    fmt.Fprintf(os.Stdout, "Response from `Actors.ListConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------[CodegenParameter{isFormParam&#x3D;false, isQueryParam&#x3D;true, isPathParam&#x3D;false, isHeaderParam&#x3D;false, isCookieParam&#x3D;false, isBodyParam&#x3D;false, isContainer&#x3D;false, isCollectionFormatMulti&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isExplode&#x3D;true, baseName&#x3D;&#39;limit&#39;, paramName&#x3D;&#39;limit&#39;, dataType&#x3D;&#39;int32&#39;, datatypeWithEnum&#x3D;&#39;null&#39;, dataFormat&#x3D;&#39;null&#39;, collectionFormat&#x3D;&#39;null&#39;, description&#x3D;&#39;Pagination limit&#39;, unescapedDescription&#x3D;&#39;Pagination limit&#39;, baseType&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;20&#39;, enumName&#x3D;&#39;null&#39;, style&#x3D;&#39;form&#39;, deepObject&#x3D;&#39;false&#39;, allowEmptyValue&#x3D;&#39;false&#39;, example&#x3D;&#39;56&#39;, jsonSchema&#x3D;&#39;{
  &quot;name&quot; : &quot;limit&quot;,
  &quot;in&quot; : &quot;query&quot;,
  &quot;description&quot; : &quot;Pagination limit&quot;,
  &quot;required&quot; : false,
  &quot;style&quot; : &quot;form&quot;,
  &quot;explode&quot; : true,
  &quot;schema&quot; : {
    &quot;maximum&quot; : 100,
    &quot;minimum&quot; : 1,
    &quot;type&quot; : &quot;integer&quot;,
    &quot;default&quot; : 20
  }
}&#39;, isString&#x3D;false, isNumeric&#x3D;false, isInteger&#x3D;true, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;true, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isAnyType&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isFile&#x3D;false, isEnum&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, mostInnerItems&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], vendorExtensions&#x3D;{x-go-example&#x3D;int32(56), x-export-param-name&#x3D;Limit}, hasValidation&#x3D;true, maxProperties&#x3D;null, minProperties&#x3D;null, isNullable&#x3D;false, isDeprecated&#x3D;false, required&#x3D;false, maximum&#x3D;&#39;100&#39;, exclusiveMaximum&#x3D;false, minimum&#x3D;&#39;1&#39;, exclusiveMinimum&#x3D;false, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, uniqueItems&#x3D;false, contentType&#x3D;null, multipleOf&#x3D;null, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false, schema&#x3D;CodegenProperty{openApiType&#x3D;&#39;integer&#39;, baseName&#x3D;&#39;limit&#39;, complexType&#x3D;&#39;null&#39;, getter&#x3D;&#39;GetLimit&#39;, setter&#x3D;&#39;setLimit&#39;, description&#x3D;&#39;null&#39;, dataType&#x3D;&#39;int32&#39;, datatypeWithEnum&#x3D;&#39;int32&#39;, dataFormat&#x3D;&#39;null&#39;, name&#x3D;&#39;Limit&#39;, min&#x3D;&#39;null&#39;, max&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;20&#39;, defaultValueWithParam&#x3D;&#39; &#x3D; data.limit;&#39;, baseType&#x3D;&#39;int32&#39;, containerType&#x3D;&#39;null&#39;, title&#x3D;&#39;null&#39;, unescapedDescription&#x3D;&#39;null&#39;, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;maximum&quot; : 100,
  &quot;minimum&quot; : 1,
  &quot;type&quot; : &quot;integer&quot;,
  &quot;default&quot; : 20
}&#39;, minimum&#x3D;&#39;1&#39;, maximum&#x3D;&#39;100&#39;, exclusiveMinimum&#x3D;false, exclusiveMaximum&#x3D;false, required&#x3D;false, deprecated&#x3D;false, hasMoreNonReadOnly&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isContainer&#x3D;false, isString&#x3D;false, isNumeric&#x3D;true, isInteger&#x3D;true, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;true, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isFile&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isEnum&#x3D;false, isReadOnly&#x3D;false, isWriteOnly&#x3D;false, isNullable&#x3D;false, isSelfReference&#x3D;false, isCircularReference&#x3D;false, isDiscriminator&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], mostInnerItems&#x3D;null, vendorExtensions&#x3D;{}, hasValidation&#x3D;true, isInherited&#x3D;false, discriminatorValue&#x3D;&#39;null&#39;, nameInCamelCase&#x3D;&#39;limit&#39;, nameInSnakeCase&#x3D;&#39;LIMIT&#39;, enumName&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, maxProperties&#x3D;null, minProperties&#x3D;null, uniqueItems&#x3D;false, multipleOf&#x3D;null, isXmlAttribute&#x3D;false, xmlPrefix&#x3D;&#39;null&#39;, xmlName&#x3D;&#39;null&#39;, xmlNamespace&#x3D;&#39;null&#39;, isXmlWrapped&#x3D;false, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false}, content&#x3D;null}, CodegenParameter{isFormParam&#x3D;false, isQueryParam&#x3D;true, isPathParam&#x3D;false, isHeaderParam&#x3D;false, isCookieParam&#x3D;false, isBodyParam&#x3D;false, isContainer&#x3D;false, isCollectionFormatMulti&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isExplode&#x3D;true, baseName&#x3D;&#39;next_token&#39;, paramName&#x3D;&#39;nextToken&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;null&#39;, dataFormat&#x3D;&#39;null&#39;, collectionFormat&#x3D;&#39;null&#39;, description&#x3D;&#39;An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request.&#39;, unescapedDescription&#x3D;&#39;An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request.&#39;, baseType&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, enumName&#x3D;&#39;null&#39;, style&#x3D;&#39;form&#39;, deepObject&#x3D;&#39;false&#39;, allowEmptyValue&#x3D;&#39;false&#39;, example&#x3D;&#39;nextToken_example&#39;, jsonSchema&#x3D;&#39;{
  &quot;name&quot; : &quot;next_token&quot;,
  &quot;in&quot; : &quot;query&quot;,
  &quot;description&quot; : &quot;An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request.&quot;,
  &quot;required&quot; : false,
  &quot;style&quot; : &quot;form&quot;,
  &quot;explode&quot; : true,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;string&quot;
  }
}&#39;, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isAnyType&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isFile&#x3D;false, isEnum&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, mostInnerItems&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], vendorExtensions&#x3D;{x-go-example&#x3D;&quot;nextToken_example&quot;, x-export-param-name&#x3D;NextToken}, hasValidation&#x3D;false, maxProperties&#x3D;null, minProperties&#x3D;null, isNullable&#x3D;false, isDeprecated&#x3D;false, required&#x3D;false, maximum&#x3D;&#39;null&#39;, exclusiveMaximum&#x3D;false, minimum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, uniqueItems&#x3D;false, contentType&#x3D;null, multipleOf&#x3D;null, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false, schema&#x3D;CodegenProperty{openApiType&#x3D;&#39;string&#39;, baseName&#x3D;&#39;next_token&#39;, complexType&#x3D;&#39;null&#39;, getter&#x3D;&#39;GetNextToken&#39;, setter&#x3D;&#39;setNextToken&#39;, description&#x3D;&#39;null&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;string&#39;, dataFormat&#x3D;&#39;null&#39;, name&#x3D;&#39;NextToken&#39;, min&#x3D;&#39;null&#39;, max&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, defaultValueWithParam&#x3D;&#39; &#x3D; data.next_token;&#39;, baseType&#x3D;&#39;string&#39;, containerType&#x3D;&#39;null&#39;, title&#x3D;&#39;null&#39;, unescapedDescription&#x3D;&#39;null&#39;, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;type&quot; : &quot;string&quot;
}&#39;, minimum&#x3D;&#39;null&#39;, maximum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, exclusiveMaximum&#x3D;false, required&#x3D;false, deprecated&#x3D;false, hasMoreNonReadOnly&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isContainer&#x3D;false, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isFile&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isEnum&#x3D;false, isReadOnly&#x3D;false, isWriteOnly&#x3D;false, isNullable&#x3D;false, isSelfReference&#x3D;false, isCircularReference&#x3D;false, isDiscriminator&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], mostInnerItems&#x3D;null, vendorExtensions&#x3D;{}, hasValidation&#x3D;false, isInherited&#x3D;false, discriminatorValue&#x3D;&#39;null&#39;, nameInCamelCase&#x3D;&#39;nextToken&#39;, nameInSnakeCase&#x3D;&#39;NEXT_TOKEN&#39;, enumName&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, maxProperties&#x3D;null, minProperties&#x3D;null, uniqueItems&#x3D;false, multipleOf&#x3D;null, isXmlAttribute&#x3D;false, xmlPrefix&#x3D;&#39;null&#39;, xmlName&#x3D;&#39;null&#39;, xmlNamespace&#x3D;&#39;null&#39;, isXmlWrapped&#x3D;false, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false}, content&#x3D;null}]
**** | [****](.md) |  | [CodegenParameter{isFormParam&#x3D;false, isQueryParam&#x3D;true, isPathParam&#x3D;false, isHeaderParam&#x3D;false, isCookieParam&#x3D;false, isBodyParam&#x3D;false, isContainer&#x3D;false, isCollectionFormatMulti&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isExplode&#x3D;true, baseName&#x3D;&#39;limit&#39;, paramName&#x3D;&#39;limit&#39;, dataType&#x3D;&#39;int32&#39;, datatypeWithEnum&#x3D;&#39;null&#39;, dataFormat&#x3D;&#39;null&#39;, collectionFormat&#x3D;&#39;null&#39;, description&#x3D;&#39;Pagination limit&#39;, unescapedDescription&#x3D;&#39;Pagination limit&#39;, baseType&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;20&#39;, enumName&#x3D;&#39;null&#39;, style&#x3D;&#39;form&#39;, deepObject&#x3D;&#39;false&#39;, allowEmptyValue&#x3D;&#39;false&#39;, example&#x3D;&#39;56&#39;, jsonSchema&#x3D;&#39;{
  &quot;name&quot; : &quot;limit&quot;,
  &quot;in&quot; : &quot;query&quot;,
  &quot;description&quot; : &quot;Pagination limit&quot;,
  &quot;required&quot; : false,
  &quot;style&quot; : &quot;form&quot;,
  &quot;explode&quot; : true,
  &quot;schema&quot; : {
    &quot;maximum&quot; : 100,
    &quot;minimum&quot; : 1,
    &quot;type&quot; : &quot;integer&quot;,
    &quot;default&quot; : 20
  }
}&#39;, isString&#x3D;false, isNumeric&#x3D;false, isInteger&#x3D;true, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;true, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isAnyType&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isFile&#x3D;false, isEnum&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, mostInnerItems&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], vendorExtensions&#x3D;{x-go-example&#x3D;int32(56), x-export-param-name&#x3D;Limit}, hasValidation&#x3D;true, maxProperties&#x3D;null, minProperties&#x3D;null, isNullable&#x3D;false, isDeprecated&#x3D;false, required&#x3D;false, maximum&#x3D;&#39;100&#39;, exclusiveMaximum&#x3D;false, minimum&#x3D;&#39;1&#39;, exclusiveMinimum&#x3D;false, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, uniqueItems&#x3D;false, contentType&#x3D;null, multipleOf&#x3D;null, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false, schema&#x3D;CodegenProperty{openApiType&#x3D;&#39;integer&#39;, baseName&#x3D;&#39;limit&#39;, complexType&#x3D;&#39;null&#39;, getter&#x3D;&#39;GetLimit&#39;, setter&#x3D;&#39;setLimit&#39;, description&#x3D;&#39;null&#39;, dataType&#x3D;&#39;int32&#39;, datatypeWithEnum&#x3D;&#39;int32&#39;, dataFormat&#x3D;&#39;null&#39;, name&#x3D;&#39;Limit&#39;, min&#x3D;&#39;null&#39;, max&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;20&#39;, defaultValueWithParam&#x3D;&#39; &#x3D; data.limit;&#39;, baseType&#x3D;&#39;int32&#39;, containerType&#x3D;&#39;null&#39;, title&#x3D;&#39;null&#39;, unescapedDescription&#x3D;&#39;null&#39;, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;maximum&quot; : 100,
  &quot;minimum&quot; : 1,
  &quot;type&quot; : &quot;integer&quot;,
  &quot;default&quot; : 20
}&#39;, minimum&#x3D;&#39;1&#39;, maximum&#x3D;&#39;100&#39;, exclusiveMinimum&#x3D;false, exclusiveMaximum&#x3D;false, required&#x3D;false, deprecated&#x3D;false, hasMoreNonReadOnly&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isContainer&#x3D;false, isString&#x3D;false, isNumeric&#x3D;true, isInteger&#x3D;true, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;true, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isFile&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isEnum&#x3D;false, isReadOnly&#x3D;false, isWriteOnly&#x3D;false, isNullable&#x3D;false, isSelfReference&#x3D;false, isCircularReference&#x3D;false, isDiscriminator&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], mostInnerItems&#x3D;null, vendorExtensions&#x3D;{}, hasValidation&#x3D;true, isInherited&#x3D;false, discriminatorValue&#x3D;&#39;null&#39;, nameInCamelCase&#x3D;&#39;limit&#39;, nameInSnakeCase&#x3D;&#39;LIMIT&#39;, enumName&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, maxProperties&#x3D;null, minProperties&#x3D;null, uniqueItems&#x3D;false, multipleOf&#x3D;null, isXmlAttribute&#x3D;false, xmlPrefix&#x3D;&#39;null&#39;, xmlName&#x3D;&#39;null&#39;, xmlNamespace&#x3D;&#39;null&#39;, isXmlWrapped&#x3D;false, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false}, content&#x3D;null}, CodegenParameter{isFormParam&#x3D;false, isQueryParam&#x3D;true, isPathParam&#x3D;false, isHeaderParam&#x3D;false, isCookieParam&#x3D;false, isBodyParam&#x3D;false, isContainer&#x3D;false, isCollectionFormatMulti&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isExplode&#x3D;true, baseName&#x3D;&#39;next_token&#39;, paramName&#x3D;&#39;nextToken&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;null&#39;, dataFormat&#x3D;&#39;null&#39;, collectionFormat&#x3D;&#39;null&#39;, description&#x3D;&#39;An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request.&#39;, unescapedDescription&#x3D;&#39;An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request.&#39;, baseType&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, enumName&#x3D;&#39;null&#39;, style&#x3D;&#39;form&#39;, deepObject&#x3D;&#39;false&#39;, allowEmptyValue&#x3D;&#39;false&#39;, example&#x3D;&#39;nextToken_example&#39;, jsonSchema&#x3D;&#39;{
  &quot;name&quot; : &quot;next_token&quot;,
  &quot;in&quot; : &quot;query&quot;,
  &quot;description&quot; : &quot;An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request.&quot;,
  &quot;required&quot; : false,
  &quot;style&quot; : &quot;form&quot;,
  &quot;explode&quot; : true,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;string&quot;
  }
}&#39;, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isAnyType&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isFile&#x3D;false, isEnum&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, mostInnerItems&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], vendorExtensions&#x3D;{x-go-example&#x3D;&quot;nextToken_example&quot;, x-export-param-name&#x3D;NextToken}, hasValidation&#x3D;false, maxProperties&#x3D;null, minProperties&#x3D;null, isNullable&#x3D;false, isDeprecated&#x3D;false, required&#x3D;false, maximum&#x3D;&#39;null&#39;, exclusiveMaximum&#x3D;false, minimum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, uniqueItems&#x3D;false, contentType&#x3D;null, multipleOf&#x3D;null, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false, schema&#x3D;CodegenProperty{openApiType&#x3D;&#39;string&#39;, baseName&#x3D;&#39;next_token&#39;, complexType&#x3D;&#39;null&#39;, getter&#x3D;&#39;GetNextToken&#39;, setter&#x3D;&#39;setNextToken&#39;, description&#x3D;&#39;null&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;string&#39;, dataFormat&#x3D;&#39;null&#39;, name&#x3D;&#39;NextToken&#39;, min&#x3D;&#39;null&#39;, max&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, defaultValueWithParam&#x3D;&#39; &#x3D; data.next_token;&#39;, baseType&#x3D;&#39;string&#39;, containerType&#x3D;&#39;null&#39;, title&#x3D;&#39;null&#39;, unescapedDescription&#x3D;&#39;null&#39;, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;type&quot; : &quot;string&quot;
}&#39;, minimum&#x3D;&#39;null&#39;, maximum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, exclusiveMaximum&#x3D;false, required&#x3D;false, deprecated&#x3D;false, hasMoreNonReadOnly&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isContainer&#x3D;false, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isFile&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isEnum&#x3D;false, isReadOnly&#x3D;false, isWriteOnly&#x3D;false, isNullable&#x3D;false, isSelfReference&#x3D;false, isCircularReference&#x3D;false, isDiscriminator&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], mostInnerItems&#x3D;null, vendorExtensions&#x3D;{}, hasValidation&#x3D;false, isInherited&#x3D;false, discriminatorValue&#x3D;&#39;null&#39;, nameInCamelCase&#x3D;&#39;nextToken&#39;, nameInSnakeCase&#x3D;&#39;NEXT_TOKEN&#39;, enumName&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, maxProperties&#x3D;null, minProperties&#x3D;null, uniqueItems&#x3D;false, multipleOf&#x3D;null, isXmlAttribute&#x3D;false, xmlPrefix&#x3D;&#39;null&#39;, xmlName&#x3D;&#39;null&#39;, xmlNamespace&#x3D;&#39;null&#39;, isXmlWrapped&#x3D;false, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false}, content&#x3D;null}]

### Return type

[**ListConnectionsResponse**](ListConnectionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkspaces

> ListWorkspacesResponse ListWorkspaces(ctx).Limit(limit).NextToken(nextToken).Execute()

List actor workspaces



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
    limit := int32(56) // int32 | Pagination limit (optional) (default to 20)
    nextToken := "nextToken_example" // string | An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Actors.ListWorkspaces(context.Background()).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Actors.ListWorkspaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkspaces`: ListWorkspacesResponse
    fmt.Fprintf(os.Stdout, "Response from `Actors.ListWorkspaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWorkspacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------[CodegenParameter{isFormParam&#x3D;false, isQueryParam&#x3D;true, isPathParam&#x3D;false, isHeaderParam&#x3D;false, isCookieParam&#x3D;false, isBodyParam&#x3D;false, isContainer&#x3D;false, isCollectionFormatMulti&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isExplode&#x3D;true, baseName&#x3D;&#39;limit&#39;, paramName&#x3D;&#39;limit&#39;, dataType&#x3D;&#39;int32&#39;, datatypeWithEnum&#x3D;&#39;null&#39;, dataFormat&#x3D;&#39;null&#39;, collectionFormat&#x3D;&#39;null&#39;, description&#x3D;&#39;Pagination limit&#39;, unescapedDescription&#x3D;&#39;Pagination limit&#39;, baseType&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;20&#39;, enumName&#x3D;&#39;null&#39;, style&#x3D;&#39;form&#39;, deepObject&#x3D;&#39;false&#39;, allowEmptyValue&#x3D;&#39;false&#39;, example&#x3D;&#39;56&#39;, jsonSchema&#x3D;&#39;{
  &quot;name&quot; : &quot;limit&quot;,
  &quot;in&quot; : &quot;query&quot;,
  &quot;description&quot; : &quot;Pagination limit&quot;,
  &quot;required&quot; : false,
  &quot;style&quot; : &quot;form&quot;,
  &quot;explode&quot; : true,
  &quot;schema&quot; : {
    &quot;maximum&quot; : 100,
    &quot;minimum&quot; : 1,
    &quot;type&quot; : &quot;integer&quot;,
    &quot;default&quot; : 20
  }
}&#39;, isString&#x3D;false, isNumeric&#x3D;false, isInteger&#x3D;true, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;true, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isAnyType&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isFile&#x3D;false, isEnum&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, mostInnerItems&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], vendorExtensions&#x3D;{x-go-example&#x3D;int32(56), x-export-param-name&#x3D;Limit}, hasValidation&#x3D;true, maxProperties&#x3D;null, minProperties&#x3D;null, isNullable&#x3D;false, isDeprecated&#x3D;false, required&#x3D;false, maximum&#x3D;&#39;100&#39;, exclusiveMaximum&#x3D;false, minimum&#x3D;&#39;1&#39;, exclusiveMinimum&#x3D;false, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, uniqueItems&#x3D;false, contentType&#x3D;null, multipleOf&#x3D;null, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false, schema&#x3D;CodegenProperty{openApiType&#x3D;&#39;integer&#39;, baseName&#x3D;&#39;limit&#39;, complexType&#x3D;&#39;null&#39;, getter&#x3D;&#39;GetLimit&#39;, setter&#x3D;&#39;setLimit&#39;, description&#x3D;&#39;null&#39;, dataType&#x3D;&#39;int32&#39;, datatypeWithEnum&#x3D;&#39;int32&#39;, dataFormat&#x3D;&#39;null&#39;, name&#x3D;&#39;Limit&#39;, min&#x3D;&#39;null&#39;, max&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;20&#39;, defaultValueWithParam&#x3D;&#39; &#x3D; data.limit;&#39;, baseType&#x3D;&#39;int32&#39;, containerType&#x3D;&#39;null&#39;, title&#x3D;&#39;null&#39;, unescapedDescription&#x3D;&#39;null&#39;, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;maximum&quot; : 100,
  &quot;minimum&quot; : 1,
  &quot;type&quot; : &quot;integer&quot;,
  &quot;default&quot; : 20
}&#39;, minimum&#x3D;&#39;1&#39;, maximum&#x3D;&#39;100&#39;, exclusiveMinimum&#x3D;false, exclusiveMaximum&#x3D;false, required&#x3D;false, deprecated&#x3D;false, hasMoreNonReadOnly&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isContainer&#x3D;false, isString&#x3D;false, isNumeric&#x3D;true, isInteger&#x3D;true, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;true, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isFile&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isEnum&#x3D;false, isReadOnly&#x3D;false, isWriteOnly&#x3D;false, isNullable&#x3D;false, isSelfReference&#x3D;false, isCircularReference&#x3D;false, isDiscriminator&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], mostInnerItems&#x3D;null, vendorExtensions&#x3D;{}, hasValidation&#x3D;true, isInherited&#x3D;false, discriminatorValue&#x3D;&#39;null&#39;, nameInCamelCase&#x3D;&#39;limit&#39;, nameInSnakeCase&#x3D;&#39;LIMIT&#39;, enumName&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, maxProperties&#x3D;null, minProperties&#x3D;null, uniqueItems&#x3D;false, multipleOf&#x3D;null, isXmlAttribute&#x3D;false, xmlPrefix&#x3D;&#39;null&#39;, xmlName&#x3D;&#39;null&#39;, xmlNamespace&#x3D;&#39;null&#39;, isXmlWrapped&#x3D;false, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false}, content&#x3D;null}, CodegenParameter{isFormParam&#x3D;false, isQueryParam&#x3D;true, isPathParam&#x3D;false, isHeaderParam&#x3D;false, isCookieParam&#x3D;false, isBodyParam&#x3D;false, isContainer&#x3D;false, isCollectionFormatMulti&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isExplode&#x3D;true, baseName&#x3D;&#39;next_token&#39;, paramName&#x3D;&#39;nextToken&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;null&#39;, dataFormat&#x3D;&#39;null&#39;, collectionFormat&#x3D;&#39;null&#39;, description&#x3D;&#39;An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request.&#39;, unescapedDescription&#x3D;&#39;An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request.&#39;, baseType&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, enumName&#x3D;&#39;null&#39;, style&#x3D;&#39;form&#39;, deepObject&#x3D;&#39;false&#39;, allowEmptyValue&#x3D;&#39;false&#39;, example&#x3D;&#39;nextToken_example&#39;, jsonSchema&#x3D;&#39;{
  &quot;name&quot; : &quot;next_token&quot;,
  &quot;in&quot; : &quot;query&quot;,
  &quot;description&quot; : &quot;An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request.&quot;,
  &quot;required&quot; : false,
  &quot;style&quot; : &quot;form&quot;,
  &quot;explode&quot; : true,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;string&quot;
  }
}&#39;, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isAnyType&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isFile&#x3D;false, isEnum&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, mostInnerItems&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], vendorExtensions&#x3D;{x-go-example&#x3D;&quot;nextToken_example&quot;, x-export-param-name&#x3D;NextToken}, hasValidation&#x3D;false, maxProperties&#x3D;null, minProperties&#x3D;null, isNullable&#x3D;false, isDeprecated&#x3D;false, required&#x3D;false, maximum&#x3D;&#39;null&#39;, exclusiveMaximum&#x3D;false, minimum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, uniqueItems&#x3D;false, contentType&#x3D;null, multipleOf&#x3D;null, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false, schema&#x3D;CodegenProperty{openApiType&#x3D;&#39;string&#39;, baseName&#x3D;&#39;next_token&#39;, complexType&#x3D;&#39;null&#39;, getter&#x3D;&#39;GetNextToken&#39;, setter&#x3D;&#39;setNextToken&#39;, description&#x3D;&#39;null&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;string&#39;, dataFormat&#x3D;&#39;null&#39;, name&#x3D;&#39;NextToken&#39;, min&#x3D;&#39;null&#39;, max&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, defaultValueWithParam&#x3D;&#39; &#x3D; data.next_token;&#39;, baseType&#x3D;&#39;string&#39;, containerType&#x3D;&#39;null&#39;, title&#x3D;&#39;null&#39;, unescapedDescription&#x3D;&#39;null&#39;, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;type&quot; : &quot;string&quot;
}&#39;, minimum&#x3D;&#39;null&#39;, maximum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, exclusiveMaximum&#x3D;false, required&#x3D;false, deprecated&#x3D;false, hasMoreNonReadOnly&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isContainer&#x3D;false, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isFile&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isEnum&#x3D;false, isReadOnly&#x3D;false, isWriteOnly&#x3D;false, isNullable&#x3D;false, isSelfReference&#x3D;false, isCircularReference&#x3D;false, isDiscriminator&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], mostInnerItems&#x3D;null, vendorExtensions&#x3D;{}, hasValidation&#x3D;false, isInherited&#x3D;false, discriminatorValue&#x3D;&#39;null&#39;, nameInCamelCase&#x3D;&#39;nextToken&#39;, nameInSnakeCase&#x3D;&#39;NEXT_TOKEN&#39;, enumName&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, maxProperties&#x3D;null, minProperties&#x3D;null, uniqueItems&#x3D;false, multipleOf&#x3D;null, isXmlAttribute&#x3D;false, xmlPrefix&#x3D;&#39;null&#39;, xmlName&#x3D;&#39;null&#39;, xmlNamespace&#x3D;&#39;null&#39;, isXmlWrapped&#x3D;false, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false}, content&#x3D;null}]
**** | [****](.md) |  | [CodegenParameter{isFormParam&#x3D;false, isQueryParam&#x3D;true, isPathParam&#x3D;false, isHeaderParam&#x3D;false, isCookieParam&#x3D;false, isBodyParam&#x3D;false, isContainer&#x3D;false, isCollectionFormatMulti&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isExplode&#x3D;true, baseName&#x3D;&#39;limit&#39;, paramName&#x3D;&#39;limit&#39;, dataType&#x3D;&#39;int32&#39;, datatypeWithEnum&#x3D;&#39;null&#39;, dataFormat&#x3D;&#39;null&#39;, collectionFormat&#x3D;&#39;null&#39;, description&#x3D;&#39;Pagination limit&#39;, unescapedDescription&#x3D;&#39;Pagination limit&#39;, baseType&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;20&#39;, enumName&#x3D;&#39;null&#39;, style&#x3D;&#39;form&#39;, deepObject&#x3D;&#39;false&#39;, allowEmptyValue&#x3D;&#39;false&#39;, example&#x3D;&#39;56&#39;, jsonSchema&#x3D;&#39;{
  &quot;name&quot; : &quot;limit&quot;,
  &quot;in&quot; : &quot;query&quot;,
  &quot;description&quot; : &quot;Pagination limit&quot;,
  &quot;required&quot; : false,
  &quot;style&quot; : &quot;form&quot;,
  &quot;explode&quot; : true,
  &quot;schema&quot; : {
    &quot;maximum&quot; : 100,
    &quot;minimum&quot; : 1,
    &quot;type&quot; : &quot;integer&quot;,
    &quot;default&quot; : 20
  }
}&#39;, isString&#x3D;false, isNumeric&#x3D;false, isInteger&#x3D;true, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;true, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isAnyType&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isFile&#x3D;false, isEnum&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, mostInnerItems&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], vendorExtensions&#x3D;{x-go-example&#x3D;int32(56), x-export-param-name&#x3D;Limit}, hasValidation&#x3D;true, maxProperties&#x3D;null, minProperties&#x3D;null, isNullable&#x3D;false, isDeprecated&#x3D;false, required&#x3D;false, maximum&#x3D;&#39;100&#39;, exclusiveMaximum&#x3D;false, minimum&#x3D;&#39;1&#39;, exclusiveMinimum&#x3D;false, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, uniqueItems&#x3D;false, contentType&#x3D;null, multipleOf&#x3D;null, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false, schema&#x3D;CodegenProperty{openApiType&#x3D;&#39;integer&#39;, baseName&#x3D;&#39;limit&#39;, complexType&#x3D;&#39;null&#39;, getter&#x3D;&#39;GetLimit&#39;, setter&#x3D;&#39;setLimit&#39;, description&#x3D;&#39;null&#39;, dataType&#x3D;&#39;int32&#39;, datatypeWithEnum&#x3D;&#39;int32&#39;, dataFormat&#x3D;&#39;null&#39;, name&#x3D;&#39;Limit&#39;, min&#x3D;&#39;null&#39;, max&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;20&#39;, defaultValueWithParam&#x3D;&#39; &#x3D; data.limit;&#39;, baseType&#x3D;&#39;int32&#39;, containerType&#x3D;&#39;null&#39;, title&#x3D;&#39;null&#39;, unescapedDescription&#x3D;&#39;null&#39;, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;maximum&quot; : 100,
  &quot;minimum&quot; : 1,
  &quot;type&quot; : &quot;integer&quot;,
  &quot;default&quot; : 20
}&#39;, minimum&#x3D;&#39;1&#39;, maximum&#x3D;&#39;100&#39;, exclusiveMinimum&#x3D;false, exclusiveMaximum&#x3D;false, required&#x3D;false, deprecated&#x3D;false, hasMoreNonReadOnly&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isContainer&#x3D;false, isString&#x3D;false, isNumeric&#x3D;true, isInteger&#x3D;true, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;true, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isFile&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isEnum&#x3D;false, isReadOnly&#x3D;false, isWriteOnly&#x3D;false, isNullable&#x3D;false, isSelfReference&#x3D;false, isCircularReference&#x3D;false, isDiscriminator&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], mostInnerItems&#x3D;null, vendorExtensions&#x3D;{}, hasValidation&#x3D;true, isInherited&#x3D;false, discriminatorValue&#x3D;&#39;null&#39;, nameInCamelCase&#x3D;&#39;limit&#39;, nameInSnakeCase&#x3D;&#39;LIMIT&#39;, enumName&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, maxProperties&#x3D;null, minProperties&#x3D;null, uniqueItems&#x3D;false, multipleOf&#x3D;null, isXmlAttribute&#x3D;false, xmlPrefix&#x3D;&#39;null&#39;, xmlName&#x3D;&#39;null&#39;, xmlNamespace&#x3D;&#39;null&#39;, isXmlWrapped&#x3D;false, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false}, content&#x3D;null}, CodegenParameter{isFormParam&#x3D;false, isQueryParam&#x3D;true, isPathParam&#x3D;false, isHeaderParam&#x3D;false, isCookieParam&#x3D;false, isBodyParam&#x3D;false, isContainer&#x3D;false, isCollectionFormatMulti&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isExplode&#x3D;true, baseName&#x3D;&#39;next_token&#39;, paramName&#x3D;&#39;nextToken&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;null&#39;, dataFormat&#x3D;&#39;null&#39;, collectionFormat&#x3D;&#39;null&#39;, description&#x3D;&#39;An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request.&#39;, unescapedDescription&#x3D;&#39;An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request.&#39;, baseType&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, enumName&#x3D;&#39;null&#39;, style&#x3D;&#39;form&#39;, deepObject&#x3D;&#39;false&#39;, allowEmptyValue&#x3D;&#39;false&#39;, example&#x3D;&#39;nextToken_example&#39;, jsonSchema&#x3D;&#39;{
  &quot;name&quot; : &quot;next_token&quot;,
  &quot;in&quot; : &quot;query&quot;,
  &quot;description&quot; : &quot;An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request.&quot;,
  &quot;required&quot; : false,
  &quot;style&quot; : &quot;form&quot;,
  &quot;explode&quot; : true,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;string&quot;
  }
}&#39;, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isAnyType&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isFile&#x3D;false, isEnum&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, mostInnerItems&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], vendorExtensions&#x3D;{x-go-example&#x3D;&quot;nextToken_example&quot;, x-export-param-name&#x3D;NextToken}, hasValidation&#x3D;false, maxProperties&#x3D;null, minProperties&#x3D;null, isNullable&#x3D;false, isDeprecated&#x3D;false, required&#x3D;false, maximum&#x3D;&#39;null&#39;, exclusiveMaximum&#x3D;false, minimum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, uniqueItems&#x3D;false, contentType&#x3D;null, multipleOf&#x3D;null, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false, schema&#x3D;CodegenProperty{openApiType&#x3D;&#39;string&#39;, baseName&#x3D;&#39;next_token&#39;, complexType&#x3D;&#39;null&#39;, getter&#x3D;&#39;GetNextToken&#39;, setter&#x3D;&#39;setNextToken&#39;, description&#x3D;&#39;null&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;string&#39;, dataFormat&#x3D;&#39;null&#39;, name&#x3D;&#39;NextToken&#39;, min&#x3D;&#39;null&#39;, max&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, defaultValueWithParam&#x3D;&#39; &#x3D; data.next_token;&#39;, baseType&#x3D;&#39;string&#39;, containerType&#x3D;&#39;null&#39;, title&#x3D;&#39;null&#39;, unescapedDescription&#x3D;&#39;null&#39;, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;type&quot; : &quot;string&quot;
}&#39;, minimum&#x3D;&#39;null&#39;, maximum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, exclusiveMaximum&#x3D;false, required&#x3D;false, deprecated&#x3D;false, hasMoreNonReadOnly&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isContainer&#x3D;false, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isFile&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isEnum&#x3D;false, isReadOnly&#x3D;false, isWriteOnly&#x3D;false, isNullable&#x3D;false, isSelfReference&#x3D;false, isCircularReference&#x3D;false, isDiscriminator&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], mostInnerItems&#x3D;null, vendorExtensions&#x3D;{}, hasValidation&#x3D;false, isInherited&#x3D;false, discriminatorValue&#x3D;&#39;null&#39;, nameInCamelCase&#x3D;&#39;nextToken&#39;, nameInSnakeCase&#x3D;&#39;NEXT_TOKEN&#39;, enumName&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, maxProperties&#x3D;null, minProperties&#x3D;null, uniqueItems&#x3D;false, multipleOf&#x3D;null, isXmlAttribute&#x3D;false, xmlPrefix&#x3D;&#39;null&#39;, xmlName&#x3D;&#39;null&#39;, xmlNamespace&#x3D;&#39;null&#39;, isXmlWrapped&#x3D;false, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false}, content&#x3D;null}]

### Return type

[**ListWorkspacesResponse**](ListWorkspacesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

