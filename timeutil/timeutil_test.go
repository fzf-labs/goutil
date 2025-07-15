package timeutil

import (
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNowSQLNullTime(t *testing.T) {
	assert.Equal(t, true, NowSQLNullTime().Valid)
}

func TestTimeToSQLNullTime(t *testing.T) {
	now := time.Now()
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want sql.NullTime
	}{
		{
			name: "case1",
			args: args{
				t: now,
			},
			want: sql.NullTime{
				Time:  now,
				Valid: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, TimeToSQLNullTime(tt.args.t), "TimeToSQLNullTime(%v)", tt.args.t)
		})
	}
}

func TestTimeToHuman(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				t: time.Now().Add(time.Minute),
			},
			want: "1分钟后",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, TimeToHuman(tt.args.t), "TimeToHuman(%v)", tt.args.t)
		})
	}
}

func TestTimeToHumanShow(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				t: time.Now(),
			},
			want: "刚刚发布",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, TimeToHumanShow(tt.args.t), "TimeToHumanShow(%v)", tt.args.t)
		})
	}
}

func TestRFC3339(t *testing.T) {
	type args struct {
		tt time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				tt: time.Now(),
			},
			want: time.Now().Format(time.RFC3339),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, RFC3339(tt.args.tt), "RFC3339(%v)", tt.args.tt)
		})
	}
}
