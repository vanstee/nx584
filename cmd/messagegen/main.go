package main

import (
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"go/format"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
)

const (
	messagesInputFilename    = "messages.csv"
	messageTemplateFilename  = "message.tmpl"
	messagesTemplateFilename = "messages.tmpl"
	messagesOutputFilename   = "messages.go"
)

var (
	messageHeaders     = []string{"Number", "Name", "Length", "Transition Capable", "Acknowledged"}
	messageDataHeaders = []string{"Byte", "Bit", "Byte Name", "Bit Name"}
	skipNames          = []string{"Do Not Use", "Reserved"}
)

type Message struct {
	Number            string
	Name              string
	Length            string
	TransitionCapable string
	Acknowledged      string
	DataByteFields    []*DataByteField
	DataBitFields     []*DataBitField
}

type DataByteField struct {
	Byte          string
	Name          string
	DataBitFields []*DataBitField
}

type DataBitField struct {
	Bit  string
	Name string
}

func main() {
	var in, templates, out string
	flag.StringVar(&in, "in", "data", "directory containing csv files used to generate message go code files")
	flag.StringVar(&templates, "templates", "templates", "directory containing template files used to generate message go code files")
	flag.StringVar(&out, "out", ".", "directory to which message go code files should be written")
	flag.Parse()

	messagesInputPath := path.Join(in, messagesInputFilename)
	_, err := os.Stat(messagesInputPath)
	if os.IsNotExist(err) {
		log.Fatalf("file not found at path %s", messagesInputPath)
	}
	if err != nil {
		log.Fatalf("failed to stat %s: %v", messagesInputPath, err)
	}

	messageTemplatePath := path.Join(templates, messageTemplateFilename)
	_, err = os.Stat(messageTemplatePath)
	if os.IsNotExist(err) {
		log.Fatalf("file not found at path %s", messageTemplatePath)
	}
	if err != nil {
		log.Fatalf("failed to stat %s: %v", messageTemplatePath, err)
	}

	messageTemplateBody, err := ioutil.ReadFile(messageTemplatePath)
	if err != nil {
		log.Fatalf("failed to read %s: %v", messageTemplatePath, err)
	}

	messageTemplate, err := template.New(messageTemplateFilename).Parse(string(messageTemplateBody))
	if err != nil {
		log.Fatalf("failed to parse template %s: %v", messageTemplatePath, err)
	}

	messagesTemplatePath := path.Join(templates, messagesTemplateFilename)
	_, err = os.Stat(messagesTemplatePath)
	if os.IsNotExist(err) {
		log.Fatalf("file not found at path %s", messagesTemplatePath)
	}
	if err != nil {
		log.Fatalf("failed to stat %s: %v", messagesTemplatePath, err)
	}

	messagesTemplateBody, err := ioutil.ReadFile(messagesTemplatePath)
	if err != nil {
		log.Fatalf("failed to read %s: %v", messagesTemplatePath, err)
	}

	messagesTemplate, err := template.New(messagesTemplateFilename).Parse(string(messagesTemplateBody))
	if err != nil {
		log.Fatalf("failed to parse template %s: %v", messagesTemplatePath, err)
	}

	stat, err := os.Stat(out)
	if os.IsNotExist(err) {
		log.Fatalf("directory not found at path %s", out)
	}
	if err != nil {
		log.Fatalf("failed to stat %s: %v", out, err)
	}
	if !stat.IsDir() {
		log.Fatalf("non-directory file found at path %s", out)
	}

	messagesInputReader, err := os.Open(messagesInputPath)
	if err != nil {
		log.Fatalf("failed to read %s: %v", messagesInputPath, err)
	}
	defer messagesInputReader.Close()

	records, err := ReadRecords(messagesInputReader, messageHeaders)
	if err != nil {
		log.Fatalf("failed while parsing %s: %v", messagesInputPath, err)
	}

	var messages []*Message
	for _, record := range records {
		message := &Message{
			Number:            record[0],
			Name:              record[1],
			Length:            record[2],
			TransitionCapable: record[3],
			Acknowledged:      record[4],
		}

		skip := false
		for _, name := range skipNames {
			if name == message.Name {
				skip = true
				break
			}
		}
		if skip {
			continue
		}

		messageDataInputFilename := fmt.Sprintf("%s.csv", Underscore(message.Name))
		messageDataInputPath := path.Join(in, messageDataInputFilename)
		_, err := os.Stat(messageDataInputPath)
		if os.IsNotExist(err) {
			log.Fatalf("file not found at path %s", messageDataInputPath)
		}
		if err != nil {
			log.Fatalf("failed to stat %s: %v", messageDataInputPath, err)
		}

		messageDataInputReader, err := os.Open(messageDataInputPath)
		if err != nil {
			log.Fatalf("failed to read %s: %v", messageDataInputPath, err)
		}

		records, err := ReadRecords(messageDataInputReader, messageDataHeaders)
		if err != nil {
			log.Fatalf("failed while parsing %s: %v", messageDataInputPath, err)
		}

		var byteField *DataByteField
		for _, record := range records {
			byteNumber := record[0]
			bitNumber := record[1]
			byteName := record[2]
			bitName := record[3]

			if byteName != "" {
				byteField = &DataByteField{
					Name: byteName,
					Byte: byteNumber,
				}

				message.DataByteFields = append(message.DataByteFields, byteField)
			}

			skip := false
			for _, name := range skipNames {
				if name == bitName {
					skip = true
					break
				}
			}
			if skip {
				continue
			}

			if bitName != "" {
				bitField := &DataBitField{
					Name: bitName,
					Bit:  bitNumber,
				}

				if byteField == nil {
					log.Fatalf("failed to assign bit field to missing byte field")
				}
				byteField.DataBitFields = append(byteField.DataBitFields, bitField)
				message.DataBitFields = append(message.DataBitFields, bitField)
			}
		}

		messages = append(messages, message)
	}

	for _, message := range messages {
		filename := fmt.Sprintf("%s.go", Underscore(message.Name))
		messageOutputPath := path.Join(out, filename)
		messageOutputFile, err := os.OpenFile(messageOutputPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			log.Fatalf("failed to open %s for writing: %v", messageOutputPath, err)
		}
		formattedMessageOutput := WrappedFormatWriteCloser(messageOutputFile)

		log.Printf("Generating %s", messageOutputPath)

		err = messageTemplate.Execute(formattedMessageOutput, message)
		if err != nil {
			log.Fatalf("failed to generate template %s: %v", messageOutputPath, err)
		}

		err = formattedMessageOutput.Close()
		if err != nil {
			log.Fatalf("failed to generate template %s: %v", messageOutputPath, err)
		}
	}

	messagesOutputPath := path.Join(out, messagesOutputFilename)
	messagesOutputFile, err := os.OpenFile(messagesOutputPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("failed to open %s for writing: %v", messagesOutputPath, err)
	}
	formattedMessagesOutput := WrappedFormatWriteCloser(messagesOutputFile)
	defer formattedMessagesOutput.Close()

	log.Printf("Generating %s", messagesOutputPath)

	err = messagesTemplate.Execute(formattedMessagesOutput, messages)
	if err != nil {
		log.Fatalf("failed to generate template %s: %v", messagesOutputPath, err)
	}
}

func (m *Message) FieldifiedName() string {
	return Fieldify(m.Name)
}

func (f *DataByteField) FieldifiedName() string {
	return Fieldify(f.Name)
}

func (f *DataBitField) FieldifiedName() string {
	return Fieldify(f.Name)
}

func Underscore(name string) string {
	re := regexp.MustCompile("[^a-zA-Z0-9]+")
	lower := strings.ToLower(name)
	underscored := re.ReplaceAllString(lower, "_")
	return underscored
}

func Fieldify(name string) string {
	re := regexp.MustCompile("[^a-zA-Z0-9]+")
	lower := strings.ToLower(name)
	fields := re.ReplaceAllString(lower, " ")
	title := strings.Title(fields)
	fieldified := strings.Replace(title, " ", "", -1)

	re = regexp.MustCompile("^[0-9]")
	if re.MatchString(fieldified) {
		fieldified = fmt.Sprintf("N%s", fieldified)
	}

	return fieldified
}

func ReadRecords(reader io.Reader, header []string) ([][]string, error) {
	var records [][]string

	csvReader := csv.NewReader(reader)
	isHeader := true

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if isHeader {
			for i, column := range record {
				if header[i] != column {
					return nil, fmt.Errorf("invalid header, expected: %v, actual: %v", header, record)
				}
			}
			isHeader = false
			continue
		}

		records = append(records, record)
	}

	return records, nil
}

type FormatWriteCloser struct {
	wc  io.WriteCloser
	buf *bytes.Buffer
}

func WrappedFormatWriteCloser(wc io.WriteCloser) io.WriteCloser {
	var buf bytes.Buffer
	return &FormatWriteCloser{
		wc:  wc,
		buf: &buf,
	}
}

func (w *FormatWriteCloser) Write(p []byte) (int, error) {
	return w.buf.Write(p)
}

func (w *FormatWriteCloser) Close() error {
	bytes := w.buf.Bytes()
	formatted, err := format.Source(bytes)
	if err != nil {
		w.wc.Write(w.buf.Bytes())
		return err
	}
	_, err = w.wc.Write(formatted)
	if err != nil {
		return err
	}
	return w.wc.Close()
}
