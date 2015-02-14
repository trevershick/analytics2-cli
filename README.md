[![Stories in Ready](https://badge.waffle.io/trevershick/analytics2-cli.png?label=ready&title=Ready)](https://waffle.io/trevershick/analytics2-cli)
[![Build Status](https://travis-ci.org/trevershick/analytics2-cli.svg?branch=master)](https://travis-ci.org/trevershick/analytics2-cli)
# analytics2-cli

This repo houses the a2m tool which is a Go based command line tool for communicating with
the A2 services (API and ETL).  The project is in its infancy.

*Important Note* - this CLI only works with the `workservice-plugin` branch of the analytics2 repo.



# Installing A2M

	go install github.com/rallysoftware/analytics2-cli/a2m

	trevershick@Trevers-Mac-mini ~/projects/go
	> $ a2m

	NAME:
	   a2m - A2 Management CLI

	USAGE:
	   a2m [global options] command [command options] [arguments...]

	VERSION:
	   0.0.1

	AUTHOR:
	  Trever Shick - <tshick@rallydev.com>

	COMMANDS:
	   catalog		Show the work catalog
	   tasks		Show the tasks in-flight
	   revision-count	Show the number of revisions in a workspace's queue
	   halted		Shows if a workspace is halted
	   help, h		Shows a list of commands or help for one command

	GLOBAL OPTIONS:
	   --base, -b "http://localhost:9201/analytics-etl"	Base URL to connect to [$A2M_BASE_URL]
	   --user, -u "nobody"					User Name to connect to the web service [$A2M_USER]
	   --pass, -p "nothing"					Password to use for authentication [$A2M_PASSWORD]
	   --save, -s						Save supplied arguments to ~/.a2mrc
	   --help, -h						show help
	   --version, -v					print the version

# Viewing Active and Recent Tasks in ETL

	> $ a2m tasks --recent

	Active Tasks @ http://localhost:9201/analytics-etl/management/work/tasks
	=============================================================================
	revisionQueueTask                        COMPLETED  2015-02-08T19:42:10Z 2015-02-08T19:42:10Z

# Viewing the Task Catalog

	> $ a2m catalog

	Available Tasks in the Work Catalog @ http://localhost:9201/analytics-etl/management/work/catalog%!(EXTRA url.Values=map[])
	=============================================================================
	allAllowedValuesRefreshTask              : Refresh the allowed values metadata for all workspaces
	allowedValuesRefreshTask                 : Refresh the allowed values for the provided workspace
	haltAllWorkspacesInDatabaseTask          : Halt all workspaces by database partition
	iterationRefreshTask                     : Refresh the iterations for the provided workspace
	iterationsRefreshTask                    : Refresh the iterations for all workspaces
	processRevisionsTask                     : Processes the queued revision for the provided workspace id
	projectRefreshTask                       : Refresh the projects for the provided workspace
	projectsRefreshTask                      : Refresh the projects for all workspaces
	releaseRefreshTask                       : Refresh the releases for the provided workspace
	releasesRefreshTask                      : Refresh the releases for all workspaces
	revisionQueueTask                        : Pools the revision queue and looks for workspaces with revision to process.
	unhaltAllWorkspacesInDatabaseTask        : Unhalts all workspaces in the specified database partition.

# Getting Workspace Information

## Getting the count of revisions in the queue for a workspace

	> $ ./a2m help revision-count

	NAME:
	   revision-count - Show the number of revisions in a workspace's queue

	USAGE:
	   command revision-count [command options] [arguments...]

	OPTIONS:
	   --workspace, -w "0"	The workspace Id

	> $ ./a2m revision-count -w 41529001
	80444


## Querying for the halted status of a workspace

	> $ ./a2m halted -w 41529001
	true

## Viewing Halted Workspaces

	> $ ./a2m halted-workspaces

	Halted Workspaces
	=============================================================================
	41529001        100             false 2015-02-14T15:32:28Z manually halted

## Halting a Workspace

	> $ ./a2m halt -w 41529001
	Halted workspace 41529001
	> $ echo $?
	0

	> $ ./a2m halt -w 41529001
	Workspace 41529001 is already halted
	> $ echo $?
	1

## Unhalt a Workspace

	> $ ./a2m unhalt -w 41529001
	Workspace 41529001 unhalted.
	> $ echo $?
	0

	> $ ./a2m unhalt -w 41529001
	workspace '41529001' not in halted state
	> $ echo $?
	1

