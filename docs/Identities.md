# \Identities

All URIs are relative to *https://cloud.steampipe.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Search**](Identities.md#Search) | **Get** /identity/search | Search identity



## Search

> SearchIdentitiesResponse Search(ctx).Q(q).Execute()

Search identity



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
    q := "q_example" // string | Specify the search string.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Identities.Search(context.Background()).Q(q).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identities.Search``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Search`: SearchIdentitiesResponse
    fmt.Fprintf(os.Stdout, "Response from `Identities.Search`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------[CodegenParameter{isFormParam&#x3D;false, isQueryParam&#x3D;true, isPathParam&#x3D;false, isHeaderParam&#x3D;false, isCookieParam&#x3D;false, isBodyParam&#x3D;false, isContainer&#x3D;false, isCollectionFormatMulti&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isExplode&#x3D;true, baseName&#x3D;&#39;q&#39;, paramName&#x3D;&#39;q&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;null&#39;, dataFormat&#x3D;&#39;null&#39;, collectionFormat&#x3D;&#39;null&#39;, description&#x3D;&#39;Specify the search string.&#39;, unescapedDescription&#x3D;&#39;Specify the search string.&#39;, baseType&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, enumName&#x3D;&#39;null&#39;, style&#x3D;&#39;form&#39;, deepObject&#x3D;&#39;false&#39;, allowEmptyValue&#x3D;&#39;false&#39;, example&#x3D;&#39;q_example&#39;, jsonSchema&#x3D;&#39;{
  &quot;name&quot; : &quot;q&quot;,
  &quot;in&quot; : &quot;query&quot;,
  &quot;description&quot; : &quot;Specify the search string.&quot;,
  &quot;required&quot; : true,
  &quot;style&quot; : &quot;form&quot;,
  &quot;explode&quot; : true,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;string&quot;
  }
}&#39;, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isAnyType&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isFile&#x3D;false, isEnum&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, mostInnerItems&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], vendorExtensions&#x3D;{x-go-example&#x3D;&quot;q_example&quot;, x-export-param-name&#x3D;Q}, hasValidation&#x3D;false, maxProperties&#x3D;null, minProperties&#x3D;null, isNullable&#x3D;false, isDeprecated&#x3D;false, required&#x3D;true, maximum&#x3D;&#39;null&#39;, exclusiveMaximum&#x3D;false, minimum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, uniqueItems&#x3D;false, contentType&#x3D;null, multipleOf&#x3D;null, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false, schema&#x3D;CodegenProperty{openApiType&#x3D;&#39;string&#39;, baseName&#x3D;&#39;q&#39;, complexType&#x3D;&#39;null&#39;, getter&#x3D;&#39;GetQ&#39;, setter&#x3D;&#39;setQ&#39;, description&#x3D;&#39;null&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;string&#39;, dataFormat&#x3D;&#39;null&#39;, name&#x3D;&#39;Q&#39;, min&#x3D;&#39;null&#39;, max&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, defaultValueWithParam&#x3D;&#39; &#x3D; data.q;&#39;, baseType&#x3D;&#39;string&#39;, containerType&#x3D;&#39;null&#39;, title&#x3D;&#39;null&#39;, unescapedDescription&#x3D;&#39;null&#39;, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;type&quot; : &quot;string&quot;
}&#39;, minimum&#x3D;&#39;null&#39;, maximum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, exclusiveMaximum&#x3D;false, required&#x3D;true, deprecated&#x3D;false, hasMoreNonReadOnly&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isContainer&#x3D;false, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isFile&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isEnum&#x3D;false, isReadOnly&#x3D;false, isWriteOnly&#x3D;false, isNullable&#x3D;false, isSelfReference&#x3D;false, isCircularReference&#x3D;false, isDiscriminator&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], mostInnerItems&#x3D;null, vendorExtensions&#x3D;{}, hasValidation&#x3D;false, isInherited&#x3D;false, discriminatorValue&#x3D;&#39;null&#39;, nameInCamelCase&#x3D;&#39;q&#39;, nameInSnakeCase&#x3D;&#39;Q&#39;, enumName&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, maxProperties&#x3D;null, minProperties&#x3D;null, uniqueItems&#x3D;false, multipleOf&#x3D;null, isXmlAttribute&#x3D;false, xmlPrefix&#x3D;&#39;null&#39;, xmlName&#x3D;&#39;null&#39;, xmlNamespace&#x3D;&#39;null&#39;, isXmlWrapped&#x3D;false, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false}, content&#x3D;null}]
**** | [****](.md) |  | [CodegenParameter{isFormParam&#x3D;false, isQueryParam&#x3D;true, isPathParam&#x3D;false, isHeaderParam&#x3D;false, isCookieParam&#x3D;false, isBodyParam&#x3D;false, isContainer&#x3D;false, isCollectionFormatMulti&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isExplode&#x3D;true, baseName&#x3D;&#39;q&#39;, paramName&#x3D;&#39;q&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;null&#39;, dataFormat&#x3D;&#39;null&#39;, collectionFormat&#x3D;&#39;null&#39;, description&#x3D;&#39;Specify the search string.&#39;, unescapedDescription&#x3D;&#39;Specify the search string.&#39;, baseType&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, enumName&#x3D;&#39;null&#39;, style&#x3D;&#39;form&#39;, deepObject&#x3D;&#39;false&#39;, allowEmptyValue&#x3D;&#39;false&#39;, example&#x3D;&#39;q_example&#39;, jsonSchema&#x3D;&#39;{
  &quot;name&quot; : &quot;q&quot;,
  &quot;in&quot; : &quot;query&quot;,
  &quot;description&quot; : &quot;Specify the search string.&quot;,
  &quot;required&quot; : true,
  &quot;style&quot; : &quot;form&quot;,
  &quot;explode&quot; : true,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;string&quot;
  }
}&#39;, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isAnyType&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isFile&#x3D;false, isEnum&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, mostInnerItems&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], vendorExtensions&#x3D;{x-go-example&#x3D;&quot;q_example&quot;, x-export-param-name&#x3D;Q}, hasValidation&#x3D;false, maxProperties&#x3D;null, minProperties&#x3D;null, isNullable&#x3D;false, isDeprecated&#x3D;false, required&#x3D;true, maximum&#x3D;&#39;null&#39;, exclusiveMaximum&#x3D;false, minimum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, uniqueItems&#x3D;false, contentType&#x3D;null, multipleOf&#x3D;null, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false, schema&#x3D;CodegenProperty{openApiType&#x3D;&#39;string&#39;, baseName&#x3D;&#39;q&#39;, complexType&#x3D;&#39;null&#39;, getter&#x3D;&#39;GetQ&#39;, setter&#x3D;&#39;setQ&#39;, description&#x3D;&#39;null&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;string&#39;, dataFormat&#x3D;&#39;null&#39;, name&#x3D;&#39;Q&#39;, min&#x3D;&#39;null&#39;, max&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, defaultValueWithParam&#x3D;&#39; &#x3D; data.q;&#39;, baseType&#x3D;&#39;string&#39;, containerType&#x3D;&#39;null&#39;, title&#x3D;&#39;null&#39;, unescapedDescription&#x3D;&#39;null&#39;, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;type&quot; : &quot;string&quot;
}&#39;, minimum&#x3D;&#39;null&#39;, maximum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, exclusiveMaximum&#x3D;false, required&#x3D;true, deprecated&#x3D;false, hasMoreNonReadOnly&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isContainer&#x3D;false, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isFile&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isEnum&#x3D;false, isReadOnly&#x3D;false, isWriteOnly&#x3D;false, isNullable&#x3D;false, isSelfReference&#x3D;false, isCircularReference&#x3D;false, isDiscriminator&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], mostInnerItems&#x3D;null, vendorExtensions&#x3D;{}, hasValidation&#x3D;false, isInherited&#x3D;false, discriminatorValue&#x3D;&#39;null&#39;, nameInCamelCase&#x3D;&#39;q&#39;, nameInSnakeCase&#x3D;&#39;Q&#39;, enumName&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, maxProperties&#x3D;null, minProperties&#x3D;null, uniqueItems&#x3D;false, multipleOf&#x3D;null, isXmlAttribute&#x3D;false, xmlPrefix&#x3D;&#39;null&#39;, xmlName&#x3D;&#39;null&#39;, xmlNamespace&#x3D;&#39;null&#39;, isXmlWrapped&#x3D;false, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false}, content&#x3D;null}]

### Return type

[**SearchIdentitiesResponse**](SearchIdentitiesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

