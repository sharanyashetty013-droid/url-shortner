# URL Shortener

A simple URL shortener built in Go — turns long links into short codes and redirects you back to the original URL.

## Features

- Shorten any URL with a single API call
- MD5 hashing for unique 8-character short codes
- Instant redirect from short code to original URL

## Tech Stack

Go · `net/http` · `crypto/md5` · `encoding/json`

## Run It

```bash
go run main.go
```
Server runs on `http://localhost:3000`

## API

**Shorten a URL**
```
POST /shorten
{ "url": "https://example.com" }
```
Response: `Short URL: c8746914`

**Use it**
Visit `localhost:3000/c8746914` to be redirected to the original link.

## What I Learned

My first backend project, built from scratch as a complete beginner — Go fundamentals (structs, maps, error handling), building an HTTP server, hashing, testing with Postman, and shipping via Git/GitHub.

## Coming Next

- [ ] Persistent database storage
- [ ] JSON API responses
- [ ] Input validation
