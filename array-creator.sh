#!/bin/bash

# Define the search phrase
search_phrase="trex"

# Find all files in the repository and search for the phrase
files_with_phrase=$(grep -rl "$search_phrase" .)

# Convert the list of files into an array
IFS=$'\n' read -rd '' -a file_array <<<"$files_with_phrase"

# Print the array
echo "Files containing the phrase \"$search_phrase\":"
for file in "${file_array[@]}"; do
    echo "$file"
done