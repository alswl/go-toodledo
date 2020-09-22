package models

type Account struct {
	Userid           string `json:"userid"`
	Alias            string `json:"alias"`
	Email            string `json:"email"`
	Pro              int    `json:"pro"`
	Dateformat       int    `json:"dateformat"`
	Timezone         int    `json:"timezone"`
	Hidemonths       int    `json:"hidemonths"`
	Hotlistpriority  int    `json:"hotlistpriority"`
	Hotlistduedate   int    `json:"hotlistduedate"`
	Showtabnums      int    `json:"showtabnums"`
	LasteditFolder   int    `json:"lastedit_folder"`
	LasteditContext  int    `json:"lastedit_context"`
	LasteditGoal     int    `json:"lastedit_goal"`
	LasteditLocation int    `json:"lastedit_location"`
	LasteditTask     int    `json:"lastedit_task"`
	LastdeleteTask   int    `json:"lastdelete_task"`
	LasteditNote     int    `json:"lastedit_note"`
	LastdeleteNote   int    `json:"lastdelete_note"`
	LasteditList     int    `json:"lastedit_list"`
	LasteditOutline  int    `json:"lastedit_outline"`
}

type Folder struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Private  int    `json:"private"`
	Archived int    `json:"archived"`
	Ord      int    `json:"ord"`
}

type Context struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Private int    `json:"private"`
}

type GoalLevel int

type Goal struct {
	ID    int       `json:"id"`
	Name  string    `json:"name"`
	Level GoalLevel `json:"level"`
	// 0 or 1
	Archived    int    `json:"archived"`
	Contributes int    `json:"contributes"`
	Note        string `json:"note"`
}

type GoalAdd struct {
	// required
	Name  string `validate:"required`
	Level *GoalLevel
	// 0 or 1
	Contributes *int
	// 0 or 1
	Private *bool
	Note    *string
}

type GoalEdit struct {
	GoalAdd
}

// TODO @alswl mote fields
type Task struct {
	ID        int    `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Modified  int    `json:"modified,omitempty"`
	Completed int    `json:"completed,omitempty"`
	Folder    int    `json:"folder,omitempty"`
	Star      int    `json:"star,omitempty"`
	Ref       string `json:"ref,omitempty"`
}

type TaskQuery struct {
	before int64
	after  int64
	comp   bool
	start  int
	num    int
	fields string
}

type TaskAdd struct {
	ID        int    `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Modified  int    `json:"modified,omitempty"`
	Completed int    `json:"completed,omitempty"`
	Folder    int    `json:"folder,omitempty"`
	Star      int    `json:"star,omitempty"`
	Ref       string `json:"ref,omitempty"`
	// folder, context, goal, location, priority, status,star, duration, remind, starttime, duetime, completed, duedatemod, repeat, tag, duedate, startdate, note, parent, meta
}

type Collaborators struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Reassignable int    `json:"reassignable"`
	Sharable     int    `json:"sharable"`
}

type SavedSearch struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Bool   string `json:"bool"`
	Search struct {
		Num1 []struct {
			Field string `json:"field"`
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"1"`
		Root []struct {
			Field string `json:"field"`
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"root"`
	} `json:"search"`
}

type Note struct {
	Num      int    `json:"num,omitempty"`
	Total    int    `json:"total,omitempty"`
	ID       int    `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	Modified int    `json:"modified,omitempty"`
	Added    int    `json:"added,omitempty"`
	Folder   int    `json:"folder,omitempty"`
	Private  int    `json:"private,omitempty"`
	Text     string `json:"text,omitempty"`
}
