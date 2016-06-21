To build

<pre>
docker run --rm -it -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.6 bash
go build -v -o ping
</pre>