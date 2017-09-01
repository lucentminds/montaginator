@echo off

rem 32 bit setup.
set GOROOT=C:\go\1.9\386

rem 64 bit setup.
rem set GOROOT=C:\go\1.9\amd64

set GOPATH=%~dp0
set PATH=%PATH%;%GOROOT%\bin