#!/bin/bash
#first step
mkdir 0

#step 2 set permission

chmod 401 0

#step 3 modifid the time

touch -d "1986-01-05 00:00 UTC" 0

#loop 1 file
# create an empty file
touch 1

# set the size to 1 byte
truncate -s 1 1

#set permission to octal 604
chmod 402 1

#set the modification time
touch -d "1986-11-13 00:01 UTC" 1


#loop 2 file
touch 2
truncate -s 2 2
# set permission to octal 604
chmod 604 2
touch -d "1988-03-05 00:10 UTC" 2

#loop 3
# create a symbolic link
ln -s 0 3

# set the links timestamp
touch -h -d "1990-02-16 00:11 UTC" 3

#loop 4
touch 4
truncate -s 4 4
# set permission to octal 510
chmod 510 4
touch -d "1990-10-07 01:00 UTC" 4

#loop 5
touch 5
truncate -s 5 5
# set permission to octal 510 4
chmod 510 4
touch -d "1990-11-07 01:01 UTC" 5

# loop 6
touch 6
truncate -s 6 6
# set permission to octal 460
chmod 460 6
touch -d "1991-02-08 01:10 UTC" 6

#loop 7
touch 7
truncate -s 7 7
#set permission to octal 510
chmod 510 7
touch -d "1991-03-08 01:11 UTC" 7

#loop 8
touch 8
truncate -s 8 8
# set permission to octal 607
chmod 604 8
touch -d "1994-05-20 10:00 UTC" 8

#loop 9
touch 9
truncate -s 9 9
# set permision to octal 604 9
chmod 402 9
touch -d "1994-06-10 10:01 UTC" 9

mkdir A
# set permission to octal 401
chmod 401 A
touch -d "1995-04-10 10:10 UTC" A




# 2. '!done.tar': An exception rule meaning "Do NOT ignore done.tar".

echo -e "*\n!done.tar" > .gitignore

# Add the .gitignore file
git add .gitignore

git add done.tar





