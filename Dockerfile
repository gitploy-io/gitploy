# Build the server binary file.
FROM golang:1.17 AS server
ARG OSS=false

WORKDIR /server

COPY go.mod go.sum ./
RUN go mod download

COPY . . 
RUN if [ "${OSS}" = "false" ]; then \
        echo "Build the enterprise edition"; \
        go build -o gitploy-server ./cmd/server; \
    else \
        echo "Build the community edition"; \
        go build -o gitploy-server -tags "oss" ./cmd/server; \
    fi

# Build UI.
FROM node:14.17.0 AS ui
ARG OSS=false

WORKDIR /ui

ENV PATH /ui/node_modules/.bin:$PATH

COPY ./ui/package.json ./ui/package-lock.json ./
RUN npm install --silent

COPY ./ui ./
ENV REACT_APP_GITPLOY_OSS="${OSS}"
RUN npm run build

# Copy to the final image.
FROM golang:1.17-buster AS gitploy

WORKDIR /app

# Create DB
RUN mkdir /data

COPY --from=server --chown=root:root /server/LICENSE /server/NOTICE ./
COPY --from=server --chown=root:root /server/gitploy-server /go/bin/gitploy-server

# Copy UI output into the assets directory.
COPY --from=ui --chown=root:root /ui/build/ /app/

ENTRYPOINT [ "/go/bin/gitploy-server" ]