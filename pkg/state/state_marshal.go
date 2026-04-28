package state

import (
	"bytes"
	"encoding"
	"encoding/binary"
	"errors"
	"io"
	"strconv"

	"github.com/applejag/epic-wizard-firefly-gladiators/pkg/util"
)

var (
	ErrSaveInvalidFileMarker = errors.New("file is not a valid game save")
	ErrSaveUnexpectedField   = errors.New("unexpected game state field")
	ErrSaveTooManyFireflies  = errors.New("save contains too many fireflies")
	ErrSaveInvalidField      = errors.New("save file is invalid, perhaps corrupted")
)

type FieldError struct {
	Field string
	Err   error
}

func (err FieldError) Error() string {
	return err.Field + ": " + err.Err.Error()
}

func (err FieldError) Unwrap() error {
	return err.Err
}

type IndexedFieldError struct {
	Field string
	Index int
	Err   error
}

func (err IndexedFieldError) Error() string {
	return err.Field + "[" + strconv.Itoa(err.Index) + "]: " + err.Err.Error()
}

func (err IndexedFieldError) Unwrap() error {
	return err.Err
}

// Arbitrary number, mostly for future-proofing
const FileMarker byte = 0x37

type FieldState byte

const (
	FieldStateUnknown            FieldState = 0
	FieldStateFireflies          FieldState = 1
	FieldStateBattlesWonTotal    FieldState = 2
	FieldStateBattlesPlayedTotal FieldState = 3
	FieldStateMoney              FieldState = 4
)

var _ encoding.BinaryUnmarshaler = &GameState{}

// UnmarshalBinary implements [encoding.BinaryUnmarshaler].
func (g *GameState) UnmarshalBinary(data []byte) error {
	buf := bytes.NewBuffer(data)

	if marker, err := buf.ReadByte(); err != nil || marker != FileMarker {
		return ErrSaveInvalidFileMarker
	}

	for {
		next, err := buf.ReadByte()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		switch FieldState(next) {
		case FieldStateBattlesPlayedTotal:
			val, err := binary.ReadUvarint(buf)
			if err != nil {
				return ErrSaveInvalidField
			}
			g.BattlesPlayedTotal = int(val)
		case FieldStateBattlesWonTotal:
			val, err := binary.ReadUvarint(buf)
			if err != nil {
				return ErrSaveInvalidField
			}
			g.BattlesWonTotal = int(val)
		case FieldStateMoney:
			val, err := binary.ReadUvarint(buf)
			if err != nil {
				return ErrSaveInvalidField
			}
			g.Money = int(val)
		case FieldStateFireflies:
			val, err := binary.ReadUvarint(buf)
			if err != nil {
				return ErrSaveInvalidField
			}
			if int(val) > cap(g.Fireflies) {
				return ErrSaveTooManyFireflies
			}
			clear(g.Fireflies)
			g.Fireflies = g.Fireflies[:val]
			for i := range int(val) {
				if err := g.Fireflies[i].UnmarshalBinaryBuf(buf); err != nil {
					return err
				}
			}
		default:
			return ErrSaveUnexpectedField
		}
	}
}

func (g *GameState) WriteToBuf(b []byte) int {
	n := writeByte(b[0:], FileMarker)

	n += writeByte(b[n:], byte(FieldStateFireflies))
	n += binary.PutUvarint(b[n:], uint64(len(g.Fireflies)))
	for i := range g.Fireflies {
		n += g.Fireflies[i].WriteToBuf(b[n:])
	}

	n += writeByte(b[n:], byte(FieldStateBattlesWonTotal))
	n += binary.PutUvarint(b[n:], uint64(g.BattlesWonTotal))
	n += writeByte(b[n:], byte(FieldStateBattlesPlayedTotal))
	n += binary.PutUvarint(b[n:], uint64(g.BattlesPlayedTotal))
	n += writeByte(b[n:], byte(FieldStateMoney))
	n += binary.PutUvarint(b[n:], uint64(g.Money))
	return n
}

type FieldFirefly byte

const (
	FieldFireflyEOF           FieldFirefly = 0
	FieldFireflyID            FieldFirefly = 1
	FieldFireflyName          FieldFirefly = 2
	FieldFireflySpeed         FieldFirefly = 3
	FieldFireflyNimbleness    FieldFirefly = 4
	FieldFireflyBattlesPlayed FieldFirefly = 5
	FieldFireflyBattlesWon    FieldFirefly = 6
	FieldFireflyHat           FieldFirefly = 7
)

func (f *Firefly) UnmarshalBinaryBuf(buf *bytes.Buffer) error {
	for {
		next, err := buf.ReadByte()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		switch FieldFirefly(next) {
		case FieldFireflyID:
			val, err := binary.ReadUvarint(buf)
			if err != nil {
				return ErrSaveInvalidField
			}
			f.ID = int(val)
		case FieldFireflyName:
			val, err := binary.ReadUvarint(buf)
			if err != nil {
				return ErrSaveInvalidField
			}
			f.Name = util.Name(val)
		case FieldFireflySpeed:
			val, err := binary.ReadUvarint(buf)
			if err != nil {
				return ErrSaveInvalidField
			}
			f.Speed = int(val)
		case FieldFireflyNimbleness:
			val, err := binary.ReadUvarint(buf)
			if err != nil {
				return ErrSaveInvalidField
			}
			f.Nimbleness = int(val)
		case FieldFireflyBattlesPlayed:
			val, err := binary.ReadUvarint(buf)
			if err != nil {
				return ErrSaveInvalidField
			}
			f.BattlesPlayed = int(val)
		case FieldFireflyBattlesWon:
			val, err := binary.ReadUvarint(buf)
			if err != nil {
				return ErrSaveInvalidField
			}
			f.BattlesWon = int(val)
		case FieldFireflyHat:
			val, err := binary.ReadUvarint(buf)
			if err != nil {
				return ErrSaveInvalidField
			}
			f.Hat = int(val)
		case FieldFireflyEOF:
			return nil
		default:
			return ErrSaveUnexpectedField
		}
	}
}

func (f *Firefly) WriteToBuf(b []byte) int {
	n := writeByte(b[0:], byte(FieldFireflyID))
	n += binary.PutUvarint(b[n:], uint64(f.ID))
	n += writeByte(b[n:], byte(FieldFireflyName))
	n += binary.PutUvarint(b[n:], uint64(f.Name))
	n += writeByte(b[n:], byte(FieldFireflySpeed))
	n += binary.PutUvarint(b[n:], uint64(f.Speed))
	n += writeByte(b[n:], byte(FieldFireflyNimbleness))
	n += binary.PutUvarint(b[n:], uint64(f.Nimbleness))
	n += writeByte(b[n:], byte(FieldFireflyBattlesPlayed))
	n += binary.PutUvarint(b[n:], uint64(f.BattlesPlayed))
	n += writeByte(b[n:], byte(FieldFireflyBattlesWon))
	n += binary.PutUvarint(b[n:], uint64(f.BattlesWon))
	n += writeByte(b[n:], byte(FieldFireflyHat))
	n += binary.PutUvarint(b[n:], uint64(f.Hat))
	n += writeByte(b[n:], byte(FieldFireflyEOF))
	return n
}

func writeByte(buf []byte, b byte) int {
	buf[0] = b
	return 1
}
