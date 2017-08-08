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
func (s *Stream) Pull(stream *models.Stream, offset int, size int) ([]models.Record, error) {
	pager := "?" + strconv.Itoa(offset) + "&size=" + strconv.Itoa(size)
	raw, err := s.GetClient().Get(fmt.Sprintf(STREAM_LIST, stream.GetDevice().ID, stream.Name)+pager, nil)
	if err != nil {
		return nil, err
	}

	res := make([]models.Record, 0)
	err = s.GetClient().FromJSON(raw, res)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(res); i++ {
		res[i].SetStream(stream)
	}

	return res, nil
}

//LastUpdate fetch the last record stored
func (s *Stream) LastUpdate(stream *models.Stream) (*models.Record, error) {

	raw, err := s.GetClient().Get(fmt.Sprintf(STREAM_LAST_UPDATE, stream.GetDevice().ID, stream.Name), nil)
	if err != nil {
		return nil, err
	}

	res := &models.Record{}
	err = s.GetClient().FromJSON(raw, res)
	if err != nil {
		return nil, err
	}

	res.SetStream(stream)

	return res, nil
}

//Search stored data for a stream
func (s *Stream) Search(stream *models.Stream, q *models.DataQuery) ([]models.Record, error) {

	raw, err := s.GetClient().Post(fmt.Sprintf(STREAM_LIST, stream.GetDevice().ID, stream.Name), q, nil)
	if err != nil {
		return nil, err
	}

	res := make([]models.Record, 0)
	err = s.GetClient().FromJSON(raw, &res)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(res); i++ {
		res[i].SetStream(stream)
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

//Delete drop the data of a stream
func (s *Stream) Delete(stream *models.Stream) error {
	return s.GetClient().Delete(fmt.Sprintf(STREAM_PUSH, stream.GetDevice().ID, stream.Name), nil)
}
