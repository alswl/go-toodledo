# go-toodledo

Go CLI / TUI / SDK for Toodledo.

Status: CLI Stable / SDK Stable / TUI under development

## Usage (CLI)

```sh
> toodledo task list --context home --status nextaction
         # │ [X] │ TITLE            │     STATUS │ CONTEXT │ PRIORITY │ FOLDER  │ GOAL │ DUE        │ REPEAT      │ LENGTH │ TIMER
───────────┼─────┼──────────────────┼────────────┼─────────┼──────────┼─────────┼──────┼────────────┼─────────────┼────────┼───────
 327077755 │ [ ] │ next-action item │ NextAction │ home    │     High │ to-0128 │ b    │ 2022-02-24 │ FREQ=WEEKLY │ 1m0s   │ 20m0s
 327078471 │ [ ] │ abc2x            │ NextAction │ home    │     High │         │ c    │            │             │        │
```

More commands usage can be found in the [Manual](./docs/commands/toodledo.md).

### Quick Start

Login:

```sh
# login
toodledo auth login
# follow steps, open link your browser
toodledo auth token YOUR-CODE
# verify
toodledo auth status
```

Tasks:

```shell
# list
> toodledo task list --context home --goal goal-b
INFO[0002] Syncing tasks
         # │ [X] │ TITLE    │ STATUS │ CONTEXT │ PRIORITY │ FOLDER │ GOAL   │ DUE │ REPEAT │ LENGTH │ TIMER
───────────┼─────┼──────────┼────────┼─────────┼──────────┼────────┼────────┼─────┼────────┼────────┼───────
 330000079 │ [ ] │ cooking3 │   None │ home    │   Medium │        │ goal-b │     │        │        │

# complete
> toodledo task complete 323245685
         # │ [X] │ TITLE  │ STATUS │ CONTEXT │ PRIORITY │ FOLDER  │ GOAL    │ DUE │ REPEAT │ LENGTH │ TIMER
───────────┼─────┼────────┼────────┼─────────┼──────────┼─────────┼─────────┼─────┼────────┼────────┼───────
 320000085 │ [X] │ test-c │   None │ Not Set │      Low │ Not Set │ Not Set │   0 │        │      0 │     0

# edit
> toodledo task edit --title cooking4-8 --context a --status nextaction 334313701
         # │ [X] │ TITLE      │     STATUS │ CONTEXT │ PRIORITY │ FOLDER │ GOAL   │ DUE │ REPEAT │ LENGTH │ TIMER
───────────┼─────┼────────────┼────────────┼─────────┼──────────┼────────┼────────┼─────┼────────┼────────┼───────
 330000001 │ [ ] │ cooking4-8 │ NextAction │ a       │   Medium │ a      │ goal-b │     │        │        │
```

Goals:

```sh
> toodledo folder list
       # │ NAME                         │ ARCHIVED
─────────┼──────────────────────────────┼──────────
 3000096 │ AnalysisDesign               │        0
 4000073 │ Business                     │        0
 4000039 │ Custom-Support               │        1
 3000006 │ Coding                       │        0
 3000000 │ Family                       │        0
```

More commands usage can be found in the [Manual](./docs/toodledo.md).

## Build

```sh
> git clone https://github.com/alswl/go-toodledo.git
> make clean build
> cp ./bin/toodledo /usr/local/bin
> cp ./bin/tt /usr/local/bin
> toodledo --help
```


## One more thing

`tt` is a full-featured TUI app(under actively development).
You can get the development version from `cmd/tt`.

```
┌──────────────────────┐┌──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐
│<Contexts>            ││[ ]                                       Title   Context  Priority    Status      Goal DueStringRepe…Leng…Timer       Tag│
││ All                 ││[ ]Access your tasks anywhere with the Toodled…              Medium      None                                             │
│  a                   ││[ ]   Bookmark your task list (CTRL+D or CMD+D)                High      None                                             │
│  b                   ││[ ]               Visit our How-to Video Guides              Medium      None                                             │
│  computer            ││[ ]Import your tasks from your previous task m…              Medium      None                                             │
│  home                ││[ ]            Create your first recurring task                High      None                                             │
│  mobile              ││[ ]        Create a new task on your to-do list                 Top      None                                             │
│  outing              ││[ ]                                         abc                 Low      None                                             │
│  work                ││[ ]                            next-action item      home      HighNextAction         b2022-03-03FREQ…6000…               │
│  None                ││[ ]                                       abc2x      home      HighNextAction         c                                   │
│                      ││[ ]                                     abc-low      home       Low      None                                             │
│                      ││[ ]                                     abc-due              Medium      None          2022-01-30                         │
│                      ││[ ]                                   abc-due-2              Medium      None          2022-01-31                         │
│                      ││[ ]                                    work-t-1      work       LowNextAction                                             │
│                      ││[ ]                                       new-2                 Low      None         d                                   │
│                      ││[ ]                                       new-5                 Low      None         d                                   │
│                      ││[ ]                                     new-7-3                 Low      None         d                                   │
│                      ││[ ]                                     cooking      home    Medium      None                                             │
│                      ││[ ]                                    cooking2      home    Medium      None                                             │
│                      ││[ ]                                    cooking2      home    Medium      None                                             │
│                      ││[ ]                                    cooking2      home    Medium      None                                             │
│                      ││[ ]                                    cooking2      home    Medium      None                                             │
│                      ││[ ]                                    cooking3      home    Medium      None                                             │
│                      ││[ ]                                    cooking3      home    Medium      None    goal-b                                   │
│                      ││[ ]                                 cooking a-4         a    MediumNextAction    goal-b                    4490…          │
│                      ││[ ]                                      红楼梦                 Low      None                                             │
│                      ││[ ]                                      写作文                 Low      None                                             │
│                      ││[ ]                                      item-a                 Low      None                              2136…          │
│                      ││[ ]                                  没有 foldr                 Low      None                                             │
│                      ││[ ]                                 context a-1         a       Low      None                                             │
└──────────────────────┘└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
 tasks  INFO: tasks: 41                                                                                                                ./0  HELP(h)
```