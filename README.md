<h1 align="center">
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
$ go get -u github.com/shivangx01b/CorsMe
```
## Usage

Single Url
```plain
echo "https://example.com" | ./Corsme 
```
Multiple Url
```plain
cat http_https.txt | ./CorsMe -t 70
```
Allow wildcard .. Now if Access-Control-Allow-Origin is * it will be printed
```plain
cat http_https.txt | ./CorsMe -t 70 --wildcard
```
Tip
```plain
cat subdomains.txt | ./httprobe -c 70 -p 80,443,8080,8081,8089 | tee http_https.txt
cat http_https.txt | ./CorsMe -t 70
```
## Screenshot
![1414](https://github.com/Shivangx01b/CorsMe/blob/master/static/action.png)

## Note:

- Scanner stores the error results as "error_requests.txt"... which contains hosts which cannot be requested

## Idea for making this tools are taken from :
[CORScanner](https://github.com/chenjj/CORScanner)

[Corsy](https://github.com/s0md3v/Corsy)

[cors-blimey](https://github.com/tomnomnom/hacks/tree/master/cors-blimey)
