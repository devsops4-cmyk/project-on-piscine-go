#!/bin/bash

# lookagain for the name of the file
find . -type f -name "*.sh" -exec basename {} \; | cut -d'.' -f1 | sort -r
