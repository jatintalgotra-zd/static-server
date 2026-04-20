FROM --platform=$BUILDPLATFORM golang:1.26 AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} go build -ldflags="-s -w" -o /main .

FROM alpine:edge

RUN apk add --no-cache tzdata ca-certificates

COPY --from=build /main /main
COPY static ./static

RUN chmod +x /main

EXPOSE 8000

CMD ["/main"]
