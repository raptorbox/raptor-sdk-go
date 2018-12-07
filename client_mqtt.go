package raptor

import (
	"net/url"
	"strconv"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/raptorbox/raptor-sdk-go/models"
)

// BrokerConnection track connection status to the broker
type BrokerConnection struct {
	connected bool
	client    mqtt.Client
}

func (c *DefaultClient) connectToBroker() (mqtt.Client, error) {

	mqttClient, err := c.GetBrokerClient()
	if err != nil {
		return nil, err
	}

	if !mqttClient.IsConnected() {
		if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
			return nil, token.Error()
		}
	}

	return mqttClient, nil
}

// GetBrokerClient return a MQTT client
func (c *DefaultClient) GetBrokerClient() (mqtt.Client, error) {

	if c.brokerConnection.client != nil {
		return c.brokerConnection.client, nil
	}

	u, err := url.Parse(c.GetConfig().GetURL())
	if err != nil {
		return c.brokerConnection.client, err
	}

	mqttURI := "tcp"
	mqttPort := "1883"
	if u.Scheme == "https" {
		mqttURI += "s"
		mqttPort = "8883"
	}
	mqttURI += "://" + u.Hostname() + ":" + mqttPort

	log.Debugf("MQTT uri %s", mqttURI)
	clientid := "raptorbox_mqttjs_ans_asdfghqwerty_go_" + strconv.Itoa(time.Now().Second())
	opts := mqtt.NewClientOptions().AddBroker(mqttURI).SetClientID(clientid)
	client := mqtt.NewClient(opts)
	c.brokerConnection.client = client

	return c.brokerConnection.client, nil
}

//Subscribe to topic
func (c *DefaultClient) Subscribe(topic string, cb func(event models.Payload)) error {

	mqttClient, err := c.connectToBroker()
	if err != nil {
		return err
	}

	if token := mqttClient.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {

		parts := strings.Split(topic, "/")

		var err error
		var p models.Payload

		switch parts[0] {
		case "tree":
			p = &models.TreeNodePayload{}
			return
		case "stream":
			p = &models.StreamPayload{}
		case "inventory":
			p = &models.DevicePayload{}
		default:
			//noop
		}

		err = FromJSON(msg.Payload(), p)
		if err == nil {
			cb(p)
			return
		}

		log.Errorf("Error handling message for `%s`: %s", parts[0], string(msg.Payload()))

	}); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

//Unsubscribe from topic
func (c *DefaultClient) Unsubscribe(topic string, cb func(event models.Payload)) error {

	mqttClient, err := c.GetBrokerClient()
	if err != nil {
		return err
	}

	if !mqttClient.IsConnected() {
		return nil
	}

	if token := mqttClient.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}
