# Mnorepo
The Fortify project is setup using a monorepo strategy.

## Warning
Github actions don't integrate too well with a monorepo structure. Actions where designed to operate omultirepo setuo, where main project is in the root directory. to componsate for this issues most commands will need to be setup to change the workinf directory to location the action should run. This makes action files very verbose
