package plugins

import "errors"

var (
	Hooks *HooksHolder
)

func init() {
	Hooks = NewHooksHolder()
}

type HooksHolder struct {
	hooks map[string][]string
}

func NewHooksHolder() *HooksHolder {
	return &HooksHolder{
		hooks: make(map[string][]string),
	}
}

func (this *HooksHolder) AddHook(hook, module string) error {
	if _, exists := this.hooks[hook]; !exists {
		this.hooks[hook] = append(this.hooks[hook], module)
		return nil
	}
	return errors.New("already exists")
}

func (this *HooksHolder) Callback(hook string, v interface{}) error {
	if modules, ok := this.hooks[hook]; ok {
		for _, module := range modules {
			if err := RunModule(module, hook, v); err != nil {
				return err
			}
		}
	}
	return nil
}
