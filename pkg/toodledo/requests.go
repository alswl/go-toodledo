package toodledo

import (
	"crypto/tls"
	"encoding/json"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"time"
)

type Requests struct {
	gorequest.SuperAgent
}

func (r *Requests) Auth(accessToken string) *Requests {
	r.Param("access_token", accessToken)
	return r
}

// EndStructWithTError should be used when you want the body as a struct.
// The callbacks work the same way as with `End`, except that a struct is used instead of a string.
// the err contains general error process of Toodledo
func (r *Requests) EndStructWithTError(v interface{}, callback ...func(response gorequest.Response, v interface{}, body []byte, err []error)) (gorequest.Response, []byte, []error) {
	resp, body, errs := r.EndStruct(v, callback...)
	if errs != nil {
		firstErr := errs[0]
		if _, ok := firstErr.(*json.UnmarshalTypeError); ok {
			errResp, ok := CheckToodledoResponse(string(body))
			if !ok {
				return resp, body, []error{ApiError{
					errResp,
					resp,
					string(body),
				}}
			}
		}
		return resp, body, errs
	}
	return resp, body, errs
}

// wrapper >>>

// Returns a copy of this superagent. Useful if you want to reuse the client/settings
// concurrently.
// Note: This does a shallow copy of the parent. So you will need to be
// careful of Data provided
// Note: It also directly re-uses the client and transport. If you modify the Timeout,
// or RedirectPolicy on a clone, the clone will have a new http.client. It is recommended
// that the base request set your timeout and redirect polices, and no modification of
// the client or transport happen after cloning.
// Note: DoNotClearRequests is forced to "true" after Clone
func (r *Requests) Clone() *Requests {
	r.SuperAgent.Clone()
	return r
}

// Enable the debug mode which logs request/response detail
func (r *Requests) SetDebug(enable bool) *Requests {
	r.SuperAgent.SetDebug(enable)
	return r
}

// Enable the curlcommand mode which display a CURL command line
func (r *Requests) SetCurlCommand(enable bool) *Requests {
	r.SuperAgent.SetCurlCommand(enable)
	return r
}

// Enable the DoNotClear mode for not clearing super agent and reuse for the next request
func (r *Requests) SetDoNotClearRequests(enable bool) *Requests {
	r.SuperAgent.SetDoNotClearSuperAgent(enable)
	return r
}

func (r *Requests) SetLogger(logger gorequest.Logger) *Requests {
	r.SuperAgent.SetLogger(logger)
	return r
}

// Just a wrapper to initialize Requests instance by method string
func (r *Requests) CustomMethod(method, targetUrl string) *Requests {
	r.SuperAgent.CustomMethod(method, targetUrl)
	return r
}

func (r *Requests) Get(targetUrl string) *Requests {
	r.SuperAgent.Get(targetUrl)
	return r
}

func (r *Requests) Post(targetUrl string) *Requests {
	r.SuperAgent.Post(targetUrl)
	return r
}

func (r *Requests) Head(targetUrl string) *Requests {
	r.SuperAgent.Head(targetUrl)
	return r
}

func (r *Requests) Put(targetUrl string) *Requests {
	r.SuperAgent.Put(targetUrl)
	return r
}

func (r *Requests) Delete(targetUrl string) *Requests {
	r.SuperAgent.Delete(targetUrl)
	return r
}

func (r *Requests) Patch(targetUrl string) *Requests {
	r.SuperAgent.Patch(targetUrl)
	return r
}

func (r *Requests) Options(targetUrl string) *Requests {
	r.SuperAgent.Options(targetUrl)
	return r
}

// Set is used for setting header fields,
// this will overwrite the existed values of Header through AppendHeader().
// Example. To set `Accept` as `application/json`
//
//    gorequest.New().
//      Post("/gamelist").
//      Set("Accept", "application/json").
//      End()
func (r *Requests) Set(param string, value string) *Requests {
	r.SuperAgent.Set(param, value)
	return r
}

// AppendHeader is used for setting header fileds with multiple values,
// Example. To set `Accept` as `application/json, text/plain`
//
//    gorequest.New().
//      Post("/gamelist").
//      AppendHeader("Accept", "application/json").
//      AppendHeader("Accept", "text/plain").
//      End()
func (r *Requests) AppendHeader(param string, value string) *Requests {
	r.SuperAgent.AppendHeader(param, value)
	return r
}

// Retryable is used for setting a Retryer policy
// Example. To set Retryer policy with 5 seconds between each attempt.
//          3 max attempt.
//          And StatusBadRequest and StatusInternalServerError as RetryableStatus

//    gorequest.New().
//      Post("/gamelist").
//      Retry(3, 5 * time.seconds, http.StatusBadRequest, http.StatusInternalServerError).
//      End()
func (r *Requests) Retry(retryerCount int, retryerTime time.Duration, statusCode ...int) *Requests {
	r.SuperAgent.Retry(retryerCount, retryerTime, statusCode...)
	return r
}

// SetBasicAuth sets the basic authentication header
// Example. To set the header for username "myuser" and password "mypass"
//
//    gorequest.New()
//      Post("/gamelist").
//      SetBasicAuth("myuser", "mypass").
//      End()
func (r *Requests) SetBasicAuth(username string, password string) *Requests {
	r.SuperAgent.SetBasicAuth(username, password)
	return r
}

// AddCookie adds a cookie to the request. The behavior is the same as AddCookie on Request from net/http
func (r *Requests) AddCookie(c *http.Cookie) *Requests {
	r.SuperAgent.AddCookie(c)
	return r
}

// AddCookies is a convenient method to add multiple cookies
func (r *Requests) AddCookies(cookies []*http.Cookie) *Requests {
	r.SuperAgent.AddCookies(cookies)
	return r
}

// Type is a convenience function to specify the data type to send.
// For example, to send data as `application/x-www-form-urlencoded` :
//
//    gorequest.New().
//      Post("/recipe").
//      Type("form").
//      Send(`{ "name": "egg benedict", "category": "brunch" }`).
//      End()
//
// This will POST the body "name=egg benedict&category=brunch" to url /recipe
//
// GoRequest supports
//
//    "text/html" uses "html"
//    "application/json" uses "json"
//    "application/xml" uses "xml"
//    "text/plain" uses "text"
//    "application/x-www-form-urlencoded" uses "urlencoded", "form" or "form-data"
//
func (r *Requests) Type(typeStr string) *Requests {
	r.SuperAgent.Type(typeStr)
	return r
}

// Query function accepts either json string or strings which will form a query-string in url of GET method or body of POST method.
// For example, making "/search?query=bicycle&size=50x50&weight=20kg" using GET method:
//
//      gorequest.New().
//        Get("/search").
//        Query(`{ query: 'bicycle' }`).
//        Query(`{ size: '50x50' }`).
//        Query(`{ weight: '20kg' }`).
//        End()
//
// Or you can put multiple json values:
//
//      gorequest.New().
//        Get("/search").
//        Query(`{ query: 'bicycle', size: '50x50', weight: '20kg' }`).
//        End()
//
// Strings are also acceptable:
//
//      gorequest.New().
//        Get("/search").
//        Query("query=bicycle&size=50x50").
//        Query("weight=20kg").
//        End()
//
// Or even Mixed! :)
//
//      gorequest.New().
//        Get("/search").
//        Query("query=bicycle").
//        Query(`{ size: '50x50', weight:'20kg' }`).
//        End()
//
func (r *Requests) Query(content interface{}) *Requests {
	r.SuperAgent.Query(content)
	return r
}

// As Go conventions accepts ; as a synonym for &. (https://github.com/golang/go/issues/2210)
// Thus, Query won't accept ; in a querystring if we provide something like fields=f1;f2;f3
// This Param is then created as an alternative method to solve this.
func (r *Requests) Param(key string, value string) *Requests {
	r.SuperAgent.Param(key, value)
	return r
}

// Set TLSClientConfig for underling Transport.
// One example is you can use it to disable security check (https):
//
//      gorequest.New().TLSClientConfig(&tls.Config{ InsecureSkipVerify: true}).
//        Get("https://disable-security-check.com").
//        End()
//
func (r *Requests) TLSClientConfig(config *tls.Config) *Requests {
	r.SuperAgent.TLSClientConfig(config)
	return r
}

// Proxy function accepts a proxy url string to setup proxy url for any request.
// It provides a convenience way to setup proxy which have advantages over usual old ways.
// One example is you might try to set `http_proxy` environment. This means you are setting proxy up for all the requests.
// You will not be able to send different request with different proxy unless you change your `http_proxy` environment again.
// Another example is using Golang proxy setting. This is normal prefer way to do but too verbase compared to GoRequest's Proxy:
//
//      gorequest.New().Proxy("http://myproxy:9999").
//        Post("http://www.google.com").
//        End()
//
// To set no_proxy, just put empty string to Proxy func:
//
//      gorequest.New().Proxy("").
//        Post("http://www.google.com").
//        End()
//
func (r *Requests) Proxy(proxyUrl string) *Requests {
	r.SuperAgent.Proxy(proxyUrl)
	return r
}

// RedirectPolicy accepts a function to define how to handle redirects. If the
// policy function returns an error, the next Request is not made and the previous
// request is returned.
//
// The policy function's arguments are the Request about to be made and the
// past requests in order of oldest first.
func (r *Requests) RedirectPolicy(policy func(req gorequest.Request, via []gorequest.Request) error) *Requests {
	r.SuperAgent.RedirectPolicy(policy)
	return r
}

// Send function accepts either json string or query strings which is usually used to assign data to POST or PUT method.
// Without specifying any type, if you give Send with json data, you are doing requesting in json format:
//
//      gorequest.New().
//        Post("/search").
//        Send(`{ query: 'sushi' }`).
//        End()
//
// While if you use at least one of querystring, GoRequest understands and automatically set the Content-Type to `application/x-www-form-urlencoded`
//
//      gorequest.New().
//        Post("/search").
//        Send("query=tonkatsu").
//        End()
//
// So, if you want to strictly send json format, you need to use Type func to set it as `json` (Please see more details in Type function).
// You can also do multiple chain of Send:
//
//      gorequest.New().
//        Post("/search").
//        Send("query=bicycle&size=50x50").
//        Send(`{ wheel: '4'}`).
//        End()
//
// From v0.2.0, Send function provide another convenience way to work with Struct type. You can mix and match it with json and query string:
//
//      type BrowserVersionSupport struct {
//        Chrome string
//        Firefox string
//      }
//      ver := BrowserVersionSupport{ Chrome: "37.0.2041.6", Firefox: "30.0" }
//      gorequest.New().
//        Post("/update_version").
//        Send(ver).
//        Send(`{"Safari":"5.1.10"}`).
//        End()
//
// If you have set Type to text or Content-Type to text/plain, content will be sent as raw string in body instead of form
//
//      gorequest.New().
//        Post("/greet").
//        Type("text").
//        Send("hello world").
//        End()
//
func (r *Requests) Send(content interface{}) *Requests {
	r.SuperAgent.Send(content)
	return r
}

// SendSlice (similar to SendString) returns Requests's itself for any next chain and takes content []interface{} as a parameter.
// Its duty is to append slice of interface{} into s.SliceData ([]interface{}) which later changes into json array in the End() func.
func (r *Requests) SendSlice(content []interface{}) *Requests {
	r.SuperAgent.SendSlice(content)
	return r
}

func (r *Requests) SendMap(content interface{}) *Requests {
	r.SuperAgent.SendMap(content)
	return r
}

// SendStruct (similar to SendString) returns Requests's itself for any next chain and takes content interface{} as a parameter.
// Its duty is to transfrom interface{} (implicitly always a struct) into s.Data (map[string]interface{}) which later changes into appropriate format such as json, form, text, etc. in the End() func.
func (r *Requests) SendStruct(content interface{}) *Requests {
	r.SuperAgent.SendStruct(content)
	return r
}

// SendString returns Requests's itself for any next chain and takes content string as a parameter.
// Its duty is to transform String into s.Data (map[string]interface{}) which later changes into appropriate format such as json, form, text, etc. in the End func.
// Send implicitly uses SendString and you should use Send instead of this.
func (r *Requests) SendString(content string) *Requests {
	r.SuperAgent.SendString(content)
	return r
}

// SendFile function works only with type "multipart". The function accepts one mandatory and up to two optional arguments. The mandatory (first) argument is the file.
// The function accepts a path to a file as string:
//
//      gorequest.New().
//        Post("http://example.com").
//        Type("multipart").
//        SendFile("./example_file.ext").
//        End()
//
// File can also be a []byte slice of a already file read by eg. ioutil.ReadFile:
//
//      b, _ := ioutil.ReadFile("./example_file.ext")
//      gorequest.New().
//        Post("http://example.com").
//        Type("multipart").
//        SendFile(b).
//        End()
//
// Furthermore file can also be a os.File:
//
//      f, _ := os.Open("./example_file.ext")
//      gorequest.New().
//        Post("http://example.com").
//        Type("multipart").
//        SendFile(f).
//        End()
//
// The first optional argument (second argument overall) is the filename, which will be automatically determined when file is a string (path) or a os.File.
// When file is a []byte slice, filename defaults to "filename". In all cases the automatically determined filename can be overwritten:
//
//      b, _ := ioutil.ReadFile("./example_file.ext")
//      gorequest.New().
//        Post("http://example.com").
//        Type("multipart").
//        SendFile(b, "my_custom_filename").
//        End()
//
// The second optional argument (third argument overall) is the fieldname in the multipart/form-data request. It defaults to fileNUMBER (eg. file1), where number is ascending and starts counting at 1.
// So if you send multiple files, the fieldnames will be file1, file2, ... unless it is overwritten. If fieldname is set to "file" it will be automatically set to fileNUMBER, where number is the greatest exsiting number+1.
//
//      b, _ := ioutil.ReadFile("./example_file.ext")
//      gorequest.New().
//        Post("http://example.com").
//        Type("multipart").
//        SendFile(b, "", "my_custom_fieldname"). // filename left blank, will become "example_file.ext"
//        End()
//
func (r *Requests) SendFile(file interface{}, args ...string) *Requests {
	r.SuperAgent.SendFile(file, args...)
	return r
}

// wrapper <<<

func NewRequests() *Requests {
	agent := gorequest.New()
	return &Requests{*agent}
}
