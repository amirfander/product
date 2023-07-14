package broker

type BrokerConfiger interface {
	ConnectBroker(uri string)
}

func ConnectBroker(brokerConfig BrokerConfiger, uri string) {
	brokerConfig.ConnectBroker(uri)
}
