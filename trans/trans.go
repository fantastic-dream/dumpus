package trans

import (
	"github.com/pkg/errors"
	"net"
	"os"
	"strconv"
	"time"
	"github.com/pengcheng789/dumpus/agent"
	"github.com/pengcheng789/dumpus/conf"
	"github.com/pengcheng789/dumpus/log"
)

var logger = log.Logger

func Init(sig chan os.Signal) {
	logger.Info("Starting to send info.")

	go SendAgentInfoLoop(sig)

	logger.Info("Started sending info.")
}

// 发送数据至目标服务器。
func SendInfo(dstIpAddr string, port int, data []byte) error {
	logger.Debug("Call sending info.")
	address := dstIpAddr + ":" + strconv.Itoa(port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return errors.Wrap(err, "Connect the server failure.")
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			logger.WithError(err).Error("Close connection failure.")
		}
	}()

	_, err = conn.Write(data)
	if err != nil {
		return errors.Wrap(err, "Write data to connection failure.")
	}

	return nil
}

// 循环进行信息采集发送，直到退出信号触发。
func SendAgentInfoLoop(signal chan os.Signal) {
	config := conf.GetConfig()
	serverIpAddr := config.Server.IpAddr
	port := config.Server.Port

	for {
		select {
		case <-signal:
			logger.Info("Stop to send info.")
			return
		default:
			go func() {
				info := agent.Info()
				err := SendInfo(serverIpAddr, port, []byte(info.String()))
				if err != nil {
					logger.WithField("info", info).WithError(err).Error("Send info failure.")
				}
				logger.Debug(info)
			}()
			time.Sleep(time.Second)
		}
	}
}
