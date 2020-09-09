<h1 align="center">
  <br>
  <a href=""><img src="https://github.com/Shivangx01b/CorsMe/blob/master/static/banner.png" alt="" width="200px;"></a>
  <br>
  <img src="https://img.shields.io/github/languages/top/Shivangx01b/CorsMe?style=flat-square">
  <a href="https://goreportcard.com/report/github.com/Shivangx01b/CorsMe"><img src="https://goreportcard.com/badge/github.com/Shivangx01b/CorsMe"></a>
  <a href="https://twitter.com/intent/follow?screen_name=shivangx01b"><img src="https://img.shields.io/twitter/follow/shivangx01b?style=flat-square"></a>
</h1>

## What is CorsMe ?
A cors misconfiguration scanner tool based on golang with speed and precision in mind !

## Misconfiguration type  this scanner can check for

- Reflect Origin checks 
- Prefix Match
- Suffix Match
- Not Esacped Dots
- Null 
- ThirdParties (Like => github.io, repl.it etc.)
  - Taken from [Chenjj's github repo](https://github.com/chenjj/CORScanner/blob/master/origins.json)
- SpecialChars (Like => "}","(", etc.)
  - See more in [Advanced CORS Exploitation Techniques](https://www.corben.io/advanced-cors-techniques/)

## How to Install

```
$ go get -u -v github.com/shivangx01b/CorsMe
```
## Usage

Single Url
```plain
echo "https://example.com" | ./CorsMe 
```
Multiple Url
```plain
cat http_https.txt | ./CorsMe -t 70
```
Allow wildcard .. Now if Access-Control-Allow-Origin is * it will be printed
```plain
cat http_https.txt | ./CorsMe -t 70 -wildcard
```
Add header if required
```plain
cat http_https.txt | ./CorsMe -t 70 -wildcard -header "Cookie: Session=12cbcx...."
```
Add another method if required
```plain
cat http_https.txt | ./CorsMe -t 70 -wildcard -header "Cookie: Session=12cbcx...." -method "POST"
```

Tip
```plain
subfinder -d hackerone.com -nW -silent | ./httprobe -c 70 -p 8080,8081,8089 | tee http_https.txt
cat http_https.txt | ./CorsMe -t 70
```
## Screenshot
![1414](https://github.com/Shivangx01b/CorsMe/blob/master/static/action.png)

## Note:

- Scanner stores the error results as "error_requests.txt"... which contains hosts which cannot be requested

## Contributers
[![Twitter](https://img.shields.io/badge/twitter-@1ndianl33t-blue.svg)](https://twitter.com/1ndianl33t)

## Ideas for making this tool are taken from :
[CORScanner](https://github.com/chenjj/CORScanner)

[Corsy](https://github.com/s0md3v/Corsy)

[cors-blimey](https://github.com/tomnomnom/hacks/tree/master/cors-blimey)
