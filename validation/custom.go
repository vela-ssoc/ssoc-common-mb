package validation

import (
	"context"
	"net/url"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"unicode"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/robfig/cron/v3"
)

func Extensions() []CustomValidatorFunc {
	return []CustomValidatorFunc{
		dongFunc,
		uniqueFunc,
		requiredIfFunc,
		cronFunc,
		semverFunc,
		hostnamePortFunc,
		hostnameRFC1123Func,
		mobileFunc,
		httpURLFunc,
		wsURLFunc,
		usernameFunc,
		hostnameFunc,
		passwordFunc,
		substanceNameFunc,
		tagFunc,
		filepathFunc,
		filenameFunc,
	}
}

func dongFunc() (string, validator.FuncCtx, validator.RegisterTranslationsFunc) {
	const tag = "dong"
	valid := func(ctx context.Context, fl validator.FieldLevel) bool {
		field := fl.Field()
		if field.Kind() != reflect.String {
			return false
		}
		input := field.String()
		if sz := len(input); sz < 5 || sz > 8 {
			return false
		}
		for _, ch := range input {
			if unicode.IsUpper(ch) ||
				unicode.IsLetter(ch) ||
				unicode.IsNumber(ch) {
				continue
			}

			return false
		}

		return true
	}
	trans := func(utt ut.Translator) error {
		return utt.Add(tag, "{0}必须是一个合法的咚咚号", true)
	}

	return tag, valid, trans
}

func uniqueFunc() (string, validator.FuncCtx, validator.RegisterTranslationsFunc) {
	const tag = "unique"
	trans := func(utt ut.Translator) error {
		return utt.Add(tag, "{0}内的数据不能重复", true)
	}

	return tag, nil, trans
}

func requiredIfFunc() (string, validator.FuncCtx, validator.RegisterTranslationsFunc) {
	const tag = "required_if"
	trans := func(utt ut.Translator) error {
		return utt.Add(tag, "{0}为{1}时{2}必须填写", true)
	}

	return tag, nil, trans
}

func cronFunc() (string, validator.FuncCtx, validator.RegisterTranslationsFunc) {
	const tag = "cron"

	minuteCron := cron.NewParser(
		cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	)
	secondCron := cron.NewParser(
		cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	)

	valid := func(ctx context.Context, fl validator.FieldLevel) bool {
		field := fl.Field()
		if field.Kind() != reflect.String {
			return false
		}

		str := field.String()
		if _, err := minuteCron.Parse(str); err == nil {
			return true
		}
		_, err := secondCron.Parse(str)

		return err == nil
	}
	trans := func(utt ut.Translator) error {
		return utt.Add(tag, "{0}必须是一个合法的cron表达式", true)
	}

	return tag, valid, trans
}

// semverFunc 语义化版本号：https://semver.org/lang/zh-CN/
func semverFunc() (string, validator.FuncCtx, validator.RegisterTranslationsFunc) {
	const tag = "semver"
	trans := func(utt ut.Translator) error {
		return utt.Add(tag, "{0}必须是一个合法的版本号", true)
	}

	return tag, nil, trans
}

func hostnamePortFunc() (string, validator.FuncCtx, validator.RegisterTranslationsFunc) {
	const tag = "hostname_port"
	trans := func(utt ut.Translator) error {
		return utt.Add(tag, "{0}必须是主机名与端口号组合", true)
	}

	return tag, nil, trans
}

// hostnameRFC1123Func https://www.rfc-editor.org/rfc/rfc1123.html
func hostnameRFC1123Func() (string, validator.FuncCtx, validator.RegisterTranslationsFunc) {
	const tag = "hostname_rfc1123"
	trans := func(utt ut.Translator) error {
		return utt.Add(tag, "{0}必须是一个合法的域名", true)
	}

	return tag, nil, trans
}

// mobileFunc 中国大陆手机号校验器。
// 手机号段经常更新，尽量不要限制的太死。
// 如果要严格一点，请参考：https://www.jianshu.com/p/1e8eab706a63
func mobileFunc() (string, validator.FuncCtx, validator.RegisterTranslationsFunc) {
	tag := "mobile"
	valid := func(ctx context.Context, fl validator.FieldLevel) bool {
		field := fl.Field()
		if field.Kind() != reflect.String {
			return false
		}

		str := field.String()
		if len(str) != 11 {
			return false
		}
		for i, ch := range str {
			if i == 0 && ch != '1' {
				return false
			}
			if !unicode.IsNumber(ch) {
				return false
			}
		}

		return true
	}
	trans := func(utt ut.Translator) error {
		return utt.Add(tag, "{0}必须是合法的中国大陆手机号", true)
	}

	return tag, valid, trans
}

// httpURLFunc 必须是 http/https 协议的 URL
func httpURLFunc() (string, validator.FuncCtx, validator.RegisterTranslationsFunc) {
	const tag = "http"
	valid := func(ctx context.Context, fl validator.FieldLevel) bool {
		field := fl.Field()
		if field.Kind() != reflect.String {
			return false
		}
		str := field.String()
		pu, err := url.Parse(str)
		if err != nil {
			return false
		}
		if pu.Scheme != "http" && pu.Scheme != "https" {
			return false
		}
		if pu.Host == "" {
			return false
		}

		return true
	}
	trans := func(utt ut.Translator) error {
		return utt.Add(tag, "{0}必须是合法的http或https地址", true)
	}

	return tag, valid, trans
}

func wsURLFunc() (string, validator.FuncCtx, validator.RegisterTranslationsFunc) {
	const tag = "ws"
	valid := func(ctx context.Context, fl validator.FieldLevel) bool {
		field := fl.Field()
		if field.Kind() != reflect.String {
			return false
		}
		str := field.String()
		pu, err := url.Parse(str)
		if err != nil {
			return false
		}
		if pu.Scheme != "ws" && pu.Scheme != "wss" {
			return false
		}
		if pu.Host == "" {
			return false
		}

		return true
	}
	trans := func(utt ut.Translator) error {
		return utt.Add(tag, "{0}必须是合法的ws或wss地址", true)
	}

	return tag, valid, trans
}

// usernameFunc 用户名校验
func usernameFunc() (string, validator.FuncCtx, validator.RegisterTranslationsFunc) {
	const tag = "username"
	reg := regexp.MustCompile("^[a-zA-Z0-9._-]{2,20}$")
	valid := func(ctx context.Context, fl validator.FieldLevel) bool {
		if field := fl.Field(); field.Kind() == reflect.String {
			return reg.MatchString(field.String())
		}
		return false
	}
	trans := func(utt ut.Translator) error {
		return utt.Add(tag, "{0}必须是2-20位的数字、字母、减号、下划线组合", true)
	}

	return tag, valid, trans
}

// hostnameFunc 主机名校验器翻译
func hostnameFunc() (string, validator.FuncCtx, validator.RegisterTranslationsFunc) {
	const tag = "hostname"
	trans := func(utt ut.Translator) error {
		return utt.Add(tag, "{0}必须是合法的主机名", true)
	}

	return tag, nil, trans
}

// passwordFunc 密码校验
func passwordFunc() (string, validator.FuncCtx, validator.RegisterTranslationsFunc) {
	const tag = "password"
	valid := func(ctx context.Context, fl validator.FieldLevel) bool {
		field := fl.Field()
		if field.Kind() != reflect.String {
			return false
		}

		passwd := field.String()
		length := len(passwd)
		if length < 8 || length > 32 { // 密码8-32位
			return false
		}
		// 请参考 ASCII 码表，一目了然
		isLower := func(u rune) bool { return u >= 'a' && u <= 'z' }
		isUpper := func(u rune) bool { return u >= 'A' && u <= 'Z' }
		isNumber := func(u rune) bool { return u >= '0' && u <= '9' }
		isOther := func(u rune) bool {
			return (u >= '!' && u <= '/') ||
				(u >= ':' && u <= '@') ||
				(u >= '[' && u <= '`') ||
				(u >= '{' && u <= '~')
		}

		// 必须包含小写字母 大写字母 数字 特殊字符，且不能有其他字符
		var lower, upper, number, other bool
		for _, u := range passwd {
			if isLower(u) {
				lower = true
			} else if isUpper(u) {
				upper = true
			} else if isNumber(u) {
				number = true
			} else if isOther(u) {
				other = true
			} else {
				return false // 不符合密码要求的四种字符
			}
		}

		return lower && upper && number && other
	}
	trans := func(utt ut.Translator) error {
		return utt.Add(tag, "{0}不符合密码强度要求", true)
	}

	return tag, valid, trans
}

func substanceNameFunc() (string, validator.FuncCtx, validator.RegisterTranslationsFunc) {
	const tag = "substance_name"
	reg := regexp.MustCompile("^[a-zA-Z0-9-_@#\u4e00-\u9fa5]{1,20}$")
	valid := func(ctx context.Context, fl validator.FieldLevel) bool {
		if field := fl.Field(); field.Kind() == reflect.String {
			return reg.MatchString(field.String())
		}
		return false
	}
	trans := func(utt ut.Translator) error {
		return utt.Add(tag, "{0}不符合插件名要求", true)
	}

	return tag, valid, trans
}

// tagFunc 节点标签校验
func tagFunc() (string, validator.FuncCtx, validator.RegisterTranslationsFunc) {
	const tag = "tag"
	reg := regexp.MustCompile("^[\u4e00-\u9fa5a-zA-Z0-9@._-]{2,30}$")
	valid := func(ctx context.Context, fl validator.FieldLevel) bool {
		if field := fl.Field(); field.Kind() == reflect.String {
			return reg.MatchString(field.String())
		}
		return false
	}
	trans := func(utt ut.Translator) error {
		return utt.Add(tag, "{0}不符合标签名要求", true)
	}

	return tag, valid, trans
}

// filepathFunc 文件路径校验
func filepathFunc() (string, validator.FuncCtx, validator.RegisterTranslationsFunc) {
	const tag = "filepath"
	reg := regexp.MustCompile("^[a-zA-Z0-9./_\u4e00-\u9fa5\\\\]{3,50}$") // 一-龥
	valid := func(ctx context.Context, fl validator.FieldLevel) bool {
		if field := fl.Field(); field.Kind() == reflect.String {
			str := field.String()
			if !reg.MatchString(str) {
				return false
			}
			return !containsReservedName(str)
		}
		return false
	}
	trans := func(utt ut.Translator) error {
		return utt.Add(tag, "{0}不是合法的文件路径", true)
	}

	return tag, valid, trans
}

// filenameFunc 文件名校验器
func filenameFunc() (string, validator.FuncCtx, validator.RegisterTranslationsFunc) {
	const tag = "filename"
	// reg := regexp.MustCompile("^[a-zA-Z0-9\u4e00-\u9fa5-_.]{1,40}$")
	reg := regexp.MustCompile("^[a-zA-Z0-9._-]{1,50}$")
	valid := func(ctx context.Context, fl validator.FieldLevel) bool {
		f := fl.Field()
		if f.Kind() != reflect.String {
			return false
		}
		// 文件名不能以 . 开头
		str := f.String()
		if str == "" || str[0] == '.' {
			return false
		}
		// 要符合正则表达式
		if ok := reg.MatchString(str); !ok {
			return false
		}
		// 文件名不能是系统特殊含义字符
		ext := filepath.Ext(str) // 获取文件后缀
		if ext == ".zip" {       // zip 文件解压后的名字也不能是系统特殊含义字符
			str = str[:len(str)-4]
		}

		return !isReservedName(str)
	}
	trans := func(utt ut.Translator) error {
		return utt.Add(tag, "{0}不是一个合法的文件名", true)
	}

	return tag, valid, trans
}

// containsReservedName 判断路径中是否含有系统保留字
func containsReservedName(path string) bool {
	clean := filepath.Clean(path)
	format := strings.ReplaceAll(clean, "\\", "/")
	splits := strings.Split(format, "/")
	for _, ele := range splits {
		if isReservedName(ele) {
			return true
		}
	}
	return false
}

// reservedNames lists reserved Windows names. Search for PRN in
// https://docs.microsoft.com/en-us/windows/desktop/fileio/naming-a-file
// for details.
var reservedNames = []string{
	"CON", "PRN", "AUX", "NUL",
	"COM1", "COM2", "COM3", "COM4", "COM5", "COM6", "COM7", "COM8", "COM9",
	"LPT1", "LPT2", "LPT3", "LPT4", "LPT5", "LPT6", "LPT7", "LPT8", "LPT9",
}

// isReservedName returns true, if path is Windows reserved name.
// See reservedNames for the full list.
func isReservedName(path string) bool {
	if len(path) == 0 {
		return false
	}
	for _, reserved := range reservedNames {
		if strings.EqualFold(path, reserved) {
			return true
		}
	}
	return false
}
