package github

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGithubTimestamp_UnmarshalJSON(t *testing.T) {
	assert := require.New(t)
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		i    int64
		args args
	}{
		{
			name: "should parse unix timestamp",
			args: args{b: []byte("1554369722")},
			i:    time.Unix(1554369722, 0).Unix(),
		},
		{
			name: "should parse RFC3339 timestamp",
			args: args{b: []byte("\"2019-04-04T09:22:02.000Z\"")},
			i:    time.Unix(1554369722, 0).Unix(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var i GithubTimestamp
			err := i.UnmarshalJSON(tt.args.b)
			assert.NoError(err)
			assert.Equal(tt.i, i.Int64())
		})
	}
}

func TestGithubTimestampDoc_UnmarshalJSON(t *testing.T) {
	assert := require.New(t)

	type doc struct {
		CreatedAt GithubTimestamp `json:"created_at,omitempty"`
	}

	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		i    int64
		args args
	}{
		{
			name: "should parse unix timestamp",
			args: args{b: []byte(`{"created_at": 1554369722}`)},
			i:    time.Unix(1554369722, 0).Unix(),
		},
		{
			name: "should parse RFC3339 timestamp",
			args: args{b: []byte(`{"created_at": "2019-04-04T09:22:02.000Z"}`)},
			i:    time.Unix(1554369722, 0).Unix(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var d doc
			err := json.Unmarshal(tt.args.b, &d)
			assert.NoError(err)
			assert.Equal(tt.i, d.CreatedAt.Int64())
		})
	}
}
