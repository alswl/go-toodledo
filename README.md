# go-toodledo

Go library and Cli for Toodledo.

Status: Under Development

## Usage

```sh
> toodledo task list --context home --status nextaction
         # │ [X] │ TITLE            │     STATUS │ CONTEXT │ PRIORITY │ FOLDER  │ GOAL │ DUE        │ REPEAT      │ LENGTH │ TIMER
───────────┼─────┼──────────────────┼────────────┼─────────┼──────────┼─────────┼──────┼────────────┼─────────────┼────────┼───────
 327077755 │ [ ] │ next-action item │ NextAction │ home    │     High │ to-0128 │ b    │ 2022-02-24 │ FREQ=WEEKLY │ 1m0s   │ 20m0s
 327078471 │ [ ] │ abc2x            │ NextAction │ home    │     High │         │ c    │            │             │        │
```

More commands usage can be found in the [Manual](./docs/toodledo.md).

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
> make
> ./bin/toodledo --help
```