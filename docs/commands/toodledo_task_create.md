## toodledo task create

Create a task

```
toodledo task create [flags]
```

### Examples

```
$ toodledo tasks create --context=1 --folder=2 --goal=3 --priority=High --due_date=2020-01-01 title

```

### Options

```
      --context string    context
      --due-date string   format 2021-01-01 (omitempty,datetime=2006-01-02)
      --folder string     folder
      --goal string       goal
  -h, --help              help for create
      --priority string   priority (omitempty,oneof=Top top High high Medium medium Low low Negative negative)
      --status string     status (omitempty,oneof=None NextAction Active Planning Delegated Waiting Hold Postponed Someday Canceled Reference none nextaction active planning delegated waiting hold postponed someday canceled reference)
      --title string      title
```

### Options inherited from parent commands

```
      --access_token string   
      --config string         config file (default is $HOME/config/toodledo/conf.yaml)
```

### SEE ALSO

* [toodledo task](toodledo_task.md)	 - Manage toodledo tasks

###### Auto generated by spf13/cobra on 12-Sep-2022