package lib

import (
	"image"
	"reflect"
	"testing"
)

func TestCropByCenterAndSpoke(t *testing.T) {
	type args struct {
		rect   image.Rectangle
		center image.Point
		ratioX int
		ratioY int
		spoke  int
	}
	tests := []struct {
		name string
		args args
		want image.Rectangle
	}{
		{
			name: "ame",
			args: args{
				rect:   image.Rect(0, 0, 100, 100),
				center: image.Pt(50, 50),
				ratioX: 1,
				ratioY: 1,
				spoke:  0,
			},
			want: image.Rect(0, 0, 100, 100),
		},
		{
			name: "hmm",
			args: args{
				rect:   image.Rect(0, 0, 1400, 1970),
				center: image.Pt(1400/2, 1970/2),
				ratioX: 1,
				ratioY: 1,
				spoke:  1400 * 1970,
			},
			want: image.Rect(0, 285, 1400, 1685),
		},
		{
			name: "hmm222",
			args: args{
				rect:   image.Rect(0, 0, 1400, 1970),
				center: image.Pt(1400/2, 1970/2),
				ratioX: 1,
				ratioY: 2,
				spoke:  1400 * 1970,
			},
			want: image.Rect(207, 0, 1193, 1970),
		},

		{
			name: "ehe",
			args: args{
				rect:   image.Rect(0, 0, 300, 400),
				center: image.Pt(150, 200),
				ratioX: 1,
				ratioY: 1,
				spoke:  10,
			},
			want: image.Rect(145, 195, 155, 205),
		},
		{
			name: "spoke determine the size",
			args: args{
				rect:   image.Rect(0, 0, 300, 400),
				center: image.Pt(150, 200),
				ratioX: 1,
				ratioY: 1,
				spoke:  10,
			},
			want: image.Rect(145, 195, 155, 205),
		},
		{
			name: "spoke detection",
			args: args{
				rect:   image.Rect(0, 0, 300, 400),
				center: image.Pt(150, 200),
				ratioX: 3,
				ratioY: 1,
				spoke:  10000000,
			},
			want: image.Rect(0, 150, 300, 250),
		},
		{
			name: "spoke detection landscape, bounded by height",
			args: args{
				rect:   image.Rect(0, 0, 1600, 900),
				center: image.Pt(800, 450),
				ratioX: 1,
				ratioY: 3,
				spoke:  10000000,
			},
			want: image.Rect(650, 0, 950, 900),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CropByCenterAndSpoke(tt.args.rect, tt.args.center.X, tt.args.center.Y, tt.args.ratioX, tt.args.ratioY, tt.args.spoke); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CropByCenterAndSpoke() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cropFromPointByWidth(t *testing.T) {
	type args struct {
		ratioX   int
		ratioY   int
		size     int
		byHeight bool
		center   image.Point
	}
	tests := []struct {
		name string
		args args
		want image.Rectangle
	}{
		{
			name: "tst",
			args: args{
				ratioX:   1,
				ratioY:   1,
				size:     100,
				byHeight: false,
				center:   image.Pt(0, 0),
			},
			want: image.Rect(-50, -50, 50, 50),
		},
		{
			name: "tst",
			args: args{
				ratioX:   3,
				ratioY:   1,
				size:     300,
				byHeight: false,
				center:   image.Pt(0, 0),
			},
			want: image.Rect(-150, -50, 150, 50),
		},
		{
			name: "tst",
			args: args{
				ratioX:   3,
				ratioY:   1,
				size:     300,
				byHeight: true,
				center:   image.Pt(0, 0),
			},
			want: image.Rect(-450, -150, 450, 150),
		},
		{
			name: "tst",
			args: args{
				ratioX:   3,
				ratioY:   1,
				size:     300,
				byHeight: true,
				center:   image.Pt(500, 500),
			},
			want: image.Rect(50, 350, 950, 650),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cropFromPointByWidth(tt.args.ratioX, tt.args.ratioY, tt.args.size, tt.args.byHeight, tt.args.center.X, tt.args.center.Y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cropFromPointByWidth() = %v, want %v", got, tt.want)
			}
		})
	}
}
