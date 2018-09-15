FROM golang as test

COPY . $GOPATH/src/github.com/vanstee/nx584
WORKDIR $GOPATH/src/github.com/vanstee/nx584

RUN go get -d -v


FROM test as build

RUN CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' -o /go/bin/zones github.com/vanstee/nx584/cmd/zones


FROM scratch AS production

COPY --from=build /go/bin/zones /bin/zones
COPY --from=build /bin/sh /bin/sh

CMD /bin/zones
