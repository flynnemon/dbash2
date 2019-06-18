package prompts

import (
	"strings"
	"fmt"

	"github.com/manifoldco/promptui"
	
	"../../models"
	"../docker"
)

func DockerPrompts(containers []models.Container) models.Container {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U000021E8 {{ .Name | green }}",
		Inactive: "  {{ .Name | white }}",
		Selected: "\U00002714 Connecting to {{ .Name | white }}",
		Details: `
------------- Selected Container --------------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Image:" | faint }}	{{ .Image }}
{{ "Created At:" | faint }}	{{ .CreatedAt }}`,
	}
	
	searcher := func(input string, index int) bool {
		container := containers[index]
		name := strings.Replace(strings.ToLower(container.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)
	
		return strings.Contains(name, input)
	}
	
	prompt := promptui.Select{
		Label:     "Which container to access",
		Items:     docker.GetContainers(),
		Templates: templates,
		Size:      10,
		Searcher:  searcher,
	}
	
	i, _, err := prompt.Run()
	
	if err != nil {
		fmt.Printf("Exited %v\n", err)
	}

	return containers[i]
}