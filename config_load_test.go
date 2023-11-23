package hw_server_config

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "time"
)

func TestLoadYamlConfig(t *testing.T) {
    t.Run("valid config", func(t *testing.T) {
        s, err := LoadYamlConfig("./test_configs/config.yaml")
        assert.NoError(t, err)
        assert.Equal(t, "localhost:8080", s.Addr)
        assert.Equal(t, time.Millisecond*20, s.Timeout)
    })
    t.Run("invalid config path", func(t *testing.T) {
        _, err := LoadYamlConfig("./confi_1.yaml")
        assert.Error(t, err)
    })
    t.Run("invalid config format", func(t *testing.T) {
        _, err := LoadYamlConfig("./test_configs/config_wrong.yaml")
        assert.Error(t, err)
    })
}
