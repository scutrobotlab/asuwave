FROM node:lts AS build-env
WORKDIR /app
COPY ./package.json ./package-lock.json /app/
RUN npm ci

FROM build-env AS build
COPY ./babel.config.js ./vue.config.js ./.eslintrc.js /app/
COPY ./public /app/public/
COPY ./src /app/src/
RUN npm run build

FROM golang:latest AS binary
WORKDIR /build
COPY ./main.go ./main_release.go ./go.sum ./go.mod /build/
COPY ./internal /build/internal/
COPY ./pkg /build/pkg/
COPY --from=build /app/dist/ /build/dist/
RUN CGO_ENABLED=0 go build -tags release -ldflags="-w -s" -o asuwave

FROM alpine:latest
COPY --from=binary /build/asuwave /app/

CMD ["/app/asuwave"]
