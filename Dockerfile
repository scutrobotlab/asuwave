FROM node:lts as build-env
WORKDIR /app
COPY ./package.json ./package-lock.json /app/
RUN npm ci

FROM build-env as build
COPY ./babel.config.js ./vue.config.js /app/
COPY ./public /app/public/
COPY ./src /app/src/
RUN npm run build

FROM golang:latest
WORKDIR /build
COPY ./main.go ./go.sum ./go.mod /build/
COPY ./datautil /build/datautil/
COPY ./fromelf /build/fromelf/
COPY ./option /build/option/
COPY ./serial /build/serial/
COPY ./server /build/server/
COPY ./variable /build/variable/
RUN go build -ldflags="-w -s" -o asuwave
COPY --from=build /app/dist/ /build/dist/

CMD ./asuwave -p $PORT
