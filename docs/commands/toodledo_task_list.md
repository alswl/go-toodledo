## toodledo task list

List tasks

```
toodledo task list [flags]
```

### Examples

```
$ toodledo tasks list
$ toodledo tasks list --limit 20
$ toodledo tasks list --context Work
$ toodledo tasks list --context-id 4455
$ toodledo tasks list --folder inbox
$ toodledo tasks list --folder-id 4455
$ toodledo tasks list --goal landing-moon
$ toodledo tasks list --goal-id 4455
$ toodledo tasks list --priority High
$ toodledo tasks list --status Active
$ toodledo tasks list --due-date "2020-01-01"

```

### Options

```
      --complete                complete (omitempty)
      --context string          context
      --context-id int          context-id
      --due-date string         format 2021-01-01 (omitempty,datetime=2006-01-02)
      --folder string           folder
      --folder-id int           folder-id
      --format string           format (omitempty,oneof=name json yaml)
      --goal string             goal
      --goal-id int             goal-id
  -h, --help                    help for list
      --incomplete              incomplete (omitempty)
      --limit int32             limit
      --priority string         priority (omitempty,oneof=Top top High high Medium medium Low low Negative negative)
      --status string           status (omitempty,oneof=None NextAction Active Planning Delegated Waiting Hold Postponed Someday Canceled Reference none nextaction active planning delegated waiting hold postponed someday canceled reference)
      --sub-tasks-mode string   sub-tasks-mode (omitempty,oneof=Inline Hidden Indented inline hidden indented)
```

### Options inherited from parent commands

```
      --access_token string   
      --config string         config file (default is $HOME/config/toodledo/conf.yaml)
```

### SEE ALSO

* [toodledo task](toodledo_task.md)	 - Manage toodledo tasks

###### Auto generated by spf13/cobra on 17-Jun-2023
