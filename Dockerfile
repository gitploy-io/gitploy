# Build the server binary file.
FROM golang:1.15 AS server

WORKDIR /server

COPY go.mod go.sum ./
RUN go mod download

COPY . . 
RUN go install ./cmd/server

# Build UI.
FROM node:14.17.0 AS ui

WORKDIR /ui

ENV PATH /ui/node_modules/.bin:$PATH

COPY ./ui/package.json ./ui/package-lock.json ./
RUN npm install --silent

COPY ./ui ./
RUN npm run build

# Copy to the final image.
FROM golang:1.15-buster AS gitploy

WORKDIR /app

# Create DB
RUN mkdir /data

COPY --from=server --chown=root:root /server/LICENSE /server/NOTICE .
COPY --from=server --chown=root:root /go/bin/server /go/bin/server

# Copy UI output into the assets directory.
COPY --from=ui --chown=root:root /ui/build/ /app/

ENTRYPOINT [ "/go/bin/server" ]