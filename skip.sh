ls -l | awk 'NR % 2 == 0'

## or maybe you can use this code as well: ls -l | sed -n '2~2p'