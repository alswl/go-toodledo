package toodledo

// ErrorResponse ...
type ErrorResponse struct {
	ErrorCode int    `json:"errorCode"`
	ErrorDesc string `json:"errorDesc"`
}

// api doc: https://api.toodledo.com/3/error_codes.php
type ErrorCode int32

// UNKNOWN_ERROR ...
const (
	// General
	UNKNOWN_ERROR                                       ErrorCode = 0
	NO_ACCESS_TOKEN_WAS_GIVEN                           ErrorCode = 1 // WILL ALSO SET A 401 HTTP STATUS CODE. READ MORE...
	THE_ACCESS_TOKEN_WAS_INVALID_OR_HAD_THE_WRONG_SCOPE ErrorCode = 2 //WILL ALSO SET A 401 HTTP STATUS CODE.
	TOO_MANY_API_REQUESTS                               ErrorCode = 3 // WILL ALSO SET A 429 HTTP STATUS CODE. READ MORE...
	THE_API_IS_OFFLINE_FOR_MAINTENANCE                  ErrorCode = 4 // Will also set a 503 HTTP status code.

	// account/token.php
	ACCOUNT_SSL_CONNECTION_IS_REQUIRED_WHEN_REQUESTING_A_TOKEN ErrorCode = 101 // Read More...
	ACCOUNT_THERE_WAS_AN_ERROR_REQUESTING_A_TOKEN              ErrorCode = 102 // See error description for more details.
	ACCOUNT_TOO_MANY_TOKEN_REQUESTS                            ErrorCode = 103 // Will also set a 429 HTTP status code. Read More...

	// folders/add.php
	// folders/edit.php
	// folders/delete.php
	FOULDER_YOUR_FOLDER_MUST_HAVE_A_NAME           ErrorCode = 201
	FOULDER_A_FOLDER_WITH_THAT_NAME_ALREADY_EXISTS ErrorCode = 202
	FOULDER_MAX_FOLDERS_REACHED                    ErrorCode = 203
	FOULDER_EMPTY_ID                               ErrorCode = 204
	FOULDER_INVALID_FOLDER                         ErrorCode = 205
	FOULDER_NOTHING_WAS_EDITED                     ErrorCode = 206

	// contexts/add.php
	// contexts/edit.php
	// contexts/delete.php
	CONTEXT_YOUR_CONTEXT_MUST_HAVE_A_NAME           ErrorCode = 301
	CONTEXT_A_CONTEXT_WITH_THAT_NAME_ALREADY_EXISTS ErrorCode = 302
	CONTEXT_MAX_CONTEXTS_REACHED                    ErrorCode = 303
	CONTEXT_EMPTY_ID                                ErrorCode = 304
	CONTEXT_INVALID_CONTEXT                         ErrorCode = 305
	CONTEXT_NOTHING_WAS_EDITED                      ErrorCode = 306

	// goals/add.php
	// goals/edit.php
	// goals/delete.php
	GOAL_YOUR_GOAL_MUST_HAVE_A_NAME           ErrorCode = 401
	GOAL_A_GOAL_WITH_THAT_NAME_ALREADY_EXISTS ErrorCode = 402
	GOAL_MAX_GOALS_REACHED                    ErrorCode = 403
	GOAL_EMPTY_ID                             ErrorCode = 404
	GOAL_INVALID_GOAL                         ErrorCode = 405
	GOAL_NOTHING_WAS_EDITED                   ErrorCode = 406

	// locations/add.php
	// locations/edit.php
	// locations/delete.php
	LOCATION_YOUR_LOCATION_MUST_HAVE_A_NAME           ErrorCode = 501
	LOCATION_A_LOCATION_WITH_THAT_NAME_ALREADY_EXISTS ErrorCode = 502
	LOCATION_MAX_LOCATIONS_REACHED                    ErrorCode = 503
	LOCATION_EMPTY_ID                                 ErrorCode = 504
	LOCATION_INVALID_LOCATION                         ErrorCode = 505
	LOCATION_NOTHING_WAS_EDITED                       ErrorCode = 506

	// tasks/add.php
	// tasks/edit.php
	// tasks/delete.php
	// tasks/reassign.php
	// tasks/share.php
	TASK_YOUR_TASK_MUST_HAVE_A_TITLE                         ErrorCode = 601
	TASK_ONLY_50_TASKS_CAN_BE_ADDED_EDITED_DELETED_AT_A_TIME ErrorCode = 602
	TASK_MAX_TASKS_REACHED                                   ErrorCode = 603
	TASK_EMPTY_ID                                            ErrorCode = 604
	TASK_INVALID_TASK                                        ErrorCode = 605
	TASK_NOTHING_WAS_ADDED_EDITED                            ErrorCode = 606
	TASK_INVALID_FOLDER_ID                                   ErrorCode = 607
	TASK_INVALID_CONTEXT_ID                                  ErrorCode = 608
	TASK_INVALID_GOAL_ID                                     ErrorCode = 609
	TASK_INVALID_LOCATION_ID                                 ErrorCode = 610
	TASK_MALFORMED_REQUEST                                   ErrorCode = 611
	TASK_INVALID_PARENT_ID                                   ErrorCode = 612
	TASK_INCORRECT_FIELD_PARAMETERS                          ErrorCode = 613
	TASK_PARENT_WAS_DELETED                                  ErrorCode = 614
	TASK_INVALID_COLLABORATOR                                ErrorCode = 615
	TASK_UNABLE_TO_REASSIGN_OR_SHARE_TASK                    ErrorCode = 616
	TASK_REQUIRES_TOODLEDO_SUBSCRIPTION                      ErrorCode = 617

	// notes/add.php
	// notes/edit.php
	// notes/delete.php
	NOTE_YOUR_NOTE_MUST_HAVE_A_NAME                          ErrorCode = 701
	NOTE_ONLY_50_NOTES_CAN_BE_ADDED_EDITED_DELETED_AT_A_TIME ErrorCode = 702
	NOTE_MAX_NOTES_REACHED                                   ErrorCode = 703
	NOTE_EMPTY_ID                                            ErrorCode = 704
	NOTE_INVALID_NOTE                                        ErrorCode = 705
	NOTE_NOTHING_WAS_ADDED_EDITED                            ErrorCode = 706
	NOTE_INVALID_FOLDER_ID                                   ErrorCode = 707
	NOTE_MALFORMED_REQUEST                                   ErrorCode = 711

	// outlines/add.php
	// outlines/edit.php
	// outlines/delete.php
	OUTLINE_OUTLINE_HAD_NO_TITLE                                   ErrorCode = 801
	OUTLINE_ONLY_50_OUTLINES_CAN_BE_ADDED_EDITED_DELETED_AT_A_TIME ErrorCode = 802
	OUTLINE_MAX_OUTLINES_REACHED                                   ErrorCode = 803
	OUTLINE_EMPTY_ID___NOTHING_SENT                                ErrorCode = 804
	OUTLINE_INVALID_OUTLINE                                        ErrorCode = 805
	OUTLINE_NOTHING_WAS_ADDED_EDITED                               ErrorCode = 806
	OUTLINE_INVALID_OUTLINE_ID                                     ErrorCode = 807
	OUTLINE_MAX_NODES_REACHED                                      ErrorCode = 808
	OUTLINE_OUTLINE_ALREADY_ADDED                                  ErrorCode = 809
	OUTLINE_MALFORMED_REQUEST                                      ErrorCode = 811
	OUTLINE_REFERENCE_WAS_EMPTY                                    ErrorCode = 812
	OUTLINE_INVALID_OUTLINE_FORMAT                                 ErrorCode = 813
	OUTLINE_EDITING_WRONG_VERSION                                  ErrorCode = 814

	// lists/add.php
	// lists/edit.php
	// lists/delete.php
	LIST_LIST_HAD_NO_TITLE                                   ErrorCode = 901
	LIST_ONLY_50_LISTS_CAN_BE_ADDED_EDITED_DELETED_AT_A_TIME ErrorCode = 902
	LIST_MAX_LISTS_REACHED                                   ErrorCode = 903
	LIST_EMPTY_ID___NOTHING_SENT                             ErrorCode = 904
	LIST_INVALID_LIST                                        ErrorCode = 905
	LIST_NOTHING_WAS_ADDED_EDITED                            ErrorCode = 906
	LIST_LIST_ALREADY_ADDED                                  ErrorCode = 909
	LIST_MALFORMED_REQUEST                                   ErrorCode = 911
	LIST_REFERENCE_WAS_EMPTY                                 ErrorCode = 912
	LIST_INVALID_COLS_FORMAT                                 ErrorCode = 913
	LIST_EDITING_WRONG_VERSION                               ErrorCode = 914

	// rows/add.php
	// rows/edit.php
	// rows/delete.php
	ROW_ROW_HAD_NO_CELLS                                   ErrorCode = 1001
	ROW_ONLY_50_ROWS_CAN_BE_ADDED_EDITED_DELETED_AT_A_TIME ErrorCode = 1002
	ROW_MAX_ROWS_REACHED                                   ErrorCode = 1003
	ROW_EMPTY_ID___NOTHING_SENT                            ErrorCode = 1004
	ROW_INVALID_LIST                                       ErrorCode = 1005
	ROW_NOTHING_WAS_ADDED_EDITED                           ErrorCode = 1006
	ROW_INVALID_ROW                                        ErrorCode = 1007
	ROW_ROW_ALREADY_ADDED                                  ErrorCode = 1009
	ROW_MALFORMED_REQUEST                                  ErrorCode = 1011
	ROW_REFERENCE_WAS_EMPTY                                ErrorCode = 1012
	ROW_INVALID_CELLS_FORMAT                               ErrorCode = 1013
	ROW_EDITING_WRONG_VERSION                              ErrorCode = 1014
)
