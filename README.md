# validation
JSON based validation of YAML document
### Tools used
Generator tool used to generate schema https://jsonschema.net/ 
(generated schema was modified) \
Schema validator used is https://github.com/xeipuuv/gojsonschema \
Converter tool used is a wrapper for go-yaml, https://github.com/ghodss/yaml
<pre>
go get github.com/xeipuuv/gojsonschema
go get github.com/ghodss/yaml
</pre>
### Sample I/O

#### sample 1:
<pre>$ go run main.go
Enter path to YAML file (without any spaces) : 
./tasks/QC-CHECK-RUNNER-Q4H0.yaml
true
</pre>

#### sample 2:
<pre>$ go run main.go
Enter path to YAML file (without any spaces) : 
./exampleInvalid.yaml
false
</pre>
  

