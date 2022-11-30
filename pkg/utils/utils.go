package utils

import (
	"fmt"
	"github.com/gobeam/stringy"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os/exec"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// ChinaTimeZone is timezone of China
// TODO using configuration time zone
var ChinaTimeZone = time.FixedZone("CST", 8*3600)

// Bool2int ...
func Bool2int(input bool) (output int) {
	if input {
		output = 1
	} else {
		output = 0
	}
	return
}

// Bool2ints ...
func Bool2ints(input bool) (output string) {
	return strconv.Itoa(Bool2int(input))
}

// OpenBrowser ...
func OpenBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		logrus.Error(err)
	}
}

// BindFlagsByQuery bind flags to query, the query support basic types, and support `validate:"required"`
// obj must be a struct
func BindFlagsByQuery(cmd *cobra.Command, obj interface{}) error {
	if reflect.ValueOf(obj).Type().Kind() != reflect.Struct {
		return fmt.Errorf("%s is not a struct", reflect.ValueOf(obj).Type().Kind())
	}
	getType := reflect.TypeOf(obj)
	for i := 0; i < getType.NumField(); i++ {
		f := getType.Field(i)
		name := stringy.New(f.Name).KebabCase().ToLower()
		desc := f.Tag.Get("description")
		if desc == "" {
			desc = name
		}
		validateTags := f.Tag.Get("validate")
		if validateTags != "" {
			desc = fmt.Sprintf("%s (%s)", desc, validateTags)
		}

		short := f.Tag.Get("short")
		switch f.Type.Kind() {
		case reflect.Struct:
			if f.Type == reflect.TypeOf(time.Time{}) {
				cmd.Flags().StringP(name, "", "", desc)
			} else {
				return fmt.Errorf("%s is a struct", name)
			}
		case reflect.Bool:
			cmd.Flags().BoolP(name, short, false, desc)
		case reflect.Int:
			cmd.Flags().IntP(name, short, 0, desc)
		case reflect.String:
			cmd.Flags().StringP(name, short, "", desc)
		case reflect.Int32:
			cmd.Flags().Int32P(name, short, 0, desc)
		case reflect.Int64:
			cmd.Flags().Int64P(name, short, 0, desc)
		case reflect.Slice:
			switch f.Type.Elem().Kind() {
			case reflect.String:
				cmd.Flags().StringSliceP(name, short, []string{}, desc)
			case reflect.Int:
				cmd.Flags().IntSliceP(name, short, []int{}, desc)
			case reflect.Int32:
				cmd.Flags().Int32SliceP(name, short, []int32{}, desc)
			case reflect.Int64:
				cmd.Flags().Int64SliceP(name, short, []int64{}, desc)
			default:
				return fmt.Errorf("%s is a slice of %s, not supported", name, f.Type.Elem().Kind())
			}
		default:
			return fmt.Errorf("%s is a %s, not supported", name, f.Type.Kind())
		}

		if strings.Contains(validateTags, "required") {
			_ = cmd.MarkFlagRequired(name)
		}
	}
	return nil
}

// FillQueryByFlags fill struct by cmd query
// obj must be a pointer
func FillQueryByFlags(cmd *cobra.Command, obj interface{}) error {
	if reflect.ValueOf(obj).Type().Kind() != reflect.Ptr {
		return fmt.Errorf("%s is not a pointer", reflect.ValueOf(obj).Type().Kind())
	}

	getType := reflect.TypeOf(obj).Elem()
	getValue := reflect.ValueOf(obj).Elem()
	for i := 0; i < getType.NumField(); i++ {
		f := getType.Field(i)
		v := getValue.Field(i)
		name := stringy.New(f.Name).KebabCase().ToLower()
		//desc := f.Tag.Get("description")
		//if desc == "" {
		//	desc = name
		//}
		if !v.CanSet() {
			return fmt.Errorf("%s is not settable, please check the accessor type", name)
		}
		switch f.Type.Kind() {
		case reflect.Struct:
			if f.Type == reflect.TypeOf(time.Time{}) {
				get, err := cmd.Flags().GetString(name)
				if err != nil {
					return errors.Wrapf(err, "get %s", name)
				}
				parsed, err := time.ParseInLocation("2006-01-02", get, ChinaTimeZone)
				// if err, not set value
				if err == nil {
					v.Set(reflect.ValueOf(parsed))
				}
			} else {
				return fmt.Errorf("%s is a struct", name)
			}
		case reflect.Bool:
			get, err := cmd.Flags().GetBool(name)
			if err != nil {
				return errors.Wrapf(err, "get %s", name)
			}
			v.SetBool(get)
		case reflect.Int:
			get, err := cmd.Flags().GetInt(name)
			if err != nil {
				return errors.Wrapf(err, "get %s", name)
			}
			v.SetInt(int64(get))
		case reflect.Int32:
			get, err := cmd.Flags().GetInt32(name)
			if err != nil {
				return errors.Wrapf(err, "get %s", name)
			}
			v.Set(reflect.ValueOf(get))
		case reflect.Int64:
			get, err := cmd.Flags().GetInt64(name)
			if err != nil {
				return errors.Wrapf(err, "get %s", name)
			}
			v.Set(reflect.ValueOf(get))
		case reflect.String:
			get, err := cmd.Flags().GetString(name)
			if err != nil {
				return errors.Wrapf(err, "get %s", name)
			}
			v.SetString(get)
		case reflect.Slice:
			switch f.Type.Elem().Kind() {
			case reflect.String:
				get, err := cmd.Flags().GetStringSlice(name)
				if err != nil {
					return errors.Wrapf(err, "get %s", name)
				}
				v.Set(reflect.ValueOf(get))
			case reflect.Int:
				get, err := cmd.Flags().GetIntSlice(name)
				if err != nil {
					return errors.Wrapf(err, "get %s", name)
				}
				v.Set(reflect.ValueOf(get))
			case reflect.Int32:
				get, err := cmd.Flags().GetInt32Slice(name)
				if err != nil {
					return errors.Wrapf(err, "get %s", name)
				}
				v.Set(reflect.ValueOf(get))
			case reflect.Int64:
				get, err := cmd.Flags().GetInt64Slice(name)
				if err != nil {
					return errors.Wrapf(err, "get %s", name)
				}
				v.Set(reflect.ValueOf(get))
			default:
				return fmt.Errorf("%s is a slice of %s, not supported", name, f.Type.Elem().Kind())
			}
		default:
			return fmt.Errorf("%s is a %s, not supported", name, f.Type.Kind())
		}
	}
	return nil
}
