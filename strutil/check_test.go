package strutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	assert.Equal(t, IsEmpty(""), true)
	assert.Equal(t, IsEmpty("a"), false)
}

func TestIsNumeric(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				s: "123",
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				s: "abc",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsNumeric(tt.args.s), "IsNumeric(%v)", tt.args.s)
		})
	}
}

func TestIsAlpha(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				str: "a",
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				str: "0",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAlpha(tt.args.str); got != tt.want {
				t.Errorf("IsAlpha() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAlphaNumeric(t *testing.T) {
	assert.Equal(t, IsAlphaNumeric("哈哈"), false)
	assert.Equal(t, IsAlphaNumeric("abc"), true)
	assert.Equal(t, IsAlphaNumeric("123"), true)
	assert.Equal(t, IsAlphaNumeric("ABc123"), true)
}

func TestIsAllUpperAlpha(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				str: "ABC",
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				str: "ABc",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAllUpperAlpha(tt.args.str); got != tt.want {
				t.Errorf("IsAllUpperAlpha() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAllLowerAlpha(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				str: "abc",
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				str: "ABc",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAllLowerAlpha(tt.args.str); got != tt.want {
				t.Errorf("IsAllLowerAlpha() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsContainAlpha(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				str: "ABc123",
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				str: "123",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsContainAlpha(tt.args.str); got != tt.want {
				t.Errorf("IsContainAlpha() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsContainUpperAlpha(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				str: "ABc123",
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				str: "ccc123",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsContainUpperAlpha(tt.args.str); got != tt.want {
				t.Errorf("IsContainUpperAlpha() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsContainLowerAlpha(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				str: "ABc123",
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				str: "ABC123",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsContainLowerAlpha(tt.args.str); got != tt.want {
				t.Errorf("IsContainLowerAlpha() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsContainChineseChar(t *testing.T) {
	assert.True(t, IsContainChinese("我爱中国"))
	assert.False(t, IsContainChinese("a"))
}

func TestIsFloatStr(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				str: "0.123",
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				str: "abc",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsFloatStr(tt.args.str), "IsFloatStr(%v)", tt.args.str)
		})
	}
}

func TestIsIntStr(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				str: "123",
			},
			want: true,
		},
		{
			name: "case true",
			args: args{
				str: "abc",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsIntStr(tt.args.str), "IsIntStr(%v)", tt.args.str)
		})
	}
}

func TestIsChinese(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				s: "我爱中国",
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				s: "abc",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsChinese(tt.args.s), "IsChinese(%v)", tt.args.s)
		})
	}
}

func TestIsContainChinese(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				str: "我爱中国",
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				str: "abc",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsContainChinese(tt.args.str), "IsContainChinese(%v)", tt.args.str)
		})
	}
}

func TestIsBase64(t *testing.T) {
	type args struct {
		base64 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				base64: "YWJj",
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				base64: "abc",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsBase64(tt.args.base64), "IsBase64(%v)", tt.args.base64)
		})
	}
}

func TestIsValidUtf8(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				s: "abc",
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				s: "\xFB\xBF\xBF\xBF\xBF",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsValidUtf8(tt.args.s), "IsValidUtf8(%v)", tt.args.s)
		})
	}
}

func TestIsRuneWord(t *testing.T) {
	type args struct {
		c rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				c: 'a',
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				c: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsRuneWord(tt.args.c), "IsRuneWord(%v)", tt.args.c)
		})
	}
}

func TestIsRuneLower(t *testing.T) {
	type args struct {
		c rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				c: 'a',
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				c: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsRuneLower(tt.args.c), "IsRuneLower(%v)", tt.args.c)
		})
	}
}

func TestIsRuneUpper(t *testing.T) {
	type args struct {
		c rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				c: 'A',
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				c: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsRuneUpper(tt.args.c), "IsRuneUpper(%v)", tt.args.c)
		})
	}
}
func TestNoCaseEq(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				s: "a",
				t: "A",
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				s: "a",
				t: "1",
			},
			want: false,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NoCaseEq(tt.args.s, tt.args.t), "NoCaseEq(%v, %v)", tt.args.s, tt.args.t)
		})
	}
}

func TestHasOneSub(t *testing.T) {
	type args struct {
		s    string
		subs []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				s:    "abcdef",
				subs: []string{"a", "b"},
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				s:    "abcdef",
				subs: []string{"i"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, HasOneSub(tt.args.s, tt.args.subs), "HasOneSub(%v, %v)", tt.args.s, tt.args.subs)
		})
	}
}

func TestHasAllSubs(t *testing.T) {
	type args struct {
		s    string
		subs []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				s:    "abc",
				subs: []string{"a", "b", "c"},
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				s:    "abc",
				subs: []string{"a", "b"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, HasAllSubs(tt.args.s, tt.args.subs), "HasAllSubs(%v, %v)", tt.args.s, tt.args.subs)
		})
	}
}

func TestHasOnePrefix(t *testing.T) {
	type args struct {
		s        string
		prefixes []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case true",
			args: args{
				s:        "abc",
				prefixes: []string{"a", "b"},
			},
			want: true,
		},
		{
			name: "case false",
			args: args{
				s:        "abc",
				prefixes: []string{"c"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, HasOnePrefix(tt.args.s, tt.args.prefixes), "HasOnePrefix(%v, %v)", tt.args.s, tt.args.prefixes)
		})
	}
}
