package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
	tld "github.com/jpillora/go-tld"
	"github.com/fatih/color"
	"github.com/briandowns/spinner"
	"flag"
)


var f,  _ = os.Create("error_requests.txt")
var Threads int
 

func Banner() {
	color.HiGreen(`
 _____                ___  ___     
/  __ \               |  \/  |     
| /  \/ ___  _ __ ___ | .  . | ___ 
| |    / _ \| '__/ __|| |\/| |/ _ \
| \__/\ (_) | |  \__ \| |  | |  __/
 \____/\___/|_|  |___/\_|  |_/\___|
								   `)
	color.HiRed("                 " + "Made with <3 by @shivangx01b")
	
}

func getClient() *http.Client {          // taken from https://github.com/tomnomnom/hacks/blob/master/cors-blimey/main.go
	tr := &http.Transport{
		MaxIdleConns:    30,
		IdleConnTimeout: time.Second,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		DialContext: (&net.Dialer{
			Timeout:   time.Second * 10,
			KeepAlive: time.Second,
		}).DialContext,
	}

	re := func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	return &http.Client{
		Transport:     tr,
		CheckRedirect: re,
		Timeout:       time.Second * 10,
	}
}

func requester(c *http.Client, u string, origins []string) {
	for _, p := range origins {

		req, err := http.NewRequest("GET", u, nil)
		if err != nil {
			return
		}
		req.Header.Set("Origin", p)

		resp, err := c.Do(req)
		if resp != nil {
			io.Copy(ioutil.Discard, resp.Body)
			resp.Body.Close()
		}
		if err != nil {
			w := bufio.NewWriter(f) //stores all the urls which cannot be resolved 
			w.WriteString(u)
			w.Flush()
			return
		}

		acao := resp.Header.Get("Access-Control-Allow-Origin")
		acac := resp.Header.Get("Access-Control-Allow-Credentials")
		
		if acao == p {
			color.HiRed("\n[-] Misconfiguration found!")
			color.HiCyan("URL: %s", u)
			color.HiGreen("Access-Control-Allow-Origin: %s\n", p)
			
				if acac == "true"{
					color.HiGreen("Access-Control-Allow-Credentials: %s", acac)
					
				}							
		
		}
	}

}

func parser(u string) ([]string,  error) {
	var parsed[] string 
	url, err := tld.Parse(u)
	if err != nil {
		return []string{}, err
	} 
	subdomain := url.Subdomain
	domain := url.Domain
	tld := url.TLD
	return append(parsed, subdomain, domain, tld), nil 
}

func anyorigin() []string {
	origins := []string{
		"*",
		"http://shivangx01b.com",
		"https://shivangx01b.com",
	}
	return origins
	
}
	

func prefix(things []string) []string {
	origins := []string{"https://" + things[1] + ".shivangx01b.com", "https://" + things[1] + "." +  things[2] + ".shivangx01b.com"}	
	return origins
}

func suffix(things []string) []string {
	origins := []string{"https://" + "shivangx01b" + things[1] + "." + things[2], "https://" + "shivangx01b.com" + "." + things[1] + "." + things[2]}	
	return origins
}

func notescapedot(things []string) []string  {
	origins := []string{"https://" + things[0] + "S" + things[1] + things[2]}	
	return origins
}


func null() []string  {
	origins := []string{"null"}	
	return origins
}

func thirdparties() []string {           //taken from https://github.com/chenjj/CORScanner/blob/master/origins.json
	origins := []string{
		"https://shivangx01b.github.io",
		"http://jsbin.com",
		"https://codepen.io",
		"https://jsfiddle.net",
		"http://www.webdevout.net",
		"https://repl.it",
	}
	return origins
}

func spicalchars(things []string) []string {
	origins := []string{}
	chars := []string{"_","-","{","}","^","%60","!","~","`",";","|","&","(",")","*","'","\"","$","=","+","%0b"}
	permute := []string{"https://" + things[0] + "." + things[1] + "." + things[2] + "%s" + ".shivangx01b.com"}
	for i, per := range permute {
		for _, char := range chars{
			permute[i] = fmt.Sprintf(per, char)
			origins = append(origins, permute[i])
		}
	} 
	return origins
}

func totalwaystotest(c *http.Client, u string, )  { 	
	things, _ := parser(u)
	
	AnyOrigin := anyorigin()
	requester(c, u, AnyOrigin)
	
	Prefix := prefix(things)
	requester(c, u, Prefix)
	
	Suffix := suffix(things)
	requester(c, u, Suffix)
	
	Escaped := notescapedot(things)
	requester(c, u, Escaped)
	
	Null := null()
	requester(c, u, Null)

	Third := thirdparties()
	requester(c, u, Third)

	Specialchars := spicalchars(things)
	requester(c, u, Specialchars)

}


func ParseArguments() {
	flag.IntVar(&Threads, "t", 40, "Number of workers to use..default 40")	
	flag.Parse()
}


func main() {
	ParseArguments()
	Banner()
	color.HiBlue("\n[~] Total Tests.. 🛠")
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)  //Just some shity effects 
	s.Start()              
	time.Sleep(2 * time.Second)                                  
	s.Stop()
	color.HiYellow("\n[>] Reflect Origin.. 🔍")
	color.HiYellow("\n[>] Prefix Matches.. 🔍")
	color.HiYellow("\n[>] Suffix Matches.. 🔍")
	color.HiYellow("\n[>] Not Escaped Dots.. 🔍")
	color.HiYellow("\n[>] Null.. 🔍")
	color.HiYellow("\n[>] Common ThirdParties.. 🔍")
	color.HiYellow("\n[>] Special Chars.. 🔍")
	s.UpdateCharSet(spinner.CharSets[4]) // Ahh shit.. here we go again
	s.Restart()
	time.Sleep(2 * time.Second)
	s.Stop()
	urls := make(chan string, Threads)
	processGroup := new(sync.WaitGroup)
	processGroup.Add(Threads)

	for i := 0; i < Threads; i++ {
		c := getClient()
		go func() {
			defer processGroup.Done()
			for u := range urls {
				totalwaystotest(c, u)
			}
		}()
	}

	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		urls <- sc.Text()
	}
	close(urls)
	processGroup.Wait()
}