FROM golang:1.19-alpine3.15 AS Build
COPY . workspace
RUN cd workspace && go get && go build -o wallet.out
FROM alpine:3.15
COPY --from=Build /go/workspace/wallet.out /wallet.out
EXPOSE 3001
ENTRYPOINT ["./wallet.out"]