**basic auth**

```
$ go run main.go -token "$GITHUB_TOKEN" -auth basic &
$ git  -c http.sslVerify=false clone https://localhost/superbrothers/private.git
Cloning into 'private'...
remote: Enumerating objects: 18, done.
remote: Counting objects: 100% (18/18), done.
remote: Compressing objects: 100% (16/16), done.
remote: Total 18 (delta 0), reused 18 (delta 0), pack-reused 0
Receiving objects: 100% (18/18), 16.85 KiB | 8.43 MiB/s, done.
```

**token:**

```
$ go run main.go -token "$GITHUB_TOKEN" -auth token &
$ git  -c http.sslVerify=false clone https://localhost/superbrothers/private.git
Cloning into 'repo'...
fatal: unable to access 'https://localhost/superbrothers/private.git/': The requested URL returned error: 400```
