package config

import (
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

func New() (*Store, error) {
	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	location := path.Join(home, ".config", "tog_cli.yml")

	return &Store{
		location: location,
	}, nil
}

type Store struct {
	location string
}

func (s *Store) GetContext() (Context, error) {
	st, err := s.read()
	if err != nil {
		return Context{}, err
	}

	configContext := Context{Config: st.Config}

	stateContext, ok := st.Contexts[st.Host]
	if ok {
		configContext.AuthToken = stateContext.AuthToken
	}

	return configContext, nil
}

func (s *Store) GetConfig() (Config, error) {
	st, err := s.read()
	if err != nil {
		return Config{}, err
	}

	return st.Config, nil
}

func (s *Store) SetConfig(conf Config) error {
	st, err := s.read()
	if err != nil {
		return err
	}

	st.Config = conf
	return s.save(st)
}

func (s *Store) SetContext(host, authToken string, ttl time.Time) error {
	st, err := s.read()
	if err != nil {
		return err
	}

	stateContext, _ := st.Contexts[host]
	stateContext.Host = host
	stateContext.AuthToken = authToken
	stateContext.TTL = ttl

	if st.Contexts == nil {
		st.Contexts = map[string]context{}
	}
	st.Contexts[host] = stateContext

	return s.save(st)
}

func (s *Store) read() (state, error) {
	bytes, err := ioutil.ReadFile(s.location)
	if os.IsNotExist(err) {
		return state{}, nil
	} else if err != nil {
		return state{}, err
	}

	var st state
	if err := yaml.Unmarshal(bytes, &st); err != nil {
		return state{}, err
	}

	return st, nil
}

func (s *Store) save(st state) error {
	bytes, err := yaml.Marshal(st)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(s.location, bytes, os.ModePerm); err != nil {
		return err
	}

	return nil
}
