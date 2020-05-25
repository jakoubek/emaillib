# emaillib

Package in Go for sending emails. Encapsulates [jordan-wright/email](https://github.com/jordan-wright/email). Probably not useful for anybody else than me.

## Usage

```go
emailer := email.NewClient()
_ := emailer.Sendmail("John Doe", "jd@example.com", "Subject", "Some message")
```

## Configuration

Create a new email object with the NewClient function. It takes a _variadic_ number of configuration functions as parameters

### Configure SMTP host and port

```go
emailer := email.NewClient(
  email.WithRelayhost("smtp.example.com", 465),
)
```

### Configure authentication

```go
emailer := email.NewClient(
  email.WithAuth("your.name@example.com", "123456789abcde"),
)
```

### Configure a default sender email address (and name)

```go
emailer := email.NewClient(
  email.WithSender("Example Inc. Customer Service", "info@example.com"),
)
```

### Configure development mode without real sending

```go
emailer := email.NewClient(
  email.WithDontSend(),
)
```

### Combine multiple configuration functions

All possible configuration functions can be combined:

```go
emailer := email.NewClient(
  email.WithRelayhost("smtp.fastmail.com", 465),
  email.WithAuth("your.name@example.com", "123456789abcde"),
  email.WithSender("Example Inc. Customer Service", "info@example.com"),
  email.WithDontSend(),
)
```


## Installation

You can install this package in your project by running:

```
go get -u github.com/jakoubek/emaillib
```
