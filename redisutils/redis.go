package redisutils

import (
	"context"
	"fmt"
	"runtime"

	"github.com/go-redis/redis/v8"
)

const (
	Simple   = "simple"
	Failover = "failover"
	Cluster  = "cluster"
	Ring     = "ring"
)

func NewRedis(address []string, opts ...Option) (redis.UniversalClient, error) {
	o := &Options{
		ctx:          context.Background(),
		addresses:    address,
		shards:       nil,
		poolSize:     10 * runtime.NumCPU(),
		masterName:   "",
		db:           0,
		username:     "",
		password:     "",
		clientType:   Simple,
		dialTimeout:  0,
		readTimeout:  0,
		writeTimeout: 0,
	}
	for _, opt := range opts {
		opt(o)
	}
	return NewRedisWithOptions(o)
}

func NewRedisWithOptions(o *Options) (redis.UniversalClient, error) {
	switch o.clientType {
	default:
		fallthrough
	case Simple:
		return newSimpleRedis(o)
	case Failover:
		return newFailoverRedis(o)
	case Cluster:
		return newClusterRedis(o)
	case Ring:
		return newRingRedis(o)
	}
}

func NewSimpleRedis(opts ...Option) (*redis.Client, error) {
	o := &Options{
		poolSize: 10 * runtime.NumCPU(),
		db:       0,
	}
	for _, opt := range opts {
		opt(o)
	}
	return newSimpleRedis(o)
}

func newSimpleRedis(o *Options) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:            o.addresses[0],
		Username:        o.username,
		Password:        o.password,
		DB:              o.db,
		PoolSize:        o.poolSize,
		DialTimeout:     o.dialTimeout,
		ReadTimeout:     o.readTimeout,
		WriteTimeout:    o.writeTimeout,
		MaxRetries:      o.maxRetries,
		MinRetryBackoff: o.minRetryBackoff,
		MaxRetryBackoff: o.maxRetryBackoff,
	})
	_, err := client.Ping(o.ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed ping redis %v, %w", o.addresses, err)
	}
	return client, nil
}

func NewFailoverRedis(opts ...Option) (*redis.Client, error) {
	o := &Options{
		poolSize: 10 * runtime.NumCPU(),
		db:       0,
	}
	for _, opt := range opts {
		opt(o)
	}
	return newFailoverRedis(o)
}

func newFailoverRedis(o *Options) (*redis.Client, error) {
	failoverClient := redis.NewFailoverClient(&redis.FailoverOptions{
		SentinelAddrs:   o.addresses,
		MasterName:      o.masterName,
		DB:              o.db,
		Username:        o.username,
		Password:        o.password,
		PoolSize:        o.poolSize,
		DialTimeout:     o.dialTimeout,
		ReadTimeout:     o.readTimeout,
		WriteTimeout:    o.writeTimeout,
		MaxRetries:      o.maxRetries,
		MinRetryBackoff: o.minRetryBackoff,
		MaxRetryBackoff: o.maxRetryBackoff,
	})
	_, err := failoverClient.Ping(o.ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed ping redis %v, %w", o.addresses, err)
	}
	return failoverClient, nil
}

func NewClusterRedis(opts ...Option) (*redis.ClusterClient, error) {
	o := &Options{
		poolSize: 10 * runtime.NumCPU(),
		db:       0,
	}
	for _, opt := range opts {
		opt(o)
	}
	return newClusterRedis(o)
}

func newClusterRedis(o *Options) (*redis.ClusterClient, error) {
	clusterClient := redis.NewClusterClient(&redis.ClusterOptions{
		Username:        o.username,
		Password:        o.password,
		Addrs:           o.addresses,
		PoolSize:        o.poolSize,
		DialTimeout:     o.dialTimeout,
		ReadTimeout:     o.readTimeout,
		WriteTimeout:    o.writeTimeout,
		MaxRetries:      o.maxRetries,
		MinRetryBackoff: o.minRetryBackoff,
		MaxRetryBackoff: o.maxRetryBackoff,
	})
	_, err := clusterClient.Ping(o.ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed ping redis %v, %w", o.addresses, err)
	}
	return clusterClient, nil
}

func NewRingRedis(opts ...Option) (*redis.Ring, error) {
	o := &Options{
		poolSize: 10 * runtime.NumCPU(),
		db:       0,
	}
	for _, opt := range opts {
		opt(o)
	}
	return newRingRedis(o)
}

func newRingRedis(o *Options) (*redis.Ring, error) {
	addrs := make(map[string]string)
	for i := 0; i < len(o.addresses); i++ {
		addrs[o.shards[i]] = o.addresses[i]
	}
	ringClient := redis.NewRing(&redis.RingOptions{
		Addrs:           addrs,
		PoolSize:        o.poolSize,
		Username:        o.username,
		Password:        o.password,
		DialTimeout:     o.dialTimeout,
		ReadTimeout:     o.readTimeout,
		WriteTimeout:    o.writeTimeout,
		MaxRetries:      o.maxRetries,
		MinRetryBackoff: o.minRetryBackoff,
		MaxRetryBackoff: o.maxRetryBackoff,
	})
	_, err := ringClient.Ping(o.ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed ping redis %v, %w", o.addresses, err)
	}
	return ringClient, nil
}
