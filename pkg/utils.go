package pkg

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
	"time"
)

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

func GenerateFlagsByStructure(cmd *cobra.Command, obj interface{}) error {
	t := reflect.TypeOf(obj)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		name := stringy.New(field.Name).SnakeCase().ToLower()
		desc := field.Tag.Get("description")
		if desc == "" {
			desc = name
		}
		switch field.Type.Kind() {
		case reflect.Struct:
			if field.Type == reflect.TypeOf(time.Time{}) {
				cmd.Flags().StringP(name, "", "", desc)
			} else {
				return fmt.Errorf("%s is a struct", name)
			}
		case reflect.Bool:
			cmd.Flags().BoolP(name, field.Tag.Get("short"), false, desc)
		case reflect.Int:
			cmd.Flags().IntP(name, field.Tag.Get("short"), 0, desc)
		case reflect.String:
			cmd.Flags().StringP(name, field.Tag.Get("short"), "", desc)
		case reflect.Int32:
			cmd.Flags().Int32P(name, field.Tag.Get("short"), 0, desc)
		case reflect.Int64:
			cmd.Flags().Int64P(name, field.Tag.Get("short"), 0, desc)
		case reflect.Slice:
			switch field.Type.Elem().Kind() {
			case reflect.String:
				cmd.Flags().StringSliceP(name, field.Tag.Get("short"), []string{}, desc)
			case reflect.Int:
				cmd.Flags().IntSliceP(name, field.Tag.Get("short"), []int{}, desc)
			case reflect.Int32:
				cmd.Flags().Int32SliceP(name, field.Tag.Get("short"), []int32{}, desc)
			case reflect.Int64:
				cmd.Flags().Int64SliceP(name, field.Tag.Get("short"), []int64{}, desc)
			default:
				return fmt.Errorf("%s is a slice of %s, not supported", name, field.Type.Elem().Kind())
			}
		default:
			return fmt.Errorf("%s is a %s, not supported", name, field.Type.Kind())
		}
	}
	return nil
}

func FillQueryByStructuredCmd(cmd *cobra.Command, obj interface{}) error {
	if reflect.ValueOf(obj).Type().Kind() != reflect.Ptr {
		return fmt.Errorf("%s is not a pointer", reflect.ValueOf(obj).Type().Kind())
	}

	getType := reflect.TypeOf(obj).Elem()
	getValue := reflect.ValueOf(obj).Elem()
	for i := 0; i < getType.NumField(); i++ {
		f := getType.Field(i)
		v := getValue.Field(i)
		name := stringy.New(f.Name).SnakeCase().ToLower()
		desc := f.Tag.Get("description")
		if desc == "" {
			desc = name
		}
		if !v.CanSet() {
			return fmt.Errorf("%s is not settable, please check the accessor type", name)
		}
		switch f.Type.Kind() {
		case reflect.Struct:
			if f.Type == reflect.TypeOf(time.Time{}) {
				get, err := cmd.Flags().GetString(name)
				if err != nil {
					return errors.Wrapf(err, "failed to get %s", name)
				}
				parsed, err := time.ParseInLocation("2006-01-02", get, time.Local)
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
				return errors.Wrapf(err, "failed to get %s", name)
			}
			v.SetBool(get)
		case reflect.Int:
			get, err := cmd.Flags().GetInt(name)
			if err != nil {
				return errors.Wrapf(err, "failed to get %s", name)
			}
			v.SetInt(int64(get))
		case reflect.Int32:
			get, err := cmd.Flags().GetInt32(name)
			if err != nil {
				return errors.Wrapf(err, "failed to get %s", name)
			}
			v.Set(reflect.ValueOf(get))
		case reflect.Int64:
			get, err := cmd.Flags().GetInt64(name)
			if err != nil {
				return errors.Wrapf(err, "failed to get %s", name)
			}
			v.Set(reflect.ValueOf(get))
		case reflect.String:
			get, err := cmd.Flags().GetString(name)
			if err != nil {
				return errors.Wrapf(err, "failed to get %s", name)
			}
			v.SetString(get)
		case reflect.Slice:
			switch f.Type.Elem().Kind() {
			case reflect.String:
				get, err := cmd.Flags().GetStringSlice(name)
				if err != nil {
					return errors.Wrapf(err, "failed to get %s", name)
				}
				v.Set(reflect.ValueOf(get))
			case reflect.Int:
				get, err := cmd.Flags().GetIntSlice(name)
				if err != nil {
					return errors.Wrapf(err, "failed to get %s", name)
				}
				v.Set(reflect.ValueOf(get))
			case reflect.Int32:
				get, err := cmd.Flags().GetInt32Slice(name)
				if err != nil {
					return errors.Wrapf(err, "failed to get %s", name)
				}
				v.Set(reflect.ValueOf(get))
			case reflect.Int64:
				get, err := cmd.Flags().GetInt64Slice(name)
				if err != nil {
					return errors.Wrapf(err, "failed to get %s", name)
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
