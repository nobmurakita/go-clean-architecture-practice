package product

import (
	"testing"

	"go-clean-architecture-practice/go-pkg/ulid"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewProduct(t *testing.T) {
	ownerID := ulid.NewULID()
	type args struct {
		ownerID     string
		name        string
		description string
		price       int64
		stock       int
	}
	tests := []struct {
		name    string
		args    args
		want    *Product
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				ownerID:     ownerID,
				name:        "test",
				description: "test",
				price:       100,
				stock:       100,
			},
			want: &Product{
				ownerID:     ownerID,
				name:        "test",
				description: "test",
				price:       100,
				stock:       100,
			},
			wantErr: false,
		},
		{
			name: "異常系: オーナーIDが不正",
			args: args{
				ownerID:     "test",
				name:        "test",
				description: "test",
				price:       100,
				stock:       100,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "異常系: 商品説明が不正",
			args: args{
				ownerID:     ownerID,
				name:        "test",
				description: "",
				price:       100,
				stock:       100,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "異常系: 価格が不正",
			args: args{
				ownerID:     ownerID,
				name:        "test",
				description: "test",
				price:       0,
				stock:       100,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewProduct(tt.args.ownerID, tt.args.name, tt.args.description, tt.args.price, tt.args.stock)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			diff := cmp.Diff(
				got, tt.want,
				cmp.AllowUnexported(Product{}),
				cmpopts.IgnoreFields(Product{}, "id"),
			)
			if diff != "" {
				t.Errorf("NewProduct() = %v, want %v. error is %s", got, tt.want, err)
			}
		})
	}
}
