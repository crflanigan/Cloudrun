package cloudrun

import (
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/internal"
	"github.com/influxdata/telegraf/plugins/serializers/influx"
	"github.com/influxdata/telegraf/testutil"
	"testing"
)

// default config used by Tests
func defaultCloudrun() *CloudRun {
	return &CloudRun{
		URL:          defaultURL,
		Timeout:      internal.Duration{Duration: defaultClientTimeout},
		Method:       defaultMethod,
		ConvertPaths: true,
	}
}

// TODO: This is may only be useful as a reference
func TestCloudRun_Close(t *testing.T) {
	cr := defaultCloudrun()

	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "close success", wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := cr.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TODO: This is may only be useful as a reference
func TestCloudRun_Connect(t *testing.T) {
	cr := defaultCloudrun()

	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "connect success", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := cr.Connect(); (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TODO: This may be the most main functionality to test. Many variations.
func TestCloudRun_Write(t *testing.T) {
	cr := defaultCloudrun()

	// http_test example
	//
	//tests := []struct {
	//	name       string
	//	plugin     *HTTP
	//	statusCode int
	//	errFunc    func(t *testing.T, err error)
	//}{
	//	{
	//		name: "success",
	//		plugin: &HTTP{
	//			URL: u.String(),
	//		},
	//		statusCode: http.StatusOK,
	//		errFunc: func(t *testing.T, err error) {
	//			require.NoError(t, err)
	//		},
	//	}
	//}
	//
	// kafka_test example
	//
	//tests := []struct {
	//  name   string
	//  plugin *Kafka
	//  input  []telegraf.Metric
	//  topic  string
	//  value  string
	// }{
	//	{
	//		name: "static topic",
	//		plugin: &Kafka{
	//			Brokers:      []string{"127.0.0.1"},
	//			Topic:        "telegraf",
	//			producerFunc: NewMockProducer,
	//		},
	//		input: []telegraf.Metric{
	//			testutil.MustMetric(
	//				"cpu",
	//				map[string]string{},
	//				map[string]interface{}{
	//					"time_idle": 42.0,
	//				},
	//				time.Unix(0, 0),
	//			),
	//		},
	//		topic: "telegraf",
	//		value: "cpu time_idle=42 0\n",
	//	}
	//}

	// TODO:
	tests := []struct {
		name    string
		metrics []telegraf.Metric
		wantErr bool
	}{
		{"success", testutil.MockMetrics(), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: Common pattern tt.plugin.SetSerializer(serializer) seen in other output tests
			// 	Is it a better approach? I think it would resolve the nil pointer dereference I'm getting
			serializer := influx.NewSerializer()

			cr.serializer = serializer

			if err := cr.Write(tt.metrics); (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}