FROM node:lts as build-env
WORKDIR /app
COPY ./package.json ./package-lock.json /app/
RUN npm ci

FROM build-env as build
COPY ./public ./src ./babel.config.js ./vue.config.js /app/
RUN npm run build

FROM golang:latest
WORKDIR /build
COPY ./main.go ./go.sum ./go.mod ./datautil ./option ./serial ./server ./variable /build/
RUN go build -ldflags="-w -s" -o asuwave
COPY --from=build /app/dist/ /build/dist/

CMD ./asuwave -p $PORT
