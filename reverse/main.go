// Savchenko Arthur
// Varian with reverse service

package main
import (
	
    "io/ioutil"
   	"fmt"
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
    "strconv"
    "bytes"
)


// Transport structure
type transport struct {
     http.RoundTripper
}

var _ http.RoundTripper = &transport{}
type Mst map[string]interface{}                              


//**********************************************************
// Main 
//**********************************************************
func main() {

    
    fmt.Println("Start service on port 1999")

    // Active node
    Host  := "108.61.245.170"
    Port  := ":1999" 

    // For test with other resources
    proxy := httputil.NewSingleHostReverseProxy(&url.URL{Scheme:"http", Host:Host})
    proxy.Transport = &transport{http.DefaultTransport}

    // Proxy
    proxy.Director = func(req *http.Request) {
        // Allows
        req.Header.Set("Access-Control-Allow-Origin","*")
        req.Header.Set("Access-Control-Allow-Headers","X-Requested-With")
        req.Header.Set("X-Forwarded-For",Host)
       
        req.Host       = Host
        req.URL.Host   = Host
        req.URL.Scheme = "http"    
    }

    // Proxy 
    http.Handle("/", proxy)
 
    err := http.ListenAndServe(Port, nil)
    
    // Error
    if err != nil {
       log.Println("Error start service.",err.Error())
    }
}

//************************************************************
//  Transport
//************************************************************
func (t *transport) RoundTrip(req *http.Request) (resp *http.Response, err error) {

  resp, err = t.RoundTripper.RoundTrip(req)
  if err != nil {
     return nil, err
  }

  b, err := ioutil.ReadAll(resp.Body)
  if err != nil {
     return nil, err
  }

  err = resp.Body.Close()
  if err != nil {
     return nil, err
  }

  b         = bytes.Replace(b, []byte("server"), []byte("schmerver"), -1)
  body     := ioutil.NopCloser(bytes.NewReader(b))
  resp.Body = body
  resp.ContentLength = int64(len(b))
  resp.Header.Set("Content-Length", strconv.Itoa(len(b)))
  resp.Header.Set("Content Type", "image/png")
  
   // w.Header().Add(c)
  return resp, nil
}


//************************************************************
// Set the proxied request's host to the destination host (instead of the
// source host).  e.g. http://foo.com proxying to http://bar.com will ensure
// that the proxied requests appear to be coming from http://bar.com
//
// For both this function and queryCombiner (below), we'll be wrapping a
// Handler with our own HandlerFunc so that we can do some intermediate work
//************************************************************
func sameHost(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
           r.Host = r.URL.Host
           handler.ServeHTTP(w, r)
    })
}

//************************************************************
// Allow cross origin resource sharing
//************************************************************
func addCORS(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json;charset=utf-8")
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With")
        handler.ServeHTTP(w, r)
    })
}

/***************************************************************
  Check Eror
 ****************************************************************/
func Err(Er error, Txt string) {
    if Er != nil {
       log.Println("ERROR : " + Txt)
       return
    }
}
