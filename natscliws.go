package natscliws

import (
	"net"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	nats "github.com/nats-io/nats.go"
	stan "github.com/nats-io/stan.go"
)

// New connects to nats via websocket dialer
func New(ws, n, cluster, client string) (*MQ, error) {
	opts := []nats.Option{
		nats.SetCustomDialer(&customDialer{ws}),
		nats.ReconnectWait(1 * time.Second),
		nats.DontRandomize(),
		nats.MaxReconnects(1<<31 - 1),
	}
	nc, err := nats.Connect(n, opts...)
	if err != nil {
		return nil, err
	}
	sc, err := stan.Connect(cluster, client, stan.NatsConn(nc))
	return &MQ{sc}, err
}

type customDialer struct {
	url string
}

func (cd *customDialer) Dial(network, address string) (net.Conn, error) {

	u, _ := url.Parse(cd.url)
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}
	return c.UnderlyingConn(), nil
}

type MQ struct {
	conn stan.Conn
}

func (mq *MQ) UnderlyingConn() stan.Conn {
	return mq.conn
}
