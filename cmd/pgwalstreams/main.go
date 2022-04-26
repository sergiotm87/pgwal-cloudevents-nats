package main

import (
	"encoding/binary"
	"fmt"
	"os"

	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"github.com/jackc/pgx"
	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/sergiotm87/pgwalstreams/internal/listener"
)

var version = "0.1.0"

// logger log levels.
const (
	warningLoggerLevel = "warning"
	errorLoggerLevel   = "error"
	fatalLoggerLevel   = "fatal"
	infoLoggerLevel    = "info"
)

// getConf load config from file.
func getConf(path string) (*viper.Viper, error) {

	cfg := viper.New()

	cfg.SetConfigName("config") // name of config file (without extension)
	cfg.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	cfg.AddConfigPath("/conf/")   

	if err := cfg.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config: %w", err)
	}

	return cfg, nil
}

// initLogger init logrus preferences.
func initLogger(cfg *viper.Viper) {
	logrus.SetReportCaller(cfg.GetBool("caller"))
	if !cfg.GetBool("humanReadable") {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}

	var level logrus.Level

	switch cfg.GetString("level") {
	case warningLoggerLevel:
		level = logrus.WarnLevel
	case errorLoggerLevel:
		level = logrus.ErrorLevel
	case fatalLoggerLevel:
		level = logrus.FatalLevel
	case infoLoggerLevel:
		level = logrus.InfoLevel
	default:
		level = logrus.DebugLevel
	}

	logrus.SetLevel(level)
}


// initPgxConnections initialise db and replication connections.
func initPgxConnections(cfg *viper.Viper) (*pgx.Conn, *pgx.ReplicationConn, error) {
	pgxConf := pgx.ConnConfig{
		LogLevel: pgx.LogLevelInfo,
		Logger:   pgxLogger{},
		Host:     cfg.GetString("host"),
		Port:     uint16(cfg.GetInt("port")),
		Database: cfg.GetString("name"),
		User:     cfg.GetString("user"),
		Password: cfg.GetString("password"),
	}

	pgConn, err := pgx.Connect(pgxConf)
	if err != nil {
		return nil, nil, errors.Wrap(err, listener.ErrPostgresConnection)
	}

	rConnection, err := pgx.ReplicationConnect(pgxConf)
	if err != nil {
		return nil, nil, fmt.Errorf("%v: %w", listener.ErrReplicationConnection, err)
	}

	return pgConn, rConnection, nil
}

type pgxLogger struct{}

func (l pgxLogger) Log(level pgx.LogLevel, msg string, data map[string]interface{}) {
	logrus.Debugln(msg)
}

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "print only the version",
	}

	app := &cli.App{
		Name:    "pgwalstreams",
		Usage:   "listen postgres events",
		Version: version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Value:   "config.yml",
				Aliases: []string{"c"},
				Usage:   "path to config file",
			},
		},
		Action: func(c *cli.Context) error {
			cfg, err := getConf(c.String("config"))
			if err != nil {
				return fmt.Errorf("get config: %w", err)
			}

			initLogger(cfg.Sub("logger"))

			nc, err := nats.Connect(cfg.GetString("nats.address"))
			if err != nil {
				return fmt.Errorf("nats connection: %w", err)
			}

			// fmt.Println(cfg.Sub("database"))

			conn, rConn, err := initPgxConnections(cfg.Sub("database"))
			if err != nil {
				return fmt.Errorf("pgx connection: %w", err)
			}

			service := listener.NewWalListener(
				cfg,
				listener.NewRepository(conn),
				rConn,
				listener.NewNatsPublisher(*nc),
				listener.NewBinaryParser(binary.BigEndian),
			)

			if err := service.Process(c.Context); err != nil {
				return fmt.Errorf("service process: %w", err)
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
