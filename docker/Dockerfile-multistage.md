main.go
```
package main

import "github.com/fatih/color"
import "time"
func main(){
	color.Red("%s", time.Now())
}
```


Dockerfile
```
FROM golang:1.17

WORKDIR /src
COPY . .

RUN go build -o /bin/app .

ENTRYPOINT ["/bin/app"]
```



*1.Dockerfile optimized
```
FROM golang:1.17

WORKDIR /src
COPY go.sum go.mod ./
RUN go mod download

COPY . .

RUN go build -o /bin/app .

ENTRYPOINT ["/bin/app"]
```





docker build -t example .
docker images
REPOSITORY                   TAG       IMAGE ID       CREATED          SIZE
example                      latest    e7c5cd4137d9   3 minutes ago    1.39GB

 1133  go build .
 1134  la
 2MB

* 2.Multi stage build
 *Go hem compile bir dil, hem de static compilation a izin veriyor. sistemdeki libc gibi dependency lere bagimli degil.
 *CGO_ENABLED=0 static binary. os uzerinde calisan hic bir lib e, dosyaya ihtiyaci yok.
 Dockerfile
 ```
 FROM golang:1.17

WORKDIR /src
COPY go.sum go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /bin/app

FROM alpine
COPY --from=0 /bin/app /bin/app     -> ilk adimdan static binary yi al yeni bin/app icine kopyala
ENTRYPOINT ["/bin/app"]
```

docker images
REPOSITORY                   TAG       IMAGE ID       CREATED          SIZE
example                      latest    d40716b7b8a3   17 seconds ago   15.3MB





*3. FROM scratch. Scratch bos bir image aslinda. docker pull yapinca hicbisey olmaz. 
* Docker burda bos bir base image veriyor. bos bir image icinde python calistiramazsiniz bir cok kutuphaneye os packege ina ve python in kendisine ihtiyaci var.
* bos image icinde binary calistirilabilir. tek dosyali bir container insa etmiş oluyorum.

Dockerfile
```
FROM golang:1.17

WORKDIR /src
COPY go.sum go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /bin/app

FROM scratch
COPY --from=0 /bin/app /bin/app
ENTRYPOINT ["/bin/app"]
```

docker images 
REPOSITORY                   TAG       IMAGE ID       CREATED             SIZE
example                      scratch   b86bacb8c24a   6 minutes ago       3.23MB
example                      latest    d40716b7b8a3   12 minutes ago      15.3MB


