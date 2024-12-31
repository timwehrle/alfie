package view

import (
	"fmt"
	"github.com/timwehrle/asana/api"
	"github.com/timwehrle/asana/utils"
	"strings"
)

func FormatNames(tasks []api.Task) []string {
	names := make([]string, len(tasks))
	for i, task := range tasks {
		names[i] = fmt.Sprintf("%d. [%s] %s", i+1, utils.FormatDate(task.DueOn), task.Name)
	}
	return names
}

func FormatList(prefix string, items []string) string {
	if len(items) == 0 {
		return prefix + "None"
	}
	return prefix + strings.Join(items, ", ")
}

func FormatProjects(projects []api.Project) string {
	names := make([]string, len(projects))
	for i, project := range projects {
		names[i] = project.Name
	}
	return FormatList("Projects: ", names)
}

func FormatTags(tags []api.Tag) string {
	names := make([]string, len(tags))
	for i, tag := range tags {
		names[i] = tag.Name
	}
	return FormatList("Tags: ", names)
}

func FormatNotes(notes string) string {
	if notes == "" {
		return ""
	}
	return utils.BoldUnderline().Sprint("Description:") + "\n" + notes + "\n"
}