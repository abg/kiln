package builder

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"text/template"

	yamlConverter "github.com/ghodss/yaml"
	"gopkg.in/yaml.v2"
)

type Interpolator struct{}

type InterpolateInput struct {
	Version            string
	BOSHVariables      map[string]interface{}
	Variables          map[string]interface{}
	ReleaseManifests   map[string]interface{}
	StemcellManifests  map[string]interface{}
	StemcellManifest   interface{}
	FormTypes          map[string]interface{}
	IconImage          string
	InstanceGroups     map[string]interface{}
	Jobs               map[string]interface{}
	PropertyBlueprints map[string]interface{}
	RuntimeConfigs     map[string]interface{}
	StubReleases       bool

	PreProcess func(in []byte) ([]byte, error)
	RDelim, LDelim string
}

func (input InterpolateInput) WithDefaultMetadataPreprocessor() InterpolateInput {
	var names []string
	for name := range (Interpolator{}).functions(InterpolateInput{}) {
		names = append(names, name)
	}
	input.RDelim = "}}"
	input.LDelim = "{{"
	input.PreProcess = func(in []byte) ([]byte, error) {
		re := regexp.MustCompile(`\$\( (?P<func>((` + strings.Join(names, ")|(") + `))) (?P<line>.*)\)`)
		matches :=  re.FindAllSubmatchIndex(in, -1)
		var result []byte
		if len(matches) == 0 {
			result = in
		} else {
			for _, match := range matches {
				result = re.Expand(result, []byte(input.LDelim + ` ${func} ${line}` + input.RDelim), in, match)
			}
		}
		return result, nil
	}
	return input
}

func NewInterpolator() Interpolator {
	return Interpolator{}
}

func (i Interpolator) Interpolate(input InterpolateInput, templateYAML []byte) ([]byte, error) {
	if input.LDelim == "" {
		input.LDelim = "$("
	}
	if input.RDelim == "" {
		input.RDelim = ")"
	}

	interpolatedYAML, err := i.interpolate(input, templateYAML)
	if err != nil {
		return nil, err
	}

	prettyMetadata, err := i.prettyPrint(interpolatedYAML)
	if err != nil {
		return nil, err // un-tested
	}

	return prettyMetadata, nil
}

func (i Interpolator) functions(input InterpolateInput) template.FuncMap {
	return template.FuncMap{
		"bosh_variable": func(key string) (string, error) {
			if input.BOSHVariables == nil {
				return "", errors.New("--bosh-variables-directory must be specified")
			}
			val, ok := input.BOSHVariables[key]
			if !ok {
				return "", fmt.Errorf("could not find bosh variable with key '%s'", key)
			}
			return i.interpolateValueIntoYAML(input, val)
		},
		"form": func(key string) (string, error) {
			if input.FormTypes == nil {
				return "", errors.New("--forms-directory must be specified")
			}
			val, ok := input.FormTypes[key]
			if !ok {
				return "", fmt.Errorf("could not find form with key '%s'", key)
			}

			return i.interpolateValueIntoYAML(input, val)
		},
		"property": func(name string) (string, error) {
			if input.PropertyBlueprints == nil {
				return "", errors.New("--properties-directory must be specified")
			}
			val, ok := input.PropertyBlueprints[name]
			if !ok {
				return "", fmt.Errorf("could not find property blueprint with name '%s'", name)
			}
			return i.interpolateValueIntoYAML(input, val)
		},
		"regexReplaceAll": func(regex, inputString, replaceString string) (string, error) {
			re, err := regexp.Compile(regex)
			if err != nil {
				return "", err
			}
			return re.ReplaceAllString(inputString, replaceString), nil
		},
		"release": func(name string) (string, error) {
			if input.ReleaseManifests == nil {
				return "", errors.New("missing ReleaseManifests")
			}

			val, ok := input.ReleaseManifests[name]

			if !ok {
				if input.StubReleases {
					val = map[string]interface{}{
						"name":    name,
						"version": "UNKNOWN",
						"file":    fmt.Sprintf("%s-UNKNOWN.tgz", name),
						"sha1":    "dead8e1ea5e00dead8e1ea5ed00ead8e1ea5e000",
					}
				} else {
					return "", fmt.Errorf("could not find release with name '%s'", name)
				}
			}

			return i.interpolateValueIntoYAML(input, val)
		},
		"stemcell": func(osname ...string) (string, error) {
			if input.StemcellManifest == nil && len(input.StemcellManifests) == 0 {
				return "", errors.New("stemcell specification must be provided through either --stemcells-directory or --kilnfile")
			}

			if len(input.StemcellManifests) == 0 && len(osname) > 0 {
				return "", errors.New("$( stemcell \"<osname>\" ) cannot be used without --stemcells-directory being provided")
			}

			if len(input.StemcellManifests) > 1 && len(osname) == 0 {
				return "", errors.New("stemcell template helper requires osname argument if multiple stemcells are specified")
			}

			if len(osname) > 0 {
				return i.interpolateValueIntoYAML(input, input.StemcellManifests[osname[0]])
			}

			if len(input.StemcellManifests) == 1 {
				for _, stemcell := range input.StemcellManifests {
					return i.interpolateValueIntoYAML(input, stemcell)
				}
			}

			return i.interpolateValueIntoYAML(input, input.StemcellManifest)
		},
		"version": func() (string, error) {
			if input.Version == "" {
				return "", errors.New("--version must be specified")
			}
			return i.interpolateValueIntoYAML(input, input.Version)
		},
		"variable": func(key string) (string, error) {
			if input.Variables == nil {
				return "", errors.New("--variable or --variables-file must be specified")
			}
			val, ok := input.Variables[key]
			if !ok {
				return "", fmt.Errorf("could not find variable with key '%s'", key)
			}
			return i.interpolateValueIntoYAML(input, val)
		},
		"icon": func() (string, error) {
			if input.IconImage == "" {
				return "", fmt.Errorf("--icon must be specified")
			}
			return input.IconImage, nil
		},
		"instance_group": func(name string) (string, error) {
			if input.InstanceGroups == nil {
				return "", errors.New("--instance-groups-directory must be specified")
			}
			val, ok := input.InstanceGroups[name]
			if !ok {
				return "", fmt.Errorf("could not find instance_group with name '%s'", name)
			}

			return i.interpolateValueIntoYAML(input, val)
		},
		"job": func(name string) (string, error) {
			if input.Jobs == nil {
				return "", errors.New("--jobs-directory must be specified")
			}
			val, ok := input.Jobs[name]
			if !ok {
				return "", fmt.Errorf("could not find job with name '%s'", name)
			}

			return i.interpolateValueIntoYAML(input, val)
		},
		"runtime_config": func(name string) (string, error) {
			if input.RuntimeConfigs == nil {
				return "", errors.New("--runtime-configs-directory must be specified")
			}
			val, ok := input.RuntimeConfigs[name]
			if !ok {
				return "", fmt.Errorf("could not find runtime_config with name '%s'", name)
			}

			return i.interpolateValueIntoYAML(input, val)
		},
		"select": func(field, input string) (string, error) {
			object := map[string]interface{}{}

			err := json.Unmarshal([]byte(input), &object)
			if err != nil {
				return "", fmt.Errorf("could not JSON unmarshal %q: %s", input, err)
			}

			value, ok := object[field]
			if !ok {
				return "", fmt.Errorf("could not select %q, key does not exist", field)
			}

			output, err := json.Marshal(value)
			if err != nil {
				return "", fmt.Errorf("could not JSON marshal %q: %s", input, err) // NOTE: this cannot happen because value was unmarshalled from JSON
			}

			return string(output), nil
		},
		"tile": func() (string, error) {
			const key = "tile-name"
			val, ok := input.Variables[key]
			if !ok {
				return "", fmt.Errorf("could not find variable with key '%s'", key)
			}
			str, ok := val.(string)
			if !ok {
				return "", fmt.Errorf("variable %[1]q is %[2]T expected string: %[1]s=%[2]v", key, val)
			}
			return strings.ToLower(str), nil
		},
	}
}

func (i Interpolator) interpolate(input InterpolateInput, templateYAML []byte) ([]byte, error) {
	if input.PreProcess != nil {
		var err error
		templateYAML, err = input.PreProcess(templateYAML)
		if err != nil {
			return nil, fmt.Errorf("pre-process failure: %w", err)
		}
	}

	t, err := template.New("metadata").
		Delims(input.LDelim, input.RDelim).
		Funcs(i.functions(input)).
		Option("missingkey=error").
		Parse(string(templateYAML))

	if err != nil {
		return nil, fmt.Errorf("template parsing failed: %s", err)
	}

	var buffer bytes.Buffer
	err = t.Execute(&buffer, input.Variables)
	if err != nil {
		return nil, fmt.Errorf("template execution failed: %s", err)
	}

	return buffer.Bytes(), nil
}

func (i Interpolator) interpolateValueIntoYAML(input InterpolateInput, val interface{}) (string, error) {
	initialYAML, err := yaml.Marshal(val)
	if err != nil {
		return "", err // should never happen
	}

	interpolatedYAML, err := i.interpolate(input, initialYAML)
	if err != nil {
		return "", fmt.Errorf("unable to interpolate value: %s", err)
	}

	inlinedYAML, err := i.yamlMarshalOneLine(interpolatedYAML)
	if err != nil {
		return "", err // un-tested
	}

	return string(inlinedYAML), nil
}

// Workaround to avoid YAML indentation being incorrect when value is interpolated into the metadata
func (i Interpolator) yamlMarshalOneLine(yamlContents []byte) ([]byte, error) {
	return yamlConverter.YAMLToJSON(yamlContents)
}

func (i Interpolator) prettyPrint(inputYAML []byte) ([]byte, error) {
	var data interface{}
	err := yaml.Unmarshal(inputYAML, &data)
	if err != nil {
		return []byte{}, err // should never happen
	}

	return yaml.Marshal(data)
}
