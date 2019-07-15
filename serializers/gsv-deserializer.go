package serializers

import (
	"bufio"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/halorium/networks-example/flaw"
)

type Gsv struct {
	Separator              string
	StrictFieldMapping     bool
	StripSurroundingQuotes bool
	UnescapeInternalQuotes bool

	fieldTranslations []int
	headerFieldCount  int
	headerFields      []string
	lineCount         int
	lineFieldCount    int
	lineFields        []string
	pointerValue      reflect.Value
	scanner           *bufio.Scanner
	sliceValue        reflect.Value
	source            io.ReadCloser
	structType        reflect.Type
	target            interface{}
	structValue       reflect.Value
}

func (gsv *Gsv) Deserializer(source io.ReadCloser, target interface{}) *flaw.Error {
	defer source.Close()

	gsv.source = source
	gsv.target = target

	return gsv.validate()
}

func (gsv *Gsv) validate() *flaw.Error {
	gsv.pointerValue = reflect.ValueOf(gsv.target)

	if gsv.pointerValue.Kind() != reflect.Ptr {
		return flaw.New(
			fmt.Sprint(
				"target must be a pointer, is ",
				gsv.pointerValue.Kind(),
			),
		).Wrap("cannot validate")
	}

	gsv.sliceValue = gsv.pointerValue.Elem()

	if gsv.sliceValue.Kind() != reflect.Slice {
		return flaw.New(
			fmt.Sprint(
				"target must be a pointer to a slice, is pointer to ",
				gsv.sliceValue.Kind(),
			),
		).Wrap("cannot validate")
	}

	gsv.structType = gsv.sliceValue.Type().Elem()

	if gsv.structType.Kind() != reflect.Struct {
		return flaw.New(
			fmt.Sprint(
				"target must be a pointer to a slice of structs, is pointer to slice of ",
				gsv.structType.Kind(),
			),
		).Wrap("cannot validate")
	}

	return gsv.processHeader()
}

func (gsv *Gsv) processHeader() *flaw.Error {
	gsv.scanner = bufio.NewScanner(gsv.source)

	return gsv.readHeader()
}

func (gsv *Gsv) readHeader() *flaw.Error {
	ok, flawError := gsv.readLine()

	if !ok {
		if flawError != nil {
			return flawError.Wrap("cannot readHeader")
		}

		return flaw.New("source is empty").Wrap("cannot readHeader")
	}

	gsv.headerFieldCount = gsv.lineFieldCount
	gsv.headerFields = gsv.lineFields

	return gsv.mapHeaderFieldsToStructFields()
}

func (gsv *Gsv) mapHeaderFieldsToStructFields() *flaw.Error {
	structFieldNamePositions := map[string]int{}

	for i := 0; i < gsv.structType.NumField(); i++ {
		field := gsv.structType.Field(i)

		tag := field.Tag.Get("gsv")

		if tag == "" {
			if gsv.StrictFieldMapping {
				return flaw.New("No gsv tag found for field: " + field.Name).Wrap("cannot mapHeaderFieldsToStructFields")
			}

			continue
		}

		structFieldNamePositions[tag] = i
	}

	if len(structFieldNamePositions) > gsv.headerFieldCount {
		gsv.fieldTranslations = make([]int, len(structFieldNamePositions))
	} else {
		gsv.fieldTranslations = make([]int, gsv.headerFieldCount)
	}

	for i, tag := range gsv.headerFields {
		position, found := structFieldNamePositions[tag]

		if !found {
			if gsv.StrictFieldMapping {
				return flaw.New("input field " + tag + " not found in struct tags").Wrap("cannot mapHeaderFieldsToStructFields")
			}

			continue
		}

		gsv.fieldTranslations[i] = position
	}

	return gsv.processRecords()
}

func (gsv *Gsv) processRecords() *flaw.Error {
	for {
		ok, err := gsv.processRecord()

		if !ok {
			if err != nil {
				return flaw.From(err).Wrap("cannot processRecords")
			}

			return nil
		}
	}
}

func (gsv *Gsv) processRecord() (bool, *flaw.Error) {
	ok, flawError := gsv.readLine()

	if !ok {
		if flawError != nil {
			return ok, flawError.Wrap("cannot processRecord")
		}

		return ok, nil
	}

	return ok, gsv.validateLine()
}

func (gsv *Gsv) validateLine() *flaw.Error {
	if gsv.lineFieldCount != gsv.headerFieldCount {
		return flaw.New(
			fmt.Sprint(
				"line ",
				gsv.lineCount,
				": line field count mismatch: got ",
				gsv.lineFieldCount,
				", expected ",
				gsv.headerFieldCount,
			),
		).Wrap("cannot validateLine")
	}

	gsv.stripSurroundingQuotes()

	return nil
}

func (gsv *Gsv) stripSurroundingQuotes() {
	if gsv.StripSurroundingQuotes {
		strippedFields := []string{}

		for _, value := range gsv.lineFields {
			value = strings.TrimPrefix(value, `"`)
			value = strings.TrimSuffix(value, `"`)

			strippedFields = append(strippedFields, value)
		}

		gsv.lineFields = strippedFields
	}

	gsv.unescapeInternalQuotes()
}

func (gsv *Gsv) unescapeInternalQuotes() {
	if gsv.UnescapeInternalQuotes {
		unescapedFields := []string{}

		for _, value := range gsv.lineFields {

			value = strings.Replace(value, `\"`, `"`, -1)

			unescapedFields = append(unescapedFields, value)
		}

		gsv.lineFields = unescapedFields
	}

	gsv.newStruct()
}

func (gsv *Gsv) newStruct() {
	gsv.structValue = reflect.New(gsv.structType)

	gsv.setFields()
}

func (gsv *Gsv) setFields() {
	for i, value := range gsv.lineFields {
		reflect.Indirect(gsv.structValue).Field(gsv.fieldTranslations[i]).SetString(value)
	}

	gsv.appendTarget()
}

func (gsv *Gsv) appendTarget() {
	gsv.sliceValue.Set(
		reflect.Append(gsv.sliceValue, reflect.Indirect(gsv.structValue)),
	)
}

func (gsv *Gsv) readLine() (bool, *flaw.Error) {
	ok := gsv.scanner.Scan()

	if !ok {
		err := gsv.scanner.Err()

		if err != nil {
			return ok, flaw.From(err).Wrap("cannot readLine")
		}

		return ok, nil
	}

	gsv.lineCount++

	gsv.lineFields = strings.Split(gsv.scanner.Text(), gsv.Separator)

	gsv.lineFieldCount = len(gsv.lineFields)

	return ok, nil
}
