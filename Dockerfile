FROM golang:1.19-alpine AS gobuild
WORKDIR /src

RUN apk add --update make

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY  . ./

RUN make wasm

FROM nginx:stable-alpine
WORKDIR /usr/share/nginx/html

RUN rm -rf ./*

COPY --from=gobuild /src/web .

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]




