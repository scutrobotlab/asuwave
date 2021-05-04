FROM node:lts AS build-env
WORKDIR /app
COPY ./package.json ./package-lock.json /app/
RUN npm ci

FROM build-env AS build
COPY ./babel.config.js ./vue.config.js /app/
COPY ./public /app/public/
COPY ./src /app/src/
RUN npm run build

FROM golang:latest AS binary
WORKDIR /build
COPY ./main.go ./main_release.go ./go.sum ./go.mod /build/
COPY ./datautil /build/datautil/
COPY ./fromelf /build/fromelf/
COPY ./option /build/option/
COPY ./serial /build/serial/
COPY ./server /build/server/
COPY ./variable /build/variable/
COPY --from=build /app/dist/ /build/dist/
RUN CGO_ENABLED=0 go build -tags release -ldflags="-w -s" -o asuwave

FROM bash:latest
COPY --from=binary /build/asuwave /

CMD /asuwave -p $PORT
