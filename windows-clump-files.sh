#
# Run this script in Git Bash on Windows to ensure file names do not
# exceed 59 characters (because Go is apparently using an old system call
# to open files.
#
if test ! -d ds3/models ; then
   echo Run this from the root directory.
   exit 1
fi

(
echo
echo
echo 'package models'
echo
echo 'import ('

(find ds3/models -name *.go | 
    while read F ; do
       if grep -q '^import[^"]*".*"$' $F ; then
           grep '^import[^"]*".*"$' $F | sed 's/import[^"]*"/"/'
       else
           sed -e '1,/^import *(/d' -e '/)/,$d' $F
       fi
     done) | sed 's/ //g' | sort | uniq

echo ')'


find ds3/models -name *.go |
    while read F ; do
       if grep -q '^import.*".*"$' $F ; then
           sed -e '1,/^import.*".*"$/d' $F
       else
           sed -e '1,/i^ *)$/d' $F
       fi
       mv $F $F.orig
    done
) > ds3/clump.go
mv ds3/clump.go ds3/models

echo "Files merged to single ds3/models/clump.go"

