# mailer-api

[![Build Status](https://travis-ci.org/thecampagnards/mailer-api.svg?branch=master)](https://travis-ci.org/thecampagnards/mailer-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/thecampagnards/mailer-api)](https://goreportcard.com/report/github.com/thecampagnards/mailer-api)

This is a small REST service which send mails.

## Installation

Docker image available here <https://hub.docker.com/r/thecampagnards/mailer-api/>.
Run a mongo on localhost, then run the service with this command:

```sh
docker run -p 8080:8080 thecampagnards/mailer-api
```

You can define some env vars:

- ADMIN_USER: basic auth username (default admin)
- ADMIN_PASSWORD: basic auth password (default admin)
- MONGOHQ_URL: mongodb url (default localhost)

## Usage

You can add a server configuration by POSTing on this path `/configuration/smtp` a json like that <https://github.com/thecampagnards/mailer-api/blob/master/types/types.go#L18>.

You can check your configuration by GETing this path `/configuration/smtp`.

This is the same for the template. And the template can be a go template, the variables will be replace by the json send value.

To send a mail, you can use this path `http://localhost:8080/send?server-smtp-id=<ID>&template-id=<ID>&to=<MAILS>&cc=<MAILS>&cci=<MAILS>`.
Mails have to be split by `,`.`cc`, `cci` are not mandatory.
The body of the request will be used to replace the values of your mail template.

## Templating

The mail template can be based on go template [Check this link for more information](https://golang.org/pkg/text/template/).
All variables defined in your model will be replaced by the variable defined in the send request.
You can use HTML or plain text.
When you post your template you can minifying it (you can use [this tool](https://www.willpeavy.com/minifier/)) and then escape every json characters (you can use [this tool](https://www.freeformatter.com/json-escape.html)).
