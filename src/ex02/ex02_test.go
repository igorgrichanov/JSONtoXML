package ex02

import (
	"bufio"
	"os"
	"reflect"
	"sort"
	"testing"
)

var f1Path = "../snapshots/snapshot1.txt"
var f2Path = "../snapshots/snapshot2.txt"
var f3Path = "../snapshots/snapshot3.txt"
var f4Path = "../snapshots/snapshot4.txt"

func Test_compare(t *testing.T) {
	old, _ := os.Open(f1Path)
	defer old.Close()
	new, _ := os.Open(f2Path)
	defer new.Close()
	s1 := bufio.NewScanner(old)
	s2 := bufio.NewScanner(new)

	f3, _ := os.Open(f3Path)
	defer f3.Close()
	f4, _ := os.Open(f4Path)
	defer f4.Close()
	s3 := bufio.NewScanner(f3)
	s4 := bufio.NewScanner(f4)

	type args struct {
		s1 *bufio.Scanner
		s2 *bufio.Scanner
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 []string
	}{
		{
			name: "test 1",
			args: args{
				s1: s1,
				s2: s2,
			},
			want: []string{
				"/Users/baker/recipes/database.xml",
				"/etc/stove/config.xml",
				"/var/log/orders.log",
			},
			want1: []string{
				"/Users/igor/otchet.docx",
				"/etc/homebrew/LICENSE.txt",
				"/var/log/payments.log",
			},
		},
		{
			name: "test 2",
			args: args{
				s1: s3,
				s2: s4,
			},
			want:  []string{"/Users/baker/pokemon.avi", "/var/log/orders.log"},
			want1: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := compare(tt.args.s1, tt.args.s2)
			sort.Slice(got, func(i, j int) bool { return got[i] < got[j] })
			sort.Slice(got1, func(i, j int) bool { return got1[i] < got1[j] })
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("compare() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("compare() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
