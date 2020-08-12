package pkg

import "context"

type AccountService Service

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

func (s *AccountService) Get(ctx context.Context) (*Account, *Response, error) {
	path := "/3/account/get.php"

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return nil, nil, err
	}

	var account *Account
	resp, err := s.client.Do(ctx, req, &account)
	if err != nil {
		return nil, resp, err
	}

	return account, resp, nil
}
