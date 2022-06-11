package service

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

type BuilderService interface {
	Build(name string, goOS string, goArch string, content string) (string, error)
}

type builderService struct {
}

func NewBuilderService() BuilderService {
	return &builderService{}
}

func (b builderService) Build(name string, goOS string, goArch string, content string) (string, error) {
	source, err := saveSource(content)
	if err != nil {
		return "", err
	}

	defer os.Remove(source)

	exeFile, err := buildSource(source, name, goOS, goArch)
	if err != nil {
		return "", err
	}
	return exeFile, nil
}

func buildSource(source string, exeName string, goOS string, goArch string) (string, error) {
	exeFile, _ := ioutil.TempFile("/tmp", fmt.Sprintf("%s-%s-%s-*", exeName, goOS, goArch))
	cmd := exec.Command("go", "build", "-o", exeFile.Name(), source)

	cmd.Env = append(os.Environ(),
		fmt.Sprintf("GOOS=%s", goOS),
		fmt.Sprintf("GOARCH=%s", goArch),
	)

	err := cmd.Run()

	if err != nil {
		return "", err
	}
	return exeFile.Name(), nil
}

func saveSource(content string) (string, error) {
	sourceFile, err := ioutil.TempFile("/tmp", "toucan.*.go")
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(sourceFile.Name(), []byte(content), 0644)
	if err != nil {
		return "", err
	}
	return sourceFile.Name(), nil
}
