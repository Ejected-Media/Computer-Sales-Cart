 // # ~ . ~ Computer Sales Cart - website
// ,  ° . +
package main

import (
    "os"
    "log"
    
    "fmt"
		
	"text/template"
	"net/http"
	
//	"time"

)

// ,  ° . +
type htmlPageData struct {
    pageTitle string
    pagePath string
    pageList []pageNav
    
}

type pageNav struct {
    pageTitle string
    pageLink string
}

type pageList struct {
    pageTitle string
    pageLink string
}


type navList struct {
    Title string
    Done  bool
}


// ,  ° . +
type SOSPageData struct {
    PageTitle string
    PagePath string
    PageName string
    SOSNav     []navList
}




// ,  ° . +
func app_welcome_center_page() {


}

func hello(w http.ResponseWriter, r *http.Request) {
//	pagePath := r.URL.Path
    fmt.Fprintf(w, "Hello!")
}

func world(w http.ResponseWriter, r *http.Request) {
//	pagePath := r.URL.Path
    fmt.Fprintf(w, "World!")
}

// . testHandler,  ~ for Public Pages °
func testHandler(w http.ResponseWriter, r *http.Request) {
// ,  ° . +
    if r.URL.Path != "/" {
    	http.NotFound(w, r)
    	return
    }

// , ° . +
  pageTitle := "~ . - // - Website App"
  pagePath := r.URL.Path
  // pageType := ".."


// ,  ° . +
pageData := htmlPageData {
      pageTitle: pageTitle,
      pagePath: pagePath,
      
    //  pageList: []pageNav {
  //        { pageTitle: "one", pageLink: "one"},
//          { pageTitle: "two", pageLink: "two"},
   //       { pageTitle: "three", pageLink: "three"},
    //  },
  	
  }  //. .  pageData
  
  pageHTML := "layout_main_page.html";
  if pagePath == "/" {
      pageTitle = "Index Page"
 //     pageList = pageList
  }
 


// ,  ° . +
  pageFilePath := template.Must(
    template.ParseFiles(pageHTML))
  pageFilePath.Execute(w, pageData)
  
}  //  .  testHandler


// ,  ° . +
func worldLoader(w http.ResponseWriter, r *http.Request) {
	pagePath := r.URL.Path
    // fmt.Fprintf(w, "World!")
    pageName := "test"
    
  if pagePath == "/page/Menu" {
    pageName = "Menu Page"
    //     pageList = pageList
  }
  
  if pagePath == "/page/Location" {
    pageName = "Location Page"
    //     pageList = pageList
  }
  
  if pagePath == "/page/Equiptment" {
    pageName = "Equiptment Page"
    //     pageList = pageList
  }
  
  if pagePath == "/page/Staff" {
    pageName = "Staff Page"
    //     pageList = pageList
  }
  
  if pagePath == "/page/Services" {
    pageName = "Services Page"
    //     pageList = pageList
  }
  
  if pagePath == "/page/Music" {
    pageName = "Music Page"
    //     pageList = pageList
  }
  
  if pagePath == "/page/Setting" {
    pageName = "Setting Page"
    //     pageList = pageList
  }
  
  if pagePath == "/page/Sources" {
    pageName = "Sources Page"
    //     pageList = pageList
  }
  
  if pagePath == "/page/Marketing" {
    pageName = "Marketing Page"
    //     pageList = pageList
  }
  
  if pagePath == "/page/Operations" {
    pageName = "Operations Page"
    //     pageList = pageList
  }
  
  if pagePath == "/page/Software" {
    pageName = "Software Page"
    //     pageList = pageList
  }
  
  if pagePath == "/page/Events" {
    pageName = "Events Page"
    //     pageList = pageList
  }
  
  if pagePath == "/page/Seating" {
    pageName = "Seating Page"
    //     pageList = pageList
  }
  
  if pagePath == "/page/Hours" {
    pageName = "Hours Page"
    //     pageList = pageList
  }
  
  // ,  ° . +
    data := SOSPageData{
            PageTitle: pageName,
            PagePath: pageName,
            PageName: pageName,
            SOSNav: []navList{
                {Title: "Timer", Done: false},
                {Title: "Amount Conversion", Done: true},
                {Title: "Recipe Stuff", Done: true},
                {Title: "Camera", Done: true},
                {Title: "Notes", Done: true},
                {Title: "Session Log", Done: true},
        
            },
        }
 
 pageHTML := "layout_main_page.html";
 
  // ,  ° . +
  pageFilePath := template.Must(
    template.ParseFiles(pageHTML))
  pageFilePath.Execute(w, data)
}


//  .  html url routes 
//  .  as well as terminal cli logs

func main() {
// ,  ° . +
  appName := "~ . - Computer Sales Cart // - Website App"
  
  http.HandleFunc("/hello", hello)
    http.HandleFunc("/world", world)
  // ,  ° . +
   //  http.HandleFunc("/page/kitchen", kitchenPage_world)
    
          http.HandleFunc("/", testHandler)

// ,  ° . +
  http.HandleFunc("/Menu", worldLoader)
  http.HandleFunc("/Location", worldLoader)
  http.HandleFunc("/Equiptment", worldLoader)
  http.HandleFunc("/Staff", worldLoader)
  http.HandleFunc("/Services", worldLoader)
  http.HandleFunc("/Music", worldLoader)
  http.HandleFunc("/Setting", worldLoader)
  http.HandleFunc("/Sources", worldLoader)
  http.HandleFunc("/Marketing", worldLoader)
  http.HandleFunc("/Operations", worldLoader)
  http.HandleFunc("/Software", worldLoader)
  http.HandleFunc("/Events", worldLoader)
  http.HandleFunc("/Seating", worldLoader)
  http.HandleFunc("/Hours", worldLoader)
  

  
  // ,  ° . +
  http.HandleFunc("/one", testHandler)

// . ° ~ +
 // http.HandleFunc("/page/classSchedule", indexHandler)

// -- -
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
    log.Printf("Loading _webapp with default port")
  }
  
// ,  ° . +
  log.Printf("_webapp is active and Listening on port %s", port)

  log.Printf("// -- - %s", appName)
  log.Printf("_webapp now loaded and running at http://localhost:%s", port)

// -- - 
  if err := http.ListenAndServe(":"+port, nil); err != nil {
    log.Fatal("Error Starting the HTTP Server :", err)
    return
  }}