package logger

import (
	"bytes"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
	conf "trypto-server/config"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 전역 로거
var lg *zap.Logger
var b bytes.Buffer

// 로거 초기화 컨피그 파라메터
func InitLogger(cfg *conf.Config) (err error) {
	cf := cfg.Log
	fmt.Println("cf", cf)
	now := time.Now()
	lPath := fmt.Sprintf("%s_%s.log", cf.Fpath, now.Format("2006-01-02"))
	// 설정 옵션
	writeSyncer := getLogWriter(lPath, cf.Msize, cf.Mbackup, cf.Mage)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(cf.Level))
	if err != nil {
		return
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)
	// lg 생성
	lg = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(lg)
	return
}

func Debug(ctx ...interface{}) {

	for _, str := range ctx {
		b.WriteString(str.(string))
		b.WriteString(" ")
	}

	lg.Debug("debug", zap.String("-", b.String()))
}

// Info is a convenient alias for Root().Info
func Info(ctx ...interface{}) {

	lg.Info("info", zap.String("-", b.String()))
}

// Warn is a convenient alias for Root().Warn
func Warn(ctx ...interface{}) {

	lg.Warn("warn", zap.String("-", b.String()))
}

// Error is a convenient alias for Root().Error
func Error(ctx ...interface{}) {

	lg.Error("error", zap.String("-", b.String()))
}

// encoder 옵션 설정
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,  // 로그파일 명 지정
		MaxSize:    maxSize,   // 로그 파일 최대 사이즈
		MaxBackups: maxBackup, // 최대 로그파일 수
		MaxAge:     maxAge,    // 로그파일 저장 일수
	}
	return zapcore.AddSync(lumberJackLogger)
}

// gin 로거 대체 설정
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		lg.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// gin 리커버리 대체 설정
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					lg.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					lg.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					lg.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
