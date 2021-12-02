package dingrobot

import (
	"crypto/tls"
	"net/http"
	"testing"
)

func TestRobot_SendText(t *testing.T) {
	type fields struct {
		webHook string
		secret  string
		hclient *http.Client
	}
	type args struct {
		content   string
		atMobiles []string
		isAtAll   bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				content: "测试消息",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRobot("https://oapi.dingtalk.com/robot/send?access_token=xxxxx")
			transport := &http.Transport{
				DisableKeepAlives: true, //短链接
				TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
			}
			client := &http.Client{Transport: transport}
			r.SetHTTPClient(client)

			if err := r.SendText(tt.args.content, tt.args.atMobiles, tt.args.isAtAll); (err != nil) != tt.wantErr {
				t.Errorf("Robot.SendText() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
