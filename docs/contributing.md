# Contributing

Please take a moment to review this document in order to make the contribution process easy and effective

## Support Requests

The [discussions](https://github.com/gitploy-io/gitploy/discussions) page is the preferred channel for support requests. Please do not use the issues for personal support requests.

## Feature Requests

Feature requests are always welcome. It is up to you to make a case to convince the developers of the merits of this feature. Please provide as much detail and context as possible.

## Pull Requests

Please discuss the feature on the issue before working on any significant pull requests. And pull requests should avoid containing unrelated commits.

Please check the checklist to increase the likelihood of your pull request being accepted on time:

* Run the unit tests.
* Include unit tests when you contribute a new feature.
* Include unit tests when you contribute a bug fix to prevent regressions.

## Development

### Server

1\. Install prerequsites:

* [golang](https://golang.org/dl/)+1.17
* [mockgen](https://github.com/golang/mock)@v1.6.0

2\. Set up the `.env` file at the root directory: 

```
GITPLOY_SERVER_HOST=localhost
GITPLOY_GITHUB_CLIENT_ID=XXXXXXXXXXXXXXX
GITPLOY_GITHUB_CLIENT_SECRET=XXXXXXXXXXXXX
GITPLOY_STORE_SOURCE=file:./sqlite3.db?cache=shared&_fk=1
```

Note that if you want to interact with GitHub in the local environment, you should install tunneling tools such as [ngork](https://ngrok.com/) and expose your local server.

3\. Run the server:

```
go run ./cmd/server
```

### UI

1\. Install prerequisites:

* [node](https://nodejs.org/ko/download/)+14.17.0

2\. Install dependencies

```
cd ./ui
npm install
```

3\. Set up the `.env` file at the `ui` directory:

```
REACT_APP_GITPLOY_TOKEN=YOUR_TOKEN
REACT_APP_GITPLOY_SERVER=http://localhost
```

4\. Run:

```
npm start
```
