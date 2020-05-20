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
echo "https://example.com" | corsme 
```
Multiple Url
```plain
cat http_https.txt | corsme -t 70
```
Tip
```plain
echo $targetdomain | amass enum -passive -d  | sort -u | httprobe -c 70 -p 80,443,8080,8081,8089 | tee http_https.txt
cat http_https.txt | corsme -t 70
```

