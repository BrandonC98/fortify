# Mnorepo
The Fortify project is setup using a monorepo strategy. The reason being is to keep all the work together in one place and to explore the pros and cons of working with Monorepos.

## Warning
Github actions don't integrate too well with a monorepo structure. Actions where designed to operate multirepos, where main project is in the root directory. to componsate for this issues most commands will need to be setup to change the working directory to location the action should run. This makes action files very verbose
