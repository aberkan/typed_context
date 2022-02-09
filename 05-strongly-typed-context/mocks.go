<<<<<<< HEAD

=======
>>>>>>> 906f617 (Add 5 examples & the linter)
package main

import (
	"context"
	"fmt"
)

// ================================
// Some mock implementations to support doing the thing
// ================================
func GetContextWithAllTheMocks() MockContext {
	return MockContext{
<<<<<<< HEAD
		Context: context.Background(),
		request: &Request{ key: "mockUser" },
		database: &Database{},
		httpClient: &HttpClient{},
		secrets: &Secrets{},
		logger: &Logger{},
=======
		Context:    context.Background(),
		request:    &Request{key: "mockUser"},
		database:   &Database{},
		httpClient: &HttpClient{},
		secrets:    &Secrets{},
		logger:     &Logger{},
>>>>>>> 906f617 (Add 5 examples & the linter)
	}
}

type MockContext struct {
	context.Context
<<<<<<< HEAD
	request *Request
	database *Database
	httpClient *HttpClient
	secrets *Secrets
	logger *Logger
=======
	request    *Request
	database   *Database
	httpClient *HttpClient
	secrets    *Secrets
	logger     *Logger
>>>>>>> 906f617 (Add 5 examples & the linter)
}

func (c MockContext) Request() *Request {
	return c.request
}

func (c MockContext) Database() DatabaseInterface {
	return c.database
}

func (c MockContext) HttpClient() *HttpClient {
	return c.httpClient
}

func (c MockContext) Secrets() *Secrets {
	return c.secrets
}

func (c MockContext) Logger() *Logger {
	return c.logger
}

type Request struct {
	key DatabaseKey
}

func (r *Request) GetUserKey() (DatabaseKey, error) {
	fmt.Printf("Request getting key %v\n", r.key)
	return r.key, nil
}

type Token string

func (r *Request) GetToken() (Token, error) {
	return "a Token", nil
}

type User struct {
	name string
}

func (user *User) GetName() string {
	return user.name
}
func (*User) CanDoThing(thing string) bool {
	return true
}

type DatabaseKey string

type Database struct{}

func (*Database) Read(
<<<<<<< HEAD
	ctx interface{
=======
	ctx interface {
>>>>>>> 906f617 (Add 5 examples & the linter)
		context.Context
		SecretsContext
		LoggerContext
	},
	key DatabaseKey,
) (*User, error) {
	fmt.Printf("Database Reading %v\n", string(key))
<<<<<<< HEAD
=======
	// Mark as used so the linter doesn't complain
	// _ = ctx.Secrets()
	_ = ctx.Logger()
	_ = ctx.(context.Context)
>>>>>>> 906f617 (Add 5 examples & the linter)
	return &User{name: string(key)}, nil
}

type Secrets struct{}

<<<<<<< HEAD
type HttpClient struct {}
=======
type HttpClient struct{}
>>>>>>> 906f617 (Add 5 examples & the linter)

func (*HttpClient) Post(
	ctx interface {
		context.Context
		RequestContext
	},
	url string,
	param string,
) error {
	fmt.Printf("HTTP Posting %v?%v\n", url, param)
<<<<<<< HEAD
	return nil
}

type Logger struct {}
=======
	// Mark as used so the linter doesn't complain
	_ = ctx.Request()
	_ = ctx.(context.Context)
	return nil
}

type Logger struct{}
>>>>>>> 906f617 (Add 5 examples & the linter)
