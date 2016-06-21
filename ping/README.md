To build

<pre>
docker run --rm -it -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.6 bash
go build -v -o ping
</pre>

You probably need to build it like this:

<pre>
GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -v -o ping
</pre>
