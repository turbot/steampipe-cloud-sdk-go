# \Auth

All URIs are relative to *https://cloud.steampipe.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfirmLogin**](Auth.md#ConfirmLogin) | **Get** /login/confirm | Confirm user login
[**ConfirmSignup**](Auth.md#ConfirmSignup) | **Get** /signup/confirm | Confirm user signup
[**Login**](Auth.md#Login) | **Post** /login | User login
[**Logout**](Auth.md#Logout) | **Get** /logout/{provider} | User logout
[**Provider**](Auth.md#Provider) | **Get** /auth/{provider} | Auth Provider
[**ProviderCallback**](Auth.md#ProviderCallback) | **Get** /auth/{provider}/callback | Auth provider callback
[**Signup**](Auth.md#Signup) | **Post** /signup | User signup



## ConfirmLogin

> ConfirmLogin(ctx).T(t).Execute()

Confirm user login



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
    t := "t_example" // string | Specify the token.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Auth.ConfirmLogin(context.Background()).T(t).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.ConfirmLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfirmLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------[CodegenParameter{isFormParam&#x3D;false, isQueryParam&#x3D;true, isPathParam&#x3D;false, isHeaderParam&#x3D;false, isCookieParam&#x3D;false, isBodyParam&#x3D;false, isContainer&#x3D;false, isCollectionFormatMulti&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isExplode&#x3D;true, baseName&#x3D;&#39;t&#39;, paramName&#x3D;&#39;t&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;null&#39;, dataFormat&#x3D;&#39;null&#39;, collectionFormat&#x3D;&#39;null&#39;, description&#x3D;&#39;Specify the token.&#39;, unescapedDescription&#x3D;&#39;Specify the token.&#39;, baseType&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, enumName&#x3D;&#39;null&#39;, style&#x3D;&#39;form&#39;, deepObject&#x3D;&#39;false&#39;, allowEmptyValue&#x3D;&#39;false&#39;, example&#x3D;&#39;t_example&#39;, jsonSchema&#x3D;&#39;{
  &quot;name&quot; : &quot;t&quot;,
  &quot;in&quot; : &quot;query&quot;,
  &quot;description&quot; : &quot;Specify the token.&quot;,
  &quot;required&quot; : true,
  &quot;style&quot; : &quot;form&quot;,
  &quot;explode&quot; : true,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;string&quot;
  }
}&#39;, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isAnyType&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isFile&#x3D;false, isEnum&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, mostInnerItems&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], vendorExtensions&#x3D;{x-go-example&#x3D;&quot;t_example&quot;, x-export-param-name&#x3D;T}, hasValidation&#x3D;false, maxProperties&#x3D;null, minProperties&#x3D;null, isNullable&#x3D;false, isDeprecated&#x3D;false, required&#x3D;true, maximum&#x3D;&#39;null&#39;, exclusiveMaximum&#x3D;false, minimum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, uniqueItems&#x3D;false, contentType&#x3D;null, multipleOf&#x3D;null, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false, schema&#x3D;CodegenProperty{openApiType&#x3D;&#39;string&#39;, baseName&#x3D;&#39;t&#39;, complexType&#x3D;&#39;null&#39;, getter&#x3D;&#39;GetT&#39;, setter&#x3D;&#39;setT&#39;, description&#x3D;&#39;null&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;string&#39;, dataFormat&#x3D;&#39;null&#39;, name&#x3D;&#39;T&#39;, min&#x3D;&#39;null&#39;, max&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, defaultValueWithParam&#x3D;&#39; &#x3D; data.t;&#39;, baseType&#x3D;&#39;string&#39;, containerType&#x3D;&#39;null&#39;, title&#x3D;&#39;null&#39;, unescapedDescription&#x3D;&#39;null&#39;, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;type&quot; : &quot;string&quot;
}&#39;, minimum&#x3D;&#39;null&#39;, maximum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, exclusiveMaximum&#x3D;false, required&#x3D;true, deprecated&#x3D;false, hasMoreNonReadOnly&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isContainer&#x3D;false, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isFile&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isEnum&#x3D;false, isReadOnly&#x3D;false, isWriteOnly&#x3D;false, isNullable&#x3D;false, isSelfReference&#x3D;false, isCircularReference&#x3D;false, isDiscriminator&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], mostInnerItems&#x3D;null, vendorExtensions&#x3D;{}, hasValidation&#x3D;false, isInherited&#x3D;false, discriminatorValue&#x3D;&#39;null&#39;, nameInCamelCase&#x3D;&#39;t&#39;, nameInSnakeCase&#x3D;&#39;T&#39;, enumName&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, maxProperties&#x3D;null, minProperties&#x3D;null, uniqueItems&#x3D;false, multipleOf&#x3D;null, isXmlAttribute&#x3D;false, xmlPrefix&#x3D;&#39;null&#39;, xmlName&#x3D;&#39;null&#39;, xmlNamespace&#x3D;&#39;null&#39;, isXmlWrapped&#x3D;false, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false}, content&#x3D;null}]
**** | [****](.md) |  | [CodegenParameter{isFormParam&#x3D;false, isQueryParam&#x3D;true, isPathParam&#x3D;false, isHeaderParam&#x3D;false, isCookieParam&#x3D;false, isBodyParam&#x3D;false, isContainer&#x3D;false, isCollectionFormatMulti&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isExplode&#x3D;true, baseName&#x3D;&#39;t&#39;, paramName&#x3D;&#39;t&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;null&#39;, dataFormat&#x3D;&#39;null&#39;, collectionFormat&#x3D;&#39;null&#39;, description&#x3D;&#39;Specify the token.&#39;, unescapedDescription&#x3D;&#39;Specify the token.&#39;, baseType&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, enumName&#x3D;&#39;null&#39;, style&#x3D;&#39;form&#39;, deepObject&#x3D;&#39;false&#39;, allowEmptyValue&#x3D;&#39;false&#39;, example&#x3D;&#39;t_example&#39;, jsonSchema&#x3D;&#39;{
  &quot;name&quot; : &quot;t&quot;,
  &quot;in&quot; : &quot;query&quot;,
  &quot;description&quot; : &quot;Specify the token.&quot;,
  &quot;required&quot; : true,
  &quot;style&quot; : &quot;form&quot;,
  &quot;explode&quot; : true,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;string&quot;
  }
}&#39;, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isAnyType&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isFile&#x3D;false, isEnum&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, mostInnerItems&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], vendorExtensions&#x3D;{x-go-example&#x3D;&quot;t_example&quot;, x-export-param-name&#x3D;T}, hasValidation&#x3D;false, maxProperties&#x3D;null, minProperties&#x3D;null, isNullable&#x3D;false, isDeprecated&#x3D;false, required&#x3D;true, maximum&#x3D;&#39;null&#39;, exclusiveMaximum&#x3D;false, minimum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, uniqueItems&#x3D;false, contentType&#x3D;null, multipleOf&#x3D;null, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false, schema&#x3D;CodegenProperty{openApiType&#x3D;&#39;string&#39;, baseName&#x3D;&#39;t&#39;, complexType&#x3D;&#39;null&#39;, getter&#x3D;&#39;GetT&#39;, setter&#x3D;&#39;setT&#39;, description&#x3D;&#39;null&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;string&#39;, dataFormat&#x3D;&#39;null&#39;, name&#x3D;&#39;T&#39;, min&#x3D;&#39;null&#39;, max&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, defaultValueWithParam&#x3D;&#39; &#x3D; data.t;&#39;, baseType&#x3D;&#39;string&#39;, containerType&#x3D;&#39;null&#39;, title&#x3D;&#39;null&#39;, unescapedDescription&#x3D;&#39;null&#39;, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;type&quot; : &quot;string&quot;
}&#39;, minimum&#x3D;&#39;null&#39;, maximum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, exclusiveMaximum&#x3D;false, required&#x3D;true, deprecated&#x3D;false, hasMoreNonReadOnly&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isContainer&#x3D;false, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isFile&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isEnum&#x3D;false, isReadOnly&#x3D;false, isWriteOnly&#x3D;false, isNullable&#x3D;false, isSelfReference&#x3D;false, isCircularReference&#x3D;false, isDiscriminator&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], mostInnerItems&#x3D;null, vendorExtensions&#x3D;{}, hasValidation&#x3D;false, isInherited&#x3D;false, discriminatorValue&#x3D;&#39;null&#39;, nameInCamelCase&#x3D;&#39;t&#39;, nameInSnakeCase&#x3D;&#39;T&#39;, enumName&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, maxProperties&#x3D;null, minProperties&#x3D;null, uniqueItems&#x3D;false, multipleOf&#x3D;null, isXmlAttribute&#x3D;false, xmlPrefix&#x3D;&#39;null&#39;, xmlName&#x3D;&#39;null&#39;, xmlNamespace&#x3D;&#39;null&#39;, isXmlWrapped&#x3D;false, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false}, content&#x3D;null}]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfirmSignup

> ConfirmSignup(ctx).T(t).Execute()

Confirm user signup



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
    t := "t_example" // string | Specify the token.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Auth.ConfirmSignup(context.Background()).T(t).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.ConfirmSignup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfirmSignupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------[CodegenParameter{isFormParam&#x3D;false, isQueryParam&#x3D;true, isPathParam&#x3D;false, isHeaderParam&#x3D;false, isCookieParam&#x3D;false, isBodyParam&#x3D;false, isContainer&#x3D;false, isCollectionFormatMulti&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isExplode&#x3D;true, baseName&#x3D;&#39;t&#39;, paramName&#x3D;&#39;t&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;null&#39;, dataFormat&#x3D;&#39;null&#39;, collectionFormat&#x3D;&#39;null&#39;, description&#x3D;&#39;Specify the token.&#39;, unescapedDescription&#x3D;&#39;Specify the token.&#39;, baseType&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, enumName&#x3D;&#39;null&#39;, style&#x3D;&#39;form&#39;, deepObject&#x3D;&#39;false&#39;, allowEmptyValue&#x3D;&#39;false&#39;, example&#x3D;&#39;t_example&#39;, jsonSchema&#x3D;&#39;{
  &quot;name&quot; : &quot;t&quot;,
  &quot;in&quot; : &quot;query&quot;,
  &quot;description&quot; : &quot;Specify the token.&quot;,
  &quot;required&quot; : true,
  &quot;style&quot; : &quot;form&quot;,
  &quot;explode&quot; : true,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;string&quot;
  }
}&#39;, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isAnyType&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isFile&#x3D;false, isEnum&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, mostInnerItems&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], vendorExtensions&#x3D;{x-go-example&#x3D;&quot;t_example&quot;, x-export-param-name&#x3D;T}, hasValidation&#x3D;false, maxProperties&#x3D;null, minProperties&#x3D;null, isNullable&#x3D;false, isDeprecated&#x3D;false, required&#x3D;true, maximum&#x3D;&#39;null&#39;, exclusiveMaximum&#x3D;false, minimum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, uniqueItems&#x3D;false, contentType&#x3D;null, multipleOf&#x3D;null, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false, schema&#x3D;CodegenProperty{openApiType&#x3D;&#39;string&#39;, baseName&#x3D;&#39;t&#39;, complexType&#x3D;&#39;null&#39;, getter&#x3D;&#39;GetT&#39;, setter&#x3D;&#39;setT&#39;, description&#x3D;&#39;null&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;string&#39;, dataFormat&#x3D;&#39;null&#39;, name&#x3D;&#39;T&#39;, min&#x3D;&#39;null&#39;, max&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, defaultValueWithParam&#x3D;&#39; &#x3D; data.t;&#39;, baseType&#x3D;&#39;string&#39;, containerType&#x3D;&#39;null&#39;, title&#x3D;&#39;null&#39;, unescapedDescription&#x3D;&#39;null&#39;, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;type&quot; : &quot;string&quot;
}&#39;, minimum&#x3D;&#39;null&#39;, maximum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, exclusiveMaximum&#x3D;false, required&#x3D;true, deprecated&#x3D;false, hasMoreNonReadOnly&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isContainer&#x3D;false, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isFile&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isEnum&#x3D;false, isReadOnly&#x3D;false, isWriteOnly&#x3D;false, isNullable&#x3D;false, isSelfReference&#x3D;false, isCircularReference&#x3D;false, isDiscriminator&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], mostInnerItems&#x3D;null, vendorExtensions&#x3D;{}, hasValidation&#x3D;false, isInherited&#x3D;false, discriminatorValue&#x3D;&#39;null&#39;, nameInCamelCase&#x3D;&#39;t&#39;, nameInSnakeCase&#x3D;&#39;T&#39;, enumName&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, maxProperties&#x3D;null, minProperties&#x3D;null, uniqueItems&#x3D;false, multipleOf&#x3D;null, isXmlAttribute&#x3D;false, xmlPrefix&#x3D;&#39;null&#39;, xmlName&#x3D;&#39;null&#39;, xmlNamespace&#x3D;&#39;null&#39;, isXmlWrapped&#x3D;false, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false}, content&#x3D;null}]
**** | [****](.md) |  | [CodegenParameter{isFormParam&#x3D;false, isQueryParam&#x3D;true, isPathParam&#x3D;false, isHeaderParam&#x3D;false, isCookieParam&#x3D;false, isBodyParam&#x3D;false, isContainer&#x3D;false, isCollectionFormatMulti&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isExplode&#x3D;true, baseName&#x3D;&#39;t&#39;, paramName&#x3D;&#39;t&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;null&#39;, dataFormat&#x3D;&#39;null&#39;, collectionFormat&#x3D;&#39;null&#39;, description&#x3D;&#39;Specify the token.&#39;, unescapedDescription&#x3D;&#39;Specify the token.&#39;, baseType&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, enumName&#x3D;&#39;null&#39;, style&#x3D;&#39;form&#39;, deepObject&#x3D;&#39;false&#39;, allowEmptyValue&#x3D;&#39;false&#39;, example&#x3D;&#39;t_example&#39;, jsonSchema&#x3D;&#39;{
  &quot;name&quot; : &quot;t&quot;,
  &quot;in&quot; : &quot;query&quot;,
  &quot;description&quot; : &quot;Specify the token.&quot;,
  &quot;required&quot; : true,
  &quot;style&quot; : &quot;form&quot;,
  &quot;explode&quot; : true,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;string&quot;
  }
}&#39;, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isAnyType&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isFile&#x3D;false, isEnum&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, mostInnerItems&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], vendorExtensions&#x3D;{x-go-example&#x3D;&quot;t_example&quot;, x-export-param-name&#x3D;T}, hasValidation&#x3D;false, maxProperties&#x3D;null, minProperties&#x3D;null, isNullable&#x3D;false, isDeprecated&#x3D;false, required&#x3D;true, maximum&#x3D;&#39;null&#39;, exclusiveMaximum&#x3D;false, minimum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, uniqueItems&#x3D;false, contentType&#x3D;null, multipleOf&#x3D;null, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false, schema&#x3D;CodegenProperty{openApiType&#x3D;&#39;string&#39;, baseName&#x3D;&#39;t&#39;, complexType&#x3D;&#39;null&#39;, getter&#x3D;&#39;GetT&#39;, setter&#x3D;&#39;setT&#39;, description&#x3D;&#39;null&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;string&#39;, dataFormat&#x3D;&#39;null&#39;, name&#x3D;&#39;T&#39;, min&#x3D;&#39;null&#39;, max&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, defaultValueWithParam&#x3D;&#39; &#x3D; data.t;&#39;, baseType&#x3D;&#39;string&#39;, containerType&#x3D;&#39;null&#39;, title&#x3D;&#39;null&#39;, unescapedDescription&#x3D;&#39;null&#39;, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;type&quot; : &quot;string&quot;
}&#39;, minimum&#x3D;&#39;null&#39;, maximum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, exclusiveMaximum&#x3D;false, required&#x3D;true, deprecated&#x3D;false, hasMoreNonReadOnly&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isContainer&#x3D;false, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isFile&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isEnum&#x3D;false, isReadOnly&#x3D;false, isWriteOnly&#x3D;false, isNullable&#x3D;false, isSelfReference&#x3D;false, isCircularReference&#x3D;false, isDiscriminator&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], mostInnerItems&#x3D;null, vendorExtensions&#x3D;{}, hasValidation&#x3D;false, isInherited&#x3D;false, discriminatorValue&#x3D;&#39;null&#39;, nameInCamelCase&#x3D;&#39;t&#39;, nameInSnakeCase&#x3D;&#39;T&#39;, enumName&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, maxProperties&#x3D;null, minProperties&#x3D;null, uniqueItems&#x3D;false, multipleOf&#x3D;null, isXmlAttribute&#x3D;false, xmlPrefix&#x3D;&#39;null&#39;, xmlName&#x3D;&#39;null&#39;, xmlNamespace&#x3D;&#39;null&#39;, isXmlWrapped&#x3D;false, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false}, content&#x3D;null}]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Login

> Login(ctx).Request(request).Execute()

User login



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
    request := *openapiclient.NewUserLoginRequest("Email_example") // UserLoginRequest | The request body to login.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Auth.Login(context.Background()).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.Login``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------[CodegenParameter{isFormParam&#x3D;false, isQueryParam&#x3D;false, isPathParam&#x3D;false, isHeaderParam&#x3D;false, isCookieParam&#x3D;false, isBodyParam&#x3D;true, isContainer&#x3D;false, isCollectionFormatMulti&#x3D;false, isPrimitiveType&#x3D;false, isModel&#x3D;true, isExplode&#x3D;false, baseName&#x3D;&#39;request&#39;, paramName&#x3D;&#39;request&#39;, dataType&#x3D;&#39;UserLoginRequest&#39;, datatypeWithEnum&#x3D;&#39;null&#39;, dataFormat&#x3D;&#39;null&#39;, collectionFormat&#x3D;&#39;null&#39;, description&#x3D;&#39;The request body to login.&#39;, unescapedDescription&#x3D;&#39;null&#39;, baseType&#x3D;&#39;UserLoginRequest&#39;, defaultValue&#x3D;&#39;null&#39;, enumName&#x3D;&#39;null&#39;, style&#x3D;&#39;null&#39;, deepObject&#x3D;&#39;false&#39;, allowEmptyValue&#x3D;&#39;false&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;description&quot; : &quot;The request body to login.&quot;,
  &quot;content&quot; : {
    &quot;application/json&quot; : {
      &quot;schema&quot; : {
        &quot;$ref&quot; : &quot;#/components/schemas/UserLoginRequest&quot;
      }
    }
  },
  &quot;required&quot; : true
}&#39;, isString&#x3D;false, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isAnyType&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isFile&#x3D;false, isEnum&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, mostInnerItems&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[CodegenProperty{openApiType&#x3D;&#39;string&#39;, baseName&#x3D;&#39;email&#39;, complexType&#x3D;&#39;null&#39;, getter&#x3D;&#39;GetEmail&#39;, setter&#x3D;&#39;setEmail&#39;, description&#x3D;&#39;null&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;string&#39;, dataFormat&#x3D;&#39;null&#39;, name&#x3D;&#39;Email&#39;, min&#x3D;&#39;null&#39;, max&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, defaultValueWithParam&#x3D;&#39; &#x3D; data.email;&#39;, baseType&#x3D;&#39;string&#39;, containerType&#x3D;&#39;null&#39;, title&#x3D;&#39;null&#39;, unescapedDescription&#x3D;&#39;null&#39;, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;type&quot; : &quot;string&quot;
}&#39;, minimum&#x3D;&#39;null&#39;, maximum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, exclusiveMaximum&#x3D;false, required&#x3D;true, deprecated&#x3D;false, hasMoreNonReadOnly&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isContainer&#x3D;false, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isFile&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isEnum&#x3D;false, isReadOnly&#x3D;false, isWriteOnly&#x3D;false, isNullable&#x3D;false, isSelfReference&#x3D;false, isCircularReference&#x3D;false, isDiscriminator&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], mostInnerItems&#x3D;null, vendorExtensions&#x3D;{}, hasValidation&#x3D;false, isInherited&#x3D;false, discriminatorValue&#x3D;&#39;null&#39;, nameInCamelCase&#x3D;&#39;email&#39;, nameInSnakeCase&#x3D;&#39;EMAIL&#39;, enumName&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, maxProperties&#x3D;null, minProperties&#x3D;null, uniqueItems&#x3D;false, multipleOf&#x3D;null, isXmlAttribute&#x3D;false, xmlPrefix&#x3D;&#39;null&#39;, xmlName&#x3D;&#39;null&#39;, xmlNamespace&#x3D;&#39;null&#39;, isXmlWrapped&#x3D;false, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false}], requiredVars&#x3D;[CodegenProperty{openApiType&#x3D;&#39;string&#39;, baseName&#x3D;&#39;email&#39;, complexType&#x3D;&#39;null&#39;, getter&#x3D;&#39;GetEmail&#39;, setter&#x3D;&#39;setEmail&#39;, description&#x3D;&#39;null&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;string&#39;, dataFormat&#x3D;&#39;null&#39;, name&#x3D;&#39;Email&#39;, min&#x3D;&#39;null&#39;, max&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, defaultValueWithParam&#x3D;&#39; &#x3D; data.email;&#39;, baseType&#x3D;&#39;string&#39;, containerType&#x3D;&#39;null&#39;, title&#x3D;&#39;null&#39;, unescapedDescription&#x3D;&#39;null&#39;, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;type&quot; : &quot;string&quot;
}&#39;, minimum&#x3D;&#39;null&#39;, maximum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, exclusiveMaximum&#x3D;false, required&#x3D;true, deprecated&#x3D;false, hasMoreNonReadOnly&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isContainer&#x3D;false, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isFile&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isEnum&#x3D;false, isReadOnly&#x3D;false, isWriteOnly&#x3D;false, isNullable&#x3D;false, isSelfReference&#x3D;false, isCircularReference&#x3D;false, isDiscriminator&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], mostInnerItems&#x3D;null, vendorExtensions&#x3D;{}, hasValidation&#x3D;false, isInherited&#x3D;false, discriminatorValue&#x3D;&#39;null&#39;, nameInCamelCase&#x3D;&#39;email&#39;, nameInSnakeCase&#x3D;&#39;EMAIL&#39;, enumName&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, maxProperties&#x3D;null, minProperties&#x3D;null, uniqueItems&#x3D;false, multipleOf&#x3D;null, isXmlAttribute&#x3D;false, xmlPrefix&#x3D;&#39;null&#39;, xmlName&#x3D;&#39;null&#39;, xmlNamespace&#x3D;&#39;null&#39;, isXmlWrapped&#x3D;false, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false}], vendorExtensions&#x3D;{x-go-example&#x3D;*openapiclient.NewUserLoginRequest(&quot;Email_example&quot;), x-export-param-name&#x3D;Request}, hasValidation&#x3D;false, maxProperties&#x3D;null, minProperties&#x3D;null, isNullable&#x3D;false, isDeprecated&#x3D;false, required&#x3D;true, maximum&#x3D;&#39;null&#39;, exclusiveMaximum&#x3D;false, minimum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, uniqueItems&#x3D;false, contentType&#x3D;null, multipleOf&#x3D;null, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;true, getHasRequired&#x3D;true, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false, schema&#x3D;null, content&#x3D;{application/json&#x3D;CodegenMediaType{schema&#x3D;CodegenProperty{openApiType&#x3D;&#39;UserLoginRequest&#39;, baseName&#x3D;&#39;SchemaForRequestBodyApplicationJson&#39;, complexType&#x3D;&#39;UserLoginRequest&#39;, getter&#x3D;&#39;GetSchemaForRequestBodyApplicationJson&#39;, setter&#x3D;&#39;setSchemaForRequestBodyApplicationJson&#39;, description&#x3D;&#39;null&#39;, dataType&#x3D;&#39;UserLoginRequest&#39;, datatypeWithEnum&#x3D;&#39;UserLoginRequest&#39;, dataFormat&#x3D;&#39;null&#39;, name&#x3D;&#39;SchemaForRequestBodyApplicationJson&#39;, min&#x3D;&#39;null&#39;, max&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, defaultValueWithParam&#x3D;&#39; &#x3D; data.SchemaForRequestBodyApplicationJson;&#39;, baseType&#x3D;&#39;UserLoginRequest&#39;, containerType&#x3D;&#39;null&#39;, title&#x3D;&#39;null&#39;, unescapedDescription&#x3D;&#39;null&#39;, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;$ref&quot; : &quot;#/components/schemas/UserLoginRequest&quot;
}&#39;, minimum&#x3D;&#39;null&#39;, maximum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, exclusiveMaximum&#x3D;false, required&#x3D;false, deprecated&#x3D;false, hasMoreNonReadOnly&#x3D;false, isPrimitiveType&#x3D;false, isModel&#x3D;true, isContainer&#x3D;false, isString&#x3D;false, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isFile&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isEnum&#x3D;false, isReadOnly&#x3D;false, isWriteOnly&#x3D;false, isNullable&#x3D;false, isSelfReference&#x3D;false, isCircularReference&#x3D;false, isDiscriminator&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], mostInnerItems&#x3D;null, vendorExtensions&#x3D;{}, hasValidation&#x3D;false, isInherited&#x3D;false, discriminatorValue&#x3D;&#39;null&#39;, nameInCamelCase&#x3D;&#39;schemaForRequestBodyApplicationJson&#39;, nameInSnakeCase&#x3D;&#39;SCHEMA_FOR_REQUEST_BODY_APPLICATION_JSON&#39;, enumName&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, maxProperties&#x3D;null, minProperties&#x3D;null, uniqueItems&#x3D;false, multipleOf&#x3D;null, isXmlAttribute&#x3D;false, xmlPrefix&#x3D;&#39;null&#39;, xmlName&#x3D;&#39;null&#39;, xmlNamespace&#x3D;&#39;null&#39;, isXmlWrapped&#x3D;false, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false}, encoding&#x3D;null}}}]
**** | [****](.md) |  | [CodegenParameter{isFormParam&#x3D;false, isQueryParam&#x3D;false, isPathParam&#x3D;false, isHeaderParam&#x3D;false, isCookieParam&#x3D;false, isBodyParam&#x3D;true, isContainer&#x3D;false, isCollectionFormatMulti&#x3D;false, isPrimitiveType&#x3D;false, isModel&#x3D;true, isExplode&#x3D;false, baseName&#x3D;&#39;request&#39;, paramName&#x3D;&#39;request&#39;, dataType&#x3D;&#39;UserLoginRequest&#39;, datatypeWithEnum&#x3D;&#39;null&#39;, dataFormat&#x3D;&#39;null&#39;, collectionFormat&#x3D;&#39;null&#39;, description&#x3D;&#39;The request body to login.&#39;, unescapedDescription&#x3D;&#39;null&#39;, baseType&#x3D;&#39;UserLoginRequest&#39;, defaultValue&#x3D;&#39;null&#39;, enumName&#x3D;&#39;null&#39;, style&#x3D;&#39;null&#39;, deepObject&#x3D;&#39;false&#39;, allowEmptyValue&#x3D;&#39;false&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;description&quot; : &quot;The request body to login.&quot;,
  &quot;content&quot; : {
    &quot;application/json&quot; : {
      &quot;schema&quot; : {
        &quot;$ref&quot; : &quot;#/components/schemas/UserLoginRequest&quot;
      }
    }
  },
  &quot;required&quot; : true
}&#39;, isString&#x3D;false, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isAnyType&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isFile&#x3D;false, isEnum&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, mostInnerItems&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[CodegenProperty{openApiType&#x3D;&#39;string&#39;, baseName&#x3D;&#39;email&#39;, complexType&#x3D;&#39;null&#39;, getter&#x3D;&#39;GetEmail&#39;, setter&#x3D;&#39;setEmail&#39;, description&#x3D;&#39;null&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;string&#39;, dataFormat&#x3D;&#39;null&#39;, name&#x3D;&#39;Email&#39;, min&#x3D;&#39;null&#39;, max&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, defaultValueWithParam&#x3D;&#39; &#x3D; data.email;&#39;, baseType&#x3D;&#39;string&#39;, containerType&#x3D;&#39;null&#39;, title&#x3D;&#39;null&#39;, unescapedDescription&#x3D;&#39;null&#39;, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;type&quot; : &quot;string&quot;
}&#39;, minimum&#x3D;&#39;null&#39;, maximum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, exclusiveMaximum&#x3D;false, required&#x3D;true, deprecated&#x3D;false, hasMoreNonReadOnly&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isContainer&#x3D;false, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isFile&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isEnum&#x3D;false, isReadOnly&#x3D;false, isWriteOnly&#x3D;false, isNullable&#x3D;false, isSelfReference&#x3D;false, isCircularReference&#x3D;false, isDiscriminator&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], mostInnerItems&#x3D;null, vendorExtensions&#x3D;{}, hasValidation&#x3D;false, isInherited&#x3D;false, discriminatorValue&#x3D;&#39;null&#39;, nameInCamelCase&#x3D;&#39;email&#39;, nameInSnakeCase&#x3D;&#39;EMAIL&#39;, enumName&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, maxProperties&#x3D;null, minProperties&#x3D;null, uniqueItems&#x3D;false, multipleOf&#x3D;null, isXmlAttribute&#x3D;false, xmlPrefix&#x3D;&#39;null&#39;, xmlName&#x3D;&#39;null&#39;, xmlNamespace&#x3D;&#39;null&#39;, isXmlWrapped&#x3D;false, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false}], requiredVars&#x3D;[CodegenProperty{openApiType&#x3D;&#39;string&#39;, baseName&#x3D;&#39;email&#39;, complexType&#x3D;&#39;null&#39;, getter&#x3D;&#39;GetEmail&#39;, setter&#x3D;&#39;setEmail&#39;, description&#x3D;&#39;null&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;string&#39;, dataFormat&#x3D;&#39;null&#39;, name&#x3D;&#39;Email&#39;, min&#x3D;&#39;null&#39;, max&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, defaultValueWithParam&#x3D;&#39; &#x3D; data.email;&#39;, baseType&#x3D;&#39;string&#39;, containerType&#x3D;&#39;null&#39;, title&#x3D;&#39;null&#39;, unescapedDescription&#x3D;&#39;null&#39;, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;type&quot; : &quot;string&quot;
}&#39;, minimum&#x3D;&#39;null&#39;, maximum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, exclusiveMaximum&#x3D;false, required&#x3D;true, deprecated&#x3D;false, hasMoreNonReadOnly&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isContainer&#x3D;false, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isFile&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isEnum&#x3D;false, isReadOnly&#x3D;false, isWriteOnly&#x3D;false, isNullable&#x3D;false, isSelfReference&#x3D;false, isCircularReference&#x3D;false, isDiscriminator&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], mostInnerItems&#x3D;null, vendorExtensions&#x3D;{}, hasValidation&#x3D;false, isInherited&#x3D;false, discriminatorValue&#x3D;&#39;null&#39;, nameInCamelCase&#x3D;&#39;email&#39;, nameInSnakeCase&#x3D;&#39;EMAIL&#39;, enumName&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, maxProperties&#x3D;null, minProperties&#x3D;null, uniqueItems&#x3D;false, multipleOf&#x3D;null, isXmlAttribute&#x3D;false, xmlPrefix&#x3D;&#39;null&#39;, xmlName&#x3D;&#39;null&#39;, xmlNamespace&#x3D;&#39;null&#39;, isXmlWrapped&#x3D;false, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false}], vendorExtensions&#x3D;{x-go-example&#x3D;*openapiclient.NewUserLoginRequest(&quot;Email_example&quot;), x-export-param-name&#x3D;Request}, hasValidation&#x3D;false, maxProperties&#x3D;null, minProperties&#x3D;null, isNullable&#x3D;false, isDeprecated&#x3D;false, required&#x3D;true, maximum&#x3D;&#39;null&#39;, exclusiveMaximum&#x3D;false, minimum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, uniqueItems&#x3D;false, contentType&#x3D;null, multipleOf&#x3D;null, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;true, getHasRequired&#x3D;true, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false, schema&#x3D;null, content&#x3D;{application/json&#x3D;CodegenMediaType{schema&#x3D;CodegenProperty{openApiType&#x3D;&#39;UserLoginRequest&#39;, baseName&#x3D;&#39;SchemaForRequestBodyApplicationJson&#39;, complexType&#x3D;&#39;UserLoginRequest&#39;, getter&#x3D;&#39;GetSchemaForRequestBodyApplicationJson&#39;, setter&#x3D;&#39;setSchemaForRequestBodyApplicationJson&#39;, description&#x3D;&#39;null&#39;, dataType&#x3D;&#39;UserLoginRequest&#39;, datatypeWithEnum&#x3D;&#39;UserLoginRequest&#39;, dataFormat&#x3D;&#39;null&#39;, name&#x3D;&#39;SchemaForRequestBodyApplicationJson&#39;, min&#x3D;&#39;null&#39;, max&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, defaultValueWithParam&#x3D;&#39; &#x3D; data.SchemaForRequestBodyApplicationJson;&#39;, baseType&#x3D;&#39;UserLoginRequest&#39;, containerType&#x3D;&#39;null&#39;, title&#x3D;&#39;null&#39;, unescapedDescription&#x3D;&#39;null&#39;, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;$ref&quot; : &quot;#/components/schemas/UserLoginRequest&quot;
}&#39;, minimum&#x3D;&#39;null&#39;, maximum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, exclusiveMaximum&#x3D;false, required&#x3D;false, deprecated&#x3D;false, hasMoreNonReadOnly&#x3D;false, isPrimitiveType&#x3D;false, isModel&#x3D;true, isContainer&#x3D;false, isString&#x3D;false, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isFile&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isEnum&#x3D;false, isReadOnly&#x3D;false, isWriteOnly&#x3D;false, isNullable&#x3D;false, isSelfReference&#x3D;false, isCircularReference&#x3D;false, isDiscriminator&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], mostInnerItems&#x3D;null, vendorExtensions&#x3D;{}, hasValidation&#x3D;false, isInherited&#x3D;false, discriminatorValue&#x3D;&#39;null&#39;, nameInCamelCase&#x3D;&#39;schemaForRequestBodyApplicationJson&#39;, nameInSnakeCase&#x3D;&#39;SCHEMA_FOR_REQUEST_BODY_APPLICATION_JSON&#39;, enumName&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, maxProperties&#x3D;null, minProperties&#x3D;null, uniqueItems&#x3D;false, multipleOf&#x3D;null, isXmlAttribute&#x3D;false, xmlPrefix&#x3D;&#39;null&#39;, xmlName&#x3D;&#39;null&#39;, xmlNamespace&#x3D;&#39;null&#39;, isXmlWrapped&#x3D;false, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false}, encoding&#x3D;null}}}]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Logout

> Logout(ctx, provider).Execute()

User logout



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
    provider := "provider_example" // string | The authentication provider.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Auth.Logout(context.Background(), provider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.Logout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | The authentication provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLogoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------[CodegenParameter{isFormParam&#x3D;false, isQueryParam&#x3D;false, isPathParam&#x3D;true, isHeaderParam&#x3D;false, isCookieParam&#x3D;false, isBodyParam&#x3D;false, isContainer&#x3D;false, isCollectionFormatMulti&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isExplode&#x3D;false, baseName&#x3D;&#39;provider&#39;, paramName&#x3D;&#39;provider&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;null&#39;, dataFormat&#x3D;&#39;null&#39;, collectionFormat&#x3D;&#39;null&#39;, description&#x3D;&#39;The authentication provider.&#39;, unescapedDescription&#x3D;&#39;The authentication provider.&#39;, baseType&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, enumName&#x3D;&#39;null&#39;, style&#x3D;&#39;simple&#39;, deepObject&#x3D;&#39;false&#39;, allowEmptyValue&#x3D;&#39;false&#39;, example&#x3D;&#39;provider_example&#39;, jsonSchema&#x3D;&#39;{
  &quot;name&quot; : &quot;provider&quot;,
  &quot;in&quot; : &quot;path&quot;,
  &quot;description&quot; : &quot;The authentication provider.&quot;,
  &quot;required&quot; : true,
  &quot;style&quot; : &quot;simple&quot;,
  &quot;explode&quot; : false,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;string&quot;
  }
}&#39;, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isAnyType&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isFile&#x3D;false, isEnum&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, mostInnerItems&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], vendorExtensions&#x3D;{x-go-example&#x3D;&quot;provider_example&quot;, x-export-param-name&#x3D;Provider}, hasValidation&#x3D;false, maxProperties&#x3D;null, minProperties&#x3D;null, isNullable&#x3D;false, isDeprecated&#x3D;false, required&#x3D;true, maximum&#x3D;&#39;null&#39;, exclusiveMaximum&#x3D;false, minimum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, uniqueItems&#x3D;false, contentType&#x3D;null, multipleOf&#x3D;null, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false, schema&#x3D;CodegenProperty{openApiType&#x3D;&#39;string&#39;, baseName&#x3D;&#39;provider&#39;, complexType&#x3D;&#39;null&#39;, getter&#x3D;&#39;GetProvider&#39;, setter&#x3D;&#39;setProvider&#39;, description&#x3D;&#39;null&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;string&#39;, dataFormat&#x3D;&#39;null&#39;, name&#x3D;&#39;Provider&#39;, min&#x3D;&#39;null&#39;, max&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, defaultValueWithParam&#x3D;&#39; &#x3D; data.provider;&#39;, baseType&#x3D;&#39;string&#39;, containerType&#x3D;&#39;null&#39;, title&#x3D;&#39;null&#39;, unescapedDescription&#x3D;&#39;null&#39;, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;type&quot; : &quot;string&quot;
}&#39;, minimum&#x3D;&#39;null&#39;, maximum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, exclusiveMaximum&#x3D;false, required&#x3D;true, deprecated&#x3D;false, hasMoreNonReadOnly&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isContainer&#x3D;false, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isFile&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isEnum&#x3D;false, isReadOnly&#x3D;false, isWriteOnly&#x3D;false, isNullable&#x3D;false, isSelfReference&#x3D;false, isCircularReference&#x3D;false, isDiscriminator&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], mostInnerItems&#x3D;null, vendorExtensions&#x3D;{}, hasValidation&#x3D;false, isInherited&#x3D;false, discriminatorValue&#x3D;&#39;null&#39;, nameInCamelCase&#x3D;&#39;provider&#39;, nameInSnakeCase&#x3D;&#39;PROVIDER&#39;, enumName&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, maxProperties&#x3D;null, minProperties&#x3D;null, uniqueItems&#x3D;false, multipleOf&#x3D;null, isXmlAttribute&#x3D;false, xmlPrefix&#x3D;&#39;null&#39;, xmlName&#x3D;&#39;null&#39;, xmlNamespace&#x3D;&#39;null&#39;, isXmlWrapped&#x3D;false, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false}, content&#x3D;null}]
**** | [****](.md) |  | [CodegenParameter{isFormParam&#x3D;false, isQueryParam&#x3D;false, isPathParam&#x3D;true, isHeaderParam&#x3D;false, isCookieParam&#x3D;false, isBodyParam&#x3D;false, isContainer&#x3D;false, isCollectionFormatMulti&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isExplode&#x3D;false, baseName&#x3D;&#39;provider&#39;, paramName&#x3D;&#39;provider&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;null&#39;, dataFormat&#x3D;&#39;null&#39;, collectionFormat&#x3D;&#39;null&#39;, description&#x3D;&#39;The authentication provider.&#39;, unescapedDescription&#x3D;&#39;The authentication provider.&#39;, baseType&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, enumName&#x3D;&#39;null&#39;, style&#x3D;&#39;simple&#39;, deepObject&#x3D;&#39;false&#39;, allowEmptyValue&#x3D;&#39;false&#39;, example&#x3D;&#39;provider_example&#39;, jsonSchema&#x3D;&#39;{
  &quot;name&quot; : &quot;provider&quot;,
  &quot;in&quot; : &quot;path&quot;,
  &quot;description&quot; : &quot;The authentication provider.&quot;,
  &quot;required&quot; : true,
  &quot;style&quot; : &quot;simple&quot;,
  &quot;explode&quot; : false,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;string&quot;
  }
}&#39;, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isAnyType&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isFile&#x3D;false, isEnum&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, mostInnerItems&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], vendorExtensions&#x3D;{x-go-example&#x3D;&quot;provider_example&quot;, x-export-param-name&#x3D;Provider}, hasValidation&#x3D;false, maxProperties&#x3D;null, minProperties&#x3D;null, isNullable&#x3D;false, isDeprecated&#x3D;false, required&#x3D;true, maximum&#x3D;&#39;null&#39;, exclusiveMaximum&#x3D;false, minimum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, uniqueItems&#x3D;false, contentType&#x3D;null, multipleOf&#x3D;null, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false, schema&#x3D;CodegenProperty{openApiType&#x3D;&#39;string&#39;, baseName&#x3D;&#39;provider&#39;, complexType&#x3D;&#39;null&#39;, getter&#x3D;&#39;GetProvider&#39;, setter&#x3D;&#39;setProvider&#39;, description&#x3D;&#39;null&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;string&#39;, dataFormat&#x3D;&#39;null&#39;, name&#x3D;&#39;Provider&#39;, min&#x3D;&#39;null&#39;, max&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, defaultValueWithParam&#x3D;&#39; &#x3D; data.provider;&#39;, baseType&#x3D;&#39;string&#39;, containerType&#x3D;&#39;null&#39;, title&#x3D;&#39;null&#39;, unescapedDescription&#x3D;&#39;null&#39;, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;type&quot; : &quot;string&quot;
}&#39;, minimum&#x3D;&#39;null&#39;, maximum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, exclusiveMaximum&#x3D;false, required&#x3D;true, deprecated&#x3D;false, hasMoreNonReadOnly&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isContainer&#x3D;false, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isFile&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isEnum&#x3D;false, isReadOnly&#x3D;false, isWriteOnly&#x3D;false, isNullable&#x3D;false, isSelfReference&#x3D;false, isCircularReference&#x3D;false, isDiscriminator&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], mostInnerItems&#x3D;null, vendorExtensions&#x3D;{}, hasValidation&#x3D;false, isInherited&#x3D;false, discriminatorValue&#x3D;&#39;null&#39;, nameInCamelCase&#x3D;&#39;provider&#39;, nameInSnakeCase&#x3D;&#39;PROVIDER&#39;, enumName&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, maxProperties&#x3D;null, minProperties&#x3D;null, uniqueItems&#x3D;false, multipleOf&#x3D;null, isXmlAttribute&#x3D;false, xmlPrefix&#x3D;&#39;null&#39;, xmlName&#x3D;&#39;null&#39;, xmlNamespace&#x3D;&#39;null&#39;, isXmlWrapped&#x3D;false, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false}, content&#x3D;null}]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Provider

> Provider(ctx, provider).Execute()

Auth Provider



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
    provider := "provider_example" // string | The authentication provider.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Auth.Provider(context.Background(), provider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.Provider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | The authentication provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------[CodegenParameter{isFormParam&#x3D;false, isQueryParam&#x3D;false, isPathParam&#x3D;true, isHeaderParam&#x3D;false, isCookieParam&#x3D;false, isBodyParam&#x3D;false, isContainer&#x3D;false, isCollectionFormatMulti&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isExplode&#x3D;false, baseName&#x3D;&#39;provider&#39;, paramName&#x3D;&#39;provider&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;null&#39;, dataFormat&#x3D;&#39;null&#39;, collectionFormat&#x3D;&#39;null&#39;, description&#x3D;&#39;The authentication provider.&#39;, unescapedDescription&#x3D;&#39;The authentication provider.&#39;, baseType&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, enumName&#x3D;&#39;null&#39;, style&#x3D;&#39;simple&#39;, deepObject&#x3D;&#39;false&#39;, allowEmptyValue&#x3D;&#39;false&#39;, example&#x3D;&#39;provider_example&#39;, jsonSchema&#x3D;&#39;{
  &quot;name&quot; : &quot;provider&quot;,
  &quot;in&quot; : &quot;path&quot;,
  &quot;description&quot; : &quot;The authentication provider.&quot;,
  &quot;required&quot; : true,
  &quot;style&quot; : &quot;simple&quot;,
  &quot;explode&quot; : false,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;string&quot;
  }
}&#39;, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isAnyType&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isFile&#x3D;false, isEnum&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, mostInnerItems&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], vendorExtensions&#x3D;{x-go-example&#x3D;&quot;provider_example&quot;, x-export-param-name&#x3D;Provider}, hasValidation&#x3D;false, maxProperties&#x3D;null, minProperties&#x3D;null, isNullable&#x3D;false, isDeprecated&#x3D;false, required&#x3D;true, maximum&#x3D;&#39;null&#39;, exclusiveMaximum&#x3D;false, minimum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, uniqueItems&#x3D;false, contentType&#x3D;null, multipleOf&#x3D;null, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false, schema&#x3D;CodegenProperty{openApiType&#x3D;&#39;string&#39;, baseName&#x3D;&#39;provider&#39;, complexType&#x3D;&#39;null&#39;, getter&#x3D;&#39;GetProvider&#39;, setter&#x3D;&#39;setProvider&#39;, description&#x3D;&#39;null&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;string&#39;, dataFormat&#x3D;&#39;null&#39;, name&#x3D;&#39;Provider&#39;, min&#x3D;&#39;null&#39;, max&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, defaultValueWithParam&#x3D;&#39; &#x3D; data.provider;&#39;, baseType&#x3D;&#39;string&#39;, containerType&#x3D;&#39;null&#39;, title&#x3D;&#39;null&#39;, unescapedDescription&#x3D;&#39;null&#39;, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;type&quot; : &quot;string&quot;
}&#39;, minimum&#x3D;&#39;null&#39;, maximum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, exclusiveMaximum&#x3D;false, required&#x3D;true, deprecated&#x3D;false, hasMoreNonReadOnly&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isContainer&#x3D;false, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isFile&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isEnum&#x3D;false, isReadOnly&#x3D;false, isWriteOnly&#x3D;false, isNullable&#x3D;false, isSelfReference&#x3D;false, isCircularReference&#x3D;false, isDiscriminator&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], mostInnerItems&#x3D;null, vendorExtensions&#x3D;{}, hasValidation&#x3D;false, isInherited&#x3D;false, discriminatorValue&#x3D;&#39;null&#39;, nameInCamelCase&#x3D;&#39;provider&#39;, nameInSnakeCase&#x3D;&#39;PROVIDER&#39;, enumName&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, maxProperties&#x3D;null, minProperties&#x3D;null, uniqueItems&#x3D;false, multipleOf&#x3D;null, isXmlAttribute&#x3D;false, xmlPrefix&#x3D;&#39;null&#39;, xmlName&#x3D;&#39;null&#39;, xmlNamespace&#x3D;&#39;null&#39;, isXmlWrapped&#x3D;false, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false}, content&#x3D;null}]
**** | [****](.md) |  | [CodegenParameter{isFormParam&#x3D;false, isQueryParam&#x3D;false, isPathParam&#x3D;true, isHeaderParam&#x3D;false, isCookieParam&#x3D;false, isBodyParam&#x3D;false, isContainer&#x3D;false, isCollectionFormatMulti&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isExplode&#x3D;false, baseName&#x3D;&#39;provider&#39;, paramName&#x3D;&#39;provider&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;null&#39;, dataFormat&#x3D;&#39;null&#39;, collectionFormat&#x3D;&#39;null&#39;, description&#x3D;&#39;The authentication provider.&#39;, unescapedDescription&#x3D;&#39;The authentication provider.&#39;, baseType&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, enumName&#x3D;&#39;null&#39;, style&#x3D;&#39;simple&#39;, deepObject&#x3D;&#39;false&#39;, allowEmptyValue&#x3D;&#39;false&#39;, example&#x3D;&#39;provider_example&#39;, jsonSchema&#x3D;&#39;{
  &quot;name&quot; : &quot;provider&quot;,
  &quot;in&quot; : &quot;path&quot;,
  &quot;description&quot; : &quot;The authentication provider.&quot;,
  &quot;required&quot; : true,
  &quot;style&quot; : &quot;simple&quot;,
  &quot;explode&quot; : false,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;string&quot;
  }
}&#39;, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isAnyType&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isFile&#x3D;false, isEnum&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, mostInnerItems&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], vendorExtensions&#x3D;{x-go-example&#x3D;&quot;provider_example&quot;, x-export-param-name&#x3D;Provider}, hasValidation&#x3D;false, maxProperties&#x3D;null, minProperties&#x3D;null, isNullable&#x3D;false, isDeprecated&#x3D;false, required&#x3D;true, maximum&#x3D;&#39;null&#39;, exclusiveMaximum&#x3D;false, minimum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, uniqueItems&#x3D;false, contentType&#x3D;null, multipleOf&#x3D;null, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false, schema&#x3D;CodegenProperty{openApiType&#x3D;&#39;string&#39;, baseName&#x3D;&#39;provider&#39;, complexType&#x3D;&#39;null&#39;, getter&#x3D;&#39;GetProvider&#39;, setter&#x3D;&#39;setProvider&#39;, description&#x3D;&#39;null&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;string&#39;, dataFormat&#x3D;&#39;null&#39;, name&#x3D;&#39;Provider&#39;, min&#x3D;&#39;null&#39;, max&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, defaultValueWithParam&#x3D;&#39; &#x3D; data.provider;&#39;, baseType&#x3D;&#39;string&#39;, containerType&#x3D;&#39;null&#39;, title&#x3D;&#39;null&#39;, unescapedDescription&#x3D;&#39;null&#39;, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;type&quot; : &quot;string&quot;
}&#39;, minimum&#x3D;&#39;null&#39;, maximum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, exclusiveMaximum&#x3D;false, required&#x3D;true, deprecated&#x3D;false, hasMoreNonReadOnly&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isContainer&#x3D;false, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isFile&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isEnum&#x3D;false, isReadOnly&#x3D;false, isWriteOnly&#x3D;false, isNullable&#x3D;false, isSelfReference&#x3D;false, isCircularReference&#x3D;false, isDiscriminator&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], mostInnerItems&#x3D;null, vendorExtensions&#x3D;{}, hasValidation&#x3D;false, isInherited&#x3D;false, discriminatorValue&#x3D;&#39;null&#39;, nameInCamelCase&#x3D;&#39;provider&#39;, nameInSnakeCase&#x3D;&#39;PROVIDER&#39;, enumName&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, maxProperties&#x3D;null, minProperties&#x3D;null, uniqueItems&#x3D;false, multipleOf&#x3D;null, isXmlAttribute&#x3D;false, xmlPrefix&#x3D;&#39;null&#39;, xmlName&#x3D;&#39;null&#39;, xmlNamespace&#x3D;&#39;null&#39;, isXmlWrapped&#x3D;false, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false}, content&#x3D;null}]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderCallback

> ProviderCallback(ctx, provider).Execute()

Auth provider callback



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
    provider := "provider_example" // string | The authentication provider.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Auth.ProviderCallback(context.Background(), provider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.ProviderCallback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | The authentication provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------[CodegenParameter{isFormParam&#x3D;false, isQueryParam&#x3D;false, isPathParam&#x3D;true, isHeaderParam&#x3D;false, isCookieParam&#x3D;false, isBodyParam&#x3D;false, isContainer&#x3D;false, isCollectionFormatMulti&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isExplode&#x3D;false, baseName&#x3D;&#39;provider&#39;, paramName&#x3D;&#39;provider&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;null&#39;, dataFormat&#x3D;&#39;null&#39;, collectionFormat&#x3D;&#39;null&#39;, description&#x3D;&#39;The authentication provider.&#39;, unescapedDescription&#x3D;&#39;The authentication provider.&#39;, baseType&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, enumName&#x3D;&#39;null&#39;, style&#x3D;&#39;simple&#39;, deepObject&#x3D;&#39;false&#39;, allowEmptyValue&#x3D;&#39;false&#39;, example&#x3D;&#39;provider_example&#39;, jsonSchema&#x3D;&#39;{
  &quot;name&quot; : &quot;provider&quot;,
  &quot;in&quot; : &quot;path&quot;,
  &quot;description&quot; : &quot;The authentication provider.&quot;,
  &quot;required&quot; : true,
  &quot;style&quot; : &quot;simple&quot;,
  &quot;explode&quot; : false,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;string&quot;
  }
}&#39;, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isAnyType&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isFile&#x3D;false, isEnum&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, mostInnerItems&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], vendorExtensions&#x3D;{x-go-example&#x3D;&quot;provider_example&quot;, x-export-param-name&#x3D;Provider}, hasValidation&#x3D;false, maxProperties&#x3D;null, minProperties&#x3D;null, isNullable&#x3D;false, isDeprecated&#x3D;false, required&#x3D;true, maximum&#x3D;&#39;null&#39;, exclusiveMaximum&#x3D;false, minimum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, uniqueItems&#x3D;false, contentType&#x3D;null, multipleOf&#x3D;null, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false, schema&#x3D;CodegenProperty{openApiType&#x3D;&#39;string&#39;, baseName&#x3D;&#39;provider&#39;, complexType&#x3D;&#39;null&#39;, getter&#x3D;&#39;GetProvider&#39;, setter&#x3D;&#39;setProvider&#39;, description&#x3D;&#39;null&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;string&#39;, dataFormat&#x3D;&#39;null&#39;, name&#x3D;&#39;Provider&#39;, min&#x3D;&#39;null&#39;, max&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, defaultValueWithParam&#x3D;&#39; &#x3D; data.provider;&#39;, baseType&#x3D;&#39;string&#39;, containerType&#x3D;&#39;null&#39;, title&#x3D;&#39;null&#39;, unescapedDescription&#x3D;&#39;null&#39;, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;type&quot; : &quot;string&quot;
}&#39;, minimum&#x3D;&#39;null&#39;, maximum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, exclusiveMaximum&#x3D;false, required&#x3D;true, deprecated&#x3D;false, hasMoreNonReadOnly&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isContainer&#x3D;false, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isFile&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isEnum&#x3D;false, isReadOnly&#x3D;false, isWriteOnly&#x3D;false, isNullable&#x3D;false, isSelfReference&#x3D;false, isCircularReference&#x3D;false, isDiscriminator&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], mostInnerItems&#x3D;null, vendorExtensions&#x3D;{}, hasValidation&#x3D;false, isInherited&#x3D;false, discriminatorValue&#x3D;&#39;null&#39;, nameInCamelCase&#x3D;&#39;provider&#39;, nameInSnakeCase&#x3D;&#39;PROVIDER&#39;, enumName&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, maxProperties&#x3D;null, minProperties&#x3D;null, uniqueItems&#x3D;false, multipleOf&#x3D;null, isXmlAttribute&#x3D;false, xmlPrefix&#x3D;&#39;null&#39;, xmlName&#x3D;&#39;null&#39;, xmlNamespace&#x3D;&#39;null&#39;, isXmlWrapped&#x3D;false, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false}, content&#x3D;null}]
**** | [****](.md) |  | [CodegenParameter{isFormParam&#x3D;false, isQueryParam&#x3D;false, isPathParam&#x3D;true, isHeaderParam&#x3D;false, isCookieParam&#x3D;false, isBodyParam&#x3D;false, isContainer&#x3D;false, isCollectionFormatMulti&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isExplode&#x3D;false, baseName&#x3D;&#39;provider&#39;, paramName&#x3D;&#39;provider&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;null&#39;, dataFormat&#x3D;&#39;null&#39;, collectionFormat&#x3D;&#39;null&#39;, description&#x3D;&#39;The authentication provider.&#39;, unescapedDescription&#x3D;&#39;The authentication provider.&#39;, baseType&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, enumName&#x3D;&#39;null&#39;, style&#x3D;&#39;simple&#39;, deepObject&#x3D;&#39;false&#39;, allowEmptyValue&#x3D;&#39;false&#39;, example&#x3D;&#39;provider_example&#39;, jsonSchema&#x3D;&#39;{
  &quot;name&quot; : &quot;provider&quot;,
  &quot;in&quot; : &quot;path&quot;,
  &quot;description&quot; : &quot;The authentication provider.&quot;,
  &quot;required&quot; : true,
  &quot;style&quot; : &quot;simple&quot;,
  &quot;explode&quot; : false,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;string&quot;
  }
}&#39;, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isAnyType&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isFile&#x3D;false, isEnum&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, mostInnerItems&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], vendorExtensions&#x3D;{x-go-example&#x3D;&quot;provider_example&quot;, x-export-param-name&#x3D;Provider}, hasValidation&#x3D;false, maxProperties&#x3D;null, minProperties&#x3D;null, isNullable&#x3D;false, isDeprecated&#x3D;false, required&#x3D;true, maximum&#x3D;&#39;null&#39;, exclusiveMaximum&#x3D;false, minimum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, uniqueItems&#x3D;false, contentType&#x3D;null, multipleOf&#x3D;null, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false, schema&#x3D;CodegenProperty{openApiType&#x3D;&#39;string&#39;, baseName&#x3D;&#39;provider&#39;, complexType&#x3D;&#39;null&#39;, getter&#x3D;&#39;GetProvider&#39;, setter&#x3D;&#39;setProvider&#39;, description&#x3D;&#39;null&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;string&#39;, dataFormat&#x3D;&#39;null&#39;, name&#x3D;&#39;Provider&#39;, min&#x3D;&#39;null&#39;, max&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, defaultValueWithParam&#x3D;&#39; &#x3D; data.provider;&#39;, baseType&#x3D;&#39;string&#39;, containerType&#x3D;&#39;null&#39;, title&#x3D;&#39;null&#39;, unescapedDescription&#x3D;&#39;null&#39;, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;type&quot; : &quot;string&quot;
}&#39;, minimum&#x3D;&#39;null&#39;, maximum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, exclusiveMaximum&#x3D;false, required&#x3D;true, deprecated&#x3D;false, hasMoreNonReadOnly&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isContainer&#x3D;false, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isFile&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isEnum&#x3D;false, isReadOnly&#x3D;false, isWriteOnly&#x3D;false, isNullable&#x3D;false, isSelfReference&#x3D;false, isCircularReference&#x3D;false, isDiscriminator&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], mostInnerItems&#x3D;null, vendorExtensions&#x3D;{}, hasValidation&#x3D;false, isInherited&#x3D;false, discriminatorValue&#x3D;&#39;null&#39;, nameInCamelCase&#x3D;&#39;provider&#39;, nameInSnakeCase&#x3D;&#39;PROVIDER&#39;, enumName&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, maxProperties&#x3D;null, minProperties&#x3D;null, uniqueItems&#x3D;false, multipleOf&#x3D;null, isXmlAttribute&#x3D;false, xmlPrefix&#x3D;&#39;null&#39;, xmlName&#x3D;&#39;null&#39;, xmlNamespace&#x3D;&#39;null&#39;, isXmlWrapped&#x3D;false, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false}, content&#x3D;null}]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Signup

> Signup(ctx).Request(request).Execute()

User signup



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
    request := *openapiclient.NewUserSignupRequest("Email_example") // UserSignupRequest | The request body to signup.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Auth.Signup(context.Background()).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.Signup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------[CodegenParameter{isFormParam&#x3D;false, isQueryParam&#x3D;false, isPathParam&#x3D;false, isHeaderParam&#x3D;false, isCookieParam&#x3D;false, isBodyParam&#x3D;true, isContainer&#x3D;false, isCollectionFormatMulti&#x3D;false, isPrimitiveType&#x3D;false, isModel&#x3D;true, isExplode&#x3D;false, baseName&#x3D;&#39;request&#39;, paramName&#x3D;&#39;request&#39;, dataType&#x3D;&#39;UserSignupRequest&#39;, datatypeWithEnum&#x3D;&#39;null&#39;, dataFormat&#x3D;&#39;null&#39;, collectionFormat&#x3D;&#39;null&#39;, description&#x3D;&#39;The request body to signup.&#39;, unescapedDescription&#x3D;&#39;null&#39;, baseType&#x3D;&#39;UserSignupRequest&#39;, defaultValue&#x3D;&#39;null&#39;, enumName&#x3D;&#39;null&#39;, style&#x3D;&#39;null&#39;, deepObject&#x3D;&#39;false&#39;, allowEmptyValue&#x3D;&#39;false&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;description&quot; : &quot;The request body to signup.&quot;,
  &quot;content&quot; : {
    &quot;application/json&quot; : {
      &quot;schema&quot; : {
        &quot;$ref&quot; : &quot;#/components/schemas/UserSignupRequest&quot;
      }
    }
  },
  &quot;required&quot; : true
}&#39;, isString&#x3D;false, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isAnyType&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isFile&#x3D;false, isEnum&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, mostInnerItems&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[CodegenProperty{openApiType&#x3D;&#39;string&#39;, baseName&#x3D;&#39;email&#39;, complexType&#x3D;&#39;null&#39;, getter&#x3D;&#39;GetEmail&#39;, setter&#x3D;&#39;setEmail&#39;, description&#x3D;&#39;null&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;string&#39;, dataFormat&#x3D;&#39;null&#39;, name&#x3D;&#39;Email&#39;, min&#x3D;&#39;null&#39;, max&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, defaultValueWithParam&#x3D;&#39; &#x3D; data.email;&#39;, baseType&#x3D;&#39;string&#39;, containerType&#x3D;&#39;null&#39;, title&#x3D;&#39;null&#39;, unescapedDescription&#x3D;&#39;null&#39;, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;type&quot; : &quot;string&quot;
}&#39;, minimum&#x3D;&#39;null&#39;, maximum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, exclusiveMaximum&#x3D;false, required&#x3D;true, deprecated&#x3D;false, hasMoreNonReadOnly&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isContainer&#x3D;false, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isFile&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isEnum&#x3D;false, isReadOnly&#x3D;false, isWriteOnly&#x3D;false, isNullable&#x3D;false, isSelfReference&#x3D;false, isCircularReference&#x3D;false, isDiscriminator&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], mostInnerItems&#x3D;null, vendorExtensions&#x3D;{}, hasValidation&#x3D;false, isInherited&#x3D;false, discriminatorValue&#x3D;&#39;null&#39;, nameInCamelCase&#x3D;&#39;email&#39;, nameInSnakeCase&#x3D;&#39;EMAIL&#39;, enumName&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, maxProperties&#x3D;null, minProperties&#x3D;null, uniqueItems&#x3D;false, multipleOf&#x3D;null, isXmlAttribute&#x3D;false, xmlPrefix&#x3D;&#39;null&#39;, xmlName&#x3D;&#39;null&#39;, xmlNamespace&#x3D;&#39;null&#39;, isXmlWrapped&#x3D;false, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false}], requiredVars&#x3D;[CodegenProperty{openApiType&#x3D;&#39;string&#39;, baseName&#x3D;&#39;email&#39;, complexType&#x3D;&#39;null&#39;, getter&#x3D;&#39;GetEmail&#39;, setter&#x3D;&#39;setEmail&#39;, description&#x3D;&#39;null&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;string&#39;, dataFormat&#x3D;&#39;null&#39;, name&#x3D;&#39;Email&#39;, min&#x3D;&#39;null&#39;, max&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, defaultValueWithParam&#x3D;&#39; &#x3D; data.email;&#39;, baseType&#x3D;&#39;string&#39;, containerType&#x3D;&#39;null&#39;, title&#x3D;&#39;null&#39;, unescapedDescription&#x3D;&#39;null&#39;, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;type&quot; : &quot;string&quot;
}&#39;, minimum&#x3D;&#39;null&#39;, maximum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, exclusiveMaximum&#x3D;false, required&#x3D;true, deprecated&#x3D;false, hasMoreNonReadOnly&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isContainer&#x3D;false, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isFile&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isEnum&#x3D;false, isReadOnly&#x3D;false, isWriteOnly&#x3D;false, isNullable&#x3D;false, isSelfReference&#x3D;false, isCircularReference&#x3D;false, isDiscriminator&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], mostInnerItems&#x3D;null, vendorExtensions&#x3D;{}, hasValidation&#x3D;false, isInherited&#x3D;false, discriminatorValue&#x3D;&#39;null&#39;, nameInCamelCase&#x3D;&#39;email&#39;, nameInSnakeCase&#x3D;&#39;EMAIL&#39;, enumName&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, maxProperties&#x3D;null, minProperties&#x3D;null, uniqueItems&#x3D;false, multipleOf&#x3D;null, isXmlAttribute&#x3D;false, xmlPrefix&#x3D;&#39;null&#39;, xmlName&#x3D;&#39;null&#39;, xmlNamespace&#x3D;&#39;null&#39;, isXmlWrapped&#x3D;false, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false}], vendorExtensions&#x3D;{x-go-example&#x3D;*openapiclient.NewUserSignupRequest(&quot;Email_example&quot;), x-export-param-name&#x3D;Request}, hasValidation&#x3D;false, maxProperties&#x3D;null, minProperties&#x3D;null, isNullable&#x3D;false, isDeprecated&#x3D;false, required&#x3D;true, maximum&#x3D;&#39;null&#39;, exclusiveMaximum&#x3D;false, minimum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, uniqueItems&#x3D;false, contentType&#x3D;null, multipleOf&#x3D;null, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;true, getHasRequired&#x3D;true, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false, schema&#x3D;null, content&#x3D;{application/json&#x3D;CodegenMediaType{schema&#x3D;CodegenProperty{openApiType&#x3D;&#39;UserSignupRequest&#39;, baseName&#x3D;&#39;SchemaForRequestBodyApplicationJson&#39;, complexType&#x3D;&#39;UserSignupRequest&#39;, getter&#x3D;&#39;GetSchemaForRequestBodyApplicationJson&#39;, setter&#x3D;&#39;setSchemaForRequestBodyApplicationJson&#39;, description&#x3D;&#39;null&#39;, dataType&#x3D;&#39;UserSignupRequest&#39;, datatypeWithEnum&#x3D;&#39;UserSignupRequest&#39;, dataFormat&#x3D;&#39;null&#39;, name&#x3D;&#39;SchemaForRequestBodyApplicationJson&#39;, min&#x3D;&#39;null&#39;, max&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, defaultValueWithParam&#x3D;&#39; &#x3D; data.SchemaForRequestBodyApplicationJson;&#39;, baseType&#x3D;&#39;UserSignupRequest&#39;, containerType&#x3D;&#39;null&#39;, title&#x3D;&#39;null&#39;, unescapedDescription&#x3D;&#39;null&#39;, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;$ref&quot; : &quot;#/components/schemas/UserSignupRequest&quot;
}&#39;, minimum&#x3D;&#39;null&#39;, maximum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, exclusiveMaximum&#x3D;false, required&#x3D;false, deprecated&#x3D;false, hasMoreNonReadOnly&#x3D;false, isPrimitiveType&#x3D;false, isModel&#x3D;true, isContainer&#x3D;false, isString&#x3D;false, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isFile&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isEnum&#x3D;false, isReadOnly&#x3D;false, isWriteOnly&#x3D;false, isNullable&#x3D;false, isSelfReference&#x3D;false, isCircularReference&#x3D;false, isDiscriminator&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], mostInnerItems&#x3D;null, vendorExtensions&#x3D;{}, hasValidation&#x3D;false, isInherited&#x3D;false, discriminatorValue&#x3D;&#39;null&#39;, nameInCamelCase&#x3D;&#39;schemaForRequestBodyApplicationJson&#39;, nameInSnakeCase&#x3D;&#39;SCHEMA_FOR_REQUEST_BODY_APPLICATION_JSON&#39;, enumName&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, maxProperties&#x3D;null, minProperties&#x3D;null, uniqueItems&#x3D;false, multipleOf&#x3D;null, isXmlAttribute&#x3D;false, xmlPrefix&#x3D;&#39;null&#39;, xmlName&#x3D;&#39;null&#39;, xmlNamespace&#x3D;&#39;null&#39;, isXmlWrapped&#x3D;false, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false}, encoding&#x3D;null}}}]
**** | [****](.md) |  | [CodegenParameter{isFormParam&#x3D;false, isQueryParam&#x3D;false, isPathParam&#x3D;false, isHeaderParam&#x3D;false, isCookieParam&#x3D;false, isBodyParam&#x3D;true, isContainer&#x3D;false, isCollectionFormatMulti&#x3D;false, isPrimitiveType&#x3D;false, isModel&#x3D;true, isExplode&#x3D;false, baseName&#x3D;&#39;request&#39;, paramName&#x3D;&#39;request&#39;, dataType&#x3D;&#39;UserSignupRequest&#39;, datatypeWithEnum&#x3D;&#39;null&#39;, dataFormat&#x3D;&#39;null&#39;, collectionFormat&#x3D;&#39;null&#39;, description&#x3D;&#39;The request body to signup.&#39;, unescapedDescription&#x3D;&#39;null&#39;, baseType&#x3D;&#39;UserSignupRequest&#39;, defaultValue&#x3D;&#39;null&#39;, enumName&#x3D;&#39;null&#39;, style&#x3D;&#39;null&#39;, deepObject&#x3D;&#39;false&#39;, allowEmptyValue&#x3D;&#39;false&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;description&quot; : &quot;The request body to signup.&quot;,
  &quot;content&quot; : {
    &quot;application/json&quot; : {
      &quot;schema&quot; : {
        &quot;$ref&quot; : &quot;#/components/schemas/UserSignupRequest&quot;
      }
    }
  },
  &quot;required&quot; : true
}&#39;, isString&#x3D;false, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isAnyType&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isFile&#x3D;false, isEnum&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, mostInnerItems&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[CodegenProperty{openApiType&#x3D;&#39;string&#39;, baseName&#x3D;&#39;email&#39;, complexType&#x3D;&#39;null&#39;, getter&#x3D;&#39;GetEmail&#39;, setter&#x3D;&#39;setEmail&#39;, description&#x3D;&#39;null&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;string&#39;, dataFormat&#x3D;&#39;null&#39;, name&#x3D;&#39;Email&#39;, min&#x3D;&#39;null&#39;, max&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, defaultValueWithParam&#x3D;&#39; &#x3D; data.email;&#39;, baseType&#x3D;&#39;string&#39;, containerType&#x3D;&#39;null&#39;, title&#x3D;&#39;null&#39;, unescapedDescription&#x3D;&#39;null&#39;, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;type&quot; : &quot;string&quot;
}&#39;, minimum&#x3D;&#39;null&#39;, maximum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, exclusiveMaximum&#x3D;false, required&#x3D;true, deprecated&#x3D;false, hasMoreNonReadOnly&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isContainer&#x3D;false, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isFile&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isEnum&#x3D;false, isReadOnly&#x3D;false, isWriteOnly&#x3D;false, isNullable&#x3D;false, isSelfReference&#x3D;false, isCircularReference&#x3D;false, isDiscriminator&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], mostInnerItems&#x3D;null, vendorExtensions&#x3D;{}, hasValidation&#x3D;false, isInherited&#x3D;false, discriminatorValue&#x3D;&#39;null&#39;, nameInCamelCase&#x3D;&#39;email&#39;, nameInSnakeCase&#x3D;&#39;EMAIL&#39;, enumName&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, maxProperties&#x3D;null, minProperties&#x3D;null, uniqueItems&#x3D;false, multipleOf&#x3D;null, isXmlAttribute&#x3D;false, xmlPrefix&#x3D;&#39;null&#39;, xmlName&#x3D;&#39;null&#39;, xmlNamespace&#x3D;&#39;null&#39;, isXmlWrapped&#x3D;false, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false}], requiredVars&#x3D;[CodegenProperty{openApiType&#x3D;&#39;string&#39;, baseName&#x3D;&#39;email&#39;, complexType&#x3D;&#39;null&#39;, getter&#x3D;&#39;GetEmail&#39;, setter&#x3D;&#39;setEmail&#39;, description&#x3D;&#39;null&#39;, dataType&#x3D;&#39;string&#39;, datatypeWithEnum&#x3D;&#39;string&#39;, dataFormat&#x3D;&#39;null&#39;, name&#x3D;&#39;Email&#39;, min&#x3D;&#39;null&#39;, max&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, defaultValueWithParam&#x3D;&#39; &#x3D; data.email;&#39;, baseType&#x3D;&#39;string&#39;, containerType&#x3D;&#39;null&#39;, title&#x3D;&#39;null&#39;, unescapedDescription&#x3D;&#39;null&#39;, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;type&quot; : &quot;string&quot;
}&#39;, minimum&#x3D;&#39;null&#39;, maximum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, exclusiveMaximum&#x3D;false, required&#x3D;true, deprecated&#x3D;false, hasMoreNonReadOnly&#x3D;false, isPrimitiveType&#x3D;true, isModel&#x3D;false, isContainer&#x3D;false, isString&#x3D;true, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isFile&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isEnum&#x3D;false, isReadOnly&#x3D;false, isWriteOnly&#x3D;false, isNullable&#x3D;false, isSelfReference&#x3D;false, isCircularReference&#x3D;false, isDiscriminator&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], mostInnerItems&#x3D;null, vendorExtensions&#x3D;{}, hasValidation&#x3D;false, isInherited&#x3D;false, discriminatorValue&#x3D;&#39;null&#39;, nameInCamelCase&#x3D;&#39;email&#39;, nameInSnakeCase&#x3D;&#39;EMAIL&#39;, enumName&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, maxProperties&#x3D;null, minProperties&#x3D;null, uniqueItems&#x3D;false, multipleOf&#x3D;null, isXmlAttribute&#x3D;false, xmlPrefix&#x3D;&#39;null&#39;, xmlName&#x3D;&#39;null&#39;, xmlNamespace&#x3D;&#39;null&#39;, isXmlWrapped&#x3D;false, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false}], vendorExtensions&#x3D;{x-go-example&#x3D;*openapiclient.NewUserSignupRequest(&quot;Email_example&quot;), x-export-param-name&#x3D;Request}, hasValidation&#x3D;false, maxProperties&#x3D;null, minProperties&#x3D;null, isNullable&#x3D;false, isDeprecated&#x3D;false, required&#x3D;true, maximum&#x3D;&#39;null&#39;, exclusiveMaximum&#x3D;false, minimum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, uniqueItems&#x3D;false, contentType&#x3D;null, multipleOf&#x3D;null, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;true, getHasRequired&#x3D;true, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false, schema&#x3D;null, content&#x3D;{application/json&#x3D;CodegenMediaType{schema&#x3D;CodegenProperty{openApiType&#x3D;&#39;UserSignupRequest&#39;, baseName&#x3D;&#39;SchemaForRequestBodyApplicationJson&#39;, complexType&#x3D;&#39;UserSignupRequest&#39;, getter&#x3D;&#39;GetSchemaForRequestBodyApplicationJson&#39;, setter&#x3D;&#39;setSchemaForRequestBodyApplicationJson&#39;, description&#x3D;&#39;null&#39;, dataType&#x3D;&#39;UserSignupRequest&#39;, datatypeWithEnum&#x3D;&#39;UserSignupRequest&#39;, dataFormat&#x3D;&#39;null&#39;, name&#x3D;&#39;SchemaForRequestBodyApplicationJson&#39;, min&#x3D;&#39;null&#39;, max&#x3D;&#39;null&#39;, defaultValue&#x3D;&#39;null&#39;, defaultValueWithParam&#x3D;&#39; &#x3D; data.SchemaForRequestBodyApplicationJson;&#39;, baseType&#x3D;&#39;UserSignupRequest&#39;, containerType&#x3D;&#39;null&#39;, title&#x3D;&#39;null&#39;, unescapedDescription&#x3D;&#39;null&#39;, maxLength&#x3D;null, minLength&#x3D;null, pattern&#x3D;&#39;null&#39;, example&#x3D;&#39;null&#39;, jsonSchema&#x3D;&#39;{
  &quot;$ref&quot; : &quot;#/components/schemas/UserSignupRequest&quot;
}&#39;, minimum&#x3D;&#39;null&#39;, maximum&#x3D;&#39;null&#39;, exclusiveMinimum&#x3D;false, exclusiveMaximum&#x3D;false, required&#x3D;false, deprecated&#x3D;false, hasMoreNonReadOnly&#x3D;false, isPrimitiveType&#x3D;false, isModel&#x3D;true, isContainer&#x3D;false, isString&#x3D;false, isNumeric&#x3D;false, isInteger&#x3D;false, isShort&#x3D;false, isLong&#x3D;false, isUnboundedInteger&#x3D;false, isNumber&#x3D;false, isFloat&#x3D;false, isDouble&#x3D;false, isDecimal&#x3D;false, isByteArray&#x3D;false, isBinary&#x3D;false, isFile&#x3D;false, isBoolean&#x3D;false, isDate&#x3D;false, isDateTime&#x3D;false, isUuid&#x3D;false, isUri&#x3D;false, isEmail&#x3D;false, isFreeFormObject&#x3D;false, isArray&#x3D;false, isMap&#x3D;false, isEnum&#x3D;false, isReadOnly&#x3D;false, isWriteOnly&#x3D;false, isNullable&#x3D;false, isSelfReference&#x3D;false, isCircularReference&#x3D;false, isDiscriminator&#x3D;false, _enum&#x3D;null, allowableValues&#x3D;null, items&#x3D;null, additionalProperties&#x3D;null, vars&#x3D;[], requiredVars&#x3D;[], mostInnerItems&#x3D;null, vendorExtensions&#x3D;{}, hasValidation&#x3D;false, isInherited&#x3D;false, discriminatorValue&#x3D;&#39;null&#39;, nameInCamelCase&#x3D;&#39;schemaForRequestBodyApplicationJson&#39;, nameInSnakeCase&#x3D;&#39;SCHEMA_FOR_REQUEST_BODY_APPLICATION_JSON&#39;, enumName&#x3D;&#39;null&#39;, maxItems&#x3D;null, minItems&#x3D;null, maxProperties&#x3D;null, minProperties&#x3D;null, uniqueItems&#x3D;false, multipleOf&#x3D;null, isXmlAttribute&#x3D;false, xmlPrefix&#x3D;&#39;null&#39;, xmlName&#x3D;&#39;null&#39;, xmlNamespace&#x3D;&#39;null&#39;, isXmlWrapped&#x3D;false, isNull&#x3D;false, getAdditionalPropertiesIsAnyType&#x3D;false, getHasVars&#x3D;false, getHasRequired&#x3D;false, getHasDiscriminatorWithNonEmptyMapping&#x3D;false, composedSchemas&#x3D;null, hasMultipleTypes&#x3D;false}, encoding&#x3D;null}}}]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

