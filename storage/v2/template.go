// Copyright 2015 Prometheus Team
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package storage

import (
	"encoding/json"
	"fmt"
	tmplhtml "html/template"
	"io"
	"math"
	"net/url"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	tmpltext "text/template"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Template bundles a text and a html template instance.
type Template struct {
	text *tmpltext.Template
	html *tmplhtml.Template
}

// Option is generic modifier of the text and html templates used by a Template.
type Option func(text *tmpltext.Template, html *tmplhtml.Template)

// NewTemplate returns a new Template with the DefaultFuncs added. The DefaultFuncs
// have precedence over any added custom functions. Options allow customization
// of the text and html templates in given order.
func NewTemplate(id string, opts ...Option) *Template {
	t := &Template{
		text: tmpltext.New(id).Option("missingkey=zero"),
		html: tmplhtml.New(id).Option("missingkey=zero"),
	}

	for _, o := range opts {
		o(t.text, t.html)
	}

	t.text.Funcs(DefaultFuncs)
	t.html.Funcs(DefaultFuncs)

	return t
}

// Parse parses the given text into the template.
func (t *Template) Parse(str string) error {
	var err error
	if t.text, err = t.text.Parse(str); err == nil {
		t.html, err = t.html.Parse(str)
	}

	return err
}

// FromGlob calls ParseGlob on given path glob provided and parses into the
// template.
func (t *Template) FromGlob(path string) error {
	// ParseGlob in the template packages errors if not at least one file is
	// matched. We want to allow empty matches that may be populated later on.
	p, err := filepath.Glob(path)
	if err != nil {
		return err
	}
	if len(p) > 0 {
		if t.text, err = t.text.ParseGlob(path); err != nil {
			return err
		}
		if t.html, err = t.html.ParseGlob(path); err != nil {
			return err
		}
	}
	return nil
}

// ExecuteText needs a meaningful doc comment (TODO(fabxc)).
func (t *Template) ExecuteText(w io.Writer, tpl string, data any) error {
	if tpl == "" {
		return nil
	}

	tmpl, err := t.text.Clone()
	if err != nil {
		return err
	}

	ptpl, err := tmpl.Parse(tpl)
	if err != nil {
		return err
	}

	return ptpl.Execute(w, data)
}

// ExecuteHTML needs a meaningful doc comment (TODO(fabxc)).
func (t *Template) ExecuteHTML(w io.Writer, tpl string, data any) error {
	if tpl == "" {
		return nil
	}

	tmpl, err := t.html.Clone()
	if err != nil {
		return err
	}

	ptpl, err := tmpl.Parse(tpl)
	if err != nil {
		return err
	}

	return ptpl.Execute(w, data)
}

var DefaultFuncs = map[string]any{
	"toUpper": strings.ToUpper,
	"toLower": strings.ToLower,
	"title": func(text string) string {
		// Casers should not be shared between goroutines, instead
		// create a new caser each time this function is called.
		return cases.Title(language.AmericanEnglish).String(text)
	},
	"trimSpace": strings.TrimSpace,
	// join is equal to strings.Join but inverts the argument order
	// for easier pipelining in templates.
	"join": func(sep string, s []string) string {
		return strings.Join(s, sep)
	},
	"match": regexp.MatchString,
	"safeHtml": func(text string) tmplhtml.HTML {
		return tmplhtml.HTML(text)
	},
	"safeUrl": func(text string) tmplhtml.URL {
		return tmplhtml.URL(text)
	},
	"urlUnescape": url.QueryUnescape,
	"reReplaceAll": func(pattern, repl, text string) string {
		re := regexp.MustCompile(pattern)
		return re.ReplaceAllString(text, repl)
	},
	"stringSlice": func(s ...string) []string {
		return s
	},
	// date returns the text representation of the time in the specified format.
	"date": func(fmt string, t time.Time) string {
		return t.Format(fmt)
	},
	// tz returns the time in the timezone.
	"tz": func(name string, t time.Time) (time.Time, error) {
		loc, err := time.LoadLocation(name)
		if err != nil {
			return time.Time{}, err
		}
		return t.In(loc), nil
	},
	"now":              time.Now,
	"since":            time.Since,
	"humanizeDuration": HumanizeDuration,
	"toJson": func(v any) (string, error) {
		bytes, err := json.Marshal(v)
		if err != nil {
			return "", err
		}
		return string(bytes), nil
	},
	"list": func(args ...any) ([]any, error) {
		if args == nil {
			return []any{}, nil
		}
		return args, nil
	},
	"append": func(slice []any, args ...any) []any {
		return append(slice, args...)
	},
	"dict": func(values ...any) (map[string]any, error) {
		if len(values)%2 != 0 {
			return nil, fmt.Errorf("dict requires an even number of arguments")
		}

		res := make(map[string]any, len(values)/2)
		for i := 0; i < len(values); i += 2 {
			key, ok := values[i].(string)
			if !ok {
				return nil, fmt.Errorf("dict keys must be strings")
			}
			res[key] = values[i+1]
		}

		return res, nil
	},
}

func HumanizeDuration(i any) (string, error) {
	v, err := ConvertToFloat(i)
	if err != nil {
		return "", err
	}

	if math.IsNaN(v) || math.IsInf(v, 0) {
		return fmt.Sprintf("%.4g", v), nil
	}
	if v == 0 {
		return fmt.Sprintf("%.4gs", v), nil
	}
	if math.Abs(v) >= 1 {
		sign := ""
		if v < 0 {
			sign = "-"
			v = -v
		}
		duration := int64(v)
		seconds := duration % 60
		minutes := (duration / 60) % 60
		hours := (duration / 60 / 60) % 24
		days := duration / 60 / 60 / 24
		// For days to minutes, we display seconds as an integer.
		if days != 0 {
			return fmt.Sprintf("%s%dd %dh %dm %ds", sign, days, hours, minutes, seconds), nil
		}
		if hours != 0 {
			return fmt.Sprintf("%s%dh %dm %ds", sign, hours, minutes, seconds), nil
		}
		if minutes != 0 {
			return fmt.Sprintf("%s%dm %ds", sign, minutes, seconds), nil
		}
		// For seconds, we display 4 significant digits.
		return fmt.Sprintf("%s%.4gs", sign, v), nil
	}
	prefix := ""
	for _, p := range []string{"m", "u", "n", "p", "f", "a", "z", "y"} {
		if math.Abs(v) >= 1 {
			break
		}
		prefix = p
		v *= 1000
	}
	return fmt.Sprintf("%.4g%ss", v, prefix), nil
}

func ConvertToFloat(i any) (float64, error) {
	switch v := i.(type) {
	case float64:
		return v, nil
	case string:
		return strconv.ParseFloat(v, 64)
	case int:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case time.Duration:
		return v.Seconds(), nil
	default:
		return 0, fmt.Errorf("can't convert %T to float", v)
	}
}

// Pair is a key/value string pair.
type Pair struct {
	Name, Value string
}

// Pairs is a list of key/value string pairs.
type Pairs []Pair

// Names returns a list of names of the pairs.
func (ps Pairs) Names() []string {
	ns := make([]string, 0, len(ps))
	for _, p := range ps {
		ns = append(ns, p.Name)
	}
	return ns
}

// Values returns a list of values of the pairs.
func (ps Pairs) Values() []string {
	vs := make([]string, 0, len(ps))
	for _, p := range ps {
		vs = append(vs, p.Value)
	}
	return vs
}

func (ps Pairs) String() string {
	b := new(strings.Builder)
	for i, p := range ps {
		b.WriteString(p.Name)
		b.WriteRune('=')
		b.WriteString(p.Value)
		if i < len(ps)-1 {
			b.WriteString(", ")
		}
	}
	return b.String()
}
