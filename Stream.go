package raptor

import (
	"fmt"

	"github.com/raptorbox/raptor-sdk-go/models"
)

//CreateStream instantiate a new API client
func CreateStream(r *Raptor) *Stream {
	return &Stream{
		Raptor: r,
	}
}

//Stream API client
type Stream struct {
	Raptor *Raptor
}

//GetConfig return the configuration
func (i *Stream) GetConfig() models.Config {
	return i.Raptor.GetConfig()
}

//GetClient return a client instance
func (i *Stream) GetClient() models.Client {
	return i.Raptor.GetClient()
}

//List stored data for a stream
func (i *Stream) List(stream *models.Stream) (*[]models.Record, error) {

	raw, err := i.GetClient().Get(fmt.Sprintf(STREAM_LIST, stream.GetDevice().ID, stream.Name), nil)
	if err != nil {
		return nil, err
	}

	res := &[]models.Record{}
	err = i.GetClient().FromJSON(raw, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

//Search stored data for a stream
func (i *Stream) Search(q *models.DataQuery) (*[]models.Record, error) {

	raw, err := i.GetClient().Get(fmt.Sprintf(STREAM_LIST, stream.GetDevice().ID, stream.Name), nil)
	if err != nil {
		return nil, err
	}

	res := &[]models.Record{}
	err = i.GetClient().FromJSON(raw, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
