package raptor

import (
	"errors"
	"fmt"
	"strconv"

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
func (s *Stream) GetConfig() models.Config {
	return s.Raptor.GetConfig()
}

//GetClient return a client instance
func (s *Stream) GetClient() models.Client {
	return s.Raptor.GetClient()
}

//Pull stored data for a stream
func (s *Stream) Pull(stream *models.Stream, offset int, size int) (*[]models.Record, error) {
	pager := "?" + strconv.Itoa(offset) + "&size=" + strconv.Itoa(size)
	raw, err := s.GetClient().Get(fmt.Sprintf(STREAM_LIST, stream.GetDevice().ID, stream.Name)+pager, nil)
	if err != nil {
		return nil, err
	}

	res := &[]models.Record{}
	err = s.GetClient().FromJSON(raw, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

//Search stored data for a stream
func (s *Stream) Search(stream *models.Stream, q *models.DataQuery) (*[]models.Record, error) {

	raw, err := s.GetClient().Post(fmt.Sprintf(STREAM_LIST, stream.GetDevice().ID, stream.Name), q, nil)
	if err != nil {
		return nil, err
	}

	res := &[]models.Record{}
	err = s.GetClient().FromJSON(raw, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

//Push data to the backend
func (s *Stream) Push(r *models.Record) error {

	stream := r.GetStream()
	if stream == nil {
		return errors.New("record stream is required, use Stream.CreateRecord")
	}

	_, err := s.GetClient().Put(fmt.Sprintf(STREAM_PUSH, stream.GetDevice().ID, stream.Name), r, nil)
	return err
}
