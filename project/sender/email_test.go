package sender

import "testing"

func TestEmail_Send(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		email   *Email
		args    args
		wantErr bool
	}{
		{
			name:    "Test OK",
			email:   &Email{},
			args:    args{text: "Hello World"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email := &Email{}
			if err := email.Send(tt.args.text); (err != nil) != tt.wantErr {
				t.Errorf("Email.Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
