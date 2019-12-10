package commands

import (
	"chatterbox/lib"
	"encoding/base64"
	"io/ioutil"
	"math"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	// This allows us to build cli while developing, but removes it on the portable app
	if !lib.IsAppBinary() {
		rootCmd.AddCommand(buildCmd)
	}
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Precompiles and builds chatterbox as a portable app",
	Run: func(cmd *cobra.Command, args []string) {
		precompile()
		build()
	},
}

// NOTE - Not as robust as I like, but good enough for military work
var jsRegExp = regexp.MustCompile("<script type=\"text/javascript\" src=\"([^\"]*)\"></script>")
var cssRegExp = regexp.MustCompile("<link rel=\"stylesheet\" type=\"text/css\" href=\"([^\"]*)\" />")

func precompile() {
	// Load our index file
	indexFile, _ := readAsset("index.html")

	// Now find and inject js scripts
	js := jsRegExp.FindAllStringSubmatch(indexFile, -1)
	for _, jsScriptMatch := range js {
		jsScript, _ := readAsset(jsScriptMatch[1])
		indexFile = strings.Replace(indexFile, jsScriptMatch[0], "<script> /* ["+jsScriptMatch[1]+"] */\r\n"+jsScript+"</script>\r\n", 1)
	}

	// Now find and inject style sheets
	css := cssRegExp.FindAllStringSubmatch(indexFile, -1)
	for _, cssScriptMatch := range css {
		cssScript, _ := readAsset(cssScriptMatch[1])
		indexFile = strings.Replace(indexFile, cssScriptMatch[0], "<style> /* ["+cssScriptMatch[1]+"] */\r\n"+cssScript+"</style>\r\n", 1)
	}

	// Now make golang representation of our combined file
	encodedIndexFile := base64.StdEncoding.EncodeToString([]byte(indexFile))

	// Generate our go file, chunking the base64 string to make things easier to read
	goTemplate := `package compiled

	import (
		"encoding/base64"
	)
	
	func IndexHTML() ([]byte, error) {
		return base64.StdEncoding.DecodeString(` + "`\r\n" + chunkString(encodedIndexFile, "\r\n", 100) + "`" + `)
	}
	`

	// Now make a new content file
	writeCompiledAsset(goTemplate)
}

func build() error {
	command := exec.Command("go", "build")
	if err := command.Run(); err != nil {
		return err
	}
	return nil
}

func readAsset(filename string) (string, error) {
	file, err := os.Open("./assets/" + filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	bytesContent, err := ioutil.ReadAll(file)

	return string(bytesContent), nil
}

func writeCompiledAsset(fileContent string) error {
	err := ioutil.WriteFile("./server/compiled/indexhtml.go", []byte(fileContent), 0644)
	if err != nil {
		return err
	}
	return nil
}

// TODO - I'm pretty sure this is not efficient . . . but it is late
func chunkString(input string, seperator string, chunkSize int) string {
	// Convert input to bytes
	inputBytes := []byte(input)

	// reduce chunkSize by one to make it compatible with the for loop iterator
	chunkSize--

	// Determine the slice size
	sliceSize := int(math.Ceil(float64(len(inputBytes)) / float64(chunkSize)))

	// Make our output slice to size
	outputSlice := make([]string, sliceSize)

	// Loop through the bytes and move to the next slice once we hit a byte's iterator that divides evenly into a chunk
	sliceIteration := 0
	outputSlice[sliceIteration] = ""
	for i, inputByte := range inputBytes {
		outputSlice[sliceIteration] += string(inputByte)
		if i != 0 && i%chunkSize == 0 {
			sliceIteration++
			outputSlice[sliceIteration] = ""
		}
	}

	return strings.Join(outputSlice[:], "\r\n")
}
