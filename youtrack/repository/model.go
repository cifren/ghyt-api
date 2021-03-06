package repository

const (
	FieldStyleFields string = "color(background,foreground)"
	TagFields string = "id,name," + FieldStyleFields
	IssueFields string = "id,idReadable,summary,description,tags(" + TagFields +")"
	UserFields string = "tags(" + TagFields + ")"
)

type Issue struct {
	Id string `json:"id"`
	IdReadable string `json:"idReadable"`
	Summary string `json:"summary"`
	Description string `json:"description"`
	Tags []Tag `json:"tags"`
}
func (this Issue) HasTag(tag Tag) bool {
	for _ , value := range this.Tags {
		if value.Id == tag.Id {
			return true
		}
	}
	return false
}
func (this *Issue) AddTag(tag Tag) Issue {
	if !this.HasTag(tag) {
		this.Tags = append(this.Tags, tag)
	}

	return *this
}

type User struct {
	Tags []Tag `json:"tags"`
}

type Tag struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Color FieldStyle `json:"color"`
}

type FieldStyle struct {
	Background string `json:"background"`
	Foreground string `json:"foreground"`
}
