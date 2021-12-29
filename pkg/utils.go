package pkg

import (
	"fmt"
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
		desc := field.Tag.Get("description")
		if desc == "" {
			desc = field.Name
		}
		switch field.Type.Kind() {
		case reflect.Struct:
			if field.Type == reflect.TypeOf(time.Time{}) {
				cmd.Flags().StringP(field.Name, "", "", desc)
			} else {
				return fmt.Errorf("%s is a struct", field.Name)
			}
		case reflect.Bool:
			cmd.Flags().BoolP(field.Name, field.Tag.Get("short"), false, desc)
		case reflect.Int:
			cmd.Flags().IntP(field.Name, field.Tag.Get("short"), 0, desc)
		case reflect.String:
			cmd.Flags().StringP(field.Name, field.Tag.Get("short"), "", desc)
		case reflect.Int32:
			cmd.Flags().Int32P(field.Name, field.Tag.Get("short"), 0, desc)
		case reflect.Int64:
			cmd.Flags().Int64P(field.Name, field.Tag.Get("short"), 0, desc)
		case reflect.Slice:
			switch field.Type.Elem().Kind() {
			case reflect.String:
				cmd.Flags().StringSliceP(field.Name, field.Tag.Get("short"), []string{}, desc)
			case reflect.Int:
				cmd.Flags().IntSliceP(field.Name, field.Tag.Get("short"), []int{}, desc)
			case reflect.Int32:
				cmd.Flags().Int32SliceP(field.Name, field.Tag.Get("short"), []int32{}, desc)
			case reflect.Int64:
				cmd.Flags().Int64SliceP(field.Name, field.Tag.Get("short"), []int64{}, desc)
			default:
				return fmt.Errorf("%s is a slice of %s, not supported", field.Name, field.Type.Elem().Kind())
			}
		default:
			return fmt.Errorf("%s is a %s, not supported", field.Name, field.Type.Kind())
		}
	}
	return nil
}
