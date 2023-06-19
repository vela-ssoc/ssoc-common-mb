package elastic

import (
	"encoding/json"
	"net/http"
)

// BulkResponse is a response to a bulk execution.
//
// Example:
//
//	{
//	  "took":3,
//	  "errors":false,
//	  "items":[{
//	    "index":{
//	      "_index":"index1",
//	      "_type":"tweet",
//	      "_id":"1",
//	      "_version":3,
//	      "status":201
//	    }
//	  },{
//	    "index":{
//	      "_index":"index2",
//	      "_type":"tweet",
//	      "_id":"2",
//	      "_version":3,
//	      "status":200
//	    }
//	  },{
//	    "delete":{
//	      "_index":"index1",
//	      "_type":"tweet",
//	      "_id":"1",
//	      "_version":4,
//	      "status":200,
//	      "found":true
//	    }
//	  },{
//	    "update":{
//	      "_index":"index2",
//	      "_type":"tweet",
//	      "_id":"2",
//	      "_version":4,
//	      "status":200
//	    }
//	  }]
//	}
type BulkResponse struct {
	Took   int                            `json:"took,omitempty"`
	Errors bool                           `json:"errors,omitempty"`
	Items  []map[string]*BulkResponseItem `json:"items,omitempty"`
}

// BulkResponseItem is the result of a single bulk request.
type BulkResponseItem struct {
	Index         string        `json:"_index,omitempty"`
	Type          string        `json:"_type,omitempty"`
	ID            string        `json:"_id,omitempty"`
	Version       int64         `json:"_version,omitempty"`
	Result        string        `json:"result,omitempty"`
	Shards        *ShardsInfo   `json:"_shards,omitempty"`
	SeqNo         int64         `json:"_seq_no,omitempty"`
	PrimaryTerm   int64         `json:"_primary_term,omitempty"`
	Status        int           `json:"status,omitempty"`
	ForcedRefresh bool          `json:"forced_refresh,omitempty"`
	Error         *ErrorDetails `json:"error,omitempty"`
	GetResult     *GetResult    `json:"get,omitempty"`
}

// Indexed returns all bulk request results of "index" actions.
func (r *BulkResponse) Indexed() []*BulkResponseItem {
	return r.ByAction("index")
}

// Created returns all bulk request results of "create" actions.
func (r *BulkResponse) Created() []*BulkResponseItem {
	return r.ByAction("create")
}

// Updated returns all bulk request results of "update" actions.
func (r *BulkResponse) Updated() []*BulkResponseItem {
	return r.ByAction("update")
}

// Deleted returns all bulk request results of "delete" actions.
func (r *BulkResponse) Deleted() []*BulkResponseItem {
	return r.ByAction("delete")
}

// ByAction returns all bulk request results of a certain action,
// e.g. "index" or "delete".
func (r *BulkResponse) ByAction(action string) []*BulkResponseItem {
	if r.Items == nil {
		return nil
	}
	var items []*BulkResponseItem
	for _, item := range r.Items {
		if result, found := item[action]; found {
			items = append(items, result)
		}
	}
	return items
}

// ByID returns all bulk request results of a given document id,
// regardless of the action ("index", "delete" etc.).
func (r *BulkResponse) ByID(id string) []*BulkResponseItem {
	if r.Items == nil {
		return nil
	}
	var items []*BulkResponseItem
	for _, item := range r.Items {
		for _, result := range item {
			if result.ID == id {
				items = append(items, result)
			}
		}
	}
	return items
}

// Failed returns those items of a bulk response that have errors,
// i.e. those that don't have a status code between 200 and 299.
func (r *BulkResponse) Failed() []*BulkResponseItem {
	if r.Items == nil {
		return nil
	}
	var errors []*BulkResponseItem
	for _, item := range r.Items {
		for _, result := range item {
			if !(result.Status >= 200 && result.Status <= 299) {
				errors = append(errors, result)
			}
		}
	}
	return errors
}

// Succeeded returns those items of a bulk response that have no errors,
// i.e. those have a status code between 200 and 299.
func (r *BulkResponse) Succeeded() []*BulkResponseItem {
	if r.Items == nil {
		return nil
	}
	var succeeded []*BulkResponseItem
	for _, item := range r.Items {
		for _, result := range item {
			if result.Status >= 200 && result.Status <= 299 {
				succeeded = append(succeeded, result)
			}
		}
	}
	return succeeded
}

// ErrorDetails encapsulate error details from Elasticsearch.
// It is used in e.g. elastic.Error and elastic.BulkResponseItem.
type ErrorDetails struct {
	Type         string                   `json:"type"`
	Reason       string                   `json:"reason"`
	ResourceType string                   `json:"resource.type,omitempty"`
	ResourceId   string                   `json:"resource.id,omitempty"`
	Index        string                   `json:"index,omitempty"`
	Phase        string                   `json:"phase,omitempty"`
	Grouped      bool                     `json:"grouped,omitempty"`
	CausedBy     map[string]interface{}   `json:"caused_by,omitempty"`
	RootCause    []*ErrorDetails          `json:"root_cause,omitempty"`
	Suppressed   []*ErrorDetails          `json:"suppressed,omitempty"`
	FailedShards []map[string]interface{} `json:"failed_shards,omitempty"`
	Header       map[string]interface{}   `json:"header,omitempty"`

	// ScriptException adds the information in the following block.

	ScriptStack []string             `json:"script_stack,omitempty"` // from ScriptException
	Script      string               `json:"script,omitempty"`       // from ScriptException
	Lang        string               `json:"lang,omitempty"`         // from ScriptException
	Position    *ScriptErrorPosition `json:"position,omitempty"`     // from ScriptException (7.7+)
}

// ShardsInfo represents information from a shard.
type ShardsInfo struct {
	Total      int                              `json:"total"`
	Successful int                              `json:"successful"`
	Failed     int                              `json:"failed"`
	Failures   []*ShardOperationFailedException `json:"failures,omitempty"`
	Skipped    int                              `json:"skipped,omitempty"`
}

type ShardOperationFailedException struct {
	Shard  int                    `json:"shard,omitempty"`
	Index  string                 `json:"index,omitempty"`
	Status string                 `json:"status,omitempty"`
	Reason map[string]interface{} `json:"reason,omitempty"`

	// TODO(oe) Do we still have those?
	Node string `json:"_node,omitempty"`
	// TODO(oe) Do we still have those?
	Primary bool `json:"primary,omitempty"`
}

// ScriptErrorPosition specifies the position of the error
// in a script. It is used in ErrorDetails for scripting errors.
type ScriptErrorPosition struct {
	Offset int `json:"offset"`
	Start  int `json:"start"`
	End    int `json:"end"`
}

// GetResult is the outcome of GetService.Do.
type GetResult struct {
	Index       string                 `json:"_index"`   // index meta field
	Type        string                 `json:"_type"`    // type meta field
	Id          string                 `json:"_id"`      // id meta field
	Uid         string                 `json:"_uid"`     // uid meta field (see MapperService.java for all meta fields)
	Routing     string                 `json:"_routing"` // routing meta field
	Parent      string                 `json:"_parent"`  // parent meta field
	Version     int64                  `json:"_version"` // version number, when Version is set to true in SearchService
	SeqNo       int64                  `json:"_seq_no"`
	PrimaryTerm int64                  `json:"_primary_term"`
	Source      json.RawMessage        `json:"_source,omitempty"`
	Found       bool                   `json:"found,omitempty"`
	Fields      map[string]interface{} `json:"fields,omitempty"`
	// Error     string                 `json:"error,omitempty"` // used only in MultiGet
	// TODO double-check that MultiGet now returns details error information
	Error *ErrorDetails `json:"error,omitempty"` // only used in MultiGet
}

// Response represents a response from Elasticsearch.
type response struct {
	// StatusCode is the HTTP status code, e.g. 200.
	StatusCode int
	// Header is the HTTP header from the HTTP response.
	// Keys in the map are canonicalized (see http.CanonicalHeaderKey).
	Header http.Header
	// Body is the deserialized response body. Only available if streaming is disabled.
	Body json.RawMessage
	// DeprecationWarnings lists all deprecation warnings returned from
	// Elasticsearch.
	DeprecationWarnings []string
}

func (res response) Unmarshal(v any) error {
	return json.Unmarshal(res.Body, v)
}
