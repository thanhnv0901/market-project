package errorstrack

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"
)

const (
	MaximumStack = 50
)

type Frame struct {
	File           string  `json:"file"`                   // The LineNumber in that file
	Package        string  `json:"package"`                // The underlying ProgramCounter
	FunctionName   string  `json:"function_name"`          // The Package that contains this function
	LineNumber     int     `json:"line_number"`            // The Name of the function that contains this ProgramCounter
	Code           string  `json:"code"`                   // The Name of the function that contains this ProgramCounter
	ProgramCounter uintptr `json:"program_counter"`        // Program Counter
	ErrorMessage   string  `json:"error_mesage,omitempty"` // Error message
	NextFrame      *Frame  `json:"previous_frame"`
}
type PositionErrorsTrack int

type ErrorsTrack struct {
	Message                    string
	FlagWhereCreateErrorsTrack int
}

func New(msg string) ErrorsTrack {
	return ErrorsTrack{Message: msg}
}

func (f *Frame) String() string {
	tmp := fmt.Sprintf(`
	*** File: %s
	Package: %s
	FuncName: %s
	Line: %s: %d
	Error: %s
	`, f.File, f.Package, f.FunctionName, f.Code, f.LineNumber, f.ErrorMessage)
	return tmp
}

func (e *ErrorsTrack) PrintTrackingListLog() string {

	var (
		frame *Frame
		isOK  bool
	)

	stringBuilder := strings.Builder{}
	stringBuilder.WriteString("Errors logs tracking:")

	for i := 0; i < MaximumStack; i++ {

		frame, isOK = buildFrame(i)
		if !isOK {
			break
		}

		if i == 0 {
			frame.ErrorMessage = e.Message
		}

		stringBuilder.WriteString(frame.String())
	}

	return stringBuilder.String()
}

func (e *ErrorsTrack) PrintTrackingJSON() string {

	var (
		headFrame *Frame
		tailFrame *Frame
		frame     *Frame
		isOK      bool
	)

	headFrame = nil
	tailFrame = headFrame

	for i := 0; i < MaximumStack; i++ {

		frame, isOK = buildFrame(i)
		if !isOK {
			break
		}

		if i == 0 {
			frame.ErrorMessage = e.Message
		}

		if headFrame == nil {
			headFrame = frame
			tailFrame = headFrame
		} else {
			tailFrame.NextFrame = frame
			tailFrame = frame
		}

	}

	bytes, _ := json.Marshal(headFrame)
	return string(bytes)
}

func buildFrame(skip int) (*Frame, bool) {

	pc, file, line, ok := runtime.Caller(2 + skip)
	if !ok {
		return nil, false
	}

	var frame = new(Frame)
	packageName, functionName := packageAndName(runtime.FuncForPC(pc))

	frame.File = file
	frame.LineNumber = line
	frame.ProgramCounter = pc
	frame.Package = packageName
	frame.FunctionName = functionName
	frame.Code, _ = frame.sourceLine()

	return frame, true
}

func packageAndName(fn *runtime.Func) (string, string) {
	name := fn.Name()
	pkg := ""

	if lastslash := strings.LastIndex(name, "/"); lastslash >= 0 {
		pkg += name[:lastslash] + "/"
		name = name[lastslash+1:]
	}
	if period := strings.Index(name, "."); period >= 0 {
		pkg += name[:period]
		name = name[period+1:]
	}

	name = strings.Replace(name, "·", ".", -1)
	return pkg, name
}

// SourceLine gets the line of code (from File and Line) of the original source if possible.
func (f *Frame) sourceLine() (string, error) {
	if f.LineNumber <= 0 {
		return "???", nil
	}

	file, err := os.Open(f.File)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	currentLine := 1
	for scanner.Scan() {
		if currentLine == f.LineNumber {
			return string(bytes.Trim(scanner.Bytes(), " \t")), nil
		}
		currentLine++
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "???", nil
}
