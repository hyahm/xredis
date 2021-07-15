package xredis

import (
	"crypto/tls"
	"errors"
	"net"
	"time"

	"github.com/gomodule/redigo/redis"
)

type GoRedis struct {
	// conf   *redis.Options
	client *redis.Pool
}

type Options struct {
	// Default is tcp.
	Network string
	// host:port address.
	Addr string

	// Dialer creates new network connection and has priority over
	// Network and Addr options.
	Dialer func() (net.Conn, error)

	// Hook that is called when new connection is established.
	// OnConnect func(*Conn) error

	// Optional password. Must match the password specified in the
	// requirepass server configuration option.
	Password string
	// Database to be selected after connecting to the server.
	DB int

	// Maximum number of retries before giving up.
	// Default is to not retry failed commands.
	MaxRetries int
	// Minimum backoff between each retry.
	// Default is 8 milliseconds; -1 disables backoff.
	MinRetryBackoff time.Duration
	// Maximum backoff between each retry.
	// Default is 512 milliseconds; -1 disables backoff.
	MaxRetryBackoff time.Duration
	ConnectTimeout  time.Duration
	// Dial timeout for establishing new connections.
	// Default is 5 seconds.
	DialTimeout time.Duration
	// Timeout for socket reads. If reached, commands will fail
	// with a timeout instead of blocking. Use value -1 for no timeout and 0 for default.
	// Default is 3 seconds.
	ReadTimeout time.Duration
	// Timeout for socket writes. If reached, commands will fail
	// with a timeout instead of blocking.
	// Default is ReadTimeout.
	WriteTimeout time.Duration
	KeepAlive    time.Duration
	// Maximum number of socket connections.
	// Default is 10 connections per every CPU as reported by runtime.NumCPU.
	PoolSize int
	// Minimum number of idle connections which is useful when establishing
	// new connection is slow.
	MinIdleConns int
	// Connection age at which client retires (closes) the connection.
	// Default is to not close aged connections.
	// Amount of time client waits for connection if all connections
	// are busy before returning an error.
	// Default is ReadTimeout + 1 second.
	PoolTimeout time.Duration
	// Amount of time after which client closes idle connections.
	// Should be less than server's timeout.
	// Default is 5 minutes. -1 disables idle timeout check.
	IdleTimeout time.Duration
	// Frequency of idle checks made by idle connections reaper.
	// Default is 1 minute. -1 disables idle connections reaper,
	// but idle connections are still discarded by the client
	// if IdleTimeout is set.
	IdleCheckFrequency time.Duration

	// Enables read only queries on slave nodes.
	readOnly bool

	// TLS Config to use. When set TLS will be negotiated.
	TLSConfig *tls.Config
}

func Conn(opt Options) (*GoRedis, error) {
	pool := redis.NewPool(func() (redis.Conn, error) {
		return redis.Dial("tcp",
			opt.Addr, redis.DialPassword(opt.Password),
			redis.DialDatabase(opt.DB),
			redis.DialKeepAlive(opt.KeepAlive),
			redis.DialReadTimeout(opt.ReadTimeout),
			redis.DialTLSConfig(opt.TLSConfig),
			redis.DialWriteTimeout(opt.WriteTimeout),
			redis.DialConnectTimeout(opt.ConnectTimeout),
		)
	}, opt.PoolSize)
	gr := &GoRedis{client: pool}
	return gr, gr.Ping()
}

func (gr *GoRedis) Ping() error {
	conn := gr.client.Get()
	return conn.Err()
}

// func (gr *GoRedis) Echo(message interface{}) (string, error) {
// 	if gr == nil {
// 		panic("please conn first")
// 	}
// 	if err := gr.Ping(); err != nil {
// 		return "", err
// 	}

// 	return gr.Echo(message).Result()
// }

func (gr *GoRedis) Quit() error {
	if gr.client == nil {
		return errors.New("please conn first")
	}
	return gr.client.Close()
}
