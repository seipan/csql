package csql

import "testing"

func TestYmlFileReadYmlFile(t *testing.T) {
	t.Skip("Skip YmlFile.ReadYmlFile() test")
	cases := []struct {
		path string
		want string
	}{
		{
			path: "test.yml",
			want: "mysql://root:root@tcp(localhost:3306)/test",
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.path, func(t *testing.T) {
			t.Parallel()
			y := YmlFile{path: c.path}
			got, err := y.ReadYmlFile()
			if err != nil {
				t.Errorf("YmlFile.ReadYmlFile() error = %v", err)
				return
			}
			if got.(Yml).dns != c.want {
				t.Errorf("YmlFile.ReadYmlFile() = %v, want %v", got, c.want)
			}
		})
	}
}
