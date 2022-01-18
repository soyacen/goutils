package redisutils

import (
	"context"
	"time"
)

type Options struct {
	ctx             context.Context
	addresses       []string
	shards          []string
	poolSize        int
	masterName      string
	db              int
	username        string
	password        string
	clientType      string
	dialTimeout     time.Duration
	readTimeout     time.Duration
	writeTimeout    time.Duration
	maxRetries      int
	minRetryBackoff time.Duration
	maxRetryBackoff time.Duration
}

type Option func(*Options)

func Context(ctx context.Context) Option {
	return func(o *Options) {
		o.ctx = ctx
	}
}

func Shards(shards []string) Option {
	return func(o *Options) {
		o.shards = shards
	}
}

func PoolSize(size int) Option {
	return func(o *Options) {
		o.poolSize = size
	}
}

func MasterName(masterName string) Option {
	return func(o *Options) {
		o.masterName = masterName
	}
}

func DB(db int) Option {
	return func(o *Options) {
		o.db = db
	}
}

func Username(username string) Option {
	return func(o *Options) {
		o.username = username
	}
}

func Password(pwd string) Option {
	return func(o *Options) {
		o.password = pwd
	}
}

func DialTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.dialTimeout = timeout
	}
}

func ReadTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.readTimeout = timeout
	}
}

func WriteTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.writeTimeout = timeout
	}
}

func ClientType(clientType string) Option {
	return func(o *Options) {
		o.clientType = clientType
	}
}

func MaxRetries(maxRetries int) Option {
	return func(o *Options) {
		o.maxRetries = maxRetries
	}
}
func MinRetryBackoff(minRetryBackoff time.Duration) Option {
	return func(o *Options) {
		o.minRetryBackoff = minRetryBackoff
	}
}
func MaxRetryBackoff(maxRetryBackoff time.Duration) Option {
	return func(o *Options) {
		o.maxRetryBackoff = maxRetryBackoff
	}
}
