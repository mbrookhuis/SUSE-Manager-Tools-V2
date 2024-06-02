# SUSE-Manager-Tools version 2

**********************************************

The python version of this project has been moved to https://github.com/uyuni-project/contrib and you can the tools can be found in uyuni-tools

***********************************************

This repository contains the golang version of SUSE Manager Tools. This version is writen against GoLang 1.21.

The goal was to keep the same parameters and the same configuration file as with the python version. But there are some exceptions:
* For the python version there is a different loglevel for what is written to the file and what has been shown on screen. For the golang version the loglevel defined for file will be used for both. Separate log-levels are not possible anymore.
* Currently there will be one central logfile. The server that is being updated can be easily recognized in the log.