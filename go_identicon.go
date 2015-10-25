package identicon

import (
    "fmt"
    "crypto/sha256"

    "image"
    "image/png"
    "image/color"

    "encoding/base64"
    "encoding/binary"

    "bytes"
)

// data type that contains the identicons properties
// identicon type needs to satisfy the IconCreator interface, 
//      ie needs to implement the Create method,
//      which creates an image and sets the pixel values
type identicon struct {
    dims    int       // scales the image, depending on rows and columns values
    rows    int       // how many rows of pixels    (y axis)
    columns int       // how many columns of pixels (x axis) :
                      //      - if dims is smaller than both rows and columns,
                      //        the identicon shape is a square
                      //      - otherwise it is a rectangle and the dims value scales
                      //        the identicon
    key     string    // string provided as a parameter to the IconCreator constructor
    hash    [32]byte  // sha256 checksum for the provided string
}

// interface (ie method) to implement, so that a data structure is of type IconCreator 
//      Create method signature: 
//                          input:   none
//                          returns: base64 encoded string
type IconCreator interface {
    Create() string
}

func New(key string) IconCreator {

    return &identicon{      // struct literal. Returning a pointer,
                            //      so that the struct properties can be changed
        dims:    200,       // scales the image, depending on rows and columns values
        rows:    400,       // how many rows of pixels    (y axis)
        columns: 300,       // how many columns of pixels (x axis)
        hash:    sha256.Sum256([]byte(key)),
        key:     key,
    }
}

// identicon needs to implement Create method of the given signature 
// TODO: return []byte (buf.Bytes)
func (icon *identicon) Create() string {

    // TODO: create a method to determine the endianness of the machine
    v := binary.LittleEndian.Uint32(icon.hash[:])       // casts [32]uint8 to uint32

    // 
    rgba := color.RGBA{
		R: uint8(v),            // casts uint32 to uint8, reducing 0-4294967296 to 0-255 
		G: uint8(v >> 7),       // v/2^7   => cast to uint8
		B: uint8(v >> 14),      // v/2^14  => cast to uint8
		A: 0xff,
	}

    fmt.Println(rgba)

    // create a new image with dimensions icon.dims x icon.dims
    idimage := image.NewRGBA(image.Rectangle{image.Point{0,0},image.Point{icon.dims,icon.dims}})

    // set each pixel value, based on:
    //          a) function x*y (determines the image pattern)
    //          b) on the email (using sha256 hash algorithm and the corresponding checksum)
    //          c) RGBA colour setup
    //
    //          a) determines the image pattern
    //          b) and c) determine the image colour distribution
    for x := 0; x < icon.columns; x++ {                 // x axis
        for y := 0; y < (icon.rows); y++ {              // y axis
            coef := uint8(x*y/icon.columns*icon.rows)   // function governing the image pattern
            c := color.RGBA{rgba.R*coef, rgba.G*coef, rgba.B*coef, rgba.A}
            idimage.Set(x,y,c)
        }
    }

    // 
    var buf bytes.Buffer
	png.Encode(&buf, idimage)       // if buf.Bytes() are returned => set content header 'image/png'
                                    // then write w.Write(buf.Bytes())

    imgBase64Str := base64.StdEncoding.EncodeToString(buf.Bytes())
    return imgBase64Str
}

