package e2e

import (
	prompt "github.com/c-bata/go-prompt"
	"github.com/ktr0731/evans/adapter/inputter"
	"github.com/ktr0731/evans/tests/helper"
)

// SetPrompt replaces NewPrompt var by newPrompt which is prompter injected.
// SetPrompt returns cleanup func as the result.
// caller must call cleanup after each tests.
func SetPrompt(pmt *helper.MockPrompt) func() {
	old := inputter.NewRealPrompter
	inputter.NewRealPrompter = func(executor func(string), completer func(prompt.Document) []prompt.Suggest, opt ...prompt.Option) inputter.Prompter {
		if executor == nil {
			executor = func(in string) {
				return
			}
		}
		if completer == nil {
			completer = func(d prompt.Document) []prompt.Suggest {
				return nil
			}
		}
		pmt.Executor = executor
		pmt.Completer = completer
		p := pmt
		return p
	}
	return func() {
		inputter.NewRealPrompter = old
	}
}
