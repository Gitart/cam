// Author : Svachenko Arthur
// 

package main


import(
	    "fmt"
	    "net/http"
	    "net/http/httputil"
        "os"
        "os/signal"
        "flag"
        "syscall"
        "context"
        "time"
        "log"
        "io/ioutil"
)




// 
// Main procedure
// 
func main(){

  // Recovered
  defer func() {
    if r := recover(); r != nil {
       fmt.Println("Recovered ", r)
    }
  }()

  // Flags set
  Port := flag.String("p", "1999", "Input Port") // Port by default
  flag.Parse()


  // Handles
  http.HandleFunc("/",                     ApiHelp)                // Description and short description service
  http.HandleFunc("/api/cam/",             ApiCam)                 // Main procedure preview web camera
  http.HandleFunc("/api/cams/",            ApiCamVariant)          // Main procedure preview web camera
  

 // Settings service
  srv := &http.Server{
  	Addr: ":" + *Port,
    IdleTimeout:  20 * time.Second,
    ReadTimeout:  1  * time.Second,
    WriteTimeout: 1  * time.Second,
  }

  // srv :=makeServerFromMux()
  // Start Server
  go func() {
     FgGreen("Start web cam service.")
     FgRed("Current Port :" + *Port)


    if err := srv.ListenAndServe(); err != nil {
       log.Fatal(err)
    }
  }()

  // Graceful Shutdown
  waitForShutdown(srv)
}




//
// Get picture from web cam
//
func ApiCam(w http.ResponseWriter, r *http.Request) {
    ip := ipmain + "image.jpg"
    li := 0

    w.Header().Add("Content Type", "image/png")


// Loop 
for {
	  response, err := http.Get(ip)     
     
      if err != nil {
         fmt.Printf("Error %s", err)
      }  

       defer response.Body.Close()
       content, err := ioutil.ReadAll(response.Body)
    
       if err != nil {
          fmt.Printf("%s", err)
        }

       lc:=len(content)

       // Check update information from site about content
       if li!=lc {
          log.Printf("Change page\n")
          w.Write([]byte(content))   
          
          go RefreshPage(w,r)
           
         }else{

           // For control update information
           log.Printf("NO change page!\n")

        } 

        li=lc

        // This timeout can change for example in settings
        time.Sleep(2 * time.Second)

    }


    

    // hdrs:=r.Header.Get("Last-Modified")
    // ltms:=r.ContentLength //Header.Get("Content-Length")
    //  ua := r.Header.Get("User-Agent")

    // hdr:=response.Header.Get("Content-Length")
    // ltm:=response.Header.Get("Last-Modified")

 
    



    // // fmt.Println(response["Last-Modified"])
    // fmt.Println("Header", hdr)
    // fmt.Println("Header s", hdrs)
    // fmt.Println("Header ", ltm)
    // fmt.Println("Last Mod s->", ltms)
    // fmt.Println("User agent", ua)
    

     // dump, _ := httputil.DumpRequest(r, true)
    // fmt.Print(dump)     


    

}


/*
response, _, err := http.Get("http://golang.org/")
    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("%s", err)
            os.Exit(1)
        }
        fmt.Printf("%s\n", string(contents))
    }

*/


func RefreshPage_old(path string) func(http.ResponseWriter, *http.Request) { 
  return func (w http.ResponseWriter, r *http.Request) {
    // http.Redirect(w, r, path, http.StatusMovedPermanently)
    http.Redirect(w, r, path, 303)
    fmt.Println("OK ")
  }


}


func RefreshPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("OK ")
	http.Redirect(w, r, "/api/cam/", 303) 
}



// 
// Help page
// 
func ApiHelp(w http.ResponseWriter, r *http.Request) {
	 w.Write([]byte(hlp))
}


//
// Shootdown service
//
func waitForShutdown(srv *http.Server) {

  interruptChan := make(chan os.Signal, 1)
  signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

  // Block until we receive our signal.
  <-interruptChan

  // Create a deadline to wait for.
  ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
  defer cancel()
  srv.Shutdown(ctx)

  // Notify shutdown server
  FgRed("Shutting down server.\n")
  os.Exit(0)
}
