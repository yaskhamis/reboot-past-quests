# ls -mtF | tr -d ' '
#find -name "*.sh"

# or you can use this as well

#ls *.sh

# helped me find the files, but now how to replace the .sh?

find . -type f -name "*.sh" -exec basename {} .sh \; | sort -r