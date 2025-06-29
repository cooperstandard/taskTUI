#! /bin/bash
case $1 in 
   "start")
      go build cmd/main.go
      ./main
     ;;
   esac
