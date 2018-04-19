# build stage
FROM golang:alpine AS build-env
#ADD . /src
#RUN go build -a -ldflags "-linkmode external -extldflags -static" -o hackacraic
#RUN go build -o hackacraic
RUN apk --no-cache add git bzr mercurial
#RUN cd /src && go get -d -v
#RUN cd /src && go build -o hackacraic

FROM build-env AS staging
ENV GOPATH /go
RUN echo $GOPATH
#ADD . $GOPATH/src
WORKDIR $GOPATH/src/github.com/ancient87/hackacraic
ADD . .
RUN go get 
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o app/hackacraic
RUN ls -la
RUN ls -la app

# final stage
FROM staging as app
WORKDIR /app
COPY --from=staging /go/src/github.com/ancient87/hackacraic/app/hackacraic/ /app/
EXPOSE 80
ENTRYPOINT ./hackacraic
