package prompter

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
)

func Input(title, defaultValue string) (string, error) {
	var input string
	err := ask(&survey.Input{
		Message: title,
		Default: defaultValue,
	}, &input)
	if err != nil {
		return "", err
	}
	return input, nil
}

func Confirm(message, defaultValue string) (bool, error) {
	var confirm bool
	err := ask(&survey.Confirm{
		Message: message,
		Default: defaultValue == "No",
	}, &confirm)
	return confirm, err
}

func Token() (string, error) {
	var token string
	err := ask(&survey.Password{
		Message: "Paste your authentication token:",
	}, &token, survey.WithValidator(survey.Required))
	return token, err
}

func ask(q survey.Prompt, response interface{}, opts ...survey.AskOpt) error {
	opts = append(opts, survey.WithStdio(os.Stdin, os.Stdout, os.Stderr))
	if err := survey.AskOne(q, response, opts...); err != nil {
		return fmt.Errorf("could not prompt: %w", err)
	}
	return nil
}