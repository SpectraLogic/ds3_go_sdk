#
# Run this script in Git Bash on Windows to ensure file names do not
# exceed 59 characters (because Go is apparently using an old system call
# to open files.
#
find . -name *.go |
    while read F ; do
       LEN=`basename $F | wc -c`
        if test $LEN -gt 59 ; then
           NEWF=`echo $F | sed -e 's/NotificationRegistration/NR/' -e 's/SpectraS3/SS3/' -e 's/Target/Tg/' -e 's/Object/Ob/'`
          LEN=`basename $NEWF | wc -c`
           if test $LEN -gt 59 ; then
             echo $F will still be too long, update this script.
          else
              if [ "$NEWF" != "$F" ] ; then
                 mv $F $NEWF
              fi
          fi
       fi
    done

