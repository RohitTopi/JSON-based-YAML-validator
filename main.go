package main

import ( 
    "fmt"
    "io/ioutil"
    "path/filepath"
    "github.com/ghodss/yaml"
    "github.com/xeipuuv/gojsonschema"
)


func check(e error) {
    if e != nil {
        //panic(e)
        fmt.Println(e)
    }
}

func getDirName() string{
    // get the current src location, relative to which other file paths will be obtained
    // only the directory path is required
    filename, pathErr := filepath.Abs("./source.go")
    check(pathErr)
    return filepath.Dir(filename)
}

func getInputFilePathName() string{
    fmt.Print("Enter path to YAML file (without any spaces) : \n")
    var input string
    fmt.Scanln(&input)
    return input
}

// converts YAML to JSON and writes to a .json file
func converter(inputYAMLPath string, outputJSONPath string){
    // read YAML data to a byte array
    yamlData, readError := ioutil.ReadFile(inputYAMLPath)
    check(readError)

    // convert YAML to JSON
    jsonFormData, conversionError := yaml.YAMLToJSON(yamlData)
    check(conversionError)

    // write to a .json file
    writeError := ioutil.WriteFile(outputJSONPath, jsonFormData, 0644)
    check(writeError)
}

// validates whether the instance file conforms to the schema
func validator(inputJSONURIPath string, inputSchemaURIPath string) bool {

    // load the "schema file" and the "instance file"
    schemaLoader := gojsonschema.NewReferenceLoader(inputSchemaURIPath)
    instanceLoader := gojsonschema.NewReferenceLoader(inputJSONURIPath)

    // validate; store result and possible errors in "result"
    result, err := gojsonschema.Validate(schemaLoader, instanceLoader)
    check(err)

    // print result
    if result.Valid() {
       return true
   } else {

        // print errors if necessary
        /*
        for _, err := range result.Errors() {
        fmt.Printf("- %s\n", err)
        }
        */
    
    return false
    }
}

func main() {

    // INPUT: path of file to be used
    // inputYAMLPath := "./exampleValid.yaml"
    inputYAMLPath := getInputFilePathName()

    // convert and store JSON in a file
    outputJSONPath := "./generated.json"
    converter(inputYAMLPath, outputJSONPath)

    // get URI's using current directory
    dirName:= getDirName()
    inputJSONURIPath := "file://" + dirName + "/generated.json"
    inputSchemaURIPath := "file://" + dirName + "/schema.json"

    /*
        alternately in Windows use pathname URI 
        getDirName cannot be used here because Windows uses "\" instead of "/"
        inputJSONURIPath := "file:///E:/validator/generated.json"
        inputSchemaURIPath := "file:///E:/validator/schema.json"
    */


    // OUTPUT: true or false
    if validator(inputJSONURIPath, inputSchemaURIPath){
        fmt.Printf("true\n")
    } else {
        fmt.Printf("false\n")
    }

}