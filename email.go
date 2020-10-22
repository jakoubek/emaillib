package emaillib

import (
	"fmt"
	"net/smtp"
	"strconv"

	emaillib "github.com/jordan-wright/email"
)

// ClientConfig is a type for the variadic functions to
// configure the Client object.
type ClientConfig func(*Client)

// Client is a struct for holding the configuration of the
// email client and a pointer to the email object from the library.
type Client struct {
	host     string
	port     int
	username string
	password string
	useAuth  bool
	from     string
	dontSend bool
	emailer  *emaillib.Email
}

// WithRelayhost configures the Client with a host name
// and a port to send the email to.
func WithRelayhost(host string, port int) ClientConfig {
	return func(c *Client) {
		c.host = host
		c.port = port
	}
}

// WithSender configures the Client with a sender
// name and email address.
func WithSender(name, email string) ClientConfig {
	return func(c *Client) {
		c.from = buildEmailAddress(name, email)
	}
}

// WithAuth configures the Client with a username
// and password for authentication with the SMTP server;
// sets the corresponding flag useAuth as well.
func WithAuth(username, password string, useAuth bool) ClientConfig {
	return func(c *Client) {
		c.username = username
		c.password = password
		c.useAuth = useAuth
	}
}

// WithDontSend configures the Client with the
// dontSend flag, i.e. there is no email sent.
func WithDontSend() ClientConfig {
	return func(c *Client) {
		c.dontSend = true
	}
}

// NewClient returns a Client object. The object was
// optionally configured with a variadic list of
// ClientConfig functions.
func NewClient(opts ...ClientConfig) *Client {
	client := Client{
		from:     "",
		useAuth:  false,
		dontSend: false,
		emailer:  emaillib.NewEmail(),
	}
	for _, opt := range opts {
		opt(&client)
	}
	client.emailer.From = client.from
	return &client
}

func (c *Client) NewMessage() {
	c.emailer = emaillib.NewEmail()
	c.emailer.From = c.from
}

// Debug returns a debug string with all settings in the Client.
func (c *Client) Debug() string {
	debug := "From     : " + c.from + "\n" + "Host/Port: " + c.host + ":" + strconv.Itoa(c.port) + "\n" + "Auth?    : " + strconv.FormatBool(c.useAuth) + "\n" + "Username : " + c.username
	debug += "\n-----------------------------------\n"
	debug += "Subject  : " + c.emailer.Subject
	debug += "\n-----------------------------------\n"
	debug += "TO:\n"
	for i, to := range c.emailer.To {
		debug += fmt.Sprintf("- (%d) %s\n", i, to)
	}
	debug += "CC:\n"
	for i, to := range c.emailer.Cc {
		debug += fmt.Sprintf("- (%d) %s\n", i, to)
	}

	debug += "==================================\n"
	return debug
}

// From sets the from address
func (c *Client) From(fromName, fromEmail string) {
	c.emailer.From = buildEmailAddress(fromName, fromEmail)
}

// To adds an email address to the list of recipient addresses (TO).
func (c *Client) To(toName, toEmail string) {
	c.emailer.To = append(c.emailer.To, buildEmailAddress(toName, toEmail))
}

// CC adds an email address to the list of CC addresses.
func (c *Client) CC(toName, toEmail string) {
	c.emailer.Cc = append(c.emailer.Cc, buildEmailAddress(toName, toEmail))
}

// Subject sets the email subject to the provided string.
func (c *Client) Subject(subject string) {
	c.emailer.Subject = subject
}

// BodyText sets the plaintext of the body
func (c *Client) BodyText(bodyText string) {
	c.emailer.Text = []byte(bodyText)
}

// BodyHTML sets the HTML part of the body
func (c *Client) BodyHTML(bodyHTML string) {
	c.emailer.HTML = []byte(bodyHTML)
}

// AttachFile attaches a file to the email Client.
func (c *Client) AttachFile(filename string) {
	c.emailer.AttachFile(filename)
}

// Send sends the prepared email message
func (c *Client) Send() error {
	return c.sendSMTPMessage()
}

// Sendmail sends an email with a given subject and body text to the
// provided email address. This is a shortcut to To, Subject, ....
func (c *Client) Sendmail(toName, toEmail, subject, message string) error {

	c.To(toName, toEmail)
	c.emailer.Subject = subject
	c.emailer.Text = []byte(message)

	return c.sendSMTPMessage()
}

// sendSMTPMessage is an internal method to finally send the email
// to the SMTP server.
func (c *Client) sendSMTPMessage() error {

	if c.dontSend {
		fmt.Println("!! DONT SEND !!")
		return nil
	}

	hostnameWithPort := fmt.Sprintf("%s:%d", c.host, c.port)
	if c.useAuth {
		return c.emailer.Send(hostnameWithPort,
			smtp.PlainAuth("", c.username, c.password, c.host),
		)
	}
	return c.emailer.Send(hostnameWithPort, nil)

}

// buildEmailAddress is an internal function to create an email
// address from a provided email address and an optional name.
func buildEmailAddress(name, email string) string {
	if name == "" {
		return email
	}
	return fmt.Sprintf("%s <%s>", name, email)
}
