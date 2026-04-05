package commit

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

var GenerateCommitMessageShortcut = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Generate a commit message from the global shortcut",
	ExtraCmdArgs: []string{},
	Skip:         false,
	SetupConfig: func(config *config.AppConfig) {
		config.GetUserConfig().Git.Commit.GenerateCommand = "printf 'feat: generated commit\\n\\nGenerated body'"
	},
	SetupRepo: func(shell *Shell) {
		shell.CreateFileAndAdd("file", "file content")
	},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		t.Views().Files().Focus().Press(keys.Universal.GenerateCommitMessage)

		t.ExpectPopup().CommitMessagePanel().
			Title(Equals("Commit summary")).
			Content(Equals("feat: generated commit")).
			SwitchToDescription().
			Content(Equals("Generated body"))
	},
})
