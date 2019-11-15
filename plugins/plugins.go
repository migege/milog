package plugins

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/astaxie/beego"
)

var (
	plugindir string
)

func init() {
	plugindir = beego.AppConfig.String("plugindir")
	LoadPlugins(plugindir)
}

func getPython() string {
	return beego.AppConfig.String("python")
}

func RunModule(module, hook string, v interface{}) error {
	py := getPython()
	if py == "" {
		return errors.New("Python is missing")
	}
	cmd := exec.Command(py, "-c", fmt.Sprintf("from %s import %s; %s('%s')", module, hook, hook, v))
	out, err := cmd.Output()
	if err != nil {
		log.Printf("cmd run error: %s", cmd.String())
		return err
	}

	if strings.TrimSpace(string(out)) != "0" {
		return errors.New(fmt.Sprintf("%s.%s() returns %s", module, hook, out))
	}

	return nil
}

func LoadPlugins(plugindir string) error {
	py := getPython()
	if py == "" {
		return errors.New("Python is missing")
	}
	return filepath.Walk(plugindir, func(path string, info os.FileInfo, _ error) error {
		if info.IsDir() {
			return nil
		}

		if !strings.HasSuffix(info.Name(), ".py") {
			log.Println("only python supported at present")
			return nil
		}

		module := strings.Split(info.Name(), ".")[0]

		cmd := exec.Command(py, "-c", fmt.Sprintf("from %s import register; register()", module))
		log.Printf("cmd: %s", cmd.String())
		paths := strings.Split(os.Getenv("PYTHONPATH"), ":")
		paths = append(paths, plugindir)
		os.Setenv("PYTHONPATH", strings.Join(paths, ":"))

		out, err := cmd.Output()
		if err != nil {
			log.Printf("register plugin: %s, %s", path, err)
			return nil
		} else {
			log.Printf("register plugin: %s, succeeded", path)
		}

		var hooks []string
		json.Unmarshal(out, &hooks)
		for _, hook := range hooks {
			Hooks.AddHook(hook, module)
		}

		return nil
	})
}
