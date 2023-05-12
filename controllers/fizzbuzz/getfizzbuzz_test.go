package fizzbuzz

import (
	"reflect"
	"testing"
)

func Test_getFizzBuzz(t *testing.T) {
	type args struct {
		params fizzbuzzparams
	}
	tests := []struct {
		name    string
		args    args
		wantRes []string
		wantErr bool
	}{
		{
			name: "test KO, int1 is not an integer",
			args: args{
				params: fizzbuzzparams{
					int1: "abc",
				},
			},
			wantRes: nil,
			wantErr: true,
		},
		{
			name: "test KO, int2 is not an integer",
			args: args{
				params: fizzbuzzparams{
					int1: "3",
					int2: "abc",
				},
			},
			wantRes: nil,
			wantErr: true,
		},
		{
			name: "test KO, limit is not an integer",
			args: args{
				params: fizzbuzzparams{
					int1:  "3",
					int2:  "5",
					limit: "abc",
				},
			},
			wantRes: nil,
			wantErr: true,
		},
		{
			name: "test ok, fizzbuzz with int1 = 3 and int2 = 5, limit 100",
			args: args{
				params: fizzbuzzparams{
					string1: "fizz",
					string2: "buzz",
					int1:    "3",
					int2:    "5",
					limit:   "100",
				},
			},
			wantRes: []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz", "16", "17", "fizz", "19", "buzz", "fizz", "22", "23", "fizz", "buzz", "26", "fizz", "28", "29", "fizzbuzz", "31", "32", "fizz", "34", "buzz", "fizz", "37", "38", "fizz", "buzz", "41", "fizz", "43", "44", "fizzbuzz", "46", "47", "fizz", "49", "buzz", "fizz", "52", "53", "fizz", "buzz", "56", "fizz", "58", "59", "fizzbuzz", "61", "62", "fizz", "64", "buzz", "fizz", "67", "68", "fizz", "buzz", "71", "fizz", "73", "74", "fizzbuzz", "76", "77", "fizz", "79", "buzz", "fizz", "82", "83", "fizz", "buzz", "86", "fizz", "88", "89", "fizzbuzz", "91", "92", "fizz", "94", "buzz", "fizz", "97", "98", "fizz", "buzz"},
			wantErr: false,
		},
		{
			name: "test ok, tototata with int1 = 2 and int2 = 7, limit 20",
			args: args{
				params: fizzbuzzparams{
					string1: "toto",
					string2: "tata",
					int1:    "2",
					int2:    "7",
					limit:   "50",
				},
			},
			wantRes: []string{"1", "toto", "3", "toto", "5", "toto", "tata", "toto", "9", "toto", "11", "toto", "13", "tototata", "15", "toto", "17", "toto", "19", "toto", "tata", "toto", "23", "toto", "25", "toto", "27", "tototata", "29", "toto", "31", "toto", "33", "toto", "tata", "toto", "37", "toto", "39", "toto", "41", "tototata", "43", "toto", "45", "toto", "47", "toto", "tata", "toto"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := getFizzBuzz(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("getFizzBuzz() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("getFizzBuzz() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
