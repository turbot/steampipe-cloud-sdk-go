/*
Steampipe Cloud

Steampipe Cloud is a hosted version of Steampipe (https://steampipe.io), an open source tool to instantly query your cloud services (e.g. AWS, Azure, GCP and more) with SQL. No DB required.

API version: {{OPEN_API_VERSION}}
Contact: help@steampipe.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package steampipecloud

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// UserWorkspaceConnectionAssociationsService UserWorkspaceConnectionAssociations service
type UserWorkspaceConnectionAssociationsService service

type UserWorkspaceConnectionAssociationsApiCreateRequest struct {
	ctx             _context.Context
	ApiService      *UserWorkspaceConnectionAssociationsService
	userHandle      string
	workspaceHandle string
	request         *CreateWorkspaceConnRequest
}

// The request body for the association to be created.
func (r UserWorkspaceConnectionAssociationsApiCreateRequest) Request(request CreateWorkspaceConnRequest) UserWorkspaceConnectionAssociationsApiCreateRequest {
	r.request = &request
	return r
}

func (r UserWorkspaceConnectionAssociationsApiCreateRequest) Execute() (WorkspaceConn, *_nethttp.Response, error) {
	return r.ApiService.CreateExecute(r)
}

/*
Create Create user workspace connection association

Associate a connection with the workspace. A workspace can have multiple association.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userHandle The handle of the user where we want to create an association.
	@param workspaceHandle The handle of the workspace where the connection will be associated.
	@return UserWorkspaceConnectionAssociationsApiCreateRequest
*/
func (a *UserWorkspaceConnectionAssociationsService) Create(ctx _context.Context, userHandle string, workspaceHandle string) UserWorkspaceConnectionAssociationsApiCreateRequest {
	return UserWorkspaceConnectionAssociationsApiCreateRequest{
		ApiService:      a,
		ctx:             ctx,
		userHandle:      userHandle,
		workspaceHandle: workspaceHandle,
	}
}

// Execute executes the request
//
//	@return WorkspaceConn
func (a *UserWorkspaceConnectionAssociationsService) CreateExecute(r UserWorkspaceConnectionAssociationsApiCreateRequest) (WorkspaceConn, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue WorkspaceConn
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserWorkspaceConnectionAssociationsService.Create")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user/{user_handle}/workspace/{workspace_handle}/conn"
	localVarPath = strings.Replace(localVarPath, "{"+"user_handle"+"}", _neturl.PathEscape(parameterToString(r.userHandle, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workspace_handle"+"}", _neturl.PathEscape(parameterToString(r.workspaceHandle, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.request == nil {
		return localVarReturnValue, nil, reportError("request is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.request
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 402 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UserWorkspaceConnectionAssociationsApiDeleteRequest struct {
	ctx             _context.Context
	ApiService      *UserWorkspaceConnectionAssociationsService
	userHandle      string
	workspaceHandle string
	connHandle      string
}

func (r UserWorkspaceConnectionAssociationsApiDeleteRequest) Execute() (WorkspaceConn, *_nethttp.Response, error) {
	return r.ApiService.DeleteExecute(r)
}

/*
Delete Delete user workspace connection association

Dissociates or delete the workspace and connection association.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userHandle The handle of the user performing the action.
	@param workspaceHandle The handle of the workspace whose association needs to be deleted.
	@param connHandle The handle of the conn whose association needs to be deleted.
	@return UserWorkspaceConnectionAssociationsApiDeleteRequest
*/
func (a *UserWorkspaceConnectionAssociationsService) Delete(ctx _context.Context, userHandle string, workspaceHandle string, connHandle string) UserWorkspaceConnectionAssociationsApiDeleteRequest {
	return UserWorkspaceConnectionAssociationsApiDeleteRequest{
		ApiService:      a,
		ctx:             ctx,
		userHandle:      userHandle,
		workspaceHandle: workspaceHandle,
		connHandle:      connHandle,
	}
}

// Execute executes the request
//
//	@return WorkspaceConn
func (a *UserWorkspaceConnectionAssociationsService) DeleteExecute(r UserWorkspaceConnectionAssociationsApiDeleteRequest) (WorkspaceConn, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue WorkspaceConn
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserWorkspaceConnectionAssociationsService.Delete")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user/{user_handle}/workspace/{workspace_handle}/conn/{conn_handle}"
	localVarPath = strings.Replace(localVarPath, "{"+"user_handle"+"}", _neturl.PathEscape(parameterToString(r.userHandle, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workspace_handle"+"}", _neturl.PathEscape(parameterToString(r.workspaceHandle, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"conn_handle"+"}", _neturl.PathEscape(parameterToString(r.connHandle, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UserWorkspaceConnectionAssociationsApiGetRequest struct {
	ctx             _context.Context
	ApiService      *UserWorkspaceConnectionAssociationsService
	userHandle      string
	workspaceHandle string
	connHandle      string
}

func (r UserWorkspaceConnectionAssociationsApiGetRequest) Execute() (WorkspaceConn, *_nethttp.Response, error) {
	return r.ApiService.GetExecute(r)
}

/*
Get Get user workspace connection association

Get the details for a workspace and connection association.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userHandle The handle of the user for which you want to get the association.
	@param workspaceHandle The handle of the workspace wherethe association exist.
	@param connHandle The handle of the conn whose association details needs to be fetched.
	@return UserWorkspaceConnectionAssociationsApiGetRequest
*/
func (a *UserWorkspaceConnectionAssociationsService) Get(ctx _context.Context, userHandle string, workspaceHandle string, connHandle string) UserWorkspaceConnectionAssociationsApiGetRequest {
	return UserWorkspaceConnectionAssociationsApiGetRequest{
		ApiService:      a,
		ctx:             ctx,
		userHandle:      userHandle,
		workspaceHandle: workspaceHandle,
		connHandle:      connHandle,
	}
}

// Execute executes the request
//
//	@return WorkspaceConn
func (a *UserWorkspaceConnectionAssociationsService) GetExecute(r UserWorkspaceConnectionAssociationsApiGetRequest) (WorkspaceConn, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue WorkspaceConn
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserWorkspaceConnectionAssociationsService.Get")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user/{user_handle}/workspace/{workspace_handle}/conn/{conn_handle}"
	localVarPath = strings.Replace(localVarPath, "{"+"user_handle"+"}", _neturl.PathEscape(parameterToString(r.userHandle, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workspace_handle"+"}", _neturl.PathEscape(parameterToString(r.workspaceHandle, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"conn_handle"+"}", _neturl.PathEscape(parameterToString(r.connHandle, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UserWorkspaceConnectionAssociationsApiListRequest struct {
	ctx             _context.Context
	ApiService      *UserWorkspaceConnectionAssociationsService
	userHandle      string
	workspaceHandle string
	limit           *int32
	nextToken       *string
}

// The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25.
func (r UserWorkspaceConnectionAssociationsApiListRequest) Limit(limit int32) UserWorkspaceConnectionAssociationsApiListRequest {
	r.limit = &limit
	return r
}

// When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data.
func (r UserWorkspaceConnectionAssociationsApiListRequest) NextToken(nextToken string) UserWorkspaceConnectionAssociationsApiListRequest {
	r.nextToken = &nextToken
	return r
}

func (r UserWorkspaceConnectionAssociationsApiListRequest) Execute() (ListWorkspaceConnResponse, *_nethttp.Response, error) {
	return r.ApiService.ListExecute(r)
}

/*
List List user workspace connection associations

List the connections associated with a workspace.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userHandle The handle of the user for which you want to start listing associations.
	@param workspaceHandle The handle of the workspace where you want to list the associations.
	@return UserWorkspaceConnectionAssociationsApiListRequest
*/
func (a *UserWorkspaceConnectionAssociationsService) List(ctx _context.Context, userHandle string, workspaceHandle string) UserWorkspaceConnectionAssociationsApiListRequest {
	return UserWorkspaceConnectionAssociationsApiListRequest{
		ApiService:      a,
		ctx:             ctx,
		userHandle:      userHandle,
		workspaceHandle: workspaceHandle,
	}
}

// Execute executes the request
//
//	@return ListWorkspaceConnResponse
func (a *UserWorkspaceConnectionAssociationsService) ListExecute(r UserWorkspaceConnectionAssociationsApiListRequest) (ListWorkspaceConnResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue ListWorkspaceConnResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserWorkspaceConnectionAssociationsService.List")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user/{user_handle}/workspace/{workspace_handle}/conn"
	localVarPath = strings.Replace(localVarPath, "{"+"user_handle"+"}", _neturl.PathEscape(parameterToString(r.userHandle, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workspace_handle"+"}", _neturl.PathEscape(parameterToString(r.workspaceHandle, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.nextToken != nil {
		localVarQueryParams.Add("next_token", parameterToString(*r.nextToken, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
