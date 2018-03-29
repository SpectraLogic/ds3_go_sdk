#
# Run this script in Git Bash on Windows to ensure file names do not
# exceed 59 characters (because Go is apparently using an old system call
# to open files.
#
COUNT=0
find . -name *.go | grep ds3/models/ |
    while read F ; do
       NEWF=`dirname $F`/ds3l.${COUNT}.go
       mv $F $NEWF
       COUNT=`expr ${COUNT} + 1`
    done

