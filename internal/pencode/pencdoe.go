package pencode

import (
	"encoding/binary"
	"io"
	"io/ioutil"

	"github.com/golang/protobuf/proto"

	"github.com/pkg/errors"
)

func Append(w io.Writer, tnums []uint64) error {

	// Unmarshal the struct
	buffer := NumberedList{Numbers: tnums}

	// Marshal the struct
	tbytes, err := proto.Marshal(&buffer)
	if err != nil {
		return errors.Wrap(err, "append - could not marshal struct")
	}

	// Write the 8 bytes length
	size := make([]byte, 8, 8)
	binary.LittleEndian.PutUint64(size, uint64(len(tbytes)))
	if n, err := w.Write(size); err != nil || n != 8 {
		return errors.Wrapf(err, "append - wrote %d bytes", n)
	}

	// Write the struct
	_, err = w.Write(tbytes)
	if err != nil {
		return errors.Wrap(err, "append - could not write to file")
	}

	return nil
}

func GetNumbers(r io.Reader) ([]uint64, error) {
	// Read the file
	tbytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, errors.Wrap(err, "read - could not read from reader")
	}

	// Prepare a struct
	tints := []uint64{}

	for len(tbytes) > 0 {

		// Read the first four bytes, and pop them off
		n := binary.LittleEndian.Uint64(tbytes[:8])
		tbytes = tbytes[8:]

		// Get the protobuffer
		pbuffer := NumberedList{}
		err := proto.Unmarshal(tbytes[:n], &pbuffer)
		if err != nil {
			return nil, errors.Wrapf(err, "could not read protocol buffer at index %d", n)
		}
		// Pop those bytes off
		tbytes = tbytes[n:]
		// Append the numbers
		tints = append(tints, pbuffer.Numbers...)
	}

	return tints, nil
}
