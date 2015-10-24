package identicon

import (
    "fmt"
    "crypto/sha256"
    "image"
    "image/png"
    "image/color"
    "bytes"
    "encoding/base64"
    "encoding/binary"
)

// data type that contains
type identicon struct {
    dim     int
    rows    int
    columns int
    hash    [32]byte
    key     string
    itype   string       // png, jpeg
}

type IconCreator interface {
    Create() string
}

func New(key string) IconCreator {

    return &identicon{
        dim:     400,
        rows:    200,
        columns: 200,
        key:     key,
        hash:    sha256.Sum256([]byte(key)),
        itype:   "png",
    }
}

//func (icon *identicon) Create() []byte {
func (icon *identicon) Create() string {

    // TODO: create a method to determine the endianness of the machine
    v := binary.LittleEndian.Uint64(icon.hash[:])

    fmt.Printf("%t\n", v)

    nrgba := color.NRGBA{
		R: uint8(v),
		G: uint8(v >> 8),
		B: uint8(v >> 16),
		A: 0xff,
	}

    fmt.Println("----------------------------")
    fmt.Println(nrgba.R)
    fmt.Println(nrgba.G)
    fmt.Println(nrgba.B)
    fmt.Println(nrgba.A)
    fmt.Println("----------------------------")

    idimage := image.NewRGBA(image.Rectangle{image.Point{0,0},image.Point{icon.dim,icon.dim}})

    for x := 0; x < icon.rows; x++ {
        for y := 0; y < (icon.columns); y++ {
            coef := uint8(x*y/icon.rows*icon.columns)
            c := color.RGBA{nrgba.R*coef, nrgba.G*coef, nrgba.B*coef, nrgba.A}
            idimage.Set(x,y,c)
        }
    }

    var buf bytes.Buffer
	png.Encode(&buf, idimage)

    imgBase64Str := base64.StdEncoding.EncodeToString(buf.Bytes())

    return imgBase64Str
}

