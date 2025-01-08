package db

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/wqh/smart/school/system/internal/configs"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"time"
)

/**
* description:
* author: wqh
* date: 2025/1/8
 */
func InitRedis(db *redis.Client, config configs.Config) {
	var err error
	switch config.Redis.ConnectType {
	case "common":
		db = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port),
			Password: config.Redis.Password, // 没有密码，默认值
			DB:       config.Redis.Database, // 默认DB 0
		})
		if err = db.Ping(context.Background()).Err(); err != nil {
			log.Fatalf("redis connect error %v", err)
		}
		break
	case "TLS":
		db = redis.NewClient(&redis.Options{
			TLSConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
				ServerName: "you domain",
			},
		})
		if err = db.Ping(context.Background()).Err(); err != nil {
			log.Fatalf("redis connect error %v", err)
		}
		break
	case "SSH":
		sshConfig := &ssh.ClientConfig{
			User:            "root",
			Auth:            []ssh.AuthMethod{ssh.Password("password")},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Timeout:         15 * time.Second,
		}

		sshClient, err := ssh.Dial("tcp", "remoteIP:22", sshConfig)
		if err != nil {
			panic(err)
		}

		db = redis.NewClient(&redis.Options{
			Addr: net.JoinHostPort("127.0.0.1", "6379"),
			Dialer: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return sshClient.Dial(network, addr)
			},
			// SSH不支持超时设置，在这里禁用
			ReadTimeout:  -1,
			WriteTimeout: -1,
		})
		if err = db.Ping(context.Background()).Err(); err != nil {
			log.Fatalf("redis connect error %v", err)
		}
		break
	default:
		log.Fatalf("redis connect error %v", err)
	}
}
