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

You can add a server configuration by POSTing on this path `/configuration/smtp` a json like that [see the object type](https://github.com/thecampagnards/mailer-api/blob/master/types/types.go#L27) :

```json
{
  "Host": "...",
  "Port": 587,
  "User": "...",
  "Password": "...",
  "From": "...",
  "InsecureSkipVerify": true
}
```

You can check your configuration by GETing this path `/configuration/smtp` and retrieve the id for the sending part.

## Templating

The mail template can be based on go template [Check this link for more information](https://golang.org/pkg/text/template/).
All variables defined in your model will be replaced by the variable defined in the send mail request.
You can use HTML, markdown (use `{{ "# my text" | markdown }}` in template) or plain text. [Here](https://github.com/wildbit/postmark-templates) you can find good HTML templates.
When you post your template you can minifying it (you can use [this tool](https://www.willpeavy.com/minifier/)) and then escape every json characters (you can use [this tool](https://www.freeformatter.com/json-escape.html)).

You can also use go template in the subject of the mail.

The template and the layout can be store somewhere else and can be retrieved via the url. You can't put both url and template.

You can add a template by POSTing on this path `/configuration/template` a json or yaml like that [see the object type](https://github.com/thecampagnards/mailer-api/blob/master/types/types.go#L14) :

```json
{
  "Subject": "...",
  "Template": "...", /* or */ "TemplateURL": "...",
  "LayoutIDs": ["...", "..."],
  "Variables": {
    "MyTemplateVar": {
      "type": "this is used as information",
      "description": "this is used as information"
    }
  },
  "Description": "this is used as information"
}
```

```yaml
subject: Example
description: "this is used as information"
template: |
    This is an example <a href="https://github.com/thecampagnards">github</a>
variables:
  myTemplateVar:
    description: "this is used as information"
    type: "this is used as information"
```

You can also use `Layout`, in this layout you can define blocks `{{define "title"}}Home{{end}}` then you can use it in your mail template `{{template "title" .}}`.

You can add a layout by POSTing on this path `/configuration/layout` a json or yaml like that [see the object type](https://github.com/thecampagnards/mailer-api/blob/master/types/types.go#L28) :

```json
{
  "Layout": "{{define \"title\"}}Home{{end}}",  /* or */ "LayoutURL": "",
  "Description": "this is used as information"
}
```

```yaml
layout: {{define "title"}}Home{{end}}
description: this is used as information
```

To use this layout in your template, you just have to append the layout id in the key `LayoutIDs` of your template.

## Send

You can send mail like that (the body can also be a yaml):

```bash
curl -X POST \
-H 'Content-Type: application/json' \
-d '{"MyTemplateVar":"...."}' \
http://localhost:8080/send?server-smtp-id=...&template-id=...&to=...&cc=...&cci=...
```

`server-smtp-id` and `template-id` can be retrieve by GETing this path `/configuration/smtp` and `/configuration/template`.
Mails have to be split by `,`.`to`, `cc`, `cci` are not mandatory.

### Send with attachements

You can add files to the sended mail like that :

```bash
curl -X POST \
-H 'Content-Type: multipart/form-data' \
-F 'attachments=@/path/to/fileX' \
-F 'attachments=@/path/to/fileY' \
... \
-F 'data={"MyTemplateVar":"...."}' \
http://localhost:8080/send?server-smtp-id=...&template-id=...&to=...&cc=...&cci=...
```
